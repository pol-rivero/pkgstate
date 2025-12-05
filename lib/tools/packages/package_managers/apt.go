package packagemanagers

import (
	"strings"

	"github.com/pol-rivero/pkgstate/lib/common"
)

type Apt struct{}

func (a *Apt) GetBinaryName() string {
	return "apt"
}

func (a *Apt) GetAllInstalledPackages() ([]string, error) {
	output, err := common.RunCommandGetOutput("apt", "list", "--installed")
	if err != nil {
		return nil, err
	}
	return a.parseAptListOutput(output), nil
}

func (a *Apt) GetExplicitlyInstalledPackages() ([]string, error) {
	return common.RunCommandGetLines("apt-mark", "showmanual")
}

func (a *Apt) RemovePackages(packages []string) error {
	// Same pattern as pacman: first mark the packages as dependencies,
	// and then remove only those that became orphans.
	if err := a.markPackagesAsDependency(packages); err != nil {
		return err
	}
	return a.cleanUnusedDependencies(packages)
}

func (a *Apt) MarkPackagesAsExplicitlyInstalled(packages []string) error {
	args := append([]string{"sudo", "apt-mark", "manual"}, packages...)
	return common.RunCommand(args...)
}

func (a *Apt) InstallPackages(packages []string) error {
	args := append([]string{"sudo", "apt-get", "install", "-y", "--no-install-recommends"}, packages...)
	return common.RunCommand(args...)
}

func (a *Apt) markPackagesAsDependency(packages []string) error {
	args := append([]string{"sudo", "apt-mark", "auto"}, packages...)
	return common.RunCommand(args...)
}

func (a *Apt) cleanUnusedDependencies(packagesAllowedToBeRemoved []string) error {
	// First, get the list of packages that would be autoremoved
	output, err := common.RunCommandGetOutput("apt-get", "--dry-run", "autoremove",
		"-o", "APT::AutoRemove::RecommendsImportant=0",
		"-o", "APT::AutoRemove::SuggestsImportant=0")
	if err != nil {
		return err
	}
	orphanedPackages := a.parseAutoremoveOutput(output)
	packagesToRemove := common.IntersectionOfOrderedSlices(packagesAllowedToBeRemoved, common.Sorted(orphanedPackages))
	if len(packagesToRemove) == 0 {
		return nil
	}
	args := append([]string{"sudo", "apt-get", "autoremove", "-y",
		"-o", "APT::AutoRemove::RecommendsImportant=0",
		"-o", "APT::AutoRemove::SuggestsImportant=0"}, packagesToRemove...)
	return common.RunCommand(args...)
}

// parseAptListOutput parses the output of 'apt list --installed' command.
// The output format is: "package/source version [arch]" or "package/source version arch [upgradable]"
// We only need the package name (before the '/').
func (a *Apt) parseAptListOutput(output string) []string {
	lines := strings.Split(output, "\n")
	packages := []string{}
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || line == "Listing..." {
			continue
		}
		// Extract package name (everything before the '/')
		if idx := strings.Index(line, "/"); idx > 0 {
			packages = append(packages, line[:idx])
		}
	}
	return packages
}

// parseAutoremoveOutput parses the output of 'apt-get --dry-run autoremove' command.
// It extracts the list of packages that would be removed.
func (a *Apt) parseAutoremoveOutput(output string) []string {
	lines := strings.Split(output, "\n")
	packages := []string{}
	inRemoveSection := false
	for _, line := range lines {
		if strings.HasPrefix(line, "Remv ") {
			// Line format: "Remv package [version]" or "Remv package (version)"
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				packages = append(packages, parts[1])
			}
			inRemoveSection = true
		} else if inRemoveSection && !strings.HasPrefix(line, "Remv ") && strings.TrimSpace(line) != "" {
			// We've moved past the removal section
			inRemoveSection = false
		}
	}
	return packages
}

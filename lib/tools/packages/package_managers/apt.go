package packagemanagers

import (
	"strings"

	"github.com/pol-rivero/pkgstate/lib/common"
)

type Apt struct{}

func (a *Apt) GetBinaryName() string {
	return "apt-get"
}

func (a *Apt) GetAllInstalledPackages() ([]string, error) {
	return common.RunCommandGetLines("dpkg-query", "-f", "${Package}\n", "-W")
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
	lines, err := common.RunCommandGetLines("apt-get", "--dry-run", "autoremove",
		"-o", "APT::AutoRemove::RecommendsImportant=0",
		"-o", "APT::AutoRemove::SuggestsImportant=0")
	if err != nil {
		return err
	}
	orphanedPackages := a.parseAutoremoveOutput(lines)
	packagesToRemove := common.IntersectionOfOrderedSlices(packagesAllowedToBeRemoved, common.Sorted(orphanedPackages))
	if len(packagesToRemove) == 0 {
		return nil
	}
	args := append([]string{"sudo", "apt-get", "remove", "-y"}, packagesToRemove...)
	return common.RunCommand(args...)
}

func (a *Apt) parseAutoremoveOutput(lines []string) []string {
	packages := make([]string, 0, len(lines))
	for _, line := range lines {
		// Line format: "Remv package [version]" or "Remv package (version)"
		if strings.HasPrefix(line, "Remv ") {
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				packages = append(packages, parts[1])
			}
		}
	}
	return packages
}

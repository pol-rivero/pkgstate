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
	lines, err := common.RunCommandGetLines("dpkg-query", "-f", "${db:Status-Abbrev} ${Package}\n", "-W")
	if err != nil {
		return nil, err
	}
	return a.parseDpkgQueryOutput(lines), nil
}

func (a *Apt) GetExplicitlyInstalledPackages() ([]string, error) {
	return common.RunCommandGetLines("apt-mark", "showmanual")
}

func (a *Apt) GetOrphanedPackages() ([]string, error) {
	lines, err := common.RunCommandGetLines("apt-get", "--dry-run", "autoremove",
		"-o", "APT::AutoRemove::RecommendsImportant=0",
		"-o", "APT::AutoRemove::SuggestsImportant=0")
	if err != nil {
		return nil, err
	}
	return a.parseAutoremoveOutput(lines), nil
}

func (a *Apt) MarkPackagesAsExplicitlyInstalled(packages []string) error {
	args := append([]string{"sudo", "apt-mark", "manual"}, packages...)
	return common.RunCommand(args...)
}

func (a *Apt) MarkPackagesAsDependency(packages []string) error {
	args := append([]string{"sudo", "apt-mark", "auto"}, packages...)
	return common.RunCommand(args...)
}

func (a *Apt) InstallPackages(packages []string) error {
	args := append([]string{"sudo", "apt-get", "install", "-y", "--no-install-recommends"}, packages...)
	return common.RunCommand(args...)
}

func (a *Apt) RemovePackages(packages []string) error {
	args := append([]string{"sudo", "apt-get", "remove", "--auto-remove", "-y"}, packages...)
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

func (a *Apt) parseDpkgQueryOutput(lines []string) []string {
	packages := make([]string, 0, len(lines))
	for _, line := range lines {
		// Line format: "Status Package"
		parts := strings.Fields(line)
		if len(parts) >= 2 {
			status := parts[0]
			pkg := parts[1]
			if status == "ii" {
				packages = append(packages, pkg)
			}
		}
	}
	return packages
}

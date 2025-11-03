package packagemanagers

import "github.com/pol-rivero/pkgstate/lib/common"

type Pacman struct{}

func (p *Pacman) GetBinaryName() string {
	return "pacman"
}

func (p *Pacman) GetAllInstalledPackages() ([]string, error) {
	return common.RunCommandGetLines("pacman", "-Qq")
}

func (p *Pacman) GetExplicitlyInstalledPackages() ([]string, error) {
	return common.RunCommandGetLines("pacman", "-Qqe")
}

func (p *Pacman) RemovePackages(packages []string) error {
	// Naively running 'pacman -R' can fail if some packages are a dependency of other (non-removed) packages.
	// To avoid that, first mark the packages as dependencies, and then remove only those that became orphans.
	if err := p.markPackagesAsDependency(packages); err != nil {
		return err
	}
	return p.cleanUnusedDependencies(packages)
}

func (p *Pacman) MarkPackagesAsExplicitlyInstalled(packages []string) error {
	args := append([]string{"sudo", "pacman", "-D", "--asexplicit"}, packages...)
	return common.RunCommand(args...)
}

func (p *Pacman) InstallPackages(packages []string) error {
	// '--ask 4' to not prompt for confirmation when the package to be installed conflicts with
	// an already installed package (automatically answer "yes" to uninstall the old package)
	args := append([]string{"sudo", "pacman", "-S", "--ask", "4"}, packages...)
	return common.RunCommand(args...)
}

func (p *Pacman) markPackagesAsDependency(packages []string) error {
	args := append([]string{"sudo", "pacman", "-D", "--asdep"}, packages...)
	return common.RunCommand(args...)
}

func (p *Pacman) cleanUnusedDependencies(packagesAllowedToBeRemoved []string) error {
	// 'pacman -Qdtq' lists all the orphaned packages, UNLESS they are an optional dependency of another package.
	// Pass a second t to include those optional dependencies as well.
	orphanedPackages, err := common.RunCommandGetLines("pacman", "-Qdttq")
	if err != nil {
		return err
	}
	packagesToRemove := common.IntersectionOfOrderedSlices(packagesAllowedToBeRemoved, common.Sorted(orphanedPackages))
	if len(packagesToRemove) == 0 {
		return nil
	}
	args := append([]string{"sudo", "pacman", "-Rs", "--noconfirm"}, packagesToRemove...)
	return common.RunCommand(args...)
}

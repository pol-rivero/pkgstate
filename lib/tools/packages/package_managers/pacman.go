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

func (p *Pacman) GetOrphanedPackages() ([]string, error) {
	// 'pacman -Qdtq' lists all the orphaned packages, UNLESS they are an optional dependency of another package.
	// Pass a second t to include those optional dependencies as well.
	return common.RunCommandGetLines("pacman", "-Qdttq")
}

func (p *Pacman) MarkPackagesAsExplicitlyInstalled(packages []string) error {
	args := append([]string{"sudo", "pacman", "-D", "--asexplicit"}, packages...)
	return common.RunCommand(args...)
}

func (p *Pacman) MarkPackagesAsDependency(packages []string) error {
	args := append([]string{"sudo", "pacman", "-D", "--asdep"}, packages...)
	return common.RunCommand(args...)
}

func (p *Pacman) InstallPackages(packages []string) error {
	// '--ask 4' to not prompt for confirmation when the package to be installed conflicts with
	// an already installed package (automatically answer "yes" to uninstall the old package)
	args := append([]string{"sudo", "pacman", "-S", "--ask", "4"}, packages...)
	return common.RunCommand(args...)
}

func (p *Pacman) RemovePackages(packages []string) error {
	args := append([]string{"sudo", "pacman", "-Rs", "--noconfirm"}, packages...)
	return common.RunCommand(args...)
}

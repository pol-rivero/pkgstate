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

func (y *Pacman) RemovePackages(packages []string) error {
	args := append([]string{"sudo", "pacman", "-Rs", "--noconfirm"}, packages...)
	err := common.RunCommand(args...)
	return err
}

func (y *Pacman) MarkPackagesAsExplicitlyInstalled(packages []string) error {
	args := append([]string{"sudo", "pacman", "-D", "--asexplicit"}, packages...)
	err := common.RunCommand(args...)
	return err
}

func (y *Pacman) InstallPackages(packages []string) error {
	args := append([]string{"sudo", "pacman", "-S", "--ask", "4"}, packages...)
	err := common.RunCommand(args...)
	return err
}

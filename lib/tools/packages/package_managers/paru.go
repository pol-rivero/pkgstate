package packagemanagers

import "github.com/pol-rivero/pkgstate/lib/common"

type Paru struct {
	// paru lists 'pacman' as a dependency, so we can assume it is present.
	// pacman has no problem listing and removing packages, even if they are from AUR.
	Parent Pacman
}

func (y *Paru) GetBinaryName() string {
	return "paru"
}

func (y *Paru) GetAllInstalledPackages() ([]string, error) {
	return y.Parent.GetAllInstalledPackages()
}

func (y *Paru) GetExplicitlyInstalledPackages() ([]string, error) {
	return y.Parent.GetExplicitlyInstalledPackages()
}

func (y *Paru) RemovePackages(packages []string) error {
	return y.Parent.RemovePackages(packages)
}

func (y *Paru) MarkPackagesAsExplicitlyInstalled(packages []string) error {
	return y.Parent.MarkPackagesAsExplicitlyInstalled(packages)
}

func (y *Paru) InstallPackages(packages []string) error {
	args := append([]string{"paru", "-S"}, packages...)
	_, err := common.RunCommandGetLines(args...)
	return err
}

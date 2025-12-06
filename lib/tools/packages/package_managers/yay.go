package packagemanagers

import "github.com/pol-rivero/pkgstate/lib/common"

type Yay struct {
	// yay lists 'pacman' as a dependency, so we can assume it is present.
	// pacman has no problem listing and removing packages, even if they are from AUR.
	Parent Pacman
}

func (y *Yay) GetBinaryName() string {
	return "yay"
}

func (y *Yay) InstallPackages(packages []string) error {
	args := append([]string{"yay", "-S", "--ask", "4"}, packages...)
	err := common.RunCommand(args...)
	return err
}

// Delegate to parent

func (y *Yay) GetAllInstalledPackages() ([]string, error) {
	return y.Parent.GetAllInstalledPackages()
}

func (y *Yay) GetExplicitlyInstalledPackages() ([]string, error) {
	return y.Parent.GetExplicitlyInstalledPackages()
}

func (y *Yay) GetOrphanedPackages() ([]string, error) {
	return y.Parent.GetOrphanedPackages()
}

func (y *Yay) MarkPackagesAsExplicitlyInstalled(packages []string) error {
	return y.Parent.MarkPackagesAsExplicitlyInstalled(packages)
}

func (y *Yay) MarkPackagesAsDependency(packages []string) error {
	return y.Parent.MarkPackagesAsDependency(packages)
}

func (y *Yay) RemovePackages(packages []string) error {
	return y.Parent.RemovePackages(packages)
}

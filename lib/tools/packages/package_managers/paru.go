package packagemanagers

import "github.com/pol-rivero/pkgstate/lib/common"

type Paru struct {
	// paru lists 'pacman' as a dependency, so we can assume it is present.
	// pacman has no problem listing and removing packages, even if they are from AUR.
	Parent Pacman
}

func (p *Paru) GetBinaryName() string {
	return "paru"
}

func (p *Paru) InstallPackages(packages []string) error {
	args := append([]string{"paru", "-S", "--ask", "4"}, packages...)
	err := common.RunCommand(args...)
	return err
}

// Delegate to parent

func (p *Paru) GetAllInstalledPackages() ([]string, error) {
	return p.Parent.GetAllInstalledPackages()
}

func (p *Paru) GetExplicitlyInstalledPackages() ([]string, error) {
	return p.Parent.GetExplicitlyInstalledPackages()
}

func (p *Paru) GetOrphanedPackages() ([]string, error) {
	return p.Parent.GetOrphanedPackages()
}

func (p *Paru) MarkPackagesAsExplicitlyInstalled(packages []string) error {
	return p.Parent.MarkPackagesAsExplicitlyInstalled(packages)
}

func (p *Paru) MarkPackagesAsDependency(packages []string) error {
	return p.Parent.MarkPackagesAsDependency(packages)
}

func (p *Paru) RemovePackages(packages []string) error {
	return p.Parent.RemovePackages(packages)
}

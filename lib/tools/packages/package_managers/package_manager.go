package packagemanagers

import (
	"github.com/pol-rivero/pkgstate/lib/common"
	"github.com/pol-rivero/pkgstate/lib/common/log"
)

var PACKAGE_MANAGERS = []PackageManager{
	&Yay{},
	&Paru{},
	&Pacman{},
}

type PackageManager interface {
	GetBinaryName() string
	GetAllInstalledPackages() ([]string, error)
	GetExplicitlyInstalledPackages() ([]string, error)
	InstallPackages(packages []string) error
	RemovePackages(packages []string) error
	MarkPackagesAsExplicitlyInstalled(packages []string) error
}

func FindPreferredPackageManager() PackageManager {
	for _, pm := range PACKAGE_MANAGERS {
		if common.IsCommandAvailable(pm.GetBinaryName()) {
			return pm
		}
	}
	log.Fatal("No suitable package manager found")
	panic("Unreachable")
}

package packagemanagers

import "github.com/pol-rivero/pkgstate/lib/common/log"

var PACKAGE_MANAGERS = []PackageManager{
	&Yay{},
}

type PackageManager interface {
	GetAllInstalledPackages() ([]string, error)
	GetExplicitlyInstalledPackages() ([]string, error)
}

func FindPreferredPackageManager() PackageManager {
	// TODO: find other package managers and fallback to pacman
	for _, pm := range PACKAGE_MANAGERS {
		if true {
			return pm
		}
	}
	log.Fatal("No suitable package manager found")
	panic("Unreachable")
}

package packages

import (
	pm "github.com/pol-rivero/pkgstate/lib/tools/packages/package_managers"
)

type PackagesTool struct {
	PackageManagerCache         pm.PackageManager
	DesiredPackages             []string
	AllInstalledPackages        []string
	ExplicitlyInstalledPackages []string
}

func NewPackagesTool() *PackagesTool {
	return &PackagesTool{}
}

func (l *PackagesTool) FriendlyProcessName() string {
	return "get installed packages"
}

func (l *PackagesTool) getPackageManager() pm.PackageManager {
	if l.PackageManagerCache == nil {
		l.PackageManagerCache = pm.FindPreferredPackageManager()
	}
	return l.PackageManagerCache
}

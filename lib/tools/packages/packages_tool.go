package packages

import (
	pm "github.com/pol-rivero/pkgstate/lib/tools/packages/package_managers"
)

type PackagesTool struct {
	PackageManagerCache         pm.PackageManager
	AllInstalledPackages        []string
	ExplicitlyInstalledPackages []string
}

func (l *PackagesTool) getPackageManager() pm.PackageManager {
	if l.PackageManagerCache == nil {
		l.PackageManagerCache = pm.FindPreferredPackageManager()
	}
	return l.PackageManagerCache
}

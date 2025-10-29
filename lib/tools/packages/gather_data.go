package packages

import (
	"fmt"

	"github.com/pol-rivero/pkgstate/lib/common"
	"github.com/pol-rivero/pkgstate/lib/common/config"
)

func (l *PackagesTool) GatherData(*config.Config) error {
	packageManager := l.getPackageManager()

	// TODO: parallelize data gathering
	allPackages, err := packageManager.GetAllInstalledPackages()
	if err != nil {
		return fmt.Errorf("error getting all installed packages: %v", err)
	}
	l.AllInstalledPackages = common.Sorted(allPackages)

	explicitlyInstalled, err := packageManager.GetExplicitlyInstalledPackages()
	if err != nil {
		return fmt.Errorf("error getting explicitly installed packages: %v", err)
	}
	l.ExplicitlyInstalledPackages = common.Sorted(explicitlyInstalled)

	return nil
}

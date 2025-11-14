package packages

import (
	"fmt"

	"github.com/pol-rivero/pkgstate/lib/common/config"
)

func (l *PackagesTool) GenerateConfig(cfg *config.Config) error {
	packageManager := l.getPackageManager()

	explicitlyInstalled, err := packageManager.GetExplicitlyInstalledPackages()
	if err != nil {
		return fmt.Errorf("error getting explicitly installed packages: %v", err)
	}

	cfg.Packages = explicitlyInstalled
	return nil
}

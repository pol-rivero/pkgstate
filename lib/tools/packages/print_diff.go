package packages

import (
	"fmt"

	"github.com/pol-rivero/pkgstate/lib/common"
	"github.com/pol-rivero/pkgstate/lib/common/config"
)

const (
	GREEN  = "\033[92m"
	YELLOW = "\033[93m"
	RESET  = "\033[0m"
	BOLD   = "\033[1m"
)

func (l *PackagesTool) PrintDiff(config *config.Config) error {
	desiredPackages := common.Sorted(config.Packages)
	toInstall := common.DifferenceOfOrderedSlices(desiredPackages, l.AllInstalledPackages)
	toRemove := common.DifferenceOfOrderedSlices(l.ExplicitlyInstalledPackages, desiredPackages)

	if len(toInstall) == 0 && len(toRemove) == 0 {
		return nil
	}
	printPackageList(GREEN, "Packages to install", toInstall)
	printPackageList(YELLOW, "Unmanaged packages", toRemove)
	return nil
}

func printPackageList(color string, title string, packages []string) {
	if len(packages) == 0 {
		return
	}
	fmt.Printf("%s%s: %s%d%s -> ", color, title, BOLD, len(packages), RESET)
	for i, pkg := range packages {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Printf("%s", pkg)
	}
	fmt.Println()
}

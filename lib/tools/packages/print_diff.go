package packages

import (
	"fmt"

	"github.com/pol-rivero/pkgstate/lib/common"
	"github.com/pol-rivero/pkgstate/lib/common/colors"
)

func (l *PackagesTool) PrintDiff() {
	toInstall := common.DifferenceOfOrderedSlices(l.DesiredPackages, l.AllInstalledPackages)
	toRemove := common.DifferenceOfOrderedSlices(l.ExplicitlyInstalledPackages, l.DesiredPackages)

	if len(toInstall) == 0 && len(toRemove) == 0 {
		return
	}
	printPackageList(colors.GREEN, "Packages to install", toInstall)
	printPackageList(colors.YELLOW, "Unmanaged packages", toRemove)
}

func printPackageList(color string, title string, packages []string) {
	if len(packages) == 0 {
		return
	}
	fmt.Printf("%s%s: %s%d%s\n", color, title, colors.BOLD, len(packages), colors.RESET)
	for _, pkg := range packages {
		fmt.Printf("  - %s\n", pkg)
	}
}

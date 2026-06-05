package packages

import (
	"strings"

	"github.com/pol-rivero/pkgstate/lib/common"
	"github.com/pol-rivero/pkgstate/lib/common/log"
	"github.com/pol-rivero/pkgstate/lib/common/prompt"
	pm "github.com/pol-rivero/pkgstate/lib/tools/packages/package_managers"
	. "github.com/pol-rivero/pkgstate/lib/types"
)

func (l *PackagesTool) ApplyFixes(requestConfirmation bool, noRemove bool) ApplyFixesResult {
	toInstall := common.DifferenceOfOrderedSlices(l.DesiredPackages, l.AllInstalledPackages)
	toRemove := common.DifferenceOfOrderedSlices(l.ExplicitlyInstalledPackages, l.DesiredPackages)
	toMarkAsExplicit := l.installedAndDesiredAndNotAlreadyExplicit()

	packageManager := l.getPackageManager()
	if len(toInstall) > 0 {
		if ynPrompt(requestConfirmation, "Do you want to install the following packages?\n%s", toInstall) {
			checkErr(packageManager.InstallPackages(toInstall), "install packages")
			return ProcessAgain
		} else {
			log.Info("Skipping installation of packages.")
		}
	}
	if len(toMarkAsExplicit) > 0 {
		if ynPrompt(requestConfirmation, "Do you want to mark the following packages as explicitly installed?\n%s", toMarkAsExplicit) {
			checkErr(packageManager.MarkPackagesAsExplicitlyInstalled(toMarkAsExplicit), "mark packages as explicitly installed")
			return ProcessAgain
		} else {
			log.Info("Skipping marking packages as explicitly installed.")
		}
	}
	if len(toRemove) > 0 {
		if noRemove {
			log.Info("Skipping removal of packages (--no-remove).")
		} else if ynPrompt(requestConfirmation, "Do you want to remove the following packages?\n%s", toRemove) {
			checkErr(removePackages(packageManager, toRemove), "remove packages")
		} else {
			log.Info("Skipping removal of packages.")
		}
	}
	return Done
}

func removePackages(packageManager pm.PackageManager, packagesToPotentiallyRemove []string) error {
	// Naively calling RemovePackages can fail if some packages are a dependency of other (non-removed) packages.
	// To avoid that, first mark the packages as dependencies, and then remove only those that became orphans.
	if err := packageManager.MarkPackagesAsDependency(packagesToPotentiallyRemove); err != nil {
		return err
	}
	orphanedPackages, err := packageManager.GetOrphanedPackages()
	if err != nil {
		return err
	}
	packagesToRemove := common.IntersectionOfOrderedSlices(packagesToPotentiallyRemove, common.Sorted(orphanedPackages))
	if len(packagesToRemove) == 0 {
		return nil
	}
	return packageManager.RemovePackages(packagesToRemove)
}

func ynPrompt(requestConfirmation bool, message string, packages []string) bool {
	if requestConfirmation {
		return prompt.RequestInput("Yn", message, formatList(packages)) == 'y'
	}
	return true
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Error("Failed to %s: %v", msg, err)
	}
}

func (l *PackagesTool) installedAndDesiredAndNotAlreadyExplicit() []string {
	desiredAndNotAlreadyExplicit := common.DifferenceOfOrderedSlices(l.DesiredPackages, l.ExplicitlyInstalledPackages)
	return common.IntersectionOfOrderedSlices(desiredAndNotAlreadyExplicit, l.AllInstalledPackages)
}

func formatList(packages []string) string {
	var builder strings.Builder
	for _, pkg := range packages {
		builder.WriteString("  - ")
		builder.WriteString(pkg)
		builder.WriteString("\n")
	}
	return builder.String()
}

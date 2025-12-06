package packages

import (
	"github.com/pol-rivero/pkgstate/lib/common/log"
	"github.com/pol-rivero/pkgstate/lib/common/prompt"
)

func (l *PackagesTool) Cleanup(requestConfirmation bool) {
	packageManager := l.getPackageManager()
	orphanedPackages, err := packageManager.GetOrphanedPackages()
	if err != nil {
		log.Error("Failed to get orphaned packages: %v", err)
		return
	}
	if len(orphanedPackages) > 0 {
		prompt := "This will make your system fully declarative but might break things. Carefully review this list and " +
			"modify the configuration file to keep the packages you still need.\nDo you want to remove the following packages?\n%s"
		if ynPromptDefaultNo(requestConfirmation, prompt, orphanedPackages) {
			checkErr(removePackages(packageManager, orphanedPackages), "remove packages")
		} else {
			log.Info("Skipping removal of packages.")
		}
	}
}

func ynPromptDefaultNo(requestConfirmation bool, message string, packages []string) bool {
	if requestConfirmation {
		return prompt.RequestInput("yN", message, formatList(packages)) == 'y'
	}
	return true
}

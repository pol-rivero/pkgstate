package fix

import (
	"github.com/pol-rivero/pkgstate/lib/common/config"
	"github.com/pol-rivero/pkgstate/lib/common/log"
	"github.com/pol-rivero/pkgstate/lib/tools"
)

func Fix(noConfirm bool) {
	config := config.GetConfig()
	toolList := tools.CreateTools()
	for _, tool := range toolList {
		// Fixes from one tool may affect others, so don't gather data in parallel
		err := tool.GatherData(&config)
		if err != nil {
			log.Fatal("Failed to %s: %v", tool.FriendlyProcessName(), err)
		}
		requestConfirmation := !noConfirm
		tool.ApplyFixes(requestConfirmation)
	}
}

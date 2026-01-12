package commands

import (
	"github.com/pol-rivero/pkgstate/lib/common/config"
	"github.com/pol-rivero/pkgstate/lib/common/log"
	"github.com/pol-rivero/pkgstate/lib/tools"
	. "github.com/pol-rivero/pkgstate/lib/types"
)

func Fix(noConfirm bool) {
	config := config.GetConfig()
	toolList := tools.CreateTools()
	for _, tool := range toolList {
		// Fixes from one tool may affect others, so don't gather data in parallel
		applyFixTool(tool, config, noConfirm)
	}
}

func applyFixTool(tool tools.Tool, config config.Config, noConfirm bool) {
	err := tool.GatherData(&config)
	if err != nil {
		log.Error("Failed to %s (skipping): %v", tool.FriendlyProcessName(), err)
		return
	}
	requestConfirmation := !noConfirm
	result := tool.ApplyFixes(requestConfirmation)
	if result == ProcessAgain {
		log.Info("Refreshing '%s' and trying again", tool.FriendlyProcessName())
		applyFixTool(tool, config, noConfirm)
	}
}

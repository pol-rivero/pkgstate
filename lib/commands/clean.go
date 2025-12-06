package commands

import (
	"github.com/pol-rivero/pkgstate/lib/tools"
)

func Clean(noConfirm bool) {
	toolList := tools.CreateTools()
	requestConfirmation := !noConfirm
	for _, tool := range toolList {
		tool.Cleanup(requestConfirmation)
	}
}

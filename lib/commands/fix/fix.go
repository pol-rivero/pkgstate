package fix

import (
	"github.com/pol-rivero/pkgstate/lib/common/config"
	"github.com/pol-rivero/pkgstate/lib/tools"
)

func Fix(noConfirm bool) {
	config := config.GetConfig()
	tools := tools.GetTools(&config)
	for _, tool := range tools {
		tool.ApplyFixes(&config)
	}
}

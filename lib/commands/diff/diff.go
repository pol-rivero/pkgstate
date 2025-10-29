package diff

import (
	"github.com/pol-rivero/pkgstate/lib/common/config"
	"github.com/pol-rivero/pkgstate/lib/tools"
)

func Diff() {
	config := config.GetConfig()
	tools := tools.GetTools(&config)
	for _, tool := range tools {
		tool.PrintDiff(&config)
	}
}

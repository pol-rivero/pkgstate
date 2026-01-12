package tools

import (
	"github.com/pol-rivero/pkgstate/lib/common/config"
	"github.com/pol-rivero/pkgstate/lib/tools/groups"
	"github.com/pol-rivero/pkgstate/lib/tools/packages"
	"github.com/pol-rivero/pkgstate/lib/tools/systemd"
	. "github.com/pol-rivero/pkgstate/lib/types"
)

type Tool interface {
	FriendlyProcessName() string
	GatherData(*config.Config) error
	PrintDiff()
	ApplyFixes(requestConfirmation bool) ApplyFixesResult
	GenerateConfig(*config.Config) error
	Cleanup(requestConfirmation bool)
}

func CreateTools() []Tool {
	return []Tool{
		packages.NewPackagesTool(),
		systemd.NewSystemdTool(true),
		systemd.NewSystemdTool(false),
		groups.NewGroupsTool(),
	}
}

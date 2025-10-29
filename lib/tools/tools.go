package tools

import (
	"github.com/pol-rivero/pkgstate/lib/common/config"
	"github.com/pol-rivero/pkgstate/lib/tools/packages"
)

type Tool interface {
	GatherData(*config.Config) error
	PrintDiff(*config.Config) error
	ApplyFixes(*config.Config) error
}

func instantiateTools() []Tool {
	return []Tool{
		&packages.PackagesTool{},
	}
}

func GetTools(config *config.Config) []Tool {
	tools := instantiateTools()
	for _, tool := range tools {
		tool.GatherData(config)
	}
	return tools
}

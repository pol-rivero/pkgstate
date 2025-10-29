package tools

import (
	"sync"
	"time"

	"github.com/pol-rivero/pkgstate/lib/common/config"
	"github.com/pol-rivero/pkgstate/lib/common/log"
	"github.com/pol-rivero/pkgstate/lib/tools/packages"
	"github.com/pol-rivero/pkgstate/lib/tools/systemd"
)

type Tool interface {
	FriendlyProcessName() string
	GatherData(*config.Config) error
	PrintDiff()
	ApplyFixes(noConfirm bool)
}

func CreateTools() []Tool {
	return []Tool{
		packages.NewPackagesTool(),
		systemd.NewSystemdTool(true),
		systemd.NewSystemdTool(false),
	}
}

func GatherDataInParallel(tools []Tool, config *config.Config) {
	var waitGroup sync.WaitGroup
	initializeToolData := func(t Tool) {
		startTime := time.Now()
		err := t.GatherData(config)
		if err != nil {
			log.Fatal("Failed to %s: %v", t.FriendlyProcessName(), err)
		}
		log.Info("Time taken to %s: %s", t.FriendlyProcessName(), time.Since(startTime))
		waitGroup.Done()
	}
	for _, tool := range tools {
		waitGroup.Add(1)
		go initializeToolData(tool)
	}
	waitGroup.Wait()
}

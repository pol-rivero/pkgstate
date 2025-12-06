package diff

import (
	"sync"
	"time"

	"github.com/pol-rivero/pkgstate/lib/common/config"
	"github.com/pol-rivero/pkgstate/lib/common/log"
	"github.com/pol-rivero/pkgstate/lib/tools"
)

func Diff() {
	config := config.GetConfig()
	toolList := tools.CreateTools()

	gatherToolData := func(t tools.Tool, waitGroup *sync.WaitGroup) {
		startTime := time.Now()
		err := t.GatherData(&config)
		if err != nil {
			log.Error("Failed to %s (skipping): %v", t.FriendlyProcessName(), err)
		}
		log.Info("Time taken to %s: %s", t.FriendlyProcessName(), time.Since(startTime))
		waitGroup.Done()
	}

	waitGroups := make([]*sync.WaitGroup, len(toolList))
	for i, tool := range toolList {
		waitGroups[i] = &sync.WaitGroup{}
		waitGroups[i].Add(1)
		go gatherToolData(tool, waitGroups[i])
	}

	for i, tool := range toolList {
		waitGroups[i].Wait()
		tool.PrintDiff()
	}

}

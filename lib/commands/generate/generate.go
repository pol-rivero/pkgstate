package generate

import (
	"fmt"

	"github.com/pol-rivero/pkgstate/lib/common/config"
	"github.com/pol-rivero/pkgstate/lib/common/log"
	"github.com/pol-rivero/pkgstate/lib/tools"
)

func Generate() {
	cfg := config.Config{}

	toolList := tools.CreateTools()
	for _, tool := range toolList {
		err := tool.GenerateConfig(&cfg)
		if err != nil {
			log.Fatal("Failed to %s: %v", tool.FriendlyProcessName(), err)
		}
	}

	fmt.Printf("%v", cfg)
}

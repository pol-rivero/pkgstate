package commands

import (
	"fmt"

	"github.com/pol-rivero/pkgstate/lib/common/config"
)

func PrintConfig() {
	config := config.GetConfig()
	fmt.Printf("%v", config)
}

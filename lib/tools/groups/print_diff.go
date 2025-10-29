package groups

import (
	"fmt"

	"github.com/pol-rivero/pkgstate/lib/common/colors"
)

func (l *GroupsTool) PrintDiff() {
	if len(l.MissingGroups) == 0 {
		return
	}
	fmt.Printf("%sMissing User Groups: %s%d%s\n", colors.YELLOW, colors.BOLD, len(l.MissingGroups), colors.RESET)
	for _, group := range l.MissingGroups {
		fmt.Printf("  - %s\n", group.Name)
	}
}

package systemd

import (
	"fmt"

	"github.com/pol-rivero/pkgstate/lib/common/colors"
)

func (l *SystemdTool) PrintDiff() {
	if len(l.UnitMismatches) == 0 {
		return
	}
	fmt.Printf("%sSystemd Unit Mismatches: %s%d%s\n", colors.YELLOW, colors.BOLD, len(l.UnitMismatches), colors.RESET)
	for _, mismatch := range l.UnitMismatches {
		fmt.Printf("  - %s: %s (should be %s%s%s)\n", mismatch.UnitName, mismatch.CurrentState, colors.BOLD, mismatch.DesiredState, colors.RESET)
	}
}

package systemd

import (
	"fmt"

	"github.com/pol-rivero/pkgstate/lib/common"
	"github.com/pol-rivero/pkgstate/lib/common/log"
	"github.com/pol-rivero/pkgstate/lib/common/prompt"
	. "github.com/pol-rivero/pkgstate/lib/types"
)

func (l *SystemdTool) ApplyFixes(requestConfirmation bool, _ bool) ApplyFixesResult {
	sudo := getSudoPrefix(l.SystemScope)
	scope := getScopeFlag(l.SystemScope)
	for _, mismatch := range l.UnitMismatches {
		action, err := getActionForDesiredState(mismatch.DesiredState)
		if err != nil {
			log.Error("Failed to update unit '%s': %v", mismatch.UnitName, err)
			continue
		}

		if requestConfirmation {
			response := prompt.RequestInput("Yn", "Do you want to %s the unit '%s'?", action, mismatch.UnitName)
			if response != 'y' {
				log.Info("Skipping update of unit '%s'", mismatch.UnitName)
				continue
			}
		}

		commandAndArgs := append(sudo, "systemctl", scope, action, "--now", string(mismatch.UnitName))
		err = common.RunCommand(commandAndArgs...)
		if err != nil {
			log.Error("Failed to update unit '%s': %v", mismatch.UnitName, err)
			continue
		}
		fmt.Printf("-> Successfully updated unit '%s' to state '%s'\n", mismatch.UnitName, mismatch.DesiredState)
	}
	return Done
}

func getActionForDesiredState(desiredState SystemdUnitState) (string, error) {
	switch desiredState {
	case "enabled":
		return "enable", nil
	case "disabled":
		return "disable", nil
	case "masked":
		return "mask", nil
	default:
		return "", fmt.Errorf("the desired state '%s' cannot be applied automatically", desiredState)
	}
}

func getSudoPrefix(systemScope bool) []string {
	if systemScope {
		return []string{"sudo"}
	} else {
		return []string{}
	}
}

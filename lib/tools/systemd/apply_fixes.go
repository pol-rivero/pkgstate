package systemd

import (
	"fmt"

	"github.com/pol-rivero/pkgstate/lib/common"
	"github.com/pol-rivero/pkgstate/lib/common/log"
	"github.com/pol-rivero/pkgstate/lib/common/prompt"
)

func (l *SystemdTool) ApplyFixes(requestConfirmation bool) {
	sudo := getSudoPrefix(l.SystemScope)
	scope := getScopeFlag(l.SystemScope)
	for _, mismatch := range l.UnitMismatches {
		action, err := getActionForDesiredState(mismatch.DesiredState)
		if err != nil {
			log.Error("Failed to update unit '%s': %v", mismatch.UnitName, err)
			continue
		}

		if requestConfirmation {
			response := prompt.RequestInput("yN", "Do you want to %s the unit '%s'?", action, mismatch.UnitName)
			if response != 'y' {
				log.Info("Skipping update of unit '%s'", mismatch.UnitName)
				continue
			}
		}

		commandAndArgs := append(sudo, "systemctl", scope, action, "--now", string(mismatch.UnitName))
		_, err = common.RunCommand(commandAndArgs...)
		if err != nil {
			log.Error("Failed to update unit '%s': %v", mismatch.UnitName, err)
			continue
		}
		fmt.Printf("-> Successfully updated unit '%s' to state '%s'\n", mismatch.UnitName, mismatch.DesiredState)
	}
}

func getActionForDesiredState(desiredState SystemdUnitState) (string, error) {
	switch desiredState {
	case "enabled":
		return "enable", nil
	case "disabled":
		return "disable", nil
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

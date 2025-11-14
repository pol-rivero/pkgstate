package systemd

import (
	"github.com/pol-rivero/pkgstate/lib/common/config"
	. "github.com/pol-rivero/pkgstate/lib/types"
)

func (l *SystemdTool) GenerateConfig(cfg *config.Config) error {
	output, err := getCurrentUnits(l.SystemScope)
	if err != nil {
		return err
	}

	unitsMap := make(map[string]string)
	for _, unit := range output {
		state := NewLowercaseString(unit.State)
		// Only include units where the state differs from the default (preset)
		if (state == "enabled" || state == "disabled") && (unit.Preset == nil || state != NewLowercaseString(*unit.Preset)) {
			unitsMap[unit.UnitFile] = string(state)
		}
	}

	if l.SystemScope {
		cfg.SystemUnits = unitsMap
	} else {
		cfg.UserUnits = unitsMap
	}

	return nil
}

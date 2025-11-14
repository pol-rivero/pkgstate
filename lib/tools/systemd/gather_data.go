package systemd

import (
	"encoding/json"
	"fmt"

	"github.com/pol-rivero/pkgstate/lib/common"
	"github.com/pol-rivero/pkgstate/lib/common/config"
)

func (l *SystemdTool) GatherData(config *config.Config) error {
	desiredUnits := systemdUnitCollectionFromMap(getDesiredUnits(config, l.SystemScope))
	currentUnits, err := getCurrentUnits(l.SystemScope)
	if err != nil {
		return err
	}
	currentUnitsCollection := systemdUnitCollectionFromListUnitFilesOutput(currentUnits)
	l.UnitMismatches = getUnitMismatches(currentUnitsCollection, desiredUnits)
	return nil
}

func getCurrentUnits(systemScope bool) ([]SystemdListUnitFilesOutput, error) {
	jsonOutput, err := common.RunCommandGetOutput("systemctl", getScopeFlag(systemScope), "list-unit-files", "--all", "--no-pager", "--output=json")
	if err != nil {
		return nil, err
	}
	output := []SystemdListUnitFilesOutput{}
	if err := json.Unmarshal([]byte(jsonOutput), &output); err != nil {
		return nil, fmt.Errorf("error parsing systemctl list-unit-files output: %w", err)
	}
	return output, nil
}

func getScopeFlag(systemScope bool) string {
	if systemScope {
		return "--system"
	} else {
		return "--user"
	}
}

func getDesiredUnits(config *config.Config, systemScope bool) map[string]string {
	if systemScope {
		return config.SystemUnits
	} else {
		return config.UserUnits
	}
}

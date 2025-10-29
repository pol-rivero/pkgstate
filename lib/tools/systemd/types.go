package systemd

import (
	. "github.com/pol-rivero/pkgstate/lib/types"
)

type SystemdListUnitFilesOutput struct {
	UnitFile string `json:"unit_file"`
	State    string `json:"state"`
}

type SystemdUnitName string
type SystemdUnitState LowercaseString
type SystemdUnitCollection map[SystemdUnitName]SystemdUnitState

func systemdUnitCollectionFromMap(unitsMap map[string]string) SystemdUnitCollection {
	units := make(SystemdUnitCollection, len(unitsMap))
	for name, state := range unitsMap {
		stateLower := NewLowercaseString(state)
		units[SystemdUnitName(name)] = SystemdUnitState(stateLower)
	}
	return units
}

func systemdUnitCollectionFromListUnitFilesOutput(unitsList []SystemdListUnitFilesOutput) SystemdUnitCollection {
	units := make(SystemdUnitCollection, len(unitsList))
	for _, unit := range unitsList {
		stateLower := NewLowercaseString(unit.State)
		units[SystemdUnitName(unit.UnitFile)] = SystemdUnitState(stateLower)
	}
	return units
}

type SystemdUnitMismatch struct {
	UnitName     SystemdUnitName
	CurrentState SystemdUnitState
	DesiredState SystemdUnitState
}

func getUnitMismatches(currentUnits, desiredUnits SystemdUnitCollection) []SystemdUnitMismatch {
	var mismatches = make([]SystemdUnitMismatch, 0, len(desiredUnits))
	for unitName, currentState := range currentUnits {
		desiredState, exists := desiredUnits[unitName]
		if exists && currentState != desiredState {
			mismatches = append(mismatches, SystemdUnitMismatch{
				UnitName:     unitName,
				CurrentState: currentState,
				DesiredState: desiredState,
			})
		}
	}
	return mismatches
}

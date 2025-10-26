package config

import (
	"encoding/json"
	"fmt"

	"github.com/pol-rivero/pkgstate/lib/common/log"
)

func (config Config) String() string {
	formatString := "PACKAGES: %s\nUSER GROUPS: %s\nSYSTEMD UNITS (SYSTEM): %s\nSYSTEMD UNITS (USER): %s"
	return fmt.Sprintf(
		formatString,
		objectToString(config.Packages),
		objectToString(config.UserGroups),
		objectToString(config.SystemUnits),
		objectToString(config.UserUnits),
	)
}

func objectToString(obj any) string {
	bytes, err := json.Marshal(obj)
	if err != nil {
		log.Fatal("Error serializing object to string: %v", err)
	}
	return string(bytes)
}

package config

import (
	"maps"
	"os"
	"path"
	"strings"

	"github.com/pol-rivero/pkgstate/lib/common/log"
)

func GetConfig() Config {
	return fromDir(getConfigDir())
}

func fromDir(dir string) Config {
	config := defaultConfig()
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal("Error reading config directory (%s): %v", dir, err)
	}
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".yaml") {
			continue
		}
		fullName := dir + "/" + entry.Name()
		log.Info("Found config file: %s", fullName)
		entryConfig := fromFile(fullName)
		mergeConfigs(&config, entryConfig)
	}
	log.Info("Using merged config:\n%v", config)
	return config
}

func mergeConfigs(base *Config, override Config) {
	base.Packages = append(base.Packages, override.Packages...)
	base.UserGroups = append(base.UserGroups, override.UserGroups...)
	maps.Copy(base.SystemUnits, override.SystemUnits)
	maps.Copy(base.UserUnits, override.UserUnits)
}

func getConfigDir() string {
	customDir := os.Getenv("PKGSTATE_DIR")
	if customDir != "" {
		return customDir
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Could not determine user's home directory: %v", err)
	}
	return path.Join(homeDir, ".config", "packages")
}

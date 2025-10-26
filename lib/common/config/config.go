package config

import (
	"os"

	"github.com/pelletier/go-toml/v2"
	"github.com/pol-rivero/pkgstate/lib/common/log"
)

type Config struct {
	Packages    []string          `toml:"packages"`
	UserGroups  []string          `toml:"user_groups"`
	SystemUnits map[string]string `toml:"systemd_system"`
	UserUnits   map[string]string `toml:"systemd_user"`
}

func defaultConfig() Config {
	return Config{
		Packages:    []string{},
		UserGroups:  []string{},
		SystemUnits: map[string]string{},
		UserUnits:   map[string]string{},
	}
}

func fromFile(path string) Config {
	config := defaultConfig()
	fileContents, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Error reading config file (%s): %v", path, err)
	}
	err = toml.Unmarshal(fileContents, &config)
	if err != nil {
		log.Fatal("Error parsing config file (%s): %v", path, err)
	}
	return config
}

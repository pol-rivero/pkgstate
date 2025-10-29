package config

import (
	"os"

	"github.com/pol-rivero/pkgstate/lib/common/log"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Packages    []string          `yaml:"packages"`
	UserGroups  []string          `yaml:"groups"`
	SystemUnits map[string]string `yaml:"systemd"`
	UserUnits   map[string]string `yaml:"systemd_user"`
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
	err = yaml.Unmarshal(fileContents, &config)
	if err != nil {
		log.Fatal("Error parsing config file (%s): %v", path, err)
	}
	return config
}

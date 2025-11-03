package config

import (
	"fmt"
	"strings"

	"github.com/pol-rivero/pkgstate/lib/common"
	"gopkg.in/yaml.v3"
)

func (config Config) String() string {
	sortConfig(&config)
	writer := &strings.Builder{}
	encoder := yaml.NewEncoder(writer)
	encoder.SetIndent(2)
	err := encoder.Encode(config)
	if err != nil {
		return fmt.Sprintf("Error serializing config to YAML: %v", err)
	}
	yamlString := writer.String()
	return addSectionSeparators(yamlString)
}

func sortConfig(cfg *Config) {
	cfg.Packages = common.Sorted(cfg.Packages)
	cfg.UserGroups = common.Sorted(cfg.UserGroups)
}

func addSectionSeparators(yamlString string) string {
	result := &strings.Builder{}
	lines := strings.Split(yamlString, "\n")
	for i, line := range lines {
		if i > 0 && strings.HasSuffix(line, ":") {
			result.WriteString("\n")
		}
		result.WriteString(line)
		result.WriteString("\n")
	}
	return result.String()
}

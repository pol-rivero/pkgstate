package printconfig

import (
	"fmt"
	"strings"

	"github.com/pol-rivero/pkgstate/lib/common"
	"github.com/pol-rivero/pkgstate/lib/common/config"
	"github.com/pol-rivero/pkgstate/lib/common/log"
	"gopkg.in/yaml.v3"
)

func PrintConfig() {
	config := config.GetConfig()
	sortConfig(&config)
	yaml, err := prettyPrintConfig(config)
	if err != nil {
		log.Fatal("Failed to serialize config to YAML: %v", err)
	}
	fmt.Println(yaml)
}

func sortConfig(cfg *config.Config) {
	cfg.Packages = common.Sorted(cfg.Packages)
	cfg.UserGroups = common.Sorted(cfg.UserGroups)
}

func prettyPrintConfig(cfg config.Config) (string, error) {
	writer := &strings.Builder{}
	encoder := yaml.NewEncoder(writer)
	encoder.SetIndent(2)
	err := encoder.Encode(cfg)
	if err != nil {
		return "", err
	}
	return writer.String(), nil
}

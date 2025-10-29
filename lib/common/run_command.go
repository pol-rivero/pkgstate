package common

import (
	"os/exec"
	"strings"

	"github.com/pol-rivero/pkgstate/lib/common/log"
)

func RunCommand(name string, args ...string) (string, error) {
	log.Info("Running command: %s %s", name, strings.Join(args, " "))
	cmd := exec.Command(name, args...)
	outputBytes, err := cmd.CombinedOutput()
	output := strings.TrimSpace(string(outputBytes))
	if err != nil {
		log.Error("Command failed: %s %s\nOutput: %s", name, strings.Join(args, " "), output)
		return output, err
	}
	return output, nil
}

func RunCommandGetLines(name string, args ...string) ([]string, error) {
	output, err := RunCommand(name, args...)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(output, "\n")
	return lines, nil
}

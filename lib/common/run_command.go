package common

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/pol-rivero/pkgstate/lib/common/log"
)

func RunCommand(commandAndArgs ...string) (string, error) {
	log.Info("Running command: %s", strings.Join(commandAndArgs, " "))
	cmd := exec.Command(commandAndArgs[0], commandAndArgs[1:]...)
	outputBytes, err := cmd.CombinedOutput()
	output := strings.TrimSpace(string(outputBytes))
	if err != nil {
		return output, fmt.Errorf("command '%s' failed (%w):\n%s", strings.Join(commandAndArgs, " "), err, output)
	}
	return output, nil
}

func RunCommandGetLines(commandAndArgs ...string) ([]string, error) {
	output, err := RunCommand(commandAndArgs...)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(output, "\n")
	return lines, nil
}

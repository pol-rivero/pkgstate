package common

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/pol-rivero/pkgstate/lib/common/log"
)

func RunCommand(commandAndArgs ...string) error {
	log.Info("Running command: %s", strings.Join(commandAndArgs, " "))
	cmd := exec.Command(commandAndArgs[0], commandAndArgs[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("command '%s' failed (%w)", strings.Join(commandAndArgs, " "), err)
	}
	return nil
}

func RunCommandGetOutput(commandAndArgs ...string) (string, error) {
	log.Info("Running command: %s", strings.Join(commandAndArgs, " "))
	cmd := exec.Command(commandAndArgs[0], commandAndArgs[1:]...)
	cmd.Stdin = os.Stdin
	outputBytes, err := cmd.CombinedOutput()
	output := strings.TrimSpace(string(outputBytes))
	if err != nil {
		return output, fmt.Errorf("command '%s' failed (%w):\n%s", strings.Join(commandAndArgs, " "), err, output)
	}
	return output, nil
}

func RunCommandGetLines(commandAndArgs ...string) ([]string, error) {
	output, err := RunCommandGetOutput(commandAndArgs...)
	lines := strings.Split(output, "\n")
	nonEmptyLines := make([]string, 0, len(lines))
	for _, line := range lines {
		if line != "" {
			nonEmptyLines = append(nonEmptyLines, line)
		}
	}
	return nonEmptyLines, err
}

func IsCommandAvailable(commandName string) bool {
	_, err := exec.LookPath(commandName)
	return err == nil
}

package common

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/pol-rivero/pkgstate/lib/common/log"
)

func RunCommand(echoToTerminal bool, commandAndArgs ...string) (string, error) {
	log.Info("Running command: %s", strings.Join(commandAndArgs, " "))
	cmd := exec.Command(commandAndArgs[0], commandAndArgs[1:]...)
	cmd.Stdin = os.Stdin
	var buf bytes.Buffer
	if echoToTerminal {
		cmd.Stdout = io.MultiWriter(os.Stdout, &buf)
		cmd.Stderr = io.MultiWriter(os.Stderr, &buf)
	} else {
		cmd.Stdout = &buf
		cmd.Stderr = &buf
	}
	err := cmd.Run()
	output := strings.TrimSpace(buf.String())
	if err != nil {
		if echoToTerminal {
			return output, fmt.Errorf("command '%s' failed (%w)", strings.Join(commandAndArgs, " "), err)
		} else {
			return output, fmt.Errorf("command '%s' failed (%w):\n%s", strings.Join(commandAndArgs, " "), err, output)
		}
	}
	return output, nil
}

func RunCommandGetLines(echoToTerminal bool, commandAndArgs ...string) ([]string, error) {
	output, err := RunCommand(echoToTerminal, commandAndArgs...)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(output, "\n")
	return lines, nil
}

func IsCommandAvailable(commandName string) bool {
	_, err := exec.LookPath(commandName)
	return err == nil
}

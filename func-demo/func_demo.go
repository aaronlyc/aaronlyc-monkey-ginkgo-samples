package func_demo

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

func Exec(cmd string, args ...string) (string, error) {
	cmdPath, err := exec.LookPath(cmd)
	if err != nil {
		fmt.Errorf("exec.LookPath err: %v, cmd: %s", err, cmd)
		return "", errors.New("any")
	}

	var output []byte
	output, err = exec.Command(cmdPath, args...).CombinedOutput()
	if err != nil {
		fmt.Errorf("exec.Command.CombinedOutput err: %v, cmd: %s", err, cmd)
		return "", errors.New("any")
	}
	fmt.Println("CMD[", cmdPath, "]ARGS[", args, "]OUT[", string(output), "]")
	return string(output), nil
}

func Belong(points string, lines []string) bool {
	flag := false
	for _, line := range lines {
		flag = true
		for _, r := range points {
			if !strings.ContainsRune(line, r) {
				flag = false
				break
			}
		}
		if flag {
			return true
		}
	}
	return false
}
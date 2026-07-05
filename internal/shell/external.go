package shell

import (
	"fmt"
	"os"
	"os/exec"
)

func Run(cmd string, args []string) error {
	c := exec.Command(cmd, args...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		fmt.Printf("rune: command not found: %s\n", cmd)
		return nil
	}
	return nil
}

func Path(command string) (string, error) {
	path, err := exec.LookPath(command)
	if err != nil {
		return "", err
	}
	return path, nil
}

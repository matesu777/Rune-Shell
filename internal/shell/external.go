package shell

import (
	"os"
	"os/exec"

	"github.com/matesu777/Rune/internal/parser"
)

func Run(cmd parser.Command) error {
	command := exec.Command(cmd.Name, cmd.Args...)

	command.Stdin = os.Stdin
	command.Stdout = os.Stdout

	return command.Run()
}

func Path(command string) (string, error) {
	path, err := exec.LookPath(command)
	if err != nil {
		return "", err
	}
	return path, nil
}

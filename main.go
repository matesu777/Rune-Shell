package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/matesu777/Rune/internal/builtin"
	"github.com/matesu777/Rune/internal/shell"
)

func main() {
	reg := shell.NewRegistry()
	builtin.RegistryAll(reg)

	sh := &shell.Shell{
		Registry: reg,
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			continue
		}

		args := strings.Fields(input)

		cmd := shell.Command{
			Name: args[0],
			Args: args[1:],
		}

		if err := sh.Execute(cmd); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

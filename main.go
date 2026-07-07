package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/matesu777/Rune/internal/builtin"
	"github.com/matesu777/Rune/internal/parser"
	"github.com/matesu777/Rune/internal/shell"
)

func main() {
	reg := shell.NewRegistry()
	builtin.RegistryAll(reg)

	sh := &shell.Shell{
		Registry: reg,
	}

	reader := bufio.NewScanner(os.Stdin)

	for {

		fmt.Print("$ ")

		reader.Scan()

		input := reader.Text()

		lex := parser.NewLexer(input)

		p := parser.NewParser(lex)

		cmd := p.ParseCommand()

		if cmd.Name == "" {
			continue
		}

		err := sh.Execute(cmd)

		if err != nil {
			fmt.Println(err)
		}
		if err := reader.Err(); err != nil {
			fmt.Println("erro lendo entrada:", err)
		}
	}
}

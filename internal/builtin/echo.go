package builtin

import (
	"fmt"
	"strings"

	"github.com/matesu777/Rune/internal/shell"
)

func NewEchoCmd(r *shell.Registry) shell.BuiltinFunc {
	return func(args []string) error {
		fmt.Println(strings.Join(args, " "))
		return nil
	}
}

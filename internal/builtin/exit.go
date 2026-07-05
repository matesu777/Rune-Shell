package builtin

import (
	"os"

	"github.com/matesu777/Rune/internal/shell"
)

func NewExitCmd(r *shell.Registry) shell.BuiltinFunc {
	return func(args []string) error {
		os.Exit(0)
		return nil
	}
}

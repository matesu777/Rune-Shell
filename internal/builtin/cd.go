package builtin

import (
	"fmt"
	"os"

	"github.com/matesu777/Rune/internal/shell"
)

func NewCdCmd(r *shell.Registry) shell.BuiltinFunc {
	return func(args []string) error {
		target_dir := os.Getenv("HOME")
		if len(args) > 0 {
			target_dir = args[0]
			if target_dir == "~" {
				target_dir = os.Getenv("HOME")
			}
		}
		if err := os.Chdir(target_dir); err != nil {
			fmt.Printf("cd: %s: No such file or directory\n", target_dir)
			return nil
		}
		return nil
	}
}

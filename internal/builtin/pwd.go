package builtin

import (
	"fmt"
	"os"

	"github.com/matesu777/Rune/internal/shell"
)

func NewPwdCmd(r *shell.Registry) shell.BuiltinFunc {
	return func(args []string) error {
		dir, err := os.Getwd()
		if err != nil {
			return err
		}
		fmt.Println(dir)
		return nil
	}
}

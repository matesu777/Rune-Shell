package builtin

import (
	"fmt"

	"github.com/matesu777/Rune/internal/shell"
)

func NewTypeCmd(r *shell.Registry) shell.BuiltinFunc {
	return func(agrs []string) error {

		if len(agrs) == 0 {
			fmt.Println("usage: type <command>")
			return nil
		}
		if r.Exist(agrs[0]) {
			fmt.Printf("%s is a shell builtin\n", agrs[0])
			return nil
		}

		path, err := shell.Path(agrs[0])

		if err != nil {
			return err
		}
		fmt.Printf("%s is %s\n", agrs[0], path)
		return nil
	}
}

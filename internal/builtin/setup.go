package builtin

import "github.com/matesu777/Rune/internal/shell"

func RegistryAll(r *shell.Registry) {
	r.Add(shell.Builtin{
		Name: "type",
		Help: "Show command type",
		Run:  NewTypeCmd(r),
	})
	r.Add(shell.Builtin{
		Name: "pwd",
		Help: "Show path current",
		Run:  NewPwdCmd(r),
	})
	r.Add(shell.Builtin{
		Name: "exit",
		Help: "Close shell",
		Run:  NewExitCmd(r),
	})
	r.Add(shell.Builtin{
		Name: "cd",
		Help: "Change directory",
		Run:  NewCdCmd(r),
	})
	r.Add(shell.Builtin{
		Name: "echo",
		Help: "Print in terminal",
		Run:  NewEchoCmd(r),
	})
}

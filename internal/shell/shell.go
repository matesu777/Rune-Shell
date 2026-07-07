package shell

import (
	"github.com/matesu777/Rune/internal/parser"
)

type Shell struct {
	Registry *Registry
}

func (s *Shell) Execute(cmd parser.Command) error {
	if builtin, ok := s.Registry.Get(cmd.Name); ok {
		return builtin.Run(cmd.Args)
	}
	return Run(cmd)
}

package shell

type Shell struct {
	Registry *Registry
}

type Command struct {
	Name string
	Args []string
}

func (s *Shell) Execute(cmd Command) error {
	if builtin, ok := s.Registry.Get(cmd.Name); ok {
		return builtin.Run(cmd.Args)
	}
	return Run(cmd.Name, cmd.Args)
}

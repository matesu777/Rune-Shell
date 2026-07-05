package shell

type BuiltinFunc func([]string) error

type Builtin struct {
	Name string
	Help string
	Run  BuiltinFunc
}

type Registry struct {
	builtins map[string]Builtin
}

func NewRegistry() *Registry {
	return &Registry{
		builtins: make(map[string]Builtin),
	}
}

func (r *Registry) Add(builtin Builtin) {
	r.builtins[builtin.Name] = builtin
}

func (r *Registry) Get(cmd string) (Builtin, bool) {
	builtin, ok := r.builtins[cmd]
	return builtin, ok
}

func (r *Registry) List() []Builtin {
	list := []Builtin{}
	for _, value := range r.builtins {
		list = append(list, value)
	}
	return list
}

func (r *Registry) Exist(name string) bool {
	_, exist := r.Get(name)
	if exist {
		return true
	}
	return false
}

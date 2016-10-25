package schema

type Action struct {
	Exec ExecuteFunc
}

type ExecuteFunc func(map[string]interface{}, interface{}) error

func (a *Action) Execute(c map[string]interface{}, m interface{}) error {
	return a.Exec(c, m)
}

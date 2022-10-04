package plugins

type PlugA struct {
}

func (pa *PlugA) Execute() string {
	return "hello from plugin A"
}

package achieve

// Action ...
type Action struct {
	Provider   string
	Name       string
	RawOptions map[string]interface{}
}

type ActionProviderFunc func() ActionProvider

type ActionProvider interface {
	Execute(*Action) error
	Configure(c *ProviderConfig) error
}

type ActionProviderFactory func() (ActionProvider, error)

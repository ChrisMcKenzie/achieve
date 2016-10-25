package achieve

import "fmt"

// Config ...
type Config struct {
	Tasks           map[string]*Task
	ProviderConfigs []*ProviderConfig
}

// GetTask checks if a task exist by the given name and then returns or errors
func (c *Config) GetTask(name string) (*Task, error) {
	task, ok := c.Tasks[name]
	if !ok {
		return nil, fmt.Errorf("Task \"%s\" not found", name)
	}

	return task, nil
}

func (c *Config) GetProviderConfig(name string) (*ProviderConfig, error) {
	for _, provider := range c.ProviderConfigs {
		if provider.Name == name {
			return provider, nil
		}
	}

	return nil, fmt.Errorf("Provider Config \"%s\" not found", name)
}

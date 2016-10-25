package schema

import (
	"fmt"

	"github.com/ChrisMcKenzie/achieve/pkg"
)

// ConfigFunc is a function for configuring providers
type ConfigFunc func(*achieve.ProviderConfig) (interface{}, error)

// Provider is a struct for defining what items a Provider provides
type Provider struct {
	Actions    map[string]*Action
	ConfigFunc ConfigFunc

	meta interface{}
}

func (p *Provider) Configure(c *achieve.ProviderConfig) error {
	if p.ConfigFunc == nil {
		return nil
	}

	meta, err := p.ConfigFunc(c)
	if err != nil {
		return err
	}

	p.meta = meta
	return nil
}

func (p *Provider) Execute(info *achieve.Action) error {
	a, ok := p.Actions[info.Name]
	if !ok {
		return fmt.Errorf("unknown action type %s", info.Name)
	}

	return a.Execute(info.RawOptions, p.meta)
}

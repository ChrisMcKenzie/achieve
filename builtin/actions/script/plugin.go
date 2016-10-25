package script

import (
	"github.com/ChrisMcKenzie/achieve/pkg"
	"github.com/ChrisMcKenzie/achieve/pkg/schema"
)

// Provider is an ActionFunc for instantiating and ActionProvider for github
func Provider() achieve.ActionProvider {
	return &schema.Provider{
		Actions: map[string]*schema.Action{
			"script_run": runScript(),
		},
	}
}

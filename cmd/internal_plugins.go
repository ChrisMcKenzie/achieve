package cmd

import (
	"github.com/ChrisMcKenzie/achieve/builtin/actions/github"
	"github.com/ChrisMcKenzie/achieve/builtin/actions/script"
	"github.com/ChrisMcKenzie/achieve/pkg/plugin"
)

// InternalActions ...
var InternalActions = map[string]plugin.ActionFunc{
	"github": github.Provider,
	"script": script.Provider,
}

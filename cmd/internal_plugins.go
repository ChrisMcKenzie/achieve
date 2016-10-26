package cmd

import (
	"github.com/ChrisMcKenzie/achieve/builtin/actions/github"
	"github.com/ChrisMcKenzie/achieve/builtin/actions/script"
	"github.com/ChrisMcKenzie/achieve/pkg"
)

// InternalActions ...
var InternalActions = map[string]achieve.ActionProviderFunc{
	"github": github.Provider,
	"script": script.Provider,
}

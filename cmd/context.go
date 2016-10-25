// Copyright © 2016 Chris McKenzie <chris@chrismckenzie.io>

package cmd

import (
	"fmt"

	"github.com/ChrisMcKenzie/achieve/pkg"
)

// Context ...
type Context struct {
	Config *achieve.Config
	task   string
}

// NewContext returns a new Context
func NewContext(task string, conf *achieve.Config) *Context {
	return &Context{conf, task}
}

// Execute ...
func (ctx *Context) Execute() {
	t, err := ctx.Config.GetTask(ctx.task)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, action := range t.Actions {
		err := ctx.executeAction(action)
		if err != nil {
			fmt.Printf("ERR: %s\n", err)
		}
	}
}

func (ctx *Context) executeAction(action *achieve.Action) error {
	provider, ok := InternalActions[action.Provider]
	if !ok {
		return fmt.Errorf("Action provider %s not found\n", action.Name)
	}
	ap := provider()

	pc, _ := ctx.Config.GetProviderConfig(action.Provider)
	err := ap.Configure(pc)
	if err != nil {
		return err
	}

	return ap.Execute(action)
}

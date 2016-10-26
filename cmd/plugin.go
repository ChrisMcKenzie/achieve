// Copyright © 2016 Chris McKenzie <chris@chrismckenzie.io>
//

package cmd

import (
	"fmt"
	"log"

	acPlugin "github.com/ChrisMcKenzie/achieve/pkg/plugin"
	"github.com/spf13/cobra"
)

// pluginCmd represents the plugin command
var pluginCmd = &cobra.Command{
	Use:   "internal-plugin pluginName",
	Short: "internal plugin command",
	Long: `
	Runs an internally-compiled version of a plugin from the terraform binary.
  NOTE: this is an internal command and you should not call it yourself.
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("Wrong number of args; expected: do internal-plugin pluginName")
		}

		pluginName := args[0]

		pluginFunc, found := InternalActions[pluginName]
		if !found {
			return fmt.Errorf("Could not load provider: %s", pluginName)
		}
		log.Printf("[INFO] Starting provider plugin \"%s\"", pluginName)

		acPlugin.Serve(&acPlugin.ServeOpts{
			ActionFunc: pluginFunc,
		})
		return nil
	},
}

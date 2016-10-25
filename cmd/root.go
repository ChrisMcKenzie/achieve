// Copyright © 2016 Chris McKenzie <chris@chrismckenzie.io>

package cmd

import (
	"fmt"
	"os"

	"github.com/ChrisMcKenzie/achieve/pkg"
	"github.com/spf13/cobra"
)

var cfgFile string
var ctx *Context

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "achieve task",
	Short: "A modern tool for development task automation",
	Long:  ``,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 && args[0] == "internal-plugin" {
			return pluginCmd.RunE(pluginCmd, args[1:])
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		taskName := "default"
		if len(args) > 0 {
			taskName = args[0]
		}

		if cfgFile == "" { // enable ability to specify config file via flag
			cfgFile = "tasks.hcl"
		}

		config, err := achieve.LoadConfig(cfgFile)
		if err != nil {
			return err
		}

		ctx = NewContext(taskName, config)

		fmt.Printf("Executing Task %s\n", taskName)
		ctx.Execute()

		return nil
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $CWD/task.yaml)")
}

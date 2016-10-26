package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/ChrisMcKenzie/achieve/pkg"
	acPlugin "github.com/ChrisMcKenzie/achieve/pkg/plugin"
	"github.com/hashicorp/go-plugin"
	"github.com/kardianos/osext"
)

type Config struct {
	Providers map[string]string
}

var DefaultConfig Config

func (c *Config) LoadPlugins() error {
	if c.Providers == nil {
		c.Providers = make(map[string]string)
	}
	for name := range InternalActions {
		cmd, err := buildCmdString(name)
		if err != nil {
			return err
		}
		c.Providers[name] = cmd
	}
	return nil
}

func (c *Config) ActionProviderFactories() map[string]achieve.ActionProviderFactory {
	result := make(map[string]achieve.ActionProviderFactory)
	for k, v := range c.Providers {
		result[k] = c.actionProviderFactory(v)
	}
	return result
}

func (c *Config) actionProviderFactory(name string) achieve.ActionProviderFactory {
	var config plugin.ClientConfig
	config.Cmd = providerCmd(name)
	config.HandshakeConfig = acPlugin.Handshake
	config.Managed = true
	config.Plugins = acPlugin.PluginMap
	client := plugin.NewClient(&config)

	return func() (achieve.ActionProvider, error) {
		rpcClient, err := client.Client()
		if err != nil {
			return nil, err
		}

		raw, err := rpcClient.Dispense("action")
		if err != nil {
			return nil, err
		}

		return raw.(achieve.ActionProvider), nil
	}
}

func providerCmd(name string) *exec.Cmd {
	parts := strings.Split(name, " ")
	return exec.Command(parts[0], parts[1:]...)
}

func buildCmdString(name string) (string, error) {
	path, err := osext.Executable()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s internal-plugin %s", path, name), nil
}

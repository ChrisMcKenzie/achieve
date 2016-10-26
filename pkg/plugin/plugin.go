package plugin

import (
	"github.com/ChrisMcKenzie/achieve/pkg"
	"github.com/hashicorp/go-plugin"
)

var PluginMap = map[string]plugin.Plugin{
	"action": &ActionProviderPlugin{},
}

// Handshake is the HandshakeConfig used to configure clients and servers.
var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "AC_PLUGIN_MAGIC_COOKIE",
	MagicCookieValue: "secret-cat",
}

// ServeOpts are the configurations to serve a plugin.
type ServeOpts struct {
	ActionFunc achieve.ActionProviderFunc
}

// Serve serves a plugin. This function never returns and should be the final
// function called in the main function of the plugin.
func Serve(opts *ServeOpts) {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: Handshake,
		Plugins:         pluginMap(opts),
	})
}

// pluginMap returns the map[string]plugin.Plugin to use for configuring a plugin
// server or client.
func pluginMap(opts *ServeOpts) map[string]plugin.Plugin {
	return map[string]plugin.Plugin{
		"action": &ActionProviderPlugin{F: opts.ActionFunc},
	}
}

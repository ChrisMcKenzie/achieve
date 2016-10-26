package plugin

import (
	"net/rpc"

	"github.com/ChrisMcKenzie/achieve/pkg"
	"github.com/hashicorp/go-plugin"
)

// ActionProviderPlugin is the plugin.Plugin implementation.
type ActionProviderPlugin struct {
	F achieve.ActionProviderFunc
}

func (p *ActionProviderPlugin) Server(b *plugin.MuxBroker) (interface{}, error) {
	return &ActionProviderServer{Broker: b, Provider: p.F()}, nil
}

func (p *ActionProviderPlugin) Client(
	b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &ActionProvider{Broker: b, Client: c}, nil
}

// ActionProvider is an implementation of terraform.ResourceProvider
// that communicates over RPC.
type ActionProvider struct {
	Broker *plugin.MuxBroker
	Client *rpc.Client
}

func (p *ActionProvider) Execute(action *achieve.Action) error {
	var resp ActionProviderExecuteResponse
	arg := ActionProviderExecuteArgs{
		Action: action,
	}

	err := p.Client.Call("Plugin.Execute", arg, &resp)
	if err != nil {
		return err
	}
	return nil
}

func (p *ActionProvider) Configure(conf *achieve.ProviderConfig) error {
	var resp ActionProviderConfigureResponse
	arg := ActionProviderConfigureArgs{
		Config: conf,
	}

	err := p.Client.Call("Plugin.Configure", arg, &resp)
	if err != nil {
		return err
	}
	return nil
}

type ActionProviderServer struct {
	Broker   *plugin.MuxBroker
	Provider achieve.ActionProvider
}

func (s *ActionProviderServer) Execute(
	args *ActionProviderExecuteArgs,
	result *ActionProviderExecuteResponse) error {

	err := s.Provider.Execute(args.Action)
	*result = ActionProviderExecuteResponse{
		Error: plugin.NewBasicError(err),
	}

	return nil
}

func (s *ActionProviderServer) Configure(
	args *ActionProviderConfigureArgs,
	result *ActionProviderConfigureResponse) error {

	err := s.Provider.Configure(args.Config)
	*result = ActionProviderConfigureResponse{
		Error: plugin.NewBasicError(err),
	}

	return nil
}

type ActionProviderExecuteArgs struct {
	Action *achieve.Action
}

type ActionProviderExecuteResponse struct {
	Error *plugin.BasicError
}

type ActionProviderConfigureArgs struct {
	Config *achieve.ProviderConfig
}

type ActionProviderConfigureResponse struct {
	Error *plugin.BasicError
}

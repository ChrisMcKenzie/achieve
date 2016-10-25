package github

import (
	"fmt"

	"github.com/ChrisMcKenzie/achieve/pkg"
	"github.com/ChrisMcKenzie/achieve/pkg/schema"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// Provider is an ActionFunc for instantiating and ActionProvider for github
func Provider() achieve.ActionProvider {
	return &schema.Provider{
		Actions: map[string]*schema.Action{
			"github_create_release": createRelease(),
		},

		ConfigFunc: providerConfig,
	}
}

func providerConfig(c *achieve.ProviderConfig) (interface{}, error) {
	token, ok := c.RawConfig["token"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid token type")
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)
	return client, nil
}

package github

import (
	"fmt"
	"strings"

	"github.com/ChrisMcKenzie/achieve/pkg/schema"
	"github.com/google/go-github/github"
)

func createRelease() *schema.Action {
	return &schema.Action{
		Exec: execCreateRelease,
	}
}

func execCreateRelease(c map[string]interface{}, meta interface{}) error {
	client := meta.(*github.Client)

	if _, ok := c["repo"]; !ok {
		return fmt.Errorf("repo is required")
	}
	owner, repo := parseRepo(c["repo"].(string))

	var title string
	if v, ok := c["title"]; ok {
		title = v.(string)
	} else if v, ok := c["version"]; ok {
		title = v.(string)
	}

	var version string
	if v, ok := c["version"]; ok {
		version = v.(string)
	} else {
		return fmt.Errorf("version is required")
	}

	var targetCommitish string
	if v, ok := c["target_commitish"]; ok {
		targetCommitish = v.(string)
	}

	var preRelease bool
	if v, ok := c["pre_release"]; ok {
		preRelease = v.(bool)
	}

	release := &github.RepositoryRelease{
		TagName:         &version,
		TargetCommitish: &targetCommitish,
		Prerelease:      &preRelease,
	}

	if title != "" {
		release.Name = &title
	}

	_, _, err := client.Repositories.CreateRelease(owner, repo, release)

	return err
}

func parseRepo(path string) (string, string) {
	vals := strings.Split(path, "/")
	return vals[0], vals[1]
}

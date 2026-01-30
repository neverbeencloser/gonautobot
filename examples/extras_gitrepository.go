package main

import (
	"fmt"
	"net/url"

	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/types"
	"github.com/rs/zerolog/log"
)

// GitRepositoryExample : Example usage of the Git Repository Nautobot methods.
func (c *ex) GitRepositoryExample() {
	// Create a new git repository
	repo, err := c.Extras.GitRepositoryCreate(types.NewGitRepository{
		Name:             "example-config-repo",
		RemoteURL:        "https://github.com/example/config-repo.git",
		Branch:           "main",
		ProvidedContents: []string{"extras.configcontext"},
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create git repository")
	}
	fmt.Println("Created git repository:", repo.Name)

	// Filter git repositories by name
	repos, err := c.Extras.GitRepositoryFilter(&url.Values{"name": {"example-config-repo"}})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to filter git repositories")
	}
	first, _ := core.First[types.GitRepository](repos)
	fmt.Println("Found git repository:", first.Name)

	// Get a specific git repository by ID
	retrieved, err := c.Extras.GitRepositoryGet(repo.ID)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get git repository")
	}
	fmt.Println("Retrieved git repository:", retrieved.Name)

	// Update the git repository
	updated, err := c.Extras.GitRepositoryUpdate(repo.ID, map[string]any{
		"branch": "develop",
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to update git repository")
	}
	fmt.Println("Updated git repository branch:", updated.Branch)

	// Delete the git repository
	if err := c.Extras.GitRepositoryDelete(repo.ID); err != nil {
		log.Fatal().Err(err).Msg("Failed to delete git repository")
	}
	fmt.Println("Deleted git repository:", repo.Name)
}

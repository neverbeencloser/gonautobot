package extras

import (
	"net/url"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/types"
)

const (
	extrasEndpointGitRepository = "extras/git-repositories/"
)

// GitRepositoryGet : Get a Git Repository by UUID identifier.
func (c *Client) GitRepositoryGet(id uuid.UUID) (*types.GitRepository, error) {
	return core.Get[types.GitRepository](c.Client, extrasEndpointGitRepository, id)
}

// GitRepositoryFilter : Get a list of Git Repositories based on query parameters.
func (c *Client) GitRepositoryFilter(q *url.Values) ([]types.GitRepository, error) {
	repos := make([]types.GitRepository, 0)
	return repos, core.Paginate[types.GitRepository](c.Client, extrasEndpointGitRepository, q, &repos)
}

// GitRepositoryAll : Get all Git Repositories in Nautobot.
func (c *Client) GitRepositoryAll() ([]types.GitRepository, error) {
	repos := make([]types.GitRepository, 0)
	return repos, core.Paginate[types.GitRepository](c.Client, extrasEndpointGitRepository, nil, &repos)
}

// GitRepositoryCreate : Generate a new Git Repository record in Nautobot.
func (c *Client) GitRepositoryCreate(obj types.NewGitRepository) (*types.GitRepository, error) {
	return core.Create[types.GitRepository, types.NewGitRepository](c.Client, extrasEndpointGitRepository, obj)
}

// GitRepositoryDelete : Delete a Git Repository by UUID identifier.
func (c *Client) GitRepositoryDelete(id uuid.UUID) error {
	return core.Delete(c.Client, extrasEndpointGitRepository, id)
}

// GitRepositoryUpdate : Update an existing Git Repository record in Nautobot.
func (c *Client) GitRepositoryUpdate(id uuid.UUID, patch map[string]any) (*types.GitRepository, error) {
	return core.Update[types.GitRepository](c.Client, extrasEndpointGitRepository, id, patch)
}

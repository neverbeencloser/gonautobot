package extras

import (
	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/query"
	"github.com/neverbeencloser/gonautobot/types"
)

const (
	extrasEndpointGitRepository = "extras/git-repositories/"
)

// GitRepositoryGet : Get a Git Repository by UUID identifier.
func (c *Client) GitRepositoryGet(id uuid.UUID, opts ...query.Option) (*types.GitRepository, error) {
	return core.Get[types.GitRepository](c.Client, extrasEndpointGitRepository, id, opts...)
}

// GitRepositoryList : Get a list of Git Repositories in Nautobot, filtered by any query options.
func (c *Client) GitRepositoryList(opts ...query.Option) ([]types.GitRepository, error) {
	repos := make([]types.GitRepository, 0)
	return repos, core.Paginate[types.GitRepository](c.Client, extrasEndpointGitRepository, opts, &repos)
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

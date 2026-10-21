package extras

import (
	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/query"
	"github.com/neverbeencloser/gonautobot/types"
)

const extrasEndpointJobHook = "extras/job-hooks/"

// JobHookGet : Get a JobHook by UUID identifier.
func (c *Client) JobHookGet(id uuid.UUID, opts ...query.Option) (*types.JobHook, error) {
	return core.Get[types.JobHook](c.Client, extrasEndpointJobHook, id, opts...)
}

// JobHookList : Get a list of JobHooks in Nautobot, filtered by any query options.
func (c *Client) JobHookList(opts ...query.Option) ([]types.JobHook, error) {
	hooks := make([]types.JobHook, 0)
	return hooks, core.Paginate[types.JobHook](c.Client, extrasEndpointJobHook, opts, &hooks)
}

// JobHookCreate : Create a new JobHook.
func (c *Client) JobHookCreate(obj types.NewJobHook) (*types.JobHook, error) {
	return core.Create[types.JobHook, types.NewJobHook](c.Client, extrasEndpointJobHook, obj)
}

// JobHookDelete : Delete a JobHook by UUID identifier.
func (c *Client) JobHookDelete(id uuid.UUID) error {
	return core.Delete(c.Client, extrasEndpointJobHook, id)
}

// JobHookUpdate : Partially update a JobHook (PATCH).
func (c *Client) JobHookUpdate(id uuid.UUID, patch map[string]any) (*types.JobHook, error) {
	return core.Update[types.JobHook](c.Client, extrasEndpointJobHook, id, patch)
}

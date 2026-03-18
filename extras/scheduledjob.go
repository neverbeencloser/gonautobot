package extras

import (
	"net/url"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/types"
)

const extrasEndpointScheduledJob = "extras/scheduled-jobs/"

// ScheduledJobGet : Get a ScheduledJob by UUID identifier.
func (c *Client) ScheduledJobGet(id uuid.UUID) (*types.ScheduledJob, error) {
	return core.Get[types.ScheduledJob](c.Client, extrasEndpointScheduledJob, id)
}

// ScheduledJobFilter : List scheduled jobs matching query parameters.
func (c *Client) ScheduledJobFilter(q *url.Values) ([]types.ScheduledJob, error) {
	out := make([]types.ScheduledJob, 0)
	return out, core.Paginate[types.ScheduledJob](c.Client, extrasEndpointScheduledJob, q, &out)
}

// ScheduledJobAll : List all scheduled jobs.
func (c *Client) ScheduledJobAll() ([]types.ScheduledJob, error) {
	out := make([]types.ScheduledJob, 0)
	return out, core.Paginate[types.ScheduledJob](c.Client, extrasEndpointScheduledJob, nil, &out)
}

// ScheduledJobDelete : Delete a scheduled job by UUID.
func (c *Client) ScheduledJobDelete(id uuid.UUID) error {
	return core.Delete(c.Client, extrasEndpointScheduledJob, id)
}

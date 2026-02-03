package extras

import (
	"net/url"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/types"
)

const (
	extrasEndpointJob       = "extras/jobs/"
	extrasEndpointJobResult = "extras/job-results/"
	actionRun               = "run"
)

// JobGet : Get a Job by UUID identifier.
func (c *Client) JobGet(id uuid.UUID) (*types.Job, error) {
	return core.Get[types.Job](c.Client, extrasEndpointJob, id)
}

// JobFilter : Get a list of Jobs based on query parameters.
func (c *Client) JobFilter(q *url.Values) ([]types.Job, error) {
	jobs := make([]types.Job, 0)
	return jobs, core.Paginate[types.Job](c.Client, extrasEndpointJob, q, &jobs)
}

// JobAll : Get all Jobs in Nautobot.
func (c *Client) JobAll() ([]types.Job, error) {
	jobs := make([]types.Job, 0)
	return jobs, core.Paginate[types.Job](c.Client, extrasEndpointJob, nil, &jobs)
}

// JobUpdate : Update an existing Job record in Nautobot.
// Jobs are created by the system when code is loaded, not via API.
// Use this method to enable/disable jobs or change their settings.
func (c *Client) JobUpdate(id uuid.UUID, patch map[string]any) (*types.Job, error) {
	return core.Update[types.Job](c.Client, extrasEndpointJob, id, patch)
}

// JobRun : Run a Job by UUID identifier.
// Returns a JobRunResponse containing either a ScheduledJob (for jobs requiring approval)
// or a JobResult (for immediate execution).
// The request can include optional data parameters, commit flag, and dryrun flag.
func (c *Client) JobRun(id uuid.UUID, request types.JobRunRequest) (*types.JobRunResponse, error) {
	return core.Action[types.JobRunResponse](c.Client, extrasEndpointJob, id, actionRun, request)
}

// JobResultGet : Get a Job Result by UUID identifier.
func (c *Client) JobResultGet(id uuid.UUID) (*types.JobResult, error) {
	return core.Get[types.JobResult](c.Client, extrasEndpointJobResult, id)
}

// JobResultFilter : Get a list of Job Results based on query parameters.
func (c *Client) JobResultFilter(q *url.Values) ([]types.JobResult, error) {
	results := make([]types.JobResult, 0)
	return results, core.Paginate[types.JobResult](c.Client, extrasEndpointJobResult, q, &results)
}

// JobResultAll : Get all Job Results in Nautobot.
func (c *Client) JobResultAll() ([]types.JobResult, error) {
	results := make([]types.JobResult, 0)
	return results, core.Paginate[types.JobResult](c.Client, extrasEndpointJobResult, nil, &results)
}

// JobResultDelete : Delete a Job Result by UUID identifier.
func (c *Client) JobResultDelete(id uuid.UUID) error {
	return core.Delete(c.Client, extrasEndpointJobResult, id)
}

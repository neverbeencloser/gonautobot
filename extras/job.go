package extras

import (
	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/query"
	"github.com/neverbeencloser/gonautobot/types"
)

const (
	extrasEndpointJob       = "extras/jobs/"
	extrasEndpointJobResult = "extras/job-results/"
	extrasEndpointJobLog    = "extras/job-logs/"
	actionRun               = "run"
)

// JobGet : Get a Job by UUID identifier.
func (c *Client) JobGet(id uuid.UUID, opts ...query.Option) (*types.Job, error) {
	return core.Get[types.Job](c.Client, extrasEndpointJob, id, opts...)
}

// JobList : Get a list of Jobs in Nautobot, filtered by any query options.
func (c *Client) JobList(opts ...query.Option) ([]types.Job, error) {
	jobs := make([]types.Job, 0)
	return jobs, core.Paginate[types.Job](c.Client, extrasEndpointJob, opts, &jobs)
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
func (c *Client) JobResultGet(id uuid.UUID, opts ...query.Option) (*types.JobResult, error) {
	return core.Get[types.JobResult](c.Client, extrasEndpointJobResult, id, opts...)
}

// JobResultList : Get a list of Job Results in Nautobot, filtered by any query options.
func (c *Client) JobResultList(opts ...query.Option) ([]types.JobResult, error) {
	results := make([]types.JobResult, 0)
	return results, core.Paginate[types.JobResult](c.Client, extrasEndpointJobResult, opts, &results)
}

// JobResultDelete : Delete a Job Result by UUID identifier.
func (c *Client) JobResultDelete(id uuid.UUID) error {
	return core.Delete(c.Client, extrasEndpointJobResult, id)
}

// JobLogGet : Get a Job Log entry by UUID identifier.
func (c *Client) JobLogGet(id uuid.UUID, opts ...query.Option) (*types.JobLog, error) {
	return core.Get[types.JobLog](c.Client, extrasEndpointJobLog, id, opts...)
}

// JobLogList : Get a list of Job Log entries in Nautobot, filtered by any query options.
func (c *Client) JobLogList(opts ...query.Option) ([]types.JobLog, error) {
	logs := make([]types.JobLog, 0)
	return logs, core.Paginate[types.JobLog](c.Client, extrasEndpointJobLog, opts, &logs)
}

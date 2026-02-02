package types

import (
	"time"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/types/nested"
)

type (
	// JobRunRequest : Request body for running a job.
	JobRunRequest struct {
		Data   map[string]any `json:"data,omitempty"`
		Commit bool           `json:"commit,omitempty"`
		Dryrun bool           `json:"dryrun,omitempty"`
	}

	// JobRunResponse : Response from running a job.
	// Contains either a scheduled_job (for jobs requiring approval) or job_result (for immediate execution).
	JobRunResponse struct {
		ScheduledJob *ScheduledJob `json:"scheduled_job"`
		JobResult    *JobResult    `json:"job_result"`
	}

	// JobResult : Result of a job execution.
	JobResult struct {
		ID             uuid.UUID      `json:"id"`
		ObjectType     string         `json:"object_type"`
		Display        string         `json:"display"`
		URL            string         `json:"url"`
		NaturalSlug    string         `json:"natural_slug"`
		Status         *JobStatus     `json:"status"`
		Name           string         `json:"name"`
		TaskName       *string        `json:"task_name"`
		DateCreated    time.Time      `json:"date_created"`
		DateStarted    *time.Time     `json:"date_started"`
		DateDone       *time.Time     `json:"date_done"`
		Result         any            `json:"result"`
		Worker         *string        `json:"worker"`
		TaskArgs       []any          `json:"task_args"`
		TaskKwargs     map[string]any `json:"task_kwargs"`
		CeleryKwargs   map[string]any `json:"celery_kwargs"`
		Traceback      *string        `json:"traceback"`
		Meta           any            `json:"meta"`
		JobModel       *nested.Job    `json:"job_model"`
		User           *nested.User   `json:"user"`
		ScheduledJob   *ScheduledJob  `json:"scheduled_job"`
		CustomFields   map[string]any `json:"custom_fields"`
		ComputedFields map[string]any `json:"computed_fields"`
		Files          []any          `json:"files"`
	}

	// JobStatus : Status of a job result.
	JobStatus struct {
		Value string `json:"value"`
		Label string `json:"label"`
	}

	// ScheduledJob : A scheduled job awaiting execution or approval.
	ScheduledJob struct {
		ID               uuid.UUID    `json:"id"`
		ObjectType       string       `json:"object_type"`
		Display          string       `json:"display"`
		URL              string       `json:"url"`
		JobModel         *nested.Job  `json:"job_model"`
		Name             string       `json:"name"`
		Task             string       `json:"task"`
		Interval         string       `json:"interval"`
		Queue            string       `json:"queue"`
		User             *nested.User `json:"user"`
		StartTime        time.Time    `json:"start_time"`
		Description      string       `json:"description"`
		ApprovalRequired bool         `json:"approval_required"`
		Approved         bool         `json:"approved"`
		ApprovedAt       *time.Time   `json:"approved_at"`
		ApprovedByUser   *nested.User `json:"approved_by_user"`
		CrontabSpec      string       `json:"crontab"`
		Created          time.Time    `json:"created"`
		LastUpdated      time.Time    `json:"last_updated"`
	}

	// Job : Data type entry for a Job in Nautobot.
	Job struct {
		ID                            uuid.UUID         `json:"id"`
		ObjectType                    string            `json:"object_type"`
		Display                       string            `json:"display"`
		URL                           string            `json:"url"`
		NaturalSlug                   string            `json:"natural_slug"`
		TaskQueues                    []string          `json:"task_queues"`
		TaskQueuesOverride            bool              `json:"task_queues_override"`
		ModuleName                    string            `json:"module_name"`
		JobClassName                  string            `json:"job_class_name"`
		Grouping                      string            `json:"grouping"`
		Name                          string            `json:"name"`
		Description                   string            `json:"description"`
		Installed                     bool              `json:"installed"`
		Enabled                       bool              `json:"enabled"`
		IsJobHookReceiver             bool              `json:"is_job_hook_receiver"`
		IsJobButtonReceiver           bool              `json:"is_job_button_receiver"`
		HasSensitiveVariables         bool              `json:"has_sensitive_variables"`
		IsSingleton                   bool              `json:"is_singleton"`
		ApprovalRequired              bool              `json:"approval_required"`
		Hidden                        bool              `json:"hidden"`
		DryrunDefault                 bool              `json:"dryrun_default"`
		ReadOnly                      bool              `json:"read_only"`
		SoftTimeLimit                 float64           `json:"soft_time_limit"`
		TimeLimit                     float64           `json:"time_limit"`
		SupportsDryrun                bool              `json:"supports_dryrun"`
		GroupingOverride              bool              `json:"grouping_override"`
		NameOverride                  bool              `json:"name_override"`
		DescriptionOverride           bool              `json:"description_override"`
		ApprovalRequiredOverride      bool              `json:"approval_required_override"`
		DryrunDefaultOverride         bool              `json:"dryrun_default_override"`
		HiddenOverride                bool              `json:"hidden_override"`
		SoftTimeLimitOverride         bool              `json:"soft_time_limit_override"`
		TimeLimitOverride             bool              `json:"time_limit_override"`
		HasSensitiveVariablesOverride bool              `json:"has_sensitive_variables_override"`
		JobQueuesOverride             bool              `json:"job_queues_override"`
		DefaultJobQueueOverride       bool              `json:"default_job_queue_override"`
		IsSingletonOverride           bool              `json:"is_singleton_override"`
		DefaultJobQueue               *nested.JobQueue  `json:"default_job_queue"`
		JobQueues                     []nested.JobQueue `json:"job_queues"`
		Created                       time.Time         `json:"created"`
		LastUpdated                   time.Time         `json:"last_updated"`
		Tags                          []Tag             `json:"tags"`
		NotesURL                      string            `json:"notes_url"`
		CustomFields                  map[string]any    `json:"custom_fields"`
		ComputedFields                map[string]any    `json:"computed_fields"`
		Relationships                 map[string]any    `json:"relationships"`
	}
)

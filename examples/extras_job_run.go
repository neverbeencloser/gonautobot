package main

import (
	"net/url"

	"github.com/neverbeencloser/gonautobot/types"
	"github.com/rs/zerolog/log"
)

// JobRunExample : Example usage of running a Job in Nautobot.
// This example shows how to run a job with optional parameters.
func (c *ex) JobRunExample() {
	const jobName = "Logs Cleanup"
	jobResult, err := c.Client.Extras.JobFilter(&url.Values{"name": {jobName}})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get job")
	}
	if len(jobResult) == 0 {
		log.Fatal().Msg("No job found")
	}

	job := jobResult[0]
	log.Info().
		Str("uuid", job.ID.String()).
		Str("name", job.Name).
		Bool("enabled", job.Enabled).
		Bool("installed", job.Installed).
		Bool("supports_dryrun", job.SupportsDryrun).
		Msg("Retrieved job to run")

	// Verify the job can be run
	if !job.Enabled {
		log.Error().
			Str("uuid", job.ID.String()).
			Str("name", job.Name).
			Msg("Job is not enabled, cannot run")
		return
	}
	if !job.Installed {
		log.Error().
			Str("uuid", job.ID.String()).
			Str("name", job.Name).
			Msg("Job is not installed, cannot run")
		return
	}

	// Run the job with parameters
	resp, err := c.Extras.JobRun(job.ID, types.JobRunRequest{
		Data: map[string]any{
			"max_age":       90,
			"cleanup_types": []string{"extras.JobResult"},
		},
		Dryrun: false, // Set to true to run in dryrun mode
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to run job")
	}

	// Check if job was scheduled (requires approval) or executed immediately
	if resp.ScheduledJob != nil {
		// Job requires approval and was scheduled
		log.Info().
			Str("scheduled_job_uuid", resp.ScheduledJob.ID.String()).
			Str("name", resp.ScheduledJob.Name).
			Str("queue", resp.ScheduledJob.Queue).
			Str("interval", resp.ScheduledJob.Interval).
			Time("start_time", resp.ScheduledJob.StartTime).
			Bool("approval_required", resp.ScheduledJob.ApprovalRequired).
			Bool("approved", resp.ScheduledJob.Approved).
			Msg("Job scheduled (awaiting approval)")

		if resp.ScheduledJob.User != nil {
			log.Info().
				Str("user_id", resp.ScheduledJob.User.ID).
				Str("username", resp.ScheduledJob.User.Username).
				Msg("Scheduled job submitted by user")
		}
	}

	if resp.JobResult != nil {
		// Job was executed immediately
		log.Info().
			Str("job_result_uuid", resp.JobResult.ID.String()).
			Str("name", resp.JobResult.Name).
			Str("status", resp.JobResult.Status.Value).
			Time("date_created", resp.JobResult.DateCreated).
			Str("url", resp.JobResult.URL).
			Msg("Job result created")

		// Log job model info
		if resp.JobResult.JobModel != nil {
			log.Info().
				Str("job_model_id", resp.JobResult.JobModel.ID).
				Msg("Job model")
		}

		// Log user info
		if resp.JobResult.User != nil {
			log.Info().
				Str("user_id", resp.JobResult.User.ID).
				Msg("Job submitted by user")
		}
	}
}

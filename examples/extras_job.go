package main

import (
	"net/url"

	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/types"
	"github.com/rs/zerolog/log"
)

// JobExample : Example usage of the Job Nautobot methods.
// Note: Jobs are created by the system when code is loaded, not via API.
// This example shows how to list, filter, get, and update jobs.
func (c *ex) JobExample() {
	// Get all jobs
	jobs, err := c.Extras.JobAll()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get all jobs")
	}
	log.Info().
		Int("count", len(jobs)).
		Msg("Retrieved all jobs")

	// Filter jobs by enabled status
	enabledJobs, err := c.Extras.JobFilter(&url.Values{"enabled": {"true"}})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to filter jobs")
	}
	log.Info().
		Int("count", len(enabledJobs)).
		Str("filter", "enabled=true").
		Msg("Filtered enabled jobs")

	// Get the first job from the list
	if len(jobs) == 0 {
		log.Warn().Msg("No jobs found in Nautobot")
		return
	}

	first, _ := core.First[types.Job](jobs)
	log.Info().
		Str("uuid", first.ID.String()).
		Str("name", first.Name).
		Str("module_name", first.ModuleName).
		Str("job_class_name", first.JobClassName).
		Str("grouping", first.Grouping).
		Bool("enabled", first.Enabled).
		Bool("installed", first.Installed).
		Msg("First job from list")

	// Get a specific job by ID
	job, err := c.Extras.JobGet(first.ID)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get job")
	}
	log.Info().
		Str("uuid", job.ID.String()).
		Str("name", job.Name).
		Str("description", job.Description).
		Bool("enabled", job.Enabled).
		Bool("approval_required", job.ApprovalRequired).
		Bool("is_singleton", job.IsSingleton).
		Bool("supports_dryrun", job.SupportsDryrun).
		Float64("soft_time_limit", job.SoftTimeLimit).
		Float64("time_limit", job.TimeLimit).
		Time("last_updated", job.LastUpdated).
		Msg("Retrieved job by UUID")

	// Update job settings (e.g., disable the job)
	updated, err := c.Extras.JobUpdate(job.ID, map[string]any{
		"enabled": false,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to update job")
	}
	log.Info().
		Str("uuid", updated.ID.String()).
		Str("name", updated.Name).
		Bool("enabled", updated.Enabled).
		Time("last_updated", updated.LastUpdated).
		Msg("Updated job (disabled)")

	// Re-enable the job
	reenabled, err := c.Extras.JobUpdate(job.ID, map[string]any{
		"enabled": true,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to re-enable job")
	}
	log.Info().
		Str("uuid", reenabled.ID.String()).
		Str("name", reenabled.Name).
		Bool("enabled", reenabled.Enabled).
		Msg("Re-enabled job")
}

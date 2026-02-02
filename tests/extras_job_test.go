package nautobot_test

import (
	"net/url"
	"path"
	"testing"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/h2non/gock.v1"
)

const (
	testJobID       = "f1a2b3c4-d5e6-7890-abcd-ef1234567890"
	testJobResultID = "85b43318-e232-4494-9265-e53440f6bb7f"
)

func TestClient_JobGet(t *testing.T) {
	gock.New(testURL).Get("extras/jobs/" + testJobID + "/").Reply(200).
		File(path.Join("fixtures", "extras", "job_200_1.json"))

	id, err := uuid.Parse(testJobID)
	require.NoError(t, err)

	resp, err := testClient.Extras.JobGet(id)
	require.NoError(t, err)
	assert.Equal(t, "Device Onboarding", resp.Name)
	assert.Equal(t, id, resp.ID)
	assert.Equal(t, "nautobot_device_onboarding.jobs", resp.ModuleName)
	assert.Equal(t, "DeviceOnboarding", resp.JobClassName)
	assert.True(t, resp.Enabled)
	assert.True(t, resp.Installed)
}

func TestClient_JobFilter(t *testing.T) {
	gock.New(testURL).Get("extras/jobs/").Reply(200).
		File(path.Join("fixtures", "extras", "jobs_200_1.json"))

	q := &url.Values{}
	q.Set("enabled", "true")

	resp, err := testClient.Extras.JobFilter(q)
	require.NoError(t, err)
	assert.Len(t, resp, 2)
	assert.Equal(t, "Device Onboarding", resp[0].Name)
}

func TestClient_JobAll(t *testing.T) {
	gock.New(testURL).Get("extras/jobs/").Reply(200).
		File(path.Join("fixtures", "extras", "jobs_200_1.json"))

	resp, err := testClient.Extras.JobAll()
	require.NoError(t, err)

	assert.Len(t, resp, 2)
	assert.Equal(t, "Device Onboarding", resp[0].Name)
	assert.Equal(t, "Backup Configurations", resp[1].Name)
}

func TestClient_JobUpdate(t *testing.T) {
	updateData := map[string]any{
		"enabled": false,
	}

	gock.New(testURL).Patch("extras/jobs/" + testJobID + "/").
		JSON(updateData).
		Reply(200).
		File(path.Join("fixtures", "extras", "job_update_200_1.json"))

	id, err := uuid.Parse(testJobID)
	require.NoError(t, err)

	resp, err := testClient.Extras.JobUpdate(id, updateData)
	require.NoError(t, err)

	assert.False(t, resp.Enabled)
	assert.Equal(t, id, resp.ID)
}

func TestClient_JobRun(t *testing.T) {
	runRequest := types.JobRunRequest{
		Data: map[string]any{
			"max_age":       90,
			"cleanup_types": []string{"extras.JobResult"},
		},
		Dryrun: false,
	}

	gock.New(testURL).Post("extras/jobs/" + testJobID + "/run/").
		JSON(runRequest).
		Reply(200).
		File(path.Join("fixtures", "extras", "job_run_200_1.json"))

	id, err := uuid.Parse(testJobID)
	require.NoError(t, err)

	resp, err := testClient.Extras.JobRun(id, runRequest)
	require.NoError(t, err)

	// ScheduledJob should be nil for immediate execution
	assert.Nil(t, resp.ScheduledJob)

	// JobResult should contain the result
	require.NotNil(t, resp.JobResult)
	assert.Equal(t, "Logs Cleanup", resp.JobResult.Name)
	assert.Equal(t, "PENDING", resp.JobResult.Status.Value)
	assert.NotNil(t, resp.JobResult.JobModel)
	assert.Equal(t, "baa7a682-f727-40db-987b-289b8d14b966", resp.JobResult.JobModel.ID)
	assert.NotNil(t, resp.JobResult.User)
	assert.Equal(t, "23b16958-5cc9-4a32-bda2-9b7e47993862", resp.JobResult.User.ID)
}

func TestClient_JobResultGet(t *testing.T) {
	gock.New(testURL).Get("extras/job-results/" + testJobResultID + "/").Reply(200).
		File(path.Join("fixtures", "extras", "jobresult_200_1.json"))

	id, err := uuid.Parse(testJobResultID)
	require.NoError(t, err)

	resp, err := testClient.Extras.JobResultGet(id)
	require.NoError(t, err)
	assert.Equal(t, "Logs Cleanup", resp.Name)
	assert.Equal(t, id, resp.ID)
	assert.Equal(t, "SUCCESS", resp.Status.Value)
	assert.NotNil(t, resp.JobModel)
	assert.NotNil(t, resp.User)
}

func TestClient_JobResultFilter(t *testing.T) {
	gock.New(testURL).Get("extras/job-results/").Reply(200).
		File(path.Join("fixtures", "extras", "jobresults_200_1.json"))

	q := &url.Values{}
	q.Set("status", "SUCCESS")

	resp, err := testClient.Extras.JobResultFilter(q)
	require.NoError(t, err)
	assert.Len(t, resp, 2)
	assert.Equal(t, "Logs Cleanup", resp[0].Name)
}

func TestClient_JobResultAll(t *testing.T) {
	gock.New(testURL).Get("extras/job-results/").Reply(200).
		File(path.Join("fixtures", "extras", "jobresults_200_1.json"))

	resp, err := testClient.Extras.JobResultAll()
	require.NoError(t, err)

	assert.Len(t, resp, 2)
	assert.Equal(t, "Logs Cleanup", resp[0].Name)
	assert.Equal(t, "Device Sync", resp[1].Name)
}

func TestClient_JobResultDelete(t *testing.T) {
	gock.New(testURL).Delete("extras/job-results/" + testJobResultID + "/").Reply(204)

	id, err := uuid.Parse(testJobResultID)
	require.NoError(t, err)

	err = testClient.Extras.JobResultDelete(id)
	require.NoError(t, err)
	assert.True(t, gock.IsDone())
}

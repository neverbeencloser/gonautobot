package nautobot_test

import (
	"net/url"
	"path"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/h2non/gock.v1"
)

const testScheduledJobID = "e1f2a3b4-c5d6-7890-abcd-ef1234567890"

func TestClient_ScheduledJobGet(t *testing.T) {
	gock.New(testURL).Get("extras/scheduled-jobs/" + testScheduledJobID + "/").Reply(200).
		File(path.Join("fixtures", "extras", "scheduledjob_200_1.json"))

	id, err := uuid.Parse(testScheduledJobID)
	require.NoError(t, err)

	resp, err := testClient.Extras.ScheduledJobGet(id)
	require.NoError(t, err)

	expectedLastRunAt := time.Date(2026, 2, 1, 14, 0, 0, 0, time.UTC)
	expectedDateChanged := time.Date(2026, 2, 1, 14, 0, 0, 0, time.UTC)

	assert.Equal(t, "Logs Cleanup", resp.Name)
	assert.Equal(t, id, resp.ID)
	assert.Equal(t, "default", resp.Queue)
	assert.Equal(t, "UTC", resp.TimeZone)
	assert.Equal(t, "nautobot.extras.jobs.run_job", resp.Task)
	assert.Equal(t, "future", resp.Interval)
	assert.Empty(t, resp.Args)
	assert.NotNil(t, resp.Kwargs)
	require.NotNil(t, resp.CeleryKwargs)
	assert.Equal(t, "default", resp.CeleryKwargs["queue"])
	assert.False(t, resp.OneOff)
	assert.True(t, resp.Enabled)
	require.NotNil(t, resp.LastRunAt)
	assert.Equal(t, expectedLastRunAt, resp.LastRunAt.Truncate(time.Second))
	assert.Equal(t, 42, resp.TotalRunCount)
	require.NotNil(t, resp.DateChanged)
	assert.Equal(t, expectedDateChanged, resp.DateChanged.Truncate(time.Second))
	assert.True(t, resp.ApprovalRequired)
	assert.False(t, resp.Approved)
	require.NotNil(t, resp.JobModel)
	assert.Equal(t, "baa7a682-f727-40db-987b-289b8d14b966", resp.JobModel.ID)
	require.NotNil(t, resp.JobQueue)
	assert.Equal(t, "a1b2c3d4-e5f6-7890-abcd-ef1234567890", resp.JobQueue.ID)
	assert.Equal(t, "https://demo.nautobot.com/api/extras/job-queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890/", resp.JobQueue.URL)
	require.NotNil(t, resp.User)
	assert.Equal(t, "admin", resp.User.Username)
}

func TestClient_ScheduledJobFilter(t *testing.T) {
	gock.New(testURL).Get("extras/scheduled-jobs/").Reply(200).
		File(path.Join("fixtures", "extras", "scheduledjobs_200_1.json"))

	q := &url.Values{}
	q.Set("approval_required", "true")

	resp, err := testClient.Extras.ScheduledJobFilter(q)
	require.NoError(t, err)
	assert.Len(t, resp, 2)
	assert.Equal(t, "Logs Cleanup", resp[0].Name)
	assert.Equal(t, "default", resp[0].Queue)
	assert.Equal(t, "UTC", resp[0].TimeZone)
	assert.True(t, resp[0].Enabled)
	assert.Equal(t, 42, resp[0].TotalRunCount)
	require.NotNil(t, resp[0].JobQueue)
	assert.Equal(t, "Device Sync", resp[1].Name)
	assert.Equal(t, "K8s", resp[1].Queue)
	assert.Equal(t, "America/New_York", resp[1].TimeZone)
	assert.Nil(t, resp[1].JobQueue)
	assert.Nil(t, resp[1].LastRunAt)
	assert.Nil(t, resp[1].DateChanged)
	assert.Equal(t, 0, resp[1].TotalRunCount)
}

func TestClient_ScheduledJobAll(t *testing.T) {
	gock.New(testURL).Get("extras/scheduled-jobs/").Reply(200).
		File(path.Join("fixtures", "extras", "scheduledjobs_200_1.json"))

	resp, err := testClient.Extras.ScheduledJobAll()
	require.NoError(t, err)
	assert.Len(t, resp, 2)
	assert.Equal(t, "Logs Cleanup", resp[0].Name)
	assert.Equal(t, "Device Sync", resp[1].Name)
}

func TestClient_ScheduledJobDelete(t *testing.T) {
	gock.New(testURL).Delete("extras/scheduled-jobs/" + testScheduledJobID + "/").Reply(204)

	id, err := uuid.Parse(testScheduledJobID)
	require.NoError(t, err)

	err = testClient.Extras.ScheduledJobDelete(id)
	require.NoError(t, err)
	assert.True(t, gock.IsDone())
}

package nautobot_test

import (
	"net/url"
	"path"
	"testing"

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
	assert.Equal(t, "Logs Cleanup", resp.Name)
	assert.Equal(t, id, resp.ID)
	assert.True(t, resp.ApprovalRequired)
	assert.False(t, resp.Approved)
	require.NotNil(t, resp.JobModel)
	assert.Equal(t, "baa7a682-f727-40db-987b-289b8d14b966", resp.JobModel.ID)
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
	assert.Equal(t, "Device Sync", resp[1].Name)
}

func TestClient_ScheduledJobAll(t *testing.T) {
	gock.New(testURL).Get("extras/scheduled-jobs/").Reply(200).
		File(path.Join("fixtures", "extras", "scheduledjobs_200_1.json"))

	resp, err := testClient.Extras.ScheduledJobAll()
	require.NoError(t, err)
	assert.Len(t, resp, 2)
}

func TestClient_ScheduledJobDelete(t *testing.T) {
	gock.New(testURL).Delete("extras/scheduled-jobs/" + testScheduledJobID + "/").Reply(204)

	id, err := uuid.Parse(testScheduledJobID)
	require.NoError(t, err)

	err = testClient.Extras.ScheduledJobDelete(id)
	require.NoError(t, err)
	assert.True(t, gock.IsDone())
}

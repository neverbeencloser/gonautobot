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

const (
	testJobID = "f1a2b3c4-d5e6-7890-abcd-ef1234567890"
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

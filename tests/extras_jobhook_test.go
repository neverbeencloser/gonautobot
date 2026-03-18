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

const testJobHookID = "a1b2c3d4-e5f6-7890-abcd-ef1234567890"

func TestClient_JobHookGet(t *testing.T) {
	gock.New(testURL).Get("extras/job-hooks/" + testJobHookID + "/").Reply(200).
		File(path.Join("fixtures", "extras", "jobhook_200_1.json"))

	id, err := uuid.Parse(testJobHookID)
	require.NoError(t, err)

	resp, err := testClient.Extras.JobHookGet(id)
	require.NoError(t, err)
	assert.Equal(t, "Sync device on change", resp.Name)
	assert.Equal(t, id, resp.ID)
	assert.True(t, resp.Enabled)
	assert.True(t, resp.TypeCreate)
	assert.False(t, resp.TypeDelete)
	assert.True(t, resp.TypeUpdate)
	require.NotNil(t, resp.Job)
	assert.Equal(t, "baa7a682-f727-40db-987b-289b8d14b966", resp.Job.ID)
	assert.Equal(t, []string{"dcim.device"}, resp.ContentTypes)
}

func TestClient_JobHookFilter(t *testing.T) {
	gock.New(testURL).Get("extras/job-hooks/").Reply(200).
		File(path.Join("fixtures", "extras", "jobhooks_200_1.json"))

	q := &url.Values{}
	q.Set("enabled", "true")

	resp, err := testClient.Extras.JobHookFilter(q)
	require.NoError(t, err)
	assert.Len(t, resp, 2)
	assert.Equal(t, "Sync device on change", resp[0].Name)
	assert.Equal(t, "IPAM hook", resp[1].Name)
}

func TestClient_JobHookAll(t *testing.T) {
	gock.New(testURL).Get("extras/job-hooks/").Reply(200).
		File(path.Join("fixtures", "extras", "jobhooks_200_1.json"))

	resp, err := testClient.Extras.JobHookAll()
	require.NoError(t, err)
	assert.Len(t, resp, 2)
}

func TestClient_JobHookCreate(t *testing.T) {
	newHook := types.NewJobHook{
		Name:         "Sync device on change",
		ContentTypes: []string{"dcim.device"},
		Job:          "baa7a682-f727-40db-987b-289b8d14b966",
	}
	enabled := true
	tc, tu := true, true
	td := false
	newHook.Enabled = &enabled
	newHook.TypeCreate = &tc
	newHook.TypeDelete = &td
	newHook.TypeUpdate = &tu

	gock.New(testURL).Post("extras/job-hooks/").
		JSON(newHook).
		Reply(201).
		File(path.Join("fixtures", "extras", "jobhook_create_201_1.json"))

	resp, err := testClient.Extras.JobHookCreate(newHook)
	require.NoError(t, err)
	assert.Equal(t, newHook.Name, resp.Name)
	assert.Equal(t, uuid.MustParse(testJobHookID), resp.ID)
}

func TestClient_JobHookUpdate(t *testing.T) {
	updateData := map[string]any{
		"enabled":        false,
		"content_types":  []string{"dcim.device", "dcim.interface"},
	}

	gock.New(testURL).Patch("extras/job-hooks/" + testJobHookID + "/").
		JSON(updateData).
		Reply(200).
		File(path.Join("fixtures", "extras", "jobhook_update_200_1.json"))

	id, err := uuid.Parse(testJobHookID)
	require.NoError(t, err)

	resp, err := testClient.Extras.JobHookUpdate(id, updateData)
	require.NoError(t, err)
	assert.False(t, resp.Enabled)
	assert.Len(t, resp.ContentTypes, 2)
}

func TestClient_JobHookDelete(t *testing.T) {
	gock.New(testURL).Delete("extras/job-hooks/" + testJobHookID + "/").Reply(204)

	id, err := uuid.Parse(testJobHookID)
	require.NoError(t, err)

	err = testClient.Extras.JobHookDelete(id)
	require.NoError(t, err)
	assert.True(t, gock.IsDone())
}

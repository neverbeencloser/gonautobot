package nautobot_test

import (
	"net/url"
	"path"
	"testing"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/extras"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/h2non/gock.v1"
)

const (
	testStatusID = "f46a5d99-26e0-4cf0-82d2-55f5a05cda13"
)

func TestClient_StatusGet(t *testing.T) {
	gock.New(testURL).Get("extras/statuses/" + testStatusID + "/").Reply(200).
		File(path.Join("fixtures", "extras", "status_200_1.json"))

	id, err := uuid.Parse(testStatusID)
	require.NoError(t, err)

	resp, err := testClient.Extras.StatusGet(id)
	require.NoError(t, err)
	assert.Equal(t, "Staging", resp.Name)
	assert.Equal(t, id, resp.ID)
}

func TestClient_StatusFilter(t *testing.T) {
	gock.New(testURL).Get("extras/statuses/").Reply(200).
		File(path.Join("fixtures", "extras", "statuses_200_2.json"))

	q := &url.Values{}
	q.Set("content_types__ic", "dcim.device")

	resp, err := testClient.Extras.StatusFilter(q)
	require.NoError(t, err)
	assert.Len(t, resp, 7)
	assert.Equal(t, "Active", resp[0].Name)
}

func TestClient_StatusAll(t *testing.T) {
	gock.New(testURL).Get("extras/statuses/").Reply(200).
		File(path.Join("fixtures", "extras", "statuses_200_1.json"))

	resp, err := testClient.Extras.StatusAll()
	require.NoError(t, err)
	assert.Len(t, resp, 22)
	assert.Equal(t, "Active", resp[0].Name)
}

func TestClient_StatusCreate(t *testing.T) {
	newStatus := extras.NewStatus{
		Name:         "Staging",
		ContentTypes: []string{"dcim.location", "dcim.deviceredundancygroup", "dcim.interfaceredundancygroup"},
		Color:        "2196f3",
		Description:  "Location is in the process of being staged",
	}

	gock.New(testURL).Post("extras/statuses/").
		JSON(newStatus).
		Reply(201).
		File(path.Join("fixtures", "extras", "status_create_201_1.json"))

	resp, err := testClient.Extras.StatusCreate(newStatus)
	require.NoError(t, err)
	assert.Equal(t, newStatus.Name, resp.Name)
	assert.Equal(t, newStatus.Color, resp.Color)
	assert.Equal(t, newStatus.Description, resp.Description)
	assert.Equal(t, uuid.MustParse(testStatusID), resp.ID)
}

func TestClient_StatusUpdate(t *testing.T) {
	updateData := map[string]any{
		"description": "description updated",
	}

	gock.New(testURL).Patch("extras/statuses/" + testStatusID + "/").
		JSON(updateData).
		Reply(200).
		File(path.Join("fixtures", "extras", "status_update_200_1.json"))

	id, err := uuid.Parse(testStatusID)
	require.NoError(t, err)

	resp, err := testClient.Extras.StatusUpdate(id, updateData)
	require.NoError(t, err)
	assert.Equal(t, "description updated", resp.Description)
	assert.Equal(t, id, resp.ID)
}

func TestClient_StatusDelete(t *testing.T) {
	statusID := "a706e57a-593d-4012-ba60-8469802009de"
	gock.New(testURL).Delete("extras/statuses/" + statusID + "/").Reply(204)

	id, err := uuid.Parse(statusID)
	require.NoError(t, err)

	err = testClient.Extras.StatusDelete(id)
	require.NoError(t, err)
}

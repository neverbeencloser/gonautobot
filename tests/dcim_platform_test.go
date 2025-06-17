package nautobot_test

import (
	"net/url"
	"path"
	"testing"

	"github.com/google/uuid"
	"github.com/josh-silvas/gonautobot/dcim"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/h2non/gock.v1"
)

const (
	testPlatformID = "b23f67d6-b92d-4447-b18e-beb7bd6adee7"
)

func TestClient_PlatformGet(t *testing.T) {
	defer gock.Off()
	gock.New(testURL).Get("dcim/platforms/" + testPlatformID + "/").Reply(200).
		File(path.Join("fixtures", "dcim", "platform_200_1.json"))

	id, err := uuid.Parse(testPlatformID)
	require.NoError(t, err)

	resp, err := testClient.Dcim.PlatformGet(id)
	require.NoError(t, err)
	assert.Equal(t, "eos", resp.Name)
	assert.Equal(t, id, resp.ID)
}

func TestClient_PlatformFilter(t *testing.T) {
	defer gock.Off()
	gock.New(testURL).Get("dcim/platforms/").MatchParam("name", "eos").Reply(200).
		File(path.Join("fixtures", "dcim", "platforms_200_1.json"))

	q := &url.Values{}
	q.Set("name", "eos")

	resp, err := testClient.Dcim.PlatformFilter(q)
	require.NoError(t, err)
	assert.Len(t, resp, 3)
	assert.Equal(t, "eos", resp[0].Name)
}

func TestClient_PlatformAll(t *testing.T) {
	defer gock.Off()
	gock.New(testURL).Get("dcim/platforms/").Reply(200).
		File(path.Join("fixtures", "dcim", "platforms_200_1.json"))

	resp, err := testClient.Dcim.PlatformAll()
	require.NoError(t, err)

	assert.Len(t, resp, 3)
	assert.Equal(t, "eos", resp[0].Name)
}

func TestClient_PlatformCreate(t *testing.T) {
	defer gock.Off()
	newPlatform := dcim.NewPlatform{
		Name:         "eos",
		Manufacturer: "b832f225-56af-4ff3-8e9a-66ab401c84de", // Manufacturer ID from fixture
	}

	gock.New(testURL).Post("dcim/platforms/").
		JSON(newPlatform).
		Reply(201).
		File(path.Join("fixtures", "dcim", "platform_create_201_1.json"))

	resp, err := testClient.Dcim.PlatformCreate(newPlatform)
	require.NoError(t, err)
	assert.Equal(t, newPlatform.Name, resp.Name)
	assert.Equal(t, uuid.MustParse(testPlatformID), resp.ID)
}

func TestClient_PlatformUpdate(t *testing.T) {
	defer gock.Off()
	updateData := map[string]any{
		"description": "updated description",
	}

	gock.New(testURL).Patch("dcim/platforms/" + testPlatformID + "/").
		JSON(updateData).
		Reply(200).
		File(path.Join("fixtures", "dcim", "platform_update_200_1.json"))

	id, err := uuid.Parse(testPlatformID)
	require.NoError(t, err)

	resp, err := testClient.Dcim.PlatformUpdate(id, updateData)
	require.NoError(t, err)

	assert.Equal(t, "updated description", resp.Description)
	assert.Equal(t, id, resp.ID)
}

func TestClient_PlatformDelete(t *testing.T) {
	defer gock.Off()
	gock.New(testURL).Delete("dcim/platforms/" + testPlatformID + "/").Reply(204)

	id, err := uuid.Parse(testPlatformID)
	require.NoError(t, err)

	err = testClient.Dcim.PlatformDelete(id)
	require.NoError(t, err)
}

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
	testTagID = "91283c9c-e2ba-4cfe-8f9f-ed6dd6e18ba1"
)

func TestClient_GetTag(t *testing.T) {
	gock.New(testURL).Get("extras/tags/91283c9c-e2ba-4cfe-8f9f-ed6dd6e18ba1/").Reply(200).
		File(path.Join("fixtures", "extras", "tag_200_1.json"))

	resp, err := testClient.Extras.TagGet(uuid.MustParse(testTagID))
	require.NoError(t, err)
	assert.Equal(t, "sup-720", resp.Name)
}

func TestClient_GetTags(t *testing.T) {
	gock.New(testURL).Get("extras/tags/").Reply(200).
		File(path.Join("fixtures", "extras", "tags_200_1.json"))

	resp, err := testClient.Extras.TagFilter(&url.Values{"offset": {"50"}})
	require.NoError(t, err)
	assert.Len(t, resp, 1)
}

func TestClient_TagAll(t *testing.T) {
	gock.New(testURL).Get("extras/tags/").Reply(200).
		File(path.Join("fixtures", "extras", "tags_200_1.json"))

	resp, err := testClient.Extras.TagAll()
	require.NoError(t, err)
	assert.Len(t, resp, 1)
}

func TestClient_TagCreate(t *testing.T) {
	newTag := types.NewTag{
		Name:         "sup-720",
		ContentTypes: []string{},
		Color:        "607d8b",
		Description:  "Supervisor Engine 720",
	}

	gock.New(testURL).Post("extras/tags/").
		JSON(newTag).
		Reply(201).
		File(path.Join("fixtures", "extras", "tag_create_201_1.json"))

	resp, err := testClient.Extras.TagCreate(newTag)
	require.NoError(t, err)
	assert.Equal(t, newTag.Name, resp.Name)
	assert.Equal(t, newTag.Color, resp.Color)
	assert.Equal(t, newTag.Description, resp.Description)
	assert.Equal(t, uuid.MustParse(testTagID), resp.ID)
}

func TestClient_TagUpdate(t *testing.T) {
	updateData := map[string]any{
		"description": "updated description",
	}

	gock.New(testURL).Patch("extras/tags/" + testTagID + "/").
		JSON(updateData).
		Reply(200).
		File(path.Join("fixtures", "extras", "tag_update_200_1.json"))

	id, err := uuid.Parse(testTagID)
	require.NoError(t, err)

	resp, err := testClient.Extras.TagUpdate(id, updateData)
	require.NoError(t, err)

	assert.Equal(t, "updated description", resp.Description)
	assert.Equal(t, id, resp.ID)
}

func TestClient_TagDelete(t *testing.T) {
	gock.New(testURL).Delete("extras/tags/" + testTagID + "/").Reply(204)

	id, err := uuid.Parse(testTagID)
	require.NoError(t, err)

	err = testClient.Extras.TagDelete(id)
	require.NoError(t, err)
	assert.True(t, gock.IsDone())
}

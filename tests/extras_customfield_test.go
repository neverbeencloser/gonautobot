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
	testCustomFieldID = "b6e143ff-c980-4c1f-a2b7-70a59042a188"
)

func TestClient_CustomFieldGet(t *testing.T) {
	defer gock.Off()
	gock.New(testURL).Get("extras/custom-fields/" + testCustomFieldID + "/").Reply(200).
		File(path.Join("fixtures", "extras", "customfield_200_1.json"))

	id, err := uuid.Parse(testCustomFieldID)
	require.NoError(t, err)

	resp, err := testClient.Extras.CustomFieldGet(id)
	require.NoError(t, err)
	assert.Equal(t, "Terraform", resp.Label)
	assert.False(t, resp.Default.(bool))
	assert.Equal(t, id, resp.ID)
}

func TestClient_CustomFieldFilter(t *testing.T) {
	defer gock.Off()
	gock.New(testURL).Get("extras/custom-fields/").Reply(200).
		File(path.Join("fixtures", "extras", "customfields_200_1.json"))

	q := &url.Values{}

	resp, err := testClient.Extras.CustomFieldFilter(q)
	require.NoError(t, err)
	assert.Len(t, resp, 3)
	assert.Equal(t, resp[2].ID, uuid.MustParse(testCustomFieldID))
}

func TestClient_CustomFieldAll(t *testing.T) {
	defer gock.Off()
	gock.New(testURL).Get("extras/custom-fields/").Reply(200).
		File(path.Join("fixtures", "extras", "customfields_200_1.json"))

	resp, err := testClient.Extras.CustomFieldAll()
	require.NoError(t, err)

	assert.Len(t, resp, 3)
	assert.Equal(t, "Custom Thing1", resp[0].Label)
	assert.Equal(t, "Terraform", resp[2].Label)
	// check Default fields
	assert.Nil(t, resp[0].Default)
	assert.Equal(t, map[string]any{"stuff": "things"}, resp[1].Default.(map[string]any))
	assert.False(t, resp[2].Default.(bool))
}

func TestClient_CustomFieldCreate(t *testing.T) {
	defer gock.Off()
	newCf := types.NewCustomField{
		ContentTypes: []string{"dcim.device", "dcim.location", "dcim.locationtype", "tenancy.tenant"},
		Key:          "terraform",
		Label:        "Terraform",
		Type:         "boolean",
	}

	gock.New(testURL).Post("extras/custom-fields/").
		JSON(newCf).
		Reply(201).
		File(path.Join("fixtures", "extras", "customfield_create_200_1.json"))

	resp, err := testClient.Extras.CustomFieldCreate(newCf)
	require.NoError(t, err)

	assert.Equal(t, "Terraform", resp.Label)
	assert.Equal(t, "terraform", resp.Key)
	assert.Equal(t, uuid.MustParse(testCustomFieldID), resp.ID)
}

func TestClient_CustomFieldUpdate(t *testing.T) {
	defer gock.Off()
	updateData := map[string]any{
		"description": "updated description",
		"weight":      250,
	}

	gock.New(testURL).Patch("extras/custom-fields/" + testCustomFieldID + "/").
		JSON(updateData).
		Reply(200).
		File(path.Join("fixtures", "extras", "customfield_update_200_1.json"))

	id, err := uuid.Parse(testCustomFieldID)
	require.NoError(t, err)

	resp, err := testClient.Extras.CustomFieldUpdate(id, updateData)
	require.NoError(t, err)

	assert.Equal(t, "updated description", resp.Description)
	assert.Equal(t, 250, resp.Weight)
	assert.Equal(t, id, resp.ID)
}

func TestClient_CustomFieldDelete(t *testing.T) {
	defer gock.Off()
	gock.New(testURL).Delete("extras/custom-fields/" + testCustomFieldID + "/").Reply(204)

	id, err := uuid.Parse(testCustomFieldID)
	require.NoError(t, err)

	err = testClient.Extras.CustomFieldDelete(id)
	require.NoError(t, err)
}

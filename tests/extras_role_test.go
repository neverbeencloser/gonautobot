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
	testRoleID = "0eb7e306-b960-4199-8040-08ce947f2e3a"
)

func TestClient_RoleGet(t *testing.T) {
	gock.New(testURL).Get("extras/roles/" + testRoleID + "/").Reply(200).
		File(path.Join("fixtures", "extras", "role_200_1.json"))

	id, err := uuid.Parse(testRoleID)
	require.NoError(t, err)

	resp, err := testClient.Extras.RoleGet(id)
	require.NoError(t, err)
	assert.Equal(t, "Compute", resp.Name)
	assert.Equal(t, id, resp.ID)
}

func TestClient_RoleFilter(t *testing.T) {
	gock.New(testURL).Get("extras/roles/").Reply(200).
		File(path.Join("fixtures", "extras", "roles_200_2.json"))

	q := &url.Values{}
	q.Set("content_types", "dcim.device")

	resp, err := testClient.Extras.RoleFilter(q)
	require.NoError(t, err)
	assert.Len(t, resp, 2)
	assert.Equal(t, "Compute", resp[0].Name)
}

func TestClient_RoleAll(t *testing.T) {
	gock.New(testURL).Get("extras/roles/").Reply(200).
		File(path.Join("fixtures", "extras", "roles_200_1.json"))

	resp, err := testClient.Extras.RoleAll()
	require.NoError(t, err)

	assert.Len(t, resp, 14)
	assert.Equal(t, "Administrative", resp[0].Name)
}

func TestClient_RoleCreate(t *testing.T) {
	newRole := extras.NewRole{
		Name:         "Compute",
		ContentTypes: []string{"dcim.device"},
		Color:        "9e9e9e",
	}

	gock.New(testURL).Post("extras/roles/").
		JSON(newRole).
		Reply(201).
		File(path.Join("fixtures", "extras", "role_create_201_1.json"))

	resp, err := testClient.Extras.RoleCreate(newRole)
	require.NoError(t, err)
	assert.Equal(t, newRole.Name, resp.Name)
	assert.Equal(t, newRole.Color, resp.Color)
	assert.Equal(t, uuid.MustParse(testRoleID), resp.ID)
}

func TestClient_RoleUpdate(t *testing.T) {
	updateData := map[string]any{
		"description": "updated description",
		"weight":      200,
	}

	gock.New(testURL).Patch("extras/roles/" + testRoleID + "/").
		JSON(updateData).
		Reply(200).
		File(path.Join("fixtures", "extras", "role_update_200_1.json"))

	id, err := uuid.Parse(testRoleID)
	require.NoError(t, err)

	resp, err := testClient.Extras.RoleUpdate(id, updateData)
	require.NoError(t, err)

	assert.Equal(t, "updated description", resp.Description)
	assert.Equal(t, 200, *resp.Weight)
	assert.Equal(t, id, resp.ID)
}

func TestClient_RoleDelete(t *testing.T) {
	gock.New(testURL).Delete("extras/roles/" + testRoleID + "/").Reply(204)

	id, err := uuid.Parse(testRoleID)
	require.NoError(t, err)

	err = testClient.Extras.RoleDelete(id)
	require.NoError(t, err)
	assert.True(t, gock.IsDone())
}

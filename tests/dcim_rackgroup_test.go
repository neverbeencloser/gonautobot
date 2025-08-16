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
	testRackGroupID = "7e2c45bf-2d17-45ca-9edf-a0dd3259ed66"
)

func TestClient_RackGroupGet(t *testing.T) {
	gock.New(testURL).Get("dcim/rack-groups/" + testRackGroupID + "/").Reply(200).
		File(path.Join("fixtures", "dcim", "rackgroup_200_1.json"))

	id, err := uuid.Parse(testRackGroupID)
	require.NoError(t, err)

	resp, err := testClient.Dcim.RackGroupGet(id)
	require.NoError(t, err)
	assert.Equal(t, "rackgroup1", resp.Name)
	assert.Equal(t, id, resp.ID)
}

func TestClient_RackGroupFilter(t *testing.T) {
	gock.New(testURL).Get("dcim/rack-groups/").Reply(200).
		File(path.Join("fixtures", "dcim", "rackgroups_200_1.json"))

	q := &url.Values{}
	q.Set("location", "b840f97e-e6e9-4b8b-b7ce-0e30f3424201")

	resp, err := testClient.Dcim.RackGroupFilter(q)
	require.NoError(t, err)
	assert.Len(t, resp, 3)
	assert.Equal(t, "rackgroup1", resp[0].Name)
}

func TestClient_RackGroupAll(t *testing.T) {
	gock.New(testURL).Get("dcim/rack-groups/").Reply(200).
		File(path.Join("fixtures", "dcim", "rackgroups_200_1.json"))

	resp, err := testClient.Dcim.RackGroupAll()
	require.NoError(t, err)

	assert.Len(t, resp, 3)
	assert.Equal(t, "rackgroup1", resp[0].Name)
}

func TestClient_RackGroupCreate(t *testing.T) {
	newRackGroup := types.NewRackGroup{
		Name:     "rackgroup1",
		Location: "b840f97e-e6e9-4b8b-b7ce-0e30f3424201",
	}

	gock.New(testURL).Post("dcim/rack-groups/").
		JSON(newRackGroup).
		Reply(201).
		File(path.Join("fixtures", "dcim", "rackgroup_create_201_1.json"))

	resp, err := testClient.Dcim.RackGroupCreate(newRackGroup)
	require.NoError(t, err)
	assert.Equal(t, newRackGroup.Name, resp.Name)
	assert.Equal(t, uuid.MustParse(testRackGroupID), resp.ID)
}

func TestClient_RackGroupUpdate(t *testing.T) {
	updateData := map[string]any{
		"description": "updated description",
	}

	gock.New(testURL).Patch("dcim/rack-groups/" + testRackGroupID + "/").
		JSON(updateData).
		Reply(200).
		File(path.Join("fixtures", "dcim", "rackgroup_update_200_1.json"))

	id, err := uuid.Parse(testRackGroupID)
	require.NoError(t, err)

	resp, err := testClient.Dcim.RackGroupUpdate(id, updateData)
	require.NoError(t, err)

	assert.Equal(t, "updated description", resp.Description)
	assert.Equal(t, id, resp.ID)
}

func TestClient_RackGroupDelete(t *testing.T) {
	gock.New(testURL).Delete("dcim/rack-groups/" + testRackGroupID + "/").Reply(204)

	id, err := uuid.Parse(testRackGroupID)
	require.NoError(t, err)

	err = testClient.Dcim.RackGroupDelete(id)
	require.NoError(t, err)
	assert.True(t, gock.IsDone())
}

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
	testVLANID = "01acc3fb-d449-4f36-9789-df5503206528"
)

func TestClient_VLANGet(t *testing.T) {
	defer gock.Off()

	gock.New(testURL).Get("ipam/vlans/" + testVLANID + "/").Reply(200).
		File(path.Join("fixtures", "ipam", "vlan_200_1.json"))

	id, err := uuid.Parse(testVLANID)
	require.NoError(t, err)

	resp, err := testClient.Ipam.VLANGet(id)
	require.NoError(t, err)
	assert.Equal(t, "dfw01-103-mgmt", resp.Name)
	assert.Equal(t, id, resp.ID)
}

func TestClient_VLANFilter(t *testing.T) {
	defer gock.Off()

	gock.New(testURL).Get("ipam/vlans/").Reply(200).
		File(path.Join("fixtures", "ipam", "vlans_200_1.json"))

	q := &url.Values{}
	q.Set("name", "dfw01-103-mgmt")

	resp, err := testClient.Ipam.VLANFilter(q)
	require.NoError(t, err)

	assert.NotEmpty(t, resp)
	assert.Len(t, resp, 2)
	assert.Equal(t, resp[0].ID, uuid.MustParse(testVLANID))
}

func TestClient_VLANAll(t *testing.T) {
	defer gock.Off()

	gock.New(testURL).Get("ipam/vlans/").Reply(200).
		File(path.Join("fixtures", "ipam", "vlans_200_1.json"))

	resp, err := testClient.Ipam.VLANAll()
	require.NoError(t, err)

	assert.NotEmpty(t, resp)
	assert.Len(t, resp, 2)
	assert.Equal(t, resp[0].ID, uuid.MustParse(testVLANID))
}

func TestClient_VLANCreate(t *testing.T) {
	defer gock.Off()

	newVLAN := types.NewVLAN{
		Name:        "dfw01-103-mgmt",
		Status:      "9f38bab4-4b47-4e77-b50c-fda62817b2db",
		VID:         99,
		Description: "",
		Role:        "5f965b05-bca7-48a4-b5d0-c3360c102c46",
		Tenant:      "1f7fbd07-111a-4091-81d0-f34db26d961d",
	}

	gock.New(testURL).Post("ipam/vlans/").
		JSON(newVLAN).
		Reply(201).
		File(path.Join("fixtures", "ipam", "vlan_200_1.json"))

	resp, err := testClient.Ipam.VLANCreate(newVLAN)
	require.NoError(t, err)

	assert.Equal(t, newVLAN.Name, resp.Name)
	assert.Equal(t, newVLAN.Description, resp.Description)
	assert.Equal(t, uuid.MustParse(testVLANID), resp.ID)
}

func TestClient_VLANUpdate(t *testing.T) {
	defer gock.Off()

	updateData := map[string]any{
		"description": "updated description",
	}

	gock.New(testURL).Patch("ipam/vlans/" + testVLANID + "/").
		JSON(updateData).
		Reply(200).
		File(path.Join("fixtures", "ipam", "vlan_update_200_1.json"))

	id, err := uuid.Parse(testVLANID)
	require.NoError(t, err)

	resp, err := testClient.Ipam.VLANUpdate(id, updateData)
	require.NoError(t, err)

	assert.Equal(t, "updated description", resp.Description)
	assert.Equal(t, id, resp.ID)
}

func TestClient_VLANDelete(t *testing.T) {
	defer gock.Off()

	gock.New(testURL).Delete("ipam/vlans/" + testVLANID + "/").Reply(204)

	id, err := uuid.Parse(testVLANID)
	require.NoError(t, err)

	err = testClient.Ipam.VLANDelete(id)
	require.NoError(t, err)
	assert.True(t, gock.IsDone())
}

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
	testRackID = "11ed8c56-6b0f-4689-9ca4-134e59860102"
)

func TestClient_RackGet(t *testing.T) {
	gock.New(testURL).Get("dcim/racks/" + testRackID + "/").Reply(200).
		File(path.Join("fixtures", "dcim", "rack_200_1.json"))

	id, err := uuid.Parse(testRackID)
	require.NoError(t, err)

	resp, err := testClient.Dcim.RackGet(id)
	require.NoError(t, err)
	assert.Equal(t, "A12-24", resp.Name)
	assert.Equal(t, id, resp.ID)
}

func TestClient_RackFilter(t *testing.T) {
	gock.New(testURL).Get("dcim/racks/").Reply(200).
		File(path.Join("fixtures", "dcim", "racks_200_1.json"))

	q := &url.Values{}
	q.Set("location", "96af6f0f-c573-48d7-8ee1-f4e10545c14c")

	resp, err := testClient.Dcim.RackFilter(q)
	require.NoError(t, err)
	assert.Len(t, resp, 3)
	assert.Equal(t, "A11-13", resp[0].Name)
}

func TestClient_RackAll(t *testing.T) {
	gock.New(testURL).Get("dcim/racks/").Reply(200).
		File(path.Join("fixtures", "dcim", "racks_200_1.json"))

	resp, err := testClient.Dcim.RackAll()
	require.NoError(t, err)

	assert.Len(t, resp, 3)
	assert.Equal(t, "A11-13", resp[0].Name)
}

func TestClient_RackCreate(t *testing.T) {
	newRack := types.NewRack{
		Location: "96af6f0f-c573-48d7-8ee1-f4e10545c14c",
		Name:     "A12-24",
		Status:   "7069f6d4-e69e-4b7b-a5d1-ac1d4fc59353",
		Width:    19,
		UHeight:  42,
		AssetTag: "asfd23",
		Serial:   "2345",
		Type:     "4-post-frame",
	}

	gock.New(testURL).Post("dcim/racks/").
		JSON(newRack).
		Reply(201).
		File(path.Join("fixtures", "dcim", "rack_create_201_1.json"))

	resp, err := testClient.Dcim.RackCreate(newRack)
	require.NoError(t, err)
	assert.Equal(t, newRack.Name, resp.Name)
	assert.Equal(t, newRack.AssetTag, *resp.AssetTag)
	assert.Equal(t, uuid.MustParse(testRackID), resp.ID)
}

func TestClient_RackUpdate(t *testing.T) {
	updateData := map[string]any{
		"asset_tag": "updated-value",
	}

	gock.New(testURL).Patch("dcim/racks/" + testRackID + "/").
		JSON(updateData).
		Reply(200).
		File(path.Join("fixtures", "dcim", "rack_update_200_1.json"))

	id, err := uuid.Parse(testRackID)
	require.NoError(t, err)

	resp, err := testClient.Dcim.RackUpdate(id, updateData)
	require.NoError(t, err)

	assert.Equal(t, "updated-value", *resp.AssetTag)
	assert.Equal(t, id, resp.ID)
}

func TestClient_RackDelete(t *testing.T) {
	gock.New(testURL).Delete("dcim/racks/" + testRackID + "/").Reply(204)

	id, err := uuid.Parse(testRackID)
	require.NoError(t, err)

	err = testClient.Dcim.RackDelete(id)
	require.NoError(t, err)
	assert.True(t, gock.IsDone())
}

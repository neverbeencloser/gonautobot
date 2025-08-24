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
	testVRFID = "2d1376b7-2eca-4641-ac30-0e6ff7613ebc"
)

func TestClient_VRFGet(t *testing.T) {
	defer gock.Off()

	gock.New(testURL).Get("ipam/vrfs/" + testVRFID + "/").Reply(200).
		File(path.Join("fixtures", "ipam", "vrf_200_1.json"))

	id, err := uuid.Parse(testVRFID)
	require.NoError(t, err)

	resp, err := testClient.Ipam.VRFGet(id)
	require.NoError(t, err)
	assert.Equal(t, "CadetBlue8029", resp.Name)
	assert.Equal(t, id, resp.ID)
}

func TestClient_VRFFilter(t *testing.T) {
	defer gock.Off()

	gock.New(testURL).Get("ipam/vrfs/").Reply(200).
		File(path.Join("fixtures", "ipam", "vrfs_200_1.json"))

	q := &url.Values{}
	q.Set("name", "CadetBlue8029")

	resp, err := testClient.Ipam.VRFFilter(q)
	require.NoError(t, err)

	assert.NotEmpty(t, resp)
	assert.Len(t, resp, 3)
	assert.Equal(t, resp[0].ID, uuid.MustParse(testVRFID))
}

func TestClient_VRFAll(t *testing.T) {
	defer gock.Off()

	gock.New(testURL).Get("ipam/vrfs/").Reply(200).
		File(path.Join("fixtures", "ipam", "vrfs_200_1.json"))

	resp, err := testClient.Ipam.VRFAll()
	require.NoError(t, err)

	assert.NotEmpty(t, resp)
	assert.Len(t, resp, 3)
	assert.Equal(t, resp[0].ID, uuid.MustParse(testVRFID))
}

func TestClient_VRFCreate(t *testing.T) {
	defer gock.Off()

	newVRF := types.NewVRF{
		Name:        "CadetBlue8029",
		Namespace:   "e7bf1338-ea50-487a-ae37-56c939187d9e",
		RD:          "4280249921:43428",
		Description: "Bag house campaign authority. Piece list or Congress. Coach street to only early.\nBook bed think off. Thank enjoy morning.\nTo point machine above suggest structure woman. See nor at more admit.",
	}

	gock.New(testURL).Post("ipam/vrfs/").
		JSON(newVRF).
		Reply(201).
		File(path.Join("fixtures", "ipam", "vrf_200_1.json"))

	resp, err := testClient.Ipam.VRFCreate(newVRF)
	require.NoError(t, err)

	assert.Equal(t, newVRF.Name, resp.Name)
	assert.Equal(t, newVRF.Description, resp.Description)
	assert.Equal(t, uuid.MustParse(testVRFID), resp.ID)
}

func TestClient_VRFUpdate(t *testing.T) {
	defer gock.Off()

	updateData := map[string]any{
		"description": "updated description",
	}

	gock.New(testURL).Patch("ipam/vrfs/" + testVRFID + "/").
		JSON(updateData).
		Reply(200).
		File(path.Join("fixtures", "ipam", "vrf_200_update_1.json"))

	id, err := uuid.Parse(testVRFID)
	require.NoError(t, err)

	resp, err := testClient.Ipam.VRFUpdate(id, updateData)
	require.NoError(t, err)

	assert.Equal(t, "updated description.", resp.Description)
	assert.Equal(t, id, resp.ID)
}

func TestClient_VRFDelete(t *testing.T) {
	defer gock.Off()

	gock.New(testURL).Delete("ipam/vrfs/" + testVRFID + "/").Reply(204)

	id, err := uuid.Parse(testVRFID)
	require.NoError(t, err)

	err = testClient.Ipam.VRFDelete(id)
	require.NoError(t, err)
	assert.True(t, gock.IsDone())
}
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
	testDeviceFamilyID = "8985db57-2dcd-4bce-8129-5f86e3a65174"
)

func TestClient_DeviceFamilyGet(t *testing.T) {
	defer gock.Off()
	gock.New(testURL).Get("dcim/device-families/" + testDeviceFamilyID + "/").Reply(200).
		File(path.Join("fixtures", "dcim", "devicefamily_200_1.json"))

	id, err := uuid.Parse(testDeviceFamilyID)
	require.NoError(t, err)

	resp, err := testClient.Dcim.DeviceFamilyGet(id)
	require.NoError(t, err)
	assert.Equal(t, "Family1", resp.Name)
	assert.Equal(t, id, resp.ID)
	assert.True(t, gock.IsDone())
}

func TestClient_DeviceFamilyFilter(t *testing.T) {
	defer gock.Off()
	gock.New(testURL).Get("dcim/device-families/").MatchParam("name", "Family1").Reply(200).
		File(path.Join("fixtures", "dcim", "devicefamilies_200_1.json"))

	q := &url.Values{}
	q.Set("name", "Family1")

	resp, err := testClient.Dcim.DeviceFamilyFilter(q)
	require.NoError(t, err)
	assert.Len(t, resp, 2)
	assert.Equal(t, "Family1", resp[0].Name)
	assert.True(t, gock.IsDone())
}

func TestClient_DeviceFamilyAll(t *testing.T) {
	defer gock.Off()
	gock.New(testURL).Get("dcim/device-families/").Reply(200).
		File(path.Join("fixtures", "dcim", "devicefamilies_200_1.json"))

	resp, err := testClient.Dcim.DeviceFamilyAll()
	require.NoError(t, err)

	assert.Len(t, resp, 2)
	assert.Equal(t, "Family1", resp[0].Name)
	assert.True(t, gock.IsDone())
}

func TestClient_DeviceFamilyCreate(t *testing.T) {
	defer gock.Off()
	newDeviceFamily := types.NewDeviceFamily{
		Name:        "Family1",
		Description: "A family for devices",
	}

	gock.New(testURL).Post("dcim/device-families/").
		JSON(newDeviceFamily).
		Reply(201).
		File(path.Join("fixtures", "dcim", "devicefamily_create_201_1.json"))

	resp, err := testClient.Dcim.DeviceFamilyCreate(newDeviceFamily)
	require.NoError(t, err)
	assert.Equal(t, newDeviceFamily.Name, resp.Name)
	assert.Equal(t, newDeviceFamily.Description, resp.Description)
	assert.Equal(t, uuid.MustParse(testDeviceFamilyID), resp.ID)
	assert.True(t, gock.IsDone())
}

func TestClient_DeviceFamilyUpdate(t *testing.T) {
	defer gock.Off()
	updateData := map[string]any{
		"description": "updated description",
	}

	gock.New(testURL).Patch("dcim/device-families/" + testDeviceFamilyID + "/").
		JSON(updateData).
		Reply(200).
		File(path.Join("fixtures", "dcim", "devicefamily_update_200_1.json"))

	id, err := uuid.Parse(testDeviceFamilyID)
	require.NoError(t, err)

	resp, err := testClient.Dcim.DeviceFamilyUpdate(id, updateData)
	require.NoError(t, err)

	assert.Equal(t, "updated description", resp.Description)
	assert.Equal(t, id, resp.ID)
	assert.True(t, gock.IsDone())
}

func TestClient_DeviceFamilyDelete(t *testing.T) {
	defer gock.Off()
	gock.New(testURL).Delete("dcim/device-families/" + testDeviceFamilyID + "/").Reply(204)

	id, err := uuid.Parse(testDeviceFamilyID)
	require.NoError(t, err)

	err = testClient.Dcim.DeviceFamilyDelete(id)
	require.NoError(t, err)
	assert.True(t, gock.IsDone())
}

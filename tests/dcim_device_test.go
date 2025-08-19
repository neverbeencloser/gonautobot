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
	testDeviceID = "89b2ac3b-1853-4eeb-9ea6-6a081999bd3c"
)

func TestClient_GetDevice(t *testing.T) {
	deviceID := uuid.MustParse(testDeviceID)
	gock.New(testURL).Get("dcim/devices/" + testDeviceID + "/").Reply(200).
		File(path.Join("fixtures", "dcim", "device_200_1.json"))

	resp, err := testClient.Dcim.DeviceGet(deviceID)
	require.NoError(t, err)
	assert.Equal(t, "ams01-dist-01", resp.Name)
	assert.Equal(t, deviceID, resp.ID)
}

func TestClient_GetDevices(t *testing.T) {
	gock.New(testURL).Get("dcim/devices/").Reply(200).
		File(path.Join("fixtures", "dcim", "devices_200_1.json"))

	resp, err := testClient.Dcim.DeviceFilter(&url.Values{"offset": {"50"}})
	require.NoError(t, err)
	assert.Len(t, resp, 2)
}

func TestClient_DeviceAll(t *testing.T) {
	gock.New(testURL).Get("dcim/devices/").Reply(200).
		File(path.Join("fixtures", "dcim", "devices_200_1.json"))

	resp, err := testClient.Dcim.DeviceAll()
	require.NoError(t, err)

	assert.Len(t, resp, 2)
	assert.Equal(t, "aaa00-comptegateway-01", resp[0].Name)
}

func TestClient_DeviceCreate(t *testing.T) {
	newDevice := types.NewDevice{
		Name:       "ams01-dist-01",
		Role:       "40567487-6328-4dac-b7b5-b789d1154bf0",
		Status:     "9f38bab4-4b47-4e77-b50c-fda62817b2db",
		DeviceType: "4bf23e23-4eb1-4fae-961c-edd6f8cbaaf1",
		Location:   "9e39051b-e968-4016-b0cf-63a5607375de",
		Tenant:     "1f7fbd07-111a-4091-81d0-f34db26d961d",
		Platform:   "aa07ca99-b973-4870-9b44-e1ea48c23cc9",
	}

	gock.New(testURL).Post("dcim/devices/").
		Reply(201).
		File(path.Join("fixtures", "dcim", "device_200_1.json"))

	resp, err := testClient.Dcim.DeviceCreate(newDevice)
	require.NoError(t, err)
	assert.Equal(t, newDevice.Name, resp.Name)
	assert.Equal(t, uuid.MustParse(testDeviceID), resp.ID)
}

func TestClient_DeviceUpdate(t *testing.T) {
	updateData := map[string]any{
		"comments": "updated comments",
	}

	gock.New(testURL).Patch("dcim/devices/" + testDeviceID + "/").
		Reply(200).
		File(path.Join("fixtures", "dcim", "device_updated_200_1.json"))

	deviceID := uuid.MustParse(testDeviceID)
	resp, err := testClient.Dcim.DeviceUpdate(deviceID, updateData)
	require.NoError(t, err)

	assert.Equal(t, deviceID, resp.ID)
	assert.Equal(t, "updated comments", resp.Comments)
}

func TestClient_DeviceDelete(t *testing.T) {
	gock.New(testURL).Delete("dcim/devices/" + testDeviceID + "/").Reply(204)

	deviceID := uuid.MustParse(testDeviceID)
	err := testClient.Dcim.DeviceDelete(deviceID)
	require.NoError(t, err)
}

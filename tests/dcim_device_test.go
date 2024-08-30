package nautobot_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/url"
	"path"
	"testing"

	"gopkg.in/h2non/gock.v1"
)

func TestClient_GetDevice(t *testing.T) {
	gock.New(testURL).Get("dcim/devices/89b2ac3b-1853-4eeb-9ea6-6a081999bd3c/").Reply(200).
		File(path.Join("fixtures", "dcim", "device_200_1.json"))

	resp, err := testClient.Dcim.GetDevice("89b2ac3b-1853-4eeb-9ea6-6a081999bd3c", nil)
	require.NoError(t, err)
	assert.Equal(t, "ams01-dist-01", resp.Name)
}

func TestClient_GetDevices(t *testing.T) {
	gock.New(testURL).Get("dcim/devices/").Reply(200).
		File(path.Join("fixtures", "dcim", "devices_200_1.json"))

	resp, err := testClient.Dcim.GetDevices(&url.Values{"offset": {"50"}})
	require.NoError(t, err)
	assert.Len(t, resp, 2)
}

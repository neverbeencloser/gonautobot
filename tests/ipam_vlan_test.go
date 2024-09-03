package nautobot_test

import (
	"github.com/stretchr/testify/require"
	"net/url"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestClient_GetVLAN(t *testing.T) {
	gock.New(testURL).Get("ipam/vlans/01acc3fb-d449-4f36-9789-df5503206528/").Reply(200).
		File(path.Join("fixtures", "ipam", "vlan_200_1.json"))

	resp, err := testClient.Ipam.VLANGet("01acc3fb-d449-4f36-9789-df5503206528")
	require.NoError(t, err)
	assert.Equal(t, 99, resp.VID)
}

func TestClient_GetVLANs(t *testing.T) {
	gock.New(testURL).Get("ipam/vlans/").Reply(200).
		File(path.Join("fixtures", "ipam", "vlans_200_1.json"))

	resp, err := testClient.Ipam.VLANFilter(&url.Values{"offset": {"50"}})
	require.NoError(t, err)
	assert.Len(t, resp, 2)
}

package nautobot_test

import (
	"net/url"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"gopkg.in/h2non/gock.v1"
)

func TestClient_GetInterface(t *testing.T) {
	gock.New(testURL).Get("dcim/interfaces/1ce0df0c-5221-4b4e-a7e5-d52810828041/").Reply(200).
		File(path.Join("fixtures", "dcim", "interface_200_1.json"))

	resp, err := testClient.Dcim.InterfaceGet("1ce0df0c-5221-4b4e-a7e5-d52810828041")
	require.NoError(t, err)
	assert.Equal(t, "Ethernet1", resp.Name)
}

func TestClient_GetInterfaces(t *testing.T) {
	gock.New(testURL).Get("dcim/interfaces/").Reply(200).
		File(path.Join("fixtures", "dcim", "interfaces_200_1.json"))

	resp, err := testClient.Dcim.InterfaceFilter(&url.Values{"offset": {"50"}})
	require.NoError(t, err)
	assert.Len(t, resp, 2)
}

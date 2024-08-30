package nautobot_test

import (
	"github.com/stretchr/testify/require"
	"net/url"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestClient_GetIPAddress(t *testing.T) {
	gock.New(testURL).Get("ipam/ip-addresses/50476c26-4349-4669-bc40-fa3baedd72fb/").Reply(200).
		File(path.Join("fixtures", "ipam", "ip_address_200_1.json"))

	resp, err := testClient.Ipam.GetIPAddress("50476c26-4349-4669-bc40-fa3baedd72fb", nil)
	require.NoError(t, err)
	assert.Equal(t, "2.1.1.1/32", resp.Address)
}

func TestClient_GetIPAddresses(t *testing.T) {
	gock.New(testURL).Get("ipam/ip-addresses/").Reply(200).
		File(path.Join("fixtures", "ipam", "ip_addresses_200_1.json"))

	resp, err := testClient.Ipam.GetIPAddresses(&url.Values{"offset": {"50"}})
	require.NoError(t, err)
	assert.Len(t, resp, 2)
}

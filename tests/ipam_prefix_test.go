package nautobot_test

import (
	"github.com/josh-silvas/gonautobot/ipam"
	"github.com/stretchr/testify/require"
	"net/url"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestClient_GetPrefix(t *testing.T) {
	gock.New(testURL).Get("ipam/prefixes/900326be-92a5-4392-8178-929256081580/").Reply(200).
		File(path.Join("fixtures", "ipam", "prefix_200_1.json"))

	resp, err := testClient.Ipam.PrefixGet("900326be-92a5-4392-8178-929256081580")
	require.NoError(t, err)
	assert.Equal(t, "10.0.0.0/24", resp.Prefix)
}

func TestClient_GetPrefixes(t *testing.T) {
	gock.New(testURL).Get("ipam/prefixes/").Reply(200).
		File(path.Join("fixtures", "ipam", "prefixes_200_1.json"))

	resp, err := testClient.Ipam.PrefixFilter(&url.Values{"offset": {"50"}})
	require.NoError(t, err)
	assert.Len(t, resp, 2)
}

func TestClient_GetPrefixAvailableIPs(t *testing.T) {
	gock.New(testURL).Get("ipam/prefixes/900326be-92a5-4392-8178-929256081580/available-ips/").Reply(200).
		File(path.Join("fixtures", "ipam", "prefix_available_ips_200_1.json"))

	resp, err := testClient.Ipam.GetPrefixAvailableIPs("900326be-92a5-4392-8178-929256081580", nil)
	require.NoError(t, err)
	assert.Len(t, resp, 50)
}

func TestClient_GetPrefixAvailablePrefixes(t *testing.T) {
	gock.New(testURL).Get("ipam/prefixes/900326be-92a5-4392-8178-929256081580/available-prefixes/").Reply(200).
		File(path.Join("fixtures", "ipam", "prefix_available_prefixes_200_1.json"))

	resp, err := testClient.Ipam.GetPrefixAvailablePrefixes("900326be-92a5-4392-8178-929256081580", nil)
	require.NoError(t, err)
	assert.Len(t, resp, 1)
}

func TestClient_CreatePrefix(t *testing.T) {
	req := ipam.NewPrefix{
		Description: "Test Prefix",
		IsPool:      false,
		Prefix:      "10.197.27.0/29",
		Status:      "Active",
	}
	gock.New(testURL).Post("ipam/prefixes/").JSON(req).Reply(201).
		File(path.Join("fixtures", "ipam", "prefix_create_response_201_1.json"))

	resp, err := testClient.Ipam.PrefixCreate(&req)
	require.NoError(t, err)
	assert.Equal(t, req.Prefix, resp.Prefix)
	assert.Equal(t, req.Description, resp.Description)
	assert.Equal(t, req.IsPool, resp.IsPool)
}

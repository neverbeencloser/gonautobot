package nautobot_test

import (
	"github.com/stretchr/testify/require"
	"net/url"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestClient_GetRegion(t *testing.T) {
	gock.New(testURL).Get("dcim/regions/e0755d33-90fb-5b4f-b6e1-6b03c50acd2e/").Reply(200).
		File(path.Join("fixtures", "dcim", "region_200_1.json"))

	resp, err := testClient.Dcim.RegionGet("e0755d33-90fb-5b4f-b6e1-6b03c50acd2e")
	require.NoError(t, err)
	assert.Equal(t, "United States", resp.Name)
}

func TestClient_GetRegions(t *testing.T) {
	gock.New(testURL).Get("dcim/regions/").Reply(200).
		File(path.Join("fixtures", "dcim", "regions_200_1.json"))

	resp, err := testClient.Dcim.RegionFilter(&url.Values{"offset": {"50"}})
	require.NoError(t, err)
	assert.Len(t, resp, 1)
}

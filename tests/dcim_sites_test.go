package nautobot_test

import (
	"github.com/stretchr/testify/require"
	"net/url"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestClient_GetSite(t *testing.T) {
	gock.New(testURL).Get("dcim/sites/e0755d33-90fb-5b4f-b6e1-6b03c50acd2e/").Reply(200).
		File(path.Join("fixtures", "dcim", "site_200_1.json"))

	resp, err := testClient.Dcim.GetSite("e0755d33-90fb-5b4f-b6e1-6b03c50acd2e", nil)
	require.NoError(t, err)
	assert.Equal(t, "Casita", resp.Name)
}

func TestClient_GetSites(t *testing.T) {
	gock.New(testURL).Get("dcim/sites/").Reply(200).
		File(path.Join("fixtures", "dcim", "sites_200_1.json"))

	resp, err := testClient.Dcim.GetSites(&url.Values{"offset": {"50"}})
	require.NoError(t, err)
	assert.Len(t, resp, 1)
}

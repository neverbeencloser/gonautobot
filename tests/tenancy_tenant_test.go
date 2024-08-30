package nautobot_test

import (
	"github.com/stretchr/testify/require"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestClient_GetTenants(t *testing.T) {
	gock.New(testURL).Get("tenancy/tenants/").Reply(200).
		File(path.Join("fixtures", "tenancy", "tenants_200_1.json"))

	resp, err := testClient.Tenancy.GetTenants(nil)
	require.NoError(t, err)
	assert.Len(t, resp, 2)
}

func TestClient_GetTenantGroups(t *testing.T) {
	gock.New(testURL).Get("tenancy/tenant-groups/").Reply(200).
		File(path.Join("fixtures", "tenancy", "tenant_groups_200_1.json"))

	resp, err := testClient.Tenancy.GetTenantGroups(nil)
	require.NoError(t, err)
	assert.Len(t, resp, 1)
}

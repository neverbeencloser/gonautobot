package nautobot_test

import (
	"path"
	"testing"

	"github.com/google/uuid"
	"github.com/josh-silvas/gonautobot/tenancy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/h2non/gock.v1"
)

const (
	testTenantID = "fa973d85-f491-4253-9c90-49a62f90d657"
)

func TestClient_TenantsFilter(t *testing.T) {
	gock.New(testURL).Get("tenancy/tenants/").Reply(200).
		File(path.Join("fixtures", "tenancy", "tenants_200_1.json"))

	resp, err := testClient.Tenancy.TenantFilter(nil)
	require.NoError(t, err)
	assert.Len(t, resp, 2)
}

func TestClient_TenantsAll(t *testing.T) {
	gock.New(testURL).Get("tenancy/tenants/").Reply(200).
		File(path.Join("fixtures", "tenancy", "tenants_200_1.json"))

	resp, err := testClient.Tenancy.TenantAll()
	require.NoError(t, err)
	assert.Len(t, resp, 2)
}

func TestClient_GetTenantGroups(t *testing.T) {
	gock.New(testURL).Get("tenancy/tenant-groups/").Reply(200).
		File(path.Join("fixtures", "tenancy", "tenant_groups_200_1.json"))

	resp, err := testClient.Tenancy.TenantGroupFilter(nil)
	require.NoError(t, err)
	assert.Len(t, resp, 1)
}

func TestClient_TenantCreate(t *testing.T) {
	newTenant := tenancy.NewTenant{
		Name: "Customer2",
	}

	gock.New(testURL).Post("tenancy/tenants/").
		JSON(newTenant).
		Reply(201).
		File(path.Join("fixtures", "tenancy", "tenant_create_201_1.json"))

	resp, err := testClient.Tenancy.TenantCreate(newTenant)
	require.NoError(t, err)
	assert.Equal(t, newTenant.Name, resp.Name)
	assert.Equal(t, uuid.MustParse(testTenantID), resp.ID)
}

func TestClient_TenantUpdate(t *testing.T) {
	updateData := map[string]any{
		"description": "updated description",
		"comments":    "updated comments",
	}

	gock.New(testURL).Patch("tenancy/tenants/" + testTenantID + "/").
		JSON(updateData).
		Reply(200).
		File(path.Join("fixtures", "tenancy", "tenant_update_200_1.json"))

	id, err := uuid.Parse(testTenantID)
	require.NoError(t, err)

	resp, err := testClient.Tenancy.TenantUpdate(id, updateData)
	require.NoError(t, err)

	assert.Equal(t, "updated description", resp.Description)
	assert.Equal(t, id, resp.ID)
}

func TestClient_TenantDelete(t *testing.T) {
	gock.New(testURL).Delete("tenancy/tenants/" + testTenantID + "/").Reply(204)

	id, err := uuid.Parse(testTenantID)
	require.NoError(t, err)

	err = testClient.Tenancy.TenantDelete(id)
	require.NoError(t, err)
	assert.True(t, gock.IsDone())
}

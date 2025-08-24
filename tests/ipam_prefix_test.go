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
	testPrefixID = "900326be-92a5-4392-8178-929256081580"
)

func TestClient_PrefixGet(t *testing.T) {
	defer gock.Off()

	gock.New(testURL).Get("ipam/prefixes/" + testPrefixID + "/").Reply(200).
		File(path.Join("fixtures", "ipam", "prefix_200_1.json"))

	id, err := uuid.Parse(testPrefixID)
	require.NoError(t, err)

	resp, err := testClient.Ipam.PrefixGet(id)
	require.NoError(t, err)
	assert.Equal(t, "10.0.0.0/24", resp.Prefix)
	assert.Equal(t, id, resp.ID)
}

func TestClient_PrefixFilter(t *testing.T) {
	defer gock.Off()

	gock.New(testURL).Get("ipam/prefixes/").Reply(200).
		File(path.Join("fixtures", "ipam", "prefixes_200_1.json"))

	q := &url.Values{}
	q.Set("offset", "50")

	resp, err := testClient.Ipam.PrefixFilter(q)
	require.NoError(t, err)

	assert.NotEmpty(t, resp)
	assert.Len(t, resp, 2)
}

func TestClient_PrefixAll(t *testing.T) {
	defer gock.Off()

	gock.New(testURL).Get("ipam/prefixes/").Reply(200).
		File(path.Join("fixtures", "ipam", "prefixes_200_1.json"))

	resp, err := testClient.Ipam.PrefixAll()
	require.NoError(t, err)

	assert.NotEmpty(t, resp)
	assert.Len(t, resp, 2)
}

func TestClient_PrefixCreate(t *testing.T) {
	defer gock.Off()

	newPrefix := types.NewPrefix{
		Description: "Test Prefix",
		Prefix:      "10.197.27.0/29",
		Status:      "Active",
	}

	gock.New(testURL).Post("ipam/prefixes/").
		JSON(newPrefix).
		Reply(201).
		File(path.Join("fixtures", "ipam", "prefix_create_response_201_1.json"))

	resp, err := testClient.Ipam.PrefixCreate(newPrefix)
	require.NoError(t, err)

	assert.Equal(t, newPrefix.Prefix, resp.Prefix)
	assert.Equal(t, newPrefix.Description, resp.Description)
}

func TestClient_PrefixUpdate(t *testing.T) {
	defer gock.Off()

	updateData := map[string]any{
		"description": "updated description",
	}

	gock.New(testURL).Patch("ipam/prefixes/" + testPrefixID + "/").
		JSON(updateData).
		Reply(200).
		File(path.Join("fixtures", "ipam", "prefix_200_update_1.json"))

	id, err := uuid.Parse(testPrefixID)
	require.NoError(t, err)

	resp, err := testClient.Ipam.PrefixUpdate(id, updateData)
	require.NoError(t, err)

	assert.Equal(t, "updated description", resp.Description)
	assert.Equal(t, id, resp.ID)
}

func TestClient_PrefixDelete(t *testing.T) {
	defer gock.Off()

	gock.New(testURL).Delete("ipam/prefixes/" + testPrefixID + "/").Reply(204)

	id, err := uuid.Parse(testPrefixID)
	require.NoError(t, err)

	err = testClient.Ipam.PrefixDelete(id)
	require.NoError(t, err)
	assert.True(t, gock.IsDone())
}

func TestClient_PrefixAvailableIPs(t *testing.T) {
	defer gock.Off()

	gock.New(testURL).Get("ipam/prefixes/" + testPrefixID + "/available-ips/").Reply(200).
		File(path.Join("fixtures", "ipam", "prefix_available_ips_200_1.json"))

	id, err := uuid.Parse(testPrefixID)
	require.NoError(t, err)

	resp, err := testClient.Ipam.GetPrefixAvailableIPs(id, nil)
	require.NoError(t, err)
	assert.Len(t, resp, 50)
}

func TestClient_PrefixAvailablePrefixes(t *testing.T) {
	defer gock.Off()

	gock.New(testURL).Get("ipam/prefixes/" + testPrefixID + "/available-prefixes/").Reply(200).
		File(path.Join("fixtures", "ipam", "prefix_available_prefixes_200_1.json"))

	id, err := uuid.Parse(testPrefixID)
	require.NoError(t, err)

	resp, err := testClient.Ipam.GetPrefixAvailablePrefixes(id, nil)
	require.NoError(t, err)
	assert.Len(t, resp, 1)
}

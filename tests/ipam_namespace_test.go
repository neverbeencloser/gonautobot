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
	testNamespaceID = "e7bf1338-ea50-487a-ae37-56c939187d9e"
)

func TestClient_NamespaceGet(t *testing.T) {
	defer gock.Off()

	gock.New(testURL).Get("ipam/namespaces/" + testNamespaceID + "/").Reply(200).
		File(path.Join("fixtures", "ipam", "namespace_200_1.json"))

	id, err := uuid.Parse(testNamespaceID)
	require.NoError(t, err)

	resp, err := testClient.Ipam.NamespaceGet(id)
	require.NoError(t, err)
	assert.Equal(t, "Art old radio.", resp.Name)
	assert.Equal(t, id, resp.ID)
}

func TestClient_NamespaceFilter(t *testing.T) {
	defer gock.Off()

	gock.New(testURL).Get("ipam/namespaces/").Reply(200).
		File(path.Join("fixtures", "ipam", "namespaces_200_1.json"))

	q := &url.Values{}
	q.Set("name", "Art old radio.")

	resp, err := testClient.Ipam.NamespaceFilter(q)
	require.NoError(t, err)

	assert.NotEmpty(t, resp)
	assert.Len(t, resp, 3)
	assert.Equal(t, resp[0].ID, uuid.MustParse(testNamespaceID))
}

func TestClient_NamespaceAll(t *testing.T) {
	defer gock.Off()

	gock.New(testURL).Get("ipam/namespaces/").Reply(200).
		File(path.Join("fixtures", "ipam", "namespaces_200_1.json"))

	resp, err := testClient.Ipam.NamespaceAll()
	require.NoError(t, err)

	assert.NotEmpty(t, resp)
	assert.Len(t, resp, 3)
	assert.Equal(t, resp[0].ID, uuid.MustParse(testNamespaceID))
}

func TestClient_NamespaceCreate(t *testing.T) {
	defer gock.Off()

	newNamespace := types.NewNamespace{
		Name:        "Art old radio.",
		Description: "",
	}

	gock.New(testURL).Post("ipam/namespaces/").
		JSON(newNamespace).
		Reply(201).
		File(path.Join("fixtures", "ipam", "namespace_create_201_1.json"))

	resp, err := testClient.Ipam.NamespaceCreate(newNamespace)
	require.NoError(t, err)

	assert.Equal(t, newNamespace.Name, resp.Name)
	assert.Equal(t, newNamespace.Description, resp.Description)
	assert.Equal(t, uuid.MustParse(testNamespaceID), resp.ID)
}

func TestClient_NamespaceUpdate(t *testing.T) {
	defer gock.Off()

	updateData := map[string]any{
		"description": "updated description",
	}

	gock.New(testURL).Patch("ipam/namespaces/" + testNamespaceID + "/").
		JSON(updateData).
		Reply(200).
		File(path.Join("fixtures", "ipam", "namespace_update_200_1.json"))

	id, err := uuid.Parse(testNamespaceID)
	require.NoError(t, err)

	resp, err := testClient.Ipam.NamespaceUpdate(id, updateData)
	require.NoError(t, err)

	assert.Equal(t, "updated description", resp.Description)
	assert.Equal(t, id, resp.ID)
}

func TestClient_NamespaceDelete(t *testing.T) {
	defer gock.Off()

	gock.New(testURL).Delete("ipam/namespaces/" + testNamespaceID + "/").Reply(204)

	id, err := uuid.Parse(testNamespaceID)
	require.NoError(t, err)

	err = testClient.Ipam.NamespaceDelete(id)
	require.NoError(t, err)
	assert.True(t, gock.IsDone())
}

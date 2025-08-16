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
	testManufacturerID = "acc4c522-cf77-46bf-bdf3-d0e4a69b4d1e"
)

func TestClient_ManufacturerGet(t *testing.T) {
	defer gock.Off()
	gock.New(testURL).Get("dcim/manufacturers/" + testManufacturerID + "/").Reply(200).
		File(path.Join("fixtures", "dcim", "manufacturer_200_1.json"))

	id, err := uuid.Parse(testManufacturerID)
	require.NoError(t, err)

	resp, err := testClient.Dcim.ManufacturerGet(id)
	require.NoError(t, err)
	assert.Equal(t, "AWS", resp.Name)
	assert.Equal(t, id, resp.ID)
}

func TestClient_ManufacturerFilter(t *testing.T) {
	defer gock.Off()
	gock.New(testURL).Get("dcim/manufacturers/").MatchParam("name", "AWS").Reply(200).
		File(path.Join("fixtures", "dcim", "manufacturers_200_2.json"))

	q := &url.Values{}
	q.Set("name", "AWS")

	resp, err := testClient.Dcim.ManufacturerFilter(q)
	require.NoError(t, err)

	assert.Len(t, resp, 1)
	assert.Equal(t, "AWS", resp[0].Name)
}

func TestClient_ManufacturerAll(t *testing.T) {
	defer gock.Off()
	gock.New(testURL).Get("dcim/manufacturers/").Reply(200).
		File(path.Join("fixtures", "dcim", "manufacturers_200_1.json"))

	resp, err := testClient.Dcim.ManufacturerAll()
	require.NoError(t, err)

	assert.Len(t, resp, 2)
	assert.Equal(t, "AWS", resp[0].Name)
}

func TestClient_ManufacturerCreate(t *testing.T) {
	defer gock.Off()
	newManufacturer := types.NewManufacturer{
		Name:        "AWS",
		Description: "Amazon Web Services",
	}

	gock.New(testURL).Post("dcim/manufacturers/").
		JSON(newManufacturer).
		Reply(201).
		File(path.Join("fixtures", "dcim", "manufacturer_create_200_1.json"))

	resp, err := testClient.Dcim.ManufacturerCreate(newManufacturer)
	require.NoError(t, err)
	assert.Equal(t, newManufacturer.Name, resp.Name)
	assert.Equal(t, newManufacturer.Description, resp.Description)
	assert.Equal(t, uuid.MustParse(testManufacturerID), resp.ID)
}

func TestClient_ManufacturerUpdate(t *testing.T) {
	defer gock.Off()
	updateData := map[string]any{
		"description": "updated description",
	}

	gock.New(testURL).Patch("dcim/manufacturers/" + testManufacturerID + "/").
		JSON(updateData).
		Reply(200).
		File(path.Join("fixtures", "dcim", "manufacturer_update_200_1.json"))

	id, err := uuid.Parse(testManufacturerID)
	require.NoError(t, err)

	resp, err := testClient.Dcim.ManufacturerUpdate(id, updateData)
	require.NoError(t, err)

	assert.Equal(t, "updated description", resp.Description)
	assert.Equal(t, id, resp.ID)
}

func TestClient_ManufacturerDelete(t *testing.T) {
	defer gock.Off()
	gock.New(testURL).Delete("dcim/manufacturers/" + testManufacturerID + "/").Reply(204)

	id, err := uuid.Parse(testManufacturerID)
	require.NoError(t, err)

	err = testClient.Dcim.ManufacturerDelete(id)
	require.NoError(t, err)
}

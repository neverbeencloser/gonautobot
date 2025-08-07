package nautobot_test

import (
	"net/url"
	"path"
	"testing"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/dcim"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/h2non/gock.v1"
)

const (
	testSoftwareVersionID = "bd0d3850-385d-47c5-8117-75de4c92d6c6"
)

func TestClient_SoftwareVersionGet(t *testing.T) {
	gock.New(testURL).Get("dcim/software-versions/" + testSoftwareVersionID + "/").Reply(200).
		File(path.Join("fixtures", "dcim", "softwareversion_200_1.json"))

	id, err := uuid.Parse(testSoftwareVersionID)
	require.NoError(t, err)

	resp, err := testClient.Dcim.SoftwareVersionGet(id)
	require.NoError(t, err)
	assert.Equal(t, "6.15.2", resp.Version)
	assert.Equal(t, id, resp.ID)
}

func TestClient_SoftwareVersionFilter(t *testing.T) {
	gock.New(testURL).Get("dcim/software-versions/").Reply(200).
		File(path.Join("fixtures", "dcim", "softwareversions_200_1.json"))

	q := &url.Values{}
	q.Set("status", "active")

	resp, err := testClient.Dcim.SoftwareVersionFilter(q)
	require.NoError(t, err)
	assert.Len(t, resp, 2)
	assert.Equal(t, "4.34.0F", resp[0].Version)
}

func TestClient_SoftwareVersionAll(t *testing.T) {
	gock.New(testURL).Get("dcim/software-versions/").Reply(200).
		File(path.Join("fixtures", "dcim", "softwareversions_200_1.json"))

	resp, err := testClient.Dcim.SoftwareVersionAll()
	require.NoError(t, err)

	assert.Len(t, resp, 2)
	assert.Equal(t, "4.34.0F", resp[0].Version)
}

func TestClient_SoftwareVersionCreate(t *testing.T) {
	newSoftwareVersion := dcim.NewSoftwareVersion{
		Platform: "9a47828b-0f74-42b0-8d30-99fd9806cb1f",
		Status:   "active",
		Version:  "6.15.2",
	}

	gock.New(testURL).Post("dcim/software-versions/").
		JSON(newSoftwareVersion).
		Reply(201).
		File(path.Join("fixtures", "dcim", "softwareversion_create_200_1.json"))

	resp, err := testClient.Dcim.SoftwareVersionCreate(newSoftwareVersion)
	require.NoError(t, err)
	assert.Equal(t, newSoftwareVersion.Version, resp.Version)
	assert.Equal(t, uuid.MustParse(testSoftwareVersionID), resp.ID)
}

func TestClient_SoftwareVersionUpdate(t *testing.T) {
	updateData := map[string]any{
		"version": "6.15.3",
	}

	gock.New(testURL).Patch("dcim/software-versions/" + testSoftwareVersionID + "/").
		JSON(updateData).
		Reply(200).
		File(path.Join("fixtures", "dcim", "softwareversion_update_200_1.json"))

	id, err := uuid.Parse(testSoftwareVersionID)
	require.NoError(t, err)

	resp, err := testClient.Dcim.SoftwareVersionUpdate(id, updateData)
	require.NoError(t, err)

	assert.Equal(t, "6.15.2", resp.Version)
	assert.Equal(t, id, resp.ID)
}

func TestClient_SoftwareVersionDelete(t *testing.T) {
	gock.New(testURL).Delete("dcim/software-versions/" + testSoftwareVersionID + "/").Reply(204)

	id, err := uuid.Parse(testSoftwareVersionID)
	require.NoError(t, err)

	err = testClient.Dcim.SoftwareVersionDelete(id)
	require.NoError(t, err)
	assert.True(t, gock.IsDone())
}

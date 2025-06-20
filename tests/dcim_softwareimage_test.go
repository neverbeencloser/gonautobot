package nautobot_test

import (
	"net/url"
	"path"
	"testing"

	"github.com/google/uuid"
	"github.com/josh-silvas/gonautobot/dcim"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/h2non/gock.v1"
)

const (
	testSoftwareImageFileID = "9f80dd3f-0394-45e0-a759-c6c43f0a6b76"
)

func TestClient_SoftwareImageFileGet(t *testing.T) {
	gock.New(testURL).Get("dcim/software-image-files/" + testSoftwareImageFileID + "/").Reply(200).
		File(path.Join("fixtures", "dcim", "softwareimage_200_1.json"))

	id, err := uuid.Parse(testSoftwareImageFileID)
	require.NoError(t, err)

	resp, err := testClient.Dcim.SoftwareImageFileGet(id)
	require.NoError(t, err)
	assert.Equal(t, "efwwefwf", resp.ImageFileName)
	assert.Equal(t, id, resp.ID)
}

func TestClient_SoftwareImageFileFilter(t *testing.T) {
	gock.New(testURL).Get("dcim/software-image-files/").Reply(200).
		File(path.Join("fixtures", "dcim", "softwareimages_200_1.json"))

	q := &url.Values{}
	q.Set("status", "active")

	resp, err := testClient.Dcim.SoftwareImageFileFilter(q)
	require.NoError(t, err)
	assert.Len(t, resp, 2)
	assert.Equal(t, "asdfsdf", resp[0].ImageFileName)
}

func TestClient_SoftwareImageFileAll(t *testing.T) {
	gock.New(testURL).Get("dcim/software-image-files/").Reply(200).
		File(path.Join("fixtures", "dcim", "softwareimages_200_1.json"))

	resp, err := testClient.Dcim.SoftwareImageFileAll()
	require.NoError(t, err)

	assert.Len(t, resp, 2)
	assert.Equal(t, "asdfsdf", resp[0].ImageFileName)
}

func TestClient_SoftwareImageFileCreate(t *testing.T) {
	newSIF := dcim.NewSoftwareImageFile{
		ImageFileName:   "efwwefwf",
		SoftwareVersion: "bd0d3850-385d-47c5-8117-75de4c92d6c6",
		Status:          "active",
	}

	gock.New(testURL).Post("dcim/software-image-files/").
		JSON(newSIF).
		Reply(201).
		File(path.Join("fixtures", "dcim", "softwareimage_create_201_1.json"))

	resp, err := testClient.Dcim.SoftwareImageFileCreate(newSIF)
	require.NoError(t, err)
	assert.Equal(t, newSIF.ImageFileName, resp.ImageFileName)
	assert.Equal(t, uuid.MustParse(testSoftwareImageFileID), resp.ID)
}

func TestClient_SoftwareImageFileUpdate(t *testing.T) {
	updateData := map[string]any{
		"image_file_name": "updated_name",
	}

	gock.New(testURL).Patch("dcim/software-image-files/" + testSoftwareImageFileID + "/").
		JSON(updateData).
		Reply(200).
		File(path.Join("fixtures", "dcim", "softwareimage_update_200_1.json"))

	id, err := uuid.Parse(testSoftwareImageFileID)
	require.NoError(t, err)

	resp, err := testClient.Dcim.SoftwareImageFileUpdate(id, updateData)
	require.NoError(t, err)

	assert.Equal(t, "efwwefwf", resp.ImageFileName)
	assert.Equal(t, id, resp.ID)
}

func TestClient_SoftwareImageFileDelete(t *testing.T) {
	gock.New(testURL).Delete("dcim/software-image-files/" + testSoftwareImageFileID + "/").Reply(204)

	id, err := uuid.Parse(testSoftwareImageFileID)
	require.NoError(t, err)

	err = testClient.Dcim.SoftwareImageFileDelete(id)
	require.NoError(t, err)
}

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
	testLocationID       = "aec16331-8ab4-4fdb-ba6e-8e978be45b77"
	testParentLocationID = "aec16331-8ab4-4fdb-ba6e-8e978be45b77"
)

func TestClient_LocationGet(t *testing.T) {
	defer gock.Off()

	gock.New(testURL).Get("dcim/locations/" + testLocationID + "/").Reply(200).
		File(path.Join("fixtures", "dcim", "location_200_1.json"))

	id, err := uuid.Parse(testLocationID)
	require.NoError(t, err)

	resp, err := testClient.Dcim.LocationGet(id)
	require.NoError(t, err)
	assert.Equal(t, "us-west-1", resp.Name)
	assert.Equal(t, id, resp.ID)
	assert.Equal(t, uuid.MustParse("7e222b0f-efe9-4bb3-81a0-fb9e81db4ca6"), resp.Status.ID)
}

func TestClient_LocationFilter(t *testing.T) {
	defer gock.Off()

	gock.New(testURL).Get("dcim/locations/").Reply(200).
		File(path.Join("fixtures", "dcim", "locations_200_1.json"))

	q := &url.Values{}
	q.Set("content_types", "dcim.device")

	resp, err := testClient.Dcim.LocationFilter(q)
	require.NoError(t, err)

	assert.NotEmpty(t, resp)
	assert.Len(t, resp, 3)
	assert.Equal(t, resp[0].ID, uuid.MustParse(testLocationID))
}

func TestClient_LocationAll(t *testing.T) {
	defer gock.Off()

	gock.New(testURL).Get("dcim/locations/").Reply(200).
		File(path.Join("fixtures", "dcim", "locations_200_1.json"))

	resp, err := testClient.Dcim.LocationAll()
	require.NoError(t, err)

	assert.NotEmpty(t, resp)
	assert.Len(t, resp, 3)
	assert.Equal(t, resp[0].ID, uuid.MustParse(testLocationID))
}

func TestClient_LocationCreate(t *testing.T) {
	defer gock.Off()

	newLocation := dcim.NewLocation{
		Name:         "us-west-1",
		LocationType: "1d464670-8fd2-4d1e-bd9d-8811bf2d1fe0",
		Status:       "7e222b0f-efe9-4bb3-81a0-fb9e81db4ca6",
	}

	gock.New(testURL).Post("dcim/locations/").
		JSON(newLocation).
		Reply(200).
		File(path.Join("fixtures", "dcim", "location_create_201_1.json"))

	resp, err := testClient.Dcim.LocationCreate(newLocation)
	require.NoError(t, err)

	assert.Equal(t, newLocation.Name, resp.Name)
	assert.Equal(t, uuid.MustParse(newLocation.LocationType), resp.LocationType.ID)
	assert.Equal(t, uuid.MustParse(newLocation.Status), resp.Status.ID)
	assert.Equal(t, uuid.MustParse(testLocationID), resp.ID)
}

func TestClient_LocationUpdate(t *testing.T) {
	defer gock.Off()

	updateData := map[string]any{
		"description": "updated description",
	}

	gock.New(testURL).Patch("dcim/locations/" + testLocationID + "/").
		JSON(updateData).
		Reply(200).
		File(path.Join("fixtures", "dcim", "location_update_200_1.json"))

	id, err := uuid.Parse(testLocationID)
	require.NoError(t, err)

	resp, err := testClient.Dcim.LocationUpdate(id, updateData)
	require.NoError(t, err)

	assert.Equal(t, "updated description", resp.Description)
	assert.Equal(t, id, resp.ID)
}

func TestClient_LocationDelete(t *testing.T) {
	defer gock.Off()

	gock.New(testURL).Delete("dcim/locations/" + testLocationID + "/").Reply(204)

	id, err := uuid.Parse(testLocationID)
	require.NoError(t, err)

	err = testClient.Dcim.LocationDelete(id)
	require.NoError(t, err)
	assert.True(t, gock.IsDone())
}

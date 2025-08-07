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
	testLocationTypeID       = "efb32b11-e80a-4880-9f96-9a1b16be9896"
	testParentLocationTypeID = "e9ad1520-23d9-48e7-8de8-c44a88411aa4"
)

func TestClient_LocationTypeGet(t *testing.T) {
	defer gock.Off()

	gock.New(testURL).Get("dcim/location-types/" + testLocationTypeID + "/").Reply(200).
		File(path.Join("fixtures", "dcim", "location-type_200_1.json"))

	id, err := uuid.Parse(testLocationTypeID)
	require.NoError(t, err)

	resp, err := testClient.Dcim.LocationTypeGet(id)
	require.NoError(t, err)
	assert.Equal(t, "Site", resp.Name)
	assert.Equal(t, id, resp.ID)
}

func TestClient_LocationTypeFilter(t *testing.T) {
	defer gock.Off()

	gock.New(testURL).Get("dcim/location-types/").Reply(200).
		File(path.Join("fixtures", "dcim", "location-types_200_1.json"))

	q := &url.Values{}
	q.Set("content_types", "dcim.device")

	resp, err := testClient.Dcim.LocationTypeFilter(q)
	require.NoError(t, err)

	assert.NotEmpty(t, resp)
	assert.Len(t, resp, 3)
	assert.Equal(t, resp[2].ID, uuid.MustParse(testLocationTypeID))
}

func TestClient_LocationTypeAll(t *testing.T) {
	defer gock.Off()

	gock.New(testURL).Get("dcim/location-types/").Reply(200).
		File(path.Join("fixtures", "dcim", "location-types_200_1.json"))

	resp, err := testClient.Dcim.LocationTypeAll()
	require.NoError(t, err)

	assert.NotEmpty(t, resp)
	assert.Len(t, resp, 3)
	assert.Equal(t, resp[2].ID, uuid.MustParse(testLocationTypeID))
}

func TestClient_LocationTypeCreate(t *testing.T) {
	defer gock.Off()

	parentID := testParentLocationTypeID
	newLocationType := dcim.NewLocationType{
		Name:         "Site",
		ContentTypes: []string{"dcim.device", "ipam.prefix", "ipam.vlan", "ipam.vlangroup"},
		Nestable:     true,
		Parent:       parentID,
		Description:  "",
	}

	gock.New(testURL).Post("dcim/location-types/").
		JSON(newLocationType).
		Reply(200).
		File(path.Join("fixtures", "dcim", "location-type_create_200_1.json"))

	resp, err := testClient.Dcim.LocationTypeCreate(newLocationType)
	require.NoError(t, err)

	assert.Equal(t, newLocationType.Name, resp.Name)
	assert.Equal(t, newLocationType.Nestable, resp.Nestable)
	assert.Equal(t, newLocationType.ContentTypes, resp.ContentTypes)
	assert.Equal(t, uuid.MustParse(testLocationTypeID), resp.ID)

	require.NotNil(t, resp.Parent)
	assert.Equal(t, uuid.MustParse(testParentLocationTypeID), resp.Parent.ID)
}

func TestClient_LocationTypeUpdate(t *testing.T) {
	defer gock.Off()

	updateData := map[string]any{
		"description": "updated description",
	}

	gock.New(testURL).Patch("dcim/location-types/" + testLocationTypeID + "/").
		JSON(updateData).
		Reply(200).
		File(path.Join("fixtures", "dcim", "location-type_update_200_1.json"))

	id, err := uuid.Parse(testLocationTypeID)
	require.NoError(t, err)

	resp, err := testClient.Dcim.LocationTypeUpdate(id, updateData)
	require.NoError(t, err)

	assert.Equal(t, "updated description", resp.Description)
	assert.Equal(t, id, resp.ID)
}

func TestClient_LocationTypeDelete(t *testing.T) {
	defer gock.Off()

	gock.New(testURL).Delete("dcim/location-types/" + testLocationTypeID + "/").Reply(204)

	id, err := uuid.Parse(testLocationTypeID)
	require.NoError(t, err)

	err = testClient.Dcim.LocationTypeDelete(id)
	require.NoError(t, err)
	assert.True(t, gock.IsDone())
}

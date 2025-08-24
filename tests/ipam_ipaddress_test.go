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
	testIPAddressID = "50476c26-4349-4669-bc40-fa3baedd72fb"
)

func TestClient_IPAddressGet(t *testing.T) {
	defer gock.Off()

	gock.New(testURL).Get("ipam/ip-addresses/" + testIPAddressID + "/").Reply(200).
		File(path.Join("fixtures", "ipam", "ip_address_200_1.json"))

	id, err := uuid.Parse(testIPAddressID)
	require.NoError(t, err)

	resp, err := testClient.Ipam.IPAddressGet(id)
	require.NoError(t, err)
	assert.Equal(t, "2.1.1.1/32", resp.Address)
	assert.Equal(t, id, resp.ID)
}

func TestClient_IPAddressFilter(t *testing.T) {
	defer gock.Off()

	gock.New(testURL).Get("ipam/ip-addresses/").Reply(200).
		File(path.Join("fixtures", "ipam", "ip_addresses_200_1.json"))

	q := &url.Values{}
	q.Set("address", "2.1.1.1/32")

	resp, err := testClient.Ipam.IPAddressFilter(q)
	require.NoError(t, err)

	assert.NotEmpty(t, resp)
	assert.Len(t, resp, 2)
}

func TestClient_IPAddressAll(t *testing.T) {
	defer gock.Off()

	gock.New(testURL).Get("ipam/ip-addresses/").Reply(200).
		File(path.Join("fixtures", "ipam", "ip_addresses_200_1.json"))

	resp, err := testClient.Ipam.IPAddressAll()
	require.NoError(t, err)

	assert.NotEmpty(t, resp)
	assert.Len(t, resp, 2)
}

func TestClient_IPAddressCreate(t *testing.T) {
	defer gock.Off()

	newIPAddress := types.NewIPAddress{
		Address:     "2.1.1.1/32",
		Namespace:   "global",
		Status:      "active",
		Description: "",
	}

	gock.New(testURL).Post("ipam/ip-addresses/").
		JSON(newIPAddress).
		Reply(201).
		File(path.Join("fixtures", "ipam", "ip_address_200_1.json"))

	resp, err := testClient.Ipam.IPAddressCreate(newIPAddress)
	require.NoError(t, err)

	assert.Equal(t, newIPAddress.Address, resp.Address)
	assert.Equal(t, newIPAddress.Description, resp.Description)
	assert.Equal(t, uuid.MustParse(testIPAddressID), resp.ID)
}

func TestClient_IPAddressUpdate(t *testing.T) {
	defer gock.Off()

	updateData := map[string]any{
		"description": "updated description",
	}

	gock.New(testURL).Patch("ipam/ip-addresses/" + testIPAddressID + "/").
		JSON(updateData).
		Reply(200).
		File(path.Join("fixtures", "ipam", "ip_address_update_200_1.json"))

	id, err := uuid.Parse(testIPAddressID)
	require.NoError(t, err)

	resp, err := testClient.Ipam.IPAddressUpdate(id, updateData)
	require.NoError(t, err)

	assert.Equal(t, "updated description", resp.Description)
	assert.Equal(t, id, resp.ID)
}

func TestClient_IPAddressDelete(t *testing.T) {
	defer gock.Off()

	gock.New(testURL).Delete("ipam/ip-addresses/" + testIPAddressID + "/").Reply(204)

	id, err := uuid.Parse(testIPAddressID)
	require.NoError(t, err)

	err = testClient.Ipam.IPAddressDelete(id)
	require.NoError(t, err)
	assert.True(t, gock.IsDone())
}

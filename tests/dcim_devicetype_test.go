package nautobot_test

import (
	"errors"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/dcim"
	"github.com/neverbeencloser/gonautobot/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/h2non/gock.v1"
)

const (
	testDeviceTypeID = "a5286991-7396-47f9-9ca1-33ca3be9d2ec"
)

func TestClient_DeviceTypeGet(t *testing.T) {
	gock.New(testURL).Get("dcim/device-types/" + testDeviceTypeID + "/").Reply(200).
		File(path.Join("fixtures", "dcim", "devicetype_200_1.json"))

	id, err := uuid.Parse(testDeviceTypeID)
	require.NoError(t, err)

	resp, err := testClient.Dcim.DeviceTypeGet(id)
	require.NoError(t, err)
	assert.Equal(t, "t2.micro", resp.Model)
	assert.Equal(t, id, resp.ID)
}

func TestClient_DeviceTypeFilter(t *testing.T) {
	gock.New(testURL).Get("dcim/device-types/").Reply(200).
		File(path.Join("fixtures", "dcim", "devicetypes_200_1.json"))

	q := &url.Values{}
	q.Set("manufacturer", "AWS")

	resp, err := testClient.Dcim.DeviceTypeFilter(q)
	require.NoError(t, err)
	assert.Len(t, resp, 3)
	assert.Equal(t, "t2.micro", resp[0].Model)
}

func TestClient_DeviceTypeAll(t *testing.T) {
	gock.New(testURL).Get("dcim/device-types/").Reply(200).
		File(path.Join("fixtures", "dcim", "devicetypes_200_1.json"))

	resp, err := testClient.Dcim.DeviceTypeAll()
	require.NoError(t, err)

	assert.Len(t, resp, 3)
	assert.Equal(t, "t2.micro", resp[0].Model)
}

func TestClient_DeviceTypeCreate(t *testing.T) {
	newDeviceType := dcim.NewDeviceType{
		Model:        "t2.micro",
		Manufacturer: "acc4c522-cf77-46bf-bdf3-d0e4a69b4d1e",
		UHeight:      1,
		IsFullDepth:  true,
	}

	gock.New(testURL).Post("dcim/device-types/").
		Reply(201).
		File(path.Join("fixtures", "dcim", "devicetype_create_201_1.json"))

	resp, err := testClient.Dcim.DeviceTypeCreate(newDeviceType)
	require.NoError(t, err)
	assert.Equal(t, newDeviceType.Model, resp.Model)
	assert.Equal(t, uuid.MustParse(testDeviceTypeID), resp.ID)
}

func TestClient_DeviceTypeUpdate(t *testing.T) {
	updateData := dcim.NewDeviceType{
		Comments: "updated comments",
	}

	gock.New(testURL).Patch("dcim/device-types/" + testDeviceTypeID + "/").
		Reply(200).
		File(path.Join("fixtures", "dcim", "devicetype_update_200_1.json"))

	id, err := uuid.Parse(testDeviceTypeID)
	require.NoError(t, err)

	resp, err := testClient.Dcim.DeviceTypeUpdate(id, updateData)
	require.NoError(t, err)

	assert.Equal(t, id, resp.ID)
}

func TestClient_DeviceTypeDelete(t *testing.T) {
	gock.New(testURL).Delete("dcim/device-types/" + testDeviceTypeID + "/").Reply(204)

	id, err := uuid.Parse(testDeviceTypeID)
	require.NoError(t, err)

	err = testClient.Dcim.DeviceTypeDelete(id)
	require.NoError(t, err)
}

func TestCreateMultipartBody(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "test-image-*.png")
	require.NoError(t, err)
	defer func() {
		_ = os.Remove(tmpFile.Name()) // nolint: errcheck
	}()

	_, err = tmpFile.WriteString("dummy image data")
	require.NoError(t, err)
	err = tmpFile.Close()
	require.NoError(t, err)

	newDeviceType := dcim.NewDeviceType{
		Model:        "test-model",
		Manufacturer: "test-manufacturer",
		UHeight:      1,
		CustomFields: map[string]any{
			"cf_text": "test value",
		},
	}
	newDeviceType.SetFrontImage(tmpFile.Name())

	body, contentType, err := shared.NewMultipartBody(newDeviceType)
	require.NoError(t, err)

	boundary := strings.Split(contentType, "boundary=")[1]
	mr := multipart.NewReader(body, boundary)

	parts := make(map[string][]byte)
	for {
		p, err := mr.NextPart()
		if errors.Is(err, io.EOF) {
			break
		}
		require.NoError(t, err)

		partBody, err := io.ReadAll(p)
		require.NoError(t, err)

		parts[p.FormName()] = partBody
	}

	assert.Equal(t, newDeviceType.Model, string(parts["model"]))
	assert.Equal(t, newDeviceType.Manufacturer, string(parts["manufacturer"]))
	assert.Equal(t, "1", string(parts["u_height"]))
	assert.Equal(t, `{"cf_text":"test value"}`, string(parts["custom_fields"]))

	_, ok := parts["front_image"]
	assert.True(t, ok, "front_image part should exist")
}

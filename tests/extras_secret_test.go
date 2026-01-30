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
	testSecretID = "d4e5f6a7-b8c9-0123-4567-890abcdef012"
)

func TestClient_SecretGet(t *testing.T) {
	gock.New(testURL).Get("extras/secrets/" + testSecretID + "/").Reply(200).
		File(path.Join("fixtures", "extras", "secret_200_1.json"))

	id, err := uuid.Parse(testSecretID)
	require.NoError(t, err)

	resp, err := testClient.Extras.SecretGet(id)
	require.NoError(t, err)
	assert.Equal(t, "aws-credentials", resp.Name)
	assert.Equal(t, id, resp.ID)
	assert.Equal(t, "environment-variable", resp.Provider)
	assert.Equal(t, "AWS API credentials for automation", resp.Description)
}

func TestClient_SecretFilter(t *testing.T) {
	gock.New(testURL).Get("extras/secrets/").Reply(200).
		File(path.Join("fixtures", "extras", "secrets_200_1.json"))

	q := &url.Values{}
	q.Set("provider", "environment-variable")

	resp, err := testClient.Extras.SecretFilter(q)
	require.NoError(t, err)
	assert.Len(t, resp, 2)
	assert.Equal(t, "aws-credentials", resp[0].Name)
}

func TestClient_SecretAll(t *testing.T) {
	gock.New(testURL).Get("extras/secrets/").Reply(200).
		File(path.Join("fixtures", "extras", "secrets_200_1.json"))

	resp, err := testClient.Extras.SecretAll()
	require.NoError(t, err)

	assert.Len(t, resp, 2)
	assert.Equal(t, "aws-credentials", resp[0].Name)
	assert.Equal(t, "github-token", resp[1].Name)
}

func TestClient_SecretCreate(t *testing.T) {
	newSecret := types.NewSecret{
		Name:     "aws-credentials",
		Provider: "environment-variable",
		Parameters: map[string]any{
			"variable": "AWS_SECRET_ACCESS_KEY",
		},
	}

	gock.New(testURL).Post("extras/secrets/").
		JSON(newSecret).
		Reply(201).
		File(path.Join("fixtures", "extras", "secret_create_201_1.json"))

	resp, err := testClient.Extras.SecretCreate(newSecret)
	require.NoError(t, err)
	assert.Equal(t, newSecret.Name, resp.Name)
	assert.Equal(t, newSecret.Provider, resp.Provider)
	assert.Equal(t, uuid.MustParse(testSecretID), resp.ID)
}

func TestClient_SecretUpdate(t *testing.T) {
	updateData := map[string]any{
		"description": "Updated description for AWS credentials",
	}

	gock.New(testURL).Patch("extras/secrets/" + testSecretID + "/").
		JSON(updateData).
		Reply(200).
		File(path.Join("fixtures", "extras", "secret_update_200_1.json"))

	id, err := uuid.Parse(testSecretID)
	require.NoError(t, err)

	resp, err := testClient.Extras.SecretUpdate(id, updateData)
	require.NoError(t, err)

	assert.Equal(t, "Updated description for AWS credentials", resp.Description)
	assert.Equal(t, id, resp.ID)
}

func TestClient_SecretDelete(t *testing.T) {
	gock.New(testURL).Delete("extras/secrets/" + testSecretID + "/").Reply(204)

	id, err := uuid.Parse(testSecretID)
	require.NoError(t, err)

	err = testClient.Extras.SecretDelete(id)
	require.NoError(t, err)
	assert.True(t, gock.IsDone())
}

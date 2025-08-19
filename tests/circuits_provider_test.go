package nautobot_test

import (
	"path"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestClient_GetProvider(t *testing.T) {
	gock.New(testURL).Get("circuits/providers/").Reply(200).
		File(path.Join("fixtures", "circuits", "provider_200_1.json"))

	circuit, err := testClient.Circuits.ProviderGet(uuid.MustParse("8e384236-9ff5-4f71-8418-1d607177e10e"))
	require.NoError(t, err)
	assert.Equal(t, "8e384236-9ff5-4f71-8418-1d607177e10e", circuit.ID.String())
	assert.Equal(t, 7018, circuit.Asn)
}

func TestClient_GetProviders(t *testing.T) {
	gock.New(testURL).Get("circuits/providers/").Reply(200).
		File(path.Join("fixtures", "circuits", "providers_200_1.json"))

	circuits, err := testClient.Circuits.ProviderFilter(nil)
	require.NoError(t, err)
	assert.Len(t, circuits, 2)
}

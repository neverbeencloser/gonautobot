package nautobot_test

import (
	"github.com/stretchr/testify/require"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestClient_GetProviders(t *testing.T) {
	gock.New(testURL).Get("circuits/providers/").Reply(200).
		File(path.Join("fixtures", "circuits", "providers_200_1.json"))

	circuits, err := testClient.Circuits.GetProviders(nil)
	require.NoError(t, err)
	assert.Len(t, circuits, 2)
}

package nautobot_test

import (
	"github.com/stretchr/testify/require"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestClient_GetCircuits(t *testing.T) {
	gock.New(testURL).Get("circuits/circuits/").Reply(200).
		File(path.Join("fixtures", "circuits", "circuits_200_1.json"))

	circuits, err := testClient.Circuits.GetCircuits(nil)
	require.NoError(t, err)
	assert.Len(t, circuits, 2)
}

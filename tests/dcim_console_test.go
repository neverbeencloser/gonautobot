package nautobot_test

import (
	"github.com/stretchr/testify/require"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestClient_GetConsoleConnections(t *testing.T) {
	gock.New(testURL).Get("dcim/console-connections/").Reply(200).
		File(path.Join("fixtures", "dcim", "console_connections_200_1.json"))

	resp, err := testClient.Dcim.ConsoleConnectionFilter(nil)
	require.NoError(t, err)
	assert.Len(t, resp, 1)
}

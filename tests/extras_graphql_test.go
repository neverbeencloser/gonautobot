package nautobot_test

import (
	"github.com/josh-silvas/gonautobot/extras"
	"github.com/stretchr/testify/require"
	"gopkg.in/h2non/gock.v1"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GraphQL(t *testing.T) {
	q := &extras.GraphQLQuery{
		Query: "{devices(name: \"ams01-dist-01\") {interfaces {name}}}",
	}
	gock.New(testURL).Post("graphql/").JSON(q).Reply(200).
		File(path.Join("fixtures", "extras", "graphql_200_1.json"))

	type (
		interfaces struct {
			Name string `json:"name"`
		}
		devices struct {
			Interfaces []interfaces `json:"interfaces"`
		}
		testStruct struct {
			Devices []devices `json:"devices"`
		}
	)
	resp := new(testStruct)

	err := testClient.Extras.GraphQL(q, resp)
	require.NoError(t, err)
	assert.Equal(t, "Ethernet1", resp.Devices[0].Interfaces[0].Name)
	assert.Len(t, resp.Devices[0].Interfaces, 48)
}

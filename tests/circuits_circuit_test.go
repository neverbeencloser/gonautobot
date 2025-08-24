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
	testCircuitID = "6b2c4c96-2b3e-4533-b085-9fe31a58dbdc"
)

func TestClient_CircuitGet(t *testing.T) {
	circuitID := uuid.MustParse(testCircuitID)
	gock.New(testURL).Get("circuits/circuits/" + testCircuitID + "/").Reply(200).
		File(path.Join("fixtures", "circuits", "circuit_200_1.json"))

	resp, err := testClient.Circuits.CircuitGet(circuitID)
	require.NoError(t, err)
	assert.Equal(t, circuitID, resp.ID)
	assert.Equal(t, "ntt-104265404093023273", resp.CircuitID)
}

func TestClient_CircuitFilter(t *testing.T) {
	gock.New(testURL).Get("circuits/circuits/").Reply(200).
		File(path.Join("fixtures", "circuits", "circuits_200_1.json"))

	circuits, err := testClient.Circuits.CircuitFilter(&url.Values{"offset": {"0"}})
	require.NoError(t, err)
	assert.Len(t, circuits, 2)
}

func TestClient_CircuitAll(t *testing.T) {
	gock.New(testURL).Get("circuits/circuits/").Reply(200).
		File(path.Join("fixtures", "circuits", "circuits_200_1.json"))

	circuits, err := testClient.Circuits.CircuitAll()
	require.NoError(t, err)
	assert.Len(t, circuits, 2)
}

func TestClient_CircuitCreate(t *testing.T) {
	newCircuit := types.CircuitRequest{
		CircuitID: "TEST-001",
		Provider:  "provider-uuid",
		Type:      "type-uuid",
		Status:    "status-uuid",
	}

	gock.New(testURL).Post("circuits/circuits/").
		Reply(201).
		File(path.Join("fixtures", "circuits", "circuit_200_1.json"))

	resp, err := testClient.Circuits.CircuitCreate(newCircuit)
	require.NoError(t, err)
	assert.Equal(t, uuid.MustParse(testCircuitID), resp.ID)
}

func TestClient_CircuitUpdate(t *testing.T) {
	updateData := map[string]any{
		"description": "updated description",
	}

	gock.New(testURL).Patch("circuits/circuits/" + testCircuitID + "/").
		Reply(200).
		File(path.Join("fixtures", "circuits", "circuit_updated_200_1.json"))

	circuitID := uuid.MustParse(testCircuitID)
	resp, err := testClient.Circuits.CircuitUpdate(circuitID, updateData)
	require.NoError(t, err)
	assert.Equal(t, circuitID, resp.ID)
	assert.Equal(t, "updated description", resp.Description)
}

func TestClient_CircuitDelete(t *testing.T) {
	gock.New(testURL).Delete("circuits/circuits/" + testCircuitID + "/").Reply(204)

	circuitID := uuid.MustParse(testCircuitID)
	err := testClient.Circuits.CircuitDelete(circuitID)
	require.NoError(t, err)
}

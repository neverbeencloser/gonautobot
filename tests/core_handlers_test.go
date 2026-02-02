package nautobot_test

import (
	"path"
	"testing"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/h2non/gock.v1"
)

// TestCore_Action tests the generic Action handler for resource action endpoints.
func TestCore_Action(t *testing.T) {
	type ActionRequest struct {
		Data   map[string]any `json:"data,omitempty"`
		Dryrun bool           `json:"dryrun,omitempty"`
	}

	testID := uuid.MustParse("f1a2b3c4-d5e6-7890-abcd-ef1234567890")
	actionRequest := ActionRequest{
		Data:   map[string]any{"key": "value"},
		Dryrun: false,
	}

	gock.New(testURL).Post("extras/jobs/" + testID.String() + "/run/").
		JSON(actionRequest).
		Reply(200).
		File(path.Join("fixtures", "extras", "job_run_200_1.json"))

	// Use the core.Action handler directly through the client's Request (core.Client)
	resp, err := core.Action[map[string]any](testClient.Request, "extras/jobs/", testID, "run", actionRequest)
	require.NoError(t, err)
	require.NotNil(t, resp)

	// Verify the response was parsed
	result := *resp
	assert.NotNil(t, result["job_result"])
}

// TestCore_Action_NilID tests that Action returns an error when given a nil UUID.
func TestCore_Action_NilID(t *testing.T) {
	type ActionRequest struct {
		Data map[string]any `json:"data,omitempty"`
	}

	_, err := core.Action[map[string]any](testClient.Request, "extras/jobs/", uuid.Nil, "run", ActionRequest{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "ID is missing or nil")
}

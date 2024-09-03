package nautobot_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/url"
	"path"
	"testing"

	"gopkg.in/h2non/gock.v1"
)

func TestClient_GetTag(t *testing.T) {
	gock.New(testURL).Get("extras/tags/91283c9c-e2ba-4cfe-8f9f-ed6dd6e18ba1/").Reply(200).
		File(path.Join("fixtures", "extras", "tag_200_1.json"))

	resp, err := testClient.Extras.TagGet("91283c9c-e2ba-4cfe-8f9f-ed6dd6e18ba1")
	require.NoError(t, err)
	assert.Equal(t, "sup-720", resp.Name)
}

func TestClient_GetTags(t *testing.T) {
	gock.New(testURL).Get("extras/tags/").Reply(200).
		File(path.Join("fixtures", "extras", "tags_200_1.json"))

	resp, err := testClient.Extras.TagFilter(&url.Values{"offset": {"50"}})
	require.NoError(t, err)
	assert.Len(t, resp, 1)
}

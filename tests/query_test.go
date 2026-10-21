package nautobot_test

import (
	"net/http"
	"net/url"
	"path"
	"testing"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/query"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/h2non/gock.v1"
)

const testQueryID = "1ce0df0c-5221-4b4e-a7e5-d52810828041"

// matchQuery : match only requests whose query parameters exactly equal want.
func matchQuery(t *testing.T, want url.Values) gock.MatchFunc {
	t.Helper()
	return func(req *http.Request, _ *gock.Request) (bool, error) {
		return assert.Equal(t, want, req.URL.Query()), nil
	}
}

func TestQueryOption_GetNoOptions(t *testing.T) {
	t.Cleanup(gock.Flush)
	gock.New(testURL).Get("extras/roles/" + testRoleID + "/").
		AddMatcher(matchQuery(t, url.Values{})).
		Reply(200).
		File(path.Join("fixtures", "extras", "role_200_1.json"))

	_, err := testClient.Extras.RoleGet(uuid.MustParse(testRoleID))
	require.NoError(t, err)
	assert.True(t, gock.IsDone())
}

func TestQueryOption_GetDepthM2M(t *testing.T) {
	t.Cleanup(gock.Flush)
	gock.New(testURL).Get("dcim/interfaces/" + testQueryID + "/").
		AddMatcher(matchQuery(t, url.Values{"depth": {"1"}, "exclude_m2m": {"false"}})).
		Reply(200).
		File(path.Join("fixtures", "dcim", "interface_200_1.json"))

	_, err := testClient.Dcim.InterfaceGet(uuid.MustParse(testQueryID), query.WithDepth(1), query.WithM2M())
	require.NoError(t, err)
	assert.True(t, gock.IsDone())
}

func TestQueryOption_GetWithValues(t *testing.T) {
	t.Cleanup(gock.Flush)
	gock.New(testURL).Get("dcim/devices/" + testQueryID + "/").
		AddMatcher(matchQuery(t, url.Values{"include": {"config_context"}})).
		Reply(200).
		File(path.Join("fixtures", "dcim", "device_200_1.json"))

	_, err := testClient.Dcim.DeviceGet(uuid.MustParse(testQueryID), query.WithValues(url.Values{"include": {"config_context"}}))
	require.NoError(t, err)
	assert.True(t, gock.IsDone())
}

func TestQueryOption_ListLaterOptionWins(t *testing.T) {
	t.Cleanup(gock.Flush)
	gock.New(testURL).Get("extras/roles/").
		AddMatcher(matchQuery(t, url.Values{"depth": {"2"}, "name": {"core", "edge"}})).
		Reply(200).
		File(path.Join("fixtures", "extras", "roles_200_1.json"))

	q := url.Values{"depth": {"0"}, "name": {"core", "edge"}}
	_, err := testClient.Extras.RoleList(query.WithValues(q), query.WithDepth(2))
	require.NoError(t, err)
	assert.True(t, gock.IsDone())
}

func TestQueryOption_ListPaginatesWithFilters(t *testing.T) {
	t.Cleanup(gock.Flush)
	gock.New(testURL).Get("dcim/devices/").
		AddMatcher(matchQuery(t, url.Values{"name__ic": {"core"}})).
		Reply(200).
		JSON(map[string]any{
			"count":   2,
			"next":    testURL + "dcim/devices/?limit=1&name__ic=core&offset=1",
			"results": []map[string]any{{"id": testQueryID, "name": "core-01"}},
		})
	gock.New(testURL).Get("dcim/devices/").
		AddMatcher(matchQuery(t, url.Values{"limit": {"1"}, "name__ic": {"core"}, "offset": {"1"}})).
		Reply(200).
		JSON(map[string]any{
			"count":   2,
			"next":    nil,
			"results": []map[string]any{{"id": testRoleID, "name": "core-02"}},
		})

	q := url.Values{"name__ic": {"core"}}
	resp, err := testClient.Dcim.DeviceList(query.WithValues(q))
	require.NoError(t, err)
	require.Len(t, resp, 2)
	assert.Equal(t, "core-01", resp[0].Name)
	assert.Equal(t, "core-02", resp[1].Name)
	assert.Equal(t, url.Values{"name__ic": {"core"}}, q)
	assert.True(t, gock.IsDone())
}

func TestQueryOption_GetNilOption(t *testing.T) {
	t.Cleanup(gock.Flush)
	gock.New(testURL).Get("extras/roles/" + testRoleID + "/").
		AddMatcher(matchQuery(t, url.Values{"depth": {"2"}})).
		Reply(200).
		File(path.Join("fixtures", "extras", "role_200_1.json"))

	_, err := testClient.Extras.RoleGet(uuid.MustParse(testRoleID), nil, query.WithDepth(2))
	require.NoError(t, err)
	assert.True(t, gock.IsDone())
}

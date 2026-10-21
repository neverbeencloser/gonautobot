package query_test

import (
	"net/url"
	"testing"

	"github.com/neverbeencloser/gonautobot/query"
	"github.com/stretchr/testify/assert"
)

func TestNew_NoOptions(t *testing.T) {
	assert.Equal(t, url.Values{}, query.New())
}

func TestNew_DepthM2M(t *testing.T) {
	q := query.New(query.WithDepth(1), query.WithM2M())
	assert.Equal(t, url.Values{"depth": {"1"}, "exclude_m2m": {"false"}}, q)
}

func TestNew_LaterOptionWins(t *testing.T) {
	q := query.New(
		query.WithValues(url.Values{"depth": {"0"}, "name": {"core", "edge"}}),
		query.WithDepth(2),
		query.WithValues(url.Values{"name": {"spine"}}),
	)
	assert.Equal(t, url.Values{"depth": {"2"}, "name": {"spine"}}, q)
}

func TestNew_NilOption(t *testing.T) {
	q := query.New(nil, query.WithDepth(2), nil)
	assert.Equal(t, url.Values{"depth": {"2"}}, q)
}

func TestWithValues_DoesNotAliasCaller(t *testing.T) {
	v := url.Values{"name": {"core"}}
	q := query.New(query.WithValues(v))
	q.Add("name", "edge")
	q["name"][0] = "changed"

	assert.Equal(t, url.Values{"name": {"core"}}, v)
}

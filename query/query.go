package query

import (
	"net/url"
	"slices"
	"strconv"
)

// Option : Optional query parameters for read requests.
//
// Options are applied in order; a later option replaces any key set by an earlier one.
type Option func(url.Values)

// WithDepth : Set the nested serialization depth (`depth=N`) of related objects.
func WithDepth(depth int) Option {
	return func(q url.Values) {
		q.Set("depth", strconv.Itoa(depth))
	}
}

// WithM2M : Include many-to-many fields (`exclude_m2m=false`), which Nautobot omits by default.
func WithM2M() Option {
	return func(q url.Values) {
		q.Set("exclude_m2m", "false")
	}
}

// WithValues : Merge arbitrary query parameters into the request; each key in v replaces
// any values previously set for that key.
func WithValues(v url.Values) Option {
	return func(q url.Values) {
		for k, vals := range v {
			q[k] = slices.Clone(vals)
		}
	}
}

// New : Build request query parameters from the provided options.
func New(opts ...Option) url.Values {
	q := url.Values{}
	for _, opt := range opts {
		if opt != nil {
			opt(q)
		}
	}
	return q
}

package dcim

import (
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/types"
)

const (
	dcimEndpointSoftwareVersion = "dcim/software-versions/"
)

type (
	// SoftwareVersion represents a software version in Nautobot.
	SoftwareVersion struct {
		ID               uuid.UUID      `json:"id"`
		Alias            string         `json:"alias"`
		Created          time.Time      `json:"created"`
		CustomFields     map[string]any `json:"custom_fields"`
		Display          string         `json:"display"`
		DocumentationURL string         `json:"documentation_url"`
		EndOfSupportDate *string        `json:"end_of_support_date"`
		LastUpdated      time.Time      `json:"last_updated"`
		LongTermSupport  bool           `json:"long_term_support"`
		NaturalSlug      string         `json:"natural_slug"`
		NotesURL         string         `json:"notes_url"`
		ObjectType       string         `json:"object_type"`
		Platform         Platform       `json:"platform"`
		PreRelease       bool           `json:"pre_release"`
		ReleaseDate      *string        `json:"release_date"`
		Status           types.Status   `json:"status"`
		Tags             []types.Tag    `json:"tags"`
		URL              string         `json:"url"`
		Version          string         `json:"version"`
	}

	// NewSoftwareVersion represents the data required to create a new software version.
	NewSoftwareVersion struct {
		Platform         string         `json:"platform"`
		Status           string         `json:"status"`
		Version          string         `json:"version"`
		Alias            string         `json:"alias,omitempty"`
		CustomFields     map[string]any `json:"custom_fields,omitempty"`
		DocumentationURL string         `json:"documentation_url,omitempty"`
		EndOfSupportDate string         `json:"end_of_support_date,omitempty"`
		LongTermSupport  bool           `json:"long_term_support"`
		PreRelease       bool           `json:"pre_release"`
		ReleaseDate      string         `json:"release_date,omitempty"`
		Tags             []string       `json:"tags,omitempty"`
	}
)

// SoftwareVersionGet : Get a SoftwareVersion by UUID identifier.
func (c *Client) SoftwareVersionGet(id uuid.UUID) (*SoftwareVersion, error) {
	return core.Get[SoftwareVersion](c.Client, dcimEndpointSoftwareVersion, id)
}

// SoftwareVersionFilter : Get a list of SoftwareVersions based on query parameters.
func (c *Client) SoftwareVersionFilter(q *url.Values) ([]SoftwareVersion, error) {
	softwareVersions := make([]SoftwareVersion, 0)
	return softwareVersions, core.Paginate[SoftwareVersion](c.Client, dcimEndpointSoftwareVersion, q, &softwareVersions)
}

// SoftwareVersionAll : Get all SoftwareVersions in Nautobot.
func (c *Client) SoftwareVersionAll() ([]SoftwareVersion, error) {
	softwareVersions := make([]SoftwareVersion, 0)
	return softwareVersions, core.Paginate[SoftwareVersion](c.Client, dcimEndpointSoftwareVersion, nil, &softwareVersions)
}

// SoftwareVersionCreate : Generate a new SoftwareVersion record in Nautobot.
func (c *Client) SoftwareVersionCreate(obj NewSoftwareVersion) (*SoftwareVersion, error) {
	return core.Create[SoftwareVersion, NewSoftwareVersion](c.Client, dcimEndpointSoftwareVersion, obj)
}

// SoftwareVersionDelete : Delete a SoftwareVersion by UUID identifier.
func (c *Client) SoftwareVersionDelete(id uuid.UUID) error {
	return core.Delete(c.Client, dcimEndpointSoftwareVersion, id)
}

// SoftwareVersionUpdate : Update an existing SoftwareVersion record in Nautobot.
func (c *Client) SoftwareVersionUpdate(id uuid.UUID, patch map[string]any) (*SoftwareVersion, error) {
	return core.Update[SoftwareVersion](c.Client, dcimEndpointSoftwareVersion, id, patch)
}

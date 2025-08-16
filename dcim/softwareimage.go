package dcim

import (
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/types"
)

const (
	dcimEndpointSoftwareImageFile = "dcim/software-image-files/"
)

type (
	// SoftwareImageFile represents a software image file in Nautobot.
	SoftwareImageFile struct {
		ID                  uuid.UUID       `json:"id"`
		Created             time.Time       `json:"created"`
		CustomFields        map[string]any  `json:"custom_fields"`
		DefaultImage        bool            `json:"default_image"`
		Display             string          `json:"display"`
		DownloadURL         string          `json:"download_url"`
		ExternalIntegration *types.Object   `json:"external_integration"`
		HashingAlgorithm    string          `json:"hashing_algorithm"`
		ImageFileChecksum   string          `json:"image_file_checksum"`
		ImageFileName       string          `json:"image_file_name"`
		ImageFileSize       *int            `json:"image_file_size"`
		LastUpdated         time.Time       `json:"last_updated"`
		NaturalSlug         string          `json:"natural_slug"`
		NotesURL            string          `json:"notes_url"`
		ObjectType          string          `json:"object_type"`
		SoftwareVersion     SoftwareVersion `json:"software_version"`
		Status              types.Status    `json:"status"`
		Tags                []types.Tag     `json:"tags"`
		URL                 string          `json:"url"`
	}

	// NewSoftwareImageFile represents the data required to create a new software image file.
	NewSoftwareImageFile struct {
		ImageFileName       string         `json:"image_file_name"`
		SoftwareVersion     string         `json:"software_version"`
		Status              string         `json:"status"`
		CustomFields        map[string]any `json:"custom_fields,omitempty"`
		DefaultImage        bool           `json:"default_image"`
		DownloadURL         string         `json:"download_url,omitempty"`
		ExternalIntegration string         `json:"external_integration,omitempty"`
		HashingAlgorithm    string         `json:"hashing_algorithm,omitempty"`
		ImageFileChecksum   string         `json:"image_file_checksum,omitempty"`
		ImageFileSize       int            `json:"image_file_size,omitempty"`
		Relationships       map[string]any `json:"relationships,omitempty"`
		Tags                []string       `json:"tags,omitempty"`
	}
)

// SoftwareImageFileGet : Get a SoftwareImageFile by UUID identifier.
func (c *Client) SoftwareImageFileGet(id uuid.UUID) (*SoftwareImageFile, error) {
	return core.Get[SoftwareImageFile](c.Client, dcimEndpointSoftwareImageFile, id)
}

// SoftwareImageFileFilter : Get a list of SoftwareImageFiles based on query parameters.
func (c *Client) SoftwareImageFileFilter(q *url.Values) ([]SoftwareImageFile, error) {
	sifs := make([]SoftwareImageFile, 0)
	return sifs, core.Paginate[SoftwareImageFile](c.Client, dcimEndpointSoftwareImageFile, q, &sifs)
}

// SoftwareImageFileAll : Get all SoftwareImageFiles in Nautobot.
func (c *Client) SoftwareImageFileAll() ([]SoftwareImageFile, error) {
	sifs := make([]SoftwareImageFile, 0)
	return sifs, core.Paginate[SoftwareImageFile](c.Client, dcimEndpointSoftwareImageFile, nil, &sifs)
}

// SoftwareImageFileCreate : Generate a new SoftwareImageFile record in Nautobot.
func (c *Client) SoftwareImageFileCreate(obj NewSoftwareImageFile) (*SoftwareImageFile, error) {
	return core.Create[SoftwareImageFile, NewSoftwareImageFile](c.Client, dcimEndpointSoftwareImageFile, obj)
}

// SoftwareImageFileDelete : Delete a SoftwareImageFile by UUID identifier.
func (c *Client) SoftwareImageFileDelete(id uuid.UUID) error {
	return core.Delete(c.Client, dcimEndpointSoftwareImageFile, id)
}

// SoftwareImageFileUpdate : Update an existing SoftwareImageFile record in Nautobot.
func (c *Client) SoftwareImageFileUpdate(id uuid.UUID, patch map[string]any) (*SoftwareImageFile, error) {
	return core.Update[SoftwareImageFile](c.Client, dcimEndpointSoftwareImageFile, id, patch)
}

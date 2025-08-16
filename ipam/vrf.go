package ipam

import (
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/dcim"
	"github.com/neverbeencloser/gonautobot/tenancy"
	"github.com/neverbeencloser/gonautobot/types"
)

const (
	ipamEndpointVRF = "ipam/vrfs/"
)

type (
	// VRF : Data type entry for a VRF in Nautobot.
	VRF struct {
		ID                    uuid.UUID       `json:"id"`
		Created               time.Time       `json:"created"`
		CustomFields          map[string]any  `json:"custom_fields"`
		Description           string          `json:"description"`
		Devices               []dcim.Device   `json:"devices"`
		Display               string          `json:"display"`
		ExportTargets         []types.Object  `json:"export_targets"`
		ImportTargets         []types.Object  `json:"import_targets"`
		LastUpdated           time.Time       `json:"last_updated"`
		Name                  string          `json:"name"`
		Namespace             Namespace       `json:"namespace"`
		NaturalSlug           string          `json:"natural_slug"`
		NotesURL              string          `json:"notes_url"`
		ObjectType            string          `json:"object_type"`
		Prefixes              []Prefix        `json:"prefixes"`
		RD                    string          `json:"rd"`
		Status                *types.Status   `json:"status"`
		Tags                  []types.Tag     `json:"tags"`
		Tenant                *tenancy.Tenant `json:"tenant"`
		URL                   string          `json:"url"`
		VirtualDeviceContexts []types.Object  `json:"virtual_device_contexts"`
		VirtualMachines       []types.Object  `json:"virtual_machines"`
	}

	// NewVRF : Structured input for a new VRF record in Nautobot.
	NewVRF struct {
		Name          string         `json:"name"`
		Namespace     string         `json:"namespace"`
		RD            string         `json:"rd"`
		CustomFields  map[string]any `json:"custom_fields,omitempty"`
		Description   string         `json:"description,omitempty"`
		ExportTargets []string       `json:"export_targets,omitempty"`
		ImportTargets []string       `json:"import_targets,omitempty"`
		Relationships map[string]any `json:"relationships,omitempty"`
		Status        string         `json:"status,omitempty"`
		Tags          []string       `json:"tags,omitempty"`
		Tenant        string         `json:"tenant,omitempty"`
	}
)

// VRFGet : Get a VRF by UUID identifier.
func (c *Client) VRFGet(id uuid.UUID) (*VRF, error) {
	return core.Get[VRF](c.Client, ipamEndpointVRF, id)
}

// VRFFilter : Get a list of VRFs based on query parameters.
func (c *Client) VRFFilter(q *url.Values) ([]VRF, error) {
	vrfs := make([]VRF, 0)
	return vrfs, core.Paginate[VRF](c.Client, ipamEndpointVRF, q, &vrfs)
}

// VRFAll : Get all VRFs in Nautobot.
func (c *Client) VRFAll() ([]VRF, error) {
	vrfs := make([]VRF, 0)
	return vrfs, core.Paginate[VRF](c.Client, ipamEndpointVRF, nil, &vrfs)
}

// VRFCreate : Generate a new VRF record in Nautobot.
func (c *Client) VRFCreate(obj NewVRF) (*VRF, error) {
	return core.Create[VRF, NewVRF](c.Client, ipamEndpointVRF, obj)
}

// VRFDelete : Delete a VRF by UUID identifier.
func (c *Client) VRFDelete(id uuid.UUID) error {
	return core.Delete(c.Client, ipamEndpointVRF, id)
}

// VRFUpdate : Update an existing VRF record in Nautobot.
func (c *Client) VRFUpdate(id uuid.UUID, patch map[string]any) (*VRF, error) {
	return core.Update[VRF](c.Client, ipamEndpointVRF, id, patch)
}

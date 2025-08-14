package ipam

import (
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/dcim"
	"github.com/neverbeencloser/gonautobot/extras"
)

const (
	ipamEndpointNamespace = "ipam/namespaces/"
)

type (
	// Namespace : Represents a namespace in Nautobot.
	Namespace struct {
		ID           uuid.UUID      `json:"id"`
		Created      time.Time      `json:"created"`
		CustomFields map[string]any `json:"custom_fields"`
		Description  string         `json:"description"`
		Display      string         `json:"display"`
		LastUpdated  time.Time      `json:"last_updated"`
		Location     *dcim.Location `json:"location"`
		Name         string         `json:"name"`
		NaturalSlug  string         `json:"natural_slug"`
		NotesURL     string         `json:"notes_url"`
		ObjectType   string         `json:"object_type"`
		Tags         []extras.Tag   `json:"tags"`
		URL          string         `json:"url"`
	}

	// NewNamespace : Represents a new namespace to be created in Nautobot.
	NewNamespace struct {
		Name          string         `json:"name"`
		CustomFields  map[string]any `json:"custom_fields,omitempty"`
		Description   string         `json:"description,omitempty"`
		Location      string         `json:"location,omitempty"`
		Relationships map[string]any `json:"relationships,omitempty"`
		Tags          []string       `json:"tags,omitempty"`
	}
)

// NamespaceGet : Get a Namespace by UUID identifier.
func (c *Client) NamespaceGet(id uuid.UUID) (*Namespace, error) {
	return core.Get[Namespace](c.Client, ipamEndpointNamespace, id)
}

// NamespaceFilter : Get a list of Namespaces based on query parameters.
func (c *Client) NamespaceFilter(q *url.Values) ([]Namespace, error) {
	namespaces := make([]Namespace, 0)
	return namespaces, core.Paginate[Namespace](c.Client, ipamEndpointNamespace, q, &namespaces)
}

// NamespaceAll : Get all Namespaces in Nautobot.
func (c *Client) NamespaceAll() ([]Namespace, error) {
	namespaces := make([]Namespace, 0)
	return namespaces, core.Paginate[Namespace](c.Client, ipamEndpointNamespace, nil, &namespaces)
}

// NamespaceCreate : Generate a new Namespace record in Nautobot.
func (c *Client) NamespaceCreate(obj NewNamespace) (*Namespace, error) {
	return core.Create[Namespace, NewNamespace](c.Client, ipamEndpointNamespace, obj)
}

// NamespaceUpdate : Update an existing Namespace record in Nautobot.
func (c *Client) NamespaceUpdate(id uuid.UUID, patch map[string]any) (*Namespace, error) {
	return core.Update[Namespace](c.Client, ipamEndpointNamespace, id, patch)
}

// NamespaceDelete : Delete a Namespace by UUID identifier.
func (c *Client) NamespaceDelete(id uuid.UUID) error {
	return core.Delete(c.Client, ipamEndpointNamespace, id)
}

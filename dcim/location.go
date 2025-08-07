package dcim

import (
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/extras"
	"github.com/neverbeencloser/gonautobot/tenancy"
)

const (
	dcimEndpointLocation = "dcim/locations/"
)

type (
	// Location : Represents a location in Nautobot.
	Location struct {
		ID              uuid.UUID       `json:"id"`
		ASN             *int            `json:"asn"`
		CircuitCount    int             `json:"circuit_count"`
		Comments        string          `json:"comments"`
		ContactEmail    string          `json:"contact_email"`
		ContactName     string          `json:"contact_name"`
		ContactPhone    string          `json:"contact_phone"`
		Created         time.Time       `json:"created"`
		CustomFields    map[string]any  `json:"custom_fields"`
		Description     string          `json:"description"`
		DeviceCount     int             `json:"device_count"`
		Display         string          `json:"display"`
		Facility        string          `json:"facility"`
		LastUpdated     time.Time       `json:"last_updated"`
		Latitude        *float64        `json:"latitude,string"`
		LocationType    LocationType    `json:"location_type"`
		Longitude       *float64        `json:"longitude,string"`
		Name            string          `json:"name"`
		NaturalSlug     string          `json:"natural_slug"`
		NotesURL        string          `json:"notes_url"`
		ObjectType      string          `json:"object_type"`
		Parent          *Location       `json:"parent"`
		PhysicalAddress string          `json:"physical_address"`
		PrefixCount     int             `json:"prefix_count"`
		RackCount       int             `json:"rack_count"`
		ShippingAddress string          `json:"shipping_address"`
		Status          extras.Status   `json:"status"`
		Tags            []extras.Tag    `json:"tags"`
		Tenant          *tenancy.Tenant `json:"tenant"`
		TimeZone        *string         `json:"time_zone"`
		TreeDepth       *int            `json:"tree_depth"`
		URL             string          `json:"url"`
		VMCount         int             `json:"virtual_machine_count"`
		VLANCount       int             `json:"vlan_count"`
	}

	// NewLocation : Represents a new location to be created in Nautobot.
	NewLocation struct {
		Name            string         `json:"name"`
		LocationType    string         `json:"location_type"`
		Status          string         `json:"status"`
		ASN             int            `json:"asn,omitempty"`
		Comments        string         `json:"comments,omitempty"`
		ContactEmail    string         `json:"contact_email,omitempty"`
		ContactName     string         `json:"contact_name,omitempty"`
		ContactPhone    string         `json:"contact_phone,omitempty"`
		CustomFields    map[string]any `json:"custom_fields,omitempty"`
		Description     string         `json:"description,omitempty"`
		Facility        string         `json:"facility,omitempty"`
		Latitude        *float64       `json:"latitude,omitempty"`
		Longitude       *float64       `json:"longitude,omitempty"`
		Parent          string         `json:"parent,omitempty"`
		PhysicalAddress string         `json:"physical_address,omitempty"`
		Relationships   map[string]any `json:"relationships,omitempty"`
		ShippingAddress string         `json:"shipping_address,omitempty"`
		Tags            []string       `json:"tags,omitempty"`
		Tenant          string         `json:"tenant,omitempty"`
		TimeZone        string         `json:"time_zone,omitempty"`
	}
)

// LocationGet : Get a Location by UUID identifier.
func (c *Client) LocationGet(id uuid.UUID) (*Location, error) {
	return core.Get[Location](c.Client, dcimEndpointLocation, id)
}

// LocationFilter : Get a list of Locations based on query parameters.
func (c *Client) LocationFilter(q *url.Values) ([]Location, error) {
	locationTypes := make([]Location, 0)
	return locationTypes, core.Paginate[Location](c.Client, dcimEndpointLocation, q, &locationTypes)
}

// LocationAll : Get all Locations in Nautobot.
func (c *Client) LocationAll() ([]Location, error) {
	locationTypes := make([]Location, 0)
	return locationTypes, core.Paginate[Location](c.Client, dcimEndpointLocation, nil, &locationTypes)
}

// LocationCreate : Generate a new Location record in Nautobot.
func (c *Client) LocationCreate(obj NewLocation) (*Location, error) {
	return core.Create[Location, NewLocation](c.Client, dcimEndpointLocation, obj)
}

// LocationUpdate : Update an existing Location record in Nautobot.
func (c *Client) LocationUpdate(id uuid.UUID, patch map[string]any) (*Location, error) {
	return core.Update[Location](c.Client, dcimEndpointLocation, id, patch)
}

// LocationDelete : Delete a Location by UUID identifier.
func (c *Client) LocationDelete(id uuid.UUID) error {
	return core.Delete(c.Client, dcimEndpointLocation, id)
}

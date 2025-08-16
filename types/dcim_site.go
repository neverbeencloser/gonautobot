package types

import (
	"encoding/json"

	"github.com/neverbeencloser/gonautobot/types/nested"
)

type (
	// Site : Data representation of a Site model in Nautobot.
	Site struct {
		ASN                 *int            `json:"asn"`
		CircuitCount        int             `json:"circuit_count"`
		Comments            string          `json:"comments" datastore:",noindex"`
		ContactEmail        string          `json:"contact_email"`
		ContactName         string          `json:"contact_name"`
		ContactPhone        string          `json:"contact_phone"`
		Created             string          `json:"created"`
		CustomFields        map[string]any  `json:"custom_fields"`
		Description         string          `json:"description"`
		DeviceCount         int             `json:"device_count"`
		Display             string          `json:"display"`
		Facility            string          `json:"facility"`
		ID                  string          `json:"id"`
		LastUpdated         string          `json:"last_updated"`
		Latitude            json.RawMessage `json:"latitude"`
		Longitude           json.RawMessage `json:"longitude"`
		Name                string          `json:"name"`
		NotesURL            string          `json:"notes_url"`
		PhysicalAddress     string          `json:"physical_address"`
		PrefixCount         int             `json:"prefix_count"`
		RackCount           int             `json:"rack_count"`
		Region              *Region         `json:"region"`
		ShippingAddress     string          `json:"shipping_address"`
		Slug                string          `json:"slug"`
		Status              *LabelValue     `json:"status"`
		Tags                []Tag           `json:"tags"`
		Tenant              *nested.Tenant  `json:"tenant"`
		TimeZone            json.RawMessage `json:"time_zone"`
		URL                 string          `json:"url"`
		VirtualMachineCount int             `json:"virtualmachine_count"`
		VLANCount           int             `json:"vlan_count"`
	}
)

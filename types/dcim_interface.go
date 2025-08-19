package types

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/types/nested"
)

type (
	// DeviceInterface : defines an interface as represented in Nautobot
	//
	// CablePeer and ConnectedEndpoint must be decoded dynamically; they are dependent
	// on the value of 'CablePeerType' and 'ConnectedEndpointType', e.g., "dcim.interface"
	DeviceInterface struct {
		ID                         uuid.UUID        `json:"id"`
		Bridge                     *DeviceInterface `json:"bridge"`
		Cable                      *nested.Cable    `json:"cable"`
		CablePeer                  json.RawMessage  `json:"cable_peer"`
		CablePeerType              *string          `json:"cable_peer_type"`
		ConnectedEndpoint          json.RawMessage  `json:"connected_endpoint"`
		ConnectedEndpointReachable bool             `json:"connected_endpoint_reachable"`
		ConnectedEndpointType      *string          `json:"connected_endpoint_type"`
		CountIPAddresses           int              `json:"count_ipaddresses"`
		Created                    string           `json:"created"`
		CustomFields               map[string]any   `json:"custom_fields"`
		Description                string           `json:"description"`
		Device                     *Device          `json:"device"`
		Display                    string           `json:"display"`
		Enabled                    bool             `json:"enabled"`
		Label                      string           `json:"label"`
		Lag                        *DeviceInterface `json:"lag"`
		LastUpdated                string           `json:"last_updated"`
		MACAddress                 *string          `json:"mac_address"`
		MgmtOnly                   bool             `json:"mgmt_only"`
		Mode                       *LabelValue      `json:"mode"`
		MTU                        *int             `json:"mtu"`
		Name                       string           `json:"name"`
		NotesURL                   string           `json:"notes_url"`
		ParentInterface            *DeviceInterface `json:"parent_interface"`
		TaggedVLANs                []VLAN           `json:"tagged_vlans"`
		Tags                       []Tag            `json:"tags"`
		Type                       *LabelValue      `json:"type"`
		UntaggedVLAN               *VLAN            `json:"untagged_vlan"`
		URL                        string           `json:"url"`
	}
)

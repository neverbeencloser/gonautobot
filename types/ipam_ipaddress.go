package types

import (
	"time"

	"github.com/google/uuid"
)

const (
	// IPAddressTypeDHCP : IP Address assigned by DHCP
	IPAddressTypeDHCP IPAddressType = "dhcp"
	// IPAddressTypeHost : IP Address assigned statically to a host
	IPAddressTypeHost IPAddressType = "host"
	// IPAddressTypeSLAAC : IP Address assigned by SLAAC
	IPAddressTypeSLAAC IPAddressType = "slaac"
)

type (
	// IPAddress : defines an IP Address entry in Nautobot
	IPAddress struct {
		ID             uuid.UUID         `json:"id"`
		Address        string            `json:"address"`
		Created        time.Time         `json:"created"`
		CustomFields   map[string]any    `json:"custom_fields"`
		Description    string            `json:"description"`
		Display        string            `json:"display"`
		DNSName        string            `json:"dns_name"`
		Host           string            `json:"host"`
		Interfaces     []DeviceInterface `json:"interfaces"`
		IPVersion      int               `json:"ip_version"`
		LastUpdated    time.Time         `json:"last_updated"`
		MaskLength     int               `json:"mask_length"`
		NATInside      *IPAddress        `json:"nat_inside"`
		NATOutsideList []IPAddress       `json:"nat_outside"`
		NaturalSlug    string            `json:"natural_slug"`
		NotesURL       string            `json:"notes_url"`
		ObjectType     string            `json:"object_type"`
		Parent         Prefix            `json:"parent"`
		Role           *Role             `json:"role"`
		Status         *Status           `json:"status"`
		Tags           []Tag             `json:"tags"`
		Tenant         *Tenant           `json:"tenant"`
		Type           string            `json:"type"`
		URL            string            `json:"url"`
		VMInterfaces   []Object          `json:"vm_interfaces"` // TODO: update to VMInterface type when available
	}

	// NewIPAddress : defines the fields required to create a new IP Address entry in Nautobot
	NewIPAddress struct {
		Address       string         `json:"address"`
		Namespace     string         `json:"namespace"`
		Status        string         `json:"status"`
		CustomFields  map[string]any `json:"custom_fields,omitempty"`
		Description   string         `json:"description,omitempty"`
		DNSName       string         `json:"dns_name,omitempty"`
		IPVersion     int            `json:"ip_version,omitempty"`
		NATInside     string         `json:"nat_inside,omitempty"` // IPAddress UUID
		Parent        string         `json:"parent,omitempty"`     // Prefix UUID
		Relationships map[string]any `json:"relationships,omitempty"`
		Role          string         `json:"role,omitempty"`
		Tenant        string         `json:"tenant,omitempty"`
		Tags          []string       `json:"tags,omitempty"`
		Type          IPAddressType  `json:"type,omitempty"`
	}

	// IPAddressType : Enum for NewIPAddress.Type field
	IPAddressType string
)

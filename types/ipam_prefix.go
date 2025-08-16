package types

import "github.com/neverbeencloser/gonautobot/types/nested"

type (
	// Prefix : defines an IPAM Prefix entry in Nautobot
	Prefix struct {
		Created      string                 `json:"created"`
		CustomFields map[string]interface{} `json:"custom_fields"`
		Description  string                 `json:"description"`
		Display      string                 `json:"display"`
		Family       *LabelValueInt         `json:"family"`
		Group        *nested.SecretsGroup   `json:"group"`
		ID           string                 `json:"id"`
		IsPool       bool                   `json:"is_pool"`
		LastUpdated  string                 `json:"last_updated"`
		Location     *nested.Location       `json:"location"`
		Name         string                 `json:"name"`
		NotesURL     string                 `json:"notes_url"`
		Prefix       string                 `json:"prefix"`
		Role         *nested.Role           `json:"role"`
		Site         *nested.Site           `json:"site"`
		Status       *LabelValue            `json:"status"`
		Tags         []Tag                  `json:"tags"`
		Tenant       *nested.Tenant         `json:"tenant"`
		URL          string                 `json:"url"`
		VLAN         *nested.VLAN           `json:"vlan"`
		VRF          *nested.VRF            `json:"vrf"`
	}

	// PrefixAvailableIP : stub IPAddress entry returned by the /prefixes/:id/available-ips/ methods
	PrefixAvailableIP struct {
		Address string `json:"address"`
		Family  int    `json:"family"`
		VRF     *VRF   `json:"vrf"`
	}

	// PrefixAvailablePrefix : stub Prefix entry returned by the /prefixes/:id/available-prefixes/ methods
	PrefixAvailablePrefix struct {
		Family int    `json:"family"`
		Prefix string `json:"prefix"`
		VRF    *VRF   `json:"vrf"`
	}

	// NewPrefix : The data structure required to create a new Prefix object in Nautobot.
	NewPrefix struct {
		Description string `json:"description"`
		ID          string `json:"id"`
		IsPool      bool   `json:"is_pool"`
		Prefix      string `json:"prefix"`
		Status      string `json:"status"`
		Tags        []Tag  `json:"tags"`
		VRF         *VRF   `json:"vrf"`
	}

	// PrefixUpdate : defines writable fields to update a Prefix entry in nautobot
	PrefixUpdate struct {
		Prefix      string `json:"prefix"`
		IsPool      bool   `json:"is_pool"`
		Tags        []Tag  `json:"tags"`
		Description string `json:"description"`
		VRF         *VRF   `json:"vrf"`
		Site        string `json:"site"`
		Location    string `json:"location"`
		Tenant      string `json:"tenant"`
		VLAN        string `json:"vlan"`
	}

	// PrefixesAvailablePrefixesCreateRequest : required parameters for PrefixesAvailablePrefixesCreate
	PrefixesAvailablePrefixesCreateRequest struct {
		PrefixLength int    `json:"prefix_length"`
		Status       Status `json:"status"`
	}
)

package types

import "github.com/neverbeencloser/gonautobot/types/nested"

type (
	// VLAN : defines a vlan entry in Nautobot
	VLAN struct {
		ID           string                 `json:"id"`
		Created      string                 `json:"created"`
		CustomFields map[string]interface{} `json:"custom_fields"`
		Description  string                 `json:"description"`
		Display      string                 `json:"display"`
		Group        *nested.SecretsGroup   `json:"group"`
		LastUpdated  string                 `json:"last_updated"`
		Location     *nested.Location       `json:"location"`
		Name         string                 `json:"name"`
		NotesURL     string                 `json:"notes_url"`
		PrefixCount  int                    `json:"prefix_count"`
		Role         *nested.Role           `json:"role"`
		Site         *nested.Site           `json:"site"`
		Status       *LabelValue            `json:"status"`
		Tags         []Tag                  `json:"tags"`
		Tenant       *nested.Tenant         `json:"tenant"`
		URL          string                 `json:"url"`
		VID          int                    `json:"vid"`
	}
)

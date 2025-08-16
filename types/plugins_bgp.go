package types

import "github.com/neverbeencloser/gonautobot/types/nested"

type (
	// AutonomousSystem : Autonomous System data struct representation.
	AutonomousSystem struct {
		ASN          int                    `json:"asn"`
		Created      string                 `json:"created"`
		CustomFields map[string]interface{} `json:"custom_fields"`
		Description  string                 `json:"description"`
		Display      string                 `json:"display"`
		ID           string                 `json:"id"`
		LastUpdated  string                 `json:"last_updated"`
		Provider     *nested.Provider       `json:"provider"`
		Status       *LabelValue            `json:"status"`
		Tags         []Tag                  `json:"tags"`
		URL          string                 `json:"url"`
	}

	// RoutingInstance : Routing Instance data representation in Nautobot.
	RoutingInstance struct {
		AutonomousSystem *AutonomousSystem        `json:"autonomous_system"`
		Created          string                   `json:"created"`
		CustomFields     map[string]interface{}   `json:"custom_fields"`
		Description      string                   `json:"description"`
		Device           *nested.Device           `json:"device"`
		Display          string                   `json:"display"`
		Endpoints        []nested.BGPPeerEndpoint `json:"endpoints"`
		ExtraAttributes  map[string]interface{}   `json:"extra_attributes"`
		ID               string                   `json:"id"`
		LastUpdated      string                   `json:"last_updated"`
		RouterID         *nested.IPAddress        `json:"router_id"`
		Status           *LabelValue              `json:"status"`
		Tags             []Tag                    `json:"tags"`
		URL              string                   `json:"url"`
	}
)

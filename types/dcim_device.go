package types

import (
	"time"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/types/nested"
)

type (
	// Device : Data structure to represent a device record in Nautobot.
	Device struct {
		ID                 uuid.UUID                   `json:"id"`
		AssetTag           string                      `json:"asset_tag"`
		Cluster            *Cluster                    `json:"cluster"`
		Comments           string                      `json:"comments" datastore:",noindex"`
		ConfigContext      map[string]any              `json:"config_context"`
		Created            time.Time                   `json:"created"`
		CustomFields       map[string]any              `json:"custom_fields"`
		DeviceRole         *Role                       `json:"device_role"`
		DeviceType         *DeviceType                 `json:"device_type"`
		Display            string                      `json:"display"`
		Face               *LabelValue                 `json:"face"`
		LastUpdated        time.Time                   `json:"last_updated"`
		LocalContextSchema *nested.ConfigContextSchema `json:"local_context_schema"`
		LocalContextData   map[string]any              `json:"local_context_data"`
		Location           *Location                   `json:"location"`
		Name               string                      `json:"name"`
		NotesURL           string                      `json:"notes_url"`
		ParentDevice       *Device                     `json:"parent_device"`
		Platform           *Platform                   `json:"platform"`
		Position           *int                        `json:"position"`
		PrimaryIP          *IPAddress                  `json:"primary_ip"`
		PrimaryIP4         *IPAddress                  `json:"primary_ip4"`
		PrimaryIP6         *IPAddress                  `json:"primary_ip6"`
		Rack               *Rack                       `json:"rack"`
		SecretsGroup       *nested.SecretsGroup        `json:"secrets_group"`
		Serial             string                      `json:"serial"`
		Site               *Site                       `json:"site"`
		Status             *LabelValue                 `json:"status"`
		Tags               []Tag                       `json:"tags"`
		Tenant             *Tenant                     `json:"tenant"`
		URL                string                      `json:"url"`
		VCPosition         *int                        `json:"vc_position"`
		VCPriority         *int                        `json:"vc_priority"`
		VirtualChassis     *nested.VirtualChassis      `json:"virtual_chassis"`
	}

	// NewDevice : Structured input for a new Device record in Nautobot.
	NewDevice struct {
		Name                         string         `json:"name"`
		Role                         string         `json:"role"`
		Status                       string         `json:"status"`
		DeviceType                   string         `json:"device_type"`
		Location                     string         `json:"location,omitempty"`
		Site                         string         `json:"site,omitempty"`
		Tenant                       string         `json:"tenant,omitempty"`
		Platform                     string         `json:"platform,omitempty"`
		Serial                       string         `json:"serial,omitempty"`
		AssetTag                     string         `json:"asset_tag,omitempty"`
		Position                     int            `json:"position,omitempty"`
		Face                         string         `json:"face,omitempty"`
		VcPosition                   int            `json:"vc_position,omitempty"`
		VcPriority                   int            `json:"vc_priority,omitempty"`
		Comments                     string         `json:"comments,omitempty"`
		Rack                         string         `json:"rack,omitempty"`
		PrimaryIP4                   string         `json:"primary_ip4,omitempty"`
		PrimaryIP6                   string         `json:"primary_ip6,omitempty"`
		Cluster                      string         `json:"cluster,omitempty"`
		VirtualChassis               string         `json:"virtual_chassis,omitempty"`
		DeviceRedundancyGroup        string         `json:"device_redundancy_group,omitempty"`
		SoftwareVersion              string         `json:"software_version,omitempty"`
		SecretsGroup                 string         `json:"secrets_group,omitempty"`
		ControllerManagedDeviceGroup string         `json:"controller_managed_device_group,omitempty"`
		SoftwareImageFiles           []string       `json:"software_image_files,omitempty"`
		CustomFields                 map[string]any `json:"custom_fields,omitempty"`
		Tags                         []string       `json:"tags,omitempty"`
		ParentBay                    string         `json:"parent_bay,omitempty"`
	}
)

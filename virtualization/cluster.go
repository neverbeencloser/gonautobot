package virtualization

import (
	"time"

	"github.com/neverbeencloser/gonautobot/dcim"
	"github.com/neverbeencloser/gonautobot/tenancy"
)

// Cluster : Cerebro Cluster data representation in Nautobot.
type Cluster struct {
	ID          string `json:"id"`
	ObjectType  string `json:"object_type"`
	Display     string `json:"display"`
	URL         string `json:"url"`
	NaturalSlug string `json:"natural_slug"`
	Name        string `json:"name"`
	Comments    string `json:"comments"`
	ClusterType struct {
		ID         string `json:"id"`
		ObjectType string `json:"object_type"`
		URL        string `json:"url"`
	} `json:"cluster_type"`
	ClusterGroup any             `json:"cluster_group"`
	Tenant       *tenancy.Tenant `json:"tenant"`
	Location     *dcim.Location  `json:"location"`
	Created      time.Time       `json:"created"`
	LastUpdated  time.Time       `json:"last_updated"`
	NotesURL     string          `json:"notes_url"`
	CustomFields map[string]any  `json:"custom_fields"`
}

package types

import (
	"time"

	"github.com/neverbeencloser/gonautobot/types/nested"
)

type (
	// Circuit : defines a circuit entry in Nautobot
	Circuit struct {
		ID           string             `json:"id"`
		Display      string             `json:"display"`
		URL          string             `json:"url"`
		CircuitID    string             `json:"cid"`
		Provider     nested.Provider    `json:"provider"`
		Type         nested.CircuitType `json:"type"`
		Status       LabelValue         `json:"status"`
		Tenant       nested.Tenant      `json:"tenant"`
		InstallDate  string             `json:"install_date"`
		CommitRate   int                `json:"commit_rate"`
		Description  string             `json:"description"`
		TerminationA Termination        `json:"termination_a"`
		TerminationZ Termination        `json:"termination_z"`
		Comments     string             `json:"comments"`
		Created      string             `json:"created"`
		LastUpdated  time.Time          `json:"last_updated"`
		Tags         []Tag              `json:"tags"`
		NotesURL     string             `json:"notes_url"`
		CustomFields map[string]any     `json:"custom_fields"`
	}

	// CircuitRequest : Models a new circuit entry in Nautobot.
	CircuitRequest struct {
		CircuitID    string         `json:"cid"`
		Provider     string         `json:"provider"`
		Type         string         `json:"circuit_type"`
		Status       string         `json:"status"`
		Tenant       string         `json:"tenant,omitempty"`
		InstallDate  string         `json:"install_date,omitempty"`
		CommitRate   int            `json:"commit_rate,omitempty"`
		Description  string         `json:"description,omitempty"`
		TerminationA Termination    `json:"termination_a,omitempty"`
		TerminationZ Termination    `json:"termination_z,omitempty"`
		Comments     string         `json:"comments,omitempty"`
		Created      string         `json:"created,omitempty"`
		Tags         []Tag          `json:"tags,omitempty"`
		CustomFields map[string]any `json:"custom_fields,omitempty"`
	}
)

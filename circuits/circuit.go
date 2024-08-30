package circuits

import (
	"encoding/json"
	"fmt"
	"github.com/josh-silvas/gonautobot/extras"
	"github.com/josh-silvas/gonautobot/shared"
	"github.com/josh-silvas/gonautobot/shared/nested"
	"net/http"
	"net/url"
	"time"
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
		Status       shared.LabelValue  `json:"status"`
		Tenant       nested.Tenant      `json:"tenant"`
		InstallDate  string             `json:"install_date"`
		CommitRate   int                `json:"commit_rate"`
		Description  string             `json:"description"`
		TerminationA Termination        `json:"termination_a"`
		TerminationZ Termination        `json:"termination_z"`
		Comments     string             `json:"comments"`
		Created      string             `json:"created"`
		LastUpdated  time.Time          `json:"last_updated"`
		Tags         []extras.Tag       `json:"tags"`
		NotesURL     string             `json:"notes_url"`
		CustomFields map[string]any     `json:"custom_fields"`
	}
)

// GetCircuits : Go function to process requests for the endpoint: /api/circuits/circuits/
//
// https://demo.nautobot.com/api/docs/#/circuits/circuits_circuits_list
func (c *Client) GetCircuits(q *url.Values) ([]Circuit, error) {
	req, err := c.Request(http.MethodGet, "circuits/circuits/", nil, q)
	if err != nil {
		return nil, err
	}

	resp := new(shared.ResponseList)
	ret := make([]Circuit, 0)
	if err = c.UnmarshalDo(req, resp); err != nil {
		return ret, err
	}

	if err = json.Unmarshal(resp.Results, &ret); err != nil {
		err = fmt.Errorf("GetCircuits.error.json.Unmarshal(%w)", err)
	}
	return ret, err
}

package extras

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/josh-silvas/gonautobot/core"
)

type (
	// Status : Data type entry for a Status in Nautobot.
	Status struct {
		ID           uuid.UUID      `json:"id"`
		Color        string         `json:"color"`
		ContentTypes []string       `json:"content_types"`
		Created      time.Time      `json:"created"`
		CustomFields map[string]any `json:"custom_fields"`
		Description  string         `json:"description"`
		Display      string         `json:"display"`
		LastUpdated  time.Time      `json:"last_updated"`
		Name         string         `json:"name"`
		NaturalSlug  string         `json:"natural_slug"`
		NotesURL     string         `json:"notes_url"`
		ObjectType   string         `json:"object_type"`
		URL          string         `json:"url"`
	}

	// NewStatus : Structured input for a new Status record in Nautobot.
	NewStatus struct {
		Name         string         `json:"name"`
		ContentTypes []string       `json:"content_types"`
		Color        string         `json:"color,omitempty"`
		Description  string         `json:"description,omitempty"`
		CustomFields map[string]any `json:"custom_fields,omitempty"`
	}
)

// StatusGet : Go function to process requests for the endpoint: /api/extras/statuses/:id/
func (c *Client) StatusGet(id uuid.UUID) (*Status, error) {
	if id == uuid.Nil {
		return nil, errors.New("StatusGet.error.ID(ID is missing or nil)")
	}
	req, err := c.Request(http.MethodGet, fmt.Sprintf("extras/statuses/%s/", id), nil, nil)
	if err != nil {
		return nil, err
	}

	ret := new(Status)
	err = c.UnmarshalDo(req, ret)
	return ret, err
}

// StatusFilter : Go function to process requests for the endpoint: /api/extras/statuses/
func (c *Client) StatusFilter(q *url.Values) ([]Status, error) {
	statuses := make([]Status, 0)
	return statuses, core.Paginate[Status](c.Client, "extras/statuses/", q, &statuses)
}

// StatusAll : Go function to process requests for the endpoint: /api/extras/statuses/
func (c *Client) StatusAll() ([]Status, error) {
	statuses := make([]Status, 0)
	return statuses, core.Paginate[Status](c.Client, "extras/statuses/", nil, &statuses)
}

// StatusCreate : Generate a new Status record in Nautobot.
func (c *Client) StatusCreate(obj NewStatus) (Status, error) {
	var r Status
	req, err := c.Request(http.MethodPost, "extras/statuses/", obj, nil)
	if err != nil {
		return r, err
	}

	if err := c.UnmarshalDo(req, &r); err != nil {
		return r, fmt.Errorf("StatusCreate.error.UnmarshalDo(%w)", err)
	}
	return r, nil
}

// StatusDelete : Delete a Status by UUID identifier.
func (c *Client) StatusDelete(id uuid.UUID) error {
	if id == uuid.Nil {
		return errors.New("StatusDelete.error.ID(ID is missing or nil)")
	}
	req, err := c.Request(http.MethodDelete, fmt.Sprintf("extras/statuses/%s/", id), nil, nil)
	if err != nil {
		return err
	}
	return c.UnmarshalDo(req, nil)
}

// StatusUpdate : Update an existing Status record in Nautobot.
func (c *Client) StatusUpdate(id uuid.UUID, patch map[string]any) (Status, error) {
	var r Status
	if id == uuid.Nil {
		return r, errors.New("StatusUpdate.error.ID(ID is missing or nil)")
	}
	req, err := c.Request(http.MethodPatch, fmt.Sprintf("extras/statuses/%s/", id), patch, nil)
	if err != nil {
		return r, err
	}
	if err := c.UnmarshalDo(req, &r); err != nil {
		return r, fmt.Errorf("StatusUpdate.error.UnmarshalDo(%w)", err)
	}
	return r, nil
}

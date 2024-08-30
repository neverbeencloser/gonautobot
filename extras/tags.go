package extras

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/josh-silvas/gonautobot/shared"
	"net/http"
	"net/url"
)

type (
	// Tag : Data type entry for a Tag in Nautobot.
	Tag struct {
		ID          string `json:"id"`
		Color       string `json:"color"`
		Created     string `json:"created"`
		Display     string `json:"display"`
		LastUpdated string `json:"last_updated"`
		Name        string `json:"name"`
		Slug        string `json:"slug"`
		URL         string `json:"url"`
	}
)

// GetTag : Fetches a tag from Nautobot by the slug name.
func (c *Client) GetTag(uuid string, q *url.Values) (Tag, error) {
	req, err := c.Request(http.MethodGet, fmt.Sprintf("extras/tags/%s/", url.PathEscape(uuid)), nil, q)
	if err != nil {
		return Tag{}, err
	}

	ret := new(Tag)
	if err = c.UnmarshalDo(req, ret); err != nil {
		return Tag{}, err
	}

	return *ret, nil
}

// GetTags : Returns an array of tags from Nautobot filtered by the q (query values)
func (c *Client) GetTags(q *url.Values) ([]Tag, error) {
	req, err := c.Request(http.MethodGet, "extras/tags/", nil, q)
	if err != nil {
		return nil, err
	}

	resp := new(shared.ResponseList)
	ret := make([]Tag, 0)
	if err := c.UnmarshalDo(req, resp); err != nil {
		return nil, fmt.Errorf("GetTags.error.UnmarshalDo(%w)", err)
	}
	if err = json.Unmarshal(resp.Results, &ret); err != nil {
		return nil, fmt.Errorf("GetTags.error.json.Unmarshal(%w)", err)
	}
	return ret, nil
}

// CreateTag creates a new tag if it doesn't exist and returns an error if it does.
// It does not check if the tag exists. That can be done with TagExists().
func (c *Client) CreateTag(tag Tag) (*Tag, error) {
	req, err := c.Request(http.MethodPost, "extras/tags/", tag, nil)
	if err != nil {
		return nil, err
	}

	var r Tag
	if err := c.UnmarshalDo(req, &r); err != nil {
		return nil, fmt.Errorf("CreateTag.error.UnmarshalDo(%w)", err)
	}
	return &r, nil
}

// DeleteTag : Extras method to delete a tag by UUIDv4 identifier.
func (c *Client) DeleteTag(uuid string) error {
	if uuid == "" {
		return errors.New("DeleteTag.error.UUID(UUIDv4 is missing)")
	}
	req, err := c.Request(http.MethodDelete, "extras/tags/", shared.BulkDelete{{ID: uuid}}, nil)
	if err != nil {
		return err
	}
	return c.UnmarshalDo(req, nil)
}

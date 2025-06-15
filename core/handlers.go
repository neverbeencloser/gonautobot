package core

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

// Create : Generic function to perform a POST request to Nautobot.
func Create[T any, R any](c *Client, uri string, body R) (*T, error) {
	var resp T

	req, err := c.Request(http.MethodPost, uri, body, nil)
	if err != nil {
		return nil, err
	}
	if err := c.UnmarshalDo(req, &resp); err != nil {
		return nil, fmt.Errorf("%s Create.error.UnmarshalDo(%w)", uri, err)
	}
	return &resp, nil
}

// Delete : Generic function to perform a DELETE request to Nautobot.
func Delete(c *Client, uri string, id uuid.UUID) error {
	if id == uuid.Nil {
		return fmt.Errorf("%s Delete.error.ID(ID is missing or nil)", uri)
	}
	req, err := c.Request(http.MethodDelete, fmt.Sprintf("%s%s/", uri, id), nil, nil)
	if err != nil {
		return err
	}
	return c.UnmarshalDo(req, nil)
}

// Get : Generic function to perform a GET request to Nautobot for a single object.
func Get[T any](c *Client, uri string, id uuid.UUID) (*T, error) {
	if id == uuid.Nil {
		return nil, fmt.Errorf("%s Get.error.ID(ID is missing or nil)", uri)
	}

	req, err := c.Request(http.MethodGet, fmt.Sprintf("%s%s/", uri, id), nil, nil)
	if err != nil {
		return nil, err
	}
	var resp T
	if err := c.UnmarshalDo(req, &resp); err != nil {
		return nil, fmt.Errorf("%s Get.error.UnmarshalDo(%w)", uri, err)
	}
	return &resp, nil
}

// Update : Generic function to perform a PATCH request to Nautobot.
func Update[T any](c *Client, uri string, id uuid.UUID, patch map[string]any) (*T, error) {
	var resp T
	if id == uuid.Nil {
		return nil, fmt.Errorf("%s Update.error.ID(ID is missing or nil)", uri)
	}

	req, err := c.Request(http.MethodPatch, fmt.Sprintf("%s%s/", uri, id), patch, nil)
	if err != nil {
		return nil, err
	}
	if err := c.UnmarshalDo(req, &resp); err != nil {
		return nil, fmt.Errorf("%s Update.error.UnmarshalDo(%w)", uri, err)
	}
	return &resp, nil
}

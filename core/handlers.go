package core

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/shared"
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

// CreateMultipart : Generic function to perform a POST request with multipart/form-data to Nautobot.
//
// This should only be used on API endpoints that support
// multipart form data and image uploads, such as 'dcim/device-types/'.
func CreateMultipart[T any, R any](c *Client, uri string, body R) (*T, error) {
	var resp T

	m, contentType, err := shared.NewMultipartBody(body)
	if err != nil {
		return nil, fmt.Errorf("%s CreateMultipart.error.NewMultipartBody(%w)", uri, err)
	}

	req, err := c.Request(http.MethodPost, uri, m, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)

	if err := c.UnmarshalDo(req, &resp); err != nil {
		return nil, fmt.Errorf("%s CreateMultipart.error.UnmarshalDo(%w)", uri, err)
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

// Action : Generic function to perform a POST request to a resource action endpoint.
// This is used for endpoints like {uri}{id}/{action}/ (e.g., extras/jobs/{id}/run/).
func Action[T any, R any](c *Client, uri string, id uuid.UUID, action string, body R) (*T, error) {
	var resp T
	if id == uuid.Nil {
		return nil, fmt.Errorf("%s Action.error.ID(ID is missing or nil)", uri)
	}

	// Build the action URI: {base_uri}{uuid}/{action}/ (e.g., "extras/jobs/abc-123/run/")
	actionURI := fmt.Sprintf("%s%s/%s/", uri, id, action)
	req, err := c.Request(http.MethodPost, actionURI, body, nil)
	if err != nil {
		return nil, err
	}
	if err := c.UnmarshalDo(req, &resp); err != nil {
		return nil, fmt.Errorf("%s Action.error.UnmarshalDo(%w)", uri, err)
	}
	return &resp, nil
}

// UpdateMultipart : Generic function to perform a PATCH request with multipart/form-data to Nautobot.
//
// This should only be used on API endpoints that support
// multipart form data and image uploads, such as 'dcim/device-types/'.
func UpdateMultipart[T any, R any](c *Client, uri string, id uuid.UUID, body R) (*T, error) {
	var resp T

	m, contentType, err := shared.NewMultipartBody(body)
	if err != nil {
		return nil, fmt.Errorf("%s UpdateMultipart.error.NewMultipartBody(%w)", uri, err)
	}

	req, err := c.Request(http.MethodPatch, fmt.Sprintf("%s%s/", uri, id), m, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)

	if err := c.UnmarshalDo(req, &resp); err != nil {
		return nil, fmt.Errorf("%s UpdateMultipart.error.UnmarshalDo(%w)", uri, err)
	}
	return &resp, nil
}

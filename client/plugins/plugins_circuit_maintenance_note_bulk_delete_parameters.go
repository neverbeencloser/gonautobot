package plugins

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewPluginsCircuitMaintenanceNoteBulkDeleteParams creates a new PluginsCircuitMaintenanceNoteBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsCircuitMaintenanceNoteBulkDeleteParams() *PluginsCircuitMaintenanceNoteBulkDeleteParams {
	return &PluginsCircuitMaintenanceNoteBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsCircuitMaintenanceNoteBulkDeleteParamsWithTimeout creates a new PluginsCircuitMaintenanceNoteBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewPluginsCircuitMaintenanceNoteBulkDeleteParamsWithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceNoteBulkDeleteParams {
	return &PluginsCircuitMaintenanceNoteBulkDeleteParams{
		timeout: timeout,
	}
}

// NewPluginsCircuitMaintenanceNoteBulkDeleteParamsWithContext creates a new PluginsCircuitMaintenanceNoteBulkDeleteParams object
// with the ability to set a context for a request.
func NewPluginsCircuitMaintenanceNoteBulkDeleteParamsWithContext(ctx context.Context) *PluginsCircuitMaintenanceNoteBulkDeleteParams {
	return &PluginsCircuitMaintenanceNoteBulkDeleteParams{
		Context: ctx,
	}
}

// NewPluginsCircuitMaintenanceNoteBulkDeleteParamsWithHTTPClient creates a new PluginsCircuitMaintenanceNoteBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsCircuitMaintenanceNoteBulkDeleteParamsWithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceNoteBulkDeleteParams {
	return &PluginsCircuitMaintenanceNoteBulkDeleteParams{
		HTTPClient: client,
	}
}

/* PluginsCircuitMaintenanceNoteBulkDeleteParams contains all the parameters to send to the API endpoint
   for the plugins circuit maintenance note bulk delete operation.

   Typically these are written to a http.Request.
*/
type PluginsCircuitMaintenanceNoteBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins circuit maintenance note bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceNoteBulkDeleteParams) WithDefaults() *PluginsCircuitMaintenanceNoteBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins circuit maintenance note bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceNoteBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins circuit maintenance note bulk delete params
func (o *PluginsCircuitMaintenanceNoteBulkDeleteParams) WithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceNoteBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins circuit maintenance note bulk delete params
func (o *PluginsCircuitMaintenanceNoteBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins circuit maintenance note bulk delete params
func (o *PluginsCircuitMaintenanceNoteBulkDeleteParams) WithContext(ctx context.Context) *PluginsCircuitMaintenanceNoteBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins circuit maintenance note bulk delete params
func (o *PluginsCircuitMaintenanceNoteBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins circuit maintenance note bulk delete params
func (o *PluginsCircuitMaintenanceNoteBulkDeleteParams) WithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceNoteBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins circuit maintenance note bulk delete params
func (o *PluginsCircuitMaintenanceNoteBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsCircuitMaintenanceNoteBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

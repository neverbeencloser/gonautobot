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

// NewPluginsCircuitMaintenanceCircuitimpactBulkDeleteParams creates a new PluginsCircuitMaintenanceCircuitimpactBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsCircuitMaintenanceCircuitimpactBulkDeleteParams() *PluginsCircuitMaintenanceCircuitimpactBulkDeleteParams {
	return &PluginsCircuitMaintenanceCircuitimpactBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsCircuitMaintenanceCircuitimpactBulkDeleteParamsWithTimeout creates a new PluginsCircuitMaintenanceCircuitimpactBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewPluginsCircuitMaintenanceCircuitimpactBulkDeleteParamsWithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceCircuitimpactBulkDeleteParams {
	return &PluginsCircuitMaintenanceCircuitimpactBulkDeleteParams{
		timeout: timeout,
	}
}

// NewPluginsCircuitMaintenanceCircuitimpactBulkDeleteParamsWithContext creates a new PluginsCircuitMaintenanceCircuitimpactBulkDeleteParams object
// with the ability to set a context for a request.
func NewPluginsCircuitMaintenanceCircuitimpactBulkDeleteParamsWithContext(ctx context.Context) *PluginsCircuitMaintenanceCircuitimpactBulkDeleteParams {
	return &PluginsCircuitMaintenanceCircuitimpactBulkDeleteParams{
		Context: ctx,
	}
}

// NewPluginsCircuitMaintenanceCircuitimpactBulkDeleteParamsWithHTTPClient creates a new PluginsCircuitMaintenanceCircuitimpactBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsCircuitMaintenanceCircuitimpactBulkDeleteParamsWithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceCircuitimpactBulkDeleteParams {
	return &PluginsCircuitMaintenanceCircuitimpactBulkDeleteParams{
		HTTPClient: client,
	}
}

/* PluginsCircuitMaintenanceCircuitimpactBulkDeleteParams contains all the parameters to send to the API endpoint
   for the plugins circuit maintenance circuitimpact bulk delete operation.

   Typically these are written to a http.Request.
*/
type PluginsCircuitMaintenanceCircuitimpactBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins circuit maintenance circuitimpact bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceCircuitimpactBulkDeleteParams) WithDefaults() *PluginsCircuitMaintenanceCircuitimpactBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins circuit maintenance circuitimpact bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceCircuitimpactBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins circuit maintenance circuitimpact bulk delete params
func (o *PluginsCircuitMaintenanceCircuitimpactBulkDeleteParams) WithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceCircuitimpactBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins circuit maintenance circuitimpact bulk delete params
func (o *PluginsCircuitMaintenanceCircuitimpactBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins circuit maintenance circuitimpact bulk delete params
func (o *PluginsCircuitMaintenanceCircuitimpactBulkDeleteParams) WithContext(ctx context.Context) *PluginsCircuitMaintenanceCircuitimpactBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins circuit maintenance circuitimpact bulk delete params
func (o *PluginsCircuitMaintenanceCircuitimpactBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins circuit maintenance circuitimpact bulk delete params
func (o *PluginsCircuitMaintenanceCircuitimpactBulkDeleteParams) WithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceCircuitimpactBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins circuit maintenance circuitimpact bulk delete params
func (o *PluginsCircuitMaintenanceCircuitimpactBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsCircuitMaintenanceCircuitimpactBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

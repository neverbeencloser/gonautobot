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

// NewPluginsCircuitMaintenanceMaintenanceDeleteParams creates a new PluginsCircuitMaintenanceMaintenanceDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsCircuitMaintenanceMaintenanceDeleteParams() *PluginsCircuitMaintenanceMaintenanceDeleteParams {
	return &PluginsCircuitMaintenanceMaintenanceDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsCircuitMaintenanceMaintenanceDeleteParamsWithTimeout creates a new PluginsCircuitMaintenanceMaintenanceDeleteParams object
// with the ability to set a timeout on a request.
func NewPluginsCircuitMaintenanceMaintenanceDeleteParamsWithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceMaintenanceDeleteParams {
	return &PluginsCircuitMaintenanceMaintenanceDeleteParams{
		timeout: timeout,
	}
}

// NewPluginsCircuitMaintenanceMaintenanceDeleteParamsWithContext creates a new PluginsCircuitMaintenanceMaintenanceDeleteParams object
// with the ability to set a context for a request.
func NewPluginsCircuitMaintenanceMaintenanceDeleteParamsWithContext(ctx context.Context) *PluginsCircuitMaintenanceMaintenanceDeleteParams {
	return &PluginsCircuitMaintenanceMaintenanceDeleteParams{
		Context: ctx,
	}
}

// NewPluginsCircuitMaintenanceMaintenanceDeleteParamsWithHTTPClient creates a new PluginsCircuitMaintenanceMaintenanceDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsCircuitMaintenanceMaintenanceDeleteParamsWithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceMaintenanceDeleteParams {
	return &PluginsCircuitMaintenanceMaintenanceDeleteParams{
		HTTPClient: client,
	}
}

/* PluginsCircuitMaintenanceMaintenanceDeleteParams contains all the parameters to send to the API endpoint
   for the plugins circuit maintenance maintenance delete operation.

   Typically these are written to a http.Request.
*/
type PluginsCircuitMaintenanceMaintenanceDeleteParams struct {

	/* ID.

	   A UUID string identifying this circuit maintenance.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins circuit maintenance maintenance delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceMaintenanceDeleteParams) WithDefaults() *PluginsCircuitMaintenanceMaintenanceDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins circuit maintenance maintenance delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceMaintenanceDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins circuit maintenance maintenance delete params
func (o *PluginsCircuitMaintenanceMaintenanceDeleteParams) WithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceMaintenanceDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins circuit maintenance maintenance delete params
func (o *PluginsCircuitMaintenanceMaintenanceDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins circuit maintenance maintenance delete params
func (o *PluginsCircuitMaintenanceMaintenanceDeleteParams) WithContext(ctx context.Context) *PluginsCircuitMaintenanceMaintenanceDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins circuit maintenance maintenance delete params
func (o *PluginsCircuitMaintenanceMaintenanceDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins circuit maintenance maintenance delete params
func (o *PluginsCircuitMaintenanceMaintenanceDeleteParams) WithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceMaintenanceDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins circuit maintenance maintenance delete params
func (o *PluginsCircuitMaintenanceMaintenanceDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the plugins circuit maintenance maintenance delete params
func (o *PluginsCircuitMaintenanceMaintenanceDeleteParams) WithID(id strfmt.UUID) *PluginsCircuitMaintenanceMaintenanceDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins circuit maintenance maintenance delete params
func (o *PluginsCircuitMaintenanceMaintenanceDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsCircuitMaintenanceMaintenanceDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

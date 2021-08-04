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

// NewPluginsCircuitMaintenanceMaintenanceReadParams creates a new PluginsCircuitMaintenanceMaintenanceReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsCircuitMaintenanceMaintenanceReadParams() *PluginsCircuitMaintenanceMaintenanceReadParams {
	return &PluginsCircuitMaintenanceMaintenanceReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsCircuitMaintenanceMaintenanceReadParamsWithTimeout creates a new PluginsCircuitMaintenanceMaintenanceReadParams object
// with the ability to set a timeout on a request.
func NewPluginsCircuitMaintenanceMaintenanceReadParamsWithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceMaintenanceReadParams {
	return &PluginsCircuitMaintenanceMaintenanceReadParams{
		timeout: timeout,
	}
}

// NewPluginsCircuitMaintenanceMaintenanceReadParamsWithContext creates a new PluginsCircuitMaintenanceMaintenanceReadParams object
// with the ability to set a context for a request.
func NewPluginsCircuitMaintenanceMaintenanceReadParamsWithContext(ctx context.Context) *PluginsCircuitMaintenanceMaintenanceReadParams {
	return &PluginsCircuitMaintenanceMaintenanceReadParams{
		Context: ctx,
	}
}

// NewPluginsCircuitMaintenanceMaintenanceReadParamsWithHTTPClient creates a new PluginsCircuitMaintenanceMaintenanceReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsCircuitMaintenanceMaintenanceReadParamsWithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceMaintenanceReadParams {
	return &PluginsCircuitMaintenanceMaintenanceReadParams{
		HTTPClient: client,
	}
}

/* PluginsCircuitMaintenanceMaintenanceReadParams contains all the parameters to send to the API endpoint
   for the plugins circuit maintenance maintenance read operation.

   Typically these are written to a http.Request.
*/
type PluginsCircuitMaintenanceMaintenanceReadParams struct {

	/* ID.

	   A UUID string identifying this circuit maintenance.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins circuit maintenance maintenance read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceMaintenanceReadParams) WithDefaults() *PluginsCircuitMaintenanceMaintenanceReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins circuit maintenance maintenance read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceMaintenanceReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins circuit maintenance maintenance read params
func (o *PluginsCircuitMaintenanceMaintenanceReadParams) WithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceMaintenanceReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins circuit maintenance maintenance read params
func (o *PluginsCircuitMaintenanceMaintenanceReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins circuit maintenance maintenance read params
func (o *PluginsCircuitMaintenanceMaintenanceReadParams) WithContext(ctx context.Context) *PluginsCircuitMaintenanceMaintenanceReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins circuit maintenance maintenance read params
func (o *PluginsCircuitMaintenanceMaintenanceReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins circuit maintenance maintenance read params
func (o *PluginsCircuitMaintenanceMaintenanceReadParams) WithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceMaintenanceReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins circuit maintenance maintenance read params
func (o *PluginsCircuitMaintenanceMaintenanceReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the plugins circuit maintenance maintenance read params
func (o *PluginsCircuitMaintenanceMaintenanceReadParams) WithID(id strfmt.UUID) *PluginsCircuitMaintenanceMaintenanceReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins circuit maintenance maintenance read params
func (o *PluginsCircuitMaintenanceMaintenanceReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsCircuitMaintenanceMaintenanceReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

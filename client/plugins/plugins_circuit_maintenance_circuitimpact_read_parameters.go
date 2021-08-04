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

// NewPluginsCircuitMaintenanceCircuitimpactReadParams creates a new PluginsCircuitMaintenanceCircuitimpactReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsCircuitMaintenanceCircuitimpactReadParams() *PluginsCircuitMaintenanceCircuitimpactReadParams {
	return &PluginsCircuitMaintenanceCircuitimpactReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsCircuitMaintenanceCircuitimpactReadParamsWithTimeout creates a new PluginsCircuitMaintenanceCircuitimpactReadParams object
// with the ability to set a timeout on a request.
func NewPluginsCircuitMaintenanceCircuitimpactReadParamsWithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceCircuitimpactReadParams {
	return &PluginsCircuitMaintenanceCircuitimpactReadParams{
		timeout: timeout,
	}
}

// NewPluginsCircuitMaintenanceCircuitimpactReadParamsWithContext creates a new PluginsCircuitMaintenanceCircuitimpactReadParams object
// with the ability to set a context for a request.
func NewPluginsCircuitMaintenanceCircuitimpactReadParamsWithContext(ctx context.Context) *PluginsCircuitMaintenanceCircuitimpactReadParams {
	return &PluginsCircuitMaintenanceCircuitimpactReadParams{
		Context: ctx,
	}
}

// NewPluginsCircuitMaintenanceCircuitimpactReadParamsWithHTTPClient creates a new PluginsCircuitMaintenanceCircuitimpactReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsCircuitMaintenanceCircuitimpactReadParamsWithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceCircuitimpactReadParams {
	return &PluginsCircuitMaintenanceCircuitimpactReadParams{
		HTTPClient: client,
	}
}

/* PluginsCircuitMaintenanceCircuitimpactReadParams contains all the parameters to send to the API endpoint
   for the plugins circuit maintenance circuitimpact read operation.

   Typically these are written to a http.Request.
*/
type PluginsCircuitMaintenanceCircuitimpactReadParams struct {

	/* ID.

	   A UUID string identifying this circuit impact.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins circuit maintenance circuitimpact read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceCircuitimpactReadParams) WithDefaults() *PluginsCircuitMaintenanceCircuitimpactReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins circuit maintenance circuitimpact read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceCircuitimpactReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins circuit maintenance circuitimpact read params
func (o *PluginsCircuitMaintenanceCircuitimpactReadParams) WithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceCircuitimpactReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins circuit maintenance circuitimpact read params
func (o *PluginsCircuitMaintenanceCircuitimpactReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins circuit maintenance circuitimpact read params
func (o *PluginsCircuitMaintenanceCircuitimpactReadParams) WithContext(ctx context.Context) *PluginsCircuitMaintenanceCircuitimpactReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins circuit maintenance circuitimpact read params
func (o *PluginsCircuitMaintenanceCircuitimpactReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins circuit maintenance circuitimpact read params
func (o *PluginsCircuitMaintenanceCircuitimpactReadParams) WithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceCircuitimpactReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins circuit maintenance circuitimpact read params
func (o *PluginsCircuitMaintenanceCircuitimpactReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the plugins circuit maintenance circuitimpact read params
func (o *PluginsCircuitMaintenanceCircuitimpactReadParams) WithID(id strfmt.UUID) *PluginsCircuitMaintenanceCircuitimpactReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins circuit maintenance circuitimpact read params
func (o *PluginsCircuitMaintenanceCircuitimpactReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsCircuitMaintenanceCircuitimpactReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

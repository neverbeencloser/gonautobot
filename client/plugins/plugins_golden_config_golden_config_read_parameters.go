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

// NewPluginsGoldenConfigGoldenConfigReadParams creates a new PluginsGoldenConfigGoldenConfigReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigGoldenConfigReadParams() *PluginsGoldenConfigGoldenConfigReadParams {
	return &PluginsGoldenConfigGoldenConfigReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigGoldenConfigReadParamsWithTimeout creates a new PluginsGoldenConfigGoldenConfigReadParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigGoldenConfigReadParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigGoldenConfigReadParams {
	return &PluginsGoldenConfigGoldenConfigReadParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigGoldenConfigReadParamsWithContext creates a new PluginsGoldenConfigGoldenConfigReadParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigGoldenConfigReadParamsWithContext(ctx context.Context) *PluginsGoldenConfigGoldenConfigReadParams {
	return &PluginsGoldenConfigGoldenConfigReadParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigGoldenConfigReadParamsWithHTTPClient creates a new PluginsGoldenConfigGoldenConfigReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigGoldenConfigReadParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigGoldenConfigReadParams {
	return &PluginsGoldenConfigGoldenConfigReadParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigGoldenConfigReadParams contains all the parameters to send to the API endpoint
   for the plugins golden config golden config read operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigGoldenConfigReadParams struct {

	/* ID.

	   A UUID string identifying this golden config.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config golden config read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigGoldenConfigReadParams) WithDefaults() *PluginsGoldenConfigGoldenConfigReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config golden config read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigGoldenConfigReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config golden config read params
func (o *PluginsGoldenConfigGoldenConfigReadParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigGoldenConfigReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config golden config read params
func (o *PluginsGoldenConfigGoldenConfigReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config golden config read params
func (o *PluginsGoldenConfigGoldenConfigReadParams) WithContext(ctx context.Context) *PluginsGoldenConfigGoldenConfigReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config golden config read params
func (o *PluginsGoldenConfigGoldenConfigReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config golden config read params
func (o *PluginsGoldenConfigGoldenConfigReadParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigGoldenConfigReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config golden config read params
func (o *PluginsGoldenConfigGoldenConfigReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the plugins golden config golden config read params
func (o *PluginsGoldenConfigGoldenConfigReadParams) WithID(id strfmt.UUID) *PluginsGoldenConfigGoldenConfigReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins golden config golden config read params
func (o *PluginsGoldenConfigGoldenConfigReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigGoldenConfigReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

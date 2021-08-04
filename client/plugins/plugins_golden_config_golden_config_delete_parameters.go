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

// NewPluginsGoldenConfigGoldenConfigDeleteParams creates a new PluginsGoldenConfigGoldenConfigDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigGoldenConfigDeleteParams() *PluginsGoldenConfigGoldenConfigDeleteParams {
	return &PluginsGoldenConfigGoldenConfigDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigGoldenConfigDeleteParamsWithTimeout creates a new PluginsGoldenConfigGoldenConfigDeleteParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigGoldenConfigDeleteParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigGoldenConfigDeleteParams {
	return &PluginsGoldenConfigGoldenConfigDeleteParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigGoldenConfigDeleteParamsWithContext creates a new PluginsGoldenConfigGoldenConfigDeleteParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigGoldenConfigDeleteParamsWithContext(ctx context.Context) *PluginsGoldenConfigGoldenConfigDeleteParams {
	return &PluginsGoldenConfigGoldenConfigDeleteParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigGoldenConfigDeleteParamsWithHTTPClient creates a new PluginsGoldenConfigGoldenConfigDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigGoldenConfigDeleteParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigGoldenConfigDeleteParams {
	return &PluginsGoldenConfigGoldenConfigDeleteParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigGoldenConfigDeleteParams contains all the parameters to send to the API endpoint
   for the plugins golden config golden config delete operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigGoldenConfigDeleteParams struct {

	/* ID.

	   A UUID string identifying this golden config.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config golden config delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigGoldenConfigDeleteParams) WithDefaults() *PluginsGoldenConfigGoldenConfigDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config golden config delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigGoldenConfigDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config golden config delete params
func (o *PluginsGoldenConfigGoldenConfigDeleteParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigGoldenConfigDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config golden config delete params
func (o *PluginsGoldenConfigGoldenConfigDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config golden config delete params
func (o *PluginsGoldenConfigGoldenConfigDeleteParams) WithContext(ctx context.Context) *PluginsGoldenConfigGoldenConfigDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config golden config delete params
func (o *PluginsGoldenConfigGoldenConfigDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config golden config delete params
func (o *PluginsGoldenConfigGoldenConfigDeleteParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigGoldenConfigDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config golden config delete params
func (o *PluginsGoldenConfigGoldenConfigDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the plugins golden config golden config delete params
func (o *PluginsGoldenConfigGoldenConfigDeleteParams) WithID(id strfmt.UUID) *PluginsGoldenConfigGoldenConfigDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins golden config golden config delete params
func (o *PluginsGoldenConfigGoldenConfigDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigGoldenConfigDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

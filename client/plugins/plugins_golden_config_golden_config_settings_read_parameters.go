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

// NewPluginsGoldenConfigGoldenConfigSettingsReadParams creates a new PluginsGoldenConfigGoldenConfigSettingsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigGoldenConfigSettingsReadParams() *PluginsGoldenConfigGoldenConfigSettingsReadParams {
	return &PluginsGoldenConfigGoldenConfigSettingsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigGoldenConfigSettingsReadParamsWithTimeout creates a new PluginsGoldenConfigGoldenConfigSettingsReadParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigGoldenConfigSettingsReadParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigGoldenConfigSettingsReadParams {
	return &PluginsGoldenConfigGoldenConfigSettingsReadParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigGoldenConfigSettingsReadParamsWithContext creates a new PluginsGoldenConfigGoldenConfigSettingsReadParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigGoldenConfigSettingsReadParamsWithContext(ctx context.Context) *PluginsGoldenConfigGoldenConfigSettingsReadParams {
	return &PluginsGoldenConfigGoldenConfigSettingsReadParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigGoldenConfigSettingsReadParamsWithHTTPClient creates a new PluginsGoldenConfigGoldenConfigSettingsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigGoldenConfigSettingsReadParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigGoldenConfigSettingsReadParams {
	return &PluginsGoldenConfigGoldenConfigSettingsReadParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigGoldenConfigSettingsReadParams contains all the parameters to send to the API endpoint
   for the plugins golden config golden config settings read operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigGoldenConfigSettingsReadParams struct {

	/* ID.

	   A UUID string identifying this golden config setting.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config golden config settings read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigGoldenConfigSettingsReadParams) WithDefaults() *PluginsGoldenConfigGoldenConfigSettingsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config golden config settings read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigGoldenConfigSettingsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config golden config settings read params
func (o *PluginsGoldenConfigGoldenConfigSettingsReadParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigGoldenConfigSettingsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config golden config settings read params
func (o *PluginsGoldenConfigGoldenConfigSettingsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config golden config settings read params
func (o *PluginsGoldenConfigGoldenConfigSettingsReadParams) WithContext(ctx context.Context) *PluginsGoldenConfigGoldenConfigSettingsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config golden config settings read params
func (o *PluginsGoldenConfigGoldenConfigSettingsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config golden config settings read params
func (o *PluginsGoldenConfigGoldenConfigSettingsReadParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigGoldenConfigSettingsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config golden config settings read params
func (o *PluginsGoldenConfigGoldenConfigSettingsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the plugins golden config golden config settings read params
func (o *PluginsGoldenConfigGoldenConfigSettingsReadParams) WithID(id strfmt.UUID) *PluginsGoldenConfigGoldenConfigSettingsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins golden config golden config settings read params
func (o *PluginsGoldenConfigGoldenConfigSettingsReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigGoldenConfigSettingsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

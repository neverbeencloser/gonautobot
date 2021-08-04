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

// NewPluginsGoldenConfigGoldenConfigSettingsDeleteParams creates a new PluginsGoldenConfigGoldenConfigSettingsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigGoldenConfigSettingsDeleteParams() *PluginsGoldenConfigGoldenConfigSettingsDeleteParams {
	return &PluginsGoldenConfigGoldenConfigSettingsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigGoldenConfigSettingsDeleteParamsWithTimeout creates a new PluginsGoldenConfigGoldenConfigSettingsDeleteParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigGoldenConfigSettingsDeleteParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigGoldenConfigSettingsDeleteParams {
	return &PluginsGoldenConfigGoldenConfigSettingsDeleteParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigGoldenConfigSettingsDeleteParamsWithContext creates a new PluginsGoldenConfigGoldenConfigSettingsDeleteParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigGoldenConfigSettingsDeleteParamsWithContext(ctx context.Context) *PluginsGoldenConfigGoldenConfigSettingsDeleteParams {
	return &PluginsGoldenConfigGoldenConfigSettingsDeleteParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigGoldenConfigSettingsDeleteParamsWithHTTPClient creates a new PluginsGoldenConfigGoldenConfigSettingsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigGoldenConfigSettingsDeleteParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigGoldenConfigSettingsDeleteParams {
	return &PluginsGoldenConfigGoldenConfigSettingsDeleteParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigGoldenConfigSettingsDeleteParams contains all the parameters to send to the API endpoint
   for the plugins golden config golden config settings delete operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigGoldenConfigSettingsDeleteParams struct {

	/* ID.

	   A UUID string identifying this golden config setting.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config golden config settings delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigGoldenConfigSettingsDeleteParams) WithDefaults() *PluginsGoldenConfigGoldenConfigSettingsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config golden config settings delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigGoldenConfigSettingsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config golden config settings delete params
func (o *PluginsGoldenConfigGoldenConfigSettingsDeleteParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigGoldenConfigSettingsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config golden config settings delete params
func (o *PluginsGoldenConfigGoldenConfigSettingsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config golden config settings delete params
func (o *PluginsGoldenConfigGoldenConfigSettingsDeleteParams) WithContext(ctx context.Context) *PluginsGoldenConfigGoldenConfigSettingsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config golden config settings delete params
func (o *PluginsGoldenConfigGoldenConfigSettingsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config golden config settings delete params
func (o *PluginsGoldenConfigGoldenConfigSettingsDeleteParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigGoldenConfigSettingsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config golden config settings delete params
func (o *PluginsGoldenConfigGoldenConfigSettingsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the plugins golden config golden config settings delete params
func (o *PluginsGoldenConfigGoldenConfigSettingsDeleteParams) WithID(id strfmt.UUID) *PluginsGoldenConfigGoldenConfigSettingsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins golden config golden config settings delete params
func (o *PluginsGoldenConfigGoldenConfigSettingsDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigGoldenConfigSettingsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

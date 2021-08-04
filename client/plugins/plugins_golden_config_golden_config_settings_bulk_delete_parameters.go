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

// NewPluginsGoldenConfigGoldenConfigSettingsBulkDeleteParams creates a new PluginsGoldenConfigGoldenConfigSettingsBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigGoldenConfigSettingsBulkDeleteParams() *PluginsGoldenConfigGoldenConfigSettingsBulkDeleteParams {
	return &PluginsGoldenConfigGoldenConfigSettingsBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigGoldenConfigSettingsBulkDeleteParamsWithTimeout creates a new PluginsGoldenConfigGoldenConfigSettingsBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigGoldenConfigSettingsBulkDeleteParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigGoldenConfigSettingsBulkDeleteParams {
	return &PluginsGoldenConfigGoldenConfigSettingsBulkDeleteParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigGoldenConfigSettingsBulkDeleteParamsWithContext creates a new PluginsGoldenConfigGoldenConfigSettingsBulkDeleteParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigGoldenConfigSettingsBulkDeleteParamsWithContext(ctx context.Context) *PluginsGoldenConfigGoldenConfigSettingsBulkDeleteParams {
	return &PluginsGoldenConfigGoldenConfigSettingsBulkDeleteParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigGoldenConfigSettingsBulkDeleteParamsWithHTTPClient creates a new PluginsGoldenConfigGoldenConfigSettingsBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigGoldenConfigSettingsBulkDeleteParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigGoldenConfigSettingsBulkDeleteParams {
	return &PluginsGoldenConfigGoldenConfigSettingsBulkDeleteParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigGoldenConfigSettingsBulkDeleteParams contains all the parameters to send to the API endpoint
   for the plugins golden config golden config settings bulk delete operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigGoldenConfigSettingsBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config golden config settings bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigGoldenConfigSettingsBulkDeleteParams) WithDefaults() *PluginsGoldenConfigGoldenConfigSettingsBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config golden config settings bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigGoldenConfigSettingsBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config golden config settings bulk delete params
func (o *PluginsGoldenConfigGoldenConfigSettingsBulkDeleteParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigGoldenConfigSettingsBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config golden config settings bulk delete params
func (o *PluginsGoldenConfigGoldenConfigSettingsBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config golden config settings bulk delete params
func (o *PluginsGoldenConfigGoldenConfigSettingsBulkDeleteParams) WithContext(ctx context.Context) *PluginsGoldenConfigGoldenConfigSettingsBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config golden config settings bulk delete params
func (o *PluginsGoldenConfigGoldenConfigSettingsBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config golden config settings bulk delete params
func (o *PluginsGoldenConfigGoldenConfigSettingsBulkDeleteParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigGoldenConfigSettingsBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config golden config settings bulk delete params
func (o *PluginsGoldenConfigGoldenConfigSettingsBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigGoldenConfigSettingsBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

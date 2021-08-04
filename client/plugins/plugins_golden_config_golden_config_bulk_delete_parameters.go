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

// NewPluginsGoldenConfigGoldenConfigBulkDeleteParams creates a new PluginsGoldenConfigGoldenConfigBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigGoldenConfigBulkDeleteParams() *PluginsGoldenConfigGoldenConfigBulkDeleteParams {
	return &PluginsGoldenConfigGoldenConfigBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigGoldenConfigBulkDeleteParamsWithTimeout creates a new PluginsGoldenConfigGoldenConfigBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigGoldenConfigBulkDeleteParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigGoldenConfigBulkDeleteParams {
	return &PluginsGoldenConfigGoldenConfigBulkDeleteParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigGoldenConfigBulkDeleteParamsWithContext creates a new PluginsGoldenConfigGoldenConfigBulkDeleteParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigGoldenConfigBulkDeleteParamsWithContext(ctx context.Context) *PluginsGoldenConfigGoldenConfigBulkDeleteParams {
	return &PluginsGoldenConfigGoldenConfigBulkDeleteParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigGoldenConfigBulkDeleteParamsWithHTTPClient creates a new PluginsGoldenConfigGoldenConfigBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigGoldenConfigBulkDeleteParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigGoldenConfigBulkDeleteParams {
	return &PluginsGoldenConfigGoldenConfigBulkDeleteParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigGoldenConfigBulkDeleteParams contains all the parameters to send to the API endpoint
   for the plugins golden config golden config bulk delete operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigGoldenConfigBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config golden config bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigGoldenConfigBulkDeleteParams) WithDefaults() *PluginsGoldenConfigGoldenConfigBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config golden config bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigGoldenConfigBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config golden config bulk delete params
func (o *PluginsGoldenConfigGoldenConfigBulkDeleteParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigGoldenConfigBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config golden config bulk delete params
func (o *PluginsGoldenConfigGoldenConfigBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config golden config bulk delete params
func (o *PluginsGoldenConfigGoldenConfigBulkDeleteParams) WithContext(ctx context.Context) *PluginsGoldenConfigGoldenConfigBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config golden config bulk delete params
func (o *PluginsGoldenConfigGoldenConfigBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config golden config bulk delete params
func (o *PluginsGoldenConfigGoldenConfigBulkDeleteParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigGoldenConfigBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config golden config bulk delete params
func (o *PluginsGoldenConfigGoldenConfigBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigGoldenConfigBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

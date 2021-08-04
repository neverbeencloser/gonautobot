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

// NewPluginsGoldenConfigConfigReplaceBulkDeleteParams creates a new PluginsGoldenConfigConfigReplaceBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigConfigReplaceBulkDeleteParams() *PluginsGoldenConfigConfigReplaceBulkDeleteParams {
	return &PluginsGoldenConfigConfigReplaceBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigConfigReplaceBulkDeleteParamsWithTimeout creates a new PluginsGoldenConfigConfigReplaceBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigConfigReplaceBulkDeleteParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigConfigReplaceBulkDeleteParams {
	return &PluginsGoldenConfigConfigReplaceBulkDeleteParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigConfigReplaceBulkDeleteParamsWithContext creates a new PluginsGoldenConfigConfigReplaceBulkDeleteParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigConfigReplaceBulkDeleteParamsWithContext(ctx context.Context) *PluginsGoldenConfigConfigReplaceBulkDeleteParams {
	return &PluginsGoldenConfigConfigReplaceBulkDeleteParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigConfigReplaceBulkDeleteParamsWithHTTPClient creates a new PluginsGoldenConfigConfigReplaceBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigConfigReplaceBulkDeleteParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigConfigReplaceBulkDeleteParams {
	return &PluginsGoldenConfigConfigReplaceBulkDeleteParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigConfigReplaceBulkDeleteParams contains all the parameters to send to the API endpoint
   for the plugins golden config config replace bulk delete operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigConfigReplaceBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config config replace bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigConfigReplaceBulkDeleteParams) WithDefaults() *PluginsGoldenConfigConfigReplaceBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config config replace bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigConfigReplaceBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config config replace bulk delete params
func (o *PluginsGoldenConfigConfigReplaceBulkDeleteParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigConfigReplaceBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config config replace bulk delete params
func (o *PluginsGoldenConfigConfigReplaceBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config config replace bulk delete params
func (o *PluginsGoldenConfigConfigReplaceBulkDeleteParams) WithContext(ctx context.Context) *PluginsGoldenConfigConfigReplaceBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config config replace bulk delete params
func (o *PluginsGoldenConfigConfigReplaceBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config config replace bulk delete params
func (o *PluginsGoldenConfigConfigReplaceBulkDeleteParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigConfigReplaceBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config config replace bulk delete params
func (o *PluginsGoldenConfigConfigReplaceBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigConfigReplaceBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

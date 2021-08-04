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

// NewPluginsDataValidationEngineRulesRegexBulkDeleteParams creates a new PluginsDataValidationEngineRulesRegexBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsDataValidationEngineRulesRegexBulkDeleteParams() *PluginsDataValidationEngineRulesRegexBulkDeleteParams {
	return &PluginsDataValidationEngineRulesRegexBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsDataValidationEngineRulesRegexBulkDeleteParamsWithTimeout creates a new PluginsDataValidationEngineRulesRegexBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewPluginsDataValidationEngineRulesRegexBulkDeleteParamsWithTimeout(timeout time.Duration) *PluginsDataValidationEngineRulesRegexBulkDeleteParams {
	return &PluginsDataValidationEngineRulesRegexBulkDeleteParams{
		timeout: timeout,
	}
}

// NewPluginsDataValidationEngineRulesRegexBulkDeleteParamsWithContext creates a new PluginsDataValidationEngineRulesRegexBulkDeleteParams object
// with the ability to set a context for a request.
func NewPluginsDataValidationEngineRulesRegexBulkDeleteParamsWithContext(ctx context.Context) *PluginsDataValidationEngineRulesRegexBulkDeleteParams {
	return &PluginsDataValidationEngineRulesRegexBulkDeleteParams{
		Context: ctx,
	}
}

// NewPluginsDataValidationEngineRulesRegexBulkDeleteParamsWithHTTPClient creates a new PluginsDataValidationEngineRulesRegexBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsDataValidationEngineRulesRegexBulkDeleteParamsWithHTTPClient(client *http.Client) *PluginsDataValidationEngineRulesRegexBulkDeleteParams {
	return &PluginsDataValidationEngineRulesRegexBulkDeleteParams{
		HTTPClient: client,
	}
}

/* PluginsDataValidationEngineRulesRegexBulkDeleteParams contains all the parameters to send to the API endpoint
   for the plugins data validation engine rules regex bulk delete operation.

   Typically these are written to a http.Request.
*/
type PluginsDataValidationEngineRulesRegexBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins data validation engine rules regex bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsDataValidationEngineRulesRegexBulkDeleteParams) WithDefaults() *PluginsDataValidationEngineRulesRegexBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins data validation engine rules regex bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsDataValidationEngineRulesRegexBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins data validation engine rules regex bulk delete params
func (o *PluginsDataValidationEngineRulesRegexBulkDeleteParams) WithTimeout(timeout time.Duration) *PluginsDataValidationEngineRulesRegexBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins data validation engine rules regex bulk delete params
func (o *PluginsDataValidationEngineRulesRegexBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins data validation engine rules regex bulk delete params
func (o *PluginsDataValidationEngineRulesRegexBulkDeleteParams) WithContext(ctx context.Context) *PluginsDataValidationEngineRulesRegexBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins data validation engine rules regex bulk delete params
func (o *PluginsDataValidationEngineRulesRegexBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins data validation engine rules regex bulk delete params
func (o *PluginsDataValidationEngineRulesRegexBulkDeleteParams) WithHTTPClient(client *http.Client) *PluginsDataValidationEngineRulesRegexBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins data validation engine rules regex bulk delete params
func (o *PluginsDataValidationEngineRulesRegexBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsDataValidationEngineRulesRegexBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

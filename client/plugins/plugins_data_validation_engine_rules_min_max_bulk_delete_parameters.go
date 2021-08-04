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

// NewPluginsDataValidationEngineRulesMinMaxBulkDeleteParams creates a new PluginsDataValidationEngineRulesMinMaxBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsDataValidationEngineRulesMinMaxBulkDeleteParams() *PluginsDataValidationEngineRulesMinMaxBulkDeleteParams {
	return &PluginsDataValidationEngineRulesMinMaxBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsDataValidationEngineRulesMinMaxBulkDeleteParamsWithTimeout creates a new PluginsDataValidationEngineRulesMinMaxBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewPluginsDataValidationEngineRulesMinMaxBulkDeleteParamsWithTimeout(timeout time.Duration) *PluginsDataValidationEngineRulesMinMaxBulkDeleteParams {
	return &PluginsDataValidationEngineRulesMinMaxBulkDeleteParams{
		timeout: timeout,
	}
}

// NewPluginsDataValidationEngineRulesMinMaxBulkDeleteParamsWithContext creates a new PluginsDataValidationEngineRulesMinMaxBulkDeleteParams object
// with the ability to set a context for a request.
func NewPluginsDataValidationEngineRulesMinMaxBulkDeleteParamsWithContext(ctx context.Context) *PluginsDataValidationEngineRulesMinMaxBulkDeleteParams {
	return &PluginsDataValidationEngineRulesMinMaxBulkDeleteParams{
		Context: ctx,
	}
}

// NewPluginsDataValidationEngineRulesMinMaxBulkDeleteParamsWithHTTPClient creates a new PluginsDataValidationEngineRulesMinMaxBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsDataValidationEngineRulesMinMaxBulkDeleteParamsWithHTTPClient(client *http.Client) *PluginsDataValidationEngineRulesMinMaxBulkDeleteParams {
	return &PluginsDataValidationEngineRulesMinMaxBulkDeleteParams{
		HTTPClient: client,
	}
}

/* PluginsDataValidationEngineRulesMinMaxBulkDeleteParams contains all the parameters to send to the API endpoint
   for the plugins data validation engine rules min max bulk delete operation.

   Typically these are written to a http.Request.
*/
type PluginsDataValidationEngineRulesMinMaxBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins data validation engine rules min max bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsDataValidationEngineRulesMinMaxBulkDeleteParams) WithDefaults() *PluginsDataValidationEngineRulesMinMaxBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins data validation engine rules min max bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsDataValidationEngineRulesMinMaxBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins data validation engine rules min max bulk delete params
func (o *PluginsDataValidationEngineRulesMinMaxBulkDeleteParams) WithTimeout(timeout time.Duration) *PluginsDataValidationEngineRulesMinMaxBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins data validation engine rules min max bulk delete params
func (o *PluginsDataValidationEngineRulesMinMaxBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins data validation engine rules min max bulk delete params
func (o *PluginsDataValidationEngineRulesMinMaxBulkDeleteParams) WithContext(ctx context.Context) *PluginsDataValidationEngineRulesMinMaxBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins data validation engine rules min max bulk delete params
func (o *PluginsDataValidationEngineRulesMinMaxBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins data validation engine rules min max bulk delete params
func (o *PluginsDataValidationEngineRulesMinMaxBulkDeleteParams) WithHTTPClient(client *http.Client) *PluginsDataValidationEngineRulesMinMaxBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins data validation engine rules min max bulk delete params
func (o *PluginsDataValidationEngineRulesMinMaxBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsDataValidationEngineRulesMinMaxBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

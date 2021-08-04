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

// NewPluginsGoldenConfigComplianceRuleBulkDeleteParams creates a new PluginsGoldenConfigComplianceRuleBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigComplianceRuleBulkDeleteParams() *PluginsGoldenConfigComplianceRuleBulkDeleteParams {
	return &PluginsGoldenConfigComplianceRuleBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigComplianceRuleBulkDeleteParamsWithTimeout creates a new PluginsGoldenConfigComplianceRuleBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigComplianceRuleBulkDeleteParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigComplianceRuleBulkDeleteParams {
	return &PluginsGoldenConfigComplianceRuleBulkDeleteParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigComplianceRuleBulkDeleteParamsWithContext creates a new PluginsGoldenConfigComplianceRuleBulkDeleteParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigComplianceRuleBulkDeleteParamsWithContext(ctx context.Context) *PluginsGoldenConfigComplianceRuleBulkDeleteParams {
	return &PluginsGoldenConfigComplianceRuleBulkDeleteParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigComplianceRuleBulkDeleteParamsWithHTTPClient creates a new PluginsGoldenConfigComplianceRuleBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigComplianceRuleBulkDeleteParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigComplianceRuleBulkDeleteParams {
	return &PluginsGoldenConfigComplianceRuleBulkDeleteParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigComplianceRuleBulkDeleteParams contains all the parameters to send to the API endpoint
   for the plugins golden config compliance rule bulk delete operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigComplianceRuleBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config compliance rule bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigComplianceRuleBulkDeleteParams) WithDefaults() *PluginsGoldenConfigComplianceRuleBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config compliance rule bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigComplianceRuleBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config compliance rule bulk delete params
func (o *PluginsGoldenConfigComplianceRuleBulkDeleteParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigComplianceRuleBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config compliance rule bulk delete params
func (o *PluginsGoldenConfigComplianceRuleBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config compliance rule bulk delete params
func (o *PluginsGoldenConfigComplianceRuleBulkDeleteParams) WithContext(ctx context.Context) *PluginsGoldenConfigComplianceRuleBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config compliance rule bulk delete params
func (o *PluginsGoldenConfigComplianceRuleBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config compliance rule bulk delete params
func (o *PluginsGoldenConfigComplianceRuleBulkDeleteParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigComplianceRuleBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config compliance rule bulk delete params
func (o *PluginsGoldenConfigComplianceRuleBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigComplianceRuleBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

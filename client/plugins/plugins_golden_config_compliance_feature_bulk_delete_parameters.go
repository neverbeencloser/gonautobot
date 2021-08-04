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

// NewPluginsGoldenConfigComplianceFeatureBulkDeleteParams creates a new PluginsGoldenConfigComplianceFeatureBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigComplianceFeatureBulkDeleteParams() *PluginsGoldenConfigComplianceFeatureBulkDeleteParams {
	return &PluginsGoldenConfigComplianceFeatureBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigComplianceFeatureBulkDeleteParamsWithTimeout creates a new PluginsGoldenConfigComplianceFeatureBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigComplianceFeatureBulkDeleteParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigComplianceFeatureBulkDeleteParams {
	return &PluginsGoldenConfigComplianceFeatureBulkDeleteParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigComplianceFeatureBulkDeleteParamsWithContext creates a new PluginsGoldenConfigComplianceFeatureBulkDeleteParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigComplianceFeatureBulkDeleteParamsWithContext(ctx context.Context) *PluginsGoldenConfigComplianceFeatureBulkDeleteParams {
	return &PluginsGoldenConfigComplianceFeatureBulkDeleteParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigComplianceFeatureBulkDeleteParamsWithHTTPClient creates a new PluginsGoldenConfigComplianceFeatureBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigComplianceFeatureBulkDeleteParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigComplianceFeatureBulkDeleteParams {
	return &PluginsGoldenConfigComplianceFeatureBulkDeleteParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigComplianceFeatureBulkDeleteParams contains all the parameters to send to the API endpoint
   for the plugins golden config compliance feature bulk delete operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigComplianceFeatureBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config compliance feature bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigComplianceFeatureBulkDeleteParams) WithDefaults() *PluginsGoldenConfigComplianceFeatureBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config compliance feature bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigComplianceFeatureBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config compliance feature bulk delete params
func (o *PluginsGoldenConfigComplianceFeatureBulkDeleteParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigComplianceFeatureBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config compliance feature bulk delete params
func (o *PluginsGoldenConfigComplianceFeatureBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config compliance feature bulk delete params
func (o *PluginsGoldenConfigComplianceFeatureBulkDeleteParams) WithContext(ctx context.Context) *PluginsGoldenConfigComplianceFeatureBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config compliance feature bulk delete params
func (o *PluginsGoldenConfigComplianceFeatureBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config compliance feature bulk delete params
func (o *PluginsGoldenConfigComplianceFeatureBulkDeleteParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigComplianceFeatureBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config compliance feature bulk delete params
func (o *PluginsGoldenConfigComplianceFeatureBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigComplianceFeatureBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

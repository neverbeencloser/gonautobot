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

// NewPluginsGoldenConfigComplianceFeatureDeleteParams creates a new PluginsGoldenConfigComplianceFeatureDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigComplianceFeatureDeleteParams() *PluginsGoldenConfigComplianceFeatureDeleteParams {
	return &PluginsGoldenConfigComplianceFeatureDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigComplianceFeatureDeleteParamsWithTimeout creates a new PluginsGoldenConfigComplianceFeatureDeleteParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigComplianceFeatureDeleteParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigComplianceFeatureDeleteParams {
	return &PluginsGoldenConfigComplianceFeatureDeleteParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigComplianceFeatureDeleteParamsWithContext creates a new PluginsGoldenConfigComplianceFeatureDeleteParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigComplianceFeatureDeleteParamsWithContext(ctx context.Context) *PluginsGoldenConfigComplianceFeatureDeleteParams {
	return &PluginsGoldenConfigComplianceFeatureDeleteParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigComplianceFeatureDeleteParamsWithHTTPClient creates a new PluginsGoldenConfigComplianceFeatureDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigComplianceFeatureDeleteParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigComplianceFeatureDeleteParams {
	return &PluginsGoldenConfigComplianceFeatureDeleteParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigComplianceFeatureDeleteParams contains all the parameters to send to the API endpoint
   for the plugins golden config compliance feature delete operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigComplianceFeatureDeleteParams struct {

	/* ID.

	   A UUID string identifying this compliance feature.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config compliance feature delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigComplianceFeatureDeleteParams) WithDefaults() *PluginsGoldenConfigComplianceFeatureDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config compliance feature delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigComplianceFeatureDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config compliance feature delete params
func (o *PluginsGoldenConfigComplianceFeatureDeleteParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigComplianceFeatureDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config compliance feature delete params
func (o *PluginsGoldenConfigComplianceFeatureDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config compliance feature delete params
func (o *PluginsGoldenConfigComplianceFeatureDeleteParams) WithContext(ctx context.Context) *PluginsGoldenConfigComplianceFeatureDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config compliance feature delete params
func (o *PluginsGoldenConfigComplianceFeatureDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config compliance feature delete params
func (o *PluginsGoldenConfigComplianceFeatureDeleteParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigComplianceFeatureDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config compliance feature delete params
func (o *PluginsGoldenConfigComplianceFeatureDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the plugins golden config compliance feature delete params
func (o *PluginsGoldenConfigComplianceFeatureDeleteParams) WithID(id strfmt.UUID) *PluginsGoldenConfigComplianceFeatureDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins golden config compliance feature delete params
func (o *PluginsGoldenConfigComplianceFeatureDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigComplianceFeatureDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

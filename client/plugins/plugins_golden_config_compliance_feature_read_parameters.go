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

// NewPluginsGoldenConfigComplianceFeatureReadParams creates a new PluginsGoldenConfigComplianceFeatureReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigComplianceFeatureReadParams() *PluginsGoldenConfigComplianceFeatureReadParams {
	return &PluginsGoldenConfigComplianceFeatureReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigComplianceFeatureReadParamsWithTimeout creates a new PluginsGoldenConfigComplianceFeatureReadParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigComplianceFeatureReadParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigComplianceFeatureReadParams {
	return &PluginsGoldenConfigComplianceFeatureReadParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigComplianceFeatureReadParamsWithContext creates a new PluginsGoldenConfigComplianceFeatureReadParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigComplianceFeatureReadParamsWithContext(ctx context.Context) *PluginsGoldenConfigComplianceFeatureReadParams {
	return &PluginsGoldenConfigComplianceFeatureReadParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigComplianceFeatureReadParamsWithHTTPClient creates a new PluginsGoldenConfigComplianceFeatureReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigComplianceFeatureReadParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigComplianceFeatureReadParams {
	return &PluginsGoldenConfigComplianceFeatureReadParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigComplianceFeatureReadParams contains all the parameters to send to the API endpoint
   for the plugins golden config compliance feature read operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigComplianceFeatureReadParams struct {

	/* ID.

	   A UUID string identifying this compliance feature.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config compliance feature read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigComplianceFeatureReadParams) WithDefaults() *PluginsGoldenConfigComplianceFeatureReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config compliance feature read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigComplianceFeatureReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config compliance feature read params
func (o *PluginsGoldenConfigComplianceFeatureReadParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigComplianceFeatureReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config compliance feature read params
func (o *PluginsGoldenConfigComplianceFeatureReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config compliance feature read params
func (o *PluginsGoldenConfigComplianceFeatureReadParams) WithContext(ctx context.Context) *PluginsGoldenConfigComplianceFeatureReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config compliance feature read params
func (o *PluginsGoldenConfigComplianceFeatureReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config compliance feature read params
func (o *PluginsGoldenConfigComplianceFeatureReadParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigComplianceFeatureReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config compliance feature read params
func (o *PluginsGoldenConfigComplianceFeatureReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the plugins golden config compliance feature read params
func (o *PluginsGoldenConfigComplianceFeatureReadParams) WithID(id strfmt.UUID) *PluginsGoldenConfigComplianceFeatureReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins golden config compliance feature read params
func (o *PluginsGoldenConfigComplianceFeatureReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigComplianceFeatureReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

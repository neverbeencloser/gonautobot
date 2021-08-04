package plugins

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// NewPluginsGoldenConfigComplianceFeatureCreateParams creates a new PluginsGoldenConfigComplianceFeatureCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigComplianceFeatureCreateParams() *PluginsGoldenConfigComplianceFeatureCreateParams {
	return &PluginsGoldenConfigComplianceFeatureCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigComplianceFeatureCreateParamsWithTimeout creates a new PluginsGoldenConfigComplianceFeatureCreateParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigComplianceFeatureCreateParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigComplianceFeatureCreateParams {
	return &PluginsGoldenConfigComplianceFeatureCreateParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigComplianceFeatureCreateParamsWithContext creates a new PluginsGoldenConfigComplianceFeatureCreateParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigComplianceFeatureCreateParamsWithContext(ctx context.Context) *PluginsGoldenConfigComplianceFeatureCreateParams {
	return &PluginsGoldenConfigComplianceFeatureCreateParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigComplianceFeatureCreateParamsWithHTTPClient creates a new PluginsGoldenConfigComplianceFeatureCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigComplianceFeatureCreateParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigComplianceFeatureCreateParams {
	return &PluginsGoldenConfigComplianceFeatureCreateParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigComplianceFeatureCreateParams contains all the parameters to send to the API endpoint
   for the plugins golden config compliance feature create operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigComplianceFeatureCreateParams struct {

	// Data.
	Data *models.ComplianceFeature

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config compliance feature create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigComplianceFeatureCreateParams) WithDefaults() *PluginsGoldenConfigComplianceFeatureCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config compliance feature create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigComplianceFeatureCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config compliance feature create params
func (o *PluginsGoldenConfigComplianceFeatureCreateParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigComplianceFeatureCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config compliance feature create params
func (o *PluginsGoldenConfigComplianceFeatureCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config compliance feature create params
func (o *PluginsGoldenConfigComplianceFeatureCreateParams) WithContext(ctx context.Context) *PluginsGoldenConfigComplianceFeatureCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config compliance feature create params
func (o *PluginsGoldenConfigComplianceFeatureCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config compliance feature create params
func (o *PluginsGoldenConfigComplianceFeatureCreateParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigComplianceFeatureCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config compliance feature create params
func (o *PluginsGoldenConfigComplianceFeatureCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins golden config compliance feature create params
func (o *PluginsGoldenConfigComplianceFeatureCreateParams) WithData(data *models.ComplianceFeature) *PluginsGoldenConfigComplianceFeatureCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins golden config compliance feature create params
func (o *PluginsGoldenConfigComplianceFeatureCreateParams) SetData(data *models.ComplianceFeature) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigComplianceFeatureCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

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

// NewPluginsGoldenConfigComplianceFeatureUpdateParams creates a new PluginsGoldenConfigComplianceFeatureUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigComplianceFeatureUpdateParams() *PluginsGoldenConfigComplianceFeatureUpdateParams {
	return &PluginsGoldenConfigComplianceFeatureUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigComplianceFeatureUpdateParamsWithTimeout creates a new PluginsGoldenConfigComplianceFeatureUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigComplianceFeatureUpdateParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigComplianceFeatureUpdateParams {
	return &PluginsGoldenConfigComplianceFeatureUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigComplianceFeatureUpdateParamsWithContext creates a new PluginsGoldenConfigComplianceFeatureUpdateParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigComplianceFeatureUpdateParamsWithContext(ctx context.Context) *PluginsGoldenConfigComplianceFeatureUpdateParams {
	return &PluginsGoldenConfigComplianceFeatureUpdateParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigComplianceFeatureUpdateParamsWithHTTPClient creates a new PluginsGoldenConfigComplianceFeatureUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigComplianceFeatureUpdateParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigComplianceFeatureUpdateParams {
	return &PluginsGoldenConfigComplianceFeatureUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigComplianceFeatureUpdateParams contains all the parameters to send to the API endpoint
   for the plugins golden config compliance feature update operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigComplianceFeatureUpdateParams struct {

	// Data.
	Data *models.ComplianceFeature

	/* ID.

	   A UUID string identifying this compliance feature.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config compliance feature update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigComplianceFeatureUpdateParams) WithDefaults() *PluginsGoldenConfigComplianceFeatureUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config compliance feature update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigComplianceFeatureUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config compliance feature update params
func (o *PluginsGoldenConfigComplianceFeatureUpdateParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigComplianceFeatureUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config compliance feature update params
func (o *PluginsGoldenConfigComplianceFeatureUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config compliance feature update params
func (o *PluginsGoldenConfigComplianceFeatureUpdateParams) WithContext(ctx context.Context) *PluginsGoldenConfigComplianceFeatureUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config compliance feature update params
func (o *PluginsGoldenConfigComplianceFeatureUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config compliance feature update params
func (o *PluginsGoldenConfigComplianceFeatureUpdateParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigComplianceFeatureUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config compliance feature update params
func (o *PluginsGoldenConfigComplianceFeatureUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins golden config compliance feature update params
func (o *PluginsGoldenConfigComplianceFeatureUpdateParams) WithData(data *models.ComplianceFeature) *PluginsGoldenConfigComplianceFeatureUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins golden config compliance feature update params
func (o *PluginsGoldenConfigComplianceFeatureUpdateParams) SetData(data *models.ComplianceFeature) {
	o.Data = data
}

// WithID adds the id to the plugins golden config compliance feature update params
func (o *PluginsGoldenConfigComplianceFeatureUpdateParams) WithID(id strfmt.UUID) *PluginsGoldenConfigComplianceFeatureUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins golden config compliance feature update params
func (o *PluginsGoldenConfigComplianceFeatureUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigComplianceFeatureUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

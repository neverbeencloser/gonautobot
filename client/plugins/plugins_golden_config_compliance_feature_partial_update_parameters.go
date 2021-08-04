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

// NewPluginsGoldenConfigComplianceFeaturePartialUpdateParams creates a new PluginsGoldenConfigComplianceFeaturePartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigComplianceFeaturePartialUpdateParams() *PluginsGoldenConfigComplianceFeaturePartialUpdateParams {
	return &PluginsGoldenConfigComplianceFeaturePartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigComplianceFeaturePartialUpdateParamsWithTimeout creates a new PluginsGoldenConfigComplianceFeaturePartialUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigComplianceFeaturePartialUpdateParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigComplianceFeaturePartialUpdateParams {
	return &PluginsGoldenConfigComplianceFeaturePartialUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigComplianceFeaturePartialUpdateParamsWithContext creates a new PluginsGoldenConfigComplianceFeaturePartialUpdateParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigComplianceFeaturePartialUpdateParamsWithContext(ctx context.Context) *PluginsGoldenConfigComplianceFeaturePartialUpdateParams {
	return &PluginsGoldenConfigComplianceFeaturePartialUpdateParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigComplianceFeaturePartialUpdateParamsWithHTTPClient creates a new PluginsGoldenConfigComplianceFeaturePartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigComplianceFeaturePartialUpdateParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigComplianceFeaturePartialUpdateParams {
	return &PluginsGoldenConfigComplianceFeaturePartialUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigComplianceFeaturePartialUpdateParams contains all the parameters to send to the API endpoint
   for the plugins golden config compliance feature partial update operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigComplianceFeaturePartialUpdateParams struct {

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

// WithDefaults hydrates default values in the plugins golden config compliance feature partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigComplianceFeaturePartialUpdateParams) WithDefaults() *PluginsGoldenConfigComplianceFeaturePartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config compliance feature partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigComplianceFeaturePartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config compliance feature partial update params
func (o *PluginsGoldenConfigComplianceFeaturePartialUpdateParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigComplianceFeaturePartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config compliance feature partial update params
func (o *PluginsGoldenConfigComplianceFeaturePartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config compliance feature partial update params
func (o *PluginsGoldenConfigComplianceFeaturePartialUpdateParams) WithContext(ctx context.Context) *PluginsGoldenConfigComplianceFeaturePartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config compliance feature partial update params
func (o *PluginsGoldenConfigComplianceFeaturePartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config compliance feature partial update params
func (o *PluginsGoldenConfigComplianceFeaturePartialUpdateParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigComplianceFeaturePartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config compliance feature partial update params
func (o *PluginsGoldenConfigComplianceFeaturePartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins golden config compliance feature partial update params
func (o *PluginsGoldenConfigComplianceFeaturePartialUpdateParams) WithData(data *models.ComplianceFeature) *PluginsGoldenConfigComplianceFeaturePartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins golden config compliance feature partial update params
func (o *PluginsGoldenConfigComplianceFeaturePartialUpdateParams) SetData(data *models.ComplianceFeature) {
	o.Data = data
}

// WithID adds the id to the plugins golden config compliance feature partial update params
func (o *PluginsGoldenConfigComplianceFeaturePartialUpdateParams) WithID(id strfmt.UUID) *PluginsGoldenConfigComplianceFeaturePartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins golden config compliance feature partial update params
func (o *PluginsGoldenConfigComplianceFeaturePartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigComplianceFeaturePartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

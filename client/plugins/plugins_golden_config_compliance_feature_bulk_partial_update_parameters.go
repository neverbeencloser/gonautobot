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

// NewPluginsGoldenConfigComplianceFeatureBulkPartialUpdateParams creates a new PluginsGoldenConfigComplianceFeatureBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigComplianceFeatureBulkPartialUpdateParams() *PluginsGoldenConfigComplianceFeatureBulkPartialUpdateParams {
	return &PluginsGoldenConfigComplianceFeatureBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigComplianceFeatureBulkPartialUpdateParamsWithTimeout creates a new PluginsGoldenConfigComplianceFeatureBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigComplianceFeatureBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigComplianceFeatureBulkPartialUpdateParams {
	return &PluginsGoldenConfigComplianceFeatureBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigComplianceFeatureBulkPartialUpdateParamsWithContext creates a new PluginsGoldenConfigComplianceFeatureBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigComplianceFeatureBulkPartialUpdateParamsWithContext(ctx context.Context) *PluginsGoldenConfigComplianceFeatureBulkPartialUpdateParams {
	return &PluginsGoldenConfigComplianceFeatureBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigComplianceFeatureBulkPartialUpdateParamsWithHTTPClient creates a new PluginsGoldenConfigComplianceFeatureBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigComplianceFeatureBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigComplianceFeatureBulkPartialUpdateParams {
	return &PluginsGoldenConfigComplianceFeatureBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigComplianceFeatureBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the plugins golden config compliance feature bulk partial update operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigComplianceFeatureBulkPartialUpdateParams struct {

	// Data.
	Data *models.ComplianceFeature

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config compliance feature bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigComplianceFeatureBulkPartialUpdateParams) WithDefaults() *PluginsGoldenConfigComplianceFeatureBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config compliance feature bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigComplianceFeatureBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config compliance feature bulk partial update params
func (o *PluginsGoldenConfigComplianceFeatureBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigComplianceFeatureBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config compliance feature bulk partial update params
func (o *PluginsGoldenConfigComplianceFeatureBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config compliance feature bulk partial update params
func (o *PluginsGoldenConfigComplianceFeatureBulkPartialUpdateParams) WithContext(ctx context.Context) *PluginsGoldenConfigComplianceFeatureBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config compliance feature bulk partial update params
func (o *PluginsGoldenConfigComplianceFeatureBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config compliance feature bulk partial update params
func (o *PluginsGoldenConfigComplianceFeatureBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigComplianceFeatureBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config compliance feature bulk partial update params
func (o *PluginsGoldenConfigComplianceFeatureBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins golden config compliance feature bulk partial update params
func (o *PluginsGoldenConfigComplianceFeatureBulkPartialUpdateParams) WithData(data *models.ComplianceFeature) *PluginsGoldenConfigComplianceFeatureBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins golden config compliance feature bulk partial update params
func (o *PluginsGoldenConfigComplianceFeatureBulkPartialUpdateParams) SetData(data *models.ComplianceFeature) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigComplianceFeatureBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

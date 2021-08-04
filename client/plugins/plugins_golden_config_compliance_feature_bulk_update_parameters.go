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

// NewPluginsGoldenConfigComplianceFeatureBulkUpdateParams creates a new PluginsGoldenConfigComplianceFeatureBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigComplianceFeatureBulkUpdateParams() *PluginsGoldenConfigComplianceFeatureBulkUpdateParams {
	return &PluginsGoldenConfigComplianceFeatureBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigComplianceFeatureBulkUpdateParamsWithTimeout creates a new PluginsGoldenConfigComplianceFeatureBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigComplianceFeatureBulkUpdateParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigComplianceFeatureBulkUpdateParams {
	return &PluginsGoldenConfigComplianceFeatureBulkUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigComplianceFeatureBulkUpdateParamsWithContext creates a new PluginsGoldenConfigComplianceFeatureBulkUpdateParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigComplianceFeatureBulkUpdateParamsWithContext(ctx context.Context) *PluginsGoldenConfigComplianceFeatureBulkUpdateParams {
	return &PluginsGoldenConfigComplianceFeatureBulkUpdateParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigComplianceFeatureBulkUpdateParamsWithHTTPClient creates a new PluginsGoldenConfigComplianceFeatureBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigComplianceFeatureBulkUpdateParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigComplianceFeatureBulkUpdateParams {
	return &PluginsGoldenConfigComplianceFeatureBulkUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigComplianceFeatureBulkUpdateParams contains all the parameters to send to the API endpoint
   for the plugins golden config compliance feature bulk update operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigComplianceFeatureBulkUpdateParams struct {

	// Data.
	Data *models.ComplianceFeature

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config compliance feature bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigComplianceFeatureBulkUpdateParams) WithDefaults() *PluginsGoldenConfigComplianceFeatureBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config compliance feature bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigComplianceFeatureBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config compliance feature bulk update params
func (o *PluginsGoldenConfigComplianceFeatureBulkUpdateParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigComplianceFeatureBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config compliance feature bulk update params
func (o *PluginsGoldenConfigComplianceFeatureBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config compliance feature bulk update params
func (o *PluginsGoldenConfigComplianceFeatureBulkUpdateParams) WithContext(ctx context.Context) *PluginsGoldenConfigComplianceFeatureBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config compliance feature bulk update params
func (o *PluginsGoldenConfigComplianceFeatureBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config compliance feature bulk update params
func (o *PluginsGoldenConfigComplianceFeatureBulkUpdateParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigComplianceFeatureBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config compliance feature bulk update params
func (o *PluginsGoldenConfigComplianceFeatureBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins golden config compliance feature bulk update params
func (o *PluginsGoldenConfigComplianceFeatureBulkUpdateParams) WithData(data *models.ComplianceFeature) *PluginsGoldenConfigComplianceFeatureBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins golden config compliance feature bulk update params
func (o *PluginsGoldenConfigComplianceFeatureBulkUpdateParams) SetData(data *models.ComplianceFeature) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigComplianceFeatureBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

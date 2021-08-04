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

// NewPluginsGoldenConfigComplianceRuleBulkPartialUpdateParams creates a new PluginsGoldenConfigComplianceRuleBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigComplianceRuleBulkPartialUpdateParams() *PluginsGoldenConfigComplianceRuleBulkPartialUpdateParams {
	return &PluginsGoldenConfigComplianceRuleBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigComplianceRuleBulkPartialUpdateParamsWithTimeout creates a new PluginsGoldenConfigComplianceRuleBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigComplianceRuleBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigComplianceRuleBulkPartialUpdateParams {
	return &PluginsGoldenConfigComplianceRuleBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigComplianceRuleBulkPartialUpdateParamsWithContext creates a new PluginsGoldenConfigComplianceRuleBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigComplianceRuleBulkPartialUpdateParamsWithContext(ctx context.Context) *PluginsGoldenConfigComplianceRuleBulkPartialUpdateParams {
	return &PluginsGoldenConfigComplianceRuleBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigComplianceRuleBulkPartialUpdateParamsWithHTTPClient creates a new PluginsGoldenConfigComplianceRuleBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigComplianceRuleBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigComplianceRuleBulkPartialUpdateParams {
	return &PluginsGoldenConfigComplianceRuleBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigComplianceRuleBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the plugins golden config compliance rule bulk partial update operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigComplianceRuleBulkPartialUpdateParams struct {

	// Data.
	Data *models.ComplianceRule

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config compliance rule bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigComplianceRuleBulkPartialUpdateParams) WithDefaults() *PluginsGoldenConfigComplianceRuleBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config compliance rule bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigComplianceRuleBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config compliance rule bulk partial update params
func (o *PluginsGoldenConfigComplianceRuleBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigComplianceRuleBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config compliance rule bulk partial update params
func (o *PluginsGoldenConfigComplianceRuleBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config compliance rule bulk partial update params
func (o *PluginsGoldenConfigComplianceRuleBulkPartialUpdateParams) WithContext(ctx context.Context) *PluginsGoldenConfigComplianceRuleBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config compliance rule bulk partial update params
func (o *PluginsGoldenConfigComplianceRuleBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config compliance rule bulk partial update params
func (o *PluginsGoldenConfigComplianceRuleBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigComplianceRuleBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config compliance rule bulk partial update params
func (o *PluginsGoldenConfigComplianceRuleBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins golden config compliance rule bulk partial update params
func (o *PluginsGoldenConfigComplianceRuleBulkPartialUpdateParams) WithData(data *models.ComplianceRule) *PluginsGoldenConfigComplianceRuleBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins golden config compliance rule bulk partial update params
func (o *PluginsGoldenConfigComplianceRuleBulkPartialUpdateParams) SetData(data *models.ComplianceRule) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigComplianceRuleBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

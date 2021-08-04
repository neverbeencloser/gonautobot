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

// NewPluginsGoldenConfigComplianceRuleBulkUpdateParams creates a new PluginsGoldenConfigComplianceRuleBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigComplianceRuleBulkUpdateParams() *PluginsGoldenConfigComplianceRuleBulkUpdateParams {
	return &PluginsGoldenConfigComplianceRuleBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigComplianceRuleBulkUpdateParamsWithTimeout creates a new PluginsGoldenConfigComplianceRuleBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigComplianceRuleBulkUpdateParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigComplianceRuleBulkUpdateParams {
	return &PluginsGoldenConfigComplianceRuleBulkUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigComplianceRuleBulkUpdateParamsWithContext creates a new PluginsGoldenConfigComplianceRuleBulkUpdateParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigComplianceRuleBulkUpdateParamsWithContext(ctx context.Context) *PluginsGoldenConfigComplianceRuleBulkUpdateParams {
	return &PluginsGoldenConfigComplianceRuleBulkUpdateParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigComplianceRuleBulkUpdateParamsWithHTTPClient creates a new PluginsGoldenConfigComplianceRuleBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigComplianceRuleBulkUpdateParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigComplianceRuleBulkUpdateParams {
	return &PluginsGoldenConfigComplianceRuleBulkUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigComplianceRuleBulkUpdateParams contains all the parameters to send to the API endpoint
   for the plugins golden config compliance rule bulk update operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigComplianceRuleBulkUpdateParams struct {

	// Data.
	Data *models.ComplianceRule

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config compliance rule bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigComplianceRuleBulkUpdateParams) WithDefaults() *PluginsGoldenConfigComplianceRuleBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config compliance rule bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigComplianceRuleBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config compliance rule bulk update params
func (o *PluginsGoldenConfigComplianceRuleBulkUpdateParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigComplianceRuleBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config compliance rule bulk update params
func (o *PluginsGoldenConfigComplianceRuleBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config compliance rule bulk update params
func (o *PluginsGoldenConfigComplianceRuleBulkUpdateParams) WithContext(ctx context.Context) *PluginsGoldenConfigComplianceRuleBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config compliance rule bulk update params
func (o *PluginsGoldenConfigComplianceRuleBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config compliance rule bulk update params
func (o *PluginsGoldenConfigComplianceRuleBulkUpdateParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigComplianceRuleBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config compliance rule bulk update params
func (o *PluginsGoldenConfigComplianceRuleBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins golden config compliance rule bulk update params
func (o *PluginsGoldenConfigComplianceRuleBulkUpdateParams) WithData(data *models.ComplianceRule) *PluginsGoldenConfigComplianceRuleBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins golden config compliance rule bulk update params
func (o *PluginsGoldenConfigComplianceRuleBulkUpdateParams) SetData(data *models.ComplianceRule) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigComplianceRuleBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

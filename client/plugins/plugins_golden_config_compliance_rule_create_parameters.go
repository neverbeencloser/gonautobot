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

// NewPluginsGoldenConfigComplianceRuleCreateParams creates a new PluginsGoldenConfigComplianceRuleCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigComplianceRuleCreateParams() *PluginsGoldenConfigComplianceRuleCreateParams {
	return &PluginsGoldenConfigComplianceRuleCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigComplianceRuleCreateParamsWithTimeout creates a new PluginsGoldenConfigComplianceRuleCreateParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigComplianceRuleCreateParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigComplianceRuleCreateParams {
	return &PluginsGoldenConfigComplianceRuleCreateParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigComplianceRuleCreateParamsWithContext creates a new PluginsGoldenConfigComplianceRuleCreateParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigComplianceRuleCreateParamsWithContext(ctx context.Context) *PluginsGoldenConfigComplianceRuleCreateParams {
	return &PluginsGoldenConfigComplianceRuleCreateParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigComplianceRuleCreateParamsWithHTTPClient creates a new PluginsGoldenConfigComplianceRuleCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigComplianceRuleCreateParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigComplianceRuleCreateParams {
	return &PluginsGoldenConfigComplianceRuleCreateParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigComplianceRuleCreateParams contains all the parameters to send to the API endpoint
   for the plugins golden config compliance rule create operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigComplianceRuleCreateParams struct {

	// Data.
	Data *models.ComplianceRule

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config compliance rule create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigComplianceRuleCreateParams) WithDefaults() *PluginsGoldenConfigComplianceRuleCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config compliance rule create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigComplianceRuleCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config compliance rule create params
func (o *PluginsGoldenConfigComplianceRuleCreateParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigComplianceRuleCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config compliance rule create params
func (o *PluginsGoldenConfigComplianceRuleCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config compliance rule create params
func (o *PluginsGoldenConfigComplianceRuleCreateParams) WithContext(ctx context.Context) *PluginsGoldenConfigComplianceRuleCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config compliance rule create params
func (o *PluginsGoldenConfigComplianceRuleCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config compliance rule create params
func (o *PluginsGoldenConfigComplianceRuleCreateParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigComplianceRuleCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config compliance rule create params
func (o *PluginsGoldenConfigComplianceRuleCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins golden config compliance rule create params
func (o *PluginsGoldenConfigComplianceRuleCreateParams) WithData(data *models.ComplianceRule) *PluginsGoldenConfigComplianceRuleCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins golden config compliance rule create params
func (o *PluginsGoldenConfigComplianceRuleCreateParams) SetData(data *models.ComplianceRule) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigComplianceRuleCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

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

// NewPluginsGoldenConfigComplianceRuleUpdateParams creates a new PluginsGoldenConfigComplianceRuleUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigComplianceRuleUpdateParams() *PluginsGoldenConfigComplianceRuleUpdateParams {
	return &PluginsGoldenConfigComplianceRuleUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigComplianceRuleUpdateParamsWithTimeout creates a new PluginsGoldenConfigComplianceRuleUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigComplianceRuleUpdateParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigComplianceRuleUpdateParams {
	return &PluginsGoldenConfigComplianceRuleUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigComplianceRuleUpdateParamsWithContext creates a new PluginsGoldenConfigComplianceRuleUpdateParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigComplianceRuleUpdateParamsWithContext(ctx context.Context) *PluginsGoldenConfigComplianceRuleUpdateParams {
	return &PluginsGoldenConfigComplianceRuleUpdateParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigComplianceRuleUpdateParamsWithHTTPClient creates a new PluginsGoldenConfigComplianceRuleUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigComplianceRuleUpdateParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigComplianceRuleUpdateParams {
	return &PluginsGoldenConfigComplianceRuleUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigComplianceRuleUpdateParams contains all the parameters to send to the API endpoint
   for the plugins golden config compliance rule update operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigComplianceRuleUpdateParams struct {

	// Data.
	Data *models.ComplianceRule

	/* ID.

	   A UUID string identifying this compliance rule.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config compliance rule update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigComplianceRuleUpdateParams) WithDefaults() *PluginsGoldenConfigComplianceRuleUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config compliance rule update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigComplianceRuleUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config compliance rule update params
func (o *PluginsGoldenConfigComplianceRuleUpdateParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigComplianceRuleUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config compliance rule update params
func (o *PluginsGoldenConfigComplianceRuleUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config compliance rule update params
func (o *PluginsGoldenConfigComplianceRuleUpdateParams) WithContext(ctx context.Context) *PluginsGoldenConfigComplianceRuleUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config compliance rule update params
func (o *PluginsGoldenConfigComplianceRuleUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config compliance rule update params
func (o *PluginsGoldenConfigComplianceRuleUpdateParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigComplianceRuleUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config compliance rule update params
func (o *PluginsGoldenConfigComplianceRuleUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins golden config compliance rule update params
func (o *PluginsGoldenConfigComplianceRuleUpdateParams) WithData(data *models.ComplianceRule) *PluginsGoldenConfigComplianceRuleUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins golden config compliance rule update params
func (o *PluginsGoldenConfigComplianceRuleUpdateParams) SetData(data *models.ComplianceRule) {
	o.Data = data
}

// WithID adds the id to the plugins golden config compliance rule update params
func (o *PluginsGoldenConfigComplianceRuleUpdateParams) WithID(id strfmt.UUID) *PluginsGoldenConfigComplianceRuleUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins golden config compliance rule update params
func (o *PluginsGoldenConfigComplianceRuleUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigComplianceRuleUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

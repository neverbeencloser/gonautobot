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

// NewPluginsGoldenConfigComplianceRulePartialUpdateParams creates a new PluginsGoldenConfigComplianceRulePartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigComplianceRulePartialUpdateParams() *PluginsGoldenConfigComplianceRulePartialUpdateParams {
	return &PluginsGoldenConfigComplianceRulePartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigComplianceRulePartialUpdateParamsWithTimeout creates a new PluginsGoldenConfigComplianceRulePartialUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigComplianceRulePartialUpdateParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigComplianceRulePartialUpdateParams {
	return &PluginsGoldenConfigComplianceRulePartialUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigComplianceRulePartialUpdateParamsWithContext creates a new PluginsGoldenConfigComplianceRulePartialUpdateParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigComplianceRulePartialUpdateParamsWithContext(ctx context.Context) *PluginsGoldenConfigComplianceRulePartialUpdateParams {
	return &PluginsGoldenConfigComplianceRulePartialUpdateParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigComplianceRulePartialUpdateParamsWithHTTPClient creates a new PluginsGoldenConfigComplianceRulePartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigComplianceRulePartialUpdateParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigComplianceRulePartialUpdateParams {
	return &PluginsGoldenConfigComplianceRulePartialUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigComplianceRulePartialUpdateParams contains all the parameters to send to the API endpoint
   for the plugins golden config compliance rule partial update operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigComplianceRulePartialUpdateParams struct {

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

// WithDefaults hydrates default values in the plugins golden config compliance rule partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigComplianceRulePartialUpdateParams) WithDefaults() *PluginsGoldenConfigComplianceRulePartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config compliance rule partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigComplianceRulePartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config compliance rule partial update params
func (o *PluginsGoldenConfigComplianceRulePartialUpdateParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigComplianceRulePartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config compliance rule partial update params
func (o *PluginsGoldenConfigComplianceRulePartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config compliance rule partial update params
func (o *PluginsGoldenConfigComplianceRulePartialUpdateParams) WithContext(ctx context.Context) *PluginsGoldenConfigComplianceRulePartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config compliance rule partial update params
func (o *PluginsGoldenConfigComplianceRulePartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config compliance rule partial update params
func (o *PluginsGoldenConfigComplianceRulePartialUpdateParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigComplianceRulePartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config compliance rule partial update params
func (o *PluginsGoldenConfigComplianceRulePartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins golden config compliance rule partial update params
func (o *PluginsGoldenConfigComplianceRulePartialUpdateParams) WithData(data *models.ComplianceRule) *PluginsGoldenConfigComplianceRulePartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins golden config compliance rule partial update params
func (o *PluginsGoldenConfigComplianceRulePartialUpdateParams) SetData(data *models.ComplianceRule) {
	o.Data = data
}

// WithID adds the id to the plugins golden config compliance rule partial update params
func (o *PluginsGoldenConfigComplianceRulePartialUpdateParams) WithID(id strfmt.UUID) *PluginsGoldenConfigComplianceRulePartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins golden config compliance rule partial update params
func (o *PluginsGoldenConfigComplianceRulePartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigComplianceRulePartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

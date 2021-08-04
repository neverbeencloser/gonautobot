package plugins

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewPluginsGoldenConfigComplianceRuleDeleteParams creates a new PluginsGoldenConfigComplianceRuleDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigComplianceRuleDeleteParams() *PluginsGoldenConfigComplianceRuleDeleteParams {
	return &PluginsGoldenConfigComplianceRuleDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigComplianceRuleDeleteParamsWithTimeout creates a new PluginsGoldenConfigComplianceRuleDeleteParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigComplianceRuleDeleteParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigComplianceRuleDeleteParams {
	return &PluginsGoldenConfigComplianceRuleDeleteParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigComplianceRuleDeleteParamsWithContext creates a new PluginsGoldenConfigComplianceRuleDeleteParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigComplianceRuleDeleteParamsWithContext(ctx context.Context) *PluginsGoldenConfigComplianceRuleDeleteParams {
	return &PluginsGoldenConfigComplianceRuleDeleteParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigComplianceRuleDeleteParamsWithHTTPClient creates a new PluginsGoldenConfigComplianceRuleDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigComplianceRuleDeleteParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigComplianceRuleDeleteParams {
	return &PluginsGoldenConfigComplianceRuleDeleteParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigComplianceRuleDeleteParams contains all the parameters to send to the API endpoint
   for the plugins golden config compliance rule delete operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigComplianceRuleDeleteParams struct {

	/* ID.

	   A UUID string identifying this compliance rule.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config compliance rule delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigComplianceRuleDeleteParams) WithDefaults() *PluginsGoldenConfigComplianceRuleDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config compliance rule delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigComplianceRuleDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config compliance rule delete params
func (o *PluginsGoldenConfigComplianceRuleDeleteParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigComplianceRuleDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config compliance rule delete params
func (o *PluginsGoldenConfigComplianceRuleDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config compliance rule delete params
func (o *PluginsGoldenConfigComplianceRuleDeleteParams) WithContext(ctx context.Context) *PluginsGoldenConfigComplianceRuleDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config compliance rule delete params
func (o *PluginsGoldenConfigComplianceRuleDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config compliance rule delete params
func (o *PluginsGoldenConfigComplianceRuleDeleteParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigComplianceRuleDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config compliance rule delete params
func (o *PluginsGoldenConfigComplianceRuleDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the plugins golden config compliance rule delete params
func (o *PluginsGoldenConfigComplianceRuleDeleteParams) WithID(id strfmt.UUID) *PluginsGoldenConfigComplianceRuleDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins golden config compliance rule delete params
func (o *PluginsGoldenConfigComplianceRuleDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigComplianceRuleDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

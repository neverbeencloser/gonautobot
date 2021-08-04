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

// NewPluginsGoldenConfigComplianceRuleReadParams creates a new PluginsGoldenConfigComplianceRuleReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigComplianceRuleReadParams() *PluginsGoldenConfigComplianceRuleReadParams {
	return &PluginsGoldenConfigComplianceRuleReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigComplianceRuleReadParamsWithTimeout creates a new PluginsGoldenConfigComplianceRuleReadParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigComplianceRuleReadParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigComplianceRuleReadParams {
	return &PluginsGoldenConfigComplianceRuleReadParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigComplianceRuleReadParamsWithContext creates a new PluginsGoldenConfigComplianceRuleReadParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigComplianceRuleReadParamsWithContext(ctx context.Context) *PluginsGoldenConfigComplianceRuleReadParams {
	return &PluginsGoldenConfigComplianceRuleReadParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigComplianceRuleReadParamsWithHTTPClient creates a new PluginsGoldenConfigComplianceRuleReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigComplianceRuleReadParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigComplianceRuleReadParams {
	return &PluginsGoldenConfigComplianceRuleReadParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigComplianceRuleReadParams contains all the parameters to send to the API endpoint
   for the plugins golden config compliance rule read operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigComplianceRuleReadParams struct {

	/* ID.

	   A UUID string identifying this compliance rule.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config compliance rule read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigComplianceRuleReadParams) WithDefaults() *PluginsGoldenConfigComplianceRuleReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config compliance rule read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigComplianceRuleReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config compliance rule read params
func (o *PluginsGoldenConfigComplianceRuleReadParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigComplianceRuleReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config compliance rule read params
func (o *PluginsGoldenConfigComplianceRuleReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config compliance rule read params
func (o *PluginsGoldenConfigComplianceRuleReadParams) WithContext(ctx context.Context) *PluginsGoldenConfigComplianceRuleReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config compliance rule read params
func (o *PluginsGoldenConfigComplianceRuleReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config compliance rule read params
func (o *PluginsGoldenConfigComplianceRuleReadParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigComplianceRuleReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config compliance rule read params
func (o *PluginsGoldenConfigComplianceRuleReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the plugins golden config compliance rule read params
func (o *PluginsGoldenConfigComplianceRuleReadParams) WithID(id strfmt.UUID) *PluginsGoldenConfigComplianceRuleReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins golden config compliance rule read params
func (o *PluginsGoldenConfigComplianceRuleReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigComplianceRuleReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

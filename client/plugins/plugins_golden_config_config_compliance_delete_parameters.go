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

// NewPluginsGoldenConfigConfigComplianceDeleteParams creates a new PluginsGoldenConfigConfigComplianceDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigConfigComplianceDeleteParams() *PluginsGoldenConfigConfigComplianceDeleteParams {
	return &PluginsGoldenConfigConfigComplianceDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigConfigComplianceDeleteParamsWithTimeout creates a new PluginsGoldenConfigConfigComplianceDeleteParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigConfigComplianceDeleteParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigConfigComplianceDeleteParams {
	return &PluginsGoldenConfigConfigComplianceDeleteParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigConfigComplianceDeleteParamsWithContext creates a new PluginsGoldenConfigConfigComplianceDeleteParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigConfigComplianceDeleteParamsWithContext(ctx context.Context) *PluginsGoldenConfigConfigComplianceDeleteParams {
	return &PluginsGoldenConfigConfigComplianceDeleteParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigConfigComplianceDeleteParamsWithHTTPClient creates a new PluginsGoldenConfigConfigComplianceDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigConfigComplianceDeleteParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigConfigComplianceDeleteParams {
	return &PluginsGoldenConfigConfigComplianceDeleteParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigConfigComplianceDeleteParams contains all the parameters to send to the API endpoint
   for the plugins golden config config compliance delete operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigConfigComplianceDeleteParams struct {

	/* ID.

	   A UUID string identifying this config compliance.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config config compliance delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigConfigComplianceDeleteParams) WithDefaults() *PluginsGoldenConfigConfigComplianceDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config config compliance delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigConfigComplianceDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config config compliance delete params
func (o *PluginsGoldenConfigConfigComplianceDeleteParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigConfigComplianceDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config config compliance delete params
func (o *PluginsGoldenConfigConfigComplianceDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config config compliance delete params
func (o *PluginsGoldenConfigConfigComplianceDeleteParams) WithContext(ctx context.Context) *PluginsGoldenConfigConfigComplianceDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config config compliance delete params
func (o *PluginsGoldenConfigConfigComplianceDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config config compliance delete params
func (o *PluginsGoldenConfigConfigComplianceDeleteParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigConfigComplianceDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config config compliance delete params
func (o *PluginsGoldenConfigConfigComplianceDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the plugins golden config config compliance delete params
func (o *PluginsGoldenConfigConfigComplianceDeleteParams) WithID(id strfmt.UUID) *PluginsGoldenConfigConfigComplianceDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins golden config config compliance delete params
func (o *PluginsGoldenConfigConfigComplianceDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigConfigComplianceDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

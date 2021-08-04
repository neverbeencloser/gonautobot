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

// NewPluginsDataValidationEngineRulesRegexDeleteParams creates a new PluginsDataValidationEngineRulesRegexDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsDataValidationEngineRulesRegexDeleteParams() *PluginsDataValidationEngineRulesRegexDeleteParams {
	return &PluginsDataValidationEngineRulesRegexDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsDataValidationEngineRulesRegexDeleteParamsWithTimeout creates a new PluginsDataValidationEngineRulesRegexDeleteParams object
// with the ability to set a timeout on a request.
func NewPluginsDataValidationEngineRulesRegexDeleteParamsWithTimeout(timeout time.Duration) *PluginsDataValidationEngineRulesRegexDeleteParams {
	return &PluginsDataValidationEngineRulesRegexDeleteParams{
		timeout: timeout,
	}
}

// NewPluginsDataValidationEngineRulesRegexDeleteParamsWithContext creates a new PluginsDataValidationEngineRulesRegexDeleteParams object
// with the ability to set a context for a request.
func NewPluginsDataValidationEngineRulesRegexDeleteParamsWithContext(ctx context.Context) *PluginsDataValidationEngineRulesRegexDeleteParams {
	return &PluginsDataValidationEngineRulesRegexDeleteParams{
		Context: ctx,
	}
}

// NewPluginsDataValidationEngineRulesRegexDeleteParamsWithHTTPClient creates a new PluginsDataValidationEngineRulesRegexDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsDataValidationEngineRulesRegexDeleteParamsWithHTTPClient(client *http.Client) *PluginsDataValidationEngineRulesRegexDeleteParams {
	return &PluginsDataValidationEngineRulesRegexDeleteParams{
		HTTPClient: client,
	}
}

/* PluginsDataValidationEngineRulesRegexDeleteParams contains all the parameters to send to the API endpoint
   for the plugins data validation engine rules regex delete operation.

   Typically these are written to a http.Request.
*/
type PluginsDataValidationEngineRulesRegexDeleteParams struct {

	/* ID.

	   A UUID string identifying this regular expression validation rule.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins data validation engine rules regex delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsDataValidationEngineRulesRegexDeleteParams) WithDefaults() *PluginsDataValidationEngineRulesRegexDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins data validation engine rules regex delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsDataValidationEngineRulesRegexDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins data validation engine rules regex delete params
func (o *PluginsDataValidationEngineRulesRegexDeleteParams) WithTimeout(timeout time.Duration) *PluginsDataValidationEngineRulesRegexDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins data validation engine rules regex delete params
func (o *PluginsDataValidationEngineRulesRegexDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins data validation engine rules regex delete params
func (o *PluginsDataValidationEngineRulesRegexDeleteParams) WithContext(ctx context.Context) *PluginsDataValidationEngineRulesRegexDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins data validation engine rules regex delete params
func (o *PluginsDataValidationEngineRulesRegexDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins data validation engine rules regex delete params
func (o *PluginsDataValidationEngineRulesRegexDeleteParams) WithHTTPClient(client *http.Client) *PluginsDataValidationEngineRulesRegexDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins data validation engine rules regex delete params
func (o *PluginsDataValidationEngineRulesRegexDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the plugins data validation engine rules regex delete params
func (o *PluginsDataValidationEngineRulesRegexDeleteParams) WithID(id strfmt.UUID) *PluginsDataValidationEngineRulesRegexDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins data validation engine rules regex delete params
func (o *PluginsDataValidationEngineRulesRegexDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsDataValidationEngineRulesRegexDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

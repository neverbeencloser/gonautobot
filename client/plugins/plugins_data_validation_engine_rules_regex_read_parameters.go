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

// NewPluginsDataValidationEngineRulesRegexReadParams creates a new PluginsDataValidationEngineRulesRegexReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsDataValidationEngineRulesRegexReadParams() *PluginsDataValidationEngineRulesRegexReadParams {
	return &PluginsDataValidationEngineRulesRegexReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsDataValidationEngineRulesRegexReadParamsWithTimeout creates a new PluginsDataValidationEngineRulesRegexReadParams object
// with the ability to set a timeout on a request.
func NewPluginsDataValidationEngineRulesRegexReadParamsWithTimeout(timeout time.Duration) *PluginsDataValidationEngineRulesRegexReadParams {
	return &PluginsDataValidationEngineRulesRegexReadParams{
		timeout: timeout,
	}
}

// NewPluginsDataValidationEngineRulesRegexReadParamsWithContext creates a new PluginsDataValidationEngineRulesRegexReadParams object
// with the ability to set a context for a request.
func NewPluginsDataValidationEngineRulesRegexReadParamsWithContext(ctx context.Context) *PluginsDataValidationEngineRulesRegexReadParams {
	return &PluginsDataValidationEngineRulesRegexReadParams{
		Context: ctx,
	}
}

// NewPluginsDataValidationEngineRulesRegexReadParamsWithHTTPClient creates a new PluginsDataValidationEngineRulesRegexReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsDataValidationEngineRulesRegexReadParamsWithHTTPClient(client *http.Client) *PluginsDataValidationEngineRulesRegexReadParams {
	return &PluginsDataValidationEngineRulesRegexReadParams{
		HTTPClient: client,
	}
}

/* PluginsDataValidationEngineRulesRegexReadParams contains all the parameters to send to the API endpoint
   for the plugins data validation engine rules regex read operation.

   Typically these are written to a http.Request.
*/
type PluginsDataValidationEngineRulesRegexReadParams struct {

	/* ID.

	   A UUID string identifying this regular expression validation rule.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins data validation engine rules regex read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsDataValidationEngineRulesRegexReadParams) WithDefaults() *PluginsDataValidationEngineRulesRegexReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins data validation engine rules regex read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsDataValidationEngineRulesRegexReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins data validation engine rules regex read params
func (o *PluginsDataValidationEngineRulesRegexReadParams) WithTimeout(timeout time.Duration) *PluginsDataValidationEngineRulesRegexReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins data validation engine rules regex read params
func (o *PluginsDataValidationEngineRulesRegexReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins data validation engine rules regex read params
func (o *PluginsDataValidationEngineRulesRegexReadParams) WithContext(ctx context.Context) *PluginsDataValidationEngineRulesRegexReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins data validation engine rules regex read params
func (o *PluginsDataValidationEngineRulesRegexReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins data validation engine rules regex read params
func (o *PluginsDataValidationEngineRulesRegexReadParams) WithHTTPClient(client *http.Client) *PluginsDataValidationEngineRulesRegexReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins data validation engine rules regex read params
func (o *PluginsDataValidationEngineRulesRegexReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the plugins data validation engine rules regex read params
func (o *PluginsDataValidationEngineRulesRegexReadParams) WithID(id strfmt.UUID) *PluginsDataValidationEngineRulesRegexReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins data validation engine rules regex read params
func (o *PluginsDataValidationEngineRulesRegexReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsDataValidationEngineRulesRegexReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

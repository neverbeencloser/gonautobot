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

// NewPluginsDataValidationEngineRulesMinMaxDeleteParams creates a new PluginsDataValidationEngineRulesMinMaxDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsDataValidationEngineRulesMinMaxDeleteParams() *PluginsDataValidationEngineRulesMinMaxDeleteParams {
	return &PluginsDataValidationEngineRulesMinMaxDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsDataValidationEngineRulesMinMaxDeleteParamsWithTimeout creates a new PluginsDataValidationEngineRulesMinMaxDeleteParams object
// with the ability to set a timeout on a request.
func NewPluginsDataValidationEngineRulesMinMaxDeleteParamsWithTimeout(timeout time.Duration) *PluginsDataValidationEngineRulesMinMaxDeleteParams {
	return &PluginsDataValidationEngineRulesMinMaxDeleteParams{
		timeout: timeout,
	}
}

// NewPluginsDataValidationEngineRulesMinMaxDeleteParamsWithContext creates a new PluginsDataValidationEngineRulesMinMaxDeleteParams object
// with the ability to set a context for a request.
func NewPluginsDataValidationEngineRulesMinMaxDeleteParamsWithContext(ctx context.Context) *PluginsDataValidationEngineRulesMinMaxDeleteParams {
	return &PluginsDataValidationEngineRulesMinMaxDeleteParams{
		Context: ctx,
	}
}

// NewPluginsDataValidationEngineRulesMinMaxDeleteParamsWithHTTPClient creates a new PluginsDataValidationEngineRulesMinMaxDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsDataValidationEngineRulesMinMaxDeleteParamsWithHTTPClient(client *http.Client) *PluginsDataValidationEngineRulesMinMaxDeleteParams {
	return &PluginsDataValidationEngineRulesMinMaxDeleteParams{
		HTTPClient: client,
	}
}

/* PluginsDataValidationEngineRulesMinMaxDeleteParams contains all the parameters to send to the API endpoint
   for the plugins data validation engine rules min max delete operation.

   Typically these are written to a http.Request.
*/
type PluginsDataValidationEngineRulesMinMaxDeleteParams struct {

	/* ID.

	   A UUID string identifying this min max validation rule.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins data validation engine rules min max delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsDataValidationEngineRulesMinMaxDeleteParams) WithDefaults() *PluginsDataValidationEngineRulesMinMaxDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins data validation engine rules min max delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsDataValidationEngineRulesMinMaxDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins data validation engine rules min max delete params
func (o *PluginsDataValidationEngineRulesMinMaxDeleteParams) WithTimeout(timeout time.Duration) *PluginsDataValidationEngineRulesMinMaxDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins data validation engine rules min max delete params
func (o *PluginsDataValidationEngineRulesMinMaxDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins data validation engine rules min max delete params
func (o *PluginsDataValidationEngineRulesMinMaxDeleteParams) WithContext(ctx context.Context) *PluginsDataValidationEngineRulesMinMaxDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins data validation engine rules min max delete params
func (o *PluginsDataValidationEngineRulesMinMaxDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins data validation engine rules min max delete params
func (o *PluginsDataValidationEngineRulesMinMaxDeleteParams) WithHTTPClient(client *http.Client) *PluginsDataValidationEngineRulesMinMaxDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins data validation engine rules min max delete params
func (o *PluginsDataValidationEngineRulesMinMaxDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the plugins data validation engine rules min max delete params
func (o *PluginsDataValidationEngineRulesMinMaxDeleteParams) WithID(id strfmt.UUID) *PluginsDataValidationEngineRulesMinMaxDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins data validation engine rules min max delete params
func (o *PluginsDataValidationEngineRulesMinMaxDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsDataValidationEngineRulesMinMaxDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

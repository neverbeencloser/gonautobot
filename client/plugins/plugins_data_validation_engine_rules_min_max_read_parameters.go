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

// NewPluginsDataValidationEngineRulesMinMaxReadParams creates a new PluginsDataValidationEngineRulesMinMaxReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsDataValidationEngineRulesMinMaxReadParams() *PluginsDataValidationEngineRulesMinMaxReadParams {
	return &PluginsDataValidationEngineRulesMinMaxReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsDataValidationEngineRulesMinMaxReadParamsWithTimeout creates a new PluginsDataValidationEngineRulesMinMaxReadParams object
// with the ability to set a timeout on a request.
func NewPluginsDataValidationEngineRulesMinMaxReadParamsWithTimeout(timeout time.Duration) *PluginsDataValidationEngineRulesMinMaxReadParams {
	return &PluginsDataValidationEngineRulesMinMaxReadParams{
		timeout: timeout,
	}
}

// NewPluginsDataValidationEngineRulesMinMaxReadParamsWithContext creates a new PluginsDataValidationEngineRulesMinMaxReadParams object
// with the ability to set a context for a request.
func NewPluginsDataValidationEngineRulesMinMaxReadParamsWithContext(ctx context.Context) *PluginsDataValidationEngineRulesMinMaxReadParams {
	return &PluginsDataValidationEngineRulesMinMaxReadParams{
		Context: ctx,
	}
}

// NewPluginsDataValidationEngineRulesMinMaxReadParamsWithHTTPClient creates a new PluginsDataValidationEngineRulesMinMaxReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsDataValidationEngineRulesMinMaxReadParamsWithHTTPClient(client *http.Client) *PluginsDataValidationEngineRulesMinMaxReadParams {
	return &PluginsDataValidationEngineRulesMinMaxReadParams{
		HTTPClient: client,
	}
}

/* PluginsDataValidationEngineRulesMinMaxReadParams contains all the parameters to send to the API endpoint
   for the plugins data validation engine rules min max read operation.

   Typically these are written to a http.Request.
*/
type PluginsDataValidationEngineRulesMinMaxReadParams struct {

	/* ID.

	   A UUID string identifying this min max validation rule.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins data validation engine rules min max read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsDataValidationEngineRulesMinMaxReadParams) WithDefaults() *PluginsDataValidationEngineRulesMinMaxReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins data validation engine rules min max read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsDataValidationEngineRulesMinMaxReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins data validation engine rules min max read params
func (o *PluginsDataValidationEngineRulesMinMaxReadParams) WithTimeout(timeout time.Duration) *PluginsDataValidationEngineRulesMinMaxReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins data validation engine rules min max read params
func (o *PluginsDataValidationEngineRulesMinMaxReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins data validation engine rules min max read params
func (o *PluginsDataValidationEngineRulesMinMaxReadParams) WithContext(ctx context.Context) *PluginsDataValidationEngineRulesMinMaxReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins data validation engine rules min max read params
func (o *PluginsDataValidationEngineRulesMinMaxReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins data validation engine rules min max read params
func (o *PluginsDataValidationEngineRulesMinMaxReadParams) WithHTTPClient(client *http.Client) *PluginsDataValidationEngineRulesMinMaxReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins data validation engine rules min max read params
func (o *PluginsDataValidationEngineRulesMinMaxReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the plugins data validation engine rules min max read params
func (o *PluginsDataValidationEngineRulesMinMaxReadParams) WithID(id strfmt.UUID) *PluginsDataValidationEngineRulesMinMaxReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins data validation engine rules min max read params
func (o *PluginsDataValidationEngineRulesMinMaxReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsDataValidationEngineRulesMinMaxReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

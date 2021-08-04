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

// NewPluginsDataValidationEngineRulesMinMaxCreateParams creates a new PluginsDataValidationEngineRulesMinMaxCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsDataValidationEngineRulesMinMaxCreateParams() *PluginsDataValidationEngineRulesMinMaxCreateParams {
	return &PluginsDataValidationEngineRulesMinMaxCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsDataValidationEngineRulesMinMaxCreateParamsWithTimeout creates a new PluginsDataValidationEngineRulesMinMaxCreateParams object
// with the ability to set a timeout on a request.
func NewPluginsDataValidationEngineRulesMinMaxCreateParamsWithTimeout(timeout time.Duration) *PluginsDataValidationEngineRulesMinMaxCreateParams {
	return &PluginsDataValidationEngineRulesMinMaxCreateParams{
		timeout: timeout,
	}
}

// NewPluginsDataValidationEngineRulesMinMaxCreateParamsWithContext creates a new PluginsDataValidationEngineRulesMinMaxCreateParams object
// with the ability to set a context for a request.
func NewPluginsDataValidationEngineRulesMinMaxCreateParamsWithContext(ctx context.Context) *PluginsDataValidationEngineRulesMinMaxCreateParams {
	return &PluginsDataValidationEngineRulesMinMaxCreateParams{
		Context: ctx,
	}
}

// NewPluginsDataValidationEngineRulesMinMaxCreateParamsWithHTTPClient creates a new PluginsDataValidationEngineRulesMinMaxCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsDataValidationEngineRulesMinMaxCreateParamsWithHTTPClient(client *http.Client) *PluginsDataValidationEngineRulesMinMaxCreateParams {
	return &PluginsDataValidationEngineRulesMinMaxCreateParams{
		HTTPClient: client,
	}
}

/* PluginsDataValidationEngineRulesMinMaxCreateParams contains all the parameters to send to the API endpoint
   for the plugins data validation engine rules min max create operation.

   Typically these are written to a http.Request.
*/
type PluginsDataValidationEngineRulesMinMaxCreateParams struct {

	// Data.
	Data *models.MinMaxValidationRule

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins data validation engine rules min max create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsDataValidationEngineRulesMinMaxCreateParams) WithDefaults() *PluginsDataValidationEngineRulesMinMaxCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins data validation engine rules min max create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsDataValidationEngineRulesMinMaxCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins data validation engine rules min max create params
func (o *PluginsDataValidationEngineRulesMinMaxCreateParams) WithTimeout(timeout time.Duration) *PluginsDataValidationEngineRulesMinMaxCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins data validation engine rules min max create params
func (o *PluginsDataValidationEngineRulesMinMaxCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins data validation engine rules min max create params
func (o *PluginsDataValidationEngineRulesMinMaxCreateParams) WithContext(ctx context.Context) *PluginsDataValidationEngineRulesMinMaxCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins data validation engine rules min max create params
func (o *PluginsDataValidationEngineRulesMinMaxCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins data validation engine rules min max create params
func (o *PluginsDataValidationEngineRulesMinMaxCreateParams) WithHTTPClient(client *http.Client) *PluginsDataValidationEngineRulesMinMaxCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins data validation engine rules min max create params
func (o *PluginsDataValidationEngineRulesMinMaxCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins data validation engine rules min max create params
func (o *PluginsDataValidationEngineRulesMinMaxCreateParams) WithData(data *models.MinMaxValidationRule) *PluginsDataValidationEngineRulesMinMaxCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins data validation engine rules min max create params
func (o *PluginsDataValidationEngineRulesMinMaxCreateParams) SetData(data *models.MinMaxValidationRule) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsDataValidationEngineRulesMinMaxCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

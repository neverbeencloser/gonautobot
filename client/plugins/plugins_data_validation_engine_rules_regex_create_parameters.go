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

// NewPluginsDataValidationEngineRulesRegexCreateParams creates a new PluginsDataValidationEngineRulesRegexCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsDataValidationEngineRulesRegexCreateParams() *PluginsDataValidationEngineRulesRegexCreateParams {
	return &PluginsDataValidationEngineRulesRegexCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsDataValidationEngineRulesRegexCreateParamsWithTimeout creates a new PluginsDataValidationEngineRulesRegexCreateParams object
// with the ability to set a timeout on a request.
func NewPluginsDataValidationEngineRulesRegexCreateParamsWithTimeout(timeout time.Duration) *PluginsDataValidationEngineRulesRegexCreateParams {
	return &PluginsDataValidationEngineRulesRegexCreateParams{
		timeout: timeout,
	}
}

// NewPluginsDataValidationEngineRulesRegexCreateParamsWithContext creates a new PluginsDataValidationEngineRulesRegexCreateParams object
// with the ability to set a context for a request.
func NewPluginsDataValidationEngineRulesRegexCreateParamsWithContext(ctx context.Context) *PluginsDataValidationEngineRulesRegexCreateParams {
	return &PluginsDataValidationEngineRulesRegexCreateParams{
		Context: ctx,
	}
}

// NewPluginsDataValidationEngineRulesRegexCreateParamsWithHTTPClient creates a new PluginsDataValidationEngineRulesRegexCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsDataValidationEngineRulesRegexCreateParamsWithHTTPClient(client *http.Client) *PluginsDataValidationEngineRulesRegexCreateParams {
	return &PluginsDataValidationEngineRulesRegexCreateParams{
		HTTPClient: client,
	}
}

/* PluginsDataValidationEngineRulesRegexCreateParams contains all the parameters to send to the API endpoint
   for the plugins data validation engine rules regex create operation.

   Typically these are written to a http.Request.
*/
type PluginsDataValidationEngineRulesRegexCreateParams struct {

	// Data.
	Data *models.RegularExpressionValidationRule

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins data validation engine rules regex create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsDataValidationEngineRulesRegexCreateParams) WithDefaults() *PluginsDataValidationEngineRulesRegexCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins data validation engine rules regex create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsDataValidationEngineRulesRegexCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins data validation engine rules regex create params
func (o *PluginsDataValidationEngineRulesRegexCreateParams) WithTimeout(timeout time.Duration) *PluginsDataValidationEngineRulesRegexCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins data validation engine rules regex create params
func (o *PluginsDataValidationEngineRulesRegexCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins data validation engine rules regex create params
func (o *PluginsDataValidationEngineRulesRegexCreateParams) WithContext(ctx context.Context) *PluginsDataValidationEngineRulesRegexCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins data validation engine rules regex create params
func (o *PluginsDataValidationEngineRulesRegexCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins data validation engine rules regex create params
func (o *PluginsDataValidationEngineRulesRegexCreateParams) WithHTTPClient(client *http.Client) *PluginsDataValidationEngineRulesRegexCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins data validation engine rules regex create params
func (o *PluginsDataValidationEngineRulesRegexCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins data validation engine rules regex create params
func (o *PluginsDataValidationEngineRulesRegexCreateParams) WithData(data *models.RegularExpressionValidationRule) *PluginsDataValidationEngineRulesRegexCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins data validation engine rules regex create params
func (o *PluginsDataValidationEngineRulesRegexCreateParams) SetData(data *models.RegularExpressionValidationRule) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsDataValidationEngineRulesRegexCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

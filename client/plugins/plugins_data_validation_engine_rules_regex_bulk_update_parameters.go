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

// NewPluginsDataValidationEngineRulesRegexBulkUpdateParams creates a new PluginsDataValidationEngineRulesRegexBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsDataValidationEngineRulesRegexBulkUpdateParams() *PluginsDataValidationEngineRulesRegexBulkUpdateParams {
	return &PluginsDataValidationEngineRulesRegexBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsDataValidationEngineRulesRegexBulkUpdateParamsWithTimeout creates a new PluginsDataValidationEngineRulesRegexBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsDataValidationEngineRulesRegexBulkUpdateParamsWithTimeout(timeout time.Duration) *PluginsDataValidationEngineRulesRegexBulkUpdateParams {
	return &PluginsDataValidationEngineRulesRegexBulkUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsDataValidationEngineRulesRegexBulkUpdateParamsWithContext creates a new PluginsDataValidationEngineRulesRegexBulkUpdateParams object
// with the ability to set a context for a request.
func NewPluginsDataValidationEngineRulesRegexBulkUpdateParamsWithContext(ctx context.Context) *PluginsDataValidationEngineRulesRegexBulkUpdateParams {
	return &PluginsDataValidationEngineRulesRegexBulkUpdateParams{
		Context: ctx,
	}
}

// NewPluginsDataValidationEngineRulesRegexBulkUpdateParamsWithHTTPClient creates a new PluginsDataValidationEngineRulesRegexBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsDataValidationEngineRulesRegexBulkUpdateParamsWithHTTPClient(client *http.Client) *PluginsDataValidationEngineRulesRegexBulkUpdateParams {
	return &PluginsDataValidationEngineRulesRegexBulkUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsDataValidationEngineRulesRegexBulkUpdateParams contains all the parameters to send to the API endpoint
   for the plugins data validation engine rules regex bulk update operation.

   Typically these are written to a http.Request.
*/
type PluginsDataValidationEngineRulesRegexBulkUpdateParams struct {

	// Data.
	Data *models.RegularExpressionValidationRule

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins data validation engine rules regex bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsDataValidationEngineRulesRegexBulkUpdateParams) WithDefaults() *PluginsDataValidationEngineRulesRegexBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins data validation engine rules regex bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsDataValidationEngineRulesRegexBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins data validation engine rules regex bulk update params
func (o *PluginsDataValidationEngineRulesRegexBulkUpdateParams) WithTimeout(timeout time.Duration) *PluginsDataValidationEngineRulesRegexBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins data validation engine rules regex bulk update params
func (o *PluginsDataValidationEngineRulesRegexBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins data validation engine rules regex bulk update params
func (o *PluginsDataValidationEngineRulesRegexBulkUpdateParams) WithContext(ctx context.Context) *PluginsDataValidationEngineRulesRegexBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins data validation engine rules regex bulk update params
func (o *PluginsDataValidationEngineRulesRegexBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins data validation engine rules regex bulk update params
func (o *PluginsDataValidationEngineRulesRegexBulkUpdateParams) WithHTTPClient(client *http.Client) *PluginsDataValidationEngineRulesRegexBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins data validation engine rules regex bulk update params
func (o *PluginsDataValidationEngineRulesRegexBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins data validation engine rules regex bulk update params
func (o *PluginsDataValidationEngineRulesRegexBulkUpdateParams) WithData(data *models.RegularExpressionValidationRule) *PluginsDataValidationEngineRulesRegexBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins data validation engine rules regex bulk update params
func (o *PluginsDataValidationEngineRulesRegexBulkUpdateParams) SetData(data *models.RegularExpressionValidationRule) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsDataValidationEngineRulesRegexBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

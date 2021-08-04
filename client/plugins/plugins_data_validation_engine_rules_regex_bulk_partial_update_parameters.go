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

// NewPluginsDataValidationEngineRulesRegexBulkPartialUpdateParams creates a new PluginsDataValidationEngineRulesRegexBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsDataValidationEngineRulesRegexBulkPartialUpdateParams() *PluginsDataValidationEngineRulesRegexBulkPartialUpdateParams {
	return &PluginsDataValidationEngineRulesRegexBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsDataValidationEngineRulesRegexBulkPartialUpdateParamsWithTimeout creates a new PluginsDataValidationEngineRulesRegexBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsDataValidationEngineRulesRegexBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *PluginsDataValidationEngineRulesRegexBulkPartialUpdateParams {
	return &PluginsDataValidationEngineRulesRegexBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsDataValidationEngineRulesRegexBulkPartialUpdateParamsWithContext creates a new PluginsDataValidationEngineRulesRegexBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewPluginsDataValidationEngineRulesRegexBulkPartialUpdateParamsWithContext(ctx context.Context) *PluginsDataValidationEngineRulesRegexBulkPartialUpdateParams {
	return &PluginsDataValidationEngineRulesRegexBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewPluginsDataValidationEngineRulesRegexBulkPartialUpdateParamsWithHTTPClient creates a new PluginsDataValidationEngineRulesRegexBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsDataValidationEngineRulesRegexBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *PluginsDataValidationEngineRulesRegexBulkPartialUpdateParams {
	return &PluginsDataValidationEngineRulesRegexBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsDataValidationEngineRulesRegexBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the plugins data validation engine rules regex bulk partial update operation.

   Typically these are written to a http.Request.
*/
type PluginsDataValidationEngineRulesRegexBulkPartialUpdateParams struct {

	// Data.
	Data *models.RegularExpressionValidationRule

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins data validation engine rules regex bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsDataValidationEngineRulesRegexBulkPartialUpdateParams) WithDefaults() *PluginsDataValidationEngineRulesRegexBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins data validation engine rules regex bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsDataValidationEngineRulesRegexBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins data validation engine rules regex bulk partial update params
func (o *PluginsDataValidationEngineRulesRegexBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *PluginsDataValidationEngineRulesRegexBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins data validation engine rules regex bulk partial update params
func (o *PluginsDataValidationEngineRulesRegexBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins data validation engine rules regex bulk partial update params
func (o *PluginsDataValidationEngineRulesRegexBulkPartialUpdateParams) WithContext(ctx context.Context) *PluginsDataValidationEngineRulesRegexBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins data validation engine rules regex bulk partial update params
func (o *PluginsDataValidationEngineRulesRegexBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins data validation engine rules regex bulk partial update params
func (o *PluginsDataValidationEngineRulesRegexBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *PluginsDataValidationEngineRulesRegexBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins data validation engine rules regex bulk partial update params
func (o *PluginsDataValidationEngineRulesRegexBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins data validation engine rules regex bulk partial update params
func (o *PluginsDataValidationEngineRulesRegexBulkPartialUpdateParams) WithData(data *models.RegularExpressionValidationRule) *PluginsDataValidationEngineRulesRegexBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins data validation engine rules regex bulk partial update params
func (o *PluginsDataValidationEngineRulesRegexBulkPartialUpdateParams) SetData(data *models.RegularExpressionValidationRule) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsDataValidationEngineRulesRegexBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

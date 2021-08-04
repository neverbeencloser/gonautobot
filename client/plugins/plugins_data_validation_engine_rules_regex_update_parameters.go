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

// NewPluginsDataValidationEngineRulesRegexUpdateParams creates a new PluginsDataValidationEngineRulesRegexUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsDataValidationEngineRulesRegexUpdateParams() *PluginsDataValidationEngineRulesRegexUpdateParams {
	return &PluginsDataValidationEngineRulesRegexUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsDataValidationEngineRulesRegexUpdateParamsWithTimeout creates a new PluginsDataValidationEngineRulesRegexUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsDataValidationEngineRulesRegexUpdateParamsWithTimeout(timeout time.Duration) *PluginsDataValidationEngineRulesRegexUpdateParams {
	return &PluginsDataValidationEngineRulesRegexUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsDataValidationEngineRulesRegexUpdateParamsWithContext creates a new PluginsDataValidationEngineRulesRegexUpdateParams object
// with the ability to set a context for a request.
func NewPluginsDataValidationEngineRulesRegexUpdateParamsWithContext(ctx context.Context) *PluginsDataValidationEngineRulesRegexUpdateParams {
	return &PluginsDataValidationEngineRulesRegexUpdateParams{
		Context: ctx,
	}
}

// NewPluginsDataValidationEngineRulesRegexUpdateParamsWithHTTPClient creates a new PluginsDataValidationEngineRulesRegexUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsDataValidationEngineRulesRegexUpdateParamsWithHTTPClient(client *http.Client) *PluginsDataValidationEngineRulesRegexUpdateParams {
	return &PluginsDataValidationEngineRulesRegexUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsDataValidationEngineRulesRegexUpdateParams contains all the parameters to send to the API endpoint
   for the plugins data validation engine rules regex update operation.

   Typically these are written to a http.Request.
*/
type PluginsDataValidationEngineRulesRegexUpdateParams struct {

	// Data.
	Data *models.RegularExpressionValidationRule

	/* ID.

	   A UUID string identifying this regular expression validation rule.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins data validation engine rules regex update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsDataValidationEngineRulesRegexUpdateParams) WithDefaults() *PluginsDataValidationEngineRulesRegexUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins data validation engine rules regex update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsDataValidationEngineRulesRegexUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins data validation engine rules regex update params
func (o *PluginsDataValidationEngineRulesRegexUpdateParams) WithTimeout(timeout time.Duration) *PluginsDataValidationEngineRulesRegexUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins data validation engine rules regex update params
func (o *PluginsDataValidationEngineRulesRegexUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins data validation engine rules regex update params
func (o *PluginsDataValidationEngineRulesRegexUpdateParams) WithContext(ctx context.Context) *PluginsDataValidationEngineRulesRegexUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins data validation engine rules regex update params
func (o *PluginsDataValidationEngineRulesRegexUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins data validation engine rules regex update params
func (o *PluginsDataValidationEngineRulesRegexUpdateParams) WithHTTPClient(client *http.Client) *PluginsDataValidationEngineRulesRegexUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins data validation engine rules regex update params
func (o *PluginsDataValidationEngineRulesRegexUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins data validation engine rules regex update params
func (o *PluginsDataValidationEngineRulesRegexUpdateParams) WithData(data *models.RegularExpressionValidationRule) *PluginsDataValidationEngineRulesRegexUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins data validation engine rules regex update params
func (o *PluginsDataValidationEngineRulesRegexUpdateParams) SetData(data *models.RegularExpressionValidationRule) {
	o.Data = data
}

// WithID adds the id to the plugins data validation engine rules regex update params
func (o *PluginsDataValidationEngineRulesRegexUpdateParams) WithID(id strfmt.UUID) *PluginsDataValidationEngineRulesRegexUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins data validation engine rules regex update params
func (o *PluginsDataValidationEngineRulesRegexUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsDataValidationEngineRulesRegexUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

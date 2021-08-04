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

// NewPluginsDataValidationEngineRulesRegexPartialUpdateParams creates a new PluginsDataValidationEngineRulesRegexPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsDataValidationEngineRulesRegexPartialUpdateParams() *PluginsDataValidationEngineRulesRegexPartialUpdateParams {
	return &PluginsDataValidationEngineRulesRegexPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsDataValidationEngineRulesRegexPartialUpdateParamsWithTimeout creates a new PluginsDataValidationEngineRulesRegexPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsDataValidationEngineRulesRegexPartialUpdateParamsWithTimeout(timeout time.Duration) *PluginsDataValidationEngineRulesRegexPartialUpdateParams {
	return &PluginsDataValidationEngineRulesRegexPartialUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsDataValidationEngineRulesRegexPartialUpdateParamsWithContext creates a new PluginsDataValidationEngineRulesRegexPartialUpdateParams object
// with the ability to set a context for a request.
func NewPluginsDataValidationEngineRulesRegexPartialUpdateParamsWithContext(ctx context.Context) *PluginsDataValidationEngineRulesRegexPartialUpdateParams {
	return &PluginsDataValidationEngineRulesRegexPartialUpdateParams{
		Context: ctx,
	}
}

// NewPluginsDataValidationEngineRulesRegexPartialUpdateParamsWithHTTPClient creates a new PluginsDataValidationEngineRulesRegexPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsDataValidationEngineRulesRegexPartialUpdateParamsWithHTTPClient(client *http.Client) *PluginsDataValidationEngineRulesRegexPartialUpdateParams {
	return &PluginsDataValidationEngineRulesRegexPartialUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsDataValidationEngineRulesRegexPartialUpdateParams contains all the parameters to send to the API endpoint
   for the plugins data validation engine rules regex partial update operation.

   Typically these are written to a http.Request.
*/
type PluginsDataValidationEngineRulesRegexPartialUpdateParams struct {

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

// WithDefaults hydrates default values in the plugins data validation engine rules regex partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsDataValidationEngineRulesRegexPartialUpdateParams) WithDefaults() *PluginsDataValidationEngineRulesRegexPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins data validation engine rules regex partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsDataValidationEngineRulesRegexPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins data validation engine rules regex partial update params
func (o *PluginsDataValidationEngineRulesRegexPartialUpdateParams) WithTimeout(timeout time.Duration) *PluginsDataValidationEngineRulesRegexPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins data validation engine rules regex partial update params
func (o *PluginsDataValidationEngineRulesRegexPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins data validation engine rules regex partial update params
func (o *PluginsDataValidationEngineRulesRegexPartialUpdateParams) WithContext(ctx context.Context) *PluginsDataValidationEngineRulesRegexPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins data validation engine rules regex partial update params
func (o *PluginsDataValidationEngineRulesRegexPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins data validation engine rules regex partial update params
func (o *PluginsDataValidationEngineRulesRegexPartialUpdateParams) WithHTTPClient(client *http.Client) *PluginsDataValidationEngineRulesRegexPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins data validation engine rules regex partial update params
func (o *PluginsDataValidationEngineRulesRegexPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins data validation engine rules regex partial update params
func (o *PluginsDataValidationEngineRulesRegexPartialUpdateParams) WithData(data *models.RegularExpressionValidationRule) *PluginsDataValidationEngineRulesRegexPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins data validation engine rules regex partial update params
func (o *PluginsDataValidationEngineRulesRegexPartialUpdateParams) SetData(data *models.RegularExpressionValidationRule) {
	o.Data = data
}

// WithID adds the id to the plugins data validation engine rules regex partial update params
func (o *PluginsDataValidationEngineRulesRegexPartialUpdateParams) WithID(id strfmt.UUID) *PluginsDataValidationEngineRulesRegexPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins data validation engine rules regex partial update params
func (o *PluginsDataValidationEngineRulesRegexPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsDataValidationEngineRulesRegexPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

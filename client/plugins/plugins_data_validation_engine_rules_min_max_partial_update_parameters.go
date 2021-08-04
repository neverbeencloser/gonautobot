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

// NewPluginsDataValidationEngineRulesMinMaxPartialUpdateParams creates a new PluginsDataValidationEngineRulesMinMaxPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsDataValidationEngineRulesMinMaxPartialUpdateParams() *PluginsDataValidationEngineRulesMinMaxPartialUpdateParams {
	return &PluginsDataValidationEngineRulesMinMaxPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsDataValidationEngineRulesMinMaxPartialUpdateParamsWithTimeout creates a new PluginsDataValidationEngineRulesMinMaxPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsDataValidationEngineRulesMinMaxPartialUpdateParamsWithTimeout(timeout time.Duration) *PluginsDataValidationEngineRulesMinMaxPartialUpdateParams {
	return &PluginsDataValidationEngineRulesMinMaxPartialUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsDataValidationEngineRulesMinMaxPartialUpdateParamsWithContext creates a new PluginsDataValidationEngineRulesMinMaxPartialUpdateParams object
// with the ability to set a context for a request.
func NewPluginsDataValidationEngineRulesMinMaxPartialUpdateParamsWithContext(ctx context.Context) *PluginsDataValidationEngineRulesMinMaxPartialUpdateParams {
	return &PluginsDataValidationEngineRulesMinMaxPartialUpdateParams{
		Context: ctx,
	}
}

// NewPluginsDataValidationEngineRulesMinMaxPartialUpdateParamsWithHTTPClient creates a new PluginsDataValidationEngineRulesMinMaxPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsDataValidationEngineRulesMinMaxPartialUpdateParamsWithHTTPClient(client *http.Client) *PluginsDataValidationEngineRulesMinMaxPartialUpdateParams {
	return &PluginsDataValidationEngineRulesMinMaxPartialUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsDataValidationEngineRulesMinMaxPartialUpdateParams contains all the parameters to send to the API endpoint
   for the plugins data validation engine rules min max partial update operation.

   Typically these are written to a http.Request.
*/
type PluginsDataValidationEngineRulesMinMaxPartialUpdateParams struct {

	// Data.
	Data *models.MinMaxValidationRule

	/* ID.

	   A UUID string identifying this min max validation rule.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins data validation engine rules min max partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsDataValidationEngineRulesMinMaxPartialUpdateParams) WithDefaults() *PluginsDataValidationEngineRulesMinMaxPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins data validation engine rules min max partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsDataValidationEngineRulesMinMaxPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins data validation engine rules min max partial update params
func (o *PluginsDataValidationEngineRulesMinMaxPartialUpdateParams) WithTimeout(timeout time.Duration) *PluginsDataValidationEngineRulesMinMaxPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins data validation engine rules min max partial update params
func (o *PluginsDataValidationEngineRulesMinMaxPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins data validation engine rules min max partial update params
func (o *PluginsDataValidationEngineRulesMinMaxPartialUpdateParams) WithContext(ctx context.Context) *PluginsDataValidationEngineRulesMinMaxPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins data validation engine rules min max partial update params
func (o *PluginsDataValidationEngineRulesMinMaxPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins data validation engine rules min max partial update params
func (o *PluginsDataValidationEngineRulesMinMaxPartialUpdateParams) WithHTTPClient(client *http.Client) *PluginsDataValidationEngineRulesMinMaxPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins data validation engine rules min max partial update params
func (o *PluginsDataValidationEngineRulesMinMaxPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins data validation engine rules min max partial update params
func (o *PluginsDataValidationEngineRulesMinMaxPartialUpdateParams) WithData(data *models.MinMaxValidationRule) *PluginsDataValidationEngineRulesMinMaxPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins data validation engine rules min max partial update params
func (o *PluginsDataValidationEngineRulesMinMaxPartialUpdateParams) SetData(data *models.MinMaxValidationRule) {
	o.Data = data
}

// WithID adds the id to the plugins data validation engine rules min max partial update params
func (o *PluginsDataValidationEngineRulesMinMaxPartialUpdateParams) WithID(id strfmt.UUID) *PluginsDataValidationEngineRulesMinMaxPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins data validation engine rules min max partial update params
func (o *PluginsDataValidationEngineRulesMinMaxPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsDataValidationEngineRulesMinMaxPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

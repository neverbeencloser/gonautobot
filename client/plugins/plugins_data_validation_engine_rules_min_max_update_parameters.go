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

// NewPluginsDataValidationEngineRulesMinMaxUpdateParams creates a new PluginsDataValidationEngineRulesMinMaxUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsDataValidationEngineRulesMinMaxUpdateParams() *PluginsDataValidationEngineRulesMinMaxUpdateParams {
	return &PluginsDataValidationEngineRulesMinMaxUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsDataValidationEngineRulesMinMaxUpdateParamsWithTimeout creates a new PluginsDataValidationEngineRulesMinMaxUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsDataValidationEngineRulesMinMaxUpdateParamsWithTimeout(timeout time.Duration) *PluginsDataValidationEngineRulesMinMaxUpdateParams {
	return &PluginsDataValidationEngineRulesMinMaxUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsDataValidationEngineRulesMinMaxUpdateParamsWithContext creates a new PluginsDataValidationEngineRulesMinMaxUpdateParams object
// with the ability to set a context for a request.
func NewPluginsDataValidationEngineRulesMinMaxUpdateParamsWithContext(ctx context.Context) *PluginsDataValidationEngineRulesMinMaxUpdateParams {
	return &PluginsDataValidationEngineRulesMinMaxUpdateParams{
		Context: ctx,
	}
}

// NewPluginsDataValidationEngineRulesMinMaxUpdateParamsWithHTTPClient creates a new PluginsDataValidationEngineRulesMinMaxUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsDataValidationEngineRulesMinMaxUpdateParamsWithHTTPClient(client *http.Client) *PluginsDataValidationEngineRulesMinMaxUpdateParams {
	return &PluginsDataValidationEngineRulesMinMaxUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsDataValidationEngineRulesMinMaxUpdateParams contains all the parameters to send to the API endpoint
   for the plugins data validation engine rules min max update operation.

   Typically these are written to a http.Request.
*/
type PluginsDataValidationEngineRulesMinMaxUpdateParams struct {

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

// WithDefaults hydrates default values in the plugins data validation engine rules min max update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsDataValidationEngineRulesMinMaxUpdateParams) WithDefaults() *PluginsDataValidationEngineRulesMinMaxUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins data validation engine rules min max update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsDataValidationEngineRulesMinMaxUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins data validation engine rules min max update params
func (o *PluginsDataValidationEngineRulesMinMaxUpdateParams) WithTimeout(timeout time.Duration) *PluginsDataValidationEngineRulesMinMaxUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins data validation engine rules min max update params
func (o *PluginsDataValidationEngineRulesMinMaxUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins data validation engine rules min max update params
func (o *PluginsDataValidationEngineRulesMinMaxUpdateParams) WithContext(ctx context.Context) *PluginsDataValidationEngineRulesMinMaxUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins data validation engine rules min max update params
func (o *PluginsDataValidationEngineRulesMinMaxUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins data validation engine rules min max update params
func (o *PluginsDataValidationEngineRulesMinMaxUpdateParams) WithHTTPClient(client *http.Client) *PluginsDataValidationEngineRulesMinMaxUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins data validation engine rules min max update params
func (o *PluginsDataValidationEngineRulesMinMaxUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins data validation engine rules min max update params
func (o *PluginsDataValidationEngineRulesMinMaxUpdateParams) WithData(data *models.MinMaxValidationRule) *PluginsDataValidationEngineRulesMinMaxUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins data validation engine rules min max update params
func (o *PluginsDataValidationEngineRulesMinMaxUpdateParams) SetData(data *models.MinMaxValidationRule) {
	o.Data = data
}

// WithID adds the id to the plugins data validation engine rules min max update params
func (o *PluginsDataValidationEngineRulesMinMaxUpdateParams) WithID(id strfmt.UUID) *PluginsDataValidationEngineRulesMinMaxUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins data validation engine rules min max update params
func (o *PluginsDataValidationEngineRulesMinMaxUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsDataValidationEngineRulesMinMaxUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

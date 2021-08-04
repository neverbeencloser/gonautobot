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

// NewPluginsDataValidationEngineRulesMinMaxBulkPartialUpdateParams creates a new PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsDataValidationEngineRulesMinMaxBulkPartialUpdateParams() *PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateParams {
	return &PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsDataValidationEngineRulesMinMaxBulkPartialUpdateParamsWithTimeout creates a new PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsDataValidationEngineRulesMinMaxBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateParams {
	return &PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsDataValidationEngineRulesMinMaxBulkPartialUpdateParamsWithContext creates a new PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewPluginsDataValidationEngineRulesMinMaxBulkPartialUpdateParamsWithContext(ctx context.Context) *PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateParams {
	return &PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewPluginsDataValidationEngineRulesMinMaxBulkPartialUpdateParamsWithHTTPClient creates a new PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsDataValidationEngineRulesMinMaxBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateParams {
	return &PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the plugins data validation engine rules min max bulk partial update operation.

   Typically these are written to a http.Request.
*/
type PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateParams struct {

	// Data.
	Data *models.MinMaxValidationRule

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins data validation engine rules min max bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateParams) WithDefaults() *PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins data validation engine rules min max bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins data validation engine rules min max bulk partial update params
func (o *PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins data validation engine rules min max bulk partial update params
func (o *PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins data validation engine rules min max bulk partial update params
func (o *PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateParams) WithContext(ctx context.Context) *PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins data validation engine rules min max bulk partial update params
func (o *PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins data validation engine rules min max bulk partial update params
func (o *PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins data validation engine rules min max bulk partial update params
func (o *PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins data validation engine rules min max bulk partial update params
func (o *PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateParams) WithData(data *models.MinMaxValidationRule) *PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins data validation engine rules min max bulk partial update params
func (o *PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateParams) SetData(data *models.MinMaxValidationRule) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

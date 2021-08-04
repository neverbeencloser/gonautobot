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

// NewPluginsDataValidationEngineRulesMinMaxBulkUpdateParams creates a new PluginsDataValidationEngineRulesMinMaxBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsDataValidationEngineRulesMinMaxBulkUpdateParams() *PluginsDataValidationEngineRulesMinMaxBulkUpdateParams {
	return &PluginsDataValidationEngineRulesMinMaxBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsDataValidationEngineRulesMinMaxBulkUpdateParamsWithTimeout creates a new PluginsDataValidationEngineRulesMinMaxBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsDataValidationEngineRulesMinMaxBulkUpdateParamsWithTimeout(timeout time.Duration) *PluginsDataValidationEngineRulesMinMaxBulkUpdateParams {
	return &PluginsDataValidationEngineRulesMinMaxBulkUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsDataValidationEngineRulesMinMaxBulkUpdateParamsWithContext creates a new PluginsDataValidationEngineRulesMinMaxBulkUpdateParams object
// with the ability to set a context for a request.
func NewPluginsDataValidationEngineRulesMinMaxBulkUpdateParamsWithContext(ctx context.Context) *PluginsDataValidationEngineRulesMinMaxBulkUpdateParams {
	return &PluginsDataValidationEngineRulesMinMaxBulkUpdateParams{
		Context: ctx,
	}
}

// NewPluginsDataValidationEngineRulesMinMaxBulkUpdateParamsWithHTTPClient creates a new PluginsDataValidationEngineRulesMinMaxBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsDataValidationEngineRulesMinMaxBulkUpdateParamsWithHTTPClient(client *http.Client) *PluginsDataValidationEngineRulesMinMaxBulkUpdateParams {
	return &PluginsDataValidationEngineRulesMinMaxBulkUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsDataValidationEngineRulesMinMaxBulkUpdateParams contains all the parameters to send to the API endpoint
   for the plugins data validation engine rules min max bulk update operation.

   Typically these are written to a http.Request.
*/
type PluginsDataValidationEngineRulesMinMaxBulkUpdateParams struct {

	// Data.
	Data *models.MinMaxValidationRule

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins data validation engine rules min max bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsDataValidationEngineRulesMinMaxBulkUpdateParams) WithDefaults() *PluginsDataValidationEngineRulesMinMaxBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins data validation engine rules min max bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsDataValidationEngineRulesMinMaxBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins data validation engine rules min max bulk update params
func (o *PluginsDataValidationEngineRulesMinMaxBulkUpdateParams) WithTimeout(timeout time.Duration) *PluginsDataValidationEngineRulesMinMaxBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins data validation engine rules min max bulk update params
func (o *PluginsDataValidationEngineRulesMinMaxBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins data validation engine rules min max bulk update params
func (o *PluginsDataValidationEngineRulesMinMaxBulkUpdateParams) WithContext(ctx context.Context) *PluginsDataValidationEngineRulesMinMaxBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins data validation engine rules min max bulk update params
func (o *PluginsDataValidationEngineRulesMinMaxBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins data validation engine rules min max bulk update params
func (o *PluginsDataValidationEngineRulesMinMaxBulkUpdateParams) WithHTTPClient(client *http.Client) *PluginsDataValidationEngineRulesMinMaxBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins data validation engine rules min max bulk update params
func (o *PluginsDataValidationEngineRulesMinMaxBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins data validation engine rules min max bulk update params
func (o *PluginsDataValidationEngineRulesMinMaxBulkUpdateParams) WithData(data *models.MinMaxValidationRule) *PluginsDataValidationEngineRulesMinMaxBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins data validation engine rules min max bulk update params
func (o *PluginsDataValidationEngineRulesMinMaxBulkUpdateParams) SetData(data *models.MinMaxValidationRule) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsDataValidationEngineRulesMinMaxBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

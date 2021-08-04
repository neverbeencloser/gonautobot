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

// NewPluginsGoldenConfigGoldenConfigSettingsBulkUpdateParams creates a new PluginsGoldenConfigGoldenConfigSettingsBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigGoldenConfigSettingsBulkUpdateParams() *PluginsGoldenConfigGoldenConfigSettingsBulkUpdateParams {
	return &PluginsGoldenConfigGoldenConfigSettingsBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigGoldenConfigSettingsBulkUpdateParamsWithTimeout creates a new PluginsGoldenConfigGoldenConfigSettingsBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigGoldenConfigSettingsBulkUpdateParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigGoldenConfigSettingsBulkUpdateParams {
	return &PluginsGoldenConfigGoldenConfigSettingsBulkUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigGoldenConfigSettingsBulkUpdateParamsWithContext creates a new PluginsGoldenConfigGoldenConfigSettingsBulkUpdateParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigGoldenConfigSettingsBulkUpdateParamsWithContext(ctx context.Context) *PluginsGoldenConfigGoldenConfigSettingsBulkUpdateParams {
	return &PluginsGoldenConfigGoldenConfigSettingsBulkUpdateParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigGoldenConfigSettingsBulkUpdateParamsWithHTTPClient creates a new PluginsGoldenConfigGoldenConfigSettingsBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigGoldenConfigSettingsBulkUpdateParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigGoldenConfigSettingsBulkUpdateParams {
	return &PluginsGoldenConfigGoldenConfigSettingsBulkUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigGoldenConfigSettingsBulkUpdateParams contains all the parameters to send to the API endpoint
   for the plugins golden config golden config settings bulk update operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigGoldenConfigSettingsBulkUpdateParams struct {

	// Data.
	Data *models.GoldenConfigSetting

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config golden config settings bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigGoldenConfigSettingsBulkUpdateParams) WithDefaults() *PluginsGoldenConfigGoldenConfigSettingsBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config golden config settings bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigGoldenConfigSettingsBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config golden config settings bulk update params
func (o *PluginsGoldenConfigGoldenConfigSettingsBulkUpdateParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigGoldenConfigSettingsBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config golden config settings bulk update params
func (o *PluginsGoldenConfigGoldenConfigSettingsBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config golden config settings bulk update params
func (o *PluginsGoldenConfigGoldenConfigSettingsBulkUpdateParams) WithContext(ctx context.Context) *PluginsGoldenConfigGoldenConfigSettingsBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config golden config settings bulk update params
func (o *PluginsGoldenConfigGoldenConfigSettingsBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config golden config settings bulk update params
func (o *PluginsGoldenConfigGoldenConfigSettingsBulkUpdateParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigGoldenConfigSettingsBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config golden config settings bulk update params
func (o *PluginsGoldenConfigGoldenConfigSettingsBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins golden config golden config settings bulk update params
func (o *PluginsGoldenConfigGoldenConfigSettingsBulkUpdateParams) WithData(data *models.GoldenConfigSetting) *PluginsGoldenConfigGoldenConfigSettingsBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins golden config golden config settings bulk update params
func (o *PluginsGoldenConfigGoldenConfigSettingsBulkUpdateParams) SetData(data *models.GoldenConfigSetting) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigGoldenConfigSettingsBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

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

// NewPluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateParams creates a new PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateParams() *PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateParams {
	return &PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateParamsWithTimeout creates a new PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateParams {
	return &PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateParamsWithContext creates a new PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateParamsWithContext(ctx context.Context) *PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateParams {
	return &PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateParamsWithHTTPClient creates a new PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateParams {
	return &PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the plugins golden config golden config settings bulk partial update operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateParams struct {

	// Data.
	Data *models.GoldenConfigSetting

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config golden config settings bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateParams) WithDefaults() *PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config golden config settings bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config golden config settings bulk partial update params
func (o *PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config golden config settings bulk partial update params
func (o *PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config golden config settings bulk partial update params
func (o *PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateParams) WithContext(ctx context.Context) *PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config golden config settings bulk partial update params
func (o *PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config golden config settings bulk partial update params
func (o *PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config golden config settings bulk partial update params
func (o *PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins golden config golden config settings bulk partial update params
func (o *PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateParams) WithData(data *models.GoldenConfigSetting) *PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins golden config golden config settings bulk partial update params
func (o *PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateParams) SetData(data *models.GoldenConfigSetting) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

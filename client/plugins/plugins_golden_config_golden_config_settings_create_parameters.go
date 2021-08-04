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

// NewPluginsGoldenConfigGoldenConfigSettingsCreateParams creates a new PluginsGoldenConfigGoldenConfigSettingsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigGoldenConfigSettingsCreateParams() *PluginsGoldenConfigGoldenConfigSettingsCreateParams {
	return &PluginsGoldenConfigGoldenConfigSettingsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigGoldenConfigSettingsCreateParamsWithTimeout creates a new PluginsGoldenConfigGoldenConfigSettingsCreateParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigGoldenConfigSettingsCreateParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigGoldenConfigSettingsCreateParams {
	return &PluginsGoldenConfigGoldenConfigSettingsCreateParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigGoldenConfigSettingsCreateParamsWithContext creates a new PluginsGoldenConfigGoldenConfigSettingsCreateParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigGoldenConfigSettingsCreateParamsWithContext(ctx context.Context) *PluginsGoldenConfigGoldenConfigSettingsCreateParams {
	return &PluginsGoldenConfigGoldenConfigSettingsCreateParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigGoldenConfigSettingsCreateParamsWithHTTPClient creates a new PluginsGoldenConfigGoldenConfigSettingsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigGoldenConfigSettingsCreateParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigGoldenConfigSettingsCreateParams {
	return &PluginsGoldenConfigGoldenConfigSettingsCreateParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigGoldenConfigSettingsCreateParams contains all the parameters to send to the API endpoint
   for the plugins golden config golden config settings create operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigGoldenConfigSettingsCreateParams struct {

	// Data.
	Data *models.GoldenConfigSetting

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config golden config settings create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigGoldenConfigSettingsCreateParams) WithDefaults() *PluginsGoldenConfigGoldenConfigSettingsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config golden config settings create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigGoldenConfigSettingsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config golden config settings create params
func (o *PluginsGoldenConfigGoldenConfigSettingsCreateParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigGoldenConfigSettingsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config golden config settings create params
func (o *PluginsGoldenConfigGoldenConfigSettingsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config golden config settings create params
func (o *PluginsGoldenConfigGoldenConfigSettingsCreateParams) WithContext(ctx context.Context) *PluginsGoldenConfigGoldenConfigSettingsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config golden config settings create params
func (o *PluginsGoldenConfigGoldenConfigSettingsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config golden config settings create params
func (o *PluginsGoldenConfigGoldenConfigSettingsCreateParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigGoldenConfigSettingsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config golden config settings create params
func (o *PluginsGoldenConfigGoldenConfigSettingsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins golden config golden config settings create params
func (o *PluginsGoldenConfigGoldenConfigSettingsCreateParams) WithData(data *models.GoldenConfigSetting) *PluginsGoldenConfigGoldenConfigSettingsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins golden config golden config settings create params
func (o *PluginsGoldenConfigGoldenConfigSettingsCreateParams) SetData(data *models.GoldenConfigSetting) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigGoldenConfigSettingsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

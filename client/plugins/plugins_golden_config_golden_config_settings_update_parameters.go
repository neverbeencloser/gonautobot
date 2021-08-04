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

// NewPluginsGoldenConfigGoldenConfigSettingsUpdateParams creates a new PluginsGoldenConfigGoldenConfigSettingsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigGoldenConfigSettingsUpdateParams() *PluginsGoldenConfigGoldenConfigSettingsUpdateParams {
	return &PluginsGoldenConfigGoldenConfigSettingsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigGoldenConfigSettingsUpdateParamsWithTimeout creates a new PluginsGoldenConfigGoldenConfigSettingsUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigGoldenConfigSettingsUpdateParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigGoldenConfigSettingsUpdateParams {
	return &PluginsGoldenConfigGoldenConfigSettingsUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigGoldenConfigSettingsUpdateParamsWithContext creates a new PluginsGoldenConfigGoldenConfigSettingsUpdateParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigGoldenConfigSettingsUpdateParamsWithContext(ctx context.Context) *PluginsGoldenConfigGoldenConfigSettingsUpdateParams {
	return &PluginsGoldenConfigGoldenConfigSettingsUpdateParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigGoldenConfigSettingsUpdateParamsWithHTTPClient creates a new PluginsGoldenConfigGoldenConfigSettingsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigGoldenConfigSettingsUpdateParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigGoldenConfigSettingsUpdateParams {
	return &PluginsGoldenConfigGoldenConfigSettingsUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigGoldenConfigSettingsUpdateParams contains all the parameters to send to the API endpoint
   for the plugins golden config golden config settings update operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigGoldenConfigSettingsUpdateParams struct {

	// Data.
	Data *models.GoldenConfigSetting

	/* ID.

	   A UUID string identifying this golden config setting.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config golden config settings update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigGoldenConfigSettingsUpdateParams) WithDefaults() *PluginsGoldenConfigGoldenConfigSettingsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config golden config settings update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigGoldenConfigSettingsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config golden config settings update params
func (o *PluginsGoldenConfigGoldenConfigSettingsUpdateParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigGoldenConfigSettingsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config golden config settings update params
func (o *PluginsGoldenConfigGoldenConfigSettingsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config golden config settings update params
func (o *PluginsGoldenConfigGoldenConfigSettingsUpdateParams) WithContext(ctx context.Context) *PluginsGoldenConfigGoldenConfigSettingsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config golden config settings update params
func (o *PluginsGoldenConfigGoldenConfigSettingsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config golden config settings update params
func (o *PluginsGoldenConfigGoldenConfigSettingsUpdateParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigGoldenConfigSettingsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config golden config settings update params
func (o *PluginsGoldenConfigGoldenConfigSettingsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins golden config golden config settings update params
func (o *PluginsGoldenConfigGoldenConfigSettingsUpdateParams) WithData(data *models.GoldenConfigSetting) *PluginsGoldenConfigGoldenConfigSettingsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins golden config golden config settings update params
func (o *PluginsGoldenConfigGoldenConfigSettingsUpdateParams) SetData(data *models.GoldenConfigSetting) {
	o.Data = data
}

// WithID adds the id to the plugins golden config golden config settings update params
func (o *PluginsGoldenConfigGoldenConfigSettingsUpdateParams) WithID(id strfmt.UUID) *PluginsGoldenConfigGoldenConfigSettingsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins golden config golden config settings update params
func (o *PluginsGoldenConfigGoldenConfigSettingsUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigGoldenConfigSettingsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

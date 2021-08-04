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

// NewPluginsGoldenConfigGoldenConfigSettingsPartialUpdateParams creates a new PluginsGoldenConfigGoldenConfigSettingsPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigGoldenConfigSettingsPartialUpdateParams() *PluginsGoldenConfigGoldenConfigSettingsPartialUpdateParams {
	return &PluginsGoldenConfigGoldenConfigSettingsPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigGoldenConfigSettingsPartialUpdateParamsWithTimeout creates a new PluginsGoldenConfigGoldenConfigSettingsPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigGoldenConfigSettingsPartialUpdateParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigGoldenConfigSettingsPartialUpdateParams {
	return &PluginsGoldenConfigGoldenConfigSettingsPartialUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigGoldenConfigSettingsPartialUpdateParamsWithContext creates a new PluginsGoldenConfigGoldenConfigSettingsPartialUpdateParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigGoldenConfigSettingsPartialUpdateParamsWithContext(ctx context.Context) *PluginsGoldenConfigGoldenConfigSettingsPartialUpdateParams {
	return &PluginsGoldenConfigGoldenConfigSettingsPartialUpdateParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigGoldenConfigSettingsPartialUpdateParamsWithHTTPClient creates a new PluginsGoldenConfigGoldenConfigSettingsPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigGoldenConfigSettingsPartialUpdateParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigGoldenConfigSettingsPartialUpdateParams {
	return &PluginsGoldenConfigGoldenConfigSettingsPartialUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigGoldenConfigSettingsPartialUpdateParams contains all the parameters to send to the API endpoint
   for the plugins golden config golden config settings partial update operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigGoldenConfigSettingsPartialUpdateParams struct {

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

// WithDefaults hydrates default values in the plugins golden config golden config settings partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigGoldenConfigSettingsPartialUpdateParams) WithDefaults() *PluginsGoldenConfigGoldenConfigSettingsPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config golden config settings partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigGoldenConfigSettingsPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config golden config settings partial update params
func (o *PluginsGoldenConfigGoldenConfigSettingsPartialUpdateParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigGoldenConfigSettingsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config golden config settings partial update params
func (o *PluginsGoldenConfigGoldenConfigSettingsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config golden config settings partial update params
func (o *PluginsGoldenConfigGoldenConfigSettingsPartialUpdateParams) WithContext(ctx context.Context) *PluginsGoldenConfigGoldenConfigSettingsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config golden config settings partial update params
func (o *PluginsGoldenConfigGoldenConfigSettingsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config golden config settings partial update params
func (o *PluginsGoldenConfigGoldenConfigSettingsPartialUpdateParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigGoldenConfigSettingsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config golden config settings partial update params
func (o *PluginsGoldenConfigGoldenConfigSettingsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins golden config golden config settings partial update params
func (o *PluginsGoldenConfigGoldenConfigSettingsPartialUpdateParams) WithData(data *models.GoldenConfigSetting) *PluginsGoldenConfigGoldenConfigSettingsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins golden config golden config settings partial update params
func (o *PluginsGoldenConfigGoldenConfigSettingsPartialUpdateParams) SetData(data *models.GoldenConfigSetting) {
	o.Data = data
}

// WithID adds the id to the plugins golden config golden config settings partial update params
func (o *PluginsGoldenConfigGoldenConfigSettingsPartialUpdateParams) WithID(id strfmt.UUID) *PluginsGoldenConfigGoldenConfigSettingsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins golden config golden config settings partial update params
func (o *PluginsGoldenConfigGoldenConfigSettingsPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigGoldenConfigSettingsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

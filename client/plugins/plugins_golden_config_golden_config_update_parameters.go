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

// NewPluginsGoldenConfigGoldenConfigUpdateParams creates a new PluginsGoldenConfigGoldenConfigUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigGoldenConfigUpdateParams() *PluginsGoldenConfigGoldenConfigUpdateParams {
	return &PluginsGoldenConfigGoldenConfigUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigGoldenConfigUpdateParamsWithTimeout creates a new PluginsGoldenConfigGoldenConfigUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigGoldenConfigUpdateParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigGoldenConfigUpdateParams {
	return &PluginsGoldenConfigGoldenConfigUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigGoldenConfigUpdateParamsWithContext creates a new PluginsGoldenConfigGoldenConfigUpdateParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigGoldenConfigUpdateParamsWithContext(ctx context.Context) *PluginsGoldenConfigGoldenConfigUpdateParams {
	return &PluginsGoldenConfigGoldenConfigUpdateParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigGoldenConfigUpdateParamsWithHTTPClient creates a new PluginsGoldenConfigGoldenConfigUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigGoldenConfigUpdateParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigGoldenConfigUpdateParams {
	return &PluginsGoldenConfigGoldenConfigUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigGoldenConfigUpdateParams contains all the parameters to send to the API endpoint
   for the plugins golden config golden config update operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigGoldenConfigUpdateParams struct {

	// Data.
	Data *models.GoldenConfig

	/* ID.

	   A UUID string identifying this golden config.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config golden config update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigGoldenConfigUpdateParams) WithDefaults() *PluginsGoldenConfigGoldenConfigUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config golden config update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigGoldenConfigUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config golden config update params
func (o *PluginsGoldenConfigGoldenConfigUpdateParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigGoldenConfigUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config golden config update params
func (o *PluginsGoldenConfigGoldenConfigUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config golden config update params
func (o *PluginsGoldenConfigGoldenConfigUpdateParams) WithContext(ctx context.Context) *PluginsGoldenConfigGoldenConfigUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config golden config update params
func (o *PluginsGoldenConfigGoldenConfigUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config golden config update params
func (o *PluginsGoldenConfigGoldenConfigUpdateParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigGoldenConfigUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config golden config update params
func (o *PluginsGoldenConfigGoldenConfigUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins golden config golden config update params
func (o *PluginsGoldenConfigGoldenConfigUpdateParams) WithData(data *models.GoldenConfig) *PluginsGoldenConfigGoldenConfigUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins golden config golden config update params
func (o *PluginsGoldenConfigGoldenConfigUpdateParams) SetData(data *models.GoldenConfig) {
	o.Data = data
}

// WithID adds the id to the plugins golden config golden config update params
func (o *PluginsGoldenConfigGoldenConfigUpdateParams) WithID(id strfmt.UUID) *PluginsGoldenConfigGoldenConfigUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins golden config golden config update params
func (o *PluginsGoldenConfigGoldenConfigUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigGoldenConfigUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

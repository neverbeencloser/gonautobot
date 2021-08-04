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

// NewPluginsGoldenConfigGoldenConfigPartialUpdateParams creates a new PluginsGoldenConfigGoldenConfigPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigGoldenConfigPartialUpdateParams() *PluginsGoldenConfigGoldenConfigPartialUpdateParams {
	return &PluginsGoldenConfigGoldenConfigPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigGoldenConfigPartialUpdateParamsWithTimeout creates a new PluginsGoldenConfigGoldenConfigPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigGoldenConfigPartialUpdateParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigGoldenConfigPartialUpdateParams {
	return &PluginsGoldenConfigGoldenConfigPartialUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigGoldenConfigPartialUpdateParamsWithContext creates a new PluginsGoldenConfigGoldenConfigPartialUpdateParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigGoldenConfigPartialUpdateParamsWithContext(ctx context.Context) *PluginsGoldenConfigGoldenConfigPartialUpdateParams {
	return &PluginsGoldenConfigGoldenConfigPartialUpdateParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigGoldenConfigPartialUpdateParamsWithHTTPClient creates a new PluginsGoldenConfigGoldenConfigPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigGoldenConfigPartialUpdateParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigGoldenConfigPartialUpdateParams {
	return &PluginsGoldenConfigGoldenConfigPartialUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigGoldenConfigPartialUpdateParams contains all the parameters to send to the API endpoint
   for the plugins golden config golden config partial update operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigGoldenConfigPartialUpdateParams struct {

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

// WithDefaults hydrates default values in the plugins golden config golden config partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigGoldenConfigPartialUpdateParams) WithDefaults() *PluginsGoldenConfigGoldenConfigPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config golden config partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigGoldenConfigPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config golden config partial update params
func (o *PluginsGoldenConfigGoldenConfigPartialUpdateParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigGoldenConfigPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config golden config partial update params
func (o *PluginsGoldenConfigGoldenConfigPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config golden config partial update params
func (o *PluginsGoldenConfigGoldenConfigPartialUpdateParams) WithContext(ctx context.Context) *PluginsGoldenConfigGoldenConfigPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config golden config partial update params
func (o *PluginsGoldenConfigGoldenConfigPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config golden config partial update params
func (o *PluginsGoldenConfigGoldenConfigPartialUpdateParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigGoldenConfigPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config golden config partial update params
func (o *PluginsGoldenConfigGoldenConfigPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins golden config golden config partial update params
func (o *PluginsGoldenConfigGoldenConfigPartialUpdateParams) WithData(data *models.GoldenConfig) *PluginsGoldenConfigGoldenConfigPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins golden config golden config partial update params
func (o *PluginsGoldenConfigGoldenConfigPartialUpdateParams) SetData(data *models.GoldenConfig) {
	o.Data = data
}

// WithID adds the id to the plugins golden config golden config partial update params
func (o *PluginsGoldenConfigGoldenConfigPartialUpdateParams) WithID(id strfmt.UUID) *PluginsGoldenConfigGoldenConfigPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins golden config golden config partial update params
func (o *PluginsGoldenConfigGoldenConfigPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigGoldenConfigPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

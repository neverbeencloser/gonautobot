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

// NewPluginsGoldenConfigConfigRemovePartialUpdateParams creates a new PluginsGoldenConfigConfigRemovePartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigConfigRemovePartialUpdateParams() *PluginsGoldenConfigConfigRemovePartialUpdateParams {
	return &PluginsGoldenConfigConfigRemovePartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigConfigRemovePartialUpdateParamsWithTimeout creates a new PluginsGoldenConfigConfigRemovePartialUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigConfigRemovePartialUpdateParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigConfigRemovePartialUpdateParams {
	return &PluginsGoldenConfigConfigRemovePartialUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigConfigRemovePartialUpdateParamsWithContext creates a new PluginsGoldenConfigConfigRemovePartialUpdateParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigConfigRemovePartialUpdateParamsWithContext(ctx context.Context) *PluginsGoldenConfigConfigRemovePartialUpdateParams {
	return &PluginsGoldenConfigConfigRemovePartialUpdateParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigConfigRemovePartialUpdateParamsWithHTTPClient creates a new PluginsGoldenConfigConfigRemovePartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigConfigRemovePartialUpdateParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigConfigRemovePartialUpdateParams {
	return &PluginsGoldenConfigConfigRemovePartialUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigConfigRemovePartialUpdateParams contains all the parameters to send to the API endpoint
   for the plugins golden config config remove partial update operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigConfigRemovePartialUpdateParams struct {

	// Data.
	Data *models.ConfigRemove

	/* ID.

	   A UUID string identifying this config remove.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config config remove partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigConfigRemovePartialUpdateParams) WithDefaults() *PluginsGoldenConfigConfigRemovePartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config config remove partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigConfigRemovePartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config config remove partial update params
func (o *PluginsGoldenConfigConfigRemovePartialUpdateParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigConfigRemovePartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config config remove partial update params
func (o *PluginsGoldenConfigConfigRemovePartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config config remove partial update params
func (o *PluginsGoldenConfigConfigRemovePartialUpdateParams) WithContext(ctx context.Context) *PluginsGoldenConfigConfigRemovePartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config config remove partial update params
func (o *PluginsGoldenConfigConfigRemovePartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config config remove partial update params
func (o *PluginsGoldenConfigConfigRemovePartialUpdateParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigConfigRemovePartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config config remove partial update params
func (o *PluginsGoldenConfigConfigRemovePartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins golden config config remove partial update params
func (o *PluginsGoldenConfigConfigRemovePartialUpdateParams) WithData(data *models.ConfigRemove) *PluginsGoldenConfigConfigRemovePartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins golden config config remove partial update params
func (o *PluginsGoldenConfigConfigRemovePartialUpdateParams) SetData(data *models.ConfigRemove) {
	o.Data = data
}

// WithID adds the id to the plugins golden config config remove partial update params
func (o *PluginsGoldenConfigConfigRemovePartialUpdateParams) WithID(id strfmt.UUID) *PluginsGoldenConfigConfigRemovePartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins golden config config remove partial update params
func (o *PluginsGoldenConfigConfigRemovePartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigConfigRemovePartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

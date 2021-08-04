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

// NewPluginsGoldenConfigConfigRemoveUpdateParams creates a new PluginsGoldenConfigConfigRemoveUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigConfigRemoveUpdateParams() *PluginsGoldenConfigConfigRemoveUpdateParams {
	return &PluginsGoldenConfigConfigRemoveUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigConfigRemoveUpdateParamsWithTimeout creates a new PluginsGoldenConfigConfigRemoveUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigConfigRemoveUpdateParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigConfigRemoveUpdateParams {
	return &PluginsGoldenConfigConfigRemoveUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigConfigRemoveUpdateParamsWithContext creates a new PluginsGoldenConfigConfigRemoveUpdateParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigConfigRemoveUpdateParamsWithContext(ctx context.Context) *PluginsGoldenConfigConfigRemoveUpdateParams {
	return &PluginsGoldenConfigConfigRemoveUpdateParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigConfigRemoveUpdateParamsWithHTTPClient creates a new PluginsGoldenConfigConfigRemoveUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigConfigRemoveUpdateParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigConfigRemoveUpdateParams {
	return &PluginsGoldenConfigConfigRemoveUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigConfigRemoveUpdateParams contains all the parameters to send to the API endpoint
   for the plugins golden config config remove update operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigConfigRemoveUpdateParams struct {

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

// WithDefaults hydrates default values in the plugins golden config config remove update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigConfigRemoveUpdateParams) WithDefaults() *PluginsGoldenConfigConfigRemoveUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config config remove update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigConfigRemoveUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config config remove update params
func (o *PluginsGoldenConfigConfigRemoveUpdateParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigConfigRemoveUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config config remove update params
func (o *PluginsGoldenConfigConfigRemoveUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config config remove update params
func (o *PluginsGoldenConfigConfigRemoveUpdateParams) WithContext(ctx context.Context) *PluginsGoldenConfigConfigRemoveUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config config remove update params
func (o *PluginsGoldenConfigConfigRemoveUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config config remove update params
func (o *PluginsGoldenConfigConfigRemoveUpdateParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigConfigRemoveUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config config remove update params
func (o *PluginsGoldenConfigConfigRemoveUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins golden config config remove update params
func (o *PluginsGoldenConfigConfigRemoveUpdateParams) WithData(data *models.ConfigRemove) *PluginsGoldenConfigConfigRemoveUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins golden config config remove update params
func (o *PluginsGoldenConfigConfigRemoveUpdateParams) SetData(data *models.ConfigRemove) {
	o.Data = data
}

// WithID adds the id to the plugins golden config config remove update params
func (o *PluginsGoldenConfigConfigRemoveUpdateParams) WithID(id strfmt.UUID) *PluginsGoldenConfigConfigRemoveUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins golden config config remove update params
func (o *PluginsGoldenConfigConfigRemoveUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigConfigRemoveUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

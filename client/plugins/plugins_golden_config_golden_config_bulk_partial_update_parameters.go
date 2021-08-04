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

// NewPluginsGoldenConfigGoldenConfigBulkPartialUpdateParams creates a new PluginsGoldenConfigGoldenConfigBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigGoldenConfigBulkPartialUpdateParams() *PluginsGoldenConfigGoldenConfigBulkPartialUpdateParams {
	return &PluginsGoldenConfigGoldenConfigBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigGoldenConfigBulkPartialUpdateParamsWithTimeout creates a new PluginsGoldenConfigGoldenConfigBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigGoldenConfigBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigGoldenConfigBulkPartialUpdateParams {
	return &PluginsGoldenConfigGoldenConfigBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigGoldenConfigBulkPartialUpdateParamsWithContext creates a new PluginsGoldenConfigGoldenConfigBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigGoldenConfigBulkPartialUpdateParamsWithContext(ctx context.Context) *PluginsGoldenConfigGoldenConfigBulkPartialUpdateParams {
	return &PluginsGoldenConfigGoldenConfigBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigGoldenConfigBulkPartialUpdateParamsWithHTTPClient creates a new PluginsGoldenConfigGoldenConfigBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigGoldenConfigBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigGoldenConfigBulkPartialUpdateParams {
	return &PluginsGoldenConfigGoldenConfigBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigGoldenConfigBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the plugins golden config golden config bulk partial update operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigGoldenConfigBulkPartialUpdateParams struct {

	// Data.
	Data *models.GoldenConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config golden config bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigGoldenConfigBulkPartialUpdateParams) WithDefaults() *PluginsGoldenConfigGoldenConfigBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config golden config bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigGoldenConfigBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config golden config bulk partial update params
func (o *PluginsGoldenConfigGoldenConfigBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigGoldenConfigBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config golden config bulk partial update params
func (o *PluginsGoldenConfigGoldenConfigBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config golden config bulk partial update params
func (o *PluginsGoldenConfigGoldenConfigBulkPartialUpdateParams) WithContext(ctx context.Context) *PluginsGoldenConfigGoldenConfigBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config golden config bulk partial update params
func (o *PluginsGoldenConfigGoldenConfigBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config golden config bulk partial update params
func (o *PluginsGoldenConfigGoldenConfigBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigGoldenConfigBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config golden config bulk partial update params
func (o *PluginsGoldenConfigGoldenConfigBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins golden config golden config bulk partial update params
func (o *PluginsGoldenConfigGoldenConfigBulkPartialUpdateParams) WithData(data *models.GoldenConfig) *PluginsGoldenConfigGoldenConfigBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins golden config golden config bulk partial update params
func (o *PluginsGoldenConfigGoldenConfigBulkPartialUpdateParams) SetData(data *models.GoldenConfig) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigGoldenConfigBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

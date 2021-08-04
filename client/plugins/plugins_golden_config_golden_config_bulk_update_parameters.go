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

// NewPluginsGoldenConfigGoldenConfigBulkUpdateParams creates a new PluginsGoldenConfigGoldenConfigBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigGoldenConfigBulkUpdateParams() *PluginsGoldenConfigGoldenConfigBulkUpdateParams {
	return &PluginsGoldenConfigGoldenConfigBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigGoldenConfigBulkUpdateParamsWithTimeout creates a new PluginsGoldenConfigGoldenConfigBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigGoldenConfigBulkUpdateParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigGoldenConfigBulkUpdateParams {
	return &PluginsGoldenConfigGoldenConfigBulkUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigGoldenConfigBulkUpdateParamsWithContext creates a new PluginsGoldenConfigGoldenConfigBulkUpdateParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigGoldenConfigBulkUpdateParamsWithContext(ctx context.Context) *PluginsGoldenConfigGoldenConfigBulkUpdateParams {
	return &PluginsGoldenConfigGoldenConfigBulkUpdateParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigGoldenConfigBulkUpdateParamsWithHTTPClient creates a new PluginsGoldenConfigGoldenConfigBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigGoldenConfigBulkUpdateParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigGoldenConfigBulkUpdateParams {
	return &PluginsGoldenConfigGoldenConfigBulkUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigGoldenConfigBulkUpdateParams contains all the parameters to send to the API endpoint
   for the plugins golden config golden config bulk update operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigGoldenConfigBulkUpdateParams struct {

	// Data.
	Data *models.GoldenConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config golden config bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigGoldenConfigBulkUpdateParams) WithDefaults() *PluginsGoldenConfigGoldenConfigBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config golden config bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigGoldenConfigBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config golden config bulk update params
func (o *PluginsGoldenConfigGoldenConfigBulkUpdateParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigGoldenConfigBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config golden config bulk update params
func (o *PluginsGoldenConfigGoldenConfigBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config golden config bulk update params
func (o *PluginsGoldenConfigGoldenConfigBulkUpdateParams) WithContext(ctx context.Context) *PluginsGoldenConfigGoldenConfigBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config golden config bulk update params
func (o *PluginsGoldenConfigGoldenConfigBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config golden config bulk update params
func (o *PluginsGoldenConfigGoldenConfigBulkUpdateParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigGoldenConfigBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config golden config bulk update params
func (o *PluginsGoldenConfigGoldenConfigBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins golden config golden config bulk update params
func (o *PluginsGoldenConfigGoldenConfigBulkUpdateParams) WithData(data *models.GoldenConfig) *PluginsGoldenConfigGoldenConfigBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins golden config golden config bulk update params
func (o *PluginsGoldenConfigGoldenConfigBulkUpdateParams) SetData(data *models.GoldenConfig) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigGoldenConfigBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

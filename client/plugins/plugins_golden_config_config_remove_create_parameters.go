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

// NewPluginsGoldenConfigConfigRemoveCreateParams creates a new PluginsGoldenConfigConfigRemoveCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigConfigRemoveCreateParams() *PluginsGoldenConfigConfigRemoveCreateParams {
	return &PluginsGoldenConfigConfigRemoveCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigConfigRemoveCreateParamsWithTimeout creates a new PluginsGoldenConfigConfigRemoveCreateParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigConfigRemoveCreateParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigConfigRemoveCreateParams {
	return &PluginsGoldenConfigConfigRemoveCreateParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigConfigRemoveCreateParamsWithContext creates a new PluginsGoldenConfigConfigRemoveCreateParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigConfigRemoveCreateParamsWithContext(ctx context.Context) *PluginsGoldenConfigConfigRemoveCreateParams {
	return &PluginsGoldenConfigConfigRemoveCreateParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigConfigRemoveCreateParamsWithHTTPClient creates a new PluginsGoldenConfigConfigRemoveCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigConfigRemoveCreateParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigConfigRemoveCreateParams {
	return &PluginsGoldenConfigConfigRemoveCreateParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigConfigRemoveCreateParams contains all the parameters to send to the API endpoint
   for the plugins golden config config remove create operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigConfigRemoveCreateParams struct {

	// Data.
	Data *models.ConfigRemove

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config config remove create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigConfigRemoveCreateParams) WithDefaults() *PluginsGoldenConfigConfigRemoveCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config config remove create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigConfigRemoveCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config config remove create params
func (o *PluginsGoldenConfigConfigRemoveCreateParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigConfigRemoveCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config config remove create params
func (o *PluginsGoldenConfigConfigRemoveCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config config remove create params
func (o *PluginsGoldenConfigConfigRemoveCreateParams) WithContext(ctx context.Context) *PluginsGoldenConfigConfigRemoveCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config config remove create params
func (o *PluginsGoldenConfigConfigRemoveCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config config remove create params
func (o *PluginsGoldenConfigConfigRemoveCreateParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigConfigRemoveCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config config remove create params
func (o *PluginsGoldenConfigConfigRemoveCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins golden config config remove create params
func (o *PluginsGoldenConfigConfigRemoveCreateParams) WithData(data *models.ConfigRemove) *PluginsGoldenConfigConfigRemoveCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins golden config config remove create params
func (o *PluginsGoldenConfigConfigRemoveCreateParams) SetData(data *models.ConfigRemove) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigConfigRemoveCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

package plugins

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewPluginsGoldenConfigConfigRemoveDeleteParams creates a new PluginsGoldenConfigConfigRemoveDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigConfigRemoveDeleteParams() *PluginsGoldenConfigConfigRemoveDeleteParams {
	return &PluginsGoldenConfigConfigRemoveDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigConfigRemoveDeleteParamsWithTimeout creates a new PluginsGoldenConfigConfigRemoveDeleteParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigConfigRemoveDeleteParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigConfigRemoveDeleteParams {
	return &PluginsGoldenConfigConfigRemoveDeleteParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigConfigRemoveDeleteParamsWithContext creates a new PluginsGoldenConfigConfigRemoveDeleteParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigConfigRemoveDeleteParamsWithContext(ctx context.Context) *PluginsGoldenConfigConfigRemoveDeleteParams {
	return &PluginsGoldenConfigConfigRemoveDeleteParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigConfigRemoveDeleteParamsWithHTTPClient creates a new PluginsGoldenConfigConfigRemoveDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigConfigRemoveDeleteParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigConfigRemoveDeleteParams {
	return &PluginsGoldenConfigConfigRemoveDeleteParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigConfigRemoveDeleteParams contains all the parameters to send to the API endpoint
   for the plugins golden config config remove delete operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigConfigRemoveDeleteParams struct {

	/* ID.

	   A UUID string identifying this config remove.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config config remove delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigConfigRemoveDeleteParams) WithDefaults() *PluginsGoldenConfigConfigRemoveDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config config remove delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigConfigRemoveDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config config remove delete params
func (o *PluginsGoldenConfigConfigRemoveDeleteParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigConfigRemoveDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config config remove delete params
func (o *PluginsGoldenConfigConfigRemoveDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config config remove delete params
func (o *PluginsGoldenConfigConfigRemoveDeleteParams) WithContext(ctx context.Context) *PluginsGoldenConfigConfigRemoveDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config config remove delete params
func (o *PluginsGoldenConfigConfigRemoveDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config config remove delete params
func (o *PluginsGoldenConfigConfigRemoveDeleteParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigConfigRemoveDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config config remove delete params
func (o *PluginsGoldenConfigConfigRemoveDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the plugins golden config config remove delete params
func (o *PluginsGoldenConfigConfigRemoveDeleteParams) WithID(id strfmt.UUID) *PluginsGoldenConfigConfigRemoveDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins golden config config remove delete params
func (o *PluginsGoldenConfigConfigRemoveDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigConfigRemoveDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

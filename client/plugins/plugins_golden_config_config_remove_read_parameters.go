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

// NewPluginsGoldenConfigConfigRemoveReadParams creates a new PluginsGoldenConfigConfigRemoveReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigConfigRemoveReadParams() *PluginsGoldenConfigConfigRemoveReadParams {
	return &PluginsGoldenConfigConfigRemoveReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigConfigRemoveReadParamsWithTimeout creates a new PluginsGoldenConfigConfigRemoveReadParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigConfigRemoveReadParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigConfigRemoveReadParams {
	return &PluginsGoldenConfigConfigRemoveReadParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigConfigRemoveReadParamsWithContext creates a new PluginsGoldenConfigConfigRemoveReadParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigConfigRemoveReadParamsWithContext(ctx context.Context) *PluginsGoldenConfigConfigRemoveReadParams {
	return &PluginsGoldenConfigConfigRemoveReadParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigConfigRemoveReadParamsWithHTTPClient creates a new PluginsGoldenConfigConfigRemoveReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigConfigRemoveReadParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigConfigRemoveReadParams {
	return &PluginsGoldenConfigConfigRemoveReadParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigConfigRemoveReadParams contains all the parameters to send to the API endpoint
   for the plugins golden config config remove read operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigConfigRemoveReadParams struct {

	/* ID.

	   A UUID string identifying this config remove.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config config remove read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigConfigRemoveReadParams) WithDefaults() *PluginsGoldenConfigConfigRemoveReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config config remove read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigConfigRemoveReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config config remove read params
func (o *PluginsGoldenConfigConfigRemoveReadParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigConfigRemoveReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config config remove read params
func (o *PluginsGoldenConfigConfigRemoveReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config config remove read params
func (o *PluginsGoldenConfigConfigRemoveReadParams) WithContext(ctx context.Context) *PluginsGoldenConfigConfigRemoveReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config config remove read params
func (o *PluginsGoldenConfigConfigRemoveReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config config remove read params
func (o *PluginsGoldenConfigConfigRemoveReadParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigConfigRemoveReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config config remove read params
func (o *PluginsGoldenConfigConfigRemoveReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the plugins golden config config remove read params
func (o *PluginsGoldenConfigConfigRemoveReadParams) WithID(id strfmt.UUID) *PluginsGoldenConfigConfigRemoveReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins golden config config remove read params
func (o *PluginsGoldenConfigConfigRemoveReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigConfigRemoveReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

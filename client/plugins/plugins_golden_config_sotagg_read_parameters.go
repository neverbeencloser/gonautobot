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

// NewPluginsGoldenConfigSotaggReadParams creates a new PluginsGoldenConfigSotaggReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigSotaggReadParams() *PluginsGoldenConfigSotaggReadParams {
	return &PluginsGoldenConfigSotaggReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigSotaggReadParamsWithTimeout creates a new PluginsGoldenConfigSotaggReadParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigSotaggReadParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigSotaggReadParams {
	return &PluginsGoldenConfigSotaggReadParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigSotaggReadParamsWithContext creates a new PluginsGoldenConfigSotaggReadParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigSotaggReadParamsWithContext(ctx context.Context) *PluginsGoldenConfigSotaggReadParams {
	return &PluginsGoldenConfigSotaggReadParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigSotaggReadParamsWithHTTPClient creates a new PluginsGoldenConfigSotaggReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigSotaggReadParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigSotaggReadParams {
	return &PluginsGoldenConfigSotaggReadParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigSotaggReadParams contains all the parameters to send to the API endpoint
   for the plugins golden config sotagg read operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigSotaggReadParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config sotagg read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigSotaggReadParams) WithDefaults() *PluginsGoldenConfigSotaggReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config sotagg read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigSotaggReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config sotagg read params
func (o *PluginsGoldenConfigSotaggReadParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigSotaggReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config sotagg read params
func (o *PluginsGoldenConfigSotaggReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config sotagg read params
func (o *PluginsGoldenConfigSotaggReadParams) WithContext(ctx context.Context) *PluginsGoldenConfigSotaggReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config sotagg read params
func (o *PluginsGoldenConfigSotaggReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config sotagg read params
func (o *PluginsGoldenConfigSotaggReadParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigSotaggReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config sotagg read params
func (o *PluginsGoldenConfigSotaggReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the plugins golden config sotagg read params
func (o *PluginsGoldenConfigSotaggReadParams) WithID(id string) *PluginsGoldenConfigSotaggReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins golden config sotagg read params
func (o *PluginsGoldenConfigSotaggReadParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigSotaggReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

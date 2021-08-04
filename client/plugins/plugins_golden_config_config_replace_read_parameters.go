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

// NewPluginsGoldenConfigConfigReplaceReadParams creates a new PluginsGoldenConfigConfigReplaceReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigConfigReplaceReadParams() *PluginsGoldenConfigConfigReplaceReadParams {
	return &PluginsGoldenConfigConfigReplaceReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigConfigReplaceReadParamsWithTimeout creates a new PluginsGoldenConfigConfigReplaceReadParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigConfigReplaceReadParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigConfigReplaceReadParams {
	return &PluginsGoldenConfigConfigReplaceReadParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigConfigReplaceReadParamsWithContext creates a new PluginsGoldenConfigConfigReplaceReadParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigConfigReplaceReadParamsWithContext(ctx context.Context) *PluginsGoldenConfigConfigReplaceReadParams {
	return &PluginsGoldenConfigConfigReplaceReadParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigConfigReplaceReadParamsWithHTTPClient creates a new PluginsGoldenConfigConfigReplaceReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigConfigReplaceReadParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigConfigReplaceReadParams {
	return &PluginsGoldenConfigConfigReplaceReadParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigConfigReplaceReadParams contains all the parameters to send to the API endpoint
   for the plugins golden config config replace read operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigConfigReplaceReadParams struct {

	/* ID.

	   A UUID string identifying this config replace.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config config replace read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigConfigReplaceReadParams) WithDefaults() *PluginsGoldenConfigConfigReplaceReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config config replace read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigConfigReplaceReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config config replace read params
func (o *PluginsGoldenConfigConfigReplaceReadParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigConfigReplaceReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config config replace read params
func (o *PluginsGoldenConfigConfigReplaceReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config config replace read params
func (o *PluginsGoldenConfigConfigReplaceReadParams) WithContext(ctx context.Context) *PluginsGoldenConfigConfigReplaceReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config config replace read params
func (o *PluginsGoldenConfigConfigReplaceReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config config replace read params
func (o *PluginsGoldenConfigConfigReplaceReadParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigConfigReplaceReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config config replace read params
func (o *PluginsGoldenConfigConfigReplaceReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the plugins golden config config replace read params
func (o *PluginsGoldenConfigConfigReplaceReadParams) WithID(id strfmt.UUID) *PluginsGoldenConfigConfigReplaceReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins golden config config replace read params
func (o *PluginsGoldenConfigConfigReplaceReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigConfigReplaceReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

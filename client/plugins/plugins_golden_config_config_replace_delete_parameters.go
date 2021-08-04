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

// NewPluginsGoldenConfigConfigReplaceDeleteParams creates a new PluginsGoldenConfigConfigReplaceDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigConfigReplaceDeleteParams() *PluginsGoldenConfigConfigReplaceDeleteParams {
	return &PluginsGoldenConfigConfigReplaceDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigConfigReplaceDeleteParamsWithTimeout creates a new PluginsGoldenConfigConfigReplaceDeleteParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigConfigReplaceDeleteParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigConfigReplaceDeleteParams {
	return &PluginsGoldenConfigConfigReplaceDeleteParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigConfigReplaceDeleteParamsWithContext creates a new PluginsGoldenConfigConfigReplaceDeleteParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigConfigReplaceDeleteParamsWithContext(ctx context.Context) *PluginsGoldenConfigConfigReplaceDeleteParams {
	return &PluginsGoldenConfigConfigReplaceDeleteParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigConfigReplaceDeleteParamsWithHTTPClient creates a new PluginsGoldenConfigConfigReplaceDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigConfigReplaceDeleteParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigConfigReplaceDeleteParams {
	return &PluginsGoldenConfigConfigReplaceDeleteParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigConfigReplaceDeleteParams contains all the parameters to send to the API endpoint
   for the plugins golden config config replace delete operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigConfigReplaceDeleteParams struct {

	/* ID.

	   A UUID string identifying this config replace.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config config replace delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigConfigReplaceDeleteParams) WithDefaults() *PluginsGoldenConfigConfigReplaceDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config config replace delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigConfigReplaceDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config config replace delete params
func (o *PluginsGoldenConfigConfigReplaceDeleteParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigConfigReplaceDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config config replace delete params
func (o *PluginsGoldenConfigConfigReplaceDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config config replace delete params
func (o *PluginsGoldenConfigConfigReplaceDeleteParams) WithContext(ctx context.Context) *PluginsGoldenConfigConfigReplaceDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config config replace delete params
func (o *PluginsGoldenConfigConfigReplaceDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config config replace delete params
func (o *PluginsGoldenConfigConfigReplaceDeleteParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigConfigReplaceDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config config replace delete params
func (o *PluginsGoldenConfigConfigReplaceDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the plugins golden config config replace delete params
func (o *PluginsGoldenConfigConfigReplaceDeleteParams) WithID(id strfmt.UUID) *PluginsGoldenConfigConfigReplaceDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins golden config config replace delete params
func (o *PluginsGoldenConfigConfigReplaceDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigConfigReplaceDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

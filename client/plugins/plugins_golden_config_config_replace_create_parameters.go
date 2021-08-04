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

// NewPluginsGoldenConfigConfigReplaceCreateParams creates a new PluginsGoldenConfigConfigReplaceCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigConfigReplaceCreateParams() *PluginsGoldenConfigConfigReplaceCreateParams {
	return &PluginsGoldenConfigConfigReplaceCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigConfigReplaceCreateParamsWithTimeout creates a new PluginsGoldenConfigConfigReplaceCreateParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigConfigReplaceCreateParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigConfigReplaceCreateParams {
	return &PluginsGoldenConfigConfigReplaceCreateParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigConfigReplaceCreateParamsWithContext creates a new PluginsGoldenConfigConfigReplaceCreateParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigConfigReplaceCreateParamsWithContext(ctx context.Context) *PluginsGoldenConfigConfigReplaceCreateParams {
	return &PluginsGoldenConfigConfigReplaceCreateParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigConfigReplaceCreateParamsWithHTTPClient creates a new PluginsGoldenConfigConfigReplaceCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigConfigReplaceCreateParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigConfigReplaceCreateParams {
	return &PluginsGoldenConfigConfigReplaceCreateParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigConfigReplaceCreateParams contains all the parameters to send to the API endpoint
   for the plugins golden config config replace create operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigConfigReplaceCreateParams struct {

	// Data.
	Data *models.ConfigReplace

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config config replace create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigConfigReplaceCreateParams) WithDefaults() *PluginsGoldenConfigConfigReplaceCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config config replace create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigConfigReplaceCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config config replace create params
func (o *PluginsGoldenConfigConfigReplaceCreateParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigConfigReplaceCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config config replace create params
func (o *PluginsGoldenConfigConfigReplaceCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config config replace create params
func (o *PluginsGoldenConfigConfigReplaceCreateParams) WithContext(ctx context.Context) *PluginsGoldenConfigConfigReplaceCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config config replace create params
func (o *PluginsGoldenConfigConfigReplaceCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config config replace create params
func (o *PluginsGoldenConfigConfigReplaceCreateParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigConfigReplaceCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config config replace create params
func (o *PluginsGoldenConfigConfigReplaceCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins golden config config replace create params
func (o *PluginsGoldenConfigConfigReplaceCreateParams) WithData(data *models.ConfigReplace) *PluginsGoldenConfigConfigReplaceCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins golden config config replace create params
func (o *PluginsGoldenConfigConfigReplaceCreateParams) SetData(data *models.ConfigReplace) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigConfigReplaceCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

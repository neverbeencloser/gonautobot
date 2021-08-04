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

// NewPluginsGoldenConfigConfigReplaceBulkPartialUpdateParams creates a new PluginsGoldenConfigConfigReplaceBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigConfigReplaceBulkPartialUpdateParams() *PluginsGoldenConfigConfigReplaceBulkPartialUpdateParams {
	return &PluginsGoldenConfigConfigReplaceBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigConfigReplaceBulkPartialUpdateParamsWithTimeout creates a new PluginsGoldenConfigConfigReplaceBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigConfigReplaceBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigConfigReplaceBulkPartialUpdateParams {
	return &PluginsGoldenConfigConfigReplaceBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigConfigReplaceBulkPartialUpdateParamsWithContext creates a new PluginsGoldenConfigConfigReplaceBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigConfigReplaceBulkPartialUpdateParamsWithContext(ctx context.Context) *PluginsGoldenConfigConfigReplaceBulkPartialUpdateParams {
	return &PluginsGoldenConfigConfigReplaceBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigConfigReplaceBulkPartialUpdateParamsWithHTTPClient creates a new PluginsGoldenConfigConfigReplaceBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigConfigReplaceBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigConfigReplaceBulkPartialUpdateParams {
	return &PluginsGoldenConfigConfigReplaceBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigConfigReplaceBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the plugins golden config config replace bulk partial update operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigConfigReplaceBulkPartialUpdateParams struct {

	// Data.
	Data *models.ConfigReplace

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config config replace bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigConfigReplaceBulkPartialUpdateParams) WithDefaults() *PluginsGoldenConfigConfigReplaceBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config config replace bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigConfigReplaceBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config config replace bulk partial update params
func (o *PluginsGoldenConfigConfigReplaceBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigConfigReplaceBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config config replace bulk partial update params
func (o *PluginsGoldenConfigConfigReplaceBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config config replace bulk partial update params
func (o *PluginsGoldenConfigConfigReplaceBulkPartialUpdateParams) WithContext(ctx context.Context) *PluginsGoldenConfigConfigReplaceBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config config replace bulk partial update params
func (o *PluginsGoldenConfigConfigReplaceBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config config replace bulk partial update params
func (o *PluginsGoldenConfigConfigReplaceBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigConfigReplaceBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config config replace bulk partial update params
func (o *PluginsGoldenConfigConfigReplaceBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins golden config config replace bulk partial update params
func (o *PluginsGoldenConfigConfigReplaceBulkPartialUpdateParams) WithData(data *models.ConfigReplace) *PluginsGoldenConfigConfigReplaceBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins golden config config replace bulk partial update params
func (o *PluginsGoldenConfigConfigReplaceBulkPartialUpdateParams) SetData(data *models.ConfigReplace) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigConfigReplaceBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

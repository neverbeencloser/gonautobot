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

// NewPluginsGoldenConfigConfigReplaceBulkUpdateParams creates a new PluginsGoldenConfigConfigReplaceBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigConfigReplaceBulkUpdateParams() *PluginsGoldenConfigConfigReplaceBulkUpdateParams {
	return &PluginsGoldenConfigConfigReplaceBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigConfigReplaceBulkUpdateParamsWithTimeout creates a new PluginsGoldenConfigConfigReplaceBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigConfigReplaceBulkUpdateParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigConfigReplaceBulkUpdateParams {
	return &PluginsGoldenConfigConfigReplaceBulkUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigConfigReplaceBulkUpdateParamsWithContext creates a new PluginsGoldenConfigConfigReplaceBulkUpdateParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigConfigReplaceBulkUpdateParamsWithContext(ctx context.Context) *PluginsGoldenConfigConfigReplaceBulkUpdateParams {
	return &PluginsGoldenConfigConfigReplaceBulkUpdateParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigConfigReplaceBulkUpdateParamsWithHTTPClient creates a new PluginsGoldenConfigConfigReplaceBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigConfigReplaceBulkUpdateParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigConfigReplaceBulkUpdateParams {
	return &PluginsGoldenConfigConfigReplaceBulkUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigConfigReplaceBulkUpdateParams contains all the parameters to send to the API endpoint
   for the plugins golden config config replace bulk update operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigConfigReplaceBulkUpdateParams struct {

	// Data.
	Data *models.ConfigReplace

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config config replace bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigConfigReplaceBulkUpdateParams) WithDefaults() *PluginsGoldenConfigConfigReplaceBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config config replace bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigConfigReplaceBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config config replace bulk update params
func (o *PluginsGoldenConfigConfigReplaceBulkUpdateParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigConfigReplaceBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config config replace bulk update params
func (o *PluginsGoldenConfigConfigReplaceBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config config replace bulk update params
func (o *PluginsGoldenConfigConfigReplaceBulkUpdateParams) WithContext(ctx context.Context) *PluginsGoldenConfigConfigReplaceBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config config replace bulk update params
func (o *PluginsGoldenConfigConfigReplaceBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config config replace bulk update params
func (o *PluginsGoldenConfigConfigReplaceBulkUpdateParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigConfigReplaceBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config config replace bulk update params
func (o *PluginsGoldenConfigConfigReplaceBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins golden config config replace bulk update params
func (o *PluginsGoldenConfigConfigReplaceBulkUpdateParams) WithData(data *models.ConfigReplace) *PluginsGoldenConfigConfigReplaceBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins golden config config replace bulk update params
func (o *PluginsGoldenConfigConfigReplaceBulkUpdateParams) SetData(data *models.ConfigReplace) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigConfigReplaceBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

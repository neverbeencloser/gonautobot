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

// NewPluginsGoldenConfigConfigRemoveBulkPartialUpdateParams creates a new PluginsGoldenConfigConfigRemoveBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigConfigRemoveBulkPartialUpdateParams() *PluginsGoldenConfigConfigRemoveBulkPartialUpdateParams {
	return &PluginsGoldenConfigConfigRemoveBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigConfigRemoveBulkPartialUpdateParamsWithTimeout creates a new PluginsGoldenConfigConfigRemoveBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigConfigRemoveBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigConfigRemoveBulkPartialUpdateParams {
	return &PluginsGoldenConfigConfigRemoveBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigConfigRemoveBulkPartialUpdateParamsWithContext creates a new PluginsGoldenConfigConfigRemoveBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigConfigRemoveBulkPartialUpdateParamsWithContext(ctx context.Context) *PluginsGoldenConfigConfigRemoveBulkPartialUpdateParams {
	return &PluginsGoldenConfigConfigRemoveBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigConfigRemoveBulkPartialUpdateParamsWithHTTPClient creates a new PluginsGoldenConfigConfigRemoveBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigConfigRemoveBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigConfigRemoveBulkPartialUpdateParams {
	return &PluginsGoldenConfigConfigRemoveBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigConfigRemoveBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the plugins golden config config remove bulk partial update operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigConfigRemoveBulkPartialUpdateParams struct {

	// Data.
	Data *models.ConfigRemove

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config config remove bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigConfigRemoveBulkPartialUpdateParams) WithDefaults() *PluginsGoldenConfigConfigRemoveBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config config remove bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigConfigRemoveBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config config remove bulk partial update params
func (o *PluginsGoldenConfigConfigRemoveBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigConfigRemoveBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config config remove bulk partial update params
func (o *PluginsGoldenConfigConfigRemoveBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config config remove bulk partial update params
func (o *PluginsGoldenConfigConfigRemoveBulkPartialUpdateParams) WithContext(ctx context.Context) *PluginsGoldenConfigConfigRemoveBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config config remove bulk partial update params
func (o *PluginsGoldenConfigConfigRemoveBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config config remove bulk partial update params
func (o *PluginsGoldenConfigConfigRemoveBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigConfigRemoveBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config config remove bulk partial update params
func (o *PluginsGoldenConfigConfigRemoveBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins golden config config remove bulk partial update params
func (o *PluginsGoldenConfigConfigRemoveBulkPartialUpdateParams) WithData(data *models.ConfigRemove) *PluginsGoldenConfigConfigRemoveBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins golden config config remove bulk partial update params
func (o *PluginsGoldenConfigConfigRemoveBulkPartialUpdateParams) SetData(data *models.ConfigRemove) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigConfigRemoveBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

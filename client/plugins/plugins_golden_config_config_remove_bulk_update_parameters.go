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

// NewPluginsGoldenConfigConfigRemoveBulkUpdateParams creates a new PluginsGoldenConfigConfigRemoveBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigConfigRemoveBulkUpdateParams() *PluginsGoldenConfigConfigRemoveBulkUpdateParams {
	return &PluginsGoldenConfigConfigRemoveBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigConfigRemoveBulkUpdateParamsWithTimeout creates a new PluginsGoldenConfigConfigRemoveBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigConfigRemoveBulkUpdateParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigConfigRemoveBulkUpdateParams {
	return &PluginsGoldenConfigConfigRemoveBulkUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigConfigRemoveBulkUpdateParamsWithContext creates a new PluginsGoldenConfigConfigRemoveBulkUpdateParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigConfigRemoveBulkUpdateParamsWithContext(ctx context.Context) *PluginsGoldenConfigConfigRemoveBulkUpdateParams {
	return &PluginsGoldenConfigConfigRemoveBulkUpdateParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigConfigRemoveBulkUpdateParamsWithHTTPClient creates a new PluginsGoldenConfigConfigRemoveBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigConfigRemoveBulkUpdateParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigConfigRemoveBulkUpdateParams {
	return &PluginsGoldenConfigConfigRemoveBulkUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigConfigRemoveBulkUpdateParams contains all the parameters to send to the API endpoint
   for the plugins golden config config remove bulk update operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigConfigRemoveBulkUpdateParams struct {

	// Data.
	Data *models.ConfigRemove

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config config remove bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigConfigRemoveBulkUpdateParams) WithDefaults() *PluginsGoldenConfigConfigRemoveBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config config remove bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigConfigRemoveBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config config remove bulk update params
func (o *PluginsGoldenConfigConfigRemoveBulkUpdateParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigConfigRemoveBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config config remove bulk update params
func (o *PluginsGoldenConfigConfigRemoveBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config config remove bulk update params
func (o *PluginsGoldenConfigConfigRemoveBulkUpdateParams) WithContext(ctx context.Context) *PluginsGoldenConfigConfigRemoveBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config config remove bulk update params
func (o *PluginsGoldenConfigConfigRemoveBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config config remove bulk update params
func (o *PluginsGoldenConfigConfigRemoveBulkUpdateParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigConfigRemoveBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config config remove bulk update params
func (o *PluginsGoldenConfigConfigRemoveBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins golden config config remove bulk update params
func (o *PluginsGoldenConfigConfigRemoveBulkUpdateParams) WithData(data *models.ConfigRemove) *PluginsGoldenConfigConfigRemoveBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins golden config config remove bulk update params
func (o *PluginsGoldenConfigConfigRemoveBulkUpdateParams) SetData(data *models.ConfigRemove) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigConfigRemoveBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

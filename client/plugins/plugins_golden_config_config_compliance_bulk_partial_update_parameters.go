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

// NewPluginsGoldenConfigConfigComplianceBulkPartialUpdateParams creates a new PluginsGoldenConfigConfigComplianceBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigConfigComplianceBulkPartialUpdateParams() *PluginsGoldenConfigConfigComplianceBulkPartialUpdateParams {
	return &PluginsGoldenConfigConfigComplianceBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigConfigComplianceBulkPartialUpdateParamsWithTimeout creates a new PluginsGoldenConfigConfigComplianceBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigConfigComplianceBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigConfigComplianceBulkPartialUpdateParams {
	return &PluginsGoldenConfigConfigComplianceBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigConfigComplianceBulkPartialUpdateParamsWithContext creates a new PluginsGoldenConfigConfigComplianceBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigConfigComplianceBulkPartialUpdateParamsWithContext(ctx context.Context) *PluginsGoldenConfigConfigComplianceBulkPartialUpdateParams {
	return &PluginsGoldenConfigConfigComplianceBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigConfigComplianceBulkPartialUpdateParamsWithHTTPClient creates a new PluginsGoldenConfigConfigComplianceBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigConfigComplianceBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigConfigComplianceBulkPartialUpdateParams {
	return &PluginsGoldenConfigConfigComplianceBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigConfigComplianceBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the plugins golden config config compliance bulk partial update operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigConfigComplianceBulkPartialUpdateParams struct {

	// Data.
	Data *models.ConfigCompliance

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config config compliance bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigConfigComplianceBulkPartialUpdateParams) WithDefaults() *PluginsGoldenConfigConfigComplianceBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config config compliance bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigConfigComplianceBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config config compliance bulk partial update params
func (o *PluginsGoldenConfigConfigComplianceBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigConfigComplianceBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config config compliance bulk partial update params
func (o *PluginsGoldenConfigConfigComplianceBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config config compliance bulk partial update params
func (o *PluginsGoldenConfigConfigComplianceBulkPartialUpdateParams) WithContext(ctx context.Context) *PluginsGoldenConfigConfigComplianceBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config config compliance bulk partial update params
func (o *PluginsGoldenConfigConfigComplianceBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config config compliance bulk partial update params
func (o *PluginsGoldenConfigConfigComplianceBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigConfigComplianceBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config config compliance bulk partial update params
func (o *PluginsGoldenConfigConfigComplianceBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins golden config config compliance bulk partial update params
func (o *PluginsGoldenConfigConfigComplianceBulkPartialUpdateParams) WithData(data *models.ConfigCompliance) *PluginsGoldenConfigConfigComplianceBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins golden config config compliance bulk partial update params
func (o *PluginsGoldenConfigConfigComplianceBulkPartialUpdateParams) SetData(data *models.ConfigCompliance) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigConfigComplianceBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

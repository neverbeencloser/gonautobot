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

// NewPluginsGoldenConfigConfigComplianceBulkUpdateParams creates a new PluginsGoldenConfigConfigComplianceBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigConfigComplianceBulkUpdateParams() *PluginsGoldenConfigConfigComplianceBulkUpdateParams {
	return &PluginsGoldenConfigConfigComplianceBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigConfigComplianceBulkUpdateParamsWithTimeout creates a new PluginsGoldenConfigConfigComplianceBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigConfigComplianceBulkUpdateParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigConfigComplianceBulkUpdateParams {
	return &PluginsGoldenConfigConfigComplianceBulkUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigConfigComplianceBulkUpdateParamsWithContext creates a new PluginsGoldenConfigConfigComplianceBulkUpdateParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigConfigComplianceBulkUpdateParamsWithContext(ctx context.Context) *PluginsGoldenConfigConfigComplianceBulkUpdateParams {
	return &PluginsGoldenConfigConfigComplianceBulkUpdateParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigConfigComplianceBulkUpdateParamsWithHTTPClient creates a new PluginsGoldenConfigConfigComplianceBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigConfigComplianceBulkUpdateParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigConfigComplianceBulkUpdateParams {
	return &PluginsGoldenConfigConfigComplianceBulkUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigConfigComplianceBulkUpdateParams contains all the parameters to send to the API endpoint
   for the plugins golden config config compliance bulk update operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigConfigComplianceBulkUpdateParams struct {

	// Data.
	Data *models.ConfigCompliance

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config config compliance bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigConfigComplianceBulkUpdateParams) WithDefaults() *PluginsGoldenConfigConfigComplianceBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config config compliance bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigConfigComplianceBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config config compliance bulk update params
func (o *PluginsGoldenConfigConfigComplianceBulkUpdateParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigConfigComplianceBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config config compliance bulk update params
func (o *PluginsGoldenConfigConfigComplianceBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config config compliance bulk update params
func (o *PluginsGoldenConfigConfigComplianceBulkUpdateParams) WithContext(ctx context.Context) *PluginsGoldenConfigConfigComplianceBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config config compliance bulk update params
func (o *PluginsGoldenConfigConfigComplianceBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config config compliance bulk update params
func (o *PluginsGoldenConfigConfigComplianceBulkUpdateParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigConfigComplianceBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config config compliance bulk update params
func (o *PluginsGoldenConfigConfigComplianceBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins golden config config compliance bulk update params
func (o *PluginsGoldenConfigConfigComplianceBulkUpdateParams) WithData(data *models.ConfigCompliance) *PluginsGoldenConfigConfigComplianceBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins golden config config compliance bulk update params
func (o *PluginsGoldenConfigConfigComplianceBulkUpdateParams) SetData(data *models.ConfigCompliance) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigConfigComplianceBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

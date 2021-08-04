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

// NewPluginsGoldenConfigConfigCompliancePartialUpdateParams creates a new PluginsGoldenConfigConfigCompliancePartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigConfigCompliancePartialUpdateParams() *PluginsGoldenConfigConfigCompliancePartialUpdateParams {
	return &PluginsGoldenConfigConfigCompliancePartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigConfigCompliancePartialUpdateParamsWithTimeout creates a new PluginsGoldenConfigConfigCompliancePartialUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigConfigCompliancePartialUpdateParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigConfigCompliancePartialUpdateParams {
	return &PluginsGoldenConfigConfigCompliancePartialUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigConfigCompliancePartialUpdateParamsWithContext creates a new PluginsGoldenConfigConfigCompliancePartialUpdateParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigConfigCompliancePartialUpdateParamsWithContext(ctx context.Context) *PluginsGoldenConfigConfigCompliancePartialUpdateParams {
	return &PluginsGoldenConfigConfigCompliancePartialUpdateParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigConfigCompliancePartialUpdateParamsWithHTTPClient creates a new PluginsGoldenConfigConfigCompliancePartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigConfigCompliancePartialUpdateParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigConfigCompliancePartialUpdateParams {
	return &PluginsGoldenConfigConfigCompliancePartialUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigConfigCompliancePartialUpdateParams contains all the parameters to send to the API endpoint
   for the plugins golden config config compliance partial update operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigConfigCompliancePartialUpdateParams struct {

	// Data.
	Data *models.ConfigCompliance

	/* ID.

	   A UUID string identifying this config compliance.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config config compliance partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigConfigCompliancePartialUpdateParams) WithDefaults() *PluginsGoldenConfigConfigCompliancePartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config config compliance partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigConfigCompliancePartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config config compliance partial update params
func (o *PluginsGoldenConfigConfigCompliancePartialUpdateParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigConfigCompliancePartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config config compliance partial update params
func (o *PluginsGoldenConfigConfigCompliancePartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config config compliance partial update params
func (o *PluginsGoldenConfigConfigCompliancePartialUpdateParams) WithContext(ctx context.Context) *PluginsGoldenConfigConfigCompliancePartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config config compliance partial update params
func (o *PluginsGoldenConfigConfigCompliancePartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config config compliance partial update params
func (o *PluginsGoldenConfigConfigCompliancePartialUpdateParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigConfigCompliancePartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config config compliance partial update params
func (o *PluginsGoldenConfigConfigCompliancePartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins golden config config compliance partial update params
func (o *PluginsGoldenConfigConfigCompliancePartialUpdateParams) WithData(data *models.ConfigCompliance) *PluginsGoldenConfigConfigCompliancePartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins golden config config compliance partial update params
func (o *PluginsGoldenConfigConfigCompliancePartialUpdateParams) SetData(data *models.ConfigCompliance) {
	o.Data = data
}

// WithID adds the id to the plugins golden config config compliance partial update params
func (o *PluginsGoldenConfigConfigCompliancePartialUpdateParams) WithID(id strfmt.UUID) *PluginsGoldenConfigConfigCompliancePartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins golden config config compliance partial update params
func (o *PluginsGoldenConfigConfigCompliancePartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigConfigCompliancePartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

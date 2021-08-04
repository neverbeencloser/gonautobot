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

// NewPluginsGoldenConfigConfigReplacePartialUpdateParams creates a new PluginsGoldenConfigConfigReplacePartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigConfigReplacePartialUpdateParams() *PluginsGoldenConfigConfigReplacePartialUpdateParams {
	return &PluginsGoldenConfigConfigReplacePartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigConfigReplacePartialUpdateParamsWithTimeout creates a new PluginsGoldenConfigConfigReplacePartialUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigConfigReplacePartialUpdateParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigConfigReplacePartialUpdateParams {
	return &PluginsGoldenConfigConfigReplacePartialUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigConfigReplacePartialUpdateParamsWithContext creates a new PluginsGoldenConfigConfigReplacePartialUpdateParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigConfigReplacePartialUpdateParamsWithContext(ctx context.Context) *PluginsGoldenConfigConfigReplacePartialUpdateParams {
	return &PluginsGoldenConfigConfigReplacePartialUpdateParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigConfigReplacePartialUpdateParamsWithHTTPClient creates a new PluginsGoldenConfigConfigReplacePartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigConfigReplacePartialUpdateParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigConfigReplacePartialUpdateParams {
	return &PluginsGoldenConfigConfigReplacePartialUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigConfigReplacePartialUpdateParams contains all the parameters to send to the API endpoint
   for the plugins golden config config replace partial update operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigConfigReplacePartialUpdateParams struct {

	// Data.
	Data *models.ConfigReplace

	/* ID.

	   A UUID string identifying this config replace.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config config replace partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigConfigReplacePartialUpdateParams) WithDefaults() *PluginsGoldenConfigConfigReplacePartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config config replace partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigConfigReplacePartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config config replace partial update params
func (o *PluginsGoldenConfigConfigReplacePartialUpdateParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigConfigReplacePartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config config replace partial update params
func (o *PluginsGoldenConfigConfigReplacePartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config config replace partial update params
func (o *PluginsGoldenConfigConfigReplacePartialUpdateParams) WithContext(ctx context.Context) *PluginsGoldenConfigConfigReplacePartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config config replace partial update params
func (o *PluginsGoldenConfigConfigReplacePartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config config replace partial update params
func (o *PluginsGoldenConfigConfigReplacePartialUpdateParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigConfigReplacePartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config config replace partial update params
func (o *PluginsGoldenConfigConfigReplacePartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins golden config config replace partial update params
func (o *PluginsGoldenConfigConfigReplacePartialUpdateParams) WithData(data *models.ConfigReplace) *PluginsGoldenConfigConfigReplacePartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins golden config config replace partial update params
func (o *PluginsGoldenConfigConfigReplacePartialUpdateParams) SetData(data *models.ConfigReplace) {
	o.Data = data
}

// WithID adds the id to the plugins golden config config replace partial update params
func (o *PluginsGoldenConfigConfigReplacePartialUpdateParams) WithID(id strfmt.UUID) *PluginsGoldenConfigConfigReplacePartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins golden config config replace partial update params
func (o *PluginsGoldenConfigConfigReplacePartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigConfigReplacePartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

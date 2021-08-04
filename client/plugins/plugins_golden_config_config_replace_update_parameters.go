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

// NewPluginsGoldenConfigConfigReplaceUpdateParams creates a new PluginsGoldenConfigConfigReplaceUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigConfigReplaceUpdateParams() *PluginsGoldenConfigConfigReplaceUpdateParams {
	return &PluginsGoldenConfigConfigReplaceUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigConfigReplaceUpdateParamsWithTimeout creates a new PluginsGoldenConfigConfigReplaceUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigConfigReplaceUpdateParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigConfigReplaceUpdateParams {
	return &PluginsGoldenConfigConfigReplaceUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigConfigReplaceUpdateParamsWithContext creates a new PluginsGoldenConfigConfigReplaceUpdateParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigConfigReplaceUpdateParamsWithContext(ctx context.Context) *PluginsGoldenConfigConfigReplaceUpdateParams {
	return &PluginsGoldenConfigConfigReplaceUpdateParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigConfigReplaceUpdateParamsWithHTTPClient creates a new PluginsGoldenConfigConfigReplaceUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigConfigReplaceUpdateParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigConfigReplaceUpdateParams {
	return &PluginsGoldenConfigConfigReplaceUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigConfigReplaceUpdateParams contains all the parameters to send to the API endpoint
   for the plugins golden config config replace update operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigConfigReplaceUpdateParams struct {

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

// WithDefaults hydrates default values in the plugins golden config config replace update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigConfigReplaceUpdateParams) WithDefaults() *PluginsGoldenConfigConfigReplaceUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config config replace update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigConfigReplaceUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config config replace update params
func (o *PluginsGoldenConfigConfigReplaceUpdateParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigConfigReplaceUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config config replace update params
func (o *PluginsGoldenConfigConfigReplaceUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config config replace update params
func (o *PluginsGoldenConfigConfigReplaceUpdateParams) WithContext(ctx context.Context) *PluginsGoldenConfigConfigReplaceUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config config replace update params
func (o *PluginsGoldenConfigConfigReplaceUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config config replace update params
func (o *PluginsGoldenConfigConfigReplaceUpdateParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigConfigReplaceUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config config replace update params
func (o *PluginsGoldenConfigConfigReplaceUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins golden config config replace update params
func (o *PluginsGoldenConfigConfigReplaceUpdateParams) WithData(data *models.ConfigReplace) *PluginsGoldenConfigConfigReplaceUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins golden config config replace update params
func (o *PluginsGoldenConfigConfigReplaceUpdateParams) SetData(data *models.ConfigReplace) {
	o.Data = data
}

// WithID adds the id to the plugins golden config config replace update params
func (o *PluginsGoldenConfigConfigReplaceUpdateParams) WithID(id strfmt.UUID) *PluginsGoldenConfigConfigReplaceUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins golden config config replace update params
func (o *PluginsGoldenConfigConfigReplaceUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigConfigReplaceUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

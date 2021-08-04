package extras

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/models"
)

// NewExtrasConfigContextSchemasCreateParams creates a new ExtrasConfigContextSchemasCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasConfigContextSchemasCreateParams() *ExtrasConfigContextSchemasCreateParams {
	return &ExtrasConfigContextSchemasCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasConfigContextSchemasCreateParamsWithTimeout creates a new ExtrasConfigContextSchemasCreateParams object
// with the ability to set a timeout on a request.
func NewExtrasConfigContextSchemasCreateParamsWithTimeout(timeout time.Duration) *ExtrasConfigContextSchemasCreateParams {
	return &ExtrasConfigContextSchemasCreateParams{
		timeout: timeout,
	}
}

// NewExtrasConfigContextSchemasCreateParamsWithContext creates a new ExtrasConfigContextSchemasCreateParams object
// with the ability to set a context for a request.
func NewExtrasConfigContextSchemasCreateParamsWithContext(ctx context.Context) *ExtrasConfigContextSchemasCreateParams {
	return &ExtrasConfigContextSchemasCreateParams{
		Context: ctx,
	}
}

// NewExtrasConfigContextSchemasCreateParamsWithHTTPClient creates a new ExtrasConfigContextSchemasCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasConfigContextSchemasCreateParamsWithHTTPClient(client *http.Client) *ExtrasConfigContextSchemasCreateParams {
	return &ExtrasConfigContextSchemasCreateParams{
		HTTPClient: client,
	}
}

/* ExtrasConfigContextSchemasCreateParams contains all the parameters to send to the API endpoint
   for the extras config context schemas create operation.

   Typically these are written to a http.Request.
*/
type ExtrasConfigContextSchemasCreateParams struct {

	// Data.
	Data *models.ConfigContextSchema

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras config context schemas create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasConfigContextSchemasCreateParams) WithDefaults() *ExtrasConfigContextSchemasCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras config context schemas create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasConfigContextSchemasCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras config context schemas create params
func (o *ExtrasConfigContextSchemasCreateParams) WithTimeout(timeout time.Duration) *ExtrasConfigContextSchemasCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras config context schemas create params
func (o *ExtrasConfigContextSchemasCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras config context schemas create params
func (o *ExtrasConfigContextSchemasCreateParams) WithContext(ctx context.Context) *ExtrasConfigContextSchemasCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras config context schemas create params
func (o *ExtrasConfigContextSchemasCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras config context schemas create params
func (o *ExtrasConfigContextSchemasCreateParams) WithHTTPClient(client *http.Client) *ExtrasConfigContextSchemasCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras config context schemas create params
func (o *ExtrasConfigContextSchemasCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras config context schemas create params
func (o *ExtrasConfigContextSchemasCreateParams) WithData(data *models.ConfigContextSchema) *ExtrasConfigContextSchemasCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras config context schemas create params
func (o *ExtrasConfigContextSchemasCreateParams) SetData(data *models.ConfigContextSchema) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasConfigContextSchemasCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

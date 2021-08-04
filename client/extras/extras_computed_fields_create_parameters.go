package extras

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

// NewExtrasComputedFieldsCreateParams creates a new ExtrasComputedFieldsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasComputedFieldsCreateParams() *ExtrasComputedFieldsCreateParams {
	return &ExtrasComputedFieldsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasComputedFieldsCreateParamsWithTimeout creates a new ExtrasComputedFieldsCreateParams object
// with the ability to set a timeout on a request.
func NewExtrasComputedFieldsCreateParamsWithTimeout(timeout time.Duration) *ExtrasComputedFieldsCreateParams {
	return &ExtrasComputedFieldsCreateParams{
		timeout: timeout,
	}
}

// NewExtrasComputedFieldsCreateParamsWithContext creates a new ExtrasComputedFieldsCreateParams object
// with the ability to set a context for a request.
func NewExtrasComputedFieldsCreateParamsWithContext(ctx context.Context) *ExtrasComputedFieldsCreateParams {
	return &ExtrasComputedFieldsCreateParams{
		Context: ctx,
	}
}

// NewExtrasComputedFieldsCreateParamsWithHTTPClient creates a new ExtrasComputedFieldsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasComputedFieldsCreateParamsWithHTTPClient(client *http.Client) *ExtrasComputedFieldsCreateParams {
	return &ExtrasComputedFieldsCreateParams{
		HTTPClient: client,
	}
}

/* ExtrasComputedFieldsCreateParams contains all the parameters to send to the API endpoint
   for the extras computed fields create operation.

   Typically these are written to a http.Request.
*/
type ExtrasComputedFieldsCreateParams struct {

	// Data.
	Data *models.ComputedField

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras computed fields create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasComputedFieldsCreateParams) WithDefaults() *ExtrasComputedFieldsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras computed fields create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasComputedFieldsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras computed fields create params
func (o *ExtrasComputedFieldsCreateParams) WithTimeout(timeout time.Duration) *ExtrasComputedFieldsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras computed fields create params
func (o *ExtrasComputedFieldsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras computed fields create params
func (o *ExtrasComputedFieldsCreateParams) WithContext(ctx context.Context) *ExtrasComputedFieldsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras computed fields create params
func (o *ExtrasComputedFieldsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras computed fields create params
func (o *ExtrasComputedFieldsCreateParams) WithHTTPClient(client *http.Client) *ExtrasComputedFieldsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras computed fields create params
func (o *ExtrasComputedFieldsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras computed fields create params
func (o *ExtrasComputedFieldsCreateParams) WithData(data *models.ComputedField) *ExtrasComputedFieldsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras computed fields create params
func (o *ExtrasComputedFieldsCreateParams) SetData(data *models.ComputedField) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasComputedFieldsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

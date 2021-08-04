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

// NewExtrasComputedFieldsPartialUpdateParams creates a new ExtrasComputedFieldsPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasComputedFieldsPartialUpdateParams() *ExtrasComputedFieldsPartialUpdateParams {
	return &ExtrasComputedFieldsPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasComputedFieldsPartialUpdateParamsWithTimeout creates a new ExtrasComputedFieldsPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasComputedFieldsPartialUpdateParamsWithTimeout(timeout time.Duration) *ExtrasComputedFieldsPartialUpdateParams {
	return &ExtrasComputedFieldsPartialUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasComputedFieldsPartialUpdateParamsWithContext creates a new ExtrasComputedFieldsPartialUpdateParams object
// with the ability to set a context for a request.
func NewExtrasComputedFieldsPartialUpdateParamsWithContext(ctx context.Context) *ExtrasComputedFieldsPartialUpdateParams {
	return &ExtrasComputedFieldsPartialUpdateParams{
		Context: ctx,
	}
}

// NewExtrasComputedFieldsPartialUpdateParamsWithHTTPClient creates a new ExtrasComputedFieldsPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasComputedFieldsPartialUpdateParamsWithHTTPClient(client *http.Client) *ExtrasComputedFieldsPartialUpdateParams {
	return &ExtrasComputedFieldsPartialUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasComputedFieldsPartialUpdateParams contains all the parameters to send to the API endpoint
   for the extras computed fields partial update operation.

   Typically these are written to a http.Request.
*/
type ExtrasComputedFieldsPartialUpdateParams struct {

	// Data.
	Data *models.ComputedField

	/* ID.

	   A UUID string identifying this computed field.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras computed fields partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasComputedFieldsPartialUpdateParams) WithDefaults() *ExtrasComputedFieldsPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras computed fields partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasComputedFieldsPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras computed fields partial update params
func (o *ExtrasComputedFieldsPartialUpdateParams) WithTimeout(timeout time.Duration) *ExtrasComputedFieldsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras computed fields partial update params
func (o *ExtrasComputedFieldsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras computed fields partial update params
func (o *ExtrasComputedFieldsPartialUpdateParams) WithContext(ctx context.Context) *ExtrasComputedFieldsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras computed fields partial update params
func (o *ExtrasComputedFieldsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras computed fields partial update params
func (o *ExtrasComputedFieldsPartialUpdateParams) WithHTTPClient(client *http.Client) *ExtrasComputedFieldsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras computed fields partial update params
func (o *ExtrasComputedFieldsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras computed fields partial update params
func (o *ExtrasComputedFieldsPartialUpdateParams) WithData(data *models.ComputedField) *ExtrasComputedFieldsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras computed fields partial update params
func (o *ExtrasComputedFieldsPartialUpdateParams) SetData(data *models.ComputedField) {
	o.Data = data
}

// WithID adds the id to the extras computed fields partial update params
func (o *ExtrasComputedFieldsPartialUpdateParams) WithID(id strfmt.UUID) *ExtrasComputedFieldsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras computed fields partial update params
func (o *ExtrasComputedFieldsPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasComputedFieldsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

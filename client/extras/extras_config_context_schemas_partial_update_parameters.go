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

// NewExtrasConfigContextSchemasPartialUpdateParams creates a new ExtrasConfigContextSchemasPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasConfigContextSchemasPartialUpdateParams() *ExtrasConfigContextSchemasPartialUpdateParams {
	return &ExtrasConfigContextSchemasPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasConfigContextSchemasPartialUpdateParamsWithTimeout creates a new ExtrasConfigContextSchemasPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasConfigContextSchemasPartialUpdateParamsWithTimeout(timeout time.Duration) *ExtrasConfigContextSchemasPartialUpdateParams {
	return &ExtrasConfigContextSchemasPartialUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasConfigContextSchemasPartialUpdateParamsWithContext creates a new ExtrasConfigContextSchemasPartialUpdateParams object
// with the ability to set a context for a request.
func NewExtrasConfigContextSchemasPartialUpdateParamsWithContext(ctx context.Context) *ExtrasConfigContextSchemasPartialUpdateParams {
	return &ExtrasConfigContextSchemasPartialUpdateParams{
		Context: ctx,
	}
}

// NewExtrasConfigContextSchemasPartialUpdateParamsWithHTTPClient creates a new ExtrasConfigContextSchemasPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasConfigContextSchemasPartialUpdateParamsWithHTTPClient(client *http.Client) *ExtrasConfigContextSchemasPartialUpdateParams {
	return &ExtrasConfigContextSchemasPartialUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasConfigContextSchemasPartialUpdateParams contains all the parameters to send to the API endpoint
   for the extras config context schemas partial update operation.

   Typically these are written to a http.Request.
*/
type ExtrasConfigContextSchemasPartialUpdateParams struct {

	// Data.
	Data *models.ConfigContextSchema

	/* ID.

	   A UUID string identifying this config context schema.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras config context schemas partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasConfigContextSchemasPartialUpdateParams) WithDefaults() *ExtrasConfigContextSchemasPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras config context schemas partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasConfigContextSchemasPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras config context schemas partial update params
func (o *ExtrasConfigContextSchemasPartialUpdateParams) WithTimeout(timeout time.Duration) *ExtrasConfigContextSchemasPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras config context schemas partial update params
func (o *ExtrasConfigContextSchemasPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras config context schemas partial update params
func (o *ExtrasConfigContextSchemasPartialUpdateParams) WithContext(ctx context.Context) *ExtrasConfigContextSchemasPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras config context schemas partial update params
func (o *ExtrasConfigContextSchemasPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras config context schemas partial update params
func (o *ExtrasConfigContextSchemasPartialUpdateParams) WithHTTPClient(client *http.Client) *ExtrasConfigContextSchemasPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras config context schemas partial update params
func (o *ExtrasConfigContextSchemasPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras config context schemas partial update params
func (o *ExtrasConfigContextSchemasPartialUpdateParams) WithData(data *models.ConfigContextSchema) *ExtrasConfigContextSchemasPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras config context schemas partial update params
func (o *ExtrasConfigContextSchemasPartialUpdateParams) SetData(data *models.ConfigContextSchema) {
	o.Data = data
}

// WithID adds the id to the extras config context schemas partial update params
func (o *ExtrasConfigContextSchemasPartialUpdateParams) WithID(id strfmt.UUID) *ExtrasConfigContextSchemasPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras config context schemas partial update params
func (o *ExtrasConfigContextSchemasPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasConfigContextSchemasPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

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

// NewExtrasConfigContextSchemasUpdateParams creates a new ExtrasConfigContextSchemasUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasConfigContextSchemasUpdateParams() *ExtrasConfigContextSchemasUpdateParams {
	return &ExtrasConfigContextSchemasUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasConfigContextSchemasUpdateParamsWithTimeout creates a new ExtrasConfigContextSchemasUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasConfigContextSchemasUpdateParamsWithTimeout(timeout time.Duration) *ExtrasConfigContextSchemasUpdateParams {
	return &ExtrasConfigContextSchemasUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasConfigContextSchemasUpdateParamsWithContext creates a new ExtrasConfigContextSchemasUpdateParams object
// with the ability to set a context for a request.
func NewExtrasConfigContextSchemasUpdateParamsWithContext(ctx context.Context) *ExtrasConfigContextSchemasUpdateParams {
	return &ExtrasConfigContextSchemasUpdateParams{
		Context: ctx,
	}
}

// NewExtrasConfigContextSchemasUpdateParamsWithHTTPClient creates a new ExtrasConfigContextSchemasUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasConfigContextSchemasUpdateParamsWithHTTPClient(client *http.Client) *ExtrasConfigContextSchemasUpdateParams {
	return &ExtrasConfigContextSchemasUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasConfigContextSchemasUpdateParams contains all the parameters to send to the API endpoint
   for the extras config context schemas update operation.

   Typically these are written to a http.Request.
*/
type ExtrasConfigContextSchemasUpdateParams struct {

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

// WithDefaults hydrates default values in the extras config context schemas update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasConfigContextSchemasUpdateParams) WithDefaults() *ExtrasConfigContextSchemasUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras config context schemas update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasConfigContextSchemasUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras config context schemas update params
func (o *ExtrasConfigContextSchemasUpdateParams) WithTimeout(timeout time.Duration) *ExtrasConfigContextSchemasUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras config context schemas update params
func (o *ExtrasConfigContextSchemasUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras config context schemas update params
func (o *ExtrasConfigContextSchemasUpdateParams) WithContext(ctx context.Context) *ExtrasConfigContextSchemasUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras config context schemas update params
func (o *ExtrasConfigContextSchemasUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras config context schemas update params
func (o *ExtrasConfigContextSchemasUpdateParams) WithHTTPClient(client *http.Client) *ExtrasConfigContextSchemasUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras config context schemas update params
func (o *ExtrasConfigContextSchemasUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras config context schemas update params
func (o *ExtrasConfigContextSchemasUpdateParams) WithData(data *models.ConfigContextSchema) *ExtrasConfigContextSchemasUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras config context schemas update params
func (o *ExtrasConfigContextSchemasUpdateParams) SetData(data *models.ConfigContextSchema) {
	o.Data = data
}

// WithID adds the id to the extras config context schemas update params
func (o *ExtrasConfigContextSchemasUpdateParams) WithID(id strfmt.UUID) *ExtrasConfigContextSchemasUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras config context schemas update params
func (o *ExtrasConfigContextSchemasUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasConfigContextSchemasUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

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

// NewExtrasCustomFieldChoicesPartialUpdateParams creates a new ExtrasCustomFieldChoicesPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasCustomFieldChoicesPartialUpdateParams() *ExtrasCustomFieldChoicesPartialUpdateParams {
	return &ExtrasCustomFieldChoicesPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasCustomFieldChoicesPartialUpdateParamsWithTimeout creates a new ExtrasCustomFieldChoicesPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasCustomFieldChoicesPartialUpdateParamsWithTimeout(timeout time.Duration) *ExtrasCustomFieldChoicesPartialUpdateParams {
	return &ExtrasCustomFieldChoicesPartialUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasCustomFieldChoicesPartialUpdateParamsWithContext creates a new ExtrasCustomFieldChoicesPartialUpdateParams object
// with the ability to set a context for a request.
func NewExtrasCustomFieldChoicesPartialUpdateParamsWithContext(ctx context.Context) *ExtrasCustomFieldChoicesPartialUpdateParams {
	return &ExtrasCustomFieldChoicesPartialUpdateParams{
		Context: ctx,
	}
}

// NewExtrasCustomFieldChoicesPartialUpdateParamsWithHTTPClient creates a new ExtrasCustomFieldChoicesPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasCustomFieldChoicesPartialUpdateParamsWithHTTPClient(client *http.Client) *ExtrasCustomFieldChoicesPartialUpdateParams {
	return &ExtrasCustomFieldChoicesPartialUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasCustomFieldChoicesPartialUpdateParams contains all the parameters to send to the API endpoint
   for the extras custom field choices partial update operation.

   Typically these are written to a http.Request.
*/
type ExtrasCustomFieldChoicesPartialUpdateParams struct {

	// Data.
	Data *models.WritableCustomFieldChoice

	/* ID.

	   A UUID string identifying this custom field choice.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras custom field choices partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasCustomFieldChoicesPartialUpdateParams) WithDefaults() *ExtrasCustomFieldChoicesPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras custom field choices partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasCustomFieldChoicesPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras custom field choices partial update params
func (o *ExtrasCustomFieldChoicesPartialUpdateParams) WithTimeout(timeout time.Duration) *ExtrasCustomFieldChoicesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras custom field choices partial update params
func (o *ExtrasCustomFieldChoicesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras custom field choices partial update params
func (o *ExtrasCustomFieldChoicesPartialUpdateParams) WithContext(ctx context.Context) *ExtrasCustomFieldChoicesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras custom field choices partial update params
func (o *ExtrasCustomFieldChoicesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras custom field choices partial update params
func (o *ExtrasCustomFieldChoicesPartialUpdateParams) WithHTTPClient(client *http.Client) *ExtrasCustomFieldChoicesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras custom field choices partial update params
func (o *ExtrasCustomFieldChoicesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras custom field choices partial update params
func (o *ExtrasCustomFieldChoicesPartialUpdateParams) WithData(data *models.WritableCustomFieldChoice) *ExtrasCustomFieldChoicesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras custom field choices partial update params
func (o *ExtrasCustomFieldChoicesPartialUpdateParams) SetData(data *models.WritableCustomFieldChoice) {
	o.Data = data
}

// WithID adds the id to the extras custom field choices partial update params
func (o *ExtrasCustomFieldChoicesPartialUpdateParams) WithID(id strfmt.UUID) *ExtrasCustomFieldChoicesPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras custom field choices partial update params
func (o *ExtrasCustomFieldChoicesPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasCustomFieldChoicesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

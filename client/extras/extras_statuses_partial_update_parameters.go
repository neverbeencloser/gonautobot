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

// NewExtrasStatusesPartialUpdateParams creates a new ExtrasStatusesPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasStatusesPartialUpdateParams() *ExtrasStatusesPartialUpdateParams {
	return &ExtrasStatusesPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasStatusesPartialUpdateParamsWithTimeout creates a new ExtrasStatusesPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasStatusesPartialUpdateParamsWithTimeout(timeout time.Duration) *ExtrasStatusesPartialUpdateParams {
	return &ExtrasStatusesPartialUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasStatusesPartialUpdateParamsWithContext creates a new ExtrasStatusesPartialUpdateParams object
// with the ability to set a context for a request.
func NewExtrasStatusesPartialUpdateParamsWithContext(ctx context.Context) *ExtrasStatusesPartialUpdateParams {
	return &ExtrasStatusesPartialUpdateParams{
		Context: ctx,
	}
}

// NewExtrasStatusesPartialUpdateParamsWithHTTPClient creates a new ExtrasStatusesPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasStatusesPartialUpdateParamsWithHTTPClient(client *http.Client) *ExtrasStatusesPartialUpdateParams {
	return &ExtrasStatusesPartialUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasStatusesPartialUpdateParams contains all the parameters to send to the API endpoint
   for the extras statuses partial update operation.

   Typically these are written to a http.Request.
*/
type ExtrasStatusesPartialUpdateParams struct {

	// Data.
	Data *models.Status

	/* ID.

	   A UUID string identifying this status.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras statuses partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasStatusesPartialUpdateParams) WithDefaults() *ExtrasStatusesPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras statuses partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasStatusesPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras statuses partial update params
func (o *ExtrasStatusesPartialUpdateParams) WithTimeout(timeout time.Duration) *ExtrasStatusesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras statuses partial update params
func (o *ExtrasStatusesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras statuses partial update params
func (o *ExtrasStatusesPartialUpdateParams) WithContext(ctx context.Context) *ExtrasStatusesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras statuses partial update params
func (o *ExtrasStatusesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras statuses partial update params
func (o *ExtrasStatusesPartialUpdateParams) WithHTTPClient(client *http.Client) *ExtrasStatusesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras statuses partial update params
func (o *ExtrasStatusesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras statuses partial update params
func (o *ExtrasStatusesPartialUpdateParams) WithData(data *models.Status) *ExtrasStatusesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras statuses partial update params
func (o *ExtrasStatusesPartialUpdateParams) SetData(data *models.Status) {
	o.Data = data
}

// WithID adds the id to the extras statuses partial update params
func (o *ExtrasStatusesPartialUpdateParams) WithID(id strfmt.UUID) *ExtrasStatusesPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras statuses partial update params
func (o *ExtrasStatusesPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasStatusesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

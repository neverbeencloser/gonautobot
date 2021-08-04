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

// NewExtrasStatusesUpdateParams creates a new ExtrasStatusesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasStatusesUpdateParams() *ExtrasStatusesUpdateParams {
	return &ExtrasStatusesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasStatusesUpdateParamsWithTimeout creates a new ExtrasStatusesUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasStatusesUpdateParamsWithTimeout(timeout time.Duration) *ExtrasStatusesUpdateParams {
	return &ExtrasStatusesUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasStatusesUpdateParamsWithContext creates a new ExtrasStatusesUpdateParams object
// with the ability to set a context for a request.
func NewExtrasStatusesUpdateParamsWithContext(ctx context.Context) *ExtrasStatusesUpdateParams {
	return &ExtrasStatusesUpdateParams{
		Context: ctx,
	}
}

// NewExtrasStatusesUpdateParamsWithHTTPClient creates a new ExtrasStatusesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasStatusesUpdateParamsWithHTTPClient(client *http.Client) *ExtrasStatusesUpdateParams {
	return &ExtrasStatusesUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasStatusesUpdateParams contains all the parameters to send to the API endpoint
   for the extras statuses update operation.

   Typically these are written to a http.Request.
*/
type ExtrasStatusesUpdateParams struct {

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

// WithDefaults hydrates default values in the extras statuses update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasStatusesUpdateParams) WithDefaults() *ExtrasStatusesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras statuses update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasStatusesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras statuses update params
func (o *ExtrasStatusesUpdateParams) WithTimeout(timeout time.Duration) *ExtrasStatusesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras statuses update params
func (o *ExtrasStatusesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras statuses update params
func (o *ExtrasStatusesUpdateParams) WithContext(ctx context.Context) *ExtrasStatusesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras statuses update params
func (o *ExtrasStatusesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras statuses update params
func (o *ExtrasStatusesUpdateParams) WithHTTPClient(client *http.Client) *ExtrasStatusesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras statuses update params
func (o *ExtrasStatusesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras statuses update params
func (o *ExtrasStatusesUpdateParams) WithData(data *models.Status) *ExtrasStatusesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras statuses update params
func (o *ExtrasStatusesUpdateParams) SetData(data *models.Status) {
	o.Data = data
}

// WithID adds the id to the extras statuses update params
func (o *ExtrasStatusesUpdateParams) WithID(id strfmt.UUID) *ExtrasStatusesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras statuses update params
func (o *ExtrasStatusesUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasStatusesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

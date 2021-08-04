package extras

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewExtrasStatusesReadParams creates a new ExtrasStatusesReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasStatusesReadParams() *ExtrasStatusesReadParams {
	return &ExtrasStatusesReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasStatusesReadParamsWithTimeout creates a new ExtrasStatusesReadParams object
// with the ability to set a timeout on a request.
func NewExtrasStatusesReadParamsWithTimeout(timeout time.Duration) *ExtrasStatusesReadParams {
	return &ExtrasStatusesReadParams{
		timeout: timeout,
	}
}

// NewExtrasStatusesReadParamsWithContext creates a new ExtrasStatusesReadParams object
// with the ability to set a context for a request.
func NewExtrasStatusesReadParamsWithContext(ctx context.Context) *ExtrasStatusesReadParams {
	return &ExtrasStatusesReadParams{
		Context: ctx,
	}
}

// NewExtrasStatusesReadParamsWithHTTPClient creates a new ExtrasStatusesReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasStatusesReadParamsWithHTTPClient(client *http.Client) *ExtrasStatusesReadParams {
	return &ExtrasStatusesReadParams{
		HTTPClient: client,
	}
}

/* ExtrasStatusesReadParams contains all the parameters to send to the API endpoint
   for the extras statuses read operation.

   Typically these are written to a http.Request.
*/
type ExtrasStatusesReadParams struct {

	/* ID.

	   A UUID string identifying this status.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras statuses read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasStatusesReadParams) WithDefaults() *ExtrasStatusesReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras statuses read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasStatusesReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras statuses read params
func (o *ExtrasStatusesReadParams) WithTimeout(timeout time.Duration) *ExtrasStatusesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras statuses read params
func (o *ExtrasStatusesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras statuses read params
func (o *ExtrasStatusesReadParams) WithContext(ctx context.Context) *ExtrasStatusesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras statuses read params
func (o *ExtrasStatusesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras statuses read params
func (o *ExtrasStatusesReadParams) WithHTTPClient(client *http.Client) *ExtrasStatusesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras statuses read params
func (o *ExtrasStatusesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the extras statuses read params
func (o *ExtrasStatusesReadParams) WithID(id strfmt.UUID) *ExtrasStatusesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras statuses read params
func (o *ExtrasStatusesReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasStatusesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

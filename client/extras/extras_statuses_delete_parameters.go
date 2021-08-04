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

// NewExtrasStatusesDeleteParams creates a new ExtrasStatusesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasStatusesDeleteParams() *ExtrasStatusesDeleteParams {
	return &ExtrasStatusesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasStatusesDeleteParamsWithTimeout creates a new ExtrasStatusesDeleteParams object
// with the ability to set a timeout on a request.
func NewExtrasStatusesDeleteParamsWithTimeout(timeout time.Duration) *ExtrasStatusesDeleteParams {
	return &ExtrasStatusesDeleteParams{
		timeout: timeout,
	}
}

// NewExtrasStatusesDeleteParamsWithContext creates a new ExtrasStatusesDeleteParams object
// with the ability to set a context for a request.
func NewExtrasStatusesDeleteParamsWithContext(ctx context.Context) *ExtrasStatusesDeleteParams {
	return &ExtrasStatusesDeleteParams{
		Context: ctx,
	}
}

// NewExtrasStatusesDeleteParamsWithHTTPClient creates a new ExtrasStatusesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasStatusesDeleteParamsWithHTTPClient(client *http.Client) *ExtrasStatusesDeleteParams {
	return &ExtrasStatusesDeleteParams{
		HTTPClient: client,
	}
}

/* ExtrasStatusesDeleteParams contains all the parameters to send to the API endpoint
   for the extras statuses delete operation.

   Typically these are written to a http.Request.
*/
type ExtrasStatusesDeleteParams struct {

	/* ID.

	   A UUID string identifying this status.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras statuses delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasStatusesDeleteParams) WithDefaults() *ExtrasStatusesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras statuses delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasStatusesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras statuses delete params
func (o *ExtrasStatusesDeleteParams) WithTimeout(timeout time.Duration) *ExtrasStatusesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras statuses delete params
func (o *ExtrasStatusesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras statuses delete params
func (o *ExtrasStatusesDeleteParams) WithContext(ctx context.Context) *ExtrasStatusesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras statuses delete params
func (o *ExtrasStatusesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras statuses delete params
func (o *ExtrasStatusesDeleteParams) WithHTTPClient(client *http.Client) *ExtrasStatusesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras statuses delete params
func (o *ExtrasStatusesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the extras statuses delete params
func (o *ExtrasStatusesDeleteParams) WithID(id strfmt.UUID) *ExtrasStatusesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras statuses delete params
func (o *ExtrasStatusesDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasStatusesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

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

// NewExtrasComputedFieldsDeleteParams creates a new ExtrasComputedFieldsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasComputedFieldsDeleteParams() *ExtrasComputedFieldsDeleteParams {
	return &ExtrasComputedFieldsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasComputedFieldsDeleteParamsWithTimeout creates a new ExtrasComputedFieldsDeleteParams object
// with the ability to set a timeout on a request.
func NewExtrasComputedFieldsDeleteParamsWithTimeout(timeout time.Duration) *ExtrasComputedFieldsDeleteParams {
	return &ExtrasComputedFieldsDeleteParams{
		timeout: timeout,
	}
}

// NewExtrasComputedFieldsDeleteParamsWithContext creates a new ExtrasComputedFieldsDeleteParams object
// with the ability to set a context for a request.
func NewExtrasComputedFieldsDeleteParamsWithContext(ctx context.Context) *ExtrasComputedFieldsDeleteParams {
	return &ExtrasComputedFieldsDeleteParams{
		Context: ctx,
	}
}

// NewExtrasComputedFieldsDeleteParamsWithHTTPClient creates a new ExtrasComputedFieldsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasComputedFieldsDeleteParamsWithHTTPClient(client *http.Client) *ExtrasComputedFieldsDeleteParams {
	return &ExtrasComputedFieldsDeleteParams{
		HTTPClient: client,
	}
}

/* ExtrasComputedFieldsDeleteParams contains all the parameters to send to the API endpoint
   for the extras computed fields delete operation.

   Typically these are written to a http.Request.
*/
type ExtrasComputedFieldsDeleteParams struct {

	/* ID.

	   A UUID string identifying this computed field.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras computed fields delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasComputedFieldsDeleteParams) WithDefaults() *ExtrasComputedFieldsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras computed fields delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasComputedFieldsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras computed fields delete params
func (o *ExtrasComputedFieldsDeleteParams) WithTimeout(timeout time.Duration) *ExtrasComputedFieldsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras computed fields delete params
func (o *ExtrasComputedFieldsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras computed fields delete params
func (o *ExtrasComputedFieldsDeleteParams) WithContext(ctx context.Context) *ExtrasComputedFieldsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras computed fields delete params
func (o *ExtrasComputedFieldsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras computed fields delete params
func (o *ExtrasComputedFieldsDeleteParams) WithHTTPClient(client *http.Client) *ExtrasComputedFieldsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras computed fields delete params
func (o *ExtrasComputedFieldsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the extras computed fields delete params
func (o *ExtrasComputedFieldsDeleteParams) WithID(id strfmt.UUID) *ExtrasComputedFieldsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras computed fields delete params
func (o *ExtrasComputedFieldsDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasComputedFieldsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

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

// NewExtrasConfigContextSchemasDeleteParams creates a new ExtrasConfigContextSchemasDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasConfigContextSchemasDeleteParams() *ExtrasConfigContextSchemasDeleteParams {
	return &ExtrasConfigContextSchemasDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasConfigContextSchemasDeleteParamsWithTimeout creates a new ExtrasConfigContextSchemasDeleteParams object
// with the ability to set a timeout on a request.
func NewExtrasConfigContextSchemasDeleteParamsWithTimeout(timeout time.Duration) *ExtrasConfigContextSchemasDeleteParams {
	return &ExtrasConfigContextSchemasDeleteParams{
		timeout: timeout,
	}
}

// NewExtrasConfigContextSchemasDeleteParamsWithContext creates a new ExtrasConfigContextSchemasDeleteParams object
// with the ability to set a context for a request.
func NewExtrasConfigContextSchemasDeleteParamsWithContext(ctx context.Context) *ExtrasConfigContextSchemasDeleteParams {
	return &ExtrasConfigContextSchemasDeleteParams{
		Context: ctx,
	}
}

// NewExtrasConfigContextSchemasDeleteParamsWithHTTPClient creates a new ExtrasConfigContextSchemasDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasConfigContextSchemasDeleteParamsWithHTTPClient(client *http.Client) *ExtrasConfigContextSchemasDeleteParams {
	return &ExtrasConfigContextSchemasDeleteParams{
		HTTPClient: client,
	}
}

/* ExtrasConfigContextSchemasDeleteParams contains all the parameters to send to the API endpoint
   for the extras config context schemas delete operation.

   Typically these are written to a http.Request.
*/
type ExtrasConfigContextSchemasDeleteParams struct {

	/* ID.

	   A UUID string identifying this config context schema.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras config context schemas delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasConfigContextSchemasDeleteParams) WithDefaults() *ExtrasConfigContextSchemasDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras config context schemas delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasConfigContextSchemasDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras config context schemas delete params
func (o *ExtrasConfigContextSchemasDeleteParams) WithTimeout(timeout time.Duration) *ExtrasConfigContextSchemasDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras config context schemas delete params
func (o *ExtrasConfigContextSchemasDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras config context schemas delete params
func (o *ExtrasConfigContextSchemasDeleteParams) WithContext(ctx context.Context) *ExtrasConfigContextSchemasDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras config context schemas delete params
func (o *ExtrasConfigContextSchemasDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras config context schemas delete params
func (o *ExtrasConfigContextSchemasDeleteParams) WithHTTPClient(client *http.Client) *ExtrasConfigContextSchemasDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras config context schemas delete params
func (o *ExtrasConfigContextSchemasDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the extras config context schemas delete params
func (o *ExtrasConfigContextSchemasDeleteParams) WithID(id strfmt.UUID) *ExtrasConfigContextSchemasDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras config context schemas delete params
func (o *ExtrasConfigContextSchemasDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasConfigContextSchemasDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

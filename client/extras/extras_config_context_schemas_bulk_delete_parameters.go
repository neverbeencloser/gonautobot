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

// NewExtrasConfigContextSchemasBulkDeleteParams creates a new ExtrasConfigContextSchemasBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasConfigContextSchemasBulkDeleteParams() *ExtrasConfigContextSchemasBulkDeleteParams {
	return &ExtrasConfigContextSchemasBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasConfigContextSchemasBulkDeleteParamsWithTimeout creates a new ExtrasConfigContextSchemasBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewExtrasConfigContextSchemasBulkDeleteParamsWithTimeout(timeout time.Duration) *ExtrasConfigContextSchemasBulkDeleteParams {
	return &ExtrasConfigContextSchemasBulkDeleteParams{
		timeout: timeout,
	}
}

// NewExtrasConfigContextSchemasBulkDeleteParamsWithContext creates a new ExtrasConfigContextSchemasBulkDeleteParams object
// with the ability to set a context for a request.
func NewExtrasConfigContextSchemasBulkDeleteParamsWithContext(ctx context.Context) *ExtrasConfigContextSchemasBulkDeleteParams {
	return &ExtrasConfigContextSchemasBulkDeleteParams{
		Context: ctx,
	}
}

// NewExtrasConfigContextSchemasBulkDeleteParamsWithHTTPClient creates a new ExtrasConfigContextSchemasBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasConfigContextSchemasBulkDeleteParamsWithHTTPClient(client *http.Client) *ExtrasConfigContextSchemasBulkDeleteParams {
	return &ExtrasConfigContextSchemasBulkDeleteParams{
		HTTPClient: client,
	}
}

/* ExtrasConfigContextSchemasBulkDeleteParams contains all the parameters to send to the API endpoint
   for the extras config context schemas bulk delete operation.

   Typically these are written to a http.Request.
*/
type ExtrasConfigContextSchemasBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras config context schemas bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasConfigContextSchemasBulkDeleteParams) WithDefaults() *ExtrasConfigContextSchemasBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras config context schemas bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasConfigContextSchemasBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras config context schemas bulk delete params
func (o *ExtrasConfigContextSchemasBulkDeleteParams) WithTimeout(timeout time.Duration) *ExtrasConfigContextSchemasBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras config context schemas bulk delete params
func (o *ExtrasConfigContextSchemasBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras config context schemas bulk delete params
func (o *ExtrasConfigContextSchemasBulkDeleteParams) WithContext(ctx context.Context) *ExtrasConfigContextSchemasBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras config context schemas bulk delete params
func (o *ExtrasConfigContextSchemasBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras config context schemas bulk delete params
func (o *ExtrasConfigContextSchemasBulkDeleteParams) WithHTTPClient(client *http.Client) *ExtrasConfigContextSchemasBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras config context schemas bulk delete params
func (o *ExtrasConfigContextSchemasBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasConfigContextSchemasBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

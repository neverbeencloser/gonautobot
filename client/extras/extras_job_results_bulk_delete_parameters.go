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

// NewExtrasJobResultsBulkDeleteParams creates a new ExtrasJobResultsBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasJobResultsBulkDeleteParams() *ExtrasJobResultsBulkDeleteParams {
	return &ExtrasJobResultsBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasJobResultsBulkDeleteParamsWithTimeout creates a new ExtrasJobResultsBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewExtrasJobResultsBulkDeleteParamsWithTimeout(timeout time.Duration) *ExtrasJobResultsBulkDeleteParams {
	return &ExtrasJobResultsBulkDeleteParams{
		timeout: timeout,
	}
}

// NewExtrasJobResultsBulkDeleteParamsWithContext creates a new ExtrasJobResultsBulkDeleteParams object
// with the ability to set a context for a request.
func NewExtrasJobResultsBulkDeleteParamsWithContext(ctx context.Context) *ExtrasJobResultsBulkDeleteParams {
	return &ExtrasJobResultsBulkDeleteParams{
		Context: ctx,
	}
}

// NewExtrasJobResultsBulkDeleteParamsWithHTTPClient creates a new ExtrasJobResultsBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasJobResultsBulkDeleteParamsWithHTTPClient(client *http.Client) *ExtrasJobResultsBulkDeleteParams {
	return &ExtrasJobResultsBulkDeleteParams{
		HTTPClient: client,
	}
}

/* ExtrasJobResultsBulkDeleteParams contains all the parameters to send to the API endpoint
   for the extras job results bulk delete operation.

   Typically these are written to a http.Request.
*/
type ExtrasJobResultsBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras job results bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasJobResultsBulkDeleteParams) WithDefaults() *ExtrasJobResultsBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras job results bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasJobResultsBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras job results bulk delete params
func (o *ExtrasJobResultsBulkDeleteParams) WithTimeout(timeout time.Duration) *ExtrasJobResultsBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras job results bulk delete params
func (o *ExtrasJobResultsBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras job results bulk delete params
func (o *ExtrasJobResultsBulkDeleteParams) WithContext(ctx context.Context) *ExtrasJobResultsBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras job results bulk delete params
func (o *ExtrasJobResultsBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras job results bulk delete params
func (o *ExtrasJobResultsBulkDeleteParams) WithHTTPClient(client *http.Client) *ExtrasJobResultsBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras job results bulk delete params
func (o *ExtrasJobResultsBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasJobResultsBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

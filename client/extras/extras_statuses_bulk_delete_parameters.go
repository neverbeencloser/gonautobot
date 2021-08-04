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

// NewExtrasStatusesBulkDeleteParams creates a new ExtrasStatusesBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasStatusesBulkDeleteParams() *ExtrasStatusesBulkDeleteParams {
	return &ExtrasStatusesBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasStatusesBulkDeleteParamsWithTimeout creates a new ExtrasStatusesBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewExtrasStatusesBulkDeleteParamsWithTimeout(timeout time.Duration) *ExtrasStatusesBulkDeleteParams {
	return &ExtrasStatusesBulkDeleteParams{
		timeout: timeout,
	}
}

// NewExtrasStatusesBulkDeleteParamsWithContext creates a new ExtrasStatusesBulkDeleteParams object
// with the ability to set a context for a request.
func NewExtrasStatusesBulkDeleteParamsWithContext(ctx context.Context) *ExtrasStatusesBulkDeleteParams {
	return &ExtrasStatusesBulkDeleteParams{
		Context: ctx,
	}
}

// NewExtrasStatusesBulkDeleteParamsWithHTTPClient creates a new ExtrasStatusesBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasStatusesBulkDeleteParamsWithHTTPClient(client *http.Client) *ExtrasStatusesBulkDeleteParams {
	return &ExtrasStatusesBulkDeleteParams{
		HTTPClient: client,
	}
}

/* ExtrasStatusesBulkDeleteParams contains all the parameters to send to the API endpoint
   for the extras statuses bulk delete operation.

   Typically these are written to a http.Request.
*/
type ExtrasStatusesBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras statuses bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasStatusesBulkDeleteParams) WithDefaults() *ExtrasStatusesBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras statuses bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasStatusesBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras statuses bulk delete params
func (o *ExtrasStatusesBulkDeleteParams) WithTimeout(timeout time.Duration) *ExtrasStatusesBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras statuses bulk delete params
func (o *ExtrasStatusesBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras statuses bulk delete params
func (o *ExtrasStatusesBulkDeleteParams) WithContext(ctx context.Context) *ExtrasStatusesBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras statuses bulk delete params
func (o *ExtrasStatusesBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras statuses bulk delete params
func (o *ExtrasStatusesBulkDeleteParams) WithHTTPClient(client *http.Client) *ExtrasStatusesBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras statuses bulk delete params
func (o *ExtrasStatusesBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasStatusesBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

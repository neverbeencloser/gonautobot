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

// NewExtrasComputedFieldsBulkDeleteParams creates a new ExtrasComputedFieldsBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasComputedFieldsBulkDeleteParams() *ExtrasComputedFieldsBulkDeleteParams {
	return &ExtrasComputedFieldsBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasComputedFieldsBulkDeleteParamsWithTimeout creates a new ExtrasComputedFieldsBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewExtrasComputedFieldsBulkDeleteParamsWithTimeout(timeout time.Duration) *ExtrasComputedFieldsBulkDeleteParams {
	return &ExtrasComputedFieldsBulkDeleteParams{
		timeout: timeout,
	}
}

// NewExtrasComputedFieldsBulkDeleteParamsWithContext creates a new ExtrasComputedFieldsBulkDeleteParams object
// with the ability to set a context for a request.
func NewExtrasComputedFieldsBulkDeleteParamsWithContext(ctx context.Context) *ExtrasComputedFieldsBulkDeleteParams {
	return &ExtrasComputedFieldsBulkDeleteParams{
		Context: ctx,
	}
}

// NewExtrasComputedFieldsBulkDeleteParamsWithHTTPClient creates a new ExtrasComputedFieldsBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasComputedFieldsBulkDeleteParamsWithHTTPClient(client *http.Client) *ExtrasComputedFieldsBulkDeleteParams {
	return &ExtrasComputedFieldsBulkDeleteParams{
		HTTPClient: client,
	}
}

/* ExtrasComputedFieldsBulkDeleteParams contains all the parameters to send to the API endpoint
   for the extras computed fields bulk delete operation.

   Typically these are written to a http.Request.
*/
type ExtrasComputedFieldsBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras computed fields bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasComputedFieldsBulkDeleteParams) WithDefaults() *ExtrasComputedFieldsBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras computed fields bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasComputedFieldsBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras computed fields bulk delete params
func (o *ExtrasComputedFieldsBulkDeleteParams) WithTimeout(timeout time.Duration) *ExtrasComputedFieldsBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras computed fields bulk delete params
func (o *ExtrasComputedFieldsBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras computed fields bulk delete params
func (o *ExtrasComputedFieldsBulkDeleteParams) WithContext(ctx context.Context) *ExtrasComputedFieldsBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras computed fields bulk delete params
func (o *ExtrasComputedFieldsBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras computed fields bulk delete params
func (o *ExtrasComputedFieldsBulkDeleteParams) WithHTTPClient(client *http.Client) *ExtrasComputedFieldsBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras computed fields bulk delete params
func (o *ExtrasComputedFieldsBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasComputedFieldsBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

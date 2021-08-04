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

// NewExtrasJobsListParams creates a new ExtrasJobsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasJobsListParams() *ExtrasJobsListParams {
	return &ExtrasJobsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasJobsListParamsWithTimeout creates a new ExtrasJobsListParams object
// with the ability to set a timeout on a request.
func NewExtrasJobsListParamsWithTimeout(timeout time.Duration) *ExtrasJobsListParams {
	return &ExtrasJobsListParams{
		timeout: timeout,
	}
}

// NewExtrasJobsListParamsWithContext creates a new ExtrasJobsListParams object
// with the ability to set a context for a request.
func NewExtrasJobsListParamsWithContext(ctx context.Context) *ExtrasJobsListParams {
	return &ExtrasJobsListParams{
		Context: ctx,
	}
}

// NewExtrasJobsListParamsWithHTTPClient creates a new ExtrasJobsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasJobsListParamsWithHTTPClient(client *http.Client) *ExtrasJobsListParams {
	return &ExtrasJobsListParams{
		HTTPClient: client,
	}
}

/* ExtrasJobsListParams contains all the parameters to send to the API endpoint
   for the extras jobs list operation.

   Typically these are written to a http.Request.
*/
type ExtrasJobsListParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras jobs list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasJobsListParams) WithDefaults() *ExtrasJobsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras jobs list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasJobsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras jobs list params
func (o *ExtrasJobsListParams) WithTimeout(timeout time.Duration) *ExtrasJobsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras jobs list params
func (o *ExtrasJobsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras jobs list params
func (o *ExtrasJobsListParams) WithContext(ctx context.Context) *ExtrasJobsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras jobs list params
func (o *ExtrasJobsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras jobs list params
func (o *ExtrasJobsListParams) WithHTTPClient(client *http.Client) *ExtrasJobsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras jobs list params
func (o *ExtrasJobsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasJobsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

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

// NewExtrasJobsReadParams creates a new ExtrasJobsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasJobsReadParams() *ExtrasJobsReadParams {
	return &ExtrasJobsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasJobsReadParamsWithTimeout creates a new ExtrasJobsReadParams object
// with the ability to set a timeout on a request.
func NewExtrasJobsReadParamsWithTimeout(timeout time.Duration) *ExtrasJobsReadParams {
	return &ExtrasJobsReadParams{
		timeout: timeout,
	}
}

// NewExtrasJobsReadParamsWithContext creates a new ExtrasJobsReadParams object
// with the ability to set a context for a request.
func NewExtrasJobsReadParamsWithContext(ctx context.Context) *ExtrasJobsReadParams {
	return &ExtrasJobsReadParams{
		Context: ctx,
	}
}

// NewExtrasJobsReadParamsWithHTTPClient creates a new ExtrasJobsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasJobsReadParamsWithHTTPClient(client *http.Client) *ExtrasJobsReadParams {
	return &ExtrasJobsReadParams{
		HTTPClient: client,
	}
}

/* ExtrasJobsReadParams contains all the parameters to send to the API endpoint
   for the extras jobs read operation.

   Typically these are written to a http.Request.
*/
type ExtrasJobsReadParams struct {

	// ClassPath.
	ClassPath string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras jobs read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasJobsReadParams) WithDefaults() *ExtrasJobsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras jobs read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasJobsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras jobs read params
func (o *ExtrasJobsReadParams) WithTimeout(timeout time.Duration) *ExtrasJobsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras jobs read params
func (o *ExtrasJobsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras jobs read params
func (o *ExtrasJobsReadParams) WithContext(ctx context.Context) *ExtrasJobsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras jobs read params
func (o *ExtrasJobsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras jobs read params
func (o *ExtrasJobsReadParams) WithHTTPClient(client *http.Client) *ExtrasJobsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras jobs read params
func (o *ExtrasJobsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClassPath adds the classPath to the extras jobs read params
func (o *ExtrasJobsReadParams) WithClassPath(classPath string) *ExtrasJobsReadParams {
	o.SetClassPath(classPath)
	return o
}

// SetClassPath adds the classPath to the extras jobs read params
func (o *ExtrasJobsReadParams) SetClassPath(classPath string) {
	o.ClassPath = classPath
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasJobsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param class_path
	if err := r.SetPathParam("class_path", o.ClassPath); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

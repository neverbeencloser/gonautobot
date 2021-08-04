package extras

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// NewExtrasJobsRunParams creates a new ExtrasJobsRunParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasJobsRunParams() *ExtrasJobsRunParams {
	return &ExtrasJobsRunParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasJobsRunParamsWithTimeout creates a new ExtrasJobsRunParams object
// with the ability to set a timeout on a request.
func NewExtrasJobsRunParamsWithTimeout(timeout time.Duration) *ExtrasJobsRunParams {
	return &ExtrasJobsRunParams{
		timeout: timeout,
	}
}

// NewExtrasJobsRunParamsWithContext creates a new ExtrasJobsRunParams object
// with the ability to set a context for a request.
func NewExtrasJobsRunParamsWithContext(ctx context.Context) *ExtrasJobsRunParams {
	return &ExtrasJobsRunParams{
		Context: ctx,
	}
}

// NewExtrasJobsRunParamsWithHTTPClient creates a new ExtrasJobsRunParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasJobsRunParamsWithHTTPClient(client *http.Client) *ExtrasJobsRunParams {
	return &ExtrasJobsRunParams{
		HTTPClient: client,
	}
}

/* ExtrasJobsRunParams contains all the parameters to send to the API endpoint
   for the extras jobs run operation.

   Typically these are written to a http.Request.
*/
type ExtrasJobsRunParams struct {

	// ClassPath.
	ClassPath string

	// Data.
	Data *models.JobInput

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras jobs run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasJobsRunParams) WithDefaults() *ExtrasJobsRunParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras jobs run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasJobsRunParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras jobs run params
func (o *ExtrasJobsRunParams) WithTimeout(timeout time.Duration) *ExtrasJobsRunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras jobs run params
func (o *ExtrasJobsRunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras jobs run params
func (o *ExtrasJobsRunParams) WithContext(ctx context.Context) *ExtrasJobsRunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras jobs run params
func (o *ExtrasJobsRunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras jobs run params
func (o *ExtrasJobsRunParams) WithHTTPClient(client *http.Client) *ExtrasJobsRunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras jobs run params
func (o *ExtrasJobsRunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClassPath adds the classPath to the extras jobs run params
func (o *ExtrasJobsRunParams) WithClassPath(classPath string) *ExtrasJobsRunParams {
	o.SetClassPath(classPath)
	return o
}

// SetClassPath adds the classPath to the extras jobs run params
func (o *ExtrasJobsRunParams) SetClassPath(classPath string) {
	o.ClassPath = classPath
}

// WithData adds the data to the extras jobs run params
func (o *ExtrasJobsRunParams) WithData(data *models.JobInput) *ExtrasJobsRunParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras jobs run params
func (o *ExtrasJobsRunParams) SetData(data *models.JobInput) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasJobsRunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param class_path
	if err := r.SetPathParam("class_path", o.ClassPath); err != nil {
		return err
	}
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

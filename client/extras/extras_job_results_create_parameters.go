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

// NewExtrasJobResultsCreateParams creates a new ExtrasJobResultsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasJobResultsCreateParams() *ExtrasJobResultsCreateParams {
	return &ExtrasJobResultsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasJobResultsCreateParamsWithTimeout creates a new ExtrasJobResultsCreateParams object
// with the ability to set a timeout on a request.
func NewExtrasJobResultsCreateParamsWithTimeout(timeout time.Duration) *ExtrasJobResultsCreateParams {
	return &ExtrasJobResultsCreateParams{
		timeout: timeout,
	}
}

// NewExtrasJobResultsCreateParamsWithContext creates a new ExtrasJobResultsCreateParams object
// with the ability to set a context for a request.
func NewExtrasJobResultsCreateParamsWithContext(ctx context.Context) *ExtrasJobResultsCreateParams {
	return &ExtrasJobResultsCreateParams{
		Context: ctx,
	}
}

// NewExtrasJobResultsCreateParamsWithHTTPClient creates a new ExtrasJobResultsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasJobResultsCreateParamsWithHTTPClient(client *http.Client) *ExtrasJobResultsCreateParams {
	return &ExtrasJobResultsCreateParams{
		HTTPClient: client,
	}
}

/* ExtrasJobResultsCreateParams contains all the parameters to send to the API endpoint
   for the extras job results create operation.

   Typically these are written to a http.Request.
*/
type ExtrasJobResultsCreateParams struct {

	// Data.
	Data *models.WritableJobResult

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras job results create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasJobResultsCreateParams) WithDefaults() *ExtrasJobResultsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras job results create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasJobResultsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras job results create params
func (o *ExtrasJobResultsCreateParams) WithTimeout(timeout time.Duration) *ExtrasJobResultsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras job results create params
func (o *ExtrasJobResultsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras job results create params
func (o *ExtrasJobResultsCreateParams) WithContext(ctx context.Context) *ExtrasJobResultsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras job results create params
func (o *ExtrasJobResultsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras job results create params
func (o *ExtrasJobResultsCreateParams) WithHTTPClient(client *http.Client) *ExtrasJobResultsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras job results create params
func (o *ExtrasJobResultsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras job results create params
func (o *ExtrasJobResultsCreateParams) WithData(data *models.WritableJobResult) *ExtrasJobResultsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras job results create params
func (o *ExtrasJobResultsCreateParams) SetData(data *models.WritableJobResult) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasJobResultsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
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

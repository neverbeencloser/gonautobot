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

// NewExtrasJobResultsDeleteParams creates a new ExtrasJobResultsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasJobResultsDeleteParams() *ExtrasJobResultsDeleteParams {
	return &ExtrasJobResultsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasJobResultsDeleteParamsWithTimeout creates a new ExtrasJobResultsDeleteParams object
// with the ability to set a timeout on a request.
func NewExtrasJobResultsDeleteParamsWithTimeout(timeout time.Duration) *ExtrasJobResultsDeleteParams {
	return &ExtrasJobResultsDeleteParams{
		timeout: timeout,
	}
}

// NewExtrasJobResultsDeleteParamsWithContext creates a new ExtrasJobResultsDeleteParams object
// with the ability to set a context for a request.
func NewExtrasJobResultsDeleteParamsWithContext(ctx context.Context) *ExtrasJobResultsDeleteParams {
	return &ExtrasJobResultsDeleteParams{
		Context: ctx,
	}
}

// NewExtrasJobResultsDeleteParamsWithHTTPClient creates a new ExtrasJobResultsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasJobResultsDeleteParamsWithHTTPClient(client *http.Client) *ExtrasJobResultsDeleteParams {
	return &ExtrasJobResultsDeleteParams{
		HTTPClient: client,
	}
}

/* ExtrasJobResultsDeleteParams contains all the parameters to send to the API endpoint
   for the extras job results delete operation.

   Typically these are written to a http.Request.
*/
type ExtrasJobResultsDeleteParams struct {

	/* ID.

	   A UUID string identifying this job result.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras job results delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasJobResultsDeleteParams) WithDefaults() *ExtrasJobResultsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras job results delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasJobResultsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras job results delete params
func (o *ExtrasJobResultsDeleteParams) WithTimeout(timeout time.Duration) *ExtrasJobResultsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras job results delete params
func (o *ExtrasJobResultsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras job results delete params
func (o *ExtrasJobResultsDeleteParams) WithContext(ctx context.Context) *ExtrasJobResultsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras job results delete params
func (o *ExtrasJobResultsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras job results delete params
func (o *ExtrasJobResultsDeleteParams) WithHTTPClient(client *http.Client) *ExtrasJobResultsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras job results delete params
func (o *ExtrasJobResultsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the extras job results delete params
func (o *ExtrasJobResultsDeleteParams) WithID(id strfmt.UUID) *ExtrasJobResultsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras job results delete params
func (o *ExtrasJobResultsDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasJobResultsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

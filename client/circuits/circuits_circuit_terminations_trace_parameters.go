package circuits

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewCircuitsCircuitTerminationsTraceParams creates a new CircuitsCircuitTerminationsTraceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCircuitsCircuitTerminationsTraceParams() *CircuitsCircuitTerminationsTraceParams {
	return &CircuitsCircuitTerminationsTraceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsCircuitTerminationsTraceParamsWithTimeout creates a new CircuitsCircuitTerminationsTraceParams object
// with the ability to set a timeout on a request.
func NewCircuitsCircuitTerminationsTraceParamsWithTimeout(timeout time.Duration) *CircuitsCircuitTerminationsTraceParams {
	return &CircuitsCircuitTerminationsTraceParams{
		timeout: timeout,
	}
}

// NewCircuitsCircuitTerminationsTraceParamsWithContext creates a new CircuitsCircuitTerminationsTraceParams object
// with the ability to set a context for a request.
func NewCircuitsCircuitTerminationsTraceParamsWithContext(ctx context.Context) *CircuitsCircuitTerminationsTraceParams {
	return &CircuitsCircuitTerminationsTraceParams{
		Context: ctx,
	}
}

// NewCircuitsCircuitTerminationsTraceParamsWithHTTPClient creates a new CircuitsCircuitTerminationsTraceParams object
// with the ability to set a custom HTTPClient for a request.
func NewCircuitsCircuitTerminationsTraceParamsWithHTTPClient(client *http.Client) *CircuitsCircuitTerminationsTraceParams {
	return &CircuitsCircuitTerminationsTraceParams{
		HTTPClient: client,
	}
}

/* CircuitsCircuitTerminationsTraceParams contains all the parameters to send to the API endpoint
   for the circuits circuit terminations trace operation.

   Typically these are written to a http.Request.
*/
type CircuitsCircuitTerminationsTraceParams struct {

	/* ID.

	   A UUID string identifying this circuit termination.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the circuits circuit terminations trace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsCircuitTerminationsTraceParams) WithDefaults() *CircuitsCircuitTerminationsTraceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the circuits circuit terminations trace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsCircuitTerminationsTraceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the circuits circuit terminations trace params
func (o *CircuitsCircuitTerminationsTraceParams) WithTimeout(timeout time.Duration) *CircuitsCircuitTerminationsTraceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits circuit terminations trace params
func (o *CircuitsCircuitTerminationsTraceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits circuit terminations trace params
func (o *CircuitsCircuitTerminationsTraceParams) WithContext(ctx context.Context) *CircuitsCircuitTerminationsTraceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits circuit terminations trace params
func (o *CircuitsCircuitTerminationsTraceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits circuit terminations trace params
func (o *CircuitsCircuitTerminationsTraceParams) WithHTTPClient(client *http.Client) *CircuitsCircuitTerminationsTraceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits circuit terminations trace params
func (o *CircuitsCircuitTerminationsTraceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the circuits circuit terminations trace params
func (o *CircuitsCircuitTerminationsTraceParams) WithID(id strfmt.UUID) *CircuitsCircuitTerminationsTraceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the circuits circuit terminations trace params
func (o *CircuitsCircuitTerminationsTraceParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsCircuitTerminationsTraceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

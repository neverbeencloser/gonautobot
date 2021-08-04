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

// NewCircuitsProvidersDeleteParams creates a new CircuitsProvidersDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCircuitsProvidersDeleteParams() *CircuitsProvidersDeleteParams {
	return &CircuitsProvidersDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsProvidersDeleteParamsWithTimeout creates a new CircuitsProvidersDeleteParams object
// with the ability to set a timeout on a request.
func NewCircuitsProvidersDeleteParamsWithTimeout(timeout time.Duration) *CircuitsProvidersDeleteParams {
	return &CircuitsProvidersDeleteParams{
		timeout: timeout,
	}
}

// NewCircuitsProvidersDeleteParamsWithContext creates a new CircuitsProvidersDeleteParams object
// with the ability to set a context for a request.
func NewCircuitsProvidersDeleteParamsWithContext(ctx context.Context) *CircuitsProvidersDeleteParams {
	return &CircuitsProvidersDeleteParams{
		Context: ctx,
	}
}

// NewCircuitsProvidersDeleteParamsWithHTTPClient creates a new CircuitsProvidersDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewCircuitsProvidersDeleteParamsWithHTTPClient(client *http.Client) *CircuitsProvidersDeleteParams {
	return &CircuitsProvidersDeleteParams{
		HTTPClient: client,
	}
}

/* CircuitsProvidersDeleteParams contains all the parameters to send to the API endpoint
   for the circuits providers delete operation.

   Typically these are written to a http.Request.
*/
type CircuitsProvidersDeleteParams struct {

	/* ID.

	   A UUID string identifying this provider.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the circuits providers delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsProvidersDeleteParams) WithDefaults() *CircuitsProvidersDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the circuits providers delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsProvidersDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the circuits providers delete params
func (o *CircuitsProvidersDeleteParams) WithTimeout(timeout time.Duration) *CircuitsProvidersDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits providers delete params
func (o *CircuitsProvidersDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits providers delete params
func (o *CircuitsProvidersDeleteParams) WithContext(ctx context.Context) *CircuitsProvidersDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits providers delete params
func (o *CircuitsProvidersDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits providers delete params
func (o *CircuitsProvidersDeleteParams) WithHTTPClient(client *http.Client) *CircuitsProvidersDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits providers delete params
func (o *CircuitsProvidersDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the circuits providers delete params
func (o *CircuitsProvidersDeleteParams) WithID(id strfmt.UUID) *CircuitsProvidersDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the circuits providers delete params
func (o *CircuitsProvidersDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsProvidersDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

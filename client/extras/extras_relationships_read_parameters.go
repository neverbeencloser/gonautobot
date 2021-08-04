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

// NewExtrasRelationshipsReadParams creates a new ExtrasRelationshipsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasRelationshipsReadParams() *ExtrasRelationshipsReadParams {
	return &ExtrasRelationshipsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasRelationshipsReadParamsWithTimeout creates a new ExtrasRelationshipsReadParams object
// with the ability to set a timeout on a request.
func NewExtrasRelationshipsReadParamsWithTimeout(timeout time.Duration) *ExtrasRelationshipsReadParams {
	return &ExtrasRelationshipsReadParams{
		timeout: timeout,
	}
}

// NewExtrasRelationshipsReadParamsWithContext creates a new ExtrasRelationshipsReadParams object
// with the ability to set a context for a request.
func NewExtrasRelationshipsReadParamsWithContext(ctx context.Context) *ExtrasRelationshipsReadParams {
	return &ExtrasRelationshipsReadParams{
		Context: ctx,
	}
}

// NewExtrasRelationshipsReadParamsWithHTTPClient creates a new ExtrasRelationshipsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasRelationshipsReadParamsWithHTTPClient(client *http.Client) *ExtrasRelationshipsReadParams {
	return &ExtrasRelationshipsReadParams{
		HTTPClient: client,
	}
}

/* ExtrasRelationshipsReadParams contains all the parameters to send to the API endpoint
   for the extras relationships read operation.

   Typically these are written to a http.Request.
*/
type ExtrasRelationshipsReadParams struct {

	/* ID.

	   A UUID string identifying this relationship.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras relationships read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasRelationshipsReadParams) WithDefaults() *ExtrasRelationshipsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras relationships read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasRelationshipsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras relationships read params
func (o *ExtrasRelationshipsReadParams) WithTimeout(timeout time.Duration) *ExtrasRelationshipsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras relationships read params
func (o *ExtrasRelationshipsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras relationships read params
func (o *ExtrasRelationshipsReadParams) WithContext(ctx context.Context) *ExtrasRelationshipsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras relationships read params
func (o *ExtrasRelationshipsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras relationships read params
func (o *ExtrasRelationshipsReadParams) WithHTTPClient(client *http.Client) *ExtrasRelationshipsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras relationships read params
func (o *ExtrasRelationshipsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the extras relationships read params
func (o *ExtrasRelationshipsReadParams) WithID(id strfmt.UUID) *ExtrasRelationshipsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras relationships read params
func (o *ExtrasRelationshipsReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasRelationshipsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

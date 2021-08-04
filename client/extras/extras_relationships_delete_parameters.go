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

// NewExtrasRelationshipsDeleteParams creates a new ExtrasRelationshipsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasRelationshipsDeleteParams() *ExtrasRelationshipsDeleteParams {
	return &ExtrasRelationshipsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasRelationshipsDeleteParamsWithTimeout creates a new ExtrasRelationshipsDeleteParams object
// with the ability to set a timeout on a request.
func NewExtrasRelationshipsDeleteParamsWithTimeout(timeout time.Duration) *ExtrasRelationshipsDeleteParams {
	return &ExtrasRelationshipsDeleteParams{
		timeout: timeout,
	}
}

// NewExtrasRelationshipsDeleteParamsWithContext creates a new ExtrasRelationshipsDeleteParams object
// with the ability to set a context for a request.
func NewExtrasRelationshipsDeleteParamsWithContext(ctx context.Context) *ExtrasRelationshipsDeleteParams {
	return &ExtrasRelationshipsDeleteParams{
		Context: ctx,
	}
}

// NewExtrasRelationshipsDeleteParamsWithHTTPClient creates a new ExtrasRelationshipsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasRelationshipsDeleteParamsWithHTTPClient(client *http.Client) *ExtrasRelationshipsDeleteParams {
	return &ExtrasRelationshipsDeleteParams{
		HTTPClient: client,
	}
}

/* ExtrasRelationshipsDeleteParams contains all the parameters to send to the API endpoint
   for the extras relationships delete operation.

   Typically these are written to a http.Request.
*/
type ExtrasRelationshipsDeleteParams struct {

	/* ID.

	   A UUID string identifying this relationship.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras relationships delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasRelationshipsDeleteParams) WithDefaults() *ExtrasRelationshipsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras relationships delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasRelationshipsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras relationships delete params
func (o *ExtrasRelationshipsDeleteParams) WithTimeout(timeout time.Duration) *ExtrasRelationshipsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras relationships delete params
func (o *ExtrasRelationshipsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras relationships delete params
func (o *ExtrasRelationshipsDeleteParams) WithContext(ctx context.Context) *ExtrasRelationshipsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras relationships delete params
func (o *ExtrasRelationshipsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras relationships delete params
func (o *ExtrasRelationshipsDeleteParams) WithHTTPClient(client *http.Client) *ExtrasRelationshipsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras relationships delete params
func (o *ExtrasRelationshipsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the extras relationships delete params
func (o *ExtrasRelationshipsDeleteParams) WithID(id strfmt.UUID) *ExtrasRelationshipsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras relationships delete params
func (o *ExtrasRelationshipsDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasRelationshipsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

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

// NewExtrasRelationshipAssociationsReadParams creates a new ExtrasRelationshipAssociationsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasRelationshipAssociationsReadParams() *ExtrasRelationshipAssociationsReadParams {
	return &ExtrasRelationshipAssociationsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasRelationshipAssociationsReadParamsWithTimeout creates a new ExtrasRelationshipAssociationsReadParams object
// with the ability to set a timeout on a request.
func NewExtrasRelationshipAssociationsReadParamsWithTimeout(timeout time.Duration) *ExtrasRelationshipAssociationsReadParams {
	return &ExtrasRelationshipAssociationsReadParams{
		timeout: timeout,
	}
}

// NewExtrasRelationshipAssociationsReadParamsWithContext creates a new ExtrasRelationshipAssociationsReadParams object
// with the ability to set a context for a request.
func NewExtrasRelationshipAssociationsReadParamsWithContext(ctx context.Context) *ExtrasRelationshipAssociationsReadParams {
	return &ExtrasRelationshipAssociationsReadParams{
		Context: ctx,
	}
}

// NewExtrasRelationshipAssociationsReadParamsWithHTTPClient creates a new ExtrasRelationshipAssociationsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasRelationshipAssociationsReadParamsWithHTTPClient(client *http.Client) *ExtrasRelationshipAssociationsReadParams {
	return &ExtrasRelationshipAssociationsReadParams{
		HTTPClient: client,
	}
}

/* ExtrasRelationshipAssociationsReadParams contains all the parameters to send to the API endpoint
   for the extras relationship associations read operation.

   Typically these are written to a http.Request.
*/
type ExtrasRelationshipAssociationsReadParams struct {

	/* ID.

	   A UUID string identifying this relationship association.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras relationship associations read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasRelationshipAssociationsReadParams) WithDefaults() *ExtrasRelationshipAssociationsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras relationship associations read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasRelationshipAssociationsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras relationship associations read params
func (o *ExtrasRelationshipAssociationsReadParams) WithTimeout(timeout time.Duration) *ExtrasRelationshipAssociationsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras relationship associations read params
func (o *ExtrasRelationshipAssociationsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras relationship associations read params
func (o *ExtrasRelationshipAssociationsReadParams) WithContext(ctx context.Context) *ExtrasRelationshipAssociationsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras relationship associations read params
func (o *ExtrasRelationshipAssociationsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras relationship associations read params
func (o *ExtrasRelationshipAssociationsReadParams) WithHTTPClient(client *http.Client) *ExtrasRelationshipAssociationsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras relationship associations read params
func (o *ExtrasRelationshipAssociationsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the extras relationship associations read params
func (o *ExtrasRelationshipAssociationsReadParams) WithID(id strfmt.UUID) *ExtrasRelationshipAssociationsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras relationship associations read params
func (o *ExtrasRelationshipAssociationsReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasRelationshipAssociationsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

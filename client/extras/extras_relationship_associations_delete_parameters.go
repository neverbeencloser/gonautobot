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

// NewExtrasRelationshipAssociationsDeleteParams creates a new ExtrasRelationshipAssociationsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasRelationshipAssociationsDeleteParams() *ExtrasRelationshipAssociationsDeleteParams {
	return &ExtrasRelationshipAssociationsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasRelationshipAssociationsDeleteParamsWithTimeout creates a new ExtrasRelationshipAssociationsDeleteParams object
// with the ability to set a timeout on a request.
func NewExtrasRelationshipAssociationsDeleteParamsWithTimeout(timeout time.Duration) *ExtrasRelationshipAssociationsDeleteParams {
	return &ExtrasRelationshipAssociationsDeleteParams{
		timeout: timeout,
	}
}

// NewExtrasRelationshipAssociationsDeleteParamsWithContext creates a new ExtrasRelationshipAssociationsDeleteParams object
// with the ability to set a context for a request.
func NewExtrasRelationshipAssociationsDeleteParamsWithContext(ctx context.Context) *ExtrasRelationshipAssociationsDeleteParams {
	return &ExtrasRelationshipAssociationsDeleteParams{
		Context: ctx,
	}
}

// NewExtrasRelationshipAssociationsDeleteParamsWithHTTPClient creates a new ExtrasRelationshipAssociationsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasRelationshipAssociationsDeleteParamsWithHTTPClient(client *http.Client) *ExtrasRelationshipAssociationsDeleteParams {
	return &ExtrasRelationshipAssociationsDeleteParams{
		HTTPClient: client,
	}
}

/* ExtrasRelationshipAssociationsDeleteParams contains all the parameters to send to the API endpoint
   for the extras relationship associations delete operation.

   Typically these are written to a http.Request.
*/
type ExtrasRelationshipAssociationsDeleteParams struct {

	/* ID.

	   A UUID string identifying this relationship association.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras relationship associations delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasRelationshipAssociationsDeleteParams) WithDefaults() *ExtrasRelationshipAssociationsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras relationship associations delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasRelationshipAssociationsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras relationship associations delete params
func (o *ExtrasRelationshipAssociationsDeleteParams) WithTimeout(timeout time.Duration) *ExtrasRelationshipAssociationsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras relationship associations delete params
func (o *ExtrasRelationshipAssociationsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras relationship associations delete params
func (o *ExtrasRelationshipAssociationsDeleteParams) WithContext(ctx context.Context) *ExtrasRelationshipAssociationsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras relationship associations delete params
func (o *ExtrasRelationshipAssociationsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras relationship associations delete params
func (o *ExtrasRelationshipAssociationsDeleteParams) WithHTTPClient(client *http.Client) *ExtrasRelationshipAssociationsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras relationship associations delete params
func (o *ExtrasRelationshipAssociationsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the extras relationship associations delete params
func (o *ExtrasRelationshipAssociationsDeleteParams) WithID(id strfmt.UUID) *ExtrasRelationshipAssociationsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras relationship associations delete params
func (o *ExtrasRelationshipAssociationsDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasRelationshipAssociationsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

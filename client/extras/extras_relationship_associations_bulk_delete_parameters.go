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

// NewExtrasRelationshipAssociationsBulkDeleteParams creates a new ExtrasRelationshipAssociationsBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasRelationshipAssociationsBulkDeleteParams() *ExtrasRelationshipAssociationsBulkDeleteParams {
	return &ExtrasRelationshipAssociationsBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasRelationshipAssociationsBulkDeleteParamsWithTimeout creates a new ExtrasRelationshipAssociationsBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewExtrasRelationshipAssociationsBulkDeleteParamsWithTimeout(timeout time.Duration) *ExtrasRelationshipAssociationsBulkDeleteParams {
	return &ExtrasRelationshipAssociationsBulkDeleteParams{
		timeout: timeout,
	}
}

// NewExtrasRelationshipAssociationsBulkDeleteParamsWithContext creates a new ExtrasRelationshipAssociationsBulkDeleteParams object
// with the ability to set a context for a request.
func NewExtrasRelationshipAssociationsBulkDeleteParamsWithContext(ctx context.Context) *ExtrasRelationshipAssociationsBulkDeleteParams {
	return &ExtrasRelationshipAssociationsBulkDeleteParams{
		Context: ctx,
	}
}

// NewExtrasRelationshipAssociationsBulkDeleteParamsWithHTTPClient creates a new ExtrasRelationshipAssociationsBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasRelationshipAssociationsBulkDeleteParamsWithHTTPClient(client *http.Client) *ExtrasRelationshipAssociationsBulkDeleteParams {
	return &ExtrasRelationshipAssociationsBulkDeleteParams{
		HTTPClient: client,
	}
}

/* ExtrasRelationshipAssociationsBulkDeleteParams contains all the parameters to send to the API endpoint
   for the extras relationship associations bulk delete operation.

   Typically these are written to a http.Request.
*/
type ExtrasRelationshipAssociationsBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras relationship associations bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasRelationshipAssociationsBulkDeleteParams) WithDefaults() *ExtrasRelationshipAssociationsBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras relationship associations bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasRelationshipAssociationsBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras relationship associations bulk delete params
func (o *ExtrasRelationshipAssociationsBulkDeleteParams) WithTimeout(timeout time.Duration) *ExtrasRelationshipAssociationsBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras relationship associations bulk delete params
func (o *ExtrasRelationshipAssociationsBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras relationship associations bulk delete params
func (o *ExtrasRelationshipAssociationsBulkDeleteParams) WithContext(ctx context.Context) *ExtrasRelationshipAssociationsBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras relationship associations bulk delete params
func (o *ExtrasRelationshipAssociationsBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras relationship associations bulk delete params
func (o *ExtrasRelationshipAssociationsBulkDeleteParams) WithHTTPClient(client *http.Client) *ExtrasRelationshipAssociationsBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras relationship associations bulk delete params
func (o *ExtrasRelationshipAssociationsBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasRelationshipAssociationsBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

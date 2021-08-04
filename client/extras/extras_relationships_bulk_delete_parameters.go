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

// NewExtrasRelationshipsBulkDeleteParams creates a new ExtrasRelationshipsBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasRelationshipsBulkDeleteParams() *ExtrasRelationshipsBulkDeleteParams {
	return &ExtrasRelationshipsBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasRelationshipsBulkDeleteParamsWithTimeout creates a new ExtrasRelationshipsBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewExtrasRelationshipsBulkDeleteParamsWithTimeout(timeout time.Duration) *ExtrasRelationshipsBulkDeleteParams {
	return &ExtrasRelationshipsBulkDeleteParams{
		timeout: timeout,
	}
}

// NewExtrasRelationshipsBulkDeleteParamsWithContext creates a new ExtrasRelationshipsBulkDeleteParams object
// with the ability to set a context for a request.
func NewExtrasRelationshipsBulkDeleteParamsWithContext(ctx context.Context) *ExtrasRelationshipsBulkDeleteParams {
	return &ExtrasRelationshipsBulkDeleteParams{
		Context: ctx,
	}
}

// NewExtrasRelationshipsBulkDeleteParamsWithHTTPClient creates a new ExtrasRelationshipsBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasRelationshipsBulkDeleteParamsWithHTTPClient(client *http.Client) *ExtrasRelationshipsBulkDeleteParams {
	return &ExtrasRelationshipsBulkDeleteParams{
		HTTPClient: client,
	}
}

/* ExtrasRelationshipsBulkDeleteParams contains all the parameters to send to the API endpoint
   for the extras relationships bulk delete operation.

   Typically these are written to a http.Request.
*/
type ExtrasRelationshipsBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras relationships bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasRelationshipsBulkDeleteParams) WithDefaults() *ExtrasRelationshipsBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras relationships bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasRelationshipsBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras relationships bulk delete params
func (o *ExtrasRelationshipsBulkDeleteParams) WithTimeout(timeout time.Duration) *ExtrasRelationshipsBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras relationships bulk delete params
func (o *ExtrasRelationshipsBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras relationships bulk delete params
func (o *ExtrasRelationshipsBulkDeleteParams) WithContext(ctx context.Context) *ExtrasRelationshipsBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras relationships bulk delete params
func (o *ExtrasRelationshipsBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras relationships bulk delete params
func (o *ExtrasRelationshipsBulkDeleteParams) WithHTTPClient(client *http.Client) *ExtrasRelationshipsBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras relationships bulk delete params
func (o *ExtrasRelationshipsBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasRelationshipsBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

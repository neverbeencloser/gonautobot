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

// NewExtrasGraphqlQueriesBulkDeleteParams creates a new ExtrasGraphqlQueriesBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasGraphqlQueriesBulkDeleteParams() *ExtrasGraphqlQueriesBulkDeleteParams {
	return &ExtrasGraphqlQueriesBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasGraphqlQueriesBulkDeleteParamsWithTimeout creates a new ExtrasGraphqlQueriesBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewExtrasGraphqlQueriesBulkDeleteParamsWithTimeout(timeout time.Duration) *ExtrasGraphqlQueriesBulkDeleteParams {
	return &ExtrasGraphqlQueriesBulkDeleteParams{
		timeout: timeout,
	}
}

// NewExtrasGraphqlQueriesBulkDeleteParamsWithContext creates a new ExtrasGraphqlQueriesBulkDeleteParams object
// with the ability to set a context for a request.
func NewExtrasGraphqlQueriesBulkDeleteParamsWithContext(ctx context.Context) *ExtrasGraphqlQueriesBulkDeleteParams {
	return &ExtrasGraphqlQueriesBulkDeleteParams{
		Context: ctx,
	}
}

// NewExtrasGraphqlQueriesBulkDeleteParamsWithHTTPClient creates a new ExtrasGraphqlQueriesBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasGraphqlQueriesBulkDeleteParamsWithHTTPClient(client *http.Client) *ExtrasGraphqlQueriesBulkDeleteParams {
	return &ExtrasGraphqlQueriesBulkDeleteParams{
		HTTPClient: client,
	}
}

/* ExtrasGraphqlQueriesBulkDeleteParams contains all the parameters to send to the API endpoint
   for the extras graphql queries bulk delete operation.

   Typically these are written to a http.Request.
*/
type ExtrasGraphqlQueriesBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras graphql queries bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasGraphqlQueriesBulkDeleteParams) WithDefaults() *ExtrasGraphqlQueriesBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras graphql queries bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasGraphqlQueriesBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras graphql queries bulk delete params
func (o *ExtrasGraphqlQueriesBulkDeleteParams) WithTimeout(timeout time.Duration) *ExtrasGraphqlQueriesBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras graphql queries bulk delete params
func (o *ExtrasGraphqlQueriesBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras graphql queries bulk delete params
func (o *ExtrasGraphqlQueriesBulkDeleteParams) WithContext(ctx context.Context) *ExtrasGraphqlQueriesBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras graphql queries bulk delete params
func (o *ExtrasGraphqlQueriesBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras graphql queries bulk delete params
func (o *ExtrasGraphqlQueriesBulkDeleteParams) WithHTTPClient(client *http.Client) *ExtrasGraphqlQueriesBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras graphql queries bulk delete params
func (o *ExtrasGraphqlQueriesBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasGraphqlQueriesBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

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

// NewExtrasGraphqlQueriesDeleteParams creates a new ExtrasGraphqlQueriesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasGraphqlQueriesDeleteParams() *ExtrasGraphqlQueriesDeleteParams {
	return &ExtrasGraphqlQueriesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasGraphqlQueriesDeleteParamsWithTimeout creates a new ExtrasGraphqlQueriesDeleteParams object
// with the ability to set a timeout on a request.
func NewExtrasGraphqlQueriesDeleteParamsWithTimeout(timeout time.Duration) *ExtrasGraphqlQueriesDeleteParams {
	return &ExtrasGraphqlQueriesDeleteParams{
		timeout: timeout,
	}
}

// NewExtrasGraphqlQueriesDeleteParamsWithContext creates a new ExtrasGraphqlQueriesDeleteParams object
// with the ability to set a context for a request.
func NewExtrasGraphqlQueriesDeleteParamsWithContext(ctx context.Context) *ExtrasGraphqlQueriesDeleteParams {
	return &ExtrasGraphqlQueriesDeleteParams{
		Context: ctx,
	}
}

// NewExtrasGraphqlQueriesDeleteParamsWithHTTPClient creates a new ExtrasGraphqlQueriesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasGraphqlQueriesDeleteParamsWithHTTPClient(client *http.Client) *ExtrasGraphqlQueriesDeleteParams {
	return &ExtrasGraphqlQueriesDeleteParams{
		HTTPClient: client,
	}
}

/* ExtrasGraphqlQueriesDeleteParams contains all the parameters to send to the API endpoint
   for the extras graphql queries delete operation.

   Typically these are written to a http.Request.
*/
type ExtrasGraphqlQueriesDeleteParams struct {

	/* ID.

	   A UUID string identifying this GraphQL query.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras graphql queries delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasGraphqlQueriesDeleteParams) WithDefaults() *ExtrasGraphqlQueriesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras graphql queries delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasGraphqlQueriesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras graphql queries delete params
func (o *ExtrasGraphqlQueriesDeleteParams) WithTimeout(timeout time.Duration) *ExtrasGraphqlQueriesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras graphql queries delete params
func (o *ExtrasGraphqlQueriesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras graphql queries delete params
func (o *ExtrasGraphqlQueriesDeleteParams) WithContext(ctx context.Context) *ExtrasGraphqlQueriesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras graphql queries delete params
func (o *ExtrasGraphqlQueriesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras graphql queries delete params
func (o *ExtrasGraphqlQueriesDeleteParams) WithHTTPClient(client *http.Client) *ExtrasGraphqlQueriesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras graphql queries delete params
func (o *ExtrasGraphqlQueriesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the extras graphql queries delete params
func (o *ExtrasGraphqlQueriesDeleteParams) WithID(id strfmt.UUID) *ExtrasGraphqlQueriesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras graphql queries delete params
func (o *ExtrasGraphqlQueriesDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasGraphqlQueriesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

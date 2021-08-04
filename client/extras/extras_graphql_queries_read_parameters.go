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

// NewExtrasGraphqlQueriesReadParams creates a new ExtrasGraphqlQueriesReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasGraphqlQueriesReadParams() *ExtrasGraphqlQueriesReadParams {
	return &ExtrasGraphqlQueriesReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasGraphqlQueriesReadParamsWithTimeout creates a new ExtrasGraphqlQueriesReadParams object
// with the ability to set a timeout on a request.
func NewExtrasGraphqlQueriesReadParamsWithTimeout(timeout time.Duration) *ExtrasGraphqlQueriesReadParams {
	return &ExtrasGraphqlQueriesReadParams{
		timeout: timeout,
	}
}

// NewExtrasGraphqlQueriesReadParamsWithContext creates a new ExtrasGraphqlQueriesReadParams object
// with the ability to set a context for a request.
func NewExtrasGraphqlQueriesReadParamsWithContext(ctx context.Context) *ExtrasGraphqlQueriesReadParams {
	return &ExtrasGraphqlQueriesReadParams{
		Context: ctx,
	}
}

// NewExtrasGraphqlQueriesReadParamsWithHTTPClient creates a new ExtrasGraphqlQueriesReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasGraphqlQueriesReadParamsWithHTTPClient(client *http.Client) *ExtrasGraphqlQueriesReadParams {
	return &ExtrasGraphqlQueriesReadParams{
		HTTPClient: client,
	}
}

/* ExtrasGraphqlQueriesReadParams contains all the parameters to send to the API endpoint
   for the extras graphql queries read operation.

   Typically these are written to a http.Request.
*/
type ExtrasGraphqlQueriesReadParams struct {

	/* ID.

	   A UUID string identifying this GraphQL query.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras graphql queries read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasGraphqlQueriesReadParams) WithDefaults() *ExtrasGraphqlQueriesReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras graphql queries read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasGraphqlQueriesReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras graphql queries read params
func (o *ExtrasGraphqlQueriesReadParams) WithTimeout(timeout time.Duration) *ExtrasGraphqlQueriesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras graphql queries read params
func (o *ExtrasGraphqlQueriesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras graphql queries read params
func (o *ExtrasGraphqlQueriesReadParams) WithContext(ctx context.Context) *ExtrasGraphqlQueriesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras graphql queries read params
func (o *ExtrasGraphqlQueriesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras graphql queries read params
func (o *ExtrasGraphqlQueriesReadParams) WithHTTPClient(client *http.Client) *ExtrasGraphqlQueriesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras graphql queries read params
func (o *ExtrasGraphqlQueriesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the extras graphql queries read params
func (o *ExtrasGraphqlQueriesReadParams) WithID(id strfmt.UUID) *ExtrasGraphqlQueriesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras graphql queries read params
func (o *ExtrasGraphqlQueriesReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasGraphqlQueriesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

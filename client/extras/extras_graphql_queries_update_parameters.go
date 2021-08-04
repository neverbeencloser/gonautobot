package extras

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// NewExtrasGraphqlQueriesUpdateParams creates a new ExtrasGraphqlQueriesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasGraphqlQueriesUpdateParams() *ExtrasGraphqlQueriesUpdateParams {
	return &ExtrasGraphqlQueriesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasGraphqlQueriesUpdateParamsWithTimeout creates a new ExtrasGraphqlQueriesUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasGraphqlQueriesUpdateParamsWithTimeout(timeout time.Duration) *ExtrasGraphqlQueriesUpdateParams {
	return &ExtrasGraphqlQueriesUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasGraphqlQueriesUpdateParamsWithContext creates a new ExtrasGraphqlQueriesUpdateParams object
// with the ability to set a context for a request.
func NewExtrasGraphqlQueriesUpdateParamsWithContext(ctx context.Context) *ExtrasGraphqlQueriesUpdateParams {
	return &ExtrasGraphqlQueriesUpdateParams{
		Context: ctx,
	}
}

// NewExtrasGraphqlQueriesUpdateParamsWithHTTPClient creates a new ExtrasGraphqlQueriesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasGraphqlQueriesUpdateParamsWithHTTPClient(client *http.Client) *ExtrasGraphqlQueriesUpdateParams {
	return &ExtrasGraphqlQueriesUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasGraphqlQueriesUpdateParams contains all the parameters to send to the API endpoint
   for the extras graphql queries update operation.

   Typically these are written to a http.Request.
*/
type ExtrasGraphqlQueriesUpdateParams struct {

	// Data.
	Data *models.GraphQLQuery

	/* ID.

	   A UUID string identifying this GraphQL query.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras graphql queries update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasGraphqlQueriesUpdateParams) WithDefaults() *ExtrasGraphqlQueriesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras graphql queries update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasGraphqlQueriesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras graphql queries update params
func (o *ExtrasGraphqlQueriesUpdateParams) WithTimeout(timeout time.Duration) *ExtrasGraphqlQueriesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras graphql queries update params
func (o *ExtrasGraphqlQueriesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras graphql queries update params
func (o *ExtrasGraphqlQueriesUpdateParams) WithContext(ctx context.Context) *ExtrasGraphqlQueriesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras graphql queries update params
func (o *ExtrasGraphqlQueriesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras graphql queries update params
func (o *ExtrasGraphqlQueriesUpdateParams) WithHTTPClient(client *http.Client) *ExtrasGraphqlQueriesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras graphql queries update params
func (o *ExtrasGraphqlQueriesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras graphql queries update params
func (o *ExtrasGraphqlQueriesUpdateParams) WithData(data *models.GraphQLQuery) *ExtrasGraphqlQueriesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras graphql queries update params
func (o *ExtrasGraphqlQueriesUpdateParams) SetData(data *models.GraphQLQuery) {
	o.Data = data
}

// WithID adds the id to the extras graphql queries update params
func (o *ExtrasGraphqlQueriesUpdateParams) WithID(id strfmt.UUID) *ExtrasGraphqlQueriesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras graphql queries update params
func (o *ExtrasGraphqlQueriesUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasGraphqlQueriesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

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

// NewExtrasGraphqlQueriesPartialUpdateParams creates a new ExtrasGraphqlQueriesPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasGraphqlQueriesPartialUpdateParams() *ExtrasGraphqlQueriesPartialUpdateParams {
	return &ExtrasGraphqlQueriesPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasGraphqlQueriesPartialUpdateParamsWithTimeout creates a new ExtrasGraphqlQueriesPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasGraphqlQueriesPartialUpdateParamsWithTimeout(timeout time.Duration) *ExtrasGraphqlQueriesPartialUpdateParams {
	return &ExtrasGraphqlQueriesPartialUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasGraphqlQueriesPartialUpdateParamsWithContext creates a new ExtrasGraphqlQueriesPartialUpdateParams object
// with the ability to set a context for a request.
func NewExtrasGraphqlQueriesPartialUpdateParamsWithContext(ctx context.Context) *ExtrasGraphqlQueriesPartialUpdateParams {
	return &ExtrasGraphqlQueriesPartialUpdateParams{
		Context: ctx,
	}
}

// NewExtrasGraphqlQueriesPartialUpdateParamsWithHTTPClient creates a new ExtrasGraphqlQueriesPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasGraphqlQueriesPartialUpdateParamsWithHTTPClient(client *http.Client) *ExtrasGraphqlQueriesPartialUpdateParams {
	return &ExtrasGraphqlQueriesPartialUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasGraphqlQueriesPartialUpdateParams contains all the parameters to send to the API endpoint
   for the extras graphql queries partial update operation.

   Typically these are written to a http.Request.
*/
type ExtrasGraphqlQueriesPartialUpdateParams struct {

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

// WithDefaults hydrates default values in the extras graphql queries partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasGraphqlQueriesPartialUpdateParams) WithDefaults() *ExtrasGraphqlQueriesPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras graphql queries partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasGraphqlQueriesPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras graphql queries partial update params
func (o *ExtrasGraphqlQueriesPartialUpdateParams) WithTimeout(timeout time.Duration) *ExtrasGraphqlQueriesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras graphql queries partial update params
func (o *ExtrasGraphqlQueriesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras graphql queries partial update params
func (o *ExtrasGraphqlQueriesPartialUpdateParams) WithContext(ctx context.Context) *ExtrasGraphqlQueriesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras graphql queries partial update params
func (o *ExtrasGraphqlQueriesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras graphql queries partial update params
func (o *ExtrasGraphqlQueriesPartialUpdateParams) WithHTTPClient(client *http.Client) *ExtrasGraphqlQueriesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras graphql queries partial update params
func (o *ExtrasGraphqlQueriesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras graphql queries partial update params
func (o *ExtrasGraphqlQueriesPartialUpdateParams) WithData(data *models.GraphQLQuery) *ExtrasGraphqlQueriesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras graphql queries partial update params
func (o *ExtrasGraphqlQueriesPartialUpdateParams) SetData(data *models.GraphQLQuery) {
	o.Data = data
}

// WithID adds the id to the extras graphql queries partial update params
func (o *ExtrasGraphqlQueriesPartialUpdateParams) WithID(id strfmt.UUID) *ExtrasGraphqlQueriesPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras graphql queries partial update params
func (o *ExtrasGraphqlQueriesPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasGraphqlQueriesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

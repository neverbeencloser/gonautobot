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

// NewExtrasGraphqlQueriesRunParams creates a new ExtrasGraphqlQueriesRunParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasGraphqlQueriesRunParams() *ExtrasGraphqlQueriesRunParams {
	return &ExtrasGraphqlQueriesRunParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasGraphqlQueriesRunParamsWithTimeout creates a new ExtrasGraphqlQueriesRunParams object
// with the ability to set a timeout on a request.
func NewExtrasGraphqlQueriesRunParamsWithTimeout(timeout time.Duration) *ExtrasGraphqlQueriesRunParams {
	return &ExtrasGraphqlQueriesRunParams{
		timeout: timeout,
	}
}

// NewExtrasGraphqlQueriesRunParamsWithContext creates a new ExtrasGraphqlQueriesRunParams object
// with the ability to set a context for a request.
func NewExtrasGraphqlQueriesRunParamsWithContext(ctx context.Context) *ExtrasGraphqlQueriesRunParams {
	return &ExtrasGraphqlQueriesRunParams{
		Context: ctx,
	}
}

// NewExtrasGraphqlQueriesRunParamsWithHTTPClient creates a new ExtrasGraphqlQueriesRunParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasGraphqlQueriesRunParamsWithHTTPClient(client *http.Client) *ExtrasGraphqlQueriesRunParams {
	return &ExtrasGraphqlQueriesRunParams{
		HTTPClient: client,
	}
}

/* ExtrasGraphqlQueriesRunParams contains all the parameters to send to the API endpoint
   for the extras graphql queries run operation.

   Typically these are written to a http.Request.
*/
type ExtrasGraphqlQueriesRunParams struct {

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

// WithDefaults hydrates default values in the extras graphql queries run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasGraphqlQueriesRunParams) WithDefaults() *ExtrasGraphqlQueriesRunParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras graphql queries run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasGraphqlQueriesRunParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras graphql queries run params
func (o *ExtrasGraphqlQueriesRunParams) WithTimeout(timeout time.Duration) *ExtrasGraphqlQueriesRunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras graphql queries run params
func (o *ExtrasGraphqlQueriesRunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras graphql queries run params
func (o *ExtrasGraphqlQueriesRunParams) WithContext(ctx context.Context) *ExtrasGraphqlQueriesRunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras graphql queries run params
func (o *ExtrasGraphqlQueriesRunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras graphql queries run params
func (o *ExtrasGraphqlQueriesRunParams) WithHTTPClient(client *http.Client) *ExtrasGraphqlQueriesRunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras graphql queries run params
func (o *ExtrasGraphqlQueriesRunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras graphql queries run params
func (o *ExtrasGraphqlQueriesRunParams) WithData(data *models.GraphQLQuery) *ExtrasGraphqlQueriesRunParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras graphql queries run params
func (o *ExtrasGraphqlQueriesRunParams) SetData(data *models.GraphQLQuery) {
	o.Data = data
}

// WithID adds the id to the extras graphql queries run params
func (o *ExtrasGraphqlQueriesRunParams) WithID(id strfmt.UUID) *ExtrasGraphqlQueriesRunParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras graphql queries run params
func (o *ExtrasGraphqlQueriesRunParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasGraphqlQueriesRunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

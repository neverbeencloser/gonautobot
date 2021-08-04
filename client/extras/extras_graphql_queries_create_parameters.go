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

// NewExtrasGraphqlQueriesCreateParams creates a new ExtrasGraphqlQueriesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasGraphqlQueriesCreateParams() *ExtrasGraphqlQueriesCreateParams {
	return &ExtrasGraphqlQueriesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasGraphqlQueriesCreateParamsWithTimeout creates a new ExtrasGraphqlQueriesCreateParams object
// with the ability to set a timeout on a request.
func NewExtrasGraphqlQueriesCreateParamsWithTimeout(timeout time.Duration) *ExtrasGraphqlQueriesCreateParams {
	return &ExtrasGraphqlQueriesCreateParams{
		timeout: timeout,
	}
}

// NewExtrasGraphqlQueriesCreateParamsWithContext creates a new ExtrasGraphqlQueriesCreateParams object
// with the ability to set a context for a request.
func NewExtrasGraphqlQueriesCreateParamsWithContext(ctx context.Context) *ExtrasGraphqlQueriesCreateParams {
	return &ExtrasGraphqlQueriesCreateParams{
		Context: ctx,
	}
}

// NewExtrasGraphqlQueriesCreateParamsWithHTTPClient creates a new ExtrasGraphqlQueriesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasGraphqlQueriesCreateParamsWithHTTPClient(client *http.Client) *ExtrasGraphqlQueriesCreateParams {
	return &ExtrasGraphqlQueriesCreateParams{
		HTTPClient: client,
	}
}

/* ExtrasGraphqlQueriesCreateParams contains all the parameters to send to the API endpoint
   for the extras graphql queries create operation.

   Typically these are written to a http.Request.
*/
type ExtrasGraphqlQueriesCreateParams struct {

	// Data.
	Data *models.GraphQLQuery

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras graphql queries create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasGraphqlQueriesCreateParams) WithDefaults() *ExtrasGraphqlQueriesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras graphql queries create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasGraphqlQueriesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras graphql queries create params
func (o *ExtrasGraphqlQueriesCreateParams) WithTimeout(timeout time.Duration) *ExtrasGraphqlQueriesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras graphql queries create params
func (o *ExtrasGraphqlQueriesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras graphql queries create params
func (o *ExtrasGraphqlQueriesCreateParams) WithContext(ctx context.Context) *ExtrasGraphqlQueriesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras graphql queries create params
func (o *ExtrasGraphqlQueriesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras graphql queries create params
func (o *ExtrasGraphqlQueriesCreateParams) WithHTTPClient(client *http.Client) *ExtrasGraphqlQueriesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras graphql queries create params
func (o *ExtrasGraphqlQueriesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras graphql queries create params
func (o *ExtrasGraphqlQueriesCreateParams) WithData(data *models.GraphQLQuery) *ExtrasGraphqlQueriesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras graphql queries create params
func (o *ExtrasGraphqlQueriesCreateParams) SetData(data *models.GraphQLQuery) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasGraphqlQueriesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

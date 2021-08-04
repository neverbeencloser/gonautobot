package graphql

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

// NewGraphqlCreateParams creates a new GraphqlCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGraphqlCreateParams() *GraphqlCreateParams {
	return &GraphqlCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGraphqlCreateParamsWithTimeout creates a new GraphqlCreateParams object
// with the ability to set a timeout on a request.
func NewGraphqlCreateParamsWithTimeout(timeout time.Duration) *GraphqlCreateParams {
	return &GraphqlCreateParams{
		timeout: timeout,
	}
}

// NewGraphqlCreateParamsWithContext creates a new GraphqlCreateParams object
// with the ability to set a context for a request.
func NewGraphqlCreateParamsWithContext(ctx context.Context) *GraphqlCreateParams {
	return &GraphqlCreateParams{
		Context: ctx,
	}
}

// NewGraphqlCreateParamsWithHTTPClient creates a new GraphqlCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewGraphqlCreateParamsWithHTTPClient(client *http.Client) *GraphqlCreateParams {
	return &GraphqlCreateParams{
		HTTPClient: client,
	}
}

/* GraphqlCreateParams contains all the parameters to send to the API endpoint
   for the graphql create operation.

   Typically these are written to a http.Request.
*/
type GraphqlCreateParams struct {

	// Data.
	Data *models.GraphQLAPI

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the graphql create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GraphqlCreateParams) WithDefaults() *GraphqlCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the graphql create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GraphqlCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the graphql create params
func (o *GraphqlCreateParams) WithTimeout(timeout time.Duration) *GraphqlCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the graphql create params
func (o *GraphqlCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the graphql create params
func (o *GraphqlCreateParams) WithContext(ctx context.Context) *GraphqlCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the graphql create params
func (o *GraphqlCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the graphql create params
func (o *GraphqlCreateParams) WithHTTPClient(client *http.Client) *GraphqlCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the graphql create params
func (o *GraphqlCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the graphql create params
func (o *GraphqlCreateParams) WithData(data *models.GraphQLAPI) *GraphqlCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the graphql create params
func (o *GraphqlCreateParams) SetData(data *models.GraphQLAPI) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *GraphqlCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

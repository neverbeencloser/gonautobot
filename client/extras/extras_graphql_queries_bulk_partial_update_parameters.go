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

// NewExtrasGraphqlQueriesBulkPartialUpdateParams creates a new ExtrasGraphqlQueriesBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasGraphqlQueriesBulkPartialUpdateParams() *ExtrasGraphqlQueriesBulkPartialUpdateParams {
	return &ExtrasGraphqlQueriesBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasGraphqlQueriesBulkPartialUpdateParamsWithTimeout creates a new ExtrasGraphqlQueriesBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasGraphqlQueriesBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *ExtrasGraphqlQueriesBulkPartialUpdateParams {
	return &ExtrasGraphqlQueriesBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasGraphqlQueriesBulkPartialUpdateParamsWithContext creates a new ExtrasGraphqlQueriesBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewExtrasGraphqlQueriesBulkPartialUpdateParamsWithContext(ctx context.Context) *ExtrasGraphqlQueriesBulkPartialUpdateParams {
	return &ExtrasGraphqlQueriesBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewExtrasGraphqlQueriesBulkPartialUpdateParamsWithHTTPClient creates a new ExtrasGraphqlQueriesBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasGraphqlQueriesBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *ExtrasGraphqlQueriesBulkPartialUpdateParams {
	return &ExtrasGraphqlQueriesBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasGraphqlQueriesBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the extras graphql queries bulk partial update operation.

   Typically these are written to a http.Request.
*/
type ExtrasGraphqlQueriesBulkPartialUpdateParams struct {

	// Data.
	Data *models.GraphQLQuery

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras graphql queries bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasGraphqlQueriesBulkPartialUpdateParams) WithDefaults() *ExtrasGraphqlQueriesBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras graphql queries bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasGraphqlQueriesBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras graphql queries bulk partial update params
func (o *ExtrasGraphqlQueriesBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *ExtrasGraphqlQueriesBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras graphql queries bulk partial update params
func (o *ExtrasGraphqlQueriesBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras graphql queries bulk partial update params
func (o *ExtrasGraphqlQueriesBulkPartialUpdateParams) WithContext(ctx context.Context) *ExtrasGraphqlQueriesBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras graphql queries bulk partial update params
func (o *ExtrasGraphqlQueriesBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras graphql queries bulk partial update params
func (o *ExtrasGraphqlQueriesBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *ExtrasGraphqlQueriesBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras graphql queries bulk partial update params
func (o *ExtrasGraphqlQueriesBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras graphql queries bulk partial update params
func (o *ExtrasGraphqlQueriesBulkPartialUpdateParams) WithData(data *models.GraphQLQuery) *ExtrasGraphqlQueriesBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras graphql queries bulk partial update params
func (o *ExtrasGraphqlQueriesBulkPartialUpdateParams) SetData(data *models.GraphQLQuery) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasGraphqlQueriesBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

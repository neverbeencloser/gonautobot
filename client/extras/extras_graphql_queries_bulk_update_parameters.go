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

// NewExtrasGraphqlQueriesBulkUpdateParams creates a new ExtrasGraphqlQueriesBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasGraphqlQueriesBulkUpdateParams() *ExtrasGraphqlQueriesBulkUpdateParams {
	return &ExtrasGraphqlQueriesBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasGraphqlQueriesBulkUpdateParamsWithTimeout creates a new ExtrasGraphqlQueriesBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasGraphqlQueriesBulkUpdateParamsWithTimeout(timeout time.Duration) *ExtrasGraphqlQueriesBulkUpdateParams {
	return &ExtrasGraphqlQueriesBulkUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasGraphqlQueriesBulkUpdateParamsWithContext creates a new ExtrasGraphqlQueriesBulkUpdateParams object
// with the ability to set a context for a request.
func NewExtrasGraphqlQueriesBulkUpdateParamsWithContext(ctx context.Context) *ExtrasGraphqlQueriesBulkUpdateParams {
	return &ExtrasGraphqlQueriesBulkUpdateParams{
		Context: ctx,
	}
}

// NewExtrasGraphqlQueriesBulkUpdateParamsWithHTTPClient creates a new ExtrasGraphqlQueriesBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasGraphqlQueriesBulkUpdateParamsWithHTTPClient(client *http.Client) *ExtrasGraphqlQueriesBulkUpdateParams {
	return &ExtrasGraphqlQueriesBulkUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasGraphqlQueriesBulkUpdateParams contains all the parameters to send to the API endpoint
   for the extras graphql queries bulk update operation.

   Typically these are written to a http.Request.
*/
type ExtrasGraphqlQueriesBulkUpdateParams struct {

	// Data.
	Data *models.GraphQLQuery

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras graphql queries bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasGraphqlQueriesBulkUpdateParams) WithDefaults() *ExtrasGraphqlQueriesBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras graphql queries bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasGraphqlQueriesBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras graphql queries bulk update params
func (o *ExtrasGraphqlQueriesBulkUpdateParams) WithTimeout(timeout time.Duration) *ExtrasGraphqlQueriesBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras graphql queries bulk update params
func (o *ExtrasGraphqlQueriesBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras graphql queries bulk update params
func (o *ExtrasGraphqlQueriesBulkUpdateParams) WithContext(ctx context.Context) *ExtrasGraphqlQueriesBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras graphql queries bulk update params
func (o *ExtrasGraphqlQueriesBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras graphql queries bulk update params
func (o *ExtrasGraphqlQueriesBulkUpdateParams) WithHTTPClient(client *http.Client) *ExtrasGraphqlQueriesBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras graphql queries bulk update params
func (o *ExtrasGraphqlQueriesBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras graphql queries bulk update params
func (o *ExtrasGraphqlQueriesBulkUpdateParams) WithData(data *models.GraphQLQuery) *ExtrasGraphqlQueriesBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras graphql queries bulk update params
func (o *ExtrasGraphqlQueriesBulkUpdateParams) SetData(data *models.GraphQLQuery) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasGraphqlQueriesBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

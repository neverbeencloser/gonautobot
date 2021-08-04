package ipam

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

// NewIpamAggregatesPartialUpdateParams creates a new IpamAggregatesPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamAggregatesPartialUpdateParams() *IpamAggregatesPartialUpdateParams {
	return &IpamAggregatesPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamAggregatesPartialUpdateParamsWithTimeout creates a new IpamAggregatesPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamAggregatesPartialUpdateParamsWithTimeout(timeout time.Duration) *IpamAggregatesPartialUpdateParams {
	return &IpamAggregatesPartialUpdateParams{
		timeout: timeout,
	}
}

// NewIpamAggregatesPartialUpdateParamsWithContext creates a new IpamAggregatesPartialUpdateParams object
// with the ability to set a context for a request.
func NewIpamAggregatesPartialUpdateParamsWithContext(ctx context.Context) *IpamAggregatesPartialUpdateParams {
	return &IpamAggregatesPartialUpdateParams{
		Context: ctx,
	}
}

// NewIpamAggregatesPartialUpdateParamsWithHTTPClient creates a new IpamAggregatesPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamAggregatesPartialUpdateParamsWithHTTPClient(client *http.Client) *IpamAggregatesPartialUpdateParams {
	return &IpamAggregatesPartialUpdateParams{
		HTTPClient: client,
	}
}

/* IpamAggregatesPartialUpdateParams contains all the parameters to send to the API endpoint
   for the ipam aggregates partial update operation.

   Typically these are written to a http.Request.
*/
type IpamAggregatesPartialUpdateParams struct {

	// Data.
	Data *models.WritableAggregate

	/* ID.

	   A UUID string identifying this aggregate.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam aggregates partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamAggregatesPartialUpdateParams) WithDefaults() *IpamAggregatesPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam aggregates partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamAggregatesPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam aggregates partial update params
func (o *IpamAggregatesPartialUpdateParams) WithTimeout(timeout time.Duration) *IpamAggregatesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam aggregates partial update params
func (o *IpamAggregatesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam aggregates partial update params
func (o *IpamAggregatesPartialUpdateParams) WithContext(ctx context.Context) *IpamAggregatesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam aggregates partial update params
func (o *IpamAggregatesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam aggregates partial update params
func (o *IpamAggregatesPartialUpdateParams) WithHTTPClient(client *http.Client) *IpamAggregatesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam aggregates partial update params
func (o *IpamAggregatesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam aggregates partial update params
func (o *IpamAggregatesPartialUpdateParams) WithData(data *models.WritableAggregate) *IpamAggregatesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam aggregates partial update params
func (o *IpamAggregatesPartialUpdateParams) SetData(data *models.WritableAggregate) {
	o.Data = data
}

// WithID adds the id to the ipam aggregates partial update params
func (o *IpamAggregatesPartialUpdateParams) WithID(id strfmt.UUID) *IpamAggregatesPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam aggregates partial update params
func (o *IpamAggregatesPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamAggregatesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

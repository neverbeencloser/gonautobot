package ipam

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewIpamAggregatesDeleteParams creates a new IpamAggregatesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamAggregatesDeleteParams() *IpamAggregatesDeleteParams {
	return &IpamAggregatesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamAggregatesDeleteParamsWithTimeout creates a new IpamAggregatesDeleteParams object
// with the ability to set a timeout on a request.
func NewIpamAggregatesDeleteParamsWithTimeout(timeout time.Duration) *IpamAggregatesDeleteParams {
	return &IpamAggregatesDeleteParams{
		timeout: timeout,
	}
}

// NewIpamAggregatesDeleteParamsWithContext creates a new IpamAggregatesDeleteParams object
// with the ability to set a context for a request.
func NewIpamAggregatesDeleteParamsWithContext(ctx context.Context) *IpamAggregatesDeleteParams {
	return &IpamAggregatesDeleteParams{
		Context: ctx,
	}
}

// NewIpamAggregatesDeleteParamsWithHTTPClient creates a new IpamAggregatesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamAggregatesDeleteParamsWithHTTPClient(client *http.Client) *IpamAggregatesDeleteParams {
	return &IpamAggregatesDeleteParams{
		HTTPClient: client,
	}
}

/* IpamAggregatesDeleteParams contains all the parameters to send to the API endpoint
   for the ipam aggregates delete operation.

   Typically these are written to a http.Request.
*/
type IpamAggregatesDeleteParams struct {

	/* ID.

	   A UUID string identifying this aggregate.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam aggregates delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamAggregatesDeleteParams) WithDefaults() *IpamAggregatesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam aggregates delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamAggregatesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam aggregates delete params
func (o *IpamAggregatesDeleteParams) WithTimeout(timeout time.Duration) *IpamAggregatesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam aggregates delete params
func (o *IpamAggregatesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam aggregates delete params
func (o *IpamAggregatesDeleteParams) WithContext(ctx context.Context) *IpamAggregatesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam aggregates delete params
func (o *IpamAggregatesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam aggregates delete params
func (o *IpamAggregatesDeleteParams) WithHTTPClient(client *http.Client) *IpamAggregatesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam aggregates delete params
func (o *IpamAggregatesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ipam aggregates delete params
func (o *IpamAggregatesDeleteParams) WithID(id strfmt.UUID) *IpamAggregatesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam aggregates delete params
func (o *IpamAggregatesDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamAggregatesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

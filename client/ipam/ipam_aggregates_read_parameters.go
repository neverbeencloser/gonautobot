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

// NewIpamAggregatesReadParams creates a new IpamAggregatesReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamAggregatesReadParams() *IpamAggregatesReadParams {
	return &IpamAggregatesReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamAggregatesReadParamsWithTimeout creates a new IpamAggregatesReadParams object
// with the ability to set a timeout on a request.
func NewIpamAggregatesReadParamsWithTimeout(timeout time.Duration) *IpamAggregatesReadParams {
	return &IpamAggregatesReadParams{
		timeout: timeout,
	}
}

// NewIpamAggregatesReadParamsWithContext creates a new IpamAggregatesReadParams object
// with the ability to set a context for a request.
func NewIpamAggregatesReadParamsWithContext(ctx context.Context) *IpamAggregatesReadParams {
	return &IpamAggregatesReadParams{
		Context: ctx,
	}
}

// NewIpamAggregatesReadParamsWithHTTPClient creates a new IpamAggregatesReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamAggregatesReadParamsWithHTTPClient(client *http.Client) *IpamAggregatesReadParams {
	return &IpamAggregatesReadParams{
		HTTPClient: client,
	}
}

/* IpamAggregatesReadParams contains all the parameters to send to the API endpoint
   for the ipam aggregates read operation.

   Typically these are written to a http.Request.
*/
type IpamAggregatesReadParams struct {

	/* ID.

	   A UUID string identifying this aggregate.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam aggregates read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamAggregatesReadParams) WithDefaults() *IpamAggregatesReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam aggregates read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamAggregatesReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam aggregates read params
func (o *IpamAggregatesReadParams) WithTimeout(timeout time.Duration) *IpamAggregatesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam aggregates read params
func (o *IpamAggregatesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam aggregates read params
func (o *IpamAggregatesReadParams) WithContext(ctx context.Context) *IpamAggregatesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam aggregates read params
func (o *IpamAggregatesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam aggregates read params
func (o *IpamAggregatesReadParams) WithHTTPClient(client *http.Client) *IpamAggregatesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam aggregates read params
func (o *IpamAggregatesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ipam aggregates read params
func (o *IpamAggregatesReadParams) WithID(id strfmt.UUID) *IpamAggregatesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam aggregates read params
func (o *IpamAggregatesReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamAggregatesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

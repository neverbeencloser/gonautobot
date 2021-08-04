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

// NewIpamVlansDeleteParams creates a new IpamVlansDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamVlansDeleteParams() *IpamVlansDeleteParams {
	return &IpamVlansDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamVlansDeleteParamsWithTimeout creates a new IpamVlansDeleteParams object
// with the ability to set a timeout on a request.
func NewIpamVlansDeleteParamsWithTimeout(timeout time.Duration) *IpamVlansDeleteParams {
	return &IpamVlansDeleteParams{
		timeout: timeout,
	}
}

// NewIpamVlansDeleteParamsWithContext creates a new IpamVlansDeleteParams object
// with the ability to set a context for a request.
func NewIpamVlansDeleteParamsWithContext(ctx context.Context) *IpamVlansDeleteParams {
	return &IpamVlansDeleteParams{
		Context: ctx,
	}
}

// NewIpamVlansDeleteParamsWithHTTPClient creates a new IpamVlansDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamVlansDeleteParamsWithHTTPClient(client *http.Client) *IpamVlansDeleteParams {
	return &IpamVlansDeleteParams{
		HTTPClient: client,
	}
}

/* IpamVlansDeleteParams contains all the parameters to send to the API endpoint
   for the ipam vlans delete operation.

   Typically these are written to a http.Request.
*/
type IpamVlansDeleteParams struct {

	/* ID.

	   A UUID string identifying this VLAN.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam vlans delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVlansDeleteParams) WithDefaults() *IpamVlansDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam vlans delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVlansDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam vlans delete params
func (o *IpamVlansDeleteParams) WithTimeout(timeout time.Duration) *IpamVlansDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam vlans delete params
func (o *IpamVlansDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam vlans delete params
func (o *IpamVlansDeleteParams) WithContext(ctx context.Context) *IpamVlansDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam vlans delete params
func (o *IpamVlansDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam vlans delete params
func (o *IpamVlansDeleteParams) WithHTTPClient(client *http.Client) *IpamVlansDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam vlans delete params
func (o *IpamVlansDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ipam vlans delete params
func (o *IpamVlansDeleteParams) WithID(id strfmt.UUID) *IpamVlansDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam vlans delete params
func (o *IpamVlansDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamVlansDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

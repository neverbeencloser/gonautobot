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

// NewIpamIPAddressesDeleteParams creates a new IpamIPAddressesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamIPAddressesDeleteParams() *IpamIPAddressesDeleteParams {
	return &IpamIPAddressesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamIPAddressesDeleteParamsWithTimeout creates a new IpamIPAddressesDeleteParams object
// with the ability to set a timeout on a request.
func NewIpamIPAddressesDeleteParamsWithTimeout(timeout time.Duration) *IpamIPAddressesDeleteParams {
	return &IpamIPAddressesDeleteParams{
		timeout: timeout,
	}
}

// NewIpamIPAddressesDeleteParamsWithContext creates a new IpamIPAddressesDeleteParams object
// with the ability to set a context for a request.
func NewIpamIPAddressesDeleteParamsWithContext(ctx context.Context) *IpamIPAddressesDeleteParams {
	return &IpamIPAddressesDeleteParams{
		Context: ctx,
	}
}

// NewIpamIPAddressesDeleteParamsWithHTTPClient creates a new IpamIPAddressesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamIPAddressesDeleteParamsWithHTTPClient(client *http.Client) *IpamIPAddressesDeleteParams {
	return &IpamIPAddressesDeleteParams{
		HTTPClient: client,
	}
}

/* IpamIPAddressesDeleteParams contains all the parameters to send to the API endpoint
   for the ipam ip addresses delete operation.

   Typically these are written to a http.Request.
*/
type IpamIPAddressesDeleteParams struct {

	/* ID.

	   A UUID string identifying this IP address.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam ip addresses delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamIPAddressesDeleteParams) WithDefaults() *IpamIPAddressesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam ip addresses delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamIPAddressesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam ip addresses delete params
func (o *IpamIPAddressesDeleteParams) WithTimeout(timeout time.Duration) *IpamIPAddressesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam ip addresses delete params
func (o *IpamIPAddressesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam ip addresses delete params
func (o *IpamIPAddressesDeleteParams) WithContext(ctx context.Context) *IpamIPAddressesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam ip addresses delete params
func (o *IpamIPAddressesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam ip addresses delete params
func (o *IpamIPAddressesDeleteParams) WithHTTPClient(client *http.Client) *IpamIPAddressesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam ip addresses delete params
func (o *IpamIPAddressesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ipam ip addresses delete params
func (o *IpamIPAddressesDeleteParams) WithID(id strfmt.UUID) *IpamIPAddressesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam ip addresses delete params
func (o *IpamIPAddressesDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamIPAddressesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

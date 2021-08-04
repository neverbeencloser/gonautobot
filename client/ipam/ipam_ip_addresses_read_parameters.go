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

// NewIpamIPAddressesReadParams creates a new IpamIPAddressesReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamIPAddressesReadParams() *IpamIPAddressesReadParams {
	return &IpamIPAddressesReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamIPAddressesReadParamsWithTimeout creates a new IpamIPAddressesReadParams object
// with the ability to set a timeout on a request.
func NewIpamIPAddressesReadParamsWithTimeout(timeout time.Duration) *IpamIPAddressesReadParams {
	return &IpamIPAddressesReadParams{
		timeout: timeout,
	}
}

// NewIpamIPAddressesReadParamsWithContext creates a new IpamIPAddressesReadParams object
// with the ability to set a context for a request.
func NewIpamIPAddressesReadParamsWithContext(ctx context.Context) *IpamIPAddressesReadParams {
	return &IpamIPAddressesReadParams{
		Context: ctx,
	}
}

// NewIpamIPAddressesReadParamsWithHTTPClient creates a new IpamIPAddressesReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamIPAddressesReadParamsWithHTTPClient(client *http.Client) *IpamIPAddressesReadParams {
	return &IpamIPAddressesReadParams{
		HTTPClient: client,
	}
}

/* IpamIPAddressesReadParams contains all the parameters to send to the API endpoint
   for the ipam ip addresses read operation.

   Typically these are written to a http.Request.
*/
type IpamIPAddressesReadParams struct {

	/* ID.

	   A UUID string identifying this IP address.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam ip addresses read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamIPAddressesReadParams) WithDefaults() *IpamIPAddressesReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam ip addresses read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamIPAddressesReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam ip addresses read params
func (o *IpamIPAddressesReadParams) WithTimeout(timeout time.Duration) *IpamIPAddressesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam ip addresses read params
func (o *IpamIPAddressesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam ip addresses read params
func (o *IpamIPAddressesReadParams) WithContext(ctx context.Context) *IpamIPAddressesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam ip addresses read params
func (o *IpamIPAddressesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam ip addresses read params
func (o *IpamIPAddressesReadParams) WithHTTPClient(client *http.Client) *IpamIPAddressesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam ip addresses read params
func (o *IpamIPAddressesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ipam ip addresses read params
func (o *IpamIPAddressesReadParams) WithID(id strfmt.UUID) *IpamIPAddressesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam ip addresses read params
func (o *IpamIPAddressesReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamIPAddressesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

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

// NewIpamPrefixesReadParams creates a new IpamPrefixesReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamPrefixesReadParams() *IpamPrefixesReadParams {
	return &IpamPrefixesReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamPrefixesReadParamsWithTimeout creates a new IpamPrefixesReadParams object
// with the ability to set a timeout on a request.
func NewIpamPrefixesReadParamsWithTimeout(timeout time.Duration) *IpamPrefixesReadParams {
	return &IpamPrefixesReadParams{
		timeout: timeout,
	}
}

// NewIpamPrefixesReadParamsWithContext creates a new IpamPrefixesReadParams object
// with the ability to set a context for a request.
func NewIpamPrefixesReadParamsWithContext(ctx context.Context) *IpamPrefixesReadParams {
	return &IpamPrefixesReadParams{
		Context: ctx,
	}
}

// NewIpamPrefixesReadParamsWithHTTPClient creates a new IpamPrefixesReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamPrefixesReadParamsWithHTTPClient(client *http.Client) *IpamPrefixesReadParams {
	return &IpamPrefixesReadParams{
		HTTPClient: client,
	}
}

/* IpamPrefixesReadParams contains all the parameters to send to the API endpoint
   for the ipam prefixes read operation.

   Typically these are written to a http.Request.
*/
type IpamPrefixesReadParams struct {

	/* ID.

	   A UUID string identifying this prefix.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam prefixes read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamPrefixesReadParams) WithDefaults() *IpamPrefixesReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam prefixes read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamPrefixesReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam prefixes read params
func (o *IpamPrefixesReadParams) WithTimeout(timeout time.Duration) *IpamPrefixesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam prefixes read params
func (o *IpamPrefixesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam prefixes read params
func (o *IpamPrefixesReadParams) WithContext(ctx context.Context) *IpamPrefixesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam prefixes read params
func (o *IpamPrefixesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam prefixes read params
func (o *IpamPrefixesReadParams) WithHTTPClient(client *http.Client) *IpamPrefixesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam prefixes read params
func (o *IpamPrefixesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ipam prefixes read params
func (o *IpamPrefixesReadParams) WithID(id strfmt.UUID) *IpamPrefixesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam prefixes read params
func (o *IpamPrefixesReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamPrefixesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

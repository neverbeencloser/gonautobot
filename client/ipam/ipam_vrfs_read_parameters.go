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

// NewIpamVrfsReadParams creates a new IpamVrfsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamVrfsReadParams() *IpamVrfsReadParams {
	return &IpamVrfsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamVrfsReadParamsWithTimeout creates a new IpamVrfsReadParams object
// with the ability to set a timeout on a request.
func NewIpamVrfsReadParamsWithTimeout(timeout time.Duration) *IpamVrfsReadParams {
	return &IpamVrfsReadParams{
		timeout: timeout,
	}
}

// NewIpamVrfsReadParamsWithContext creates a new IpamVrfsReadParams object
// with the ability to set a context for a request.
func NewIpamVrfsReadParamsWithContext(ctx context.Context) *IpamVrfsReadParams {
	return &IpamVrfsReadParams{
		Context: ctx,
	}
}

// NewIpamVrfsReadParamsWithHTTPClient creates a new IpamVrfsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamVrfsReadParamsWithHTTPClient(client *http.Client) *IpamVrfsReadParams {
	return &IpamVrfsReadParams{
		HTTPClient: client,
	}
}

/* IpamVrfsReadParams contains all the parameters to send to the API endpoint
   for the ipam vrfs read operation.

   Typically these are written to a http.Request.
*/
type IpamVrfsReadParams struct {

	/* ID.

	   A UUID string identifying this VRF.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam vrfs read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVrfsReadParams) WithDefaults() *IpamVrfsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam vrfs read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVrfsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam vrfs read params
func (o *IpamVrfsReadParams) WithTimeout(timeout time.Duration) *IpamVrfsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam vrfs read params
func (o *IpamVrfsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam vrfs read params
func (o *IpamVrfsReadParams) WithContext(ctx context.Context) *IpamVrfsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam vrfs read params
func (o *IpamVrfsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam vrfs read params
func (o *IpamVrfsReadParams) WithHTTPClient(client *http.Client) *IpamVrfsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam vrfs read params
func (o *IpamVrfsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ipam vrfs read params
func (o *IpamVrfsReadParams) WithID(id strfmt.UUID) *IpamVrfsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam vrfs read params
func (o *IpamVrfsReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamVrfsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

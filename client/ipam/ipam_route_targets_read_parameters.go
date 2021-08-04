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

// NewIpamRouteTargetsReadParams creates a new IpamRouteTargetsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamRouteTargetsReadParams() *IpamRouteTargetsReadParams {
	return &IpamRouteTargetsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamRouteTargetsReadParamsWithTimeout creates a new IpamRouteTargetsReadParams object
// with the ability to set a timeout on a request.
func NewIpamRouteTargetsReadParamsWithTimeout(timeout time.Duration) *IpamRouteTargetsReadParams {
	return &IpamRouteTargetsReadParams{
		timeout: timeout,
	}
}

// NewIpamRouteTargetsReadParamsWithContext creates a new IpamRouteTargetsReadParams object
// with the ability to set a context for a request.
func NewIpamRouteTargetsReadParamsWithContext(ctx context.Context) *IpamRouteTargetsReadParams {
	return &IpamRouteTargetsReadParams{
		Context: ctx,
	}
}

// NewIpamRouteTargetsReadParamsWithHTTPClient creates a new IpamRouteTargetsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamRouteTargetsReadParamsWithHTTPClient(client *http.Client) *IpamRouteTargetsReadParams {
	return &IpamRouteTargetsReadParams{
		HTTPClient: client,
	}
}

/* IpamRouteTargetsReadParams contains all the parameters to send to the API endpoint
   for the ipam route targets read operation.

   Typically these are written to a http.Request.
*/
type IpamRouteTargetsReadParams struct {

	/* ID.

	   A UUID string identifying this route target.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam route targets read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamRouteTargetsReadParams) WithDefaults() *IpamRouteTargetsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam route targets read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamRouteTargetsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam route targets read params
func (o *IpamRouteTargetsReadParams) WithTimeout(timeout time.Duration) *IpamRouteTargetsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam route targets read params
func (o *IpamRouteTargetsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam route targets read params
func (o *IpamRouteTargetsReadParams) WithContext(ctx context.Context) *IpamRouteTargetsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam route targets read params
func (o *IpamRouteTargetsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam route targets read params
func (o *IpamRouteTargetsReadParams) WithHTTPClient(client *http.Client) *IpamRouteTargetsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam route targets read params
func (o *IpamRouteTargetsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ipam route targets read params
func (o *IpamRouteTargetsReadParams) WithID(id strfmt.UUID) *IpamRouteTargetsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam route targets read params
func (o *IpamRouteTargetsReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamRouteTargetsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

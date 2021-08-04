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

// NewIpamIPAddressesUpdateParams creates a new IpamIPAddressesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamIPAddressesUpdateParams() *IpamIPAddressesUpdateParams {
	return &IpamIPAddressesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamIPAddressesUpdateParamsWithTimeout creates a new IpamIPAddressesUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamIPAddressesUpdateParamsWithTimeout(timeout time.Duration) *IpamIPAddressesUpdateParams {
	return &IpamIPAddressesUpdateParams{
		timeout: timeout,
	}
}

// NewIpamIPAddressesUpdateParamsWithContext creates a new IpamIPAddressesUpdateParams object
// with the ability to set a context for a request.
func NewIpamIPAddressesUpdateParamsWithContext(ctx context.Context) *IpamIPAddressesUpdateParams {
	return &IpamIPAddressesUpdateParams{
		Context: ctx,
	}
}

// NewIpamIPAddressesUpdateParamsWithHTTPClient creates a new IpamIPAddressesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamIPAddressesUpdateParamsWithHTTPClient(client *http.Client) *IpamIPAddressesUpdateParams {
	return &IpamIPAddressesUpdateParams{
		HTTPClient: client,
	}
}

/* IpamIPAddressesUpdateParams contains all the parameters to send to the API endpoint
   for the ipam ip addresses update operation.

   Typically these are written to a http.Request.
*/
type IpamIPAddressesUpdateParams struct {

	// Data.
	Data *models.WritableIPAddress

	/* ID.

	   A UUID string identifying this IP address.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam ip addresses update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamIPAddressesUpdateParams) WithDefaults() *IpamIPAddressesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam ip addresses update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamIPAddressesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam ip addresses update params
func (o *IpamIPAddressesUpdateParams) WithTimeout(timeout time.Duration) *IpamIPAddressesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam ip addresses update params
func (o *IpamIPAddressesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam ip addresses update params
func (o *IpamIPAddressesUpdateParams) WithContext(ctx context.Context) *IpamIPAddressesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam ip addresses update params
func (o *IpamIPAddressesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam ip addresses update params
func (o *IpamIPAddressesUpdateParams) WithHTTPClient(client *http.Client) *IpamIPAddressesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam ip addresses update params
func (o *IpamIPAddressesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam ip addresses update params
func (o *IpamIPAddressesUpdateParams) WithData(data *models.WritableIPAddress) *IpamIPAddressesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam ip addresses update params
func (o *IpamIPAddressesUpdateParams) SetData(data *models.WritableIPAddress) {
	o.Data = data
}

// WithID adds the id to the ipam ip addresses update params
func (o *IpamIPAddressesUpdateParams) WithID(id strfmt.UUID) *IpamIPAddressesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam ip addresses update params
func (o *IpamIPAddressesUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamIPAddressesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

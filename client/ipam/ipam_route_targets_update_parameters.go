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

// NewIpamRouteTargetsUpdateParams creates a new IpamRouteTargetsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamRouteTargetsUpdateParams() *IpamRouteTargetsUpdateParams {
	return &IpamRouteTargetsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamRouteTargetsUpdateParamsWithTimeout creates a new IpamRouteTargetsUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamRouteTargetsUpdateParamsWithTimeout(timeout time.Duration) *IpamRouteTargetsUpdateParams {
	return &IpamRouteTargetsUpdateParams{
		timeout: timeout,
	}
}

// NewIpamRouteTargetsUpdateParamsWithContext creates a new IpamRouteTargetsUpdateParams object
// with the ability to set a context for a request.
func NewIpamRouteTargetsUpdateParamsWithContext(ctx context.Context) *IpamRouteTargetsUpdateParams {
	return &IpamRouteTargetsUpdateParams{
		Context: ctx,
	}
}

// NewIpamRouteTargetsUpdateParamsWithHTTPClient creates a new IpamRouteTargetsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamRouteTargetsUpdateParamsWithHTTPClient(client *http.Client) *IpamRouteTargetsUpdateParams {
	return &IpamRouteTargetsUpdateParams{
		HTTPClient: client,
	}
}

/* IpamRouteTargetsUpdateParams contains all the parameters to send to the API endpoint
   for the ipam route targets update operation.

   Typically these are written to a http.Request.
*/
type IpamRouteTargetsUpdateParams struct {

	// Data.
	Data *models.WritableRouteTarget

	/* ID.

	   A UUID string identifying this route target.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam route targets update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamRouteTargetsUpdateParams) WithDefaults() *IpamRouteTargetsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam route targets update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamRouteTargetsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam route targets update params
func (o *IpamRouteTargetsUpdateParams) WithTimeout(timeout time.Duration) *IpamRouteTargetsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam route targets update params
func (o *IpamRouteTargetsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam route targets update params
func (o *IpamRouteTargetsUpdateParams) WithContext(ctx context.Context) *IpamRouteTargetsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam route targets update params
func (o *IpamRouteTargetsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam route targets update params
func (o *IpamRouteTargetsUpdateParams) WithHTTPClient(client *http.Client) *IpamRouteTargetsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam route targets update params
func (o *IpamRouteTargetsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam route targets update params
func (o *IpamRouteTargetsUpdateParams) WithData(data *models.WritableRouteTarget) *IpamRouteTargetsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam route targets update params
func (o *IpamRouteTargetsUpdateParams) SetData(data *models.WritableRouteTarget) {
	o.Data = data
}

// WithID adds the id to the ipam route targets update params
func (o *IpamRouteTargetsUpdateParams) WithID(id strfmt.UUID) *IpamRouteTargetsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam route targets update params
func (o *IpamRouteTargetsUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamRouteTargetsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

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

// NewIpamRolesPartialUpdateParams creates a new IpamRolesPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamRolesPartialUpdateParams() *IpamRolesPartialUpdateParams {
	return &IpamRolesPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamRolesPartialUpdateParamsWithTimeout creates a new IpamRolesPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamRolesPartialUpdateParamsWithTimeout(timeout time.Duration) *IpamRolesPartialUpdateParams {
	return &IpamRolesPartialUpdateParams{
		timeout: timeout,
	}
}

// NewIpamRolesPartialUpdateParamsWithContext creates a new IpamRolesPartialUpdateParams object
// with the ability to set a context for a request.
func NewIpamRolesPartialUpdateParamsWithContext(ctx context.Context) *IpamRolesPartialUpdateParams {
	return &IpamRolesPartialUpdateParams{
		Context: ctx,
	}
}

// NewIpamRolesPartialUpdateParamsWithHTTPClient creates a new IpamRolesPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamRolesPartialUpdateParamsWithHTTPClient(client *http.Client) *IpamRolesPartialUpdateParams {
	return &IpamRolesPartialUpdateParams{
		HTTPClient: client,
	}
}

/* IpamRolesPartialUpdateParams contains all the parameters to send to the API endpoint
   for the ipam roles partial update operation.

   Typically these are written to a http.Request.
*/
type IpamRolesPartialUpdateParams struct {

	// Data.
	Data *models.Role

	/* ID.

	   A UUID string identifying this role.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam roles partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamRolesPartialUpdateParams) WithDefaults() *IpamRolesPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam roles partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamRolesPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam roles partial update params
func (o *IpamRolesPartialUpdateParams) WithTimeout(timeout time.Duration) *IpamRolesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam roles partial update params
func (o *IpamRolesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam roles partial update params
func (o *IpamRolesPartialUpdateParams) WithContext(ctx context.Context) *IpamRolesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam roles partial update params
func (o *IpamRolesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam roles partial update params
func (o *IpamRolesPartialUpdateParams) WithHTTPClient(client *http.Client) *IpamRolesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam roles partial update params
func (o *IpamRolesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam roles partial update params
func (o *IpamRolesPartialUpdateParams) WithData(data *models.Role) *IpamRolesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam roles partial update params
func (o *IpamRolesPartialUpdateParams) SetData(data *models.Role) {
	o.Data = data
}

// WithID adds the id to the ipam roles partial update params
func (o *IpamRolesPartialUpdateParams) WithID(id strfmt.UUID) *IpamRolesPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam roles partial update params
func (o *IpamRolesPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamRolesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

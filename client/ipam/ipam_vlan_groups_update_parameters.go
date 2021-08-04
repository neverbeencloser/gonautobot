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

// NewIpamVlanGroupsUpdateParams creates a new IpamVlanGroupsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamVlanGroupsUpdateParams() *IpamVlanGroupsUpdateParams {
	return &IpamVlanGroupsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamVlanGroupsUpdateParamsWithTimeout creates a new IpamVlanGroupsUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamVlanGroupsUpdateParamsWithTimeout(timeout time.Duration) *IpamVlanGroupsUpdateParams {
	return &IpamVlanGroupsUpdateParams{
		timeout: timeout,
	}
}

// NewIpamVlanGroupsUpdateParamsWithContext creates a new IpamVlanGroupsUpdateParams object
// with the ability to set a context for a request.
func NewIpamVlanGroupsUpdateParamsWithContext(ctx context.Context) *IpamVlanGroupsUpdateParams {
	return &IpamVlanGroupsUpdateParams{
		Context: ctx,
	}
}

// NewIpamVlanGroupsUpdateParamsWithHTTPClient creates a new IpamVlanGroupsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamVlanGroupsUpdateParamsWithHTTPClient(client *http.Client) *IpamVlanGroupsUpdateParams {
	return &IpamVlanGroupsUpdateParams{
		HTTPClient: client,
	}
}

/* IpamVlanGroupsUpdateParams contains all the parameters to send to the API endpoint
   for the ipam vlan groups update operation.

   Typically these are written to a http.Request.
*/
type IpamVlanGroupsUpdateParams struct {

	// Data.
	Data *models.WritableVLANGroup

	/* ID.

	   A UUID string identifying this VLAN group.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam vlan groups update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVlanGroupsUpdateParams) WithDefaults() *IpamVlanGroupsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam vlan groups update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVlanGroupsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam vlan groups update params
func (o *IpamVlanGroupsUpdateParams) WithTimeout(timeout time.Duration) *IpamVlanGroupsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam vlan groups update params
func (o *IpamVlanGroupsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam vlan groups update params
func (o *IpamVlanGroupsUpdateParams) WithContext(ctx context.Context) *IpamVlanGroupsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam vlan groups update params
func (o *IpamVlanGroupsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam vlan groups update params
func (o *IpamVlanGroupsUpdateParams) WithHTTPClient(client *http.Client) *IpamVlanGroupsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam vlan groups update params
func (o *IpamVlanGroupsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam vlan groups update params
func (o *IpamVlanGroupsUpdateParams) WithData(data *models.WritableVLANGroup) *IpamVlanGroupsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam vlan groups update params
func (o *IpamVlanGroupsUpdateParams) SetData(data *models.WritableVLANGroup) {
	o.Data = data
}

// WithID adds the id to the ipam vlan groups update params
func (o *IpamVlanGroupsUpdateParams) WithID(id strfmt.UUID) *IpamVlanGroupsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam vlan groups update params
func (o *IpamVlanGroupsUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamVlanGroupsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

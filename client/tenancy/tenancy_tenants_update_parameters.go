package tenancy

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

// NewTenancyTenantsUpdateParams creates a new TenancyTenantsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTenancyTenantsUpdateParams() *TenancyTenantsUpdateParams {
	return &TenancyTenantsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTenancyTenantsUpdateParamsWithTimeout creates a new TenancyTenantsUpdateParams object
// with the ability to set a timeout on a request.
func NewTenancyTenantsUpdateParamsWithTimeout(timeout time.Duration) *TenancyTenantsUpdateParams {
	return &TenancyTenantsUpdateParams{
		timeout: timeout,
	}
}

// NewTenancyTenantsUpdateParamsWithContext creates a new TenancyTenantsUpdateParams object
// with the ability to set a context for a request.
func NewTenancyTenantsUpdateParamsWithContext(ctx context.Context) *TenancyTenantsUpdateParams {
	return &TenancyTenantsUpdateParams{
		Context: ctx,
	}
}

// NewTenancyTenantsUpdateParamsWithHTTPClient creates a new TenancyTenantsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewTenancyTenantsUpdateParamsWithHTTPClient(client *http.Client) *TenancyTenantsUpdateParams {
	return &TenancyTenantsUpdateParams{
		HTTPClient: client,
	}
}

/* TenancyTenantsUpdateParams contains all the parameters to send to the API endpoint
   for the tenancy tenants update operation.

   Typically these are written to a http.Request.
*/
type TenancyTenantsUpdateParams struct {

	// Data.
	Data *models.WritableTenant

	/* ID.

	   A UUID string identifying this tenant.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tenancy tenants update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyTenantsUpdateParams) WithDefaults() *TenancyTenantsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tenancy tenants update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyTenantsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tenancy tenants update params
func (o *TenancyTenantsUpdateParams) WithTimeout(timeout time.Duration) *TenancyTenantsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tenancy tenants update params
func (o *TenancyTenantsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tenancy tenants update params
func (o *TenancyTenantsUpdateParams) WithContext(ctx context.Context) *TenancyTenantsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tenancy tenants update params
func (o *TenancyTenantsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tenancy tenants update params
func (o *TenancyTenantsUpdateParams) WithHTTPClient(client *http.Client) *TenancyTenantsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tenancy tenants update params
func (o *TenancyTenantsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the tenancy tenants update params
func (o *TenancyTenantsUpdateParams) WithData(data *models.WritableTenant) *TenancyTenantsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the tenancy tenants update params
func (o *TenancyTenantsUpdateParams) SetData(data *models.WritableTenant) {
	o.Data = data
}

// WithID adds the id to the tenancy tenants update params
func (o *TenancyTenantsUpdateParams) WithID(id strfmt.UUID) *TenancyTenantsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the tenancy tenants update params
func (o *TenancyTenantsUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *TenancyTenantsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

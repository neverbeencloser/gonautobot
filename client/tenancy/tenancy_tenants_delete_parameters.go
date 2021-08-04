package tenancy

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewTenancyTenantsDeleteParams creates a new TenancyTenantsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTenancyTenantsDeleteParams() *TenancyTenantsDeleteParams {
	return &TenancyTenantsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTenancyTenantsDeleteParamsWithTimeout creates a new TenancyTenantsDeleteParams object
// with the ability to set a timeout on a request.
func NewTenancyTenantsDeleteParamsWithTimeout(timeout time.Duration) *TenancyTenantsDeleteParams {
	return &TenancyTenantsDeleteParams{
		timeout: timeout,
	}
}

// NewTenancyTenantsDeleteParamsWithContext creates a new TenancyTenantsDeleteParams object
// with the ability to set a context for a request.
func NewTenancyTenantsDeleteParamsWithContext(ctx context.Context) *TenancyTenantsDeleteParams {
	return &TenancyTenantsDeleteParams{
		Context: ctx,
	}
}

// NewTenancyTenantsDeleteParamsWithHTTPClient creates a new TenancyTenantsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewTenancyTenantsDeleteParamsWithHTTPClient(client *http.Client) *TenancyTenantsDeleteParams {
	return &TenancyTenantsDeleteParams{
		HTTPClient: client,
	}
}

/* TenancyTenantsDeleteParams contains all the parameters to send to the API endpoint
   for the tenancy tenants delete operation.

   Typically these are written to a http.Request.
*/
type TenancyTenantsDeleteParams struct {

	/* ID.

	   A UUID string identifying this tenant.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tenancy tenants delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyTenantsDeleteParams) WithDefaults() *TenancyTenantsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tenancy tenants delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyTenantsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tenancy tenants delete params
func (o *TenancyTenantsDeleteParams) WithTimeout(timeout time.Duration) *TenancyTenantsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tenancy tenants delete params
func (o *TenancyTenantsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tenancy tenants delete params
func (o *TenancyTenantsDeleteParams) WithContext(ctx context.Context) *TenancyTenantsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tenancy tenants delete params
func (o *TenancyTenantsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tenancy tenants delete params
func (o *TenancyTenantsDeleteParams) WithHTTPClient(client *http.Client) *TenancyTenantsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tenancy tenants delete params
func (o *TenancyTenantsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the tenancy tenants delete params
func (o *TenancyTenantsDeleteParams) WithID(id strfmt.UUID) *TenancyTenantsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the tenancy tenants delete params
func (o *TenancyTenantsDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *TenancyTenantsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

package dcim

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDcimDeviceRolesDeleteParams creates a new DcimDeviceRolesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimDeviceRolesDeleteParams() *DcimDeviceRolesDeleteParams {
	return &DcimDeviceRolesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDeviceRolesDeleteParamsWithTimeout creates a new DcimDeviceRolesDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimDeviceRolesDeleteParamsWithTimeout(timeout time.Duration) *DcimDeviceRolesDeleteParams {
	return &DcimDeviceRolesDeleteParams{
		timeout: timeout,
	}
}

// NewDcimDeviceRolesDeleteParamsWithContext creates a new DcimDeviceRolesDeleteParams object
// with the ability to set a context for a request.
func NewDcimDeviceRolesDeleteParamsWithContext(ctx context.Context) *DcimDeviceRolesDeleteParams {
	return &DcimDeviceRolesDeleteParams{
		Context: ctx,
	}
}

// NewDcimDeviceRolesDeleteParamsWithHTTPClient creates a new DcimDeviceRolesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimDeviceRolesDeleteParamsWithHTTPClient(client *http.Client) *DcimDeviceRolesDeleteParams {
	return &DcimDeviceRolesDeleteParams{
		HTTPClient: client,
	}
}

/* DcimDeviceRolesDeleteParams contains all the parameters to send to the API endpoint
   for the dcim device roles delete operation.

   Typically these are written to a http.Request.
*/
type DcimDeviceRolesDeleteParams struct {

	/* ID.

	   A UUID string identifying this device role.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim device roles delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDeviceRolesDeleteParams) WithDefaults() *DcimDeviceRolesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim device roles delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDeviceRolesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim device roles delete params
func (o *DcimDeviceRolesDeleteParams) WithTimeout(timeout time.Duration) *DcimDeviceRolesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim device roles delete params
func (o *DcimDeviceRolesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim device roles delete params
func (o *DcimDeviceRolesDeleteParams) WithContext(ctx context.Context) *DcimDeviceRolesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim device roles delete params
func (o *DcimDeviceRolesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim device roles delete params
func (o *DcimDeviceRolesDeleteParams) WithHTTPClient(client *http.Client) *DcimDeviceRolesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim device roles delete params
func (o *DcimDeviceRolesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim device roles delete params
func (o *DcimDeviceRolesDeleteParams) WithID(id strfmt.UUID) *DcimDeviceRolesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim device roles delete params
func (o *DcimDeviceRolesDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDeviceRolesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

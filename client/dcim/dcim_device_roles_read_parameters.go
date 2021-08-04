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

// NewDcimDeviceRolesReadParams creates a new DcimDeviceRolesReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimDeviceRolesReadParams() *DcimDeviceRolesReadParams {
	return &DcimDeviceRolesReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDeviceRolesReadParamsWithTimeout creates a new DcimDeviceRolesReadParams object
// with the ability to set a timeout on a request.
func NewDcimDeviceRolesReadParamsWithTimeout(timeout time.Duration) *DcimDeviceRolesReadParams {
	return &DcimDeviceRolesReadParams{
		timeout: timeout,
	}
}

// NewDcimDeviceRolesReadParamsWithContext creates a new DcimDeviceRolesReadParams object
// with the ability to set a context for a request.
func NewDcimDeviceRolesReadParamsWithContext(ctx context.Context) *DcimDeviceRolesReadParams {
	return &DcimDeviceRolesReadParams{
		Context: ctx,
	}
}

// NewDcimDeviceRolesReadParamsWithHTTPClient creates a new DcimDeviceRolesReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimDeviceRolesReadParamsWithHTTPClient(client *http.Client) *DcimDeviceRolesReadParams {
	return &DcimDeviceRolesReadParams{
		HTTPClient: client,
	}
}

/* DcimDeviceRolesReadParams contains all the parameters to send to the API endpoint
   for the dcim device roles read operation.

   Typically these are written to a http.Request.
*/
type DcimDeviceRolesReadParams struct {

	/* ID.

	   A UUID string identifying this device role.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim device roles read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDeviceRolesReadParams) WithDefaults() *DcimDeviceRolesReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim device roles read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDeviceRolesReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim device roles read params
func (o *DcimDeviceRolesReadParams) WithTimeout(timeout time.Duration) *DcimDeviceRolesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim device roles read params
func (o *DcimDeviceRolesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim device roles read params
func (o *DcimDeviceRolesReadParams) WithContext(ctx context.Context) *DcimDeviceRolesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim device roles read params
func (o *DcimDeviceRolesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim device roles read params
func (o *DcimDeviceRolesReadParams) WithHTTPClient(client *http.Client) *DcimDeviceRolesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim device roles read params
func (o *DcimDeviceRolesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim device roles read params
func (o *DcimDeviceRolesReadParams) WithID(id strfmt.UUID) *DcimDeviceRolesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim device roles read params
func (o *DcimDeviceRolesReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDeviceRolesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

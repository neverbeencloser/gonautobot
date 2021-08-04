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

// NewDcimDevicesDeleteParams creates a new DcimDevicesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimDevicesDeleteParams() *DcimDevicesDeleteParams {
	return &DcimDevicesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDevicesDeleteParamsWithTimeout creates a new DcimDevicesDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimDevicesDeleteParamsWithTimeout(timeout time.Duration) *DcimDevicesDeleteParams {
	return &DcimDevicesDeleteParams{
		timeout: timeout,
	}
}

// NewDcimDevicesDeleteParamsWithContext creates a new DcimDevicesDeleteParams object
// with the ability to set a context for a request.
func NewDcimDevicesDeleteParamsWithContext(ctx context.Context) *DcimDevicesDeleteParams {
	return &DcimDevicesDeleteParams{
		Context: ctx,
	}
}

// NewDcimDevicesDeleteParamsWithHTTPClient creates a new DcimDevicesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimDevicesDeleteParamsWithHTTPClient(client *http.Client) *DcimDevicesDeleteParams {
	return &DcimDevicesDeleteParams{
		HTTPClient: client,
	}
}

/* DcimDevicesDeleteParams contains all the parameters to send to the API endpoint
   for the dcim devices delete operation.

   Typically these are written to a http.Request.
*/
type DcimDevicesDeleteParams struct {

	/* ID.

	   A UUID string identifying this device.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim devices delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDevicesDeleteParams) WithDefaults() *DcimDevicesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim devices delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDevicesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim devices delete params
func (o *DcimDevicesDeleteParams) WithTimeout(timeout time.Duration) *DcimDevicesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim devices delete params
func (o *DcimDevicesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim devices delete params
func (o *DcimDevicesDeleteParams) WithContext(ctx context.Context) *DcimDevicesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim devices delete params
func (o *DcimDevicesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim devices delete params
func (o *DcimDevicesDeleteParams) WithHTTPClient(client *http.Client) *DcimDevicesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim devices delete params
func (o *DcimDevicesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim devices delete params
func (o *DcimDevicesDeleteParams) WithID(id strfmt.UUID) *DcimDevicesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim devices delete params
func (o *DcimDevicesDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDevicesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

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

// NewDcimDevicesNapalmParams creates a new DcimDevicesNapalmParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimDevicesNapalmParams() *DcimDevicesNapalmParams {
	return &DcimDevicesNapalmParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDevicesNapalmParamsWithTimeout creates a new DcimDevicesNapalmParams object
// with the ability to set a timeout on a request.
func NewDcimDevicesNapalmParamsWithTimeout(timeout time.Duration) *DcimDevicesNapalmParams {
	return &DcimDevicesNapalmParams{
		timeout: timeout,
	}
}

// NewDcimDevicesNapalmParamsWithContext creates a new DcimDevicesNapalmParams object
// with the ability to set a context for a request.
func NewDcimDevicesNapalmParamsWithContext(ctx context.Context) *DcimDevicesNapalmParams {
	return &DcimDevicesNapalmParams{
		Context: ctx,
	}
}

// NewDcimDevicesNapalmParamsWithHTTPClient creates a new DcimDevicesNapalmParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimDevicesNapalmParamsWithHTTPClient(client *http.Client) *DcimDevicesNapalmParams {
	return &DcimDevicesNapalmParams{
		HTTPClient: client,
	}
}

/* DcimDevicesNapalmParams contains all the parameters to send to the API endpoint
   for the dcim devices napalm operation.

   Typically these are written to a http.Request.
*/
type DcimDevicesNapalmParams struct {

	/* ID.

	   A UUID string identifying this device.

	   Format: uuid
	*/
	ID strfmt.UUID

	// Method.
	Method string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim devices napalm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDevicesNapalmParams) WithDefaults() *DcimDevicesNapalmParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim devices napalm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDevicesNapalmParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim devices napalm params
func (o *DcimDevicesNapalmParams) WithTimeout(timeout time.Duration) *DcimDevicesNapalmParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim devices napalm params
func (o *DcimDevicesNapalmParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim devices napalm params
func (o *DcimDevicesNapalmParams) WithContext(ctx context.Context) *DcimDevicesNapalmParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim devices napalm params
func (o *DcimDevicesNapalmParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim devices napalm params
func (o *DcimDevicesNapalmParams) WithHTTPClient(client *http.Client) *DcimDevicesNapalmParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim devices napalm params
func (o *DcimDevicesNapalmParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim devices napalm params
func (o *DcimDevicesNapalmParams) WithID(id strfmt.UUID) *DcimDevicesNapalmParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim devices napalm params
func (o *DcimDevicesNapalmParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithMethod adds the method to the dcim devices napalm params
func (o *DcimDevicesNapalmParams) WithMethod(method string) *DcimDevicesNapalmParams {
	o.SetMethod(method)
	return o
}

// SetMethod adds the method to the dcim devices napalm params
func (o *DcimDevicesNapalmParams) SetMethod(method string) {
	o.Method = method
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDevicesNapalmParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// query param method
	qrMethod := o.Method
	qMethod := qrMethod
	if qMethod != "" {

		if err := r.SetQueryParam("method", qMethod); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

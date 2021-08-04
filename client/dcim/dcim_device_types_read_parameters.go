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

// NewDcimDeviceTypesReadParams creates a new DcimDeviceTypesReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimDeviceTypesReadParams() *DcimDeviceTypesReadParams {
	return &DcimDeviceTypesReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDeviceTypesReadParamsWithTimeout creates a new DcimDeviceTypesReadParams object
// with the ability to set a timeout on a request.
func NewDcimDeviceTypesReadParamsWithTimeout(timeout time.Duration) *DcimDeviceTypesReadParams {
	return &DcimDeviceTypesReadParams{
		timeout: timeout,
	}
}

// NewDcimDeviceTypesReadParamsWithContext creates a new DcimDeviceTypesReadParams object
// with the ability to set a context for a request.
func NewDcimDeviceTypesReadParamsWithContext(ctx context.Context) *DcimDeviceTypesReadParams {
	return &DcimDeviceTypesReadParams{
		Context: ctx,
	}
}

// NewDcimDeviceTypesReadParamsWithHTTPClient creates a new DcimDeviceTypesReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimDeviceTypesReadParamsWithHTTPClient(client *http.Client) *DcimDeviceTypesReadParams {
	return &DcimDeviceTypesReadParams{
		HTTPClient: client,
	}
}

/* DcimDeviceTypesReadParams contains all the parameters to send to the API endpoint
   for the dcim device types read operation.

   Typically these are written to a http.Request.
*/
type DcimDeviceTypesReadParams struct {

	/* ID.

	   A UUID string identifying this device type.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim device types read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDeviceTypesReadParams) WithDefaults() *DcimDeviceTypesReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim device types read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDeviceTypesReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim device types read params
func (o *DcimDeviceTypesReadParams) WithTimeout(timeout time.Duration) *DcimDeviceTypesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim device types read params
func (o *DcimDeviceTypesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim device types read params
func (o *DcimDeviceTypesReadParams) WithContext(ctx context.Context) *DcimDeviceTypesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim device types read params
func (o *DcimDeviceTypesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim device types read params
func (o *DcimDeviceTypesReadParams) WithHTTPClient(client *http.Client) *DcimDeviceTypesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim device types read params
func (o *DcimDeviceTypesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim device types read params
func (o *DcimDeviceTypesReadParams) WithID(id strfmt.UUID) *DcimDeviceTypesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim device types read params
func (o *DcimDeviceTypesReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDeviceTypesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

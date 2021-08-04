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

// NewDcimRearPortsDeleteParams creates a new DcimRearPortsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRearPortsDeleteParams() *DcimRearPortsDeleteParams {
	return &DcimRearPortsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRearPortsDeleteParamsWithTimeout creates a new DcimRearPortsDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimRearPortsDeleteParamsWithTimeout(timeout time.Duration) *DcimRearPortsDeleteParams {
	return &DcimRearPortsDeleteParams{
		timeout: timeout,
	}
}

// NewDcimRearPortsDeleteParamsWithContext creates a new DcimRearPortsDeleteParams object
// with the ability to set a context for a request.
func NewDcimRearPortsDeleteParamsWithContext(ctx context.Context) *DcimRearPortsDeleteParams {
	return &DcimRearPortsDeleteParams{
		Context: ctx,
	}
}

// NewDcimRearPortsDeleteParamsWithHTTPClient creates a new DcimRearPortsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRearPortsDeleteParamsWithHTTPClient(client *http.Client) *DcimRearPortsDeleteParams {
	return &DcimRearPortsDeleteParams{
		HTTPClient: client,
	}
}

/* DcimRearPortsDeleteParams contains all the parameters to send to the API endpoint
   for the dcim rear ports delete operation.

   Typically these are written to a http.Request.
*/
type DcimRearPortsDeleteParams struct {

	/* ID.

	   A UUID string identifying this rear port.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim rear ports delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRearPortsDeleteParams) WithDefaults() *DcimRearPortsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim rear ports delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRearPortsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim rear ports delete params
func (o *DcimRearPortsDeleteParams) WithTimeout(timeout time.Duration) *DcimRearPortsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rear ports delete params
func (o *DcimRearPortsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rear ports delete params
func (o *DcimRearPortsDeleteParams) WithContext(ctx context.Context) *DcimRearPortsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rear ports delete params
func (o *DcimRearPortsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rear ports delete params
func (o *DcimRearPortsDeleteParams) WithHTTPClient(client *http.Client) *DcimRearPortsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rear ports delete params
func (o *DcimRearPortsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim rear ports delete params
func (o *DcimRearPortsDeleteParams) WithID(id strfmt.UUID) *DcimRearPortsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim rear ports delete params
func (o *DcimRearPortsDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRearPortsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

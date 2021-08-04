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

// NewDcimFrontPortsDeleteParams creates a new DcimFrontPortsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimFrontPortsDeleteParams() *DcimFrontPortsDeleteParams {
	return &DcimFrontPortsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimFrontPortsDeleteParamsWithTimeout creates a new DcimFrontPortsDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimFrontPortsDeleteParamsWithTimeout(timeout time.Duration) *DcimFrontPortsDeleteParams {
	return &DcimFrontPortsDeleteParams{
		timeout: timeout,
	}
}

// NewDcimFrontPortsDeleteParamsWithContext creates a new DcimFrontPortsDeleteParams object
// with the ability to set a context for a request.
func NewDcimFrontPortsDeleteParamsWithContext(ctx context.Context) *DcimFrontPortsDeleteParams {
	return &DcimFrontPortsDeleteParams{
		Context: ctx,
	}
}

// NewDcimFrontPortsDeleteParamsWithHTTPClient creates a new DcimFrontPortsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimFrontPortsDeleteParamsWithHTTPClient(client *http.Client) *DcimFrontPortsDeleteParams {
	return &DcimFrontPortsDeleteParams{
		HTTPClient: client,
	}
}

/* DcimFrontPortsDeleteParams contains all the parameters to send to the API endpoint
   for the dcim front ports delete operation.

   Typically these are written to a http.Request.
*/
type DcimFrontPortsDeleteParams struct {

	/* ID.

	   A UUID string identifying this front port.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim front ports delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimFrontPortsDeleteParams) WithDefaults() *DcimFrontPortsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim front ports delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimFrontPortsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim front ports delete params
func (o *DcimFrontPortsDeleteParams) WithTimeout(timeout time.Duration) *DcimFrontPortsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim front ports delete params
func (o *DcimFrontPortsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim front ports delete params
func (o *DcimFrontPortsDeleteParams) WithContext(ctx context.Context) *DcimFrontPortsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim front ports delete params
func (o *DcimFrontPortsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim front ports delete params
func (o *DcimFrontPortsDeleteParams) WithHTTPClient(client *http.Client) *DcimFrontPortsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim front ports delete params
func (o *DcimFrontPortsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim front ports delete params
func (o *DcimFrontPortsDeleteParams) WithID(id strfmt.UUID) *DcimFrontPortsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim front ports delete params
func (o *DcimFrontPortsDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimFrontPortsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

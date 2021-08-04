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

// NewDcimVirtualChassisReadParams creates a new DcimVirtualChassisReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimVirtualChassisReadParams() *DcimVirtualChassisReadParams {
	return &DcimVirtualChassisReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimVirtualChassisReadParamsWithTimeout creates a new DcimVirtualChassisReadParams object
// with the ability to set a timeout on a request.
func NewDcimVirtualChassisReadParamsWithTimeout(timeout time.Duration) *DcimVirtualChassisReadParams {
	return &DcimVirtualChassisReadParams{
		timeout: timeout,
	}
}

// NewDcimVirtualChassisReadParamsWithContext creates a new DcimVirtualChassisReadParams object
// with the ability to set a context for a request.
func NewDcimVirtualChassisReadParamsWithContext(ctx context.Context) *DcimVirtualChassisReadParams {
	return &DcimVirtualChassisReadParams{
		Context: ctx,
	}
}

// NewDcimVirtualChassisReadParamsWithHTTPClient creates a new DcimVirtualChassisReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimVirtualChassisReadParamsWithHTTPClient(client *http.Client) *DcimVirtualChassisReadParams {
	return &DcimVirtualChassisReadParams{
		HTTPClient: client,
	}
}

/* DcimVirtualChassisReadParams contains all the parameters to send to the API endpoint
   for the dcim virtual chassis read operation.

   Typically these are written to a http.Request.
*/
type DcimVirtualChassisReadParams struct {

	/* ID.

	   A UUID string identifying this virtual chassis.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim virtual chassis read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimVirtualChassisReadParams) WithDefaults() *DcimVirtualChassisReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim virtual chassis read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimVirtualChassisReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim virtual chassis read params
func (o *DcimVirtualChassisReadParams) WithTimeout(timeout time.Duration) *DcimVirtualChassisReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim virtual chassis read params
func (o *DcimVirtualChassisReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim virtual chassis read params
func (o *DcimVirtualChassisReadParams) WithContext(ctx context.Context) *DcimVirtualChassisReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim virtual chassis read params
func (o *DcimVirtualChassisReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim virtual chassis read params
func (o *DcimVirtualChassisReadParams) WithHTTPClient(client *http.Client) *DcimVirtualChassisReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim virtual chassis read params
func (o *DcimVirtualChassisReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim virtual chassis read params
func (o *DcimVirtualChassisReadParams) WithID(id strfmt.UUID) *DcimVirtualChassisReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim virtual chassis read params
func (o *DcimVirtualChassisReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimVirtualChassisReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

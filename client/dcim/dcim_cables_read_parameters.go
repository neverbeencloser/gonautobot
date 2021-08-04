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

// NewDcimCablesReadParams creates a new DcimCablesReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimCablesReadParams() *DcimCablesReadParams {
	return &DcimCablesReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimCablesReadParamsWithTimeout creates a new DcimCablesReadParams object
// with the ability to set a timeout on a request.
func NewDcimCablesReadParamsWithTimeout(timeout time.Duration) *DcimCablesReadParams {
	return &DcimCablesReadParams{
		timeout: timeout,
	}
}

// NewDcimCablesReadParamsWithContext creates a new DcimCablesReadParams object
// with the ability to set a context for a request.
func NewDcimCablesReadParamsWithContext(ctx context.Context) *DcimCablesReadParams {
	return &DcimCablesReadParams{
		Context: ctx,
	}
}

// NewDcimCablesReadParamsWithHTTPClient creates a new DcimCablesReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimCablesReadParamsWithHTTPClient(client *http.Client) *DcimCablesReadParams {
	return &DcimCablesReadParams{
		HTTPClient: client,
	}
}

/* DcimCablesReadParams contains all the parameters to send to the API endpoint
   for the dcim cables read operation.

   Typically these are written to a http.Request.
*/
type DcimCablesReadParams struct {

	/* ID.

	   A UUID string identifying this cable.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim cables read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimCablesReadParams) WithDefaults() *DcimCablesReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim cables read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimCablesReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim cables read params
func (o *DcimCablesReadParams) WithTimeout(timeout time.Duration) *DcimCablesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim cables read params
func (o *DcimCablesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim cables read params
func (o *DcimCablesReadParams) WithContext(ctx context.Context) *DcimCablesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim cables read params
func (o *DcimCablesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim cables read params
func (o *DcimCablesReadParams) WithHTTPClient(client *http.Client) *DcimCablesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim cables read params
func (o *DcimCablesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim cables read params
func (o *DcimCablesReadParams) WithID(id strfmt.UUID) *DcimCablesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim cables read params
func (o *DcimCablesReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimCablesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

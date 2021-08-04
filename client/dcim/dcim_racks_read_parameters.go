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

// NewDcimRacksReadParams creates a new DcimRacksReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRacksReadParams() *DcimRacksReadParams {
	return &DcimRacksReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRacksReadParamsWithTimeout creates a new DcimRacksReadParams object
// with the ability to set a timeout on a request.
func NewDcimRacksReadParamsWithTimeout(timeout time.Duration) *DcimRacksReadParams {
	return &DcimRacksReadParams{
		timeout: timeout,
	}
}

// NewDcimRacksReadParamsWithContext creates a new DcimRacksReadParams object
// with the ability to set a context for a request.
func NewDcimRacksReadParamsWithContext(ctx context.Context) *DcimRacksReadParams {
	return &DcimRacksReadParams{
		Context: ctx,
	}
}

// NewDcimRacksReadParamsWithHTTPClient creates a new DcimRacksReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRacksReadParamsWithHTTPClient(client *http.Client) *DcimRacksReadParams {
	return &DcimRacksReadParams{
		HTTPClient: client,
	}
}

/* DcimRacksReadParams contains all the parameters to send to the API endpoint
   for the dcim racks read operation.

   Typically these are written to a http.Request.
*/
type DcimRacksReadParams struct {

	/* ID.

	   A UUID string identifying this rack.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim racks read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRacksReadParams) WithDefaults() *DcimRacksReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim racks read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRacksReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim racks read params
func (o *DcimRacksReadParams) WithTimeout(timeout time.Duration) *DcimRacksReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim racks read params
func (o *DcimRacksReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim racks read params
func (o *DcimRacksReadParams) WithContext(ctx context.Context) *DcimRacksReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim racks read params
func (o *DcimRacksReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim racks read params
func (o *DcimRacksReadParams) WithHTTPClient(client *http.Client) *DcimRacksReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim racks read params
func (o *DcimRacksReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim racks read params
func (o *DcimRacksReadParams) WithID(id strfmt.UUID) *DcimRacksReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim racks read params
func (o *DcimRacksReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRacksReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

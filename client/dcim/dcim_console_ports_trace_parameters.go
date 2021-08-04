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

// NewDcimConsolePortsTraceParams creates a new DcimConsolePortsTraceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimConsolePortsTraceParams() *DcimConsolePortsTraceParams {
	return &DcimConsolePortsTraceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimConsolePortsTraceParamsWithTimeout creates a new DcimConsolePortsTraceParams object
// with the ability to set a timeout on a request.
func NewDcimConsolePortsTraceParamsWithTimeout(timeout time.Duration) *DcimConsolePortsTraceParams {
	return &DcimConsolePortsTraceParams{
		timeout: timeout,
	}
}

// NewDcimConsolePortsTraceParamsWithContext creates a new DcimConsolePortsTraceParams object
// with the ability to set a context for a request.
func NewDcimConsolePortsTraceParamsWithContext(ctx context.Context) *DcimConsolePortsTraceParams {
	return &DcimConsolePortsTraceParams{
		Context: ctx,
	}
}

// NewDcimConsolePortsTraceParamsWithHTTPClient creates a new DcimConsolePortsTraceParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimConsolePortsTraceParamsWithHTTPClient(client *http.Client) *DcimConsolePortsTraceParams {
	return &DcimConsolePortsTraceParams{
		HTTPClient: client,
	}
}

/* DcimConsolePortsTraceParams contains all the parameters to send to the API endpoint
   for the dcim console ports trace operation.

   Typically these are written to a http.Request.
*/
type DcimConsolePortsTraceParams struct {

	/* ID.

	   A UUID string identifying this console port.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim console ports trace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsolePortsTraceParams) WithDefaults() *DcimConsolePortsTraceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim console ports trace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsolePortsTraceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim console ports trace params
func (o *DcimConsolePortsTraceParams) WithTimeout(timeout time.Duration) *DcimConsolePortsTraceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim console ports trace params
func (o *DcimConsolePortsTraceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim console ports trace params
func (o *DcimConsolePortsTraceParams) WithContext(ctx context.Context) *DcimConsolePortsTraceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim console ports trace params
func (o *DcimConsolePortsTraceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim console ports trace params
func (o *DcimConsolePortsTraceParams) WithHTTPClient(client *http.Client) *DcimConsolePortsTraceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim console ports trace params
func (o *DcimConsolePortsTraceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim console ports trace params
func (o *DcimConsolePortsTraceParams) WithID(id strfmt.UUID) *DcimConsolePortsTraceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim console ports trace params
func (o *DcimConsolePortsTraceParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimConsolePortsTraceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

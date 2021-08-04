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

// NewDcimConsoleServerPortsDeleteParams creates a new DcimConsoleServerPortsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimConsoleServerPortsDeleteParams() *DcimConsoleServerPortsDeleteParams {
	return &DcimConsoleServerPortsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimConsoleServerPortsDeleteParamsWithTimeout creates a new DcimConsoleServerPortsDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimConsoleServerPortsDeleteParamsWithTimeout(timeout time.Duration) *DcimConsoleServerPortsDeleteParams {
	return &DcimConsoleServerPortsDeleteParams{
		timeout: timeout,
	}
}

// NewDcimConsoleServerPortsDeleteParamsWithContext creates a new DcimConsoleServerPortsDeleteParams object
// with the ability to set a context for a request.
func NewDcimConsoleServerPortsDeleteParamsWithContext(ctx context.Context) *DcimConsoleServerPortsDeleteParams {
	return &DcimConsoleServerPortsDeleteParams{
		Context: ctx,
	}
}

// NewDcimConsoleServerPortsDeleteParamsWithHTTPClient creates a new DcimConsoleServerPortsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimConsoleServerPortsDeleteParamsWithHTTPClient(client *http.Client) *DcimConsoleServerPortsDeleteParams {
	return &DcimConsoleServerPortsDeleteParams{
		HTTPClient: client,
	}
}

/* DcimConsoleServerPortsDeleteParams contains all the parameters to send to the API endpoint
   for the dcim console server ports delete operation.

   Typically these are written to a http.Request.
*/
type DcimConsoleServerPortsDeleteParams struct {

	/* ID.

	   A UUID string identifying this console server port.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim console server ports delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsoleServerPortsDeleteParams) WithDefaults() *DcimConsoleServerPortsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim console server ports delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsoleServerPortsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim console server ports delete params
func (o *DcimConsoleServerPortsDeleteParams) WithTimeout(timeout time.Duration) *DcimConsoleServerPortsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim console server ports delete params
func (o *DcimConsoleServerPortsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim console server ports delete params
func (o *DcimConsoleServerPortsDeleteParams) WithContext(ctx context.Context) *DcimConsoleServerPortsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim console server ports delete params
func (o *DcimConsoleServerPortsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim console server ports delete params
func (o *DcimConsoleServerPortsDeleteParams) WithHTTPClient(client *http.Client) *DcimConsoleServerPortsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim console server ports delete params
func (o *DcimConsoleServerPortsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim console server ports delete params
func (o *DcimConsoleServerPortsDeleteParams) WithID(id strfmt.UUID) *DcimConsoleServerPortsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim console server ports delete params
func (o *DcimConsoleServerPortsDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimConsoleServerPortsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

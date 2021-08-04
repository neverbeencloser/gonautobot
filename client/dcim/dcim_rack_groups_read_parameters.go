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

// NewDcimRackGroupsReadParams creates a new DcimRackGroupsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRackGroupsReadParams() *DcimRackGroupsReadParams {
	return &DcimRackGroupsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRackGroupsReadParamsWithTimeout creates a new DcimRackGroupsReadParams object
// with the ability to set a timeout on a request.
func NewDcimRackGroupsReadParamsWithTimeout(timeout time.Duration) *DcimRackGroupsReadParams {
	return &DcimRackGroupsReadParams{
		timeout: timeout,
	}
}

// NewDcimRackGroupsReadParamsWithContext creates a new DcimRackGroupsReadParams object
// with the ability to set a context for a request.
func NewDcimRackGroupsReadParamsWithContext(ctx context.Context) *DcimRackGroupsReadParams {
	return &DcimRackGroupsReadParams{
		Context: ctx,
	}
}

// NewDcimRackGroupsReadParamsWithHTTPClient creates a new DcimRackGroupsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRackGroupsReadParamsWithHTTPClient(client *http.Client) *DcimRackGroupsReadParams {
	return &DcimRackGroupsReadParams{
		HTTPClient: client,
	}
}

/* DcimRackGroupsReadParams contains all the parameters to send to the API endpoint
   for the dcim rack groups read operation.

   Typically these are written to a http.Request.
*/
type DcimRackGroupsReadParams struct {

	/* ID.

	   A UUID string identifying this rack group.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim rack groups read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRackGroupsReadParams) WithDefaults() *DcimRackGroupsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim rack groups read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRackGroupsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim rack groups read params
func (o *DcimRackGroupsReadParams) WithTimeout(timeout time.Duration) *DcimRackGroupsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rack groups read params
func (o *DcimRackGroupsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rack groups read params
func (o *DcimRackGroupsReadParams) WithContext(ctx context.Context) *DcimRackGroupsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rack groups read params
func (o *DcimRackGroupsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rack groups read params
func (o *DcimRackGroupsReadParams) WithHTTPClient(client *http.Client) *DcimRackGroupsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rack groups read params
func (o *DcimRackGroupsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim rack groups read params
func (o *DcimRackGroupsReadParams) WithID(id strfmt.UUID) *DcimRackGroupsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim rack groups read params
func (o *DcimRackGroupsReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRackGroupsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

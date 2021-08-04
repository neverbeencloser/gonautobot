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

// NewDcimRackGroupsDeleteParams creates a new DcimRackGroupsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRackGroupsDeleteParams() *DcimRackGroupsDeleteParams {
	return &DcimRackGroupsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRackGroupsDeleteParamsWithTimeout creates a new DcimRackGroupsDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimRackGroupsDeleteParamsWithTimeout(timeout time.Duration) *DcimRackGroupsDeleteParams {
	return &DcimRackGroupsDeleteParams{
		timeout: timeout,
	}
}

// NewDcimRackGroupsDeleteParamsWithContext creates a new DcimRackGroupsDeleteParams object
// with the ability to set a context for a request.
func NewDcimRackGroupsDeleteParamsWithContext(ctx context.Context) *DcimRackGroupsDeleteParams {
	return &DcimRackGroupsDeleteParams{
		Context: ctx,
	}
}

// NewDcimRackGroupsDeleteParamsWithHTTPClient creates a new DcimRackGroupsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRackGroupsDeleteParamsWithHTTPClient(client *http.Client) *DcimRackGroupsDeleteParams {
	return &DcimRackGroupsDeleteParams{
		HTTPClient: client,
	}
}

/* DcimRackGroupsDeleteParams contains all the parameters to send to the API endpoint
   for the dcim rack groups delete operation.

   Typically these are written to a http.Request.
*/
type DcimRackGroupsDeleteParams struct {

	/* ID.

	   A UUID string identifying this rack group.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim rack groups delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRackGroupsDeleteParams) WithDefaults() *DcimRackGroupsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim rack groups delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRackGroupsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim rack groups delete params
func (o *DcimRackGroupsDeleteParams) WithTimeout(timeout time.Duration) *DcimRackGroupsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rack groups delete params
func (o *DcimRackGroupsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rack groups delete params
func (o *DcimRackGroupsDeleteParams) WithContext(ctx context.Context) *DcimRackGroupsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rack groups delete params
func (o *DcimRackGroupsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rack groups delete params
func (o *DcimRackGroupsDeleteParams) WithHTTPClient(client *http.Client) *DcimRackGroupsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rack groups delete params
func (o *DcimRackGroupsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim rack groups delete params
func (o *DcimRackGroupsDeleteParams) WithID(id strfmt.UUID) *DcimRackGroupsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim rack groups delete params
func (o *DcimRackGroupsDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRackGroupsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

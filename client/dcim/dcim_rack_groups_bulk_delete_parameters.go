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

// NewDcimRackGroupsBulkDeleteParams creates a new DcimRackGroupsBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRackGroupsBulkDeleteParams() *DcimRackGroupsBulkDeleteParams {
	return &DcimRackGroupsBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRackGroupsBulkDeleteParamsWithTimeout creates a new DcimRackGroupsBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimRackGroupsBulkDeleteParamsWithTimeout(timeout time.Duration) *DcimRackGroupsBulkDeleteParams {
	return &DcimRackGroupsBulkDeleteParams{
		timeout: timeout,
	}
}

// NewDcimRackGroupsBulkDeleteParamsWithContext creates a new DcimRackGroupsBulkDeleteParams object
// with the ability to set a context for a request.
func NewDcimRackGroupsBulkDeleteParamsWithContext(ctx context.Context) *DcimRackGroupsBulkDeleteParams {
	return &DcimRackGroupsBulkDeleteParams{
		Context: ctx,
	}
}

// NewDcimRackGroupsBulkDeleteParamsWithHTTPClient creates a new DcimRackGroupsBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRackGroupsBulkDeleteParamsWithHTTPClient(client *http.Client) *DcimRackGroupsBulkDeleteParams {
	return &DcimRackGroupsBulkDeleteParams{
		HTTPClient: client,
	}
}

/* DcimRackGroupsBulkDeleteParams contains all the parameters to send to the API endpoint
   for the dcim rack groups bulk delete operation.

   Typically these are written to a http.Request.
*/
type DcimRackGroupsBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim rack groups bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRackGroupsBulkDeleteParams) WithDefaults() *DcimRackGroupsBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim rack groups bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRackGroupsBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim rack groups bulk delete params
func (o *DcimRackGroupsBulkDeleteParams) WithTimeout(timeout time.Duration) *DcimRackGroupsBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rack groups bulk delete params
func (o *DcimRackGroupsBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rack groups bulk delete params
func (o *DcimRackGroupsBulkDeleteParams) WithContext(ctx context.Context) *DcimRackGroupsBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rack groups bulk delete params
func (o *DcimRackGroupsBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rack groups bulk delete params
func (o *DcimRackGroupsBulkDeleteParams) WithHTTPClient(client *http.Client) *DcimRackGroupsBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rack groups bulk delete params
func (o *DcimRackGroupsBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRackGroupsBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

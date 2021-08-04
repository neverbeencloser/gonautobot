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

// NewDcimRearPortTemplatesDeleteParams creates a new DcimRearPortTemplatesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRearPortTemplatesDeleteParams() *DcimRearPortTemplatesDeleteParams {
	return &DcimRearPortTemplatesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRearPortTemplatesDeleteParamsWithTimeout creates a new DcimRearPortTemplatesDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimRearPortTemplatesDeleteParamsWithTimeout(timeout time.Duration) *DcimRearPortTemplatesDeleteParams {
	return &DcimRearPortTemplatesDeleteParams{
		timeout: timeout,
	}
}

// NewDcimRearPortTemplatesDeleteParamsWithContext creates a new DcimRearPortTemplatesDeleteParams object
// with the ability to set a context for a request.
func NewDcimRearPortTemplatesDeleteParamsWithContext(ctx context.Context) *DcimRearPortTemplatesDeleteParams {
	return &DcimRearPortTemplatesDeleteParams{
		Context: ctx,
	}
}

// NewDcimRearPortTemplatesDeleteParamsWithHTTPClient creates a new DcimRearPortTemplatesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRearPortTemplatesDeleteParamsWithHTTPClient(client *http.Client) *DcimRearPortTemplatesDeleteParams {
	return &DcimRearPortTemplatesDeleteParams{
		HTTPClient: client,
	}
}

/* DcimRearPortTemplatesDeleteParams contains all the parameters to send to the API endpoint
   for the dcim rear port templates delete operation.

   Typically these are written to a http.Request.
*/
type DcimRearPortTemplatesDeleteParams struct {

	/* ID.

	   A UUID string identifying this rear port template.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim rear port templates delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRearPortTemplatesDeleteParams) WithDefaults() *DcimRearPortTemplatesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim rear port templates delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRearPortTemplatesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim rear port templates delete params
func (o *DcimRearPortTemplatesDeleteParams) WithTimeout(timeout time.Duration) *DcimRearPortTemplatesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rear port templates delete params
func (o *DcimRearPortTemplatesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rear port templates delete params
func (o *DcimRearPortTemplatesDeleteParams) WithContext(ctx context.Context) *DcimRearPortTemplatesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rear port templates delete params
func (o *DcimRearPortTemplatesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rear port templates delete params
func (o *DcimRearPortTemplatesDeleteParams) WithHTTPClient(client *http.Client) *DcimRearPortTemplatesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rear port templates delete params
func (o *DcimRearPortTemplatesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim rear port templates delete params
func (o *DcimRearPortTemplatesDeleteParams) WithID(id strfmt.UUID) *DcimRearPortTemplatesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim rear port templates delete params
func (o *DcimRearPortTemplatesDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRearPortTemplatesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

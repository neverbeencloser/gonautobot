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

// NewDcimDeviceBayTemplatesReadParams creates a new DcimDeviceBayTemplatesReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimDeviceBayTemplatesReadParams() *DcimDeviceBayTemplatesReadParams {
	return &DcimDeviceBayTemplatesReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDeviceBayTemplatesReadParamsWithTimeout creates a new DcimDeviceBayTemplatesReadParams object
// with the ability to set a timeout on a request.
func NewDcimDeviceBayTemplatesReadParamsWithTimeout(timeout time.Duration) *DcimDeviceBayTemplatesReadParams {
	return &DcimDeviceBayTemplatesReadParams{
		timeout: timeout,
	}
}

// NewDcimDeviceBayTemplatesReadParamsWithContext creates a new DcimDeviceBayTemplatesReadParams object
// with the ability to set a context for a request.
func NewDcimDeviceBayTemplatesReadParamsWithContext(ctx context.Context) *DcimDeviceBayTemplatesReadParams {
	return &DcimDeviceBayTemplatesReadParams{
		Context: ctx,
	}
}

// NewDcimDeviceBayTemplatesReadParamsWithHTTPClient creates a new DcimDeviceBayTemplatesReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimDeviceBayTemplatesReadParamsWithHTTPClient(client *http.Client) *DcimDeviceBayTemplatesReadParams {
	return &DcimDeviceBayTemplatesReadParams{
		HTTPClient: client,
	}
}

/* DcimDeviceBayTemplatesReadParams contains all the parameters to send to the API endpoint
   for the dcim device bay templates read operation.

   Typically these are written to a http.Request.
*/
type DcimDeviceBayTemplatesReadParams struct {

	/* ID.

	   A UUID string identifying this device bay template.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim device bay templates read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDeviceBayTemplatesReadParams) WithDefaults() *DcimDeviceBayTemplatesReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim device bay templates read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDeviceBayTemplatesReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim device bay templates read params
func (o *DcimDeviceBayTemplatesReadParams) WithTimeout(timeout time.Duration) *DcimDeviceBayTemplatesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim device bay templates read params
func (o *DcimDeviceBayTemplatesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim device bay templates read params
func (o *DcimDeviceBayTemplatesReadParams) WithContext(ctx context.Context) *DcimDeviceBayTemplatesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim device bay templates read params
func (o *DcimDeviceBayTemplatesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim device bay templates read params
func (o *DcimDeviceBayTemplatesReadParams) WithHTTPClient(client *http.Client) *DcimDeviceBayTemplatesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim device bay templates read params
func (o *DcimDeviceBayTemplatesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim device bay templates read params
func (o *DcimDeviceBayTemplatesReadParams) WithID(id strfmt.UUID) *DcimDeviceBayTemplatesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim device bay templates read params
func (o *DcimDeviceBayTemplatesReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDeviceBayTemplatesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

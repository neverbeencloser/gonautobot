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

// NewDcimDeviceBayTemplatesDeleteParams creates a new DcimDeviceBayTemplatesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimDeviceBayTemplatesDeleteParams() *DcimDeviceBayTemplatesDeleteParams {
	return &DcimDeviceBayTemplatesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDeviceBayTemplatesDeleteParamsWithTimeout creates a new DcimDeviceBayTemplatesDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimDeviceBayTemplatesDeleteParamsWithTimeout(timeout time.Duration) *DcimDeviceBayTemplatesDeleteParams {
	return &DcimDeviceBayTemplatesDeleteParams{
		timeout: timeout,
	}
}

// NewDcimDeviceBayTemplatesDeleteParamsWithContext creates a new DcimDeviceBayTemplatesDeleteParams object
// with the ability to set a context for a request.
func NewDcimDeviceBayTemplatesDeleteParamsWithContext(ctx context.Context) *DcimDeviceBayTemplatesDeleteParams {
	return &DcimDeviceBayTemplatesDeleteParams{
		Context: ctx,
	}
}

// NewDcimDeviceBayTemplatesDeleteParamsWithHTTPClient creates a new DcimDeviceBayTemplatesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimDeviceBayTemplatesDeleteParamsWithHTTPClient(client *http.Client) *DcimDeviceBayTemplatesDeleteParams {
	return &DcimDeviceBayTemplatesDeleteParams{
		HTTPClient: client,
	}
}

/* DcimDeviceBayTemplatesDeleteParams contains all the parameters to send to the API endpoint
   for the dcim device bay templates delete operation.

   Typically these are written to a http.Request.
*/
type DcimDeviceBayTemplatesDeleteParams struct {

	/* ID.

	   A UUID string identifying this device bay template.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim device bay templates delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDeviceBayTemplatesDeleteParams) WithDefaults() *DcimDeviceBayTemplatesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim device bay templates delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDeviceBayTemplatesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim device bay templates delete params
func (o *DcimDeviceBayTemplatesDeleteParams) WithTimeout(timeout time.Duration) *DcimDeviceBayTemplatesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim device bay templates delete params
func (o *DcimDeviceBayTemplatesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim device bay templates delete params
func (o *DcimDeviceBayTemplatesDeleteParams) WithContext(ctx context.Context) *DcimDeviceBayTemplatesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim device bay templates delete params
func (o *DcimDeviceBayTemplatesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim device bay templates delete params
func (o *DcimDeviceBayTemplatesDeleteParams) WithHTTPClient(client *http.Client) *DcimDeviceBayTemplatesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim device bay templates delete params
func (o *DcimDeviceBayTemplatesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim device bay templates delete params
func (o *DcimDeviceBayTemplatesDeleteParams) WithID(id strfmt.UUID) *DcimDeviceBayTemplatesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim device bay templates delete params
func (o *DcimDeviceBayTemplatesDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDeviceBayTemplatesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

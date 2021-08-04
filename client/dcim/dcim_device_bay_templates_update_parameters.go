package dcim

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// NewDcimDeviceBayTemplatesUpdateParams creates a new DcimDeviceBayTemplatesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimDeviceBayTemplatesUpdateParams() *DcimDeviceBayTemplatesUpdateParams {
	return &DcimDeviceBayTemplatesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDeviceBayTemplatesUpdateParamsWithTimeout creates a new DcimDeviceBayTemplatesUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimDeviceBayTemplatesUpdateParamsWithTimeout(timeout time.Duration) *DcimDeviceBayTemplatesUpdateParams {
	return &DcimDeviceBayTemplatesUpdateParams{
		timeout: timeout,
	}
}

// NewDcimDeviceBayTemplatesUpdateParamsWithContext creates a new DcimDeviceBayTemplatesUpdateParams object
// with the ability to set a context for a request.
func NewDcimDeviceBayTemplatesUpdateParamsWithContext(ctx context.Context) *DcimDeviceBayTemplatesUpdateParams {
	return &DcimDeviceBayTemplatesUpdateParams{
		Context: ctx,
	}
}

// NewDcimDeviceBayTemplatesUpdateParamsWithHTTPClient creates a new DcimDeviceBayTemplatesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimDeviceBayTemplatesUpdateParamsWithHTTPClient(client *http.Client) *DcimDeviceBayTemplatesUpdateParams {
	return &DcimDeviceBayTemplatesUpdateParams{
		HTTPClient: client,
	}
}

/* DcimDeviceBayTemplatesUpdateParams contains all the parameters to send to the API endpoint
   for the dcim device bay templates update operation.

   Typically these are written to a http.Request.
*/
type DcimDeviceBayTemplatesUpdateParams struct {

	// Data.
	Data *models.WritableDeviceBayTemplate

	/* ID.

	   A UUID string identifying this device bay template.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim device bay templates update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDeviceBayTemplatesUpdateParams) WithDefaults() *DcimDeviceBayTemplatesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim device bay templates update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDeviceBayTemplatesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim device bay templates update params
func (o *DcimDeviceBayTemplatesUpdateParams) WithTimeout(timeout time.Duration) *DcimDeviceBayTemplatesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim device bay templates update params
func (o *DcimDeviceBayTemplatesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim device bay templates update params
func (o *DcimDeviceBayTemplatesUpdateParams) WithContext(ctx context.Context) *DcimDeviceBayTemplatesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim device bay templates update params
func (o *DcimDeviceBayTemplatesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim device bay templates update params
func (o *DcimDeviceBayTemplatesUpdateParams) WithHTTPClient(client *http.Client) *DcimDeviceBayTemplatesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim device bay templates update params
func (o *DcimDeviceBayTemplatesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim device bay templates update params
func (o *DcimDeviceBayTemplatesUpdateParams) WithData(data *models.WritableDeviceBayTemplate) *DcimDeviceBayTemplatesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim device bay templates update params
func (o *DcimDeviceBayTemplatesUpdateParams) SetData(data *models.WritableDeviceBayTemplate) {
	o.Data = data
}

// WithID adds the id to the dcim device bay templates update params
func (o *DcimDeviceBayTemplatesUpdateParams) WithID(id strfmt.UUID) *DcimDeviceBayTemplatesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim device bay templates update params
func (o *DcimDeviceBayTemplatesUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDeviceBayTemplatesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

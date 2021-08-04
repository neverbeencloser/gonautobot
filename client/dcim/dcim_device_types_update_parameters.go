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

// NewDcimDeviceTypesUpdateParams creates a new DcimDeviceTypesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimDeviceTypesUpdateParams() *DcimDeviceTypesUpdateParams {
	return &DcimDeviceTypesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDeviceTypesUpdateParamsWithTimeout creates a new DcimDeviceTypesUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimDeviceTypesUpdateParamsWithTimeout(timeout time.Duration) *DcimDeviceTypesUpdateParams {
	return &DcimDeviceTypesUpdateParams{
		timeout: timeout,
	}
}

// NewDcimDeviceTypesUpdateParamsWithContext creates a new DcimDeviceTypesUpdateParams object
// with the ability to set a context for a request.
func NewDcimDeviceTypesUpdateParamsWithContext(ctx context.Context) *DcimDeviceTypesUpdateParams {
	return &DcimDeviceTypesUpdateParams{
		Context: ctx,
	}
}

// NewDcimDeviceTypesUpdateParamsWithHTTPClient creates a new DcimDeviceTypesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimDeviceTypesUpdateParamsWithHTTPClient(client *http.Client) *DcimDeviceTypesUpdateParams {
	return &DcimDeviceTypesUpdateParams{
		HTTPClient: client,
	}
}

/* DcimDeviceTypesUpdateParams contains all the parameters to send to the API endpoint
   for the dcim device types update operation.

   Typically these are written to a http.Request.
*/
type DcimDeviceTypesUpdateParams struct {

	// Data.
	Data *models.WritableDeviceType

	/* ID.

	   A UUID string identifying this device type.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim device types update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDeviceTypesUpdateParams) WithDefaults() *DcimDeviceTypesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim device types update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDeviceTypesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim device types update params
func (o *DcimDeviceTypesUpdateParams) WithTimeout(timeout time.Duration) *DcimDeviceTypesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim device types update params
func (o *DcimDeviceTypesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim device types update params
func (o *DcimDeviceTypesUpdateParams) WithContext(ctx context.Context) *DcimDeviceTypesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim device types update params
func (o *DcimDeviceTypesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim device types update params
func (o *DcimDeviceTypesUpdateParams) WithHTTPClient(client *http.Client) *DcimDeviceTypesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim device types update params
func (o *DcimDeviceTypesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim device types update params
func (o *DcimDeviceTypesUpdateParams) WithData(data *models.WritableDeviceType) *DcimDeviceTypesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim device types update params
func (o *DcimDeviceTypesUpdateParams) SetData(data *models.WritableDeviceType) {
	o.Data = data
}

// WithID adds the id to the dcim device types update params
func (o *DcimDeviceTypesUpdateParams) WithID(id strfmt.UUID) *DcimDeviceTypesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim device types update params
func (o *DcimDeviceTypesUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDeviceTypesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

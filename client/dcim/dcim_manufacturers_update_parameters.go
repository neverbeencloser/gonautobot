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

// NewDcimManufacturersUpdateParams creates a new DcimManufacturersUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimManufacturersUpdateParams() *DcimManufacturersUpdateParams {
	return &DcimManufacturersUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimManufacturersUpdateParamsWithTimeout creates a new DcimManufacturersUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimManufacturersUpdateParamsWithTimeout(timeout time.Duration) *DcimManufacturersUpdateParams {
	return &DcimManufacturersUpdateParams{
		timeout: timeout,
	}
}

// NewDcimManufacturersUpdateParamsWithContext creates a new DcimManufacturersUpdateParams object
// with the ability to set a context for a request.
func NewDcimManufacturersUpdateParamsWithContext(ctx context.Context) *DcimManufacturersUpdateParams {
	return &DcimManufacturersUpdateParams{
		Context: ctx,
	}
}

// NewDcimManufacturersUpdateParamsWithHTTPClient creates a new DcimManufacturersUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimManufacturersUpdateParamsWithHTTPClient(client *http.Client) *DcimManufacturersUpdateParams {
	return &DcimManufacturersUpdateParams{
		HTTPClient: client,
	}
}

/* DcimManufacturersUpdateParams contains all the parameters to send to the API endpoint
   for the dcim manufacturers update operation.

   Typically these are written to a http.Request.
*/
type DcimManufacturersUpdateParams struct {

	// Data.
	Data *models.Manufacturer

	/* ID.

	   A UUID string identifying this manufacturer.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim manufacturers update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimManufacturersUpdateParams) WithDefaults() *DcimManufacturersUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim manufacturers update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimManufacturersUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim manufacturers update params
func (o *DcimManufacturersUpdateParams) WithTimeout(timeout time.Duration) *DcimManufacturersUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim manufacturers update params
func (o *DcimManufacturersUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim manufacturers update params
func (o *DcimManufacturersUpdateParams) WithContext(ctx context.Context) *DcimManufacturersUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim manufacturers update params
func (o *DcimManufacturersUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim manufacturers update params
func (o *DcimManufacturersUpdateParams) WithHTTPClient(client *http.Client) *DcimManufacturersUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim manufacturers update params
func (o *DcimManufacturersUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim manufacturers update params
func (o *DcimManufacturersUpdateParams) WithData(data *models.Manufacturer) *DcimManufacturersUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim manufacturers update params
func (o *DcimManufacturersUpdateParams) SetData(data *models.Manufacturer) {
	o.Data = data
}

// WithID adds the id to the dcim manufacturers update params
func (o *DcimManufacturersUpdateParams) WithID(id strfmt.UUID) *DcimManufacturersUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim manufacturers update params
func (o *DcimManufacturersUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimManufacturersUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

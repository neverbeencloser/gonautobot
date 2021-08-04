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

// NewDcimRacksUpdateParams creates a new DcimRacksUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRacksUpdateParams() *DcimRacksUpdateParams {
	return &DcimRacksUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRacksUpdateParamsWithTimeout creates a new DcimRacksUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimRacksUpdateParamsWithTimeout(timeout time.Duration) *DcimRacksUpdateParams {
	return &DcimRacksUpdateParams{
		timeout: timeout,
	}
}

// NewDcimRacksUpdateParamsWithContext creates a new DcimRacksUpdateParams object
// with the ability to set a context for a request.
func NewDcimRacksUpdateParamsWithContext(ctx context.Context) *DcimRacksUpdateParams {
	return &DcimRacksUpdateParams{
		Context: ctx,
	}
}

// NewDcimRacksUpdateParamsWithHTTPClient creates a new DcimRacksUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRacksUpdateParamsWithHTTPClient(client *http.Client) *DcimRacksUpdateParams {
	return &DcimRacksUpdateParams{
		HTTPClient: client,
	}
}

/* DcimRacksUpdateParams contains all the parameters to send to the API endpoint
   for the dcim racks update operation.

   Typically these are written to a http.Request.
*/
type DcimRacksUpdateParams struct {

	// Data.
	Data *models.WritableRack

	/* ID.

	   A UUID string identifying this rack.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim racks update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRacksUpdateParams) WithDefaults() *DcimRacksUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim racks update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRacksUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim racks update params
func (o *DcimRacksUpdateParams) WithTimeout(timeout time.Duration) *DcimRacksUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim racks update params
func (o *DcimRacksUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim racks update params
func (o *DcimRacksUpdateParams) WithContext(ctx context.Context) *DcimRacksUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim racks update params
func (o *DcimRacksUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim racks update params
func (o *DcimRacksUpdateParams) WithHTTPClient(client *http.Client) *DcimRacksUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim racks update params
func (o *DcimRacksUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim racks update params
func (o *DcimRacksUpdateParams) WithData(data *models.WritableRack) *DcimRacksUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim racks update params
func (o *DcimRacksUpdateParams) SetData(data *models.WritableRack) {
	o.Data = data
}

// WithID adds the id to the dcim racks update params
func (o *DcimRacksUpdateParams) WithID(id strfmt.UUID) *DcimRacksUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim racks update params
func (o *DcimRacksUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRacksUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

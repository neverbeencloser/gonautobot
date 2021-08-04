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

// NewDcimPowerOutletsUpdateParams creates a new DcimPowerOutletsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPowerOutletsUpdateParams() *DcimPowerOutletsUpdateParams {
	return &DcimPowerOutletsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerOutletsUpdateParamsWithTimeout creates a new DcimPowerOutletsUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimPowerOutletsUpdateParamsWithTimeout(timeout time.Duration) *DcimPowerOutletsUpdateParams {
	return &DcimPowerOutletsUpdateParams{
		timeout: timeout,
	}
}

// NewDcimPowerOutletsUpdateParamsWithContext creates a new DcimPowerOutletsUpdateParams object
// with the ability to set a context for a request.
func NewDcimPowerOutletsUpdateParamsWithContext(ctx context.Context) *DcimPowerOutletsUpdateParams {
	return &DcimPowerOutletsUpdateParams{
		Context: ctx,
	}
}

// NewDcimPowerOutletsUpdateParamsWithHTTPClient creates a new DcimPowerOutletsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPowerOutletsUpdateParamsWithHTTPClient(client *http.Client) *DcimPowerOutletsUpdateParams {
	return &DcimPowerOutletsUpdateParams{
		HTTPClient: client,
	}
}

/* DcimPowerOutletsUpdateParams contains all the parameters to send to the API endpoint
   for the dcim power outlets update operation.

   Typically these are written to a http.Request.
*/
type DcimPowerOutletsUpdateParams struct {

	// Data.
	Data *models.WritablePowerOutlet

	/* ID.

	   A UUID string identifying this power outlet.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim power outlets update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerOutletsUpdateParams) WithDefaults() *DcimPowerOutletsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim power outlets update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerOutletsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim power outlets update params
func (o *DcimPowerOutletsUpdateParams) WithTimeout(timeout time.Duration) *DcimPowerOutletsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power outlets update params
func (o *DcimPowerOutletsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power outlets update params
func (o *DcimPowerOutletsUpdateParams) WithContext(ctx context.Context) *DcimPowerOutletsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power outlets update params
func (o *DcimPowerOutletsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power outlets update params
func (o *DcimPowerOutletsUpdateParams) WithHTTPClient(client *http.Client) *DcimPowerOutletsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power outlets update params
func (o *DcimPowerOutletsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim power outlets update params
func (o *DcimPowerOutletsUpdateParams) WithData(data *models.WritablePowerOutlet) *DcimPowerOutletsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim power outlets update params
func (o *DcimPowerOutletsUpdateParams) SetData(data *models.WritablePowerOutlet) {
	o.Data = data
}

// WithID adds the id to the dcim power outlets update params
func (o *DcimPowerOutletsUpdateParams) WithID(id strfmt.UUID) *DcimPowerOutletsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim power outlets update params
func (o *DcimPowerOutletsUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerOutletsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

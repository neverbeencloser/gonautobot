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

// NewDcimPowerPortTemplatesPartialUpdateParams creates a new DcimPowerPortTemplatesPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPowerPortTemplatesPartialUpdateParams() *DcimPowerPortTemplatesPartialUpdateParams {
	return &DcimPowerPortTemplatesPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerPortTemplatesPartialUpdateParamsWithTimeout creates a new DcimPowerPortTemplatesPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimPowerPortTemplatesPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimPowerPortTemplatesPartialUpdateParams {
	return &DcimPowerPortTemplatesPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimPowerPortTemplatesPartialUpdateParamsWithContext creates a new DcimPowerPortTemplatesPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimPowerPortTemplatesPartialUpdateParamsWithContext(ctx context.Context) *DcimPowerPortTemplatesPartialUpdateParams {
	return &DcimPowerPortTemplatesPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimPowerPortTemplatesPartialUpdateParamsWithHTTPClient creates a new DcimPowerPortTemplatesPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPowerPortTemplatesPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimPowerPortTemplatesPartialUpdateParams {
	return &DcimPowerPortTemplatesPartialUpdateParams{
		HTTPClient: client,
	}
}

/* DcimPowerPortTemplatesPartialUpdateParams contains all the parameters to send to the API endpoint
   for the dcim power port templates partial update operation.

   Typically these are written to a http.Request.
*/
type DcimPowerPortTemplatesPartialUpdateParams struct {

	// Data.
	Data *models.WritablePowerPortTemplate

	/* ID.

	   A UUID string identifying this power port template.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim power port templates partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerPortTemplatesPartialUpdateParams) WithDefaults() *DcimPowerPortTemplatesPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim power port templates partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerPortTemplatesPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim power port templates partial update params
func (o *DcimPowerPortTemplatesPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimPowerPortTemplatesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power port templates partial update params
func (o *DcimPowerPortTemplatesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power port templates partial update params
func (o *DcimPowerPortTemplatesPartialUpdateParams) WithContext(ctx context.Context) *DcimPowerPortTemplatesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power port templates partial update params
func (o *DcimPowerPortTemplatesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power port templates partial update params
func (o *DcimPowerPortTemplatesPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimPowerPortTemplatesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power port templates partial update params
func (o *DcimPowerPortTemplatesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim power port templates partial update params
func (o *DcimPowerPortTemplatesPartialUpdateParams) WithData(data *models.WritablePowerPortTemplate) *DcimPowerPortTemplatesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim power port templates partial update params
func (o *DcimPowerPortTemplatesPartialUpdateParams) SetData(data *models.WritablePowerPortTemplate) {
	o.Data = data
}

// WithID adds the id to the dcim power port templates partial update params
func (o *DcimPowerPortTemplatesPartialUpdateParams) WithID(id strfmt.UUID) *DcimPowerPortTemplatesPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim power port templates partial update params
func (o *DcimPowerPortTemplatesPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerPortTemplatesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

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

// NewDcimRegionsPartialUpdateParams creates a new DcimRegionsPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRegionsPartialUpdateParams() *DcimRegionsPartialUpdateParams {
	return &DcimRegionsPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRegionsPartialUpdateParamsWithTimeout creates a new DcimRegionsPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimRegionsPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimRegionsPartialUpdateParams {
	return &DcimRegionsPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimRegionsPartialUpdateParamsWithContext creates a new DcimRegionsPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimRegionsPartialUpdateParamsWithContext(ctx context.Context) *DcimRegionsPartialUpdateParams {
	return &DcimRegionsPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimRegionsPartialUpdateParamsWithHTTPClient creates a new DcimRegionsPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRegionsPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimRegionsPartialUpdateParams {
	return &DcimRegionsPartialUpdateParams{
		HTTPClient: client,
	}
}

/* DcimRegionsPartialUpdateParams contains all the parameters to send to the API endpoint
   for the dcim regions partial update operation.

   Typically these are written to a http.Request.
*/
type DcimRegionsPartialUpdateParams struct {

	// Data.
	Data *models.WritableRegion

	/* ID.

	   A UUID string identifying this region.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim regions partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRegionsPartialUpdateParams) WithDefaults() *DcimRegionsPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim regions partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRegionsPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim regions partial update params
func (o *DcimRegionsPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimRegionsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim regions partial update params
func (o *DcimRegionsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim regions partial update params
func (o *DcimRegionsPartialUpdateParams) WithContext(ctx context.Context) *DcimRegionsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim regions partial update params
func (o *DcimRegionsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim regions partial update params
func (o *DcimRegionsPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimRegionsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim regions partial update params
func (o *DcimRegionsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim regions partial update params
func (o *DcimRegionsPartialUpdateParams) WithData(data *models.WritableRegion) *DcimRegionsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim regions partial update params
func (o *DcimRegionsPartialUpdateParams) SetData(data *models.WritableRegion) {
	o.Data = data
}

// WithID adds the id to the dcim regions partial update params
func (o *DcimRegionsPartialUpdateParams) WithID(id strfmt.UUID) *DcimRegionsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim regions partial update params
func (o *DcimRegionsPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRegionsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

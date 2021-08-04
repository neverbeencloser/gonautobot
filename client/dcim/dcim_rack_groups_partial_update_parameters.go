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

// NewDcimRackGroupsPartialUpdateParams creates a new DcimRackGroupsPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRackGroupsPartialUpdateParams() *DcimRackGroupsPartialUpdateParams {
	return &DcimRackGroupsPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRackGroupsPartialUpdateParamsWithTimeout creates a new DcimRackGroupsPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimRackGroupsPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimRackGroupsPartialUpdateParams {
	return &DcimRackGroupsPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimRackGroupsPartialUpdateParamsWithContext creates a new DcimRackGroupsPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimRackGroupsPartialUpdateParamsWithContext(ctx context.Context) *DcimRackGroupsPartialUpdateParams {
	return &DcimRackGroupsPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimRackGroupsPartialUpdateParamsWithHTTPClient creates a new DcimRackGroupsPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRackGroupsPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimRackGroupsPartialUpdateParams {
	return &DcimRackGroupsPartialUpdateParams{
		HTTPClient: client,
	}
}

/* DcimRackGroupsPartialUpdateParams contains all the parameters to send to the API endpoint
   for the dcim rack groups partial update operation.

   Typically these are written to a http.Request.
*/
type DcimRackGroupsPartialUpdateParams struct {

	// Data.
	Data *models.WritableRackGroup

	/* ID.

	   A UUID string identifying this rack group.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim rack groups partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRackGroupsPartialUpdateParams) WithDefaults() *DcimRackGroupsPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim rack groups partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRackGroupsPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim rack groups partial update params
func (o *DcimRackGroupsPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimRackGroupsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rack groups partial update params
func (o *DcimRackGroupsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rack groups partial update params
func (o *DcimRackGroupsPartialUpdateParams) WithContext(ctx context.Context) *DcimRackGroupsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rack groups partial update params
func (o *DcimRackGroupsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rack groups partial update params
func (o *DcimRackGroupsPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimRackGroupsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rack groups partial update params
func (o *DcimRackGroupsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim rack groups partial update params
func (o *DcimRackGroupsPartialUpdateParams) WithData(data *models.WritableRackGroup) *DcimRackGroupsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim rack groups partial update params
func (o *DcimRackGroupsPartialUpdateParams) SetData(data *models.WritableRackGroup) {
	o.Data = data
}

// WithID adds the id to the dcim rack groups partial update params
func (o *DcimRackGroupsPartialUpdateParams) WithID(id strfmt.UUID) *DcimRackGroupsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim rack groups partial update params
func (o *DcimRackGroupsPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRackGroupsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

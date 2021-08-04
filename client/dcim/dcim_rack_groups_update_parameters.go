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

// NewDcimRackGroupsUpdateParams creates a new DcimRackGroupsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRackGroupsUpdateParams() *DcimRackGroupsUpdateParams {
	return &DcimRackGroupsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRackGroupsUpdateParamsWithTimeout creates a new DcimRackGroupsUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimRackGroupsUpdateParamsWithTimeout(timeout time.Duration) *DcimRackGroupsUpdateParams {
	return &DcimRackGroupsUpdateParams{
		timeout: timeout,
	}
}

// NewDcimRackGroupsUpdateParamsWithContext creates a new DcimRackGroupsUpdateParams object
// with the ability to set a context for a request.
func NewDcimRackGroupsUpdateParamsWithContext(ctx context.Context) *DcimRackGroupsUpdateParams {
	return &DcimRackGroupsUpdateParams{
		Context: ctx,
	}
}

// NewDcimRackGroupsUpdateParamsWithHTTPClient creates a new DcimRackGroupsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRackGroupsUpdateParamsWithHTTPClient(client *http.Client) *DcimRackGroupsUpdateParams {
	return &DcimRackGroupsUpdateParams{
		HTTPClient: client,
	}
}

/* DcimRackGroupsUpdateParams contains all the parameters to send to the API endpoint
   for the dcim rack groups update operation.

   Typically these are written to a http.Request.
*/
type DcimRackGroupsUpdateParams struct {

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

// WithDefaults hydrates default values in the dcim rack groups update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRackGroupsUpdateParams) WithDefaults() *DcimRackGroupsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim rack groups update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRackGroupsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim rack groups update params
func (o *DcimRackGroupsUpdateParams) WithTimeout(timeout time.Duration) *DcimRackGroupsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rack groups update params
func (o *DcimRackGroupsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rack groups update params
func (o *DcimRackGroupsUpdateParams) WithContext(ctx context.Context) *DcimRackGroupsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rack groups update params
func (o *DcimRackGroupsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rack groups update params
func (o *DcimRackGroupsUpdateParams) WithHTTPClient(client *http.Client) *DcimRackGroupsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rack groups update params
func (o *DcimRackGroupsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim rack groups update params
func (o *DcimRackGroupsUpdateParams) WithData(data *models.WritableRackGroup) *DcimRackGroupsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim rack groups update params
func (o *DcimRackGroupsUpdateParams) SetData(data *models.WritableRackGroup) {
	o.Data = data
}

// WithID adds the id to the dcim rack groups update params
func (o *DcimRackGroupsUpdateParams) WithID(id strfmt.UUID) *DcimRackGroupsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim rack groups update params
func (o *DcimRackGroupsUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRackGroupsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

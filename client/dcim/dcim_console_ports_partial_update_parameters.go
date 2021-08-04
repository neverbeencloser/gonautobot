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

// NewDcimConsolePortsPartialUpdateParams creates a new DcimConsolePortsPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimConsolePortsPartialUpdateParams() *DcimConsolePortsPartialUpdateParams {
	return &DcimConsolePortsPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimConsolePortsPartialUpdateParamsWithTimeout creates a new DcimConsolePortsPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimConsolePortsPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimConsolePortsPartialUpdateParams {
	return &DcimConsolePortsPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimConsolePortsPartialUpdateParamsWithContext creates a new DcimConsolePortsPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimConsolePortsPartialUpdateParamsWithContext(ctx context.Context) *DcimConsolePortsPartialUpdateParams {
	return &DcimConsolePortsPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimConsolePortsPartialUpdateParamsWithHTTPClient creates a new DcimConsolePortsPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimConsolePortsPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimConsolePortsPartialUpdateParams {
	return &DcimConsolePortsPartialUpdateParams{
		HTTPClient: client,
	}
}

/* DcimConsolePortsPartialUpdateParams contains all the parameters to send to the API endpoint
   for the dcim console ports partial update operation.

   Typically these are written to a http.Request.
*/
type DcimConsolePortsPartialUpdateParams struct {

	// Data.
	Data *models.WritableConsolePort

	/* ID.

	   A UUID string identifying this console port.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim console ports partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsolePortsPartialUpdateParams) WithDefaults() *DcimConsolePortsPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim console ports partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsolePortsPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim console ports partial update params
func (o *DcimConsolePortsPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimConsolePortsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim console ports partial update params
func (o *DcimConsolePortsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim console ports partial update params
func (o *DcimConsolePortsPartialUpdateParams) WithContext(ctx context.Context) *DcimConsolePortsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim console ports partial update params
func (o *DcimConsolePortsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim console ports partial update params
func (o *DcimConsolePortsPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimConsolePortsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim console ports partial update params
func (o *DcimConsolePortsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim console ports partial update params
func (o *DcimConsolePortsPartialUpdateParams) WithData(data *models.WritableConsolePort) *DcimConsolePortsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim console ports partial update params
func (o *DcimConsolePortsPartialUpdateParams) SetData(data *models.WritableConsolePort) {
	o.Data = data
}

// WithID adds the id to the dcim console ports partial update params
func (o *DcimConsolePortsPartialUpdateParams) WithID(id strfmt.UUID) *DcimConsolePortsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim console ports partial update params
func (o *DcimConsolePortsPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimConsolePortsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

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

// NewDcimConsolePortTemplatesUpdateParams creates a new DcimConsolePortTemplatesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimConsolePortTemplatesUpdateParams() *DcimConsolePortTemplatesUpdateParams {
	return &DcimConsolePortTemplatesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimConsolePortTemplatesUpdateParamsWithTimeout creates a new DcimConsolePortTemplatesUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimConsolePortTemplatesUpdateParamsWithTimeout(timeout time.Duration) *DcimConsolePortTemplatesUpdateParams {
	return &DcimConsolePortTemplatesUpdateParams{
		timeout: timeout,
	}
}

// NewDcimConsolePortTemplatesUpdateParamsWithContext creates a new DcimConsolePortTemplatesUpdateParams object
// with the ability to set a context for a request.
func NewDcimConsolePortTemplatesUpdateParamsWithContext(ctx context.Context) *DcimConsolePortTemplatesUpdateParams {
	return &DcimConsolePortTemplatesUpdateParams{
		Context: ctx,
	}
}

// NewDcimConsolePortTemplatesUpdateParamsWithHTTPClient creates a new DcimConsolePortTemplatesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimConsolePortTemplatesUpdateParamsWithHTTPClient(client *http.Client) *DcimConsolePortTemplatesUpdateParams {
	return &DcimConsolePortTemplatesUpdateParams{
		HTTPClient: client,
	}
}

/* DcimConsolePortTemplatesUpdateParams contains all the parameters to send to the API endpoint
   for the dcim console port templates update operation.

   Typically these are written to a http.Request.
*/
type DcimConsolePortTemplatesUpdateParams struct {

	// Data.
	Data *models.WritableConsolePortTemplate

	/* ID.

	   A UUID string identifying this console port template.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim console port templates update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsolePortTemplatesUpdateParams) WithDefaults() *DcimConsolePortTemplatesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim console port templates update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsolePortTemplatesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim console port templates update params
func (o *DcimConsolePortTemplatesUpdateParams) WithTimeout(timeout time.Duration) *DcimConsolePortTemplatesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim console port templates update params
func (o *DcimConsolePortTemplatesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim console port templates update params
func (o *DcimConsolePortTemplatesUpdateParams) WithContext(ctx context.Context) *DcimConsolePortTemplatesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim console port templates update params
func (o *DcimConsolePortTemplatesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim console port templates update params
func (o *DcimConsolePortTemplatesUpdateParams) WithHTTPClient(client *http.Client) *DcimConsolePortTemplatesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim console port templates update params
func (o *DcimConsolePortTemplatesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim console port templates update params
func (o *DcimConsolePortTemplatesUpdateParams) WithData(data *models.WritableConsolePortTemplate) *DcimConsolePortTemplatesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim console port templates update params
func (o *DcimConsolePortTemplatesUpdateParams) SetData(data *models.WritableConsolePortTemplate) {
	o.Data = data
}

// WithID adds the id to the dcim console port templates update params
func (o *DcimConsolePortTemplatesUpdateParams) WithID(id strfmt.UUID) *DcimConsolePortTemplatesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim console port templates update params
func (o *DcimConsolePortTemplatesUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimConsolePortTemplatesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

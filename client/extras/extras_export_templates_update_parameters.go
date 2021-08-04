package extras

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

// NewExtrasExportTemplatesUpdateParams creates a new ExtrasExportTemplatesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasExportTemplatesUpdateParams() *ExtrasExportTemplatesUpdateParams {
	return &ExtrasExportTemplatesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasExportTemplatesUpdateParamsWithTimeout creates a new ExtrasExportTemplatesUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasExportTemplatesUpdateParamsWithTimeout(timeout time.Duration) *ExtrasExportTemplatesUpdateParams {
	return &ExtrasExportTemplatesUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasExportTemplatesUpdateParamsWithContext creates a new ExtrasExportTemplatesUpdateParams object
// with the ability to set a context for a request.
func NewExtrasExportTemplatesUpdateParamsWithContext(ctx context.Context) *ExtrasExportTemplatesUpdateParams {
	return &ExtrasExportTemplatesUpdateParams{
		Context: ctx,
	}
}

// NewExtrasExportTemplatesUpdateParamsWithHTTPClient creates a new ExtrasExportTemplatesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasExportTemplatesUpdateParamsWithHTTPClient(client *http.Client) *ExtrasExportTemplatesUpdateParams {
	return &ExtrasExportTemplatesUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasExportTemplatesUpdateParams contains all the parameters to send to the API endpoint
   for the extras export templates update operation.

   Typically these are written to a http.Request.
*/
type ExtrasExportTemplatesUpdateParams struct {

	// Data.
	Data *models.ExportTemplate

	/* ID.

	   A UUID string identifying this export template.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras export templates update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasExportTemplatesUpdateParams) WithDefaults() *ExtrasExportTemplatesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras export templates update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasExportTemplatesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras export templates update params
func (o *ExtrasExportTemplatesUpdateParams) WithTimeout(timeout time.Duration) *ExtrasExportTemplatesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras export templates update params
func (o *ExtrasExportTemplatesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras export templates update params
func (o *ExtrasExportTemplatesUpdateParams) WithContext(ctx context.Context) *ExtrasExportTemplatesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras export templates update params
func (o *ExtrasExportTemplatesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras export templates update params
func (o *ExtrasExportTemplatesUpdateParams) WithHTTPClient(client *http.Client) *ExtrasExportTemplatesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras export templates update params
func (o *ExtrasExportTemplatesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras export templates update params
func (o *ExtrasExportTemplatesUpdateParams) WithData(data *models.ExportTemplate) *ExtrasExportTemplatesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras export templates update params
func (o *ExtrasExportTemplatesUpdateParams) SetData(data *models.ExportTemplate) {
	o.Data = data
}

// WithID adds the id to the extras export templates update params
func (o *ExtrasExportTemplatesUpdateParams) WithID(id strfmt.UUID) *ExtrasExportTemplatesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras export templates update params
func (o *ExtrasExportTemplatesUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasExportTemplatesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

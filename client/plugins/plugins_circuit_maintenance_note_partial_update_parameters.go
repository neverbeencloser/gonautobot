package plugins

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

// NewPluginsCircuitMaintenanceNotePartialUpdateParams creates a new PluginsCircuitMaintenanceNotePartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsCircuitMaintenanceNotePartialUpdateParams() *PluginsCircuitMaintenanceNotePartialUpdateParams {
	return &PluginsCircuitMaintenanceNotePartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsCircuitMaintenanceNotePartialUpdateParamsWithTimeout creates a new PluginsCircuitMaintenanceNotePartialUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsCircuitMaintenanceNotePartialUpdateParamsWithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceNotePartialUpdateParams {
	return &PluginsCircuitMaintenanceNotePartialUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsCircuitMaintenanceNotePartialUpdateParamsWithContext creates a new PluginsCircuitMaintenanceNotePartialUpdateParams object
// with the ability to set a context for a request.
func NewPluginsCircuitMaintenanceNotePartialUpdateParamsWithContext(ctx context.Context) *PluginsCircuitMaintenanceNotePartialUpdateParams {
	return &PluginsCircuitMaintenanceNotePartialUpdateParams{
		Context: ctx,
	}
}

// NewPluginsCircuitMaintenanceNotePartialUpdateParamsWithHTTPClient creates a new PluginsCircuitMaintenanceNotePartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsCircuitMaintenanceNotePartialUpdateParamsWithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceNotePartialUpdateParams {
	return &PluginsCircuitMaintenanceNotePartialUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsCircuitMaintenanceNotePartialUpdateParams contains all the parameters to send to the API endpoint
   for the plugins circuit maintenance note partial update operation.

   Typically these are written to a http.Request.
*/
type PluginsCircuitMaintenanceNotePartialUpdateParams struct {

	// Data.
	Data *models.Note

	/* ID.

	   A UUID string identifying this note.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins circuit maintenance note partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceNotePartialUpdateParams) WithDefaults() *PluginsCircuitMaintenanceNotePartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins circuit maintenance note partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceNotePartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins circuit maintenance note partial update params
func (o *PluginsCircuitMaintenanceNotePartialUpdateParams) WithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceNotePartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins circuit maintenance note partial update params
func (o *PluginsCircuitMaintenanceNotePartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins circuit maintenance note partial update params
func (o *PluginsCircuitMaintenanceNotePartialUpdateParams) WithContext(ctx context.Context) *PluginsCircuitMaintenanceNotePartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins circuit maintenance note partial update params
func (o *PluginsCircuitMaintenanceNotePartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins circuit maintenance note partial update params
func (o *PluginsCircuitMaintenanceNotePartialUpdateParams) WithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceNotePartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins circuit maintenance note partial update params
func (o *PluginsCircuitMaintenanceNotePartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins circuit maintenance note partial update params
func (o *PluginsCircuitMaintenanceNotePartialUpdateParams) WithData(data *models.Note) *PluginsCircuitMaintenanceNotePartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins circuit maintenance note partial update params
func (o *PluginsCircuitMaintenanceNotePartialUpdateParams) SetData(data *models.Note) {
	o.Data = data
}

// WithID adds the id to the plugins circuit maintenance note partial update params
func (o *PluginsCircuitMaintenanceNotePartialUpdateParams) WithID(id strfmt.UUID) *PluginsCircuitMaintenanceNotePartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins circuit maintenance note partial update params
func (o *PluginsCircuitMaintenanceNotePartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsCircuitMaintenanceNotePartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

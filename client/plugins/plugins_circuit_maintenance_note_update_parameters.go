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

// NewPluginsCircuitMaintenanceNoteUpdateParams creates a new PluginsCircuitMaintenanceNoteUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsCircuitMaintenanceNoteUpdateParams() *PluginsCircuitMaintenanceNoteUpdateParams {
	return &PluginsCircuitMaintenanceNoteUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsCircuitMaintenanceNoteUpdateParamsWithTimeout creates a new PluginsCircuitMaintenanceNoteUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsCircuitMaintenanceNoteUpdateParamsWithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceNoteUpdateParams {
	return &PluginsCircuitMaintenanceNoteUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsCircuitMaintenanceNoteUpdateParamsWithContext creates a new PluginsCircuitMaintenanceNoteUpdateParams object
// with the ability to set a context for a request.
func NewPluginsCircuitMaintenanceNoteUpdateParamsWithContext(ctx context.Context) *PluginsCircuitMaintenanceNoteUpdateParams {
	return &PluginsCircuitMaintenanceNoteUpdateParams{
		Context: ctx,
	}
}

// NewPluginsCircuitMaintenanceNoteUpdateParamsWithHTTPClient creates a new PluginsCircuitMaintenanceNoteUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsCircuitMaintenanceNoteUpdateParamsWithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceNoteUpdateParams {
	return &PluginsCircuitMaintenanceNoteUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsCircuitMaintenanceNoteUpdateParams contains all the parameters to send to the API endpoint
   for the plugins circuit maintenance note update operation.

   Typically these are written to a http.Request.
*/
type PluginsCircuitMaintenanceNoteUpdateParams struct {

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

// WithDefaults hydrates default values in the plugins circuit maintenance note update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceNoteUpdateParams) WithDefaults() *PluginsCircuitMaintenanceNoteUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins circuit maintenance note update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceNoteUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins circuit maintenance note update params
func (o *PluginsCircuitMaintenanceNoteUpdateParams) WithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceNoteUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins circuit maintenance note update params
func (o *PluginsCircuitMaintenanceNoteUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins circuit maintenance note update params
func (o *PluginsCircuitMaintenanceNoteUpdateParams) WithContext(ctx context.Context) *PluginsCircuitMaintenanceNoteUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins circuit maintenance note update params
func (o *PluginsCircuitMaintenanceNoteUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins circuit maintenance note update params
func (o *PluginsCircuitMaintenanceNoteUpdateParams) WithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceNoteUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins circuit maintenance note update params
func (o *PluginsCircuitMaintenanceNoteUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins circuit maintenance note update params
func (o *PluginsCircuitMaintenanceNoteUpdateParams) WithData(data *models.Note) *PluginsCircuitMaintenanceNoteUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins circuit maintenance note update params
func (o *PluginsCircuitMaintenanceNoteUpdateParams) SetData(data *models.Note) {
	o.Data = data
}

// WithID adds the id to the plugins circuit maintenance note update params
func (o *PluginsCircuitMaintenanceNoteUpdateParams) WithID(id strfmt.UUID) *PluginsCircuitMaintenanceNoteUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins circuit maintenance note update params
func (o *PluginsCircuitMaintenanceNoteUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsCircuitMaintenanceNoteUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

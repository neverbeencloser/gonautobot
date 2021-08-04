package plugins

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewPluginsCircuitMaintenanceNoteDeleteParams creates a new PluginsCircuitMaintenanceNoteDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsCircuitMaintenanceNoteDeleteParams() *PluginsCircuitMaintenanceNoteDeleteParams {
	return &PluginsCircuitMaintenanceNoteDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsCircuitMaintenanceNoteDeleteParamsWithTimeout creates a new PluginsCircuitMaintenanceNoteDeleteParams object
// with the ability to set a timeout on a request.
func NewPluginsCircuitMaintenanceNoteDeleteParamsWithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceNoteDeleteParams {
	return &PluginsCircuitMaintenanceNoteDeleteParams{
		timeout: timeout,
	}
}

// NewPluginsCircuitMaintenanceNoteDeleteParamsWithContext creates a new PluginsCircuitMaintenanceNoteDeleteParams object
// with the ability to set a context for a request.
func NewPluginsCircuitMaintenanceNoteDeleteParamsWithContext(ctx context.Context) *PluginsCircuitMaintenanceNoteDeleteParams {
	return &PluginsCircuitMaintenanceNoteDeleteParams{
		Context: ctx,
	}
}

// NewPluginsCircuitMaintenanceNoteDeleteParamsWithHTTPClient creates a new PluginsCircuitMaintenanceNoteDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsCircuitMaintenanceNoteDeleteParamsWithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceNoteDeleteParams {
	return &PluginsCircuitMaintenanceNoteDeleteParams{
		HTTPClient: client,
	}
}

/* PluginsCircuitMaintenanceNoteDeleteParams contains all the parameters to send to the API endpoint
   for the plugins circuit maintenance note delete operation.

   Typically these are written to a http.Request.
*/
type PluginsCircuitMaintenanceNoteDeleteParams struct {

	/* ID.

	   A UUID string identifying this note.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins circuit maintenance note delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceNoteDeleteParams) WithDefaults() *PluginsCircuitMaintenanceNoteDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins circuit maintenance note delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceNoteDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins circuit maintenance note delete params
func (o *PluginsCircuitMaintenanceNoteDeleteParams) WithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceNoteDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins circuit maintenance note delete params
func (o *PluginsCircuitMaintenanceNoteDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins circuit maintenance note delete params
func (o *PluginsCircuitMaintenanceNoteDeleteParams) WithContext(ctx context.Context) *PluginsCircuitMaintenanceNoteDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins circuit maintenance note delete params
func (o *PluginsCircuitMaintenanceNoteDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins circuit maintenance note delete params
func (o *PluginsCircuitMaintenanceNoteDeleteParams) WithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceNoteDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins circuit maintenance note delete params
func (o *PluginsCircuitMaintenanceNoteDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the plugins circuit maintenance note delete params
func (o *PluginsCircuitMaintenanceNoteDeleteParams) WithID(id strfmt.UUID) *PluginsCircuitMaintenanceNoteDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins circuit maintenance note delete params
func (o *PluginsCircuitMaintenanceNoteDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsCircuitMaintenanceNoteDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

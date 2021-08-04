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

// NewPluginsCircuitMaintenanceNoteCreateParams creates a new PluginsCircuitMaintenanceNoteCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsCircuitMaintenanceNoteCreateParams() *PluginsCircuitMaintenanceNoteCreateParams {
	return &PluginsCircuitMaintenanceNoteCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsCircuitMaintenanceNoteCreateParamsWithTimeout creates a new PluginsCircuitMaintenanceNoteCreateParams object
// with the ability to set a timeout on a request.
func NewPluginsCircuitMaintenanceNoteCreateParamsWithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceNoteCreateParams {
	return &PluginsCircuitMaintenanceNoteCreateParams{
		timeout: timeout,
	}
}

// NewPluginsCircuitMaintenanceNoteCreateParamsWithContext creates a new PluginsCircuitMaintenanceNoteCreateParams object
// with the ability to set a context for a request.
func NewPluginsCircuitMaintenanceNoteCreateParamsWithContext(ctx context.Context) *PluginsCircuitMaintenanceNoteCreateParams {
	return &PluginsCircuitMaintenanceNoteCreateParams{
		Context: ctx,
	}
}

// NewPluginsCircuitMaintenanceNoteCreateParamsWithHTTPClient creates a new PluginsCircuitMaintenanceNoteCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsCircuitMaintenanceNoteCreateParamsWithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceNoteCreateParams {
	return &PluginsCircuitMaintenanceNoteCreateParams{
		HTTPClient: client,
	}
}

/* PluginsCircuitMaintenanceNoteCreateParams contains all the parameters to send to the API endpoint
   for the plugins circuit maintenance note create operation.

   Typically these are written to a http.Request.
*/
type PluginsCircuitMaintenanceNoteCreateParams struct {

	// Data.
	Data *models.Note

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins circuit maintenance note create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceNoteCreateParams) WithDefaults() *PluginsCircuitMaintenanceNoteCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins circuit maintenance note create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceNoteCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins circuit maintenance note create params
func (o *PluginsCircuitMaintenanceNoteCreateParams) WithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceNoteCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins circuit maintenance note create params
func (o *PluginsCircuitMaintenanceNoteCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins circuit maintenance note create params
func (o *PluginsCircuitMaintenanceNoteCreateParams) WithContext(ctx context.Context) *PluginsCircuitMaintenanceNoteCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins circuit maintenance note create params
func (o *PluginsCircuitMaintenanceNoteCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins circuit maintenance note create params
func (o *PluginsCircuitMaintenanceNoteCreateParams) WithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceNoteCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins circuit maintenance note create params
func (o *PluginsCircuitMaintenanceNoteCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins circuit maintenance note create params
func (o *PluginsCircuitMaintenanceNoteCreateParams) WithData(data *models.Note) *PluginsCircuitMaintenanceNoteCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins circuit maintenance note create params
func (o *PluginsCircuitMaintenanceNoteCreateParams) SetData(data *models.Note) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsCircuitMaintenanceNoteCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

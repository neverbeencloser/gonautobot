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

// NewPluginsCircuitMaintenanceMaintenanceUpdateParams creates a new PluginsCircuitMaintenanceMaintenanceUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsCircuitMaintenanceMaintenanceUpdateParams() *PluginsCircuitMaintenanceMaintenanceUpdateParams {
	return &PluginsCircuitMaintenanceMaintenanceUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsCircuitMaintenanceMaintenanceUpdateParamsWithTimeout creates a new PluginsCircuitMaintenanceMaintenanceUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsCircuitMaintenanceMaintenanceUpdateParamsWithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceMaintenanceUpdateParams {
	return &PluginsCircuitMaintenanceMaintenanceUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsCircuitMaintenanceMaintenanceUpdateParamsWithContext creates a new PluginsCircuitMaintenanceMaintenanceUpdateParams object
// with the ability to set a context for a request.
func NewPluginsCircuitMaintenanceMaintenanceUpdateParamsWithContext(ctx context.Context) *PluginsCircuitMaintenanceMaintenanceUpdateParams {
	return &PluginsCircuitMaintenanceMaintenanceUpdateParams{
		Context: ctx,
	}
}

// NewPluginsCircuitMaintenanceMaintenanceUpdateParamsWithHTTPClient creates a new PluginsCircuitMaintenanceMaintenanceUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsCircuitMaintenanceMaintenanceUpdateParamsWithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceMaintenanceUpdateParams {
	return &PluginsCircuitMaintenanceMaintenanceUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsCircuitMaintenanceMaintenanceUpdateParams contains all the parameters to send to the API endpoint
   for the plugins circuit maintenance maintenance update operation.

   Typically these are written to a http.Request.
*/
type PluginsCircuitMaintenanceMaintenanceUpdateParams struct {

	// Data.
	Data *models.CircuitMaintenance

	/* ID.

	   A UUID string identifying this circuit maintenance.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins circuit maintenance maintenance update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceMaintenanceUpdateParams) WithDefaults() *PluginsCircuitMaintenanceMaintenanceUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins circuit maintenance maintenance update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceMaintenanceUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins circuit maintenance maintenance update params
func (o *PluginsCircuitMaintenanceMaintenanceUpdateParams) WithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceMaintenanceUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins circuit maintenance maintenance update params
func (o *PluginsCircuitMaintenanceMaintenanceUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins circuit maintenance maintenance update params
func (o *PluginsCircuitMaintenanceMaintenanceUpdateParams) WithContext(ctx context.Context) *PluginsCircuitMaintenanceMaintenanceUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins circuit maintenance maintenance update params
func (o *PluginsCircuitMaintenanceMaintenanceUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins circuit maintenance maintenance update params
func (o *PluginsCircuitMaintenanceMaintenanceUpdateParams) WithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceMaintenanceUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins circuit maintenance maintenance update params
func (o *PluginsCircuitMaintenanceMaintenanceUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins circuit maintenance maintenance update params
func (o *PluginsCircuitMaintenanceMaintenanceUpdateParams) WithData(data *models.CircuitMaintenance) *PluginsCircuitMaintenanceMaintenanceUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins circuit maintenance maintenance update params
func (o *PluginsCircuitMaintenanceMaintenanceUpdateParams) SetData(data *models.CircuitMaintenance) {
	o.Data = data
}

// WithID adds the id to the plugins circuit maintenance maintenance update params
func (o *PluginsCircuitMaintenanceMaintenanceUpdateParams) WithID(id strfmt.UUID) *PluginsCircuitMaintenanceMaintenanceUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins circuit maintenance maintenance update params
func (o *PluginsCircuitMaintenanceMaintenanceUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsCircuitMaintenanceMaintenanceUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

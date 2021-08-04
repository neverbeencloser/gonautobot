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

// NewPluginsCircuitMaintenanceMaintenancePartialUpdateParams creates a new PluginsCircuitMaintenanceMaintenancePartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsCircuitMaintenanceMaintenancePartialUpdateParams() *PluginsCircuitMaintenanceMaintenancePartialUpdateParams {
	return &PluginsCircuitMaintenanceMaintenancePartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsCircuitMaintenanceMaintenancePartialUpdateParamsWithTimeout creates a new PluginsCircuitMaintenanceMaintenancePartialUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsCircuitMaintenanceMaintenancePartialUpdateParamsWithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceMaintenancePartialUpdateParams {
	return &PluginsCircuitMaintenanceMaintenancePartialUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsCircuitMaintenanceMaintenancePartialUpdateParamsWithContext creates a new PluginsCircuitMaintenanceMaintenancePartialUpdateParams object
// with the ability to set a context for a request.
func NewPluginsCircuitMaintenanceMaintenancePartialUpdateParamsWithContext(ctx context.Context) *PluginsCircuitMaintenanceMaintenancePartialUpdateParams {
	return &PluginsCircuitMaintenanceMaintenancePartialUpdateParams{
		Context: ctx,
	}
}

// NewPluginsCircuitMaintenanceMaintenancePartialUpdateParamsWithHTTPClient creates a new PluginsCircuitMaintenanceMaintenancePartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsCircuitMaintenanceMaintenancePartialUpdateParamsWithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceMaintenancePartialUpdateParams {
	return &PluginsCircuitMaintenanceMaintenancePartialUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsCircuitMaintenanceMaintenancePartialUpdateParams contains all the parameters to send to the API endpoint
   for the plugins circuit maintenance maintenance partial update operation.

   Typically these are written to a http.Request.
*/
type PluginsCircuitMaintenanceMaintenancePartialUpdateParams struct {

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

// WithDefaults hydrates default values in the plugins circuit maintenance maintenance partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceMaintenancePartialUpdateParams) WithDefaults() *PluginsCircuitMaintenanceMaintenancePartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins circuit maintenance maintenance partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceMaintenancePartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins circuit maintenance maintenance partial update params
func (o *PluginsCircuitMaintenanceMaintenancePartialUpdateParams) WithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceMaintenancePartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins circuit maintenance maintenance partial update params
func (o *PluginsCircuitMaintenanceMaintenancePartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins circuit maintenance maintenance partial update params
func (o *PluginsCircuitMaintenanceMaintenancePartialUpdateParams) WithContext(ctx context.Context) *PluginsCircuitMaintenanceMaintenancePartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins circuit maintenance maintenance partial update params
func (o *PluginsCircuitMaintenanceMaintenancePartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins circuit maintenance maintenance partial update params
func (o *PluginsCircuitMaintenanceMaintenancePartialUpdateParams) WithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceMaintenancePartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins circuit maintenance maintenance partial update params
func (o *PluginsCircuitMaintenanceMaintenancePartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins circuit maintenance maintenance partial update params
func (o *PluginsCircuitMaintenanceMaintenancePartialUpdateParams) WithData(data *models.CircuitMaintenance) *PluginsCircuitMaintenanceMaintenancePartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins circuit maintenance maintenance partial update params
func (o *PluginsCircuitMaintenanceMaintenancePartialUpdateParams) SetData(data *models.CircuitMaintenance) {
	o.Data = data
}

// WithID adds the id to the plugins circuit maintenance maintenance partial update params
func (o *PluginsCircuitMaintenanceMaintenancePartialUpdateParams) WithID(id strfmt.UUID) *PluginsCircuitMaintenanceMaintenancePartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins circuit maintenance maintenance partial update params
func (o *PluginsCircuitMaintenanceMaintenancePartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsCircuitMaintenanceMaintenancePartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

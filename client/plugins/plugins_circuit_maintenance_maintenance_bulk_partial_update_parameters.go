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

// NewPluginsCircuitMaintenanceMaintenanceBulkPartialUpdateParams creates a new PluginsCircuitMaintenanceMaintenanceBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsCircuitMaintenanceMaintenanceBulkPartialUpdateParams() *PluginsCircuitMaintenanceMaintenanceBulkPartialUpdateParams {
	return &PluginsCircuitMaintenanceMaintenanceBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsCircuitMaintenanceMaintenanceBulkPartialUpdateParamsWithTimeout creates a new PluginsCircuitMaintenanceMaintenanceBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsCircuitMaintenanceMaintenanceBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceMaintenanceBulkPartialUpdateParams {
	return &PluginsCircuitMaintenanceMaintenanceBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsCircuitMaintenanceMaintenanceBulkPartialUpdateParamsWithContext creates a new PluginsCircuitMaintenanceMaintenanceBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewPluginsCircuitMaintenanceMaintenanceBulkPartialUpdateParamsWithContext(ctx context.Context) *PluginsCircuitMaintenanceMaintenanceBulkPartialUpdateParams {
	return &PluginsCircuitMaintenanceMaintenanceBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewPluginsCircuitMaintenanceMaintenanceBulkPartialUpdateParamsWithHTTPClient creates a new PluginsCircuitMaintenanceMaintenanceBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsCircuitMaintenanceMaintenanceBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceMaintenanceBulkPartialUpdateParams {
	return &PluginsCircuitMaintenanceMaintenanceBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsCircuitMaintenanceMaintenanceBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the plugins circuit maintenance maintenance bulk partial update operation.

   Typically these are written to a http.Request.
*/
type PluginsCircuitMaintenanceMaintenanceBulkPartialUpdateParams struct {

	// Data.
	Data *models.CircuitMaintenance

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins circuit maintenance maintenance bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceMaintenanceBulkPartialUpdateParams) WithDefaults() *PluginsCircuitMaintenanceMaintenanceBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins circuit maintenance maintenance bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceMaintenanceBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins circuit maintenance maintenance bulk partial update params
func (o *PluginsCircuitMaintenanceMaintenanceBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceMaintenanceBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins circuit maintenance maintenance bulk partial update params
func (o *PluginsCircuitMaintenanceMaintenanceBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins circuit maintenance maintenance bulk partial update params
func (o *PluginsCircuitMaintenanceMaintenanceBulkPartialUpdateParams) WithContext(ctx context.Context) *PluginsCircuitMaintenanceMaintenanceBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins circuit maintenance maintenance bulk partial update params
func (o *PluginsCircuitMaintenanceMaintenanceBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins circuit maintenance maintenance bulk partial update params
func (o *PluginsCircuitMaintenanceMaintenanceBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceMaintenanceBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins circuit maintenance maintenance bulk partial update params
func (o *PluginsCircuitMaintenanceMaintenanceBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins circuit maintenance maintenance bulk partial update params
func (o *PluginsCircuitMaintenanceMaintenanceBulkPartialUpdateParams) WithData(data *models.CircuitMaintenance) *PluginsCircuitMaintenanceMaintenanceBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins circuit maintenance maintenance bulk partial update params
func (o *PluginsCircuitMaintenanceMaintenanceBulkPartialUpdateParams) SetData(data *models.CircuitMaintenance) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsCircuitMaintenanceMaintenanceBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

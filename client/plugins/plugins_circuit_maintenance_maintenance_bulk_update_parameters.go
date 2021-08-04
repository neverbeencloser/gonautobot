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

// NewPluginsCircuitMaintenanceMaintenanceBulkUpdateParams creates a new PluginsCircuitMaintenanceMaintenanceBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsCircuitMaintenanceMaintenanceBulkUpdateParams() *PluginsCircuitMaintenanceMaintenanceBulkUpdateParams {
	return &PluginsCircuitMaintenanceMaintenanceBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsCircuitMaintenanceMaintenanceBulkUpdateParamsWithTimeout creates a new PluginsCircuitMaintenanceMaintenanceBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsCircuitMaintenanceMaintenanceBulkUpdateParamsWithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceMaintenanceBulkUpdateParams {
	return &PluginsCircuitMaintenanceMaintenanceBulkUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsCircuitMaintenanceMaintenanceBulkUpdateParamsWithContext creates a new PluginsCircuitMaintenanceMaintenanceBulkUpdateParams object
// with the ability to set a context for a request.
func NewPluginsCircuitMaintenanceMaintenanceBulkUpdateParamsWithContext(ctx context.Context) *PluginsCircuitMaintenanceMaintenanceBulkUpdateParams {
	return &PluginsCircuitMaintenanceMaintenanceBulkUpdateParams{
		Context: ctx,
	}
}

// NewPluginsCircuitMaintenanceMaintenanceBulkUpdateParamsWithHTTPClient creates a new PluginsCircuitMaintenanceMaintenanceBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsCircuitMaintenanceMaintenanceBulkUpdateParamsWithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceMaintenanceBulkUpdateParams {
	return &PluginsCircuitMaintenanceMaintenanceBulkUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsCircuitMaintenanceMaintenanceBulkUpdateParams contains all the parameters to send to the API endpoint
   for the plugins circuit maintenance maintenance bulk update operation.

   Typically these are written to a http.Request.
*/
type PluginsCircuitMaintenanceMaintenanceBulkUpdateParams struct {

	// Data.
	Data *models.CircuitMaintenance

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins circuit maintenance maintenance bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceMaintenanceBulkUpdateParams) WithDefaults() *PluginsCircuitMaintenanceMaintenanceBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins circuit maintenance maintenance bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceMaintenanceBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins circuit maintenance maintenance bulk update params
func (o *PluginsCircuitMaintenanceMaintenanceBulkUpdateParams) WithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceMaintenanceBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins circuit maintenance maintenance bulk update params
func (o *PluginsCircuitMaintenanceMaintenanceBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins circuit maintenance maintenance bulk update params
func (o *PluginsCircuitMaintenanceMaintenanceBulkUpdateParams) WithContext(ctx context.Context) *PluginsCircuitMaintenanceMaintenanceBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins circuit maintenance maintenance bulk update params
func (o *PluginsCircuitMaintenanceMaintenanceBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins circuit maintenance maintenance bulk update params
func (o *PluginsCircuitMaintenanceMaintenanceBulkUpdateParams) WithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceMaintenanceBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins circuit maintenance maintenance bulk update params
func (o *PluginsCircuitMaintenanceMaintenanceBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins circuit maintenance maintenance bulk update params
func (o *PluginsCircuitMaintenanceMaintenanceBulkUpdateParams) WithData(data *models.CircuitMaintenance) *PluginsCircuitMaintenanceMaintenanceBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins circuit maintenance maintenance bulk update params
func (o *PluginsCircuitMaintenanceMaintenanceBulkUpdateParams) SetData(data *models.CircuitMaintenance) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsCircuitMaintenanceMaintenanceBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

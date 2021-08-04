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

// NewPluginsCircuitMaintenanceCircuitimpactBulkUpdateParams creates a new PluginsCircuitMaintenanceCircuitimpactBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsCircuitMaintenanceCircuitimpactBulkUpdateParams() *PluginsCircuitMaintenanceCircuitimpactBulkUpdateParams {
	return &PluginsCircuitMaintenanceCircuitimpactBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsCircuitMaintenanceCircuitimpactBulkUpdateParamsWithTimeout creates a new PluginsCircuitMaintenanceCircuitimpactBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsCircuitMaintenanceCircuitimpactBulkUpdateParamsWithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceCircuitimpactBulkUpdateParams {
	return &PluginsCircuitMaintenanceCircuitimpactBulkUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsCircuitMaintenanceCircuitimpactBulkUpdateParamsWithContext creates a new PluginsCircuitMaintenanceCircuitimpactBulkUpdateParams object
// with the ability to set a context for a request.
func NewPluginsCircuitMaintenanceCircuitimpactBulkUpdateParamsWithContext(ctx context.Context) *PluginsCircuitMaintenanceCircuitimpactBulkUpdateParams {
	return &PluginsCircuitMaintenanceCircuitimpactBulkUpdateParams{
		Context: ctx,
	}
}

// NewPluginsCircuitMaintenanceCircuitimpactBulkUpdateParamsWithHTTPClient creates a new PluginsCircuitMaintenanceCircuitimpactBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsCircuitMaintenanceCircuitimpactBulkUpdateParamsWithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceCircuitimpactBulkUpdateParams {
	return &PluginsCircuitMaintenanceCircuitimpactBulkUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsCircuitMaintenanceCircuitimpactBulkUpdateParams contains all the parameters to send to the API endpoint
   for the plugins circuit maintenance circuitimpact bulk update operation.

   Typically these are written to a http.Request.
*/
type PluginsCircuitMaintenanceCircuitimpactBulkUpdateParams struct {

	// Data.
	Data *models.CircuitMaintenanceCircuitImpact

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins circuit maintenance circuitimpact bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceCircuitimpactBulkUpdateParams) WithDefaults() *PluginsCircuitMaintenanceCircuitimpactBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins circuit maintenance circuitimpact bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceCircuitimpactBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins circuit maintenance circuitimpact bulk update params
func (o *PluginsCircuitMaintenanceCircuitimpactBulkUpdateParams) WithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceCircuitimpactBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins circuit maintenance circuitimpact bulk update params
func (o *PluginsCircuitMaintenanceCircuitimpactBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins circuit maintenance circuitimpact bulk update params
func (o *PluginsCircuitMaintenanceCircuitimpactBulkUpdateParams) WithContext(ctx context.Context) *PluginsCircuitMaintenanceCircuitimpactBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins circuit maintenance circuitimpact bulk update params
func (o *PluginsCircuitMaintenanceCircuitimpactBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins circuit maintenance circuitimpact bulk update params
func (o *PluginsCircuitMaintenanceCircuitimpactBulkUpdateParams) WithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceCircuitimpactBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins circuit maintenance circuitimpact bulk update params
func (o *PluginsCircuitMaintenanceCircuitimpactBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins circuit maintenance circuitimpact bulk update params
func (o *PluginsCircuitMaintenanceCircuitimpactBulkUpdateParams) WithData(data *models.CircuitMaintenanceCircuitImpact) *PluginsCircuitMaintenanceCircuitimpactBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins circuit maintenance circuitimpact bulk update params
func (o *PluginsCircuitMaintenanceCircuitimpactBulkUpdateParams) SetData(data *models.CircuitMaintenanceCircuitImpact) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsCircuitMaintenanceCircuitimpactBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

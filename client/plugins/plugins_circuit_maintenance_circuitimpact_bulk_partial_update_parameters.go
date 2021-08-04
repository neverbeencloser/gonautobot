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

// NewPluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateParams creates a new PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateParams() *PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateParams {
	return &PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateParamsWithTimeout creates a new PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateParams {
	return &PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateParamsWithContext creates a new PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewPluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateParamsWithContext(ctx context.Context) *PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateParams {
	return &PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewPluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateParamsWithHTTPClient creates a new PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateParams {
	return &PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the plugins circuit maintenance circuitimpact bulk partial update operation.

   Typically these are written to a http.Request.
*/
type PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateParams struct {

	// Data.
	Data *models.CircuitMaintenanceCircuitImpact

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins circuit maintenance circuitimpact bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateParams) WithDefaults() *PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins circuit maintenance circuitimpact bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins circuit maintenance circuitimpact bulk partial update params
func (o *PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins circuit maintenance circuitimpact bulk partial update params
func (o *PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins circuit maintenance circuitimpact bulk partial update params
func (o *PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateParams) WithContext(ctx context.Context) *PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins circuit maintenance circuitimpact bulk partial update params
func (o *PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins circuit maintenance circuitimpact bulk partial update params
func (o *PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins circuit maintenance circuitimpact bulk partial update params
func (o *PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins circuit maintenance circuitimpact bulk partial update params
func (o *PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateParams) WithData(data *models.CircuitMaintenanceCircuitImpact) *PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins circuit maintenance circuitimpact bulk partial update params
func (o *PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateParams) SetData(data *models.CircuitMaintenanceCircuitImpact) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

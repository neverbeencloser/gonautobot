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

// NewPluginsCircuitMaintenanceCircuitimpactCreateParams creates a new PluginsCircuitMaintenanceCircuitimpactCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsCircuitMaintenanceCircuitimpactCreateParams() *PluginsCircuitMaintenanceCircuitimpactCreateParams {
	return &PluginsCircuitMaintenanceCircuitimpactCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsCircuitMaintenanceCircuitimpactCreateParamsWithTimeout creates a new PluginsCircuitMaintenanceCircuitimpactCreateParams object
// with the ability to set a timeout on a request.
func NewPluginsCircuitMaintenanceCircuitimpactCreateParamsWithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceCircuitimpactCreateParams {
	return &PluginsCircuitMaintenanceCircuitimpactCreateParams{
		timeout: timeout,
	}
}

// NewPluginsCircuitMaintenanceCircuitimpactCreateParamsWithContext creates a new PluginsCircuitMaintenanceCircuitimpactCreateParams object
// with the ability to set a context for a request.
func NewPluginsCircuitMaintenanceCircuitimpactCreateParamsWithContext(ctx context.Context) *PluginsCircuitMaintenanceCircuitimpactCreateParams {
	return &PluginsCircuitMaintenanceCircuitimpactCreateParams{
		Context: ctx,
	}
}

// NewPluginsCircuitMaintenanceCircuitimpactCreateParamsWithHTTPClient creates a new PluginsCircuitMaintenanceCircuitimpactCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsCircuitMaintenanceCircuitimpactCreateParamsWithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceCircuitimpactCreateParams {
	return &PluginsCircuitMaintenanceCircuitimpactCreateParams{
		HTTPClient: client,
	}
}

/* PluginsCircuitMaintenanceCircuitimpactCreateParams contains all the parameters to send to the API endpoint
   for the plugins circuit maintenance circuitimpact create operation.

   Typically these are written to a http.Request.
*/
type PluginsCircuitMaintenanceCircuitimpactCreateParams struct {

	// Data.
	Data *models.CircuitMaintenanceCircuitImpact

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins circuit maintenance circuitimpact create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceCircuitimpactCreateParams) WithDefaults() *PluginsCircuitMaintenanceCircuitimpactCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins circuit maintenance circuitimpact create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceCircuitimpactCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins circuit maintenance circuitimpact create params
func (o *PluginsCircuitMaintenanceCircuitimpactCreateParams) WithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceCircuitimpactCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins circuit maintenance circuitimpact create params
func (o *PluginsCircuitMaintenanceCircuitimpactCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins circuit maintenance circuitimpact create params
func (o *PluginsCircuitMaintenanceCircuitimpactCreateParams) WithContext(ctx context.Context) *PluginsCircuitMaintenanceCircuitimpactCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins circuit maintenance circuitimpact create params
func (o *PluginsCircuitMaintenanceCircuitimpactCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins circuit maintenance circuitimpact create params
func (o *PluginsCircuitMaintenanceCircuitimpactCreateParams) WithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceCircuitimpactCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins circuit maintenance circuitimpact create params
func (o *PluginsCircuitMaintenanceCircuitimpactCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins circuit maintenance circuitimpact create params
func (o *PluginsCircuitMaintenanceCircuitimpactCreateParams) WithData(data *models.CircuitMaintenanceCircuitImpact) *PluginsCircuitMaintenanceCircuitimpactCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins circuit maintenance circuitimpact create params
func (o *PluginsCircuitMaintenanceCircuitimpactCreateParams) SetData(data *models.CircuitMaintenanceCircuitImpact) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsCircuitMaintenanceCircuitimpactCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

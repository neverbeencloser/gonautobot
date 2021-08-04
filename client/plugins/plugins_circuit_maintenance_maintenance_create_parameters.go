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

// NewPluginsCircuitMaintenanceMaintenanceCreateParams creates a new PluginsCircuitMaintenanceMaintenanceCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsCircuitMaintenanceMaintenanceCreateParams() *PluginsCircuitMaintenanceMaintenanceCreateParams {
	return &PluginsCircuitMaintenanceMaintenanceCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsCircuitMaintenanceMaintenanceCreateParamsWithTimeout creates a new PluginsCircuitMaintenanceMaintenanceCreateParams object
// with the ability to set a timeout on a request.
func NewPluginsCircuitMaintenanceMaintenanceCreateParamsWithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceMaintenanceCreateParams {
	return &PluginsCircuitMaintenanceMaintenanceCreateParams{
		timeout: timeout,
	}
}

// NewPluginsCircuitMaintenanceMaintenanceCreateParamsWithContext creates a new PluginsCircuitMaintenanceMaintenanceCreateParams object
// with the ability to set a context for a request.
func NewPluginsCircuitMaintenanceMaintenanceCreateParamsWithContext(ctx context.Context) *PluginsCircuitMaintenanceMaintenanceCreateParams {
	return &PluginsCircuitMaintenanceMaintenanceCreateParams{
		Context: ctx,
	}
}

// NewPluginsCircuitMaintenanceMaintenanceCreateParamsWithHTTPClient creates a new PluginsCircuitMaintenanceMaintenanceCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsCircuitMaintenanceMaintenanceCreateParamsWithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceMaintenanceCreateParams {
	return &PluginsCircuitMaintenanceMaintenanceCreateParams{
		HTTPClient: client,
	}
}

/* PluginsCircuitMaintenanceMaintenanceCreateParams contains all the parameters to send to the API endpoint
   for the plugins circuit maintenance maintenance create operation.

   Typically these are written to a http.Request.
*/
type PluginsCircuitMaintenanceMaintenanceCreateParams struct {

	// Data.
	Data *models.CircuitMaintenance

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins circuit maintenance maintenance create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceMaintenanceCreateParams) WithDefaults() *PluginsCircuitMaintenanceMaintenanceCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins circuit maintenance maintenance create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceMaintenanceCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins circuit maintenance maintenance create params
func (o *PluginsCircuitMaintenanceMaintenanceCreateParams) WithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceMaintenanceCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins circuit maintenance maintenance create params
func (o *PluginsCircuitMaintenanceMaintenanceCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins circuit maintenance maintenance create params
func (o *PluginsCircuitMaintenanceMaintenanceCreateParams) WithContext(ctx context.Context) *PluginsCircuitMaintenanceMaintenanceCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins circuit maintenance maintenance create params
func (o *PluginsCircuitMaintenanceMaintenanceCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins circuit maintenance maintenance create params
func (o *PluginsCircuitMaintenanceMaintenanceCreateParams) WithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceMaintenanceCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins circuit maintenance maintenance create params
func (o *PluginsCircuitMaintenanceMaintenanceCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins circuit maintenance maintenance create params
func (o *PluginsCircuitMaintenanceMaintenanceCreateParams) WithData(data *models.CircuitMaintenance) *PluginsCircuitMaintenanceMaintenanceCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins circuit maintenance maintenance create params
func (o *PluginsCircuitMaintenanceMaintenanceCreateParams) SetData(data *models.CircuitMaintenance) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsCircuitMaintenanceMaintenanceCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

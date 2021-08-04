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

// NewPluginsCircuitMaintenanceCircuitimpactUpdateParams creates a new PluginsCircuitMaintenanceCircuitimpactUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsCircuitMaintenanceCircuitimpactUpdateParams() *PluginsCircuitMaintenanceCircuitimpactUpdateParams {
	return &PluginsCircuitMaintenanceCircuitimpactUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsCircuitMaintenanceCircuitimpactUpdateParamsWithTimeout creates a new PluginsCircuitMaintenanceCircuitimpactUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsCircuitMaintenanceCircuitimpactUpdateParamsWithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceCircuitimpactUpdateParams {
	return &PluginsCircuitMaintenanceCircuitimpactUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsCircuitMaintenanceCircuitimpactUpdateParamsWithContext creates a new PluginsCircuitMaintenanceCircuitimpactUpdateParams object
// with the ability to set a context for a request.
func NewPluginsCircuitMaintenanceCircuitimpactUpdateParamsWithContext(ctx context.Context) *PluginsCircuitMaintenanceCircuitimpactUpdateParams {
	return &PluginsCircuitMaintenanceCircuitimpactUpdateParams{
		Context: ctx,
	}
}

// NewPluginsCircuitMaintenanceCircuitimpactUpdateParamsWithHTTPClient creates a new PluginsCircuitMaintenanceCircuitimpactUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsCircuitMaintenanceCircuitimpactUpdateParamsWithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceCircuitimpactUpdateParams {
	return &PluginsCircuitMaintenanceCircuitimpactUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsCircuitMaintenanceCircuitimpactUpdateParams contains all the parameters to send to the API endpoint
   for the plugins circuit maintenance circuitimpact update operation.

   Typically these are written to a http.Request.
*/
type PluginsCircuitMaintenanceCircuitimpactUpdateParams struct {

	// Data.
	Data *models.CircuitMaintenanceCircuitImpact

	/* ID.

	   A UUID string identifying this circuit impact.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins circuit maintenance circuitimpact update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceCircuitimpactUpdateParams) WithDefaults() *PluginsCircuitMaintenanceCircuitimpactUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins circuit maintenance circuitimpact update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceCircuitimpactUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins circuit maintenance circuitimpact update params
func (o *PluginsCircuitMaintenanceCircuitimpactUpdateParams) WithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceCircuitimpactUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins circuit maintenance circuitimpact update params
func (o *PluginsCircuitMaintenanceCircuitimpactUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins circuit maintenance circuitimpact update params
func (o *PluginsCircuitMaintenanceCircuitimpactUpdateParams) WithContext(ctx context.Context) *PluginsCircuitMaintenanceCircuitimpactUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins circuit maintenance circuitimpact update params
func (o *PluginsCircuitMaintenanceCircuitimpactUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins circuit maintenance circuitimpact update params
func (o *PluginsCircuitMaintenanceCircuitimpactUpdateParams) WithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceCircuitimpactUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins circuit maintenance circuitimpact update params
func (o *PluginsCircuitMaintenanceCircuitimpactUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins circuit maintenance circuitimpact update params
func (o *PluginsCircuitMaintenanceCircuitimpactUpdateParams) WithData(data *models.CircuitMaintenanceCircuitImpact) *PluginsCircuitMaintenanceCircuitimpactUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins circuit maintenance circuitimpact update params
func (o *PluginsCircuitMaintenanceCircuitimpactUpdateParams) SetData(data *models.CircuitMaintenanceCircuitImpact) {
	o.Data = data
}

// WithID adds the id to the plugins circuit maintenance circuitimpact update params
func (o *PluginsCircuitMaintenanceCircuitimpactUpdateParams) WithID(id strfmt.UUID) *PluginsCircuitMaintenanceCircuitimpactUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins circuit maintenance circuitimpact update params
func (o *PluginsCircuitMaintenanceCircuitimpactUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsCircuitMaintenanceCircuitimpactUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

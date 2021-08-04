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

// NewPluginsCircuitMaintenanceCircuitimpactPartialUpdateParams creates a new PluginsCircuitMaintenanceCircuitimpactPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsCircuitMaintenanceCircuitimpactPartialUpdateParams() *PluginsCircuitMaintenanceCircuitimpactPartialUpdateParams {
	return &PluginsCircuitMaintenanceCircuitimpactPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsCircuitMaintenanceCircuitimpactPartialUpdateParamsWithTimeout creates a new PluginsCircuitMaintenanceCircuitimpactPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsCircuitMaintenanceCircuitimpactPartialUpdateParamsWithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceCircuitimpactPartialUpdateParams {
	return &PluginsCircuitMaintenanceCircuitimpactPartialUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsCircuitMaintenanceCircuitimpactPartialUpdateParamsWithContext creates a new PluginsCircuitMaintenanceCircuitimpactPartialUpdateParams object
// with the ability to set a context for a request.
func NewPluginsCircuitMaintenanceCircuitimpactPartialUpdateParamsWithContext(ctx context.Context) *PluginsCircuitMaintenanceCircuitimpactPartialUpdateParams {
	return &PluginsCircuitMaintenanceCircuitimpactPartialUpdateParams{
		Context: ctx,
	}
}

// NewPluginsCircuitMaintenanceCircuitimpactPartialUpdateParamsWithHTTPClient creates a new PluginsCircuitMaintenanceCircuitimpactPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsCircuitMaintenanceCircuitimpactPartialUpdateParamsWithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceCircuitimpactPartialUpdateParams {
	return &PluginsCircuitMaintenanceCircuitimpactPartialUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsCircuitMaintenanceCircuitimpactPartialUpdateParams contains all the parameters to send to the API endpoint
   for the plugins circuit maintenance circuitimpact partial update operation.

   Typically these are written to a http.Request.
*/
type PluginsCircuitMaintenanceCircuitimpactPartialUpdateParams struct {

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

// WithDefaults hydrates default values in the plugins circuit maintenance circuitimpact partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceCircuitimpactPartialUpdateParams) WithDefaults() *PluginsCircuitMaintenanceCircuitimpactPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins circuit maintenance circuitimpact partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceCircuitimpactPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins circuit maintenance circuitimpact partial update params
func (o *PluginsCircuitMaintenanceCircuitimpactPartialUpdateParams) WithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceCircuitimpactPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins circuit maintenance circuitimpact partial update params
func (o *PluginsCircuitMaintenanceCircuitimpactPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins circuit maintenance circuitimpact partial update params
func (o *PluginsCircuitMaintenanceCircuitimpactPartialUpdateParams) WithContext(ctx context.Context) *PluginsCircuitMaintenanceCircuitimpactPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins circuit maintenance circuitimpact partial update params
func (o *PluginsCircuitMaintenanceCircuitimpactPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins circuit maintenance circuitimpact partial update params
func (o *PluginsCircuitMaintenanceCircuitimpactPartialUpdateParams) WithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceCircuitimpactPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins circuit maintenance circuitimpact partial update params
func (o *PluginsCircuitMaintenanceCircuitimpactPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins circuit maintenance circuitimpact partial update params
func (o *PluginsCircuitMaintenanceCircuitimpactPartialUpdateParams) WithData(data *models.CircuitMaintenanceCircuitImpact) *PluginsCircuitMaintenanceCircuitimpactPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins circuit maintenance circuitimpact partial update params
func (o *PluginsCircuitMaintenanceCircuitimpactPartialUpdateParams) SetData(data *models.CircuitMaintenanceCircuitImpact) {
	o.Data = data
}

// WithID adds the id to the plugins circuit maintenance circuitimpact partial update params
func (o *PluginsCircuitMaintenanceCircuitimpactPartialUpdateParams) WithID(id strfmt.UUID) *PluginsCircuitMaintenanceCircuitimpactPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins circuit maintenance circuitimpact partial update params
func (o *PluginsCircuitMaintenanceCircuitimpactPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsCircuitMaintenanceCircuitimpactPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

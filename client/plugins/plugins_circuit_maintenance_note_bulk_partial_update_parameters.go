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

// NewPluginsCircuitMaintenanceNoteBulkPartialUpdateParams creates a new PluginsCircuitMaintenanceNoteBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsCircuitMaintenanceNoteBulkPartialUpdateParams() *PluginsCircuitMaintenanceNoteBulkPartialUpdateParams {
	return &PluginsCircuitMaintenanceNoteBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsCircuitMaintenanceNoteBulkPartialUpdateParamsWithTimeout creates a new PluginsCircuitMaintenanceNoteBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsCircuitMaintenanceNoteBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceNoteBulkPartialUpdateParams {
	return &PluginsCircuitMaintenanceNoteBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsCircuitMaintenanceNoteBulkPartialUpdateParamsWithContext creates a new PluginsCircuitMaintenanceNoteBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewPluginsCircuitMaintenanceNoteBulkPartialUpdateParamsWithContext(ctx context.Context) *PluginsCircuitMaintenanceNoteBulkPartialUpdateParams {
	return &PluginsCircuitMaintenanceNoteBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewPluginsCircuitMaintenanceNoteBulkPartialUpdateParamsWithHTTPClient creates a new PluginsCircuitMaintenanceNoteBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsCircuitMaintenanceNoteBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceNoteBulkPartialUpdateParams {
	return &PluginsCircuitMaintenanceNoteBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsCircuitMaintenanceNoteBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the plugins circuit maintenance note bulk partial update operation.

   Typically these are written to a http.Request.
*/
type PluginsCircuitMaintenanceNoteBulkPartialUpdateParams struct {

	// Data.
	Data *models.Note

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins circuit maintenance note bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceNoteBulkPartialUpdateParams) WithDefaults() *PluginsCircuitMaintenanceNoteBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins circuit maintenance note bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceNoteBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins circuit maintenance note bulk partial update params
func (o *PluginsCircuitMaintenanceNoteBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceNoteBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins circuit maintenance note bulk partial update params
func (o *PluginsCircuitMaintenanceNoteBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins circuit maintenance note bulk partial update params
func (o *PluginsCircuitMaintenanceNoteBulkPartialUpdateParams) WithContext(ctx context.Context) *PluginsCircuitMaintenanceNoteBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins circuit maintenance note bulk partial update params
func (o *PluginsCircuitMaintenanceNoteBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins circuit maintenance note bulk partial update params
func (o *PluginsCircuitMaintenanceNoteBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceNoteBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins circuit maintenance note bulk partial update params
func (o *PluginsCircuitMaintenanceNoteBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins circuit maintenance note bulk partial update params
func (o *PluginsCircuitMaintenanceNoteBulkPartialUpdateParams) WithData(data *models.Note) *PluginsCircuitMaintenanceNoteBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins circuit maintenance note bulk partial update params
func (o *PluginsCircuitMaintenanceNoteBulkPartialUpdateParams) SetData(data *models.Note) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsCircuitMaintenanceNoteBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

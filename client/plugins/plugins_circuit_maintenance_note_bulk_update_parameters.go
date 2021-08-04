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

// NewPluginsCircuitMaintenanceNoteBulkUpdateParams creates a new PluginsCircuitMaintenanceNoteBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsCircuitMaintenanceNoteBulkUpdateParams() *PluginsCircuitMaintenanceNoteBulkUpdateParams {
	return &PluginsCircuitMaintenanceNoteBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsCircuitMaintenanceNoteBulkUpdateParamsWithTimeout creates a new PluginsCircuitMaintenanceNoteBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsCircuitMaintenanceNoteBulkUpdateParamsWithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceNoteBulkUpdateParams {
	return &PluginsCircuitMaintenanceNoteBulkUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsCircuitMaintenanceNoteBulkUpdateParamsWithContext creates a new PluginsCircuitMaintenanceNoteBulkUpdateParams object
// with the ability to set a context for a request.
func NewPluginsCircuitMaintenanceNoteBulkUpdateParamsWithContext(ctx context.Context) *PluginsCircuitMaintenanceNoteBulkUpdateParams {
	return &PluginsCircuitMaintenanceNoteBulkUpdateParams{
		Context: ctx,
	}
}

// NewPluginsCircuitMaintenanceNoteBulkUpdateParamsWithHTTPClient creates a new PluginsCircuitMaintenanceNoteBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsCircuitMaintenanceNoteBulkUpdateParamsWithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceNoteBulkUpdateParams {
	return &PluginsCircuitMaintenanceNoteBulkUpdateParams{
		HTTPClient: client,
	}
}

/* PluginsCircuitMaintenanceNoteBulkUpdateParams contains all the parameters to send to the API endpoint
   for the plugins circuit maintenance note bulk update operation.

   Typically these are written to a http.Request.
*/
type PluginsCircuitMaintenanceNoteBulkUpdateParams struct {

	// Data.
	Data *models.Note

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins circuit maintenance note bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceNoteBulkUpdateParams) WithDefaults() *PluginsCircuitMaintenanceNoteBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins circuit maintenance note bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceNoteBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins circuit maintenance note bulk update params
func (o *PluginsCircuitMaintenanceNoteBulkUpdateParams) WithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceNoteBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins circuit maintenance note bulk update params
func (o *PluginsCircuitMaintenanceNoteBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins circuit maintenance note bulk update params
func (o *PluginsCircuitMaintenanceNoteBulkUpdateParams) WithContext(ctx context.Context) *PluginsCircuitMaintenanceNoteBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins circuit maintenance note bulk update params
func (o *PluginsCircuitMaintenanceNoteBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins circuit maintenance note bulk update params
func (o *PluginsCircuitMaintenanceNoteBulkUpdateParams) WithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceNoteBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins circuit maintenance note bulk update params
func (o *PluginsCircuitMaintenanceNoteBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins circuit maintenance note bulk update params
func (o *PluginsCircuitMaintenanceNoteBulkUpdateParams) WithData(data *models.Note) *PluginsCircuitMaintenanceNoteBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins circuit maintenance note bulk update params
func (o *PluginsCircuitMaintenanceNoteBulkUpdateParams) SetData(data *models.Note) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsCircuitMaintenanceNoteBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

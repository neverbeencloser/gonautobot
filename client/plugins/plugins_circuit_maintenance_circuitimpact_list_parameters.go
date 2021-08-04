package plugins

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewPluginsCircuitMaintenanceCircuitimpactListParams creates a new PluginsCircuitMaintenanceCircuitimpactListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsCircuitMaintenanceCircuitimpactListParams() *PluginsCircuitMaintenanceCircuitimpactListParams {
	return &PluginsCircuitMaintenanceCircuitimpactListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsCircuitMaintenanceCircuitimpactListParamsWithTimeout creates a new PluginsCircuitMaintenanceCircuitimpactListParams object
// with the ability to set a timeout on a request.
func NewPluginsCircuitMaintenanceCircuitimpactListParamsWithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceCircuitimpactListParams {
	return &PluginsCircuitMaintenanceCircuitimpactListParams{
		timeout: timeout,
	}
}

// NewPluginsCircuitMaintenanceCircuitimpactListParamsWithContext creates a new PluginsCircuitMaintenanceCircuitimpactListParams object
// with the ability to set a context for a request.
func NewPluginsCircuitMaintenanceCircuitimpactListParamsWithContext(ctx context.Context) *PluginsCircuitMaintenanceCircuitimpactListParams {
	return &PluginsCircuitMaintenanceCircuitimpactListParams{
		Context: ctx,
	}
}

// NewPluginsCircuitMaintenanceCircuitimpactListParamsWithHTTPClient creates a new PluginsCircuitMaintenanceCircuitimpactListParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsCircuitMaintenanceCircuitimpactListParamsWithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceCircuitimpactListParams {
	return &PluginsCircuitMaintenanceCircuitimpactListParams{
		HTTPClient: client,
	}
}

/* PluginsCircuitMaintenanceCircuitimpactListParams contains all the parameters to send to the API endpoint
   for the plugins circuit maintenance circuitimpact list operation.

   Typically these are written to a http.Request.
*/
type PluginsCircuitMaintenanceCircuitimpactListParams struct {

	/* Limit.

	   Number of results to return per page.
	*/
	Limit *int64

	/* Offset.

	   The initial index from which to return the results.
	*/
	Offset *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins circuit maintenance circuitimpact list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceCircuitimpactListParams) WithDefaults() *PluginsCircuitMaintenanceCircuitimpactListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins circuit maintenance circuitimpact list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceCircuitimpactListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins circuit maintenance circuitimpact list params
func (o *PluginsCircuitMaintenanceCircuitimpactListParams) WithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceCircuitimpactListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins circuit maintenance circuitimpact list params
func (o *PluginsCircuitMaintenanceCircuitimpactListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins circuit maintenance circuitimpact list params
func (o *PluginsCircuitMaintenanceCircuitimpactListParams) WithContext(ctx context.Context) *PluginsCircuitMaintenanceCircuitimpactListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins circuit maintenance circuitimpact list params
func (o *PluginsCircuitMaintenanceCircuitimpactListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins circuit maintenance circuitimpact list params
func (o *PluginsCircuitMaintenanceCircuitimpactListParams) WithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceCircuitimpactListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins circuit maintenance circuitimpact list params
func (o *PluginsCircuitMaintenanceCircuitimpactListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the plugins circuit maintenance circuitimpact list params
func (o *PluginsCircuitMaintenanceCircuitimpactListParams) WithLimit(limit *int64) *PluginsCircuitMaintenanceCircuitimpactListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the plugins circuit maintenance circuitimpact list params
func (o *PluginsCircuitMaintenanceCircuitimpactListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the plugins circuit maintenance circuitimpact list params
func (o *PluginsCircuitMaintenanceCircuitimpactListParams) WithOffset(offset *int64) *PluginsCircuitMaintenanceCircuitimpactListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the plugins circuit maintenance circuitimpact list params
func (o *PluginsCircuitMaintenanceCircuitimpactListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsCircuitMaintenanceCircuitimpactListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

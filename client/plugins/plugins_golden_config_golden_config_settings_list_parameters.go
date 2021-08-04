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

// NewPluginsGoldenConfigGoldenConfigSettingsListParams creates a new PluginsGoldenConfigGoldenConfigSettingsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigGoldenConfigSettingsListParams() *PluginsGoldenConfigGoldenConfigSettingsListParams {
	return &PluginsGoldenConfigGoldenConfigSettingsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigGoldenConfigSettingsListParamsWithTimeout creates a new PluginsGoldenConfigGoldenConfigSettingsListParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigGoldenConfigSettingsListParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigGoldenConfigSettingsListParams {
	return &PluginsGoldenConfigGoldenConfigSettingsListParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigGoldenConfigSettingsListParamsWithContext creates a new PluginsGoldenConfigGoldenConfigSettingsListParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigGoldenConfigSettingsListParamsWithContext(ctx context.Context) *PluginsGoldenConfigGoldenConfigSettingsListParams {
	return &PluginsGoldenConfigGoldenConfigSettingsListParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigGoldenConfigSettingsListParamsWithHTTPClient creates a new PluginsGoldenConfigGoldenConfigSettingsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigGoldenConfigSettingsListParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigGoldenConfigSettingsListParams {
	return &PluginsGoldenConfigGoldenConfigSettingsListParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigGoldenConfigSettingsListParams contains all the parameters to send to the API endpoint
   for the plugins golden config golden config settings list operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigGoldenConfigSettingsListParams struct {

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

// WithDefaults hydrates default values in the plugins golden config golden config settings list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigGoldenConfigSettingsListParams) WithDefaults() *PluginsGoldenConfigGoldenConfigSettingsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config golden config settings list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigGoldenConfigSettingsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config golden config settings list params
func (o *PluginsGoldenConfigGoldenConfigSettingsListParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigGoldenConfigSettingsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config golden config settings list params
func (o *PluginsGoldenConfigGoldenConfigSettingsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config golden config settings list params
func (o *PluginsGoldenConfigGoldenConfigSettingsListParams) WithContext(ctx context.Context) *PluginsGoldenConfigGoldenConfigSettingsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config golden config settings list params
func (o *PluginsGoldenConfigGoldenConfigSettingsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config golden config settings list params
func (o *PluginsGoldenConfigGoldenConfigSettingsListParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigGoldenConfigSettingsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config golden config settings list params
func (o *PluginsGoldenConfigGoldenConfigSettingsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the plugins golden config golden config settings list params
func (o *PluginsGoldenConfigGoldenConfigSettingsListParams) WithLimit(limit *int64) *PluginsGoldenConfigGoldenConfigSettingsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the plugins golden config golden config settings list params
func (o *PluginsGoldenConfigGoldenConfigSettingsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the plugins golden config golden config settings list params
func (o *PluginsGoldenConfigGoldenConfigSettingsListParams) WithOffset(offset *int64) *PluginsGoldenConfigGoldenConfigSettingsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the plugins golden config golden config settings list params
func (o *PluginsGoldenConfigGoldenConfigSettingsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigGoldenConfigSettingsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

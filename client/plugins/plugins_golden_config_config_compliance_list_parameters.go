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

// NewPluginsGoldenConfigConfigComplianceListParams creates a new PluginsGoldenConfigConfigComplianceListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigConfigComplianceListParams() *PluginsGoldenConfigConfigComplianceListParams {
	return &PluginsGoldenConfigConfigComplianceListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigConfigComplianceListParamsWithTimeout creates a new PluginsGoldenConfigConfigComplianceListParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigConfigComplianceListParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigConfigComplianceListParams {
	return &PluginsGoldenConfigConfigComplianceListParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigConfigComplianceListParamsWithContext creates a new PluginsGoldenConfigConfigComplianceListParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigConfigComplianceListParamsWithContext(ctx context.Context) *PluginsGoldenConfigConfigComplianceListParams {
	return &PluginsGoldenConfigConfigComplianceListParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigConfigComplianceListParamsWithHTTPClient creates a new PluginsGoldenConfigConfigComplianceListParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigConfigComplianceListParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigConfigComplianceListParams {
	return &PluginsGoldenConfigConfigComplianceListParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigConfigComplianceListParams contains all the parameters to send to the API endpoint
   for the plugins golden config config compliance list operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigConfigComplianceListParams struct {

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

// WithDefaults hydrates default values in the plugins golden config config compliance list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigConfigComplianceListParams) WithDefaults() *PluginsGoldenConfigConfigComplianceListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config config compliance list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigConfigComplianceListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config config compliance list params
func (o *PluginsGoldenConfigConfigComplianceListParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigConfigComplianceListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config config compliance list params
func (o *PluginsGoldenConfigConfigComplianceListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config config compliance list params
func (o *PluginsGoldenConfigConfigComplianceListParams) WithContext(ctx context.Context) *PluginsGoldenConfigConfigComplianceListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config config compliance list params
func (o *PluginsGoldenConfigConfigComplianceListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config config compliance list params
func (o *PluginsGoldenConfigConfigComplianceListParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigConfigComplianceListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config config compliance list params
func (o *PluginsGoldenConfigConfigComplianceListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the plugins golden config config compliance list params
func (o *PluginsGoldenConfigConfigComplianceListParams) WithLimit(limit *int64) *PluginsGoldenConfigConfigComplianceListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the plugins golden config config compliance list params
func (o *PluginsGoldenConfigConfigComplianceListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the plugins golden config config compliance list params
func (o *PluginsGoldenConfigConfigComplianceListParams) WithOffset(offset *int64) *PluginsGoldenConfigConfigComplianceListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the plugins golden config config compliance list params
func (o *PluginsGoldenConfigConfigComplianceListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigConfigComplianceListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

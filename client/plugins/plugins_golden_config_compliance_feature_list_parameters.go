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

// NewPluginsGoldenConfigComplianceFeatureListParams creates a new PluginsGoldenConfigComplianceFeatureListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigComplianceFeatureListParams() *PluginsGoldenConfigComplianceFeatureListParams {
	return &PluginsGoldenConfigComplianceFeatureListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigComplianceFeatureListParamsWithTimeout creates a new PluginsGoldenConfigComplianceFeatureListParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigComplianceFeatureListParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigComplianceFeatureListParams {
	return &PluginsGoldenConfigComplianceFeatureListParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigComplianceFeatureListParamsWithContext creates a new PluginsGoldenConfigComplianceFeatureListParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigComplianceFeatureListParamsWithContext(ctx context.Context) *PluginsGoldenConfigComplianceFeatureListParams {
	return &PluginsGoldenConfigComplianceFeatureListParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigComplianceFeatureListParamsWithHTTPClient creates a new PluginsGoldenConfigComplianceFeatureListParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigComplianceFeatureListParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigComplianceFeatureListParams {
	return &PluginsGoldenConfigComplianceFeatureListParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigComplianceFeatureListParams contains all the parameters to send to the API endpoint
   for the plugins golden config compliance feature list operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigComplianceFeatureListParams struct {

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

// WithDefaults hydrates default values in the plugins golden config compliance feature list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigComplianceFeatureListParams) WithDefaults() *PluginsGoldenConfigComplianceFeatureListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config compliance feature list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigComplianceFeatureListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config compliance feature list params
func (o *PluginsGoldenConfigComplianceFeatureListParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigComplianceFeatureListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config compliance feature list params
func (o *PluginsGoldenConfigComplianceFeatureListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config compliance feature list params
func (o *PluginsGoldenConfigComplianceFeatureListParams) WithContext(ctx context.Context) *PluginsGoldenConfigComplianceFeatureListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config compliance feature list params
func (o *PluginsGoldenConfigComplianceFeatureListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config compliance feature list params
func (o *PluginsGoldenConfigComplianceFeatureListParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigComplianceFeatureListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config compliance feature list params
func (o *PluginsGoldenConfigComplianceFeatureListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the plugins golden config compliance feature list params
func (o *PluginsGoldenConfigComplianceFeatureListParams) WithLimit(limit *int64) *PluginsGoldenConfigComplianceFeatureListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the plugins golden config compliance feature list params
func (o *PluginsGoldenConfigComplianceFeatureListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the plugins golden config compliance feature list params
func (o *PluginsGoldenConfigComplianceFeatureListParams) WithOffset(offset *int64) *PluginsGoldenConfigComplianceFeatureListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the plugins golden config compliance feature list params
func (o *PluginsGoldenConfigComplianceFeatureListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigComplianceFeatureListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

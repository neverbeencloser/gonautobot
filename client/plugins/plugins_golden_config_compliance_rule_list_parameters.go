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

// NewPluginsGoldenConfigComplianceRuleListParams creates a new PluginsGoldenConfigComplianceRuleListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigComplianceRuleListParams() *PluginsGoldenConfigComplianceRuleListParams {
	return &PluginsGoldenConfigComplianceRuleListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigComplianceRuleListParamsWithTimeout creates a new PluginsGoldenConfigComplianceRuleListParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigComplianceRuleListParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigComplianceRuleListParams {
	return &PluginsGoldenConfigComplianceRuleListParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigComplianceRuleListParamsWithContext creates a new PluginsGoldenConfigComplianceRuleListParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigComplianceRuleListParamsWithContext(ctx context.Context) *PluginsGoldenConfigComplianceRuleListParams {
	return &PluginsGoldenConfigComplianceRuleListParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigComplianceRuleListParamsWithHTTPClient creates a new PluginsGoldenConfigComplianceRuleListParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigComplianceRuleListParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigComplianceRuleListParams {
	return &PluginsGoldenConfigComplianceRuleListParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigComplianceRuleListParams contains all the parameters to send to the API endpoint
   for the plugins golden config compliance rule list operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigComplianceRuleListParams struct {

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

// WithDefaults hydrates default values in the plugins golden config compliance rule list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigComplianceRuleListParams) WithDefaults() *PluginsGoldenConfigComplianceRuleListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config compliance rule list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigComplianceRuleListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config compliance rule list params
func (o *PluginsGoldenConfigComplianceRuleListParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigComplianceRuleListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config compliance rule list params
func (o *PluginsGoldenConfigComplianceRuleListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config compliance rule list params
func (o *PluginsGoldenConfigComplianceRuleListParams) WithContext(ctx context.Context) *PluginsGoldenConfigComplianceRuleListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config compliance rule list params
func (o *PluginsGoldenConfigComplianceRuleListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config compliance rule list params
func (o *PluginsGoldenConfigComplianceRuleListParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigComplianceRuleListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config compliance rule list params
func (o *PluginsGoldenConfigComplianceRuleListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the plugins golden config compliance rule list params
func (o *PluginsGoldenConfigComplianceRuleListParams) WithLimit(limit *int64) *PluginsGoldenConfigComplianceRuleListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the plugins golden config compliance rule list params
func (o *PluginsGoldenConfigComplianceRuleListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the plugins golden config compliance rule list params
func (o *PluginsGoldenConfigComplianceRuleListParams) WithOffset(offset *int64) *PluginsGoldenConfigComplianceRuleListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the plugins golden config compliance rule list params
func (o *PluginsGoldenConfigComplianceRuleListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigComplianceRuleListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

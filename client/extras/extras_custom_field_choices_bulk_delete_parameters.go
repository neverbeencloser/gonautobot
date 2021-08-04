package extras

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewExtrasCustomFieldChoicesBulkDeleteParams creates a new ExtrasCustomFieldChoicesBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasCustomFieldChoicesBulkDeleteParams() *ExtrasCustomFieldChoicesBulkDeleteParams {
	return &ExtrasCustomFieldChoicesBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasCustomFieldChoicesBulkDeleteParamsWithTimeout creates a new ExtrasCustomFieldChoicesBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewExtrasCustomFieldChoicesBulkDeleteParamsWithTimeout(timeout time.Duration) *ExtrasCustomFieldChoicesBulkDeleteParams {
	return &ExtrasCustomFieldChoicesBulkDeleteParams{
		timeout: timeout,
	}
}

// NewExtrasCustomFieldChoicesBulkDeleteParamsWithContext creates a new ExtrasCustomFieldChoicesBulkDeleteParams object
// with the ability to set a context for a request.
func NewExtrasCustomFieldChoicesBulkDeleteParamsWithContext(ctx context.Context) *ExtrasCustomFieldChoicesBulkDeleteParams {
	return &ExtrasCustomFieldChoicesBulkDeleteParams{
		Context: ctx,
	}
}

// NewExtrasCustomFieldChoicesBulkDeleteParamsWithHTTPClient creates a new ExtrasCustomFieldChoicesBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasCustomFieldChoicesBulkDeleteParamsWithHTTPClient(client *http.Client) *ExtrasCustomFieldChoicesBulkDeleteParams {
	return &ExtrasCustomFieldChoicesBulkDeleteParams{
		HTTPClient: client,
	}
}

/* ExtrasCustomFieldChoicesBulkDeleteParams contains all the parameters to send to the API endpoint
   for the extras custom field choices bulk delete operation.

   Typically these are written to a http.Request.
*/
type ExtrasCustomFieldChoicesBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras custom field choices bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasCustomFieldChoicesBulkDeleteParams) WithDefaults() *ExtrasCustomFieldChoicesBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras custom field choices bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasCustomFieldChoicesBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras custom field choices bulk delete params
func (o *ExtrasCustomFieldChoicesBulkDeleteParams) WithTimeout(timeout time.Duration) *ExtrasCustomFieldChoicesBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras custom field choices bulk delete params
func (o *ExtrasCustomFieldChoicesBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras custom field choices bulk delete params
func (o *ExtrasCustomFieldChoicesBulkDeleteParams) WithContext(ctx context.Context) *ExtrasCustomFieldChoicesBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras custom field choices bulk delete params
func (o *ExtrasCustomFieldChoicesBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras custom field choices bulk delete params
func (o *ExtrasCustomFieldChoicesBulkDeleteParams) WithHTTPClient(client *http.Client) *ExtrasCustomFieldChoicesBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras custom field choices bulk delete params
func (o *ExtrasCustomFieldChoicesBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasCustomFieldChoicesBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

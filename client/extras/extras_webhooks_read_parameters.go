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

// NewExtrasWebhooksReadParams creates a new ExtrasWebhooksReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasWebhooksReadParams() *ExtrasWebhooksReadParams {
	return &ExtrasWebhooksReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasWebhooksReadParamsWithTimeout creates a new ExtrasWebhooksReadParams object
// with the ability to set a timeout on a request.
func NewExtrasWebhooksReadParamsWithTimeout(timeout time.Duration) *ExtrasWebhooksReadParams {
	return &ExtrasWebhooksReadParams{
		timeout: timeout,
	}
}

// NewExtrasWebhooksReadParamsWithContext creates a new ExtrasWebhooksReadParams object
// with the ability to set a context for a request.
func NewExtrasWebhooksReadParamsWithContext(ctx context.Context) *ExtrasWebhooksReadParams {
	return &ExtrasWebhooksReadParams{
		Context: ctx,
	}
}

// NewExtrasWebhooksReadParamsWithHTTPClient creates a new ExtrasWebhooksReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasWebhooksReadParamsWithHTTPClient(client *http.Client) *ExtrasWebhooksReadParams {
	return &ExtrasWebhooksReadParams{
		HTTPClient: client,
	}
}

/* ExtrasWebhooksReadParams contains all the parameters to send to the API endpoint
   for the extras webhooks read operation.

   Typically these are written to a http.Request.
*/
type ExtrasWebhooksReadParams struct {

	/* ID.

	   A UUID string identifying this webhook.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras webhooks read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasWebhooksReadParams) WithDefaults() *ExtrasWebhooksReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras webhooks read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasWebhooksReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras webhooks read params
func (o *ExtrasWebhooksReadParams) WithTimeout(timeout time.Duration) *ExtrasWebhooksReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras webhooks read params
func (o *ExtrasWebhooksReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras webhooks read params
func (o *ExtrasWebhooksReadParams) WithContext(ctx context.Context) *ExtrasWebhooksReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras webhooks read params
func (o *ExtrasWebhooksReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras webhooks read params
func (o *ExtrasWebhooksReadParams) WithHTTPClient(client *http.Client) *ExtrasWebhooksReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras webhooks read params
func (o *ExtrasWebhooksReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the extras webhooks read params
func (o *ExtrasWebhooksReadParams) WithID(id strfmt.UUID) *ExtrasWebhooksReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras webhooks read params
func (o *ExtrasWebhooksReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasWebhooksReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

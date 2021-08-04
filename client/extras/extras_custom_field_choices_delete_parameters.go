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

// NewExtrasCustomFieldChoicesDeleteParams creates a new ExtrasCustomFieldChoicesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasCustomFieldChoicesDeleteParams() *ExtrasCustomFieldChoicesDeleteParams {
	return &ExtrasCustomFieldChoicesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasCustomFieldChoicesDeleteParamsWithTimeout creates a new ExtrasCustomFieldChoicesDeleteParams object
// with the ability to set a timeout on a request.
func NewExtrasCustomFieldChoicesDeleteParamsWithTimeout(timeout time.Duration) *ExtrasCustomFieldChoicesDeleteParams {
	return &ExtrasCustomFieldChoicesDeleteParams{
		timeout: timeout,
	}
}

// NewExtrasCustomFieldChoicesDeleteParamsWithContext creates a new ExtrasCustomFieldChoicesDeleteParams object
// with the ability to set a context for a request.
func NewExtrasCustomFieldChoicesDeleteParamsWithContext(ctx context.Context) *ExtrasCustomFieldChoicesDeleteParams {
	return &ExtrasCustomFieldChoicesDeleteParams{
		Context: ctx,
	}
}

// NewExtrasCustomFieldChoicesDeleteParamsWithHTTPClient creates a new ExtrasCustomFieldChoicesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasCustomFieldChoicesDeleteParamsWithHTTPClient(client *http.Client) *ExtrasCustomFieldChoicesDeleteParams {
	return &ExtrasCustomFieldChoicesDeleteParams{
		HTTPClient: client,
	}
}

/* ExtrasCustomFieldChoicesDeleteParams contains all the parameters to send to the API endpoint
   for the extras custom field choices delete operation.

   Typically these are written to a http.Request.
*/
type ExtrasCustomFieldChoicesDeleteParams struct {

	/* ID.

	   A UUID string identifying this custom field choice.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras custom field choices delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasCustomFieldChoicesDeleteParams) WithDefaults() *ExtrasCustomFieldChoicesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras custom field choices delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasCustomFieldChoicesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras custom field choices delete params
func (o *ExtrasCustomFieldChoicesDeleteParams) WithTimeout(timeout time.Duration) *ExtrasCustomFieldChoicesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras custom field choices delete params
func (o *ExtrasCustomFieldChoicesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras custom field choices delete params
func (o *ExtrasCustomFieldChoicesDeleteParams) WithContext(ctx context.Context) *ExtrasCustomFieldChoicesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras custom field choices delete params
func (o *ExtrasCustomFieldChoicesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras custom field choices delete params
func (o *ExtrasCustomFieldChoicesDeleteParams) WithHTTPClient(client *http.Client) *ExtrasCustomFieldChoicesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras custom field choices delete params
func (o *ExtrasCustomFieldChoicesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the extras custom field choices delete params
func (o *ExtrasCustomFieldChoicesDeleteParams) WithID(id strfmt.UUID) *ExtrasCustomFieldChoicesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras custom field choices delete params
func (o *ExtrasCustomFieldChoicesDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasCustomFieldChoicesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

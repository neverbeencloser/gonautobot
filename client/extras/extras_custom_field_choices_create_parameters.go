package extras

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

// NewExtrasCustomFieldChoicesCreateParams creates a new ExtrasCustomFieldChoicesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasCustomFieldChoicesCreateParams() *ExtrasCustomFieldChoicesCreateParams {
	return &ExtrasCustomFieldChoicesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasCustomFieldChoicesCreateParamsWithTimeout creates a new ExtrasCustomFieldChoicesCreateParams object
// with the ability to set a timeout on a request.
func NewExtrasCustomFieldChoicesCreateParamsWithTimeout(timeout time.Duration) *ExtrasCustomFieldChoicesCreateParams {
	return &ExtrasCustomFieldChoicesCreateParams{
		timeout: timeout,
	}
}

// NewExtrasCustomFieldChoicesCreateParamsWithContext creates a new ExtrasCustomFieldChoicesCreateParams object
// with the ability to set a context for a request.
func NewExtrasCustomFieldChoicesCreateParamsWithContext(ctx context.Context) *ExtrasCustomFieldChoicesCreateParams {
	return &ExtrasCustomFieldChoicesCreateParams{
		Context: ctx,
	}
}

// NewExtrasCustomFieldChoicesCreateParamsWithHTTPClient creates a new ExtrasCustomFieldChoicesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasCustomFieldChoicesCreateParamsWithHTTPClient(client *http.Client) *ExtrasCustomFieldChoicesCreateParams {
	return &ExtrasCustomFieldChoicesCreateParams{
		HTTPClient: client,
	}
}

/* ExtrasCustomFieldChoicesCreateParams contains all the parameters to send to the API endpoint
   for the extras custom field choices create operation.

   Typically these are written to a http.Request.
*/
type ExtrasCustomFieldChoicesCreateParams struct {

	// Data.
	Data *models.WritableCustomFieldChoice

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras custom field choices create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasCustomFieldChoicesCreateParams) WithDefaults() *ExtrasCustomFieldChoicesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras custom field choices create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasCustomFieldChoicesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras custom field choices create params
func (o *ExtrasCustomFieldChoicesCreateParams) WithTimeout(timeout time.Duration) *ExtrasCustomFieldChoicesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras custom field choices create params
func (o *ExtrasCustomFieldChoicesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras custom field choices create params
func (o *ExtrasCustomFieldChoicesCreateParams) WithContext(ctx context.Context) *ExtrasCustomFieldChoicesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras custom field choices create params
func (o *ExtrasCustomFieldChoicesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras custom field choices create params
func (o *ExtrasCustomFieldChoicesCreateParams) WithHTTPClient(client *http.Client) *ExtrasCustomFieldChoicesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras custom field choices create params
func (o *ExtrasCustomFieldChoicesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras custom field choices create params
func (o *ExtrasCustomFieldChoicesCreateParams) WithData(data *models.WritableCustomFieldChoice) *ExtrasCustomFieldChoicesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras custom field choices create params
func (o *ExtrasCustomFieldChoicesCreateParams) SetData(data *models.WritableCustomFieldChoice) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasCustomFieldChoicesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

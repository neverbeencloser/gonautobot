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

// NewExtrasStatusesCreateParams creates a new ExtrasStatusesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasStatusesCreateParams() *ExtrasStatusesCreateParams {
	return &ExtrasStatusesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasStatusesCreateParamsWithTimeout creates a new ExtrasStatusesCreateParams object
// with the ability to set a timeout on a request.
func NewExtrasStatusesCreateParamsWithTimeout(timeout time.Duration) *ExtrasStatusesCreateParams {
	return &ExtrasStatusesCreateParams{
		timeout: timeout,
	}
}

// NewExtrasStatusesCreateParamsWithContext creates a new ExtrasStatusesCreateParams object
// with the ability to set a context for a request.
func NewExtrasStatusesCreateParamsWithContext(ctx context.Context) *ExtrasStatusesCreateParams {
	return &ExtrasStatusesCreateParams{
		Context: ctx,
	}
}

// NewExtrasStatusesCreateParamsWithHTTPClient creates a new ExtrasStatusesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasStatusesCreateParamsWithHTTPClient(client *http.Client) *ExtrasStatusesCreateParams {
	return &ExtrasStatusesCreateParams{
		HTTPClient: client,
	}
}

/* ExtrasStatusesCreateParams contains all the parameters to send to the API endpoint
   for the extras statuses create operation.

   Typically these are written to a http.Request.
*/
type ExtrasStatusesCreateParams struct {

	// Data.
	Data *models.Status

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras statuses create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasStatusesCreateParams) WithDefaults() *ExtrasStatusesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras statuses create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasStatusesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras statuses create params
func (o *ExtrasStatusesCreateParams) WithTimeout(timeout time.Duration) *ExtrasStatusesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras statuses create params
func (o *ExtrasStatusesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras statuses create params
func (o *ExtrasStatusesCreateParams) WithContext(ctx context.Context) *ExtrasStatusesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras statuses create params
func (o *ExtrasStatusesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras statuses create params
func (o *ExtrasStatusesCreateParams) WithHTTPClient(client *http.Client) *ExtrasStatusesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras statuses create params
func (o *ExtrasStatusesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras statuses create params
func (o *ExtrasStatusesCreateParams) WithData(data *models.Status) *ExtrasStatusesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras statuses create params
func (o *ExtrasStatusesCreateParams) SetData(data *models.Status) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasStatusesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

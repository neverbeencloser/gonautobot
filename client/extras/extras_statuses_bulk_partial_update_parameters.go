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

// NewExtrasStatusesBulkPartialUpdateParams creates a new ExtrasStatusesBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasStatusesBulkPartialUpdateParams() *ExtrasStatusesBulkPartialUpdateParams {
	return &ExtrasStatusesBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasStatusesBulkPartialUpdateParamsWithTimeout creates a new ExtrasStatusesBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasStatusesBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *ExtrasStatusesBulkPartialUpdateParams {
	return &ExtrasStatusesBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasStatusesBulkPartialUpdateParamsWithContext creates a new ExtrasStatusesBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewExtrasStatusesBulkPartialUpdateParamsWithContext(ctx context.Context) *ExtrasStatusesBulkPartialUpdateParams {
	return &ExtrasStatusesBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewExtrasStatusesBulkPartialUpdateParamsWithHTTPClient creates a new ExtrasStatusesBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasStatusesBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *ExtrasStatusesBulkPartialUpdateParams {
	return &ExtrasStatusesBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasStatusesBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the extras statuses bulk partial update operation.

   Typically these are written to a http.Request.
*/
type ExtrasStatusesBulkPartialUpdateParams struct {

	// Data.
	Data *models.Status

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras statuses bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasStatusesBulkPartialUpdateParams) WithDefaults() *ExtrasStatusesBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras statuses bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasStatusesBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras statuses bulk partial update params
func (o *ExtrasStatusesBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *ExtrasStatusesBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras statuses bulk partial update params
func (o *ExtrasStatusesBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras statuses bulk partial update params
func (o *ExtrasStatusesBulkPartialUpdateParams) WithContext(ctx context.Context) *ExtrasStatusesBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras statuses bulk partial update params
func (o *ExtrasStatusesBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras statuses bulk partial update params
func (o *ExtrasStatusesBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *ExtrasStatusesBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras statuses bulk partial update params
func (o *ExtrasStatusesBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras statuses bulk partial update params
func (o *ExtrasStatusesBulkPartialUpdateParams) WithData(data *models.Status) *ExtrasStatusesBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras statuses bulk partial update params
func (o *ExtrasStatusesBulkPartialUpdateParams) SetData(data *models.Status) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasStatusesBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

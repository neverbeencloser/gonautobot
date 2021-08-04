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

// NewExtrasStatusesBulkUpdateParams creates a new ExtrasStatusesBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasStatusesBulkUpdateParams() *ExtrasStatusesBulkUpdateParams {
	return &ExtrasStatusesBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasStatusesBulkUpdateParamsWithTimeout creates a new ExtrasStatusesBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasStatusesBulkUpdateParamsWithTimeout(timeout time.Duration) *ExtrasStatusesBulkUpdateParams {
	return &ExtrasStatusesBulkUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasStatusesBulkUpdateParamsWithContext creates a new ExtrasStatusesBulkUpdateParams object
// with the ability to set a context for a request.
func NewExtrasStatusesBulkUpdateParamsWithContext(ctx context.Context) *ExtrasStatusesBulkUpdateParams {
	return &ExtrasStatusesBulkUpdateParams{
		Context: ctx,
	}
}

// NewExtrasStatusesBulkUpdateParamsWithHTTPClient creates a new ExtrasStatusesBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasStatusesBulkUpdateParamsWithHTTPClient(client *http.Client) *ExtrasStatusesBulkUpdateParams {
	return &ExtrasStatusesBulkUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasStatusesBulkUpdateParams contains all the parameters to send to the API endpoint
   for the extras statuses bulk update operation.

   Typically these are written to a http.Request.
*/
type ExtrasStatusesBulkUpdateParams struct {

	// Data.
	Data *models.Status

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras statuses bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasStatusesBulkUpdateParams) WithDefaults() *ExtrasStatusesBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras statuses bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasStatusesBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras statuses bulk update params
func (o *ExtrasStatusesBulkUpdateParams) WithTimeout(timeout time.Duration) *ExtrasStatusesBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras statuses bulk update params
func (o *ExtrasStatusesBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras statuses bulk update params
func (o *ExtrasStatusesBulkUpdateParams) WithContext(ctx context.Context) *ExtrasStatusesBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras statuses bulk update params
func (o *ExtrasStatusesBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras statuses bulk update params
func (o *ExtrasStatusesBulkUpdateParams) WithHTTPClient(client *http.Client) *ExtrasStatusesBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras statuses bulk update params
func (o *ExtrasStatusesBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras statuses bulk update params
func (o *ExtrasStatusesBulkUpdateParams) WithData(data *models.Status) *ExtrasStatusesBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras statuses bulk update params
func (o *ExtrasStatusesBulkUpdateParams) SetData(data *models.Status) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasStatusesBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

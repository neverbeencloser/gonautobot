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

// NewExtrasJobResultsBulkPartialUpdateParams creates a new ExtrasJobResultsBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasJobResultsBulkPartialUpdateParams() *ExtrasJobResultsBulkPartialUpdateParams {
	return &ExtrasJobResultsBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasJobResultsBulkPartialUpdateParamsWithTimeout creates a new ExtrasJobResultsBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasJobResultsBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *ExtrasJobResultsBulkPartialUpdateParams {
	return &ExtrasJobResultsBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasJobResultsBulkPartialUpdateParamsWithContext creates a new ExtrasJobResultsBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewExtrasJobResultsBulkPartialUpdateParamsWithContext(ctx context.Context) *ExtrasJobResultsBulkPartialUpdateParams {
	return &ExtrasJobResultsBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewExtrasJobResultsBulkPartialUpdateParamsWithHTTPClient creates a new ExtrasJobResultsBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasJobResultsBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *ExtrasJobResultsBulkPartialUpdateParams {
	return &ExtrasJobResultsBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasJobResultsBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the extras job results bulk partial update operation.

   Typically these are written to a http.Request.
*/
type ExtrasJobResultsBulkPartialUpdateParams struct {

	// Data.
	Data *models.WritableJobResult

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras job results bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasJobResultsBulkPartialUpdateParams) WithDefaults() *ExtrasJobResultsBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras job results bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasJobResultsBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras job results bulk partial update params
func (o *ExtrasJobResultsBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *ExtrasJobResultsBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras job results bulk partial update params
func (o *ExtrasJobResultsBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras job results bulk partial update params
func (o *ExtrasJobResultsBulkPartialUpdateParams) WithContext(ctx context.Context) *ExtrasJobResultsBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras job results bulk partial update params
func (o *ExtrasJobResultsBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras job results bulk partial update params
func (o *ExtrasJobResultsBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *ExtrasJobResultsBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras job results bulk partial update params
func (o *ExtrasJobResultsBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras job results bulk partial update params
func (o *ExtrasJobResultsBulkPartialUpdateParams) WithData(data *models.WritableJobResult) *ExtrasJobResultsBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras job results bulk partial update params
func (o *ExtrasJobResultsBulkPartialUpdateParams) SetData(data *models.WritableJobResult) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasJobResultsBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

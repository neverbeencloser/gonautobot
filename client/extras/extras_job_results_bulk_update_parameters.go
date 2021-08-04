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

// NewExtrasJobResultsBulkUpdateParams creates a new ExtrasJobResultsBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasJobResultsBulkUpdateParams() *ExtrasJobResultsBulkUpdateParams {
	return &ExtrasJobResultsBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasJobResultsBulkUpdateParamsWithTimeout creates a new ExtrasJobResultsBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasJobResultsBulkUpdateParamsWithTimeout(timeout time.Duration) *ExtrasJobResultsBulkUpdateParams {
	return &ExtrasJobResultsBulkUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasJobResultsBulkUpdateParamsWithContext creates a new ExtrasJobResultsBulkUpdateParams object
// with the ability to set a context for a request.
func NewExtrasJobResultsBulkUpdateParamsWithContext(ctx context.Context) *ExtrasJobResultsBulkUpdateParams {
	return &ExtrasJobResultsBulkUpdateParams{
		Context: ctx,
	}
}

// NewExtrasJobResultsBulkUpdateParamsWithHTTPClient creates a new ExtrasJobResultsBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasJobResultsBulkUpdateParamsWithHTTPClient(client *http.Client) *ExtrasJobResultsBulkUpdateParams {
	return &ExtrasJobResultsBulkUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasJobResultsBulkUpdateParams contains all the parameters to send to the API endpoint
   for the extras job results bulk update operation.

   Typically these are written to a http.Request.
*/
type ExtrasJobResultsBulkUpdateParams struct {

	// Data.
	Data *models.WritableJobResult

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras job results bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasJobResultsBulkUpdateParams) WithDefaults() *ExtrasJobResultsBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras job results bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasJobResultsBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras job results bulk update params
func (o *ExtrasJobResultsBulkUpdateParams) WithTimeout(timeout time.Duration) *ExtrasJobResultsBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras job results bulk update params
func (o *ExtrasJobResultsBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras job results bulk update params
func (o *ExtrasJobResultsBulkUpdateParams) WithContext(ctx context.Context) *ExtrasJobResultsBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras job results bulk update params
func (o *ExtrasJobResultsBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras job results bulk update params
func (o *ExtrasJobResultsBulkUpdateParams) WithHTTPClient(client *http.Client) *ExtrasJobResultsBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras job results bulk update params
func (o *ExtrasJobResultsBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras job results bulk update params
func (o *ExtrasJobResultsBulkUpdateParams) WithData(data *models.WritableJobResult) *ExtrasJobResultsBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras job results bulk update params
func (o *ExtrasJobResultsBulkUpdateParams) SetData(data *models.WritableJobResult) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasJobResultsBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

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

// NewExtrasJobResultsUpdateParams creates a new ExtrasJobResultsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasJobResultsUpdateParams() *ExtrasJobResultsUpdateParams {
	return &ExtrasJobResultsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasJobResultsUpdateParamsWithTimeout creates a new ExtrasJobResultsUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasJobResultsUpdateParamsWithTimeout(timeout time.Duration) *ExtrasJobResultsUpdateParams {
	return &ExtrasJobResultsUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasJobResultsUpdateParamsWithContext creates a new ExtrasJobResultsUpdateParams object
// with the ability to set a context for a request.
func NewExtrasJobResultsUpdateParamsWithContext(ctx context.Context) *ExtrasJobResultsUpdateParams {
	return &ExtrasJobResultsUpdateParams{
		Context: ctx,
	}
}

// NewExtrasJobResultsUpdateParamsWithHTTPClient creates a new ExtrasJobResultsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasJobResultsUpdateParamsWithHTTPClient(client *http.Client) *ExtrasJobResultsUpdateParams {
	return &ExtrasJobResultsUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasJobResultsUpdateParams contains all the parameters to send to the API endpoint
   for the extras job results update operation.

   Typically these are written to a http.Request.
*/
type ExtrasJobResultsUpdateParams struct {

	// Data.
	Data *models.WritableJobResult

	/* ID.

	   A UUID string identifying this job result.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras job results update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasJobResultsUpdateParams) WithDefaults() *ExtrasJobResultsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras job results update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasJobResultsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras job results update params
func (o *ExtrasJobResultsUpdateParams) WithTimeout(timeout time.Duration) *ExtrasJobResultsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras job results update params
func (o *ExtrasJobResultsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras job results update params
func (o *ExtrasJobResultsUpdateParams) WithContext(ctx context.Context) *ExtrasJobResultsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras job results update params
func (o *ExtrasJobResultsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras job results update params
func (o *ExtrasJobResultsUpdateParams) WithHTTPClient(client *http.Client) *ExtrasJobResultsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras job results update params
func (o *ExtrasJobResultsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras job results update params
func (o *ExtrasJobResultsUpdateParams) WithData(data *models.WritableJobResult) *ExtrasJobResultsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras job results update params
func (o *ExtrasJobResultsUpdateParams) SetData(data *models.WritableJobResult) {
	o.Data = data
}

// WithID adds the id to the extras job results update params
func (o *ExtrasJobResultsUpdateParams) WithID(id strfmt.UUID) *ExtrasJobResultsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras job results update params
func (o *ExtrasJobResultsUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasJobResultsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

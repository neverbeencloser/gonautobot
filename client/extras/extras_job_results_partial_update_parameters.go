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

// NewExtrasJobResultsPartialUpdateParams creates a new ExtrasJobResultsPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasJobResultsPartialUpdateParams() *ExtrasJobResultsPartialUpdateParams {
	return &ExtrasJobResultsPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasJobResultsPartialUpdateParamsWithTimeout creates a new ExtrasJobResultsPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasJobResultsPartialUpdateParamsWithTimeout(timeout time.Duration) *ExtrasJobResultsPartialUpdateParams {
	return &ExtrasJobResultsPartialUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasJobResultsPartialUpdateParamsWithContext creates a new ExtrasJobResultsPartialUpdateParams object
// with the ability to set a context for a request.
func NewExtrasJobResultsPartialUpdateParamsWithContext(ctx context.Context) *ExtrasJobResultsPartialUpdateParams {
	return &ExtrasJobResultsPartialUpdateParams{
		Context: ctx,
	}
}

// NewExtrasJobResultsPartialUpdateParamsWithHTTPClient creates a new ExtrasJobResultsPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasJobResultsPartialUpdateParamsWithHTTPClient(client *http.Client) *ExtrasJobResultsPartialUpdateParams {
	return &ExtrasJobResultsPartialUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasJobResultsPartialUpdateParams contains all the parameters to send to the API endpoint
   for the extras job results partial update operation.

   Typically these are written to a http.Request.
*/
type ExtrasJobResultsPartialUpdateParams struct {

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

// WithDefaults hydrates default values in the extras job results partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasJobResultsPartialUpdateParams) WithDefaults() *ExtrasJobResultsPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras job results partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasJobResultsPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras job results partial update params
func (o *ExtrasJobResultsPartialUpdateParams) WithTimeout(timeout time.Duration) *ExtrasJobResultsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras job results partial update params
func (o *ExtrasJobResultsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras job results partial update params
func (o *ExtrasJobResultsPartialUpdateParams) WithContext(ctx context.Context) *ExtrasJobResultsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras job results partial update params
func (o *ExtrasJobResultsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras job results partial update params
func (o *ExtrasJobResultsPartialUpdateParams) WithHTTPClient(client *http.Client) *ExtrasJobResultsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras job results partial update params
func (o *ExtrasJobResultsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras job results partial update params
func (o *ExtrasJobResultsPartialUpdateParams) WithData(data *models.WritableJobResult) *ExtrasJobResultsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras job results partial update params
func (o *ExtrasJobResultsPartialUpdateParams) SetData(data *models.WritableJobResult) {
	o.Data = data
}

// WithID adds the id to the extras job results partial update params
func (o *ExtrasJobResultsPartialUpdateParams) WithID(id strfmt.UUID) *ExtrasJobResultsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras job results partial update params
func (o *ExtrasJobResultsPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasJobResultsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

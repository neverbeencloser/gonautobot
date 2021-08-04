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

// NewExtrasConfigContextsPartialUpdateParams creates a new ExtrasConfigContextsPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasConfigContextsPartialUpdateParams() *ExtrasConfigContextsPartialUpdateParams {
	return &ExtrasConfigContextsPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasConfigContextsPartialUpdateParamsWithTimeout creates a new ExtrasConfigContextsPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasConfigContextsPartialUpdateParamsWithTimeout(timeout time.Duration) *ExtrasConfigContextsPartialUpdateParams {
	return &ExtrasConfigContextsPartialUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasConfigContextsPartialUpdateParamsWithContext creates a new ExtrasConfigContextsPartialUpdateParams object
// with the ability to set a context for a request.
func NewExtrasConfigContextsPartialUpdateParamsWithContext(ctx context.Context) *ExtrasConfigContextsPartialUpdateParams {
	return &ExtrasConfigContextsPartialUpdateParams{
		Context: ctx,
	}
}

// NewExtrasConfigContextsPartialUpdateParamsWithHTTPClient creates a new ExtrasConfigContextsPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasConfigContextsPartialUpdateParamsWithHTTPClient(client *http.Client) *ExtrasConfigContextsPartialUpdateParams {
	return &ExtrasConfigContextsPartialUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasConfigContextsPartialUpdateParams contains all the parameters to send to the API endpoint
   for the extras config contexts partial update operation.

   Typically these are written to a http.Request.
*/
type ExtrasConfigContextsPartialUpdateParams struct {

	// Data.
	Data *models.WritableConfigContext

	/* ID.

	   A UUID string identifying this config context.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras config contexts partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasConfigContextsPartialUpdateParams) WithDefaults() *ExtrasConfigContextsPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras config contexts partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasConfigContextsPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras config contexts partial update params
func (o *ExtrasConfigContextsPartialUpdateParams) WithTimeout(timeout time.Duration) *ExtrasConfigContextsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras config contexts partial update params
func (o *ExtrasConfigContextsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras config contexts partial update params
func (o *ExtrasConfigContextsPartialUpdateParams) WithContext(ctx context.Context) *ExtrasConfigContextsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras config contexts partial update params
func (o *ExtrasConfigContextsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras config contexts partial update params
func (o *ExtrasConfigContextsPartialUpdateParams) WithHTTPClient(client *http.Client) *ExtrasConfigContextsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras config contexts partial update params
func (o *ExtrasConfigContextsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras config contexts partial update params
func (o *ExtrasConfigContextsPartialUpdateParams) WithData(data *models.WritableConfigContext) *ExtrasConfigContextsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras config contexts partial update params
func (o *ExtrasConfigContextsPartialUpdateParams) SetData(data *models.WritableConfigContext) {
	o.Data = data
}

// WithID adds the id to the extras config contexts partial update params
func (o *ExtrasConfigContextsPartialUpdateParams) WithID(id strfmt.UUID) *ExtrasConfigContextsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras config contexts partial update params
func (o *ExtrasConfigContextsPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasConfigContextsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

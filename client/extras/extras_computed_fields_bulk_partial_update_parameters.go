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

// NewExtrasComputedFieldsBulkPartialUpdateParams creates a new ExtrasComputedFieldsBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasComputedFieldsBulkPartialUpdateParams() *ExtrasComputedFieldsBulkPartialUpdateParams {
	return &ExtrasComputedFieldsBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasComputedFieldsBulkPartialUpdateParamsWithTimeout creates a new ExtrasComputedFieldsBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasComputedFieldsBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *ExtrasComputedFieldsBulkPartialUpdateParams {
	return &ExtrasComputedFieldsBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasComputedFieldsBulkPartialUpdateParamsWithContext creates a new ExtrasComputedFieldsBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewExtrasComputedFieldsBulkPartialUpdateParamsWithContext(ctx context.Context) *ExtrasComputedFieldsBulkPartialUpdateParams {
	return &ExtrasComputedFieldsBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewExtrasComputedFieldsBulkPartialUpdateParamsWithHTTPClient creates a new ExtrasComputedFieldsBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasComputedFieldsBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *ExtrasComputedFieldsBulkPartialUpdateParams {
	return &ExtrasComputedFieldsBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasComputedFieldsBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the extras computed fields bulk partial update operation.

   Typically these are written to a http.Request.
*/
type ExtrasComputedFieldsBulkPartialUpdateParams struct {

	// Data.
	Data *models.ComputedField

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras computed fields bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasComputedFieldsBulkPartialUpdateParams) WithDefaults() *ExtrasComputedFieldsBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras computed fields bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasComputedFieldsBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras computed fields bulk partial update params
func (o *ExtrasComputedFieldsBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *ExtrasComputedFieldsBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras computed fields bulk partial update params
func (o *ExtrasComputedFieldsBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras computed fields bulk partial update params
func (o *ExtrasComputedFieldsBulkPartialUpdateParams) WithContext(ctx context.Context) *ExtrasComputedFieldsBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras computed fields bulk partial update params
func (o *ExtrasComputedFieldsBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras computed fields bulk partial update params
func (o *ExtrasComputedFieldsBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *ExtrasComputedFieldsBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras computed fields bulk partial update params
func (o *ExtrasComputedFieldsBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras computed fields bulk partial update params
func (o *ExtrasComputedFieldsBulkPartialUpdateParams) WithData(data *models.ComputedField) *ExtrasComputedFieldsBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras computed fields bulk partial update params
func (o *ExtrasComputedFieldsBulkPartialUpdateParams) SetData(data *models.ComputedField) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasComputedFieldsBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

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

// NewExtrasComputedFieldsBulkUpdateParams creates a new ExtrasComputedFieldsBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasComputedFieldsBulkUpdateParams() *ExtrasComputedFieldsBulkUpdateParams {
	return &ExtrasComputedFieldsBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasComputedFieldsBulkUpdateParamsWithTimeout creates a new ExtrasComputedFieldsBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasComputedFieldsBulkUpdateParamsWithTimeout(timeout time.Duration) *ExtrasComputedFieldsBulkUpdateParams {
	return &ExtrasComputedFieldsBulkUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasComputedFieldsBulkUpdateParamsWithContext creates a new ExtrasComputedFieldsBulkUpdateParams object
// with the ability to set a context for a request.
func NewExtrasComputedFieldsBulkUpdateParamsWithContext(ctx context.Context) *ExtrasComputedFieldsBulkUpdateParams {
	return &ExtrasComputedFieldsBulkUpdateParams{
		Context: ctx,
	}
}

// NewExtrasComputedFieldsBulkUpdateParamsWithHTTPClient creates a new ExtrasComputedFieldsBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasComputedFieldsBulkUpdateParamsWithHTTPClient(client *http.Client) *ExtrasComputedFieldsBulkUpdateParams {
	return &ExtrasComputedFieldsBulkUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasComputedFieldsBulkUpdateParams contains all the parameters to send to the API endpoint
   for the extras computed fields bulk update operation.

   Typically these are written to a http.Request.
*/
type ExtrasComputedFieldsBulkUpdateParams struct {

	// Data.
	Data *models.ComputedField

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras computed fields bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasComputedFieldsBulkUpdateParams) WithDefaults() *ExtrasComputedFieldsBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras computed fields bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasComputedFieldsBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras computed fields bulk update params
func (o *ExtrasComputedFieldsBulkUpdateParams) WithTimeout(timeout time.Duration) *ExtrasComputedFieldsBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras computed fields bulk update params
func (o *ExtrasComputedFieldsBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras computed fields bulk update params
func (o *ExtrasComputedFieldsBulkUpdateParams) WithContext(ctx context.Context) *ExtrasComputedFieldsBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras computed fields bulk update params
func (o *ExtrasComputedFieldsBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras computed fields bulk update params
func (o *ExtrasComputedFieldsBulkUpdateParams) WithHTTPClient(client *http.Client) *ExtrasComputedFieldsBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras computed fields bulk update params
func (o *ExtrasComputedFieldsBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras computed fields bulk update params
func (o *ExtrasComputedFieldsBulkUpdateParams) WithData(data *models.ComputedField) *ExtrasComputedFieldsBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras computed fields bulk update params
func (o *ExtrasComputedFieldsBulkUpdateParams) SetData(data *models.ComputedField) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasComputedFieldsBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

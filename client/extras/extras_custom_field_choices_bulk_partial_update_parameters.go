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

// NewExtrasCustomFieldChoicesBulkPartialUpdateParams creates a new ExtrasCustomFieldChoicesBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasCustomFieldChoicesBulkPartialUpdateParams() *ExtrasCustomFieldChoicesBulkPartialUpdateParams {
	return &ExtrasCustomFieldChoicesBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasCustomFieldChoicesBulkPartialUpdateParamsWithTimeout creates a new ExtrasCustomFieldChoicesBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasCustomFieldChoicesBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *ExtrasCustomFieldChoicesBulkPartialUpdateParams {
	return &ExtrasCustomFieldChoicesBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasCustomFieldChoicesBulkPartialUpdateParamsWithContext creates a new ExtrasCustomFieldChoicesBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewExtrasCustomFieldChoicesBulkPartialUpdateParamsWithContext(ctx context.Context) *ExtrasCustomFieldChoicesBulkPartialUpdateParams {
	return &ExtrasCustomFieldChoicesBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewExtrasCustomFieldChoicesBulkPartialUpdateParamsWithHTTPClient creates a new ExtrasCustomFieldChoicesBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasCustomFieldChoicesBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *ExtrasCustomFieldChoicesBulkPartialUpdateParams {
	return &ExtrasCustomFieldChoicesBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasCustomFieldChoicesBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the extras custom field choices bulk partial update operation.

   Typically these are written to a http.Request.
*/
type ExtrasCustomFieldChoicesBulkPartialUpdateParams struct {

	// Data.
	Data *models.WritableCustomFieldChoice

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras custom field choices bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasCustomFieldChoicesBulkPartialUpdateParams) WithDefaults() *ExtrasCustomFieldChoicesBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras custom field choices bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasCustomFieldChoicesBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras custom field choices bulk partial update params
func (o *ExtrasCustomFieldChoicesBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *ExtrasCustomFieldChoicesBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras custom field choices bulk partial update params
func (o *ExtrasCustomFieldChoicesBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras custom field choices bulk partial update params
func (o *ExtrasCustomFieldChoicesBulkPartialUpdateParams) WithContext(ctx context.Context) *ExtrasCustomFieldChoicesBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras custom field choices bulk partial update params
func (o *ExtrasCustomFieldChoicesBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras custom field choices bulk partial update params
func (o *ExtrasCustomFieldChoicesBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *ExtrasCustomFieldChoicesBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras custom field choices bulk partial update params
func (o *ExtrasCustomFieldChoicesBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras custom field choices bulk partial update params
func (o *ExtrasCustomFieldChoicesBulkPartialUpdateParams) WithData(data *models.WritableCustomFieldChoice) *ExtrasCustomFieldChoicesBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras custom field choices bulk partial update params
func (o *ExtrasCustomFieldChoicesBulkPartialUpdateParams) SetData(data *models.WritableCustomFieldChoice) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasCustomFieldChoicesBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

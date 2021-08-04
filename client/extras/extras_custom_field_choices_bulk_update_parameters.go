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

// NewExtrasCustomFieldChoicesBulkUpdateParams creates a new ExtrasCustomFieldChoicesBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasCustomFieldChoicesBulkUpdateParams() *ExtrasCustomFieldChoicesBulkUpdateParams {
	return &ExtrasCustomFieldChoicesBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasCustomFieldChoicesBulkUpdateParamsWithTimeout creates a new ExtrasCustomFieldChoicesBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasCustomFieldChoicesBulkUpdateParamsWithTimeout(timeout time.Duration) *ExtrasCustomFieldChoicesBulkUpdateParams {
	return &ExtrasCustomFieldChoicesBulkUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasCustomFieldChoicesBulkUpdateParamsWithContext creates a new ExtrasCustomFieldChoicesBulkUpdateParams object
// with the ability to set a context for a request.
func NewExtrasCustomFieldChoicesBulkUpdateParamsWithContext(ctx context.Context) *ExtrasCustomFieldChoicesBulkUpdateParams {
	return &ExtrasCustomFieldChoicesBulkUpdateParams{
		Context: ctx,
	}
}

// NewExtrasCustomFieldChoicesBulkUpdateParamsWithHTTPClient creates a new ExtrasCustomFieldChoicesBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasCustomFieldChoicesBulkUpdateParamsWithHTTPClient(client *http.Client) *ExtrasCustomFieldChoicesBulkUpdateParams {
	return &ExtrasCustomFieldChoicesBulkUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasCustomFieldChoicesBulkUpdateParams contains all the parameters to send to the API endpoint
   for the extras custom field choices bulk update operation.

   Typically these are written to a http.Request.
*/
type ExtrasCustomFieldChoicesBulkUpdateParams struct {

	// Data.
	Data *models.WritableCustomFieldChoice

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras custom field choices bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasCustomFieldChoicesBulkUpdateParams) WithDefaults() *ExtrasCustomFieldChoicesBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras custom field choices bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasCustomFieldChoicesBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras custom field choices bulk update params
func (o *ExtrasCustomFieldChoicesBulkUpdateParams) WithTimeout(timeout time.Duration) *ExtrasCustomFieldChoicesBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras custom field choices bulk update params
func (o *ExtrasCustomFieldChoicesBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras custom field choices bulk update params
func (o *ExtrasCustomFieldChoicesBulkUpdateParams) WithContext(ctx context.Context) *ExtrasCustomFieldChoicesBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras custom field choices bulk update params
func (o *ExtrasCustomFieldChoicesBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras custom field choices bulk update params
func (o *ExtrasCustomFieldChoicesBulkUpdateParams) WithHTTPClient(client *http.Client) *ExtrasCustomFieldChoicesBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras custom field choices bulk update params
func (o *ExtrasCustomFieldChoicesBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras custom field choices bulk update params
func (o *ExtrasCustomFieldChoicesBulkUpdateParams) WithData(data *models.WritableCustomFieldChoice) *ExtrasCustomFieldChoicesBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras custom field choices bulk update params
func (o *ExtrasCustomFieldChoicesBulkUpdateParams) SetData(data *models.WritableCustomFieldChoice) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasCustomFieldChoicesBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

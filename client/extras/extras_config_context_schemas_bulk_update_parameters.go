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

// NewExtrasConfigContextSchemasBulkUpdateParams creates a new ExtrasConfigContextSchemasBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasConfigContextSchemasBulkUpdateParams() *ExtrasConfigContextSchemasBulkUpdateParams {
	return &ExtrasConfigContextSchemasBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasConfigContextSchemasBulkUpdateParamsWithTimeout creates a new ExtrasConfigContextSchemasBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasConfigContextSchemasBulkUpdateParamsWithTimeout(timeout time.Duration) *ExtrasConfigContextSchemasBulkUpdateParams {
	return &ExtrasConfigContextSchemasBulkUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasConfigContextSchemasBulkUpdateParamsWithContext creates a new ExtrasConfigContextSchemasBulkUpdateParams object
// with the ability to set a context for a request.
func NewExtrasConfigContextSchemasBulkUpdateParamsWithContext(ctx context.Context) *ExtrasConfigContextSchemasBulkUpdateParams {
	return &ExtrasConfigContextSchemasBulkUpdateParams{
		Context: ctx,
	}
}

// NewExtrasConfigContextSchemasBulkUpdateParamsWithHTTPClient creates a new ExtrasConfigContextSchemasBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasConfigContextSchemasBulkUpdateParamsWithHTTPClient(client *http.Client) *ExtrasConfigContextSchemasBulkUpdateParams {
	return &ExtrasConfigContextSchemasBulkUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasConfigContextSchemasBulkUpdateParams contains all the parameters to send to the API endpoint
   for the extras config context schemas bulk update operation.

   Typically these are written to a http.Request.
*/
type ExtrasConfigContextSchemasBulkUpdateParams struct {

	// Data.
	Data *models.ConfigContextSchema

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras config context schemas bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasConfigContextSchemasBulkUpdateParams) WithDefaults() *ExtrasConfigContextSchemasBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras config context schemas bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasConfigContextSchemasBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras config context schemas bulk update params
func (o *ExtrasConfigContextSchemasBulkUpdateParams) WithTimeout(timeout time.Duration) *ExtrasConfigContextSchemasBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras config context schemas bulk update params
func (o *ExtrasConfigContextSchemasBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras config context schemas bulk update params
func (o *ExtrasConfigContextSchemasBulkUpdateParams) WithContext(ctx context.Context) *ExtrasConfigContextSchemasBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras config context schemas bulk update params
func (o *ExtrasConfigContextSchemasBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras config context schemas bulk update params
func (o *ExtrasConfigContextSchemasBulkUpdateParams) WithHTTPClient(client *http.Client) *ExtrasConfigContextSchemasBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras config context schemas bulk update params
func (o *ExtrasConfigContextSchemasBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras config context schemas bulk update params
func (o *ExtrasConfigContextSchemasBulkUpdateParams) WithData(data *models.ConfigContextSchema) *ExtrasConfigContextSchemasBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras config context schemas bulk update params
func (o *ExtrasConfigContextSchemasBulkUpdateParams) SetData(data *models.ConfigContextSchema) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasConfigContextSchemasBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

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

// NewExtrasConfigContextSchemasBulkPartialUpdateParams creates a new ExtrasConfigContextSchemasBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasConfigContextSchemasBulkPartialUpdateParams() *ExtrasConfigContextSchemasBulkPartialUpdateParams {
	return &ExtrasConfigContextSchemasBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasConfigContextSchemasBulkPartialUpdateParamsWithTimeout creates a new ExtrasConfigContextSchemasBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasConfigContextSchemasBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *ExtrasConfigContextSchemasBulkPartialUpdateParams {
	return &ExtrasConfigContextSchemasBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasConfigContextSchemasBulkPartialUpdateParamsWithContext creates a new ExtrasConfigContextSchemasBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewExtrasConfigContextSchemasBulkPartialUpdateParamsWithContext(ctx context.Context) *ExtrasConfigContextSchemasBulkPartialUpdateParams {
	return &ExtrasConfigContextSchemasBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewExtrasConfigContextSchemasBulkPartialUpdateParamsWithHTTPClient creates a new ExtrasConfigContextSchemasBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasConfigContextSchemasBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *ExtrasConfigContextSchemasBulkPartialUpdateParams {
	return &ExtrasConfigContextSchemasBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasConfigContextSchemasBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the extras config context schemas bulk partial update operation.

   Typically these are written to a http.Request.
*/
type ExtrasConfigContextSchemasBulkPartialUpdateParams struct {

	// Data.
	Data *models.ConfigContextSchema

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras config context schemas bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasConfigContextSchemasBulkPartialUpdateParams) WithDefaults() *ExtrasConfigContextSchemasBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras config context schemas bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasConfigContextSchemasBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras config context schemas bulk partial update params
func (o *ExtrasConfigContextSchemasBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *ExtrasConfigContextSchemasBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras config context schemas bulk partial update params
func (o *ExtrasConfigContextSchemasBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras config context schemas bulk partial update params
func (o *ExtrasConfigContextSchemasBulkPartialUpdateParams) WithContext(ctx context.Context) *ExtrasConfigContextSchemasBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras config context schemas bulk partial update params
func (o *ExtrasConfigContextSchemasBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras config context schemas bulk partial update params
func (o *ExtrasConfigContextSchemasBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *ExtrasConfigContextSchemasBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras config context schemas bulk partial update params
func (o *ExtrasConfigContextSchemasBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras config context schemas bulk partial update params
func (o *ExtrasConfigContextSchemasBulkPartialUpdateParams) WithData(data *models.ConfigContextSchema) *ExtrasConfigContextSchemasBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras config context schemas bulk partial update params
func (o *ExtrasConfigContextSchemasBulkPartialUpdateParams) SetData(data *models.ConfigContextSchema) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasConfigContextSchemasBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

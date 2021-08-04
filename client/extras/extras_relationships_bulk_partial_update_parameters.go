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

// NewExtrasRelationshipsBulkPartialUpdateParams creates a new ExtrasRelationshipsBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasRelationshipsBulkPartialUpdateParams() *ExtrasRelationshipsBulkPartialUpdateParams {
	return &ExtrasRelationshipsBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasRelationshipsBulkPartialUpdateParamsWithTimeout creates a new ExtrasRelationshipsBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasRelationshipsBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *ExtrasRelationshipsBulkPartialUpdateParams {
	return &ExtrasRelationshipsBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasRelationshipsBulkPartialUpdateParamsWithContext creates a new ExtrasRelationshipsBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewExtrasRelationshipsBulkPartialUpdateParamsWithContext(ctx context.Context) *ExtrasRelationshipsBulkPartialUpdateParams {
	return &ExtrasRelationshipsBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewExtrasRelationshipsBulkPartialUpdateParamsWithHTTPClient creates a new ExtrasRelationshipsBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasRelationshipsBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *ExtrasRelationshipsBulkPartialUpdateParams {
	return &ExtrasRelationshipsBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasRelationshipsBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the extras relationships bulk partial update operation.

   Typically these are written to a http.Request.
*/
type ExtrasRelationshipsBulkPartialUpdateParams struct {

	// Data.
	Data *models.Relationship

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras relationships bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasRelationshipsBulkPartialUpdateParams) WithDefaults() *ExtrasRelationshipsBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras relationships bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasRelationshipsBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras relationships bulk partial update params
func (o *ExtrasRelationshipsBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *ExtrasRelationshipsBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras relationships bulk partial update params
func (o *ExtrasRelationshipsBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras relationships bulk partial update params
func (o *ExtrasRelationshipsBulkPartialUpdateParams) WithContext(ctx context.Context) *ExtrasRelationshipsBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras relationships bulk partial update params
func (o *ExtrasRelationshipsBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras relationships bulk partial update params
func (o *ExtrasRelationshipsBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *ExtrasRelationshipsBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras relationships bulk partial update params
func (o *ExtrasRelationshipsBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras relationships bulk partial update params
func (o *ExtrasRelationshipsBulkPartialUpdateParams) WithData(data *models.Relationship) *ExtrasRelationshipsBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras relationships bulk partial update params
func (o *ExtrasRelationshipsBulkPartialUpdateParams) SetData(data *models.Relationship) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasRelationshipsBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

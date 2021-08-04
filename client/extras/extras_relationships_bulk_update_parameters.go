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

// NewExtrasRelationshipsBulkUpdateParams creates a new ExtrasRelationshipsBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasRelationshipsBulkUpdateParams() *ExtrasRelationshipsBulkUpdateParams {
	return &ExtrasRelationshipsBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasRelationshipsBulkUpdateParamsWithTimeout creates a new ExtrasRelationshipsBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasRelationshipsBulkUpdateParamsWithTimeout(timeout time.Duration) *ExtrasRelationshipsBulkUpdateParams {
	return &ExtrasRelationshipsBulkUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasRelationshipsBulkUpdateParamsWithContext creates a new ExtrasRelationshipsBulkUpdateParams object
// with the ability to set a context for a request.
func NewExtrasRelationshipsBulkUpdateParamsWithContext(ctx context.Context) *ExtrasRelationshipsBulkUpdateParams {
	return &ExtrasRelationshipsBulkUpdateParams{
		Context: ctx,
	}
}

// NewExtrasRelationshipsBulkUpdateParamsWithHTTPClient creates a new ExtrasRelationshipsBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasRelationshipsBulkUpdateParamsWithHTTPClient(client *http.Client) *ExtrasRelationshipsBulkUpdateParams {
	return &ExtrasRelationshipsBulkUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasRelationshipsBulkUpdateParams contains all the parameters to send to the API endpoint
   for the extras relationships bulk update operation.

   Typically these are written to a http.Request.
*/
type ExtrasRelationshipsBulkUpdateParams struct {

	// Data.
	Data *models.Relationship

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras relationships bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasRelationshipsBulkUpdateParams) WithDefaults() *ExtrasRelationshipsBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras relationships bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasRelationshipsBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras relationships bulk update params
func (o *ExtrasRelationshipsBulkUpdateParams) WithTimeout(timeout time.Duration) *ExtrasRelationshipsBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras relationships bulk update params
func (o *ExtrasRelationshipsBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras relationships bulk update params
func (o *ExtrasRelationshipsBulkUpdateParams) WithContext(ctx context.Context) *ExtrasRelationshipsBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras relationships bulk update params
func (o *ExtrasRelationshipsBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras relationships bulk update params
func (o *ExtrasRelationshipsBulkUpdateParams) WithHTTPClient(client *http.Client) *ExtrasRelationshipsBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras relationships bulk update params
func (o *ExtrasRelationshipsBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras relationships bulk update params
func (o *ExtrasRelationshipsBulkUpdateParams) WithData(data *models.Relationship) *ExtrasRelationshipsBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras relationships bulk update params
func (o *ExtrasRelationshipsBulkUpdateParams) SetData(data *models.Relationship) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasRelationshipsBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

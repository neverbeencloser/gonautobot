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

// NewExtrasRelationshipsUpdateParams creates a new ExtrasRelationshipsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasRelationshipsUpdateParams() *ExtrasRelationshipsUpdateParams {
	return &ExtrasRelationshipsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasRelationshipsUpdateParamsWithTimeout creates a new ExtrasRelationshipsUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasRelationshipsUpdateParamsWithTimeout(timeout time.Duration) *ExtrasRelationshipsUpdateParams {
	return &ExtrasRelationshipsUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasRelationshipsUpdateParamsWithContext creates a new ExtrasRelationshipsUpdateParams object
// with the ability to set a context for a request.
func NewExtrasRelationshipsUpdateParamsWithContext(ctx context.Context) *ExtrasRelationshipsUpdateParams {
	return &ExtrasRelationshipsUpdateParams{
		Context: ctx,
	}
}

// NewExtrasRelationshipsUpdateParamsWithHTTPClient creates a new ExtrasRelationshipsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasRelationshipsUpdateParamsWithHTTPClient(client *http.Client) *ExtrasRelationshipsUpdateParams {
	return &ExtrasRelationshipsUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasRelationshipsUpdateParams contains all the parameters to send to the API endpoint
   for the extras relationships update operation.

   Typically these are written to a http.Request.
*/
type ExtrasRelationshipsUpdateParams struct {

	// Data.
	Data *models.Relationship

	/* ID.

	   A UUID string identifying this relationship.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras relationships update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasRelationshipsUpdateParams) WithDefaults() *ExtrasRelationshipsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras relationships update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasRelationshipsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras relationships update params
func (o *ExtrasRelationshipsUpdateParams) WithTimeout(timeout time.Duration) *ExtrasRelationshipsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras relationships update params
func (o *ExtrasRelationshipsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras relationships update params
func (o *ExtrasRelationshipsUpdateParams) WithContext(ctx context.Context) *ExtrasRelationshipsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras relationships update params
func (o *ExtrasRelationshipsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras relationships update params
func (o *ExtrasRelationshipsUpdateParams) WithHTTPClient(client *http.Client) *ExtrasRelationshipsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras relationships update params
func (o *ExtrasRelationshipsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras relationships update params
func (o *ExtrasRelationshipsUpdateParams) WithData(data *models.Relationship) *ExtrasRelationshipsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras relationships update params
func (o *ExtrasRelationshipsUpdateParams) SetData(data *models.Relationship) {
	o.Data = data
}

// WithID adds the id to the extras relationships update params
func (o *ExtrasRelationshipsUpdateParams) WithID(id strfmt.UUID) *ExtrasRelationshipsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras relationships update params
func (o *ExtrasRelationshipsUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasRelationshipsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

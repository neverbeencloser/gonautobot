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

// NewExtrasRelationshipsPartialUpdateParams creates a new ExtrasRelationshipsPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasRelationshipsPartialUpdateParams() *ExtrasRelationshipsPartialUpdateParams {
	return &ExtrasRelationshipsPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasRelationshipsPartialUpdateParamsWithTimeout creates a new ExtrasRelationshipsPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasRelationshipsPartialUpdateParamsWithTimeout(timeout time.Duration) *ExtrasRelationshipsPartialUpdateParams {
	return &ExtrasRelationshipsPartialUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasRelationshipsPartialUpdateParamsWithContext creates a new ExtrasRelationshipsPartialUpdateParams object
// with the ability to set a context for a request.
func NewExtrasRelationshipsPartialUpdateParamsWithContext(ctx context.Context) *ExtrasRelationshipsPartialUpdateParams {
	return &ExtrasRelationshipsPartialUpdateParams{
		Context: ctx,
	}
}

// NewExtrasRelationshipsPartialUpdateParamsWithHTTPClient creates a new ExtrasRelationshipsPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasRelationshipsPartialUpdateParamsWithHTTPClient(client *http.Client) *ExtrasRelationshipsPartialUpdateParams {
	return &ExtrasRelationshipsPartialUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasRelationshipsPartialUpdateParams contains all the parameters to send to the API endpoint
   for the extras relationships partial update operation.

   Typically these are written to a http.Request.
*/
type ExtrasRelationshipsPartialUpdateParams struct {

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

// WithDefaults hydrates default values in the extras relationships partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasRelationshipsPartialUpdateParams) WithDefaults() *ExtrasRelationshipsPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras relationships partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasRelationshipsPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras relationships partial update params
func (o *ExtrasRelationshipsPartialUpdateParams) WithTimeout(timeout time.Duration) *ExtrasRelationshipsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras relationships partial update params
func (o *ExtrasRelationshipsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras relationships partial update params
func (o *ExtrasRelationshipsPartialUpdateParams) WithContext(ctx context.Context) *ExtrasRelationshipsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras relationships partial update params
func (o *ExtrasRelationshipsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras relationships partial update params
func (o *ExtrasRelationshipsPartialUpdateParams) WithHTTPClient(client *http.Client) *ExtrasRelationshipsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras relationships partial update params
func (o *ExtrasRelationshipsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras relationships partial update params
func (o *ExtrasRelationshipsPartialUpdateParams) WithData(data *models.Relationship) *ExtrasRelationshipsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras relationships partial update params
func (o *ExtrasRelationshipsPartialUpdateParams) SetData(data *models.Relationship) {
	o.Data = data
}

// WithID adds the id to the extras relationships partial update params
func (o *ExtrasRelationshipsPartialUpdateParams) WithID(id strfmt.UUID) *ExtrasRelationshipsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras relationships partial update params
func (o *ExtrasRelationshipsPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasRelationshipsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

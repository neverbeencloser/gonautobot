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

// NewExtrasRelationshipAssociationsPartialUpdateParams creates a new ExtrasRelationshipAssociationsPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasRelationshipAssociationsPartialUpdateParams() *ExtrasRelationshipAssociationsPartialUpdateParams {
	return &ExtrasRelationshipAssociationsPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasRelationshipAssociationsPartialUpdateParamsWithTimeout creates a new ExtrasRelationshipAssociationsPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasRelationshipAssociationsPartialUpdateParamsWithTimeout(timeout time.Duration) *ExtrasRelationshipAssociationsPartialUpdateParams {
	return &ExtrasRelationshipAssociationsPartialUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasRelationshipAssociationsPartialUpdateParamsWithContext creates a new ExtrasRelationshipAssociationsPartialUpdateParams object
// with the ability to set a context for a request.
func NewExtrasRelationshipAssociationsPartialUpdateParamsWithContext(ctx context.Context) *ExtrasRelationshipAssociationsPartialUpdateParams {
	return &ExtrasRelationshipAssociationsPartialUpdateParams{
		Context: ctx,
	}
}

// NewExtrasRelationshipAssociationsPartialUpdateParamsWithHTTPClient creates a new ExtrasRelationshipAssociationsPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasRelationshipAssociationsPartialUpdateParamsWithHTTPClient(client *http.Client) *ExtrasRelationshipAssociationsPartialUpdateParams {
	return &ExtrasRelationshipAssociationsPartialUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasRelationshipAssociationsPartialUpdateParams contains all the parameters to send to the API endpoint
   for the extras relationship associations partial update operation.

   Typically these are written to a http.Request.
*/
type ExtrasRelationshipAssociationsPartialUpdateParams struct {

	// Data.
	Data *models.WritableRelationshipAssociation

	/* ID.

	   A UUID string identifying this relationship association.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras relationship associations partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasRelationshipAssociationsPartialUpdateParams) WithDefaults() *ExtrasRelationshipAssociationsPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras relationship associations partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasRelationshipAssociationsPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras relationship associations partial update params
func (o *ExtrasRelationshipAssociationsPartialUpdateParams) WithTimeout(timeout time.Duration) *ExtrasRelationshipAssociationsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras relationship associations partial update params
func (o *ExtrasRelationshipAssociationsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras relationship associations partial update params
func (o *ExtrasRelationshipAssociationsPartialUpdateParams) WithContext(ctx context.Context) *ExtrasRelationshipAssociationsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras relationship associations partial update params
func (o *ExtrasRelationshipAssociationsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras relationship associations partial update params
func (o *ExtrasRelationshipAssociationsPartialUpdateParams) WithHTTPClient(client *http.Client) *ExtrasRelationshipAssociationsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras relationship associations partial update params
func (o *ExtrasRelationshipAssociationsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras relationship associations partial update params
func (o *ExtrasRelationshipAssociationsPartialUpdateParams) WithData(data *models.WritableRelationshipAssociation) *ExtrasRelationshipAssociationsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras relationship associations partial update params
func (o *ExtrasRelationshipAssociationsPartialUpdateParams) SetData(data *models.WritableRelationshipAssociation) {
	o.Data = data
}

// WithID adds the id to the extras relationship associations partial update params
func (o *ExtrasRelationshipAssociationsPartialUpdateParams) WithID(id strfmt.UUID) *ExtrasRelationshipAssociationsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras relationship associations partial update params
func (o *ExtrasRelationshipAssociationsPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasRelationshipAssociationsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

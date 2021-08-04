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

// NewExtrasRelationshipAssociationsUpdateParams creates a new ExtrasRelationshipAssociationsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasRelationshipAssociationsUpdateParams() *ExtrasRelationshipAssociationsUpdateParams {
	return &ExtrasRelationshipAssociationsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasRelationshipAssociationsUpdateParamsWithTimeout creates a new ExtrasRelationshipAssociationsUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasRelationshipAssociationsUpdateParamsWithTimeout(timeout time.Duration) *ExtrasRelationshipAssociationsUpdateParams {
	return &ExtrasRelationshipAssociationsUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasRelationshipAssociationsUpdateParamsWithContext creates a new ExtrasRelationshipAssociationsUpdateParams object
// with the ability to set a context for a request.
func NewExtrasRelationshipAssociationsUpdateParamsWithContext(ctx context.Context) *ExtrasRelationshipAssociationsUpdateParams {
	return &ExtrasRelationshipAssociationsUpdateParams{
		Context: ctx,
	}
}

// NewExtrasRelationshipAssociationsUpdateParamsWithHTTPClient creates a new ExtrasRelationshipAssociationsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasRelationshipAssociationsUpdateParamsWithHTTPClient(client *http.Client) *ExtrasRelationshipAssociationsUpdateParams {
	return &ExtrasRelationshipAssociationsUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasRelationshipAssociationsUpdateParams contains all the parameters to send to the API endpoint
   for the extras relationship associations update operation.

   Typically these are written to a http.Request.
*/
type ExtrasRelationshipAssociationsUpdateParams struct {

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

// WithDefaults hydrates default values in the extras relationship associations update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasRelationshipAssociationsUpdateParams) WithDefaults() *ExtrasRelationshipAssociationsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras relationship associations update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasRelationshipAssociationsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras relationship associations update params
func (o *ExtrasRelationshipAssociationsUpdateParams) WithTimeout(timeout time.Duration) *ExtrasRelationshipAssociationsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras relationship associations update params
func (o *ExtrasRelationshipAssociationsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras relationship associations update params
func (o *ExtrasRelationshipAssociationsUpdateParams) WithContext(ctx context.Context) *ExtrasRelationshipAssociationsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras relationship associations update params
func (o *ExtrasRelationshipAssociationsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras relationship associations update params
func (o *ExtrasRelationshipAssociationsUpdateParams) WithHTTPClient(client *http.Client) *ExtrasRelationshipAssociationsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras relationship associations update params
func (o *ExtrasRelationshipAssociationsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras relationship associations update params
func (o *ExtrasRelationshipAssociationsUpdateParams) WithData(data *models.WritableRelationshipAssociation) *ExtrasRelationshipAssociationsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras relationship associations update params
func (o *ExtrasRelationshipAssociationsUpdateParams) SetData(data *models.WritableRelationshipAssociation) {
	o.Data = data
}

// WithID adds the id to the extras relationship associations update params
func (o *ExtrasRelationshipAssociationsUpdateParams) WithID(id strfmt.UUID) *ExtrasRelationshipAssociationsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras relationship associations update params
func (o *ExtrasRelationshipAssociationsUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasRelationshipAssociationsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

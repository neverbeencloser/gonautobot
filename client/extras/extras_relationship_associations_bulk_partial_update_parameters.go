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

// NewExtrasRelationshipAssociationsBulkPartialUpdateParams creates a new ExtrasRelationshipAssociationsBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasRelationshipAssociationsBulkPartialUpdateParams() *ExtrasRelationshipAssociationsBulkPartialUpdateParams {
	return &ExtrasRelationshipAssociationsBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasRelationshipAssociationsBulkPartialUpdateParamsWithTimeout creates a new ExtrasRelationshipAssociationsBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasRelationshipAssociationsBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *ExtrasRelationshipAssociationsBulkPartialUpdateParams {
	return &ExtrasRelationshipAssociationsBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasRelationshipAssociationsBulkPartialUpdateParamsWithContext creates a new ExtrasRelationshipAssociationsBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewExtrasRelationshipAssociationsBulkPartialUpdateParamsWithContext(ctx context.Context) *ExtrasRelationshipAssociationsBulkPartialUpdateParams {
	return &ExtrasRelationshipAssociationsBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewExtrasRelationshipAssociationsBulkPartialUpdateParamsWithHTTPClient creates a new ExtrasRelationshipAssociationsBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasRelationshipAssociationsBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *ExtrasRelationshipAssociationsBulkPartialUpdateParams {
	return &ExtrasRelationshipAssociationsBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasRelationshipAssociationsBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the extras relationship associations bulk partial update operation.

   Typically these are written to a http.Request.
*/
type ExtrasRelationshipAssociationsBulkPartialUpdateParams struct {

	// Data.
	Data *models.WritableRelationshipAssociation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras relationship associations bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasRelationshipAssociationsBulkPartialUpdateParams) WithDefaults() *ExtrasRelationshipAssociationsBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras relationship associations bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasRelationshipAssociationsBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras relationship associations bulk partial update params
func (o *ExtrasRelationshipAssociationsBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *ExtrasRelationshipAssociationsBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras relationship associations bulk partial update params
func (o *ExtrasRelationshipAssociationsBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras relationship associations bulk partial update params
func (o *ExtrasRelationshipAssociationsBulkPartialUpdateParams) WithContext(ctx context.Context) *ExtrasRelationshipAssociationsBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras relationship associations bulk partial update params
func (o *ExtrasRelationshipAssociationsBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras relationship associations bulk partial update params
func (o *ExtrasRelationshipAssociationsBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *ExtrasRelationshipAssociationsBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras relationship associations bulk partial update params
func (o *ExtrasRelationshipAssociationsBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras relationship associations bulk partial update params
func (o *ExtrasRelationshipAssociationsBulkPartialUpdateParams) WithData(data *models.WritableRelationshipAssociation) *ExtrasRelationshipAssociationsBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras relationship associations bulk partial update params
func (o *ExtrasRelationshipAssociationsBulkPartialUpdateParams) SetData(data *models.WritableRelationshipAssociation) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasRelationshipAssociationsBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

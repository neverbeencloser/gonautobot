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

// NewExtrasRelationshipAssociationsBulkUpdateParams creates a new ExtrasRelationshipAssociationsBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasRelationshipAssociationsBulkUpdateParams() *ExtrasRelationshipAssociationsBulkUpdateParams {
	return &ExtrasRelationshipAssociationsBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasRelationshipAssociationsBulkUpdateParamsWithTimeout creates a new ExtrasRelationshipAssociationsBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasRelationshipAssociationsBulkUpdateParamsWithTimeout(timeout time.Duration) *ExtrasRelationshipAssociationsBulkUpdateParams {
	return &ExtrasRelationshipAssociationsBulkUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasRelationshipAssociationsBulkUpdateParamsWithContext creates a new ExtrasRelationshipAssociationsBulkUpdateParams object
// with the ability to set a context for a request.
func NewExtrasRelationshipAssociationsBulkUpdateParamsWithContext(ctx context.Context) *ExtrasRelationshipAssociationsBulkUpdateParams {
	return &ExtrasRelationshipAssociationsBulkUpdateParams{
		Context: ctx,
	}
}

// NewExtrasRelationshipAssociationsBulkUpdateParamsWithHTTPClient creates a new ExtrasRelationshipAssociationsBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasRelationshipAssociationsBulkUpdateParamsWithHTTPClient(client *http.Client) *ExtrasRelationshipAssociationsBulkUpdateParams {
	return &ExtrasRelationshipAssociationsBulkUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasRelationshipAssociationsBulkUpdateParams contains all the parameters to send to the API endpoint
   for the extras relationship associations bulk update operation.

   Typically these are written to a http.Request.
*/
type ExtrasRelationshipAssociationsBulkUpdateParams struct {

	// Data.
	Data *models.WritableRelationshipAssociation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras relationship associations bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasRelationshipAssociationsBulkUpdateParams) WithDefaults() *ExtrasRelationshipAssociationsBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras relationship associations bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasRelationshipAssociationsBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras relationship associations bulk update params
func (o *ExtrasRelationshipAssociationsBulkUpdateParams) WithTimeout(timeout time.Duration) *ExtrasRelationshipAssociationsBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras relationship associations bulk update params
func (o *ExtrasRelationshipAssociationsBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras relationship associations bulk update params
func (o *ExtrasRelationshipAssociationsBulkUpdateParams) WithContext(ctx context.Context) *ExtrasRelationshipAssociationsBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras relationship associations bulk update params
func (o *ExtrasRelationshipAssociationsBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras relationship associations bulk update params
func (o *ExtrasRelationshipAssociationsBulkUpdateParams) WithHTTPClient(client *http.Client) *ExtrasRelationshipAssociationsBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras relationship associations bulk update params
func (o *ExtrasRelationshipAssociationsBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras relationship associations bulk update params
func (o *ExtrasRelationshipAssociationsBulkUpdateParams) WithData(data *models.WritableRelationshipAssociation) *ExtrasRelationshipAssociationsBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras relationship associations bulk update params
func (o *ExtrasRelationshipAssociationsBulkUpdateParams) SetData(data *models.WritableRelationshipAssociation) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasRelationshipAssociationsBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

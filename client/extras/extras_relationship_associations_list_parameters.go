package extras

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewExtrasRelationshipAssociationsListParams creates a new ExtrasRelationshipAssociationsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasRelationshipAssociationsListParams() *ExtrasRelationshipAssociationsListParams {
	return &ExtrasRelationshipAssociationsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasRelationshipAssociationsListParamsWithTimeout creates a new ExtrasRelationshipAssociationsListParams object
// with the ability to set a timeout on a request.
func NewExtrasRelationshipAssociationsListParamsWithTimeout(timeout time.Duration) *ExtrasRelationshipAssociationsListParams {
	return &ExtrasRelationshipAssociationsListParams{
		timeout: timeout,
	}
}

// NewExtrasRelationshipAssociationsListParamsWithContext creates a new ExtrasRelationshipAssociationsListParams object
// with the ability to set a context for a request.
func NewExtrasRelationshipAssociationsListParamsWithContext(ctx context.Context) *ExtrasRelationshipAssociationsListParams {
	return &ExtrasRelationshipAssociationsListParams{
		Context: ctx,
	}
}

// NewExtrasRelationshipAssociationsListParamsWithHTTPClient creates a new ExtrasRelationshipAssociationsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasRelationshipAssociationsListParamsWithHTTPClient(client *http.Client) *ExtrasRelationshipAssociationsListParams {
	return &ExtrasRelationshipAssociationsListParams{
		HTTPClient: client,
	}
}

/* ExtrasRelationshipAssociationsListParams contains all the parameters to send to the API endpoint
   for the extras relationship associations list operation.

   Typically these are written to a http.Request.
*/
type ExtrasRelationshipAssociationsListParams struct {

	// DestinationID.
	DestinationID *string

	// DestinationIDIc.
	DestinationIDIc *string

	// DestinationIDIe.
	DestinationIDIe *string

	// DestinationIDIew.
	DestinationIDIew *string

	// DestinationIDIsw.
	DestinationIDIsw *string

	// DestinationIDn.
	DestinationIDn *string

	// DestinationIDNic.
	DestinationIDNic *string

	// DestinationIDNie.
	DestinationIDNie *string

	// DestinationIDNiew.
	DestinationIDNiew *string

	// DestinationIDNisw.
	DestinationIDNisw *string

	// DestinationType.
	DestinationType *string

	// DestinationTypen.
	DestinationTypen *string

	// ID.
	ID *string

	// IDIc.
	IDIc *string

	// IDIe.
	IDIe *string

	// IDIew.
	IDIew *string

	// IDIsw.
	IDIsw *string

	// IDn.
	IDn *string

	// IDNic.
	IDNic *string

	// IDNie.
	IDNie *string

	// IDNiew.
	IDNiew *string

	// IDNisw.
	IDNisw *string

	/* Limit.

	   Number of results to return per page.
	*/
	Limit *int64

	/* Offset.

	   The initial index from which to return the results.
	*/
	Offset *int64

	// Relationship.
	Relationship *string

	// Relationshipn.
	Relationshipn *string

	// SourceID.
	SourceID *string

	// SourceIDIc.
	SourceIDIc *string

	// SourceIDIe.
	SourceIDIe *string

	// SourceIDIew.
	SourceIDIew *string

	// SourceIDIsw.
	SourceIDIsw *string

	// SourceIDn.
	SourceIDn *string

	// SourceIDNic.
	SourceIDNic *string

	// SourceIDNie.
	SourceIDNie *string

	// SourceIDNiew.
	SourceIDNiew *string

	// SourceIDNisw.
	SourceIDNisw *string

	// SourceType.
	SourceType *string

	// SourceTypen.
	SourceTypen *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras relationship associations list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasRelationshipAssociationsListParams) WithDefaults() *ExtrasRelationshipAssociationsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras relationship associations list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasRelationshipAssociationsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) WithTimeout(timeout time.Duration) *ExtrasRelationshipAssociationsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) WithContext(ctx context.Context) *ExtrasRelationshipAssociationsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) WithHTTPClient(client *http.Client) *ExtrasRelationshipAssociationsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDestinationID adds the destinationID to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) WithDestinationID(destinationID *string) *ExtrasRelationshipAssociationsListParams {
	o.SetDestinationID(destinationID)
	return o
}

// SetDestinationID adds the destinationId to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) SetDestinationID(destinationID *string) {
	o.DestinationID = destinationID
}

// WithDestinationIDIc adds the destinationIDIc to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) WithDestinationIDIc(destinationIDIc *string) *ExtrasRelationshipAssociationsListParams {
	o.SetDestinationIDIc(destinationIDIc)
	return o
}

// SetDestinationIDIc adds the destinationIdIc to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) SetDestinationIDIc(destinationIDIc *string) {
	o.DestinationIDIc = destinationIDIc
}

// WithDestinationIDIe adds the destinationIDIe to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) WithDestinationIDIe(destinationIDIe *string) *ExtrasRelationshipAssociationsListParams {
	o.SetDestinationIDIe(destinationIDIe)
	return o
}

// SetDestinationIDIe adds the destinationIdIe to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) SetDestinationIDIe(destinationIDIe *string) {
	o.DestinationIDIe = destinationIDIe
}

// WithDestinationIDIew adds the destinationIDIew to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) WithDestinationIDIew(destinationIDIew *string) *ExtrasRelationshipAssociationsListParams {
	o.SetDestinationIDIew(destinationIDIew)
	return o
}

// SetDestinationIDIew adds the destinationIdIew to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) SetDestinationIDIew(destinationIDIew *string) {
	o.DestinationIDIew = destinationIDIew
}

// WithDestinationIDIsw adds the destinationIDIsw to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) WithDestinationIDIsw(destinationIDIsw *string) *ExtrasRelationshipAssociationsListParams {
	o.SetDestinationIDIsw(destinationIDIsw)
	return o
}

// SetDestinationIDIsw adds the destinationIdIsw to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) SetDestinationIDIsw(destinationIDIsw *string) {
	o.DestinationIDIsw = destinationIDIsw
}

// WithDestinationIDn adds the destinationIDn to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) WithDestinationIDn(destinationIDn *string) *ExtrasRelationshipAssociationsListParams {
	o.SetDestinationIDn(destinationIDn)
	return o
}

// SetDestinationIDn adds the destinationIdN to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) SetDestinationIDn(destinationIDn *string) {
	o.DestinationIDn = destinationIDn
}

// WithDestinationIDNic adds the destinationIDNic to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) WithDestinationIDNic(destinationIDNic *string) *ExtrasRelationshipAssociationsListParams {
	o.SetDestinationIDNic(destinationIDNic)
	return o
}

// SetDestinationIDNic adds the destinationIdNic to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) SetDestinationIDNic(destinationIDNic *string) {
	o.DestinationIDNic = destinationIDNic
}

// WithDestinationIDNie adds the destinationIDNie to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) WithDestinationIDNie(destinationIDNie *string) *ExtrasRelationshipAssociationsListParams {
	o.SetDestinationIDNie(destinationIDNie)
	return o
}

// SetDestinationIDNie adds the destinationIdNie to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) SetDestinationIDNie(destinationIDNie *string) {
	o.DestinationIDNie = destinationIDNie
}

// WithDestinationIDNiew adds the destinationIDNiew to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) WithDestinationIDNiew(destinationIDNiew *string) *ExtrasRelationshipAssociationsListParams {
	o.SetDestinationIDNiew(destinationIDNiew)
	return o
}

// SetDestinationIDNiew adds the destinationIdNiew to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) SetDestinationIDNiew(destinationIDNiew *string) {
	o.DestinationIDNiew = destinationIDNiew
}

// WithDestinationIDNisw adds the destinationIDNisw to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) WithDestinationIDNisw(destinationIDNisw *string) *ExtrasRelationshipAssociationsListParams {
	o.SetDestinationIDNisw(destinationIDNisw)
	return o
}

// SetDestinationIDNisw adds the destinationIdNisw to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) SetDestinationIDNisw(destinationIDNisw *string) {
	o.DestinationIDNisw = destinationIDNisw
}

// WithDestinationType adds the destinationType to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) WithDestinationType(destinationType *string) *ExtrasRelationshipAssociationsListParams {
	o.SetDestinationType(destinationType)
	return o
}

// SetDestinationType adds the destinationType to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) SetDestinationType(destinationType *string) {
	o.DestinationType = destinationType
}

// WithDestinationTypen adds the destinationTypen to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) WithDestinationTypen(destinationTypen *string) *ExtrasRelationshipAssociationsListParams {
	o.SetDestinationTypen(destinationTypen)
	return o
}

// SetDestinationTypen adds the destinationTypeN to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) SetDestinationTypen(destinationTypen *string) {
	o.DestinationTypen = destinationTypen
}

// WithID adds the id to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) WithID(id *string) *ExtrasRelationshipAssociationsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) WithIDIc(iDIc *string) *ExtrasRelationshipAssociationsListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) WithIDIe(iDIe *string) *ExtrasRelationshipAssociationsListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) WithIDIew(iDIew *string) *ExtrasRelationshipAssociationsListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) WithIDIsw(iDIsw *string) *ExtrasRelationshipAssociationsListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) WithIDn(iDn *string) *ExtrasRelationshipAssociationsListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) WithIDNic(iDNic *string) *ExtrasRelationshipAssociationsListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) WithIDNie(iDNie *string) *ExtrasRelationshipAssociationsListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) WithIDNiew(iDNiew *string) *ExtrasRelationshipAssociationsListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) WithIDNisw(iDNisw *string) *ExtrasRelationshipAssociationsListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithLimit adds the limit to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) WithLimit(limit *int64) *ExtrasRelationshipAssociationsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) WithOffset(offset *int64) *ExtrasRelationshipAssociationsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithRelationship adds the relationship to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) WithRelationship(relationship *string) *ExtrasRelationshipAssociationsListParams {
	o.SetRelationship(relationship)
	return o
}

// SetRelationship adds the relationship to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) SetRelationship(relationship *string) {
	o.Relationship = relationship
}

// WithRelationshipn adds the relationshipn to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) WithRelationshipn(relationshipn *string) *ExtrasRelationshipAssociationsListParams {
	o.SetRelationshipn(relationshipn)
	return o
}

// SetRelationshipn adds the relationshipN to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) SetRelationshipn(relationshipn *string) {
	o.Relationshipn = relationshipn
}

// WithSourceID adds the sourceID to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) WithSourceID(sourceID *string) *ExtrasRelationshipAssociationsListParams {
	o.SetSourceID(sourceID)
	return o
}

// SetSourceID adds the sourceId to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) SetSourceID(sourceID *string) {
	o.SourceID = sourceID
}

// WithSourceIDIc adds the sourceIDIc to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) WithSourceIDIc(sourceIDIc *string) *ExtrasRelationshipAssociationsListParams {
	o.SetSourceIDIc(sourceIDIc)
	return o
}

// SetSourceIDIc adds the sourceIdIc to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) SetSourceIDIc(sourceIDIc *string) {
	o.SourceIDIc = sourceIDIc
}

// WithSourceIDIe adds the sourceIDIe to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) WithSourceIDIe(sourceIDIe *string) *ExtrasRelationshipAssociationsListParams {
	o.SetSourceIDIe(sourceIDIe)
	return o
}

// SetSourceIDIe adds the sourceIdIe to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) SetSourceIDIe(sourceIDIe *string) {
	o.SourceIDIe = sourceIDIe
}

// WithSourceIDIew adds the sourceIDIew to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) WithSourceIDIew(sourceIDIew *string) *ExtrasRelationshipAssociationsListParams {
	o.SetSourceIDIew(sourceIDIew)
	return o
}

// SetSourceIDIew adds the sourceIdIew to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) SetSourceIDIew(sourceIDIew *string) {
	o.SourceIDIew = sourceIDIew
}

// WithSourceIDIsw adds the sourceIDIsw to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) WithSourceIDIsw(sourceIDIsw *string) *ExtrasRelationshipAssociationsListParams {
	o.SetSourceIDIsw(sourceIDIsw)
	return o
}

// SetSourceIDIsw adds the sourceIdIsw to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) SetSourceIDIsw(sourceIDIsw *string) {
	o.SourceIDIsw = sourceIDIsw
}

// WithSourceIDn adds the sourceIDn to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) WithSourceIDn(sourceIDn *string) *ExtrasRelationshipAssociationsListParams {
	o.SetSourceIDn(sourceIDn)
	return o
}

// SetSourceIDn adds the sourceIdN to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) SetSourceIDn(sourceIDn *string) {
	o.SourceIDn = sourceIDn
}

// WithSourceIDNic adds the sourceIDNic to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) WithSourceIDNic(sourceIDNic *string) *ExtrasRelationshipAssociationsListParams {
	o.SetSourceIDNic(sourceIDNic)
	return o
}

// SetSourceIDNic adds the sourceIdNic to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) SetSourceIDNic(sourceIDNic *string) {
	o.SourceIDNic = sourceIDNic
}

// WithSourceIDNie adds the sourceIDNie to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) WithSourceIDNie(sourceIDNie *string) *ExtrasRelationshipAssociationsListParams {
	o.SetSourceIDNie(sourceIDNie)
	return o
}

// SetSourceIDNie adds the sourceIdNie to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) SetSourceIDNie(sourceIDNie *string) {
	o.SourceIDNie = sourceIDNie
}

// WithSourceIDNiew adds the sourceIDNiew to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) WithSourceIDNiew(sourceIDNiew *string) *ExtrasRelationshipAssociationsListParams {
	o.SetSourceIDNiew(sourceIDNiew)
	return o
}

// SetSourceIDNiew adds the sourceIdNiew to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) SetSourceIDNiew(sourceIDNiew *string) {
	o.SourceIDNiew = sourceIDNiew
}

// WithSourceIDNisw adds the sourceIDNisw to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) WithSourceIDNisw(sourceIDNisw *string) *ExtrasRelationshipAssociationsListParams {
	o.SetSourceIDNisw(sourceIDNisw)
	return o
}

// SetSourceIDNisw adds the sourceIdNisw to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) SetSourceIDNisw(sourceIDNisw *string) {
	o.SourceIDNisw = sourceIDNisw
}

// WithSourceType adds the sourceType to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) WithSourceType(sourceType *string) *ExtrasRelationshipAssociationsListParams {
	o.SetSourceType(sourceType)
	return o
}

// SetSourceType adds the sourceType to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) SetSourceType(sourceType *string) {
	o.SourceType = sourceType
}

// WithSourceTypen adds the sourceTypen to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) WithSourceTypen(sourceTypen *string) *ExtrasRelationshipAssociationsListParams {
	o.SetSourceTypen(sourceTypen)
	return o
}

// SetSourceTypen adds the sourceTypeN to the extras relationship associations list params
func (o *ExtrasRelationshipAssociationsListParams) SetSourceTypen(sourceTypen *string) {
	o.SourceTypen = sourceTypen
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasRelationshipAssociationsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DestinationID != nil {

		// query param destination_id
		var qrDestinationID string

		if o.DestinationID != nil {
			qrDestinationID = *o.DestinationID
		}
		qDestinationID := qrDestinationID
		if qDestinationID != "" {

			if err := r.SetQueryParam("destination_id", qDestinationID); err != nil {
				return err
			}
		}
	}

	if o.DestinationIDIc != nil {

		// query param destination_id__ic
		var qrDestinationIDIc string

		if o.DestinationIDIc != nil {
			qrDestinationIDIc = *o.DestinationIDIc
		}
		qDestinationIDIc := qrDestinationIDIc
		if qDestinationIDIc != "" {

			if err := r.SetQueryParam("destination_id__ic", qDestinationIDIc); err != nil {
				return err
			}
		}
	}

	if o.DestinationIDIe != nil {

		// query param destination_id__ie
		var qrDestinationIDIe string

		if o.DestinationIDIe != nil {
			qrDestinationIDIe = *o.DestinationIDIe
		}
		qDestinationIDIe := qrDestinationIDIe
		if qDestinationIDIe != "" {

			if err := r.SetQueryParam("destination_id__ie", qDestinationIDIe); err != nil {
				return err
			}
		}
	}

	if o.DestinationIDIew != nil {

		// query param destination_id__iew
		var qrDestinationIDIew string

		if o.DestinationIDIew != nil {
			qrDestinationIDIew = *o.DestinationIDIew
		}
		qDestinationIDIew := qrDestinationIDIew
		if qDestinationIDIew != "" {

			if err := r.SetQueryParam("destination_id__iew", qDestinationIDIew); err != nil {
				return err
			}
		}
	}

	if o.DestinationIDIsw != nil {

		// query param destination_id__isw
		var qrDestinationIDIsw string

		if o.DestinationIDIsw != nil {
			qrDestinationIDIsw = *o.DestinationIDIsw
		}
		qDestinationIDIsw := qrDestinationIDIsw
		if qDestinationIDIsw != "" {

			if err := r.SetQueryParam("destination_id__isw", qDestinationIDIsw); err != nil {
				return err
			}
		}
	}

	if o.DestinationIDn != nil {

		// query param destination_id__n
		var qrDestinationIDn string

		if o.DestinationIDn != nil {
			qrDestinationIDn = *o.DestinationIDn
		}
		qDestinationIDn := qrDestinationIDn
		if qDestinationIDn != "" {

			if err := r.SetQueryParam("destination_id__n", qDestinationIDn); err != nil {
				return err
			}
		}
	}

	if o.DestinationIDNic != nil {

		// query param destination_id__nic
		var qrDestinationIDNic string

		if o.DestinationIDNic != nil {
			qrDestinationIDNic = *o.DestinationIDNic
		}
		qDestinationIDNic := qrDestinationIDNic
		if qDestinationIDNic != "" {

			if err := r.SetQueryParam("destination_id__nic", qDestinationIDNic); err != nil {
				return err
			}
		}
	}

	if o.DestinationIDNie != nil {

		// query param destination_id__nie
		var qrDestinationIDNie string

		if o.DestinationIDNie != nil {
			qrDestinationIDNie = *o.DestinationIDNie
		}
		qDestinationIDNie := qrDestinationIDNie
		if qDestinationIDNie != "" {

			if err := r.SetQueryParam("destination_id__nie", qDestinationIDNie); err != nil {
				return err
			}
		}
	}

	if o.DestinationIDNiew != nil {

		// query param destination_id__niew
		var qrDestinationIDNiew string

		if o.DestinationIDNiew != nil {
			qrDestinationIDNiew = *o.DestinationIDNiew
		}
		qDestinationIDNiew := qrDestinationIDNiew
		if qDestinationIDNiew != "" {

			if err := r.SetQueryParam("destination_id__niew", qDestinationIDNiew); err != nil {
				return err
			}
		}
	}

	if o.DestinationIDNisw != nil {

		// query param destination_id__nisw
		var qrDestinationIDNisw string

		if o.DestinationIDNisw != nil {
			qrDestinationIDNisw = *o.DestinationIDNisw
		}
		qDestinationIDNisw := qrDestinationIDNisw
		if qDestinationIDNisw != "" {

			if err := r.SetQueryParam("destination_id__nisw", qDestinationIDNisw); err != nil {
				return err
			}
		}
	}

	if o.DestinationType != nil {

		// query param destination_type
		var qrDestinationType string

		if o.DestinationType != nil {
			qrDestinationType = *o.DestinationType
		}
		qDestinationType := qrDestinationType
		if qDestinationType != "" {

			if err := r.SetQueryParam("destination_type", qDestinationType); err != nil {
				return err
			}
		}
	}

	if o.DestinationTypen != nil {

		// query param destination_type__n
		var qrDestinationTypen string

		if o.DestinationTypen != nil {
			qrDestinationTypen = *o.DestinationTypen
		}
		qDestinationTypen := qrDestinationTypen
		if qDestinationTypen != "" {

			if err := r.SetQueryParam("destination_type__n", qDestinationTypen); err != nil {
				return err
			}
		}
	}

	if o.ID != nil {

		// query param id
		var qrID string

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}
	}

	if o.IDIc != nil {

		// query param id__ic
		var qrIDIc string

		if o.IDIc != nil {
			qrIDIc = *o.IDIc
		}
		qIDIc := qrIDIc
		if qIDIc != "" {

			if err := r.SetQueryParam("id__ic", qIDIc); err != nil {
				return err
			}
		}
	}

	if o.IDIe != nil {

		// query param id__ie
		var qrIDIe string

		if o.IDIe != nil {
			qrIDIe = *o.IDIe
		}
		qIDIe := qrIDIe
		if qIDIe != "" {

			if err := r.SetQueryParam("id__ie", qIDIe); err != nil {
				return err
			}
		}
	}

	if o.IDIew != nil {

		// query param id__iew
		var qrIDIew string

		if o.IDIew != nil {
			qrIDIew = *o.IDIew
		}
		qIDIew := qrIDIew
		if qIDIew != "" {

			if err := r.SetQueryParam("id__iew", qIDIew); err != nil {
				return err
			}
		}
	}

	if o.IDIsw != nil {

		// query param id__isw
		var qrIDIsw string

		if o.IDIsw != nil {
			qrIDIsw = *o.IDIsw
		}
		qIDIsw := qrIDIsw
		if qIDIsw != "" {

			if err := r.SetQueryParam("id__isw", qIDIsw); err != nil {
				return err
			}
		}
	}

	if o.IDn != nil {

		// query param id__n
		var qrIDn string

		if o.IDn != nil {
			qrIDn = *o.IDn
		}
		qIDn := qrIDn
		if qIDn != "" {

			if err := r.SetQueryParam("id__n", qIDn); err != nil {
				return err
			}
		}
	}

	if o.IDNic != nil {

		// query param id__nic
		var qrIDNic string

		if o.IDNic != nil {
			qrIDNic = *o.IDNic
		}
		qIDNic := qrIDNic
		if qIDNic != "" {

			if err := r.SetQueryParam("id__nic", qIDNic); err != nil {
				return err
			}
		}
	}

	if o.IDNie != nil {

		// query param id__nie
		var qrIDNie string

		if o.IDNie != nil {
			qrIDNie = *o.IDNie
		}
		qIDNie := qrIDNie
		if qIDNie != "" {

			if err := r.SetQueryParam("id__nie", qIDNie); err != nil {
				return err
			}
		}
	}

	if o.IDNiew != nil {

		// query param id__niew
		var qrIDNiew string

		if o.IDNiew != nil {
			qrIDNiew = *o.IDNiew
		}
		qIDNiew := qrIDNiew
		if qIDNiew != "" {

			if err := r.SetQueryParam("id__niew", qIDNiew); err != nil {
				return err
			}
		}
	}

	if o.IDNisw != nil {

		// query param id__nisw
		var qrIDNisw string

		if o.IDNisw != nil {
			qrIDNisw = *o.IDNisw
		}
		qIDNisw := qrIDNisw
		if qIDNisw != "" {

			if err := r.SetQueryParam("id__nisw", qIDNisw); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Relationship != nil {

		// query param relationship
		var qrRelationship string

		if o.Relationship != nil {
			qrRelationship = *o.Relationship
		}
		qRelationship := qrRelationship
		if qRelationship != "" {

			if err := r.SetQueryParam("relationship", qRelationship); err != nil {
				return err
			}
		}
	}

	if o.Relationshipn != nil {

		// query param relationship__n
		var qrRelationshipn string

		if o.Relationshipn != nil {
			qrRelationshipn = *o.Relationshipn
		}
		qRelationshipn := qrRelationshipn
		if qRelationshipn != "" {

			if err := r.SetQueryParam("relationship__n", qRelationshipn); err != nil {
				return err
			}
		}
	}

	if o.SourceID != nil {

		// query param source_id
		var qrSourceID string

		if o.SourceID != nil {
			qrSourceID = *o.SourceID
		}
		qSourceID := qrSourceID
		if qSourceID != "" {

			if err := r.SetQueryParam("source_id", qSourceID); err != nil {
				return err
			}
		}
	}

	if o.SourceIDIc != nil {

		// query param source_id__ic
		var qrSourceIDIc string

		if o.SourceIDIc != nil {
			qrSourceIDIc = *o.SourceIDIc
		}
		qSourceIDIc := qrSourceIDIc
		if qSourceIDIc != "" {

			if err := r.SetQueryParam("source_id__ic", qSourceIDIc); err != nil {
				return err
			}
		}
	}

	if o.SourceIDIe != nil {

		// query param source_id__ie
		var qrSourceIDIe string

		if o.SourceIDIe != nil {
			qrSourceIDIe = *o.SourceIDIe
		}
		qSourceIDIe := qrSourceIDIe
		if qSourceIDIe != "" {

			if err := r.SetQueryParam("source_id__ie", qSourceIDIe); err != nil {
				return err
			}
		}
	}

	if o.SourceIDIew != nil {

		// query param source_id__iew
		var qrSourceIDIew string

		if o.SourceIDIew != nil {
			qrSourceIDIew = *o.SourceIDIew
		}
		qSourceIDIew := qrSourceIDIew
		if qSourceIDIew != "" {

			if err := r.SetQueryParam("source_id__iew", qSourceIDIew); err != nil {
				return err
			}
		}
	}

	if o.SourceIDIsw != nil {

		// query param source_id__isw
		var qrSourceIDIsw string

		if o.SourceIDIsw != nil {
			qrSourceIDIsw = *o.SourceIDIsw
		}
		qSourceIDIsw := qrSourceIDIsw
		if qSourceIDIsw != "" {

			if err := r.SetQueryParam("source_id__isw", qSourceIDIsw); err != nil {
				return err
			}
		}
	}

	if o.SourceIDn != nil {

		// query param source_id__n
		var qrSourceIDn string

		if o.SourceIDn != nil {
			qrSourceIDn = *o.SourceIDn
		}
		qSourceIDn := qrSourceIDn
		if qSourceIDn != "" {

			if err := r.SetQueryParam("source_id__n", qSourceIDn); err != nil {
				return err
			}
		}
	}

	if o.SourceIDNic != nil {

		// query param source_id__nic
		var qrSourceIDNic string

		if o.SourceIDNic != nil {
			qrSourceIDNic = *o.SourceIDNic
		}
		qSourceIDNic := qrSourceIDNic
		if qSourceIDNic != "" {

			if err := r.SetQueryParam("source_id__nic", qSourceIDNic); err != nil {
				return err
			}
		}
	}

	if o.SourceIDNie != nil {

		// query param source_id__nie
		var qrSourceIDNie string

		if o.SourceIDNie != nil {
			qrSourceIDNie = *o.SourceIDNie
		}
		qSourceIDNie := qrSourceIDNie
		if qSourceIDNie != "" {

			if err := r.SetQueryParam("source_id__nie", qSourceIDNie); err != nil {
				return err
			}
		}
	}

	if o.SourceIDNiew != nil {

		// query param source_id__niew
		var qrSourceIDNiew string

		if o.SourceIDNiew != nil {
			qrSourceIDNiew = *o.SourceIDNiew
		}
		qSourceIDNiew := qrSourceIDNiew
		if qSourceIDNiew != "" {

			if err := r.SetQueryParam("source_id__niew", qSourceIDNiew); err != nil {
				return err
			}
		}
	}

	if o.SourceIDNisw != nil {

		// query param source_id__nisw
		var qrSourceIDNisw string

		if o.SourceIDNisw != nil {
			qrSourceIDNisw = *o.SourceIDNisw
		}
		qSourceIDNisw := qrSourceIDNisw
		if qSourceIDNisw != "" {

			if err := r.SetQueryParam("source_id__nisw", qSourceIDNisw); err != nil {
				return err
			}
		}
	}

	if o.SourceType != nil {

		// query param source_type
		var qrSourceType string

		if o.SourceType != nil {
			qrSourceType = *o.SourceType
		}
		qSourceType := qrSourceType
		if qSourceType != "" {

			if err := r.SetQueryParam("source_type", qSourceType); err != nil {
				return err
			}
		}
	}

	if o.SourceTypen != nil {

		// query param source_type__n
		var qrSourceTypen string

		if o.SourceTypen != nil {
			qrSourceTypen = *o.SourceTypen
		}
		qSourceTypen := qrSourceTypen
		if qSourceTypen != "" {

			if err := r.SetQueryParam("source_type__n", qSourceTypen); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

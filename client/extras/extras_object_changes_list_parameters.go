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

// NewExtrasObjectChangesListParams creates a new ExtrasObjectChangesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasObjectChangesListParams() *ExtrasObjectChangesListParams {
	return &ExtrasObjectChangesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasObjectChangesListParamsWithTimeout creates a new ExtrasObjectChangesListParams object
// with the ability to set a timeout on a request.
func NewExtrasObjectChangesListParamsWithTimeout(timeout time.Duration) *ExtrasObjectChangesListParams {
	return &ExtrasObjectChangesListParams{
		timeout: timeout,
	}
}

// NewExtrasObjectChangesListParamsWithContext creates a new ExtrasObjectChangesListParams object
// with the ability to set a context for a request.
func NewExtrasObjectChangesListParamsWithContext(ctx context.Context) *ExtrasObjectChangesListParams {
	return &ExtrasObjectChangesListParams{
		Context: ctx,
	}
}

// NewExtrasObjectChangesListParamsWithHTTPClient creates a new ExtrasObjectChangesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasObjectChangesListParamsWithHTTPClient(client *http.Client) *ExtrasObjectChangesListParams {
	return &ExtrasObjectChangesListParams{
		HTTPClient: client,
	}
}

/* ExtrasObjectChangesListParams contains all the parameters to send to the API endpoint
   for the extras object changes list operation.

   Typically these are written to a http.Request.
*/
type ExtrasObjectChangesListParams struct {

	// Action.
	Action *string

	// Actionn.
	Actionn *string

	// ChangedObjectID.
	ChangedObjectID *string

	// ChangedObjectIDIc.
	ChangedObjectIDIc *string

	// ChangedObjectIDIe.
	ChangedObjectIDIe *string

	// ChangedObjectIDIew.
	ChangedObjectIDIew *string

	// ChangedObjectIDIsw.
	ChangedObjectIDIsw *string

	// ChangedObjectIDn.
	ChangedObjectIDn *string

	// ChangedObjectIDNic.
	ChangedObjectIDNic *string

	// ChangedObjectIDNie.
	ChangedObjectIDNie *string

	// ChangedObjectIDNiew.
	ChangedObjectIDNiew *string

	// ChangedObjectIDNisw.
	ChangedObjectIDNisw *string

	// ChangedObjectType.
	ChangedObjectType *string

	// ChangedObjectTypen.
	ChangedObjectTypen *string

	// ChangedObjectTypeID.
	ChangedObjectTypeID *string

	// ChangedObjectTypeIDn.
	ChangedObjectTypeIDn *string

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

	// ObjectRepr.
	ObjectRepr *string

	// ObjectReprIc.
	ObjectReprIc *string

	// ObjectReprIe.
	ObjectReprIe *string

	// ObjectReprIew.
	ObjectReprIew *string

	// ObjectReprIsw.
	ObjectReprIsw *string

	// ObjectReprn.
	ObjectReprn *string

	// ObjectReprNic.
	ObjectReprNic *string

	// ObjectReprNie.
	ObjectReprNie *string

	// ObjectReprNiew.
	ObjectReprNiew *string

	// ObjectReprNisw.
	ObjectReprNisw *string

	/* Offset.

	   The initial index from which to return the results.
	*/
	Offset *int64

	// Q.
	Q *string

	// RequestID.
	RequestID *string

	// RequestIDIc.
	RequestIDIc *string

	// RequestIDIe.
	RequestIDIe *string

	// RequestIDIew.
	RequestIDIew *string

	// RequestIDIsw.
	RequestIDIsw *string

	// RequestIDn.
	RequestIDn *string

	// RequestIDNic.
	RequestIDNic *string

	// RequestIDNie.
	RequestIDNie *string

	// RequestIDNiew.
	RequestIDNiew *string

	// RequestIDNisw.
	RequestIDNisw *string

	// Time.
	Time *string

	// User.
	User *string

	// Usern.
	Usern *string

	// UserID.
	UserID *string

	// UserIDn.
	UserIDn *string

	// UserName.
	UserName *string

	// UserNameIc.
	UserNameIc *string

	// UserNameIe.
	UserNameIe *string

	// UserNameIew.
	UserNameIew *string

	// UserNameIsw.
	UserNameIsw *string

	// UserNamen.
	UserNamen *string

	// UserNameNic.
	UserNameNic *string

	// UserNameNie.
	UserNameNie *string

	// UserNameNiew.
	UserNameNiew *string

	// UserNameNisw.
	UserNameNisw *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras object changes list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasObjectChangesListParams) WithDefaults() *ExtrasObjectChangesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras object changes list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasObjectChangesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithTimeout(timeout time.Duration) *ExtrasObjectChangesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithContext(ctx context.Context) *ExtrasObjectChangesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithHTTPClient(client *http.Client) *ExtrasObjectChangesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAction adds the action to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithAction(action *string) *ExtrasObjectChangesListParams {
	o.SetAction(action)
	return o
}

// SetAction adds the action to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetAction(action *string) {
	o.Action = action
}

// WithActionn adds the actionn to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithActionn(actionn *string) *ExtrasObjectChangesListParams {
	o.SetActionn(actionn)
	return o
}

// SetActionn adds the actionN to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetActionn(actionn *string) {
	o.Actionn = actionn
}

// WithChangedObjectID adds the changedObjectID to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithChangedObjectID(changedObjectID *string) *ExtrasObjectChangesListParams {
	o.SetChangedObjectID(changedObjectID)
	return o
}

// SetChangedObjectID adds the changedObjectId to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetChangedObjectID(changedObjectID *string) {
	o.ChangedObjectID = changedObjectID
}

// WithChangedObjectIDIc adds the changedObjectIDIc to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithChangedObjectIDIc(changedObjectIDIc *string) *ExtrasObjectChangesListParams {
	o.SetChangedObjectIDIc(changedObjectIDIc)
	return o
}

// SetChangedObjectIDIc adds the changedObjectIdIc to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetChangedObjectIDIc(changedObjectIDIc *string) {
	o.ChangedObjectIDIc = changedObjectIDIc
}

// WithChangedObjectIDIe adds the changedObjectIDIe to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithChangedObjectIDIe(changedObjectIDIe *string) *ExtrasObjectChangesListParams {
	o.SetChangedObjectIDIe(changedObjectIDIe)
	return o
}

// SetChangedObjectIDIe adds the changedObjectIdIe to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetChangedObjectIDIe(changedObjectIDIe *string) {
	o.ChangedObjectIDIe = changedObjectIDIe
}

// WithChangedObjectIDIew adds the changedObjectIDIew to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithChangedObjectIDIew(changedObjectIDIew *string) *ExtrasObjectChangesListParams {
	o.SetChangedObjectIDIew(changedObjectIDIew)
	return o
}

// SetChangedObjectIDIew adds the changedObjectIdIew to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetChangedObjectIDIew(changedObjectIDIew *string) {
	o.ChangedObjectIDIew = changedObjectIDIew
}

// WithChangedObjectIDIsw adds the changedObjectIDIsw to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithChangedObjectIDIsw(changedObjectIDIsw *string) *ExtrasObjectChangesListParams {
	o.SetChangedObjectIDIsw(changedObjectIDIsw)
	return o
}

// SetChangedObjectIDIsw adds the changedObjectIdIsw to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetChangedObjectIDIsw(changedObjectIDIsw *string) {
	o.ChangedObjectIDIsw = changedObjectIDIsw
}

// WithChangedObjectIDn adds the changedObjectIDn to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithChangedObjectIDn(changedObjectIDn *string) *ExtrasObjectChangesListParams {
	o.SetChangedObjectIDn(changedObjectIDn)
	return o
}

// SetChangedObjectIDn adds the changedObjectIdN to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetChangedObjectIDn(changedObjectIDn *string) {
	o.ChangedObjectIDn = changedObjectIDn
}

// WithChangedObjectIDNic adds the changedObjectIDNic to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithChangedObjectIDNic(changedObjectIDNic *string) *ExtrasObjectChangesListParams {
	o.SetChangedObjectIDNic(changedObjectIDNic)
	return o
}

// SetChangedObjectIDNic adds the changedObjectIdNic to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetChangedObjectIDNic(changedObjectIDNic *string) {
	o.ChangedObjectIDNic = changedObjectIDNic
}

// WithChangedObjectIDNie adds the changedObjectIDNie to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithChangedObjectIDNie(changedObjectIDNie *string) *ExtrasObjectChangesListParams {
	o.SetChangedObjectIDNie(changedObjectIDNie)
	return o
}

// SetChangedObjectIDNie adds the changedObjectIdNie to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetChangedObjectIDNie(changedObjectIDNie *string) {
	o.ChangedObjectIDNie = changedObjectIDNie
}

// WithChangedObjectIDNiew adds the changedObjectIDNiew to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithChangedObjectIDNiew(changedObjectIDNiew *string) *ExtrasObjectChangesListParams {
	o.SetChangedObjectIDNiew(changedObjectIDNiew)
	return o
}

// SetChangedObjectIDNiew adds the changedObjectIdNiew to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetChangedObjectIDNiew(changedObjectIDNiew *string) {
	o.ChangedObjectIDNiew = changedObjectIDNiew
}

// WithChangedObjectIDNisw adds the changedObjectIDNisw to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithChangedObjectIDNisw(changedObjectIDNisw *string) *ExtrasObjectChangesListParams {
	o.SetChangedObjectIDNisw(changedObjectIDNisw)
	return o
}

// SetChangedObjectIDNisw adds the changedObjectIdNisw to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetChangedObjectIDNisw(changedObjectIDNisw *string) {
	o.ChangedObjectIDNisw = changedObjectIDNisw
}

// WithChangedObjectType adds the changedObjectType to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithChangedObjectType(changedObjectType *string) *ExtrasObjectChangesListParams {
	o.SetChangedObjectType(changedObjectType)
	return o
}

// SetChangedObjectType adds the changedObjectType to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetChangedObjectType(changedObjectType *string) {
	o.ChangedObjectType = changedObjectType
}

// WithChangedObjectTypen adds the changedObjectTypen to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithChangedObjectTypen(changedObjectTypen *string) *ExtrasObjectChangesListParams {
	o.SetChangedObjectTypen(changedObjectTypen)
	return o
}

// SetChangedObjectTypen adds the changedObjectTypeN to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetChangedObjectTypen(changedObjectTypen *string) {
	o.ChangedObjectTypen = changedObjectTypen
}

// WithChangedObjectTypeID adds the changedObjectTypeID to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithChangedObjectTypeID(changedObjectTypeID *string) *ExtrasObjectChangesListParams {
	o.SetChangedObjectTypeID(changedObjectTypeID)
	return o
}

// SetChangedObjectTypeID adds the changedObjectTypeId to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetChangedObjectTypeID(changedObjectTypeID *string) {
	o.ChangedObjectTypeID = changedObjectTypeID
}

// WithChangedObjectTypeIDn adds the changedObjectTypeIDn to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithChangedObjectTypeIDn(changedObjectTypeIDn *string) *ExtrasObjectChangesListParams {
	o.SetChangedObjectTypeIDn(changedObjectTypeIDn)
	return o
}

// SetChangedObjectTypeIDn adds the changedObjectTypeIdN to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetChangedObjectTypeIDn(changedObjectTypeIDn *string) {
	o.ChangedObjectTypeIDn = changedObjectTypeIDn
}

// WithID adds the id to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithID(id *string) *ExtrasObjectChangesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithIDIc(iDIc *string) *ExtrasObjectChangesListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithIDIe(iDIe *string) *ExtrasObjectChangesListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithIDIew(iDIew *string) *ExtrasObjectChangesListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithIDIsw(iDIsw *string) *ExtrasObjectChangesListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithIDn(iDn *string) *ExtrasObjectChangesListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithIDNic(iDNic *string) *ExtrasObjectChangesListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithIDNie(iDNie *string) *ExtrasObjectChangesListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithIDNiew(iDNiew *string) *ExtrasObjectChangesListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithIDNisw(iDNisw *string) *ExtrasObjectChangesListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithLimit adds the limit to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithLimit(limit *int64) *ExtrasObjectChangesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithObjectRepr adds the objectRepr to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithObjectRepr(objectRepr *string) *ExtrasObjectChangesListParams {
	o.SetObjectRepr(objectRepr)
	return o
}

// SetObjectRepr adds the objectRepr to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetObjectRepr(objectRepr *string) {
	o.ObjectRepr = objectRepr
}

// WithObjectReprIc adds the objectReprIc to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithObjectReprIc(objectReprIc *string) *ExtrasObjectChangesListParams {
	o.SetObjectReprIc(objectReprIc)
	return o
}

// SetObjectReprIc adds the objectReprIc to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetObjectReprIc(objectReprIc *string) {
	o.ObjectReprIc = objectReprIc
}

// WithObjectReprIe adds the objectReprIe to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithObjectReprIe(objectReprIe *string) *ExtrasObjectChangesListParams {
	o.SetObjectReprIe(objectReprIe)
	return o
}

// SetObjectReprIe adds the objectReprIe to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetObjectReprIe(objectReprIe *string) {
	o.ObjectReprIe = objectReprIe
}

// WithObjectReprIew adds the objectReprIew to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithObjectReprIew(objectReprIew *string) *ExtrasObjectChangesListParams {
	o.SetObjectReprIew(objectReprIew)
	return o
}

// SetObjectReprIew adds the objectReprIew to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetObjectReprIew(objectReprIew *string) {
	o.ObjectReprIew = objectReprIew
}

// WithObjectReprIsw adds the objectReprIsw to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithObjectReprIsw(objectReprIsw *string) *ExtrasObjectChangesListParams {
	o.SetObjectReprIsw(objectReprIsw)
	return o
}

// SetObjectReprIsw adds the objectReprIsw to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetObjectReprIsw(objectReprIsw *string) {
	o.ObjectReprIsw = objectReprIsw
}

// WithObjectReprn adds the objectReprn to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithObjectReprn(objectReprn *string) *ExtrasObjectChangesListParams {
	o.SetObjectReprn(objectReprn)
	return o
}

// SetObjectReprn adds the objectReprN to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetObjectReprn(objectReprn *string) {
	o.ObjectReprn = objectReprn
}

// WithObjectReprNic adds the objectReprNic to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithObjectReprNic(objectReprNic *string) *ExtrasObjectChangesListParams {
	o.SetObjectReprNic(objectReprNic)
	return o
}

// SetObjectReprNic adds the objectReprNic to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetObjectReprNic(objectReprNic *string) {
	o.ObjectReprNic = objectReprNic
}

// WithObjectReprNie adds the objectReprNie to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithObjectReprNie(objectReprNie *string) *ExtrasObjectChangesListParams {
	o.SetObjectReprNie(objectReprNie)
	return o
}

// SetObjectReprNie adds the objectReprNie to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetObjectReprNie(objectReprNie *string) {
	o.ObjectReprNie = objectReprNie
}

// WithObjectReprNiew adds the objectReprNiew to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithObjectReprNiew(objectReprNiew *string) *ExtrasObjectChangesListParams {
	o.SetObjectReprNiew(objectReprNiew)
	return o
}

// SetObjectReprNiew adds the objectReprNiew to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetObjectReprNiew(objectReprNiew *string) {
	o.ObjectReprNiew = objectReprNiew
}

// WithObjectReprNisw adds the objectReprNisw to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithObjectReprNisw(objectReprNisw *string) *ExtrasObjectChangesListParams {
	o.SetObjectReprNisw(objectReprNisw)
	return o
}

// SetObjectReprNisw adds the objectReprNisw to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetObjectReprNisw(objectReprNisw *string) {
	o.ObjectReprNisw = objectReprNisw
}

// WithOffset adds the offset to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithOffset(offset *int64) *ExtrasObjectChangesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithQ(q *string) *ExtrasObjectChangesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetQ(q *string) {
	o.Q = q
}

// WithRequestID adds the requestID to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithRequestID(requestID *string) *ExtrasObjectChangesListParams {
	o.SetRequestID(requestID)
	return o
}

// SetRequestID adds the requestId to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetRequestID(requestID *string) {
	o.RequestID = requestID
}

// WithRequestIDIc adds the requestIDIc to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithRequestIDIc(requestIDIc *string) *ExtrasObjectChangesListParams {
	o.SetRequestIDIc(requestIDIc)
	return o
}

// SetRequestIDIc adds the requestIdIc to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetRequestIDIc(requestIDIc *string) {
	o.RequestIDIc = requestIDIc
}

// WithRequestIDIe adds the requestIDIe to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithRequestIDIe(requestIDIe *string) *ExtrasObjectChangesListParams {
	o.SetRequestIDIe(requestIDIe)
	return o
}

// SetRequestIDIe adds the requestIdIe to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetRequestIDIe(requestIDIe *string) {
	o.RequestIDIe = requestIDIe
}

// WithRequestIDIew adds the requestIDIew to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithRequestIDIew(requestIDIew *string) *ExtrasObjectChangesListParams {
	o.SetRequestIDIew(requestIDIew)
	return o
}

// SetRequestIDIew adds the requestIdIew to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetRequestIDIew(requestIDIew *string) {
	o.RequestIDIew = requestIDIew
}

// WithRequestIDIsw adds the requestIDIsw to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithRequestIDIsw(requestIDIsw *string) *ExtrasObjectChangesListParams {
	o.SetRequestIDIsw(requestIDIsw)
	return o
}

// SetRequestIDIsw adds the requestIdIsw to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetRequestIDIsw(requestIDIsw *string) {
	o.RequestIDIsw = requestIDIsw
}

// WithRequestIDn adds the requestIDn to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithRequestIDn(requestIDn *string) *ExtrasObjectChangesListParams {
	o.SetRequestIDn(requestIDn)
	return o
}

// SetRequestIDn adds the requestIdN to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetRequestIDn(requestIDn *string) {
	o.RequestIDn = requestIDn
}

// WithRequestIDNic adds the requestIDNic to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithRequestIDNic(requestIDNic *string) *ExtrasObjectChangesListParams {
	o.SetRequestIDNic(requestIDNic)
	return o
}

// SetRequestIDNic adds the requestIdNic to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetRequestIDNic(requestIDNic *string) {
	o.RequestIDNic = requestIDNic
}

// WithRequestIDNie adds the requestIDNie to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithRequestIDNie(requestIDNie *string) *ExtrasObjectChangesListParams {
	o.SetRequestIDNie(requestIDNie)
	return o
}

// SetRequestIDNie adds the requestIdNie to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetRequestIDNie(requestIDNie *string) {
	o.RequestIDNie = requestIDNie
}

// WithRequestIDNiew adds the requestIDNiew to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithRequestIDNiew(requestIDNiew *string) *ExtrasObjectChangesListParams {
	o.SetRequestIDNiew(requestIDNiew)
	return o
}

// SetRequestIDNiew adds the requestIdNiew to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetRequestIDNiew(requestIDNiew *string) {
	o.RequestIDNiew = requestIDNiew
}

// WithRequestIDNisw adds the requestIDNisw to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithRequestIDNisw(requestIDNisw *string) *ExtrasObjectChangesListParams {
	o.SetRequestIDNisw(requestIDNisw)
	return o
}

// SetRequestIDNisw adds the requestIdNisw to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetRequestIDNisw(requestIDNisw *string) {
	o.RequestIDNisw = requestIDNisw
}

// WithTime adds the time to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithTime(time *string) *ExtrasObjectChangesListParams {
	o.SetTime(time)
	return o
}

// SetTime adds the time to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetTime(time *string) {
	o.Time = time
}

// WithUser adds the user to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithUser(user *string) *ExtrasObjectChangesListParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetUser(user *string) {
	o.User = user
}

// WithUsern adds the usern to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithUsern(usern *string) *ExtrasObjectChangesListParams {
	o.SetUsern(usern)
	return o
}

// SetUsern adds the userN to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetUsern(usern *string) {
	o.Usern = usern
}

// WithUserID adds the userID to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithUserID(userID *string) *ExtrasObjectChangesListParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetUserID(userID *string) {
	o.UserID = userID
}

// WithUserIDn adds the userIDn to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithUserIDn(userIDn *string) *ExtrasObjectChangesListParams {
	o.SetUserIDn(userIDn)
	return o
}

// SetUserIDn adds the userIdN to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetUserIDn(userIDn *string) {
	o.UserIDn = userIDn
}

// WithUserName adds the userName to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithUserName(userName *string) *ExtrasObjectChangesListParams {
	o.SetUserName(userName)
	return o
}

// SetUserName adds the userName to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetUserName(userName *string) {
	o.UserName = userName
}

// WithUserNameIc adds the userNameIc to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithUserNameIc(userNameIc *string) *ExtrasObjectChangesListParams {
	o.SetUserNameIc(userNameIc)
	return o
}

// SetUserNameIc adds the userNameIc to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetUserNameIc(userNameIc *string) {
	o.UserNameIc = userNameIc
}

// WithUserNameIe adds the userNameIe to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithUserNameIe(userNameIe *string) *ExtrasObjectChangesListParams {
	o.SetUserNameIe(userNameIe)
	return o
}

// SetUserNameIe adds the userNameIe to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetUserNameIe(userNameIe *string) {
	o.UserNameIe = userNameIe
}

// WithUserNameIew adds the userNameIew to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithUserNameIew(userNameIew *string) *ExtrasObjectChangesListParams {
	o.SetUserNameIew(userNameIew)
	return o
}

// SetUserNameIew adds the userNameIew to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetUserNameIew(userNameIew *string) {
	o.UserNameIew = userNameIew
}

// WithUserNameIsw adds the userNameIsw to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithUserNameIsw(userNameIsw *string) *ExtrasObjectChangesListParams {
	o.SetUserNameIsw(userNameIsw)
	return o
}

// SetUserNameIsw adds the userNameIsw to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetUserNameIsw(userNameIsw *string) {
	o.UserNameIsw = userNameIsw
}

// WithUserNamen adds the userNamen to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithUserNamen(userNamen *string) *ExtrasObjectChangesListParams {
	o.SetUserNamen(userNamen)
	return o
}

// SetUserNamen adds the userNameN to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetUserNamen(userNamen *string) {
	o.UserNamen = userNamen
}

// WithUserNameNic adds the userNameNic to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithUserNameNic(userNameNic *string) *ExtrasObjectChangesListParams {
	o.SetUserNameNic(userNameNic)
	return o
}

// SetUserNameNic adds the userNameNic to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetUserNameNic(userNameNic *string) {
	o.UserNameNic = userNameNic
}

// WithUserNameNie adds the userNameNie to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithUserNameNie(userNameNie *string) *ExtrasObjectChangesListParams {
	o.SetUserNameNie(userNameNie)
	return o
}

// SetUserNameNie adds the userNameNie to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetUserNameNie(userNameNie *string) {
	o.UserNameNie = userNameNie
}

// WithUserNameNiew adds the userNameNiew to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithUserNameNiew(userNameNiew *string) *ExtrasObjectChangesListParams {
	o.SetUserNameNiew(userNameNiew)
	return o
}

// SetUserNameNiew adds the userNameNiew to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetUserNameNiew(userNameNiew *string) {
	o.UserNameNiew = userNameNiew
}

// WithUserNameNisw adds the userNameNisw to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithUserNameNisw(userNameNisw *string) *ExtrasObjectChangesListParams {
	o.SetUserNameNisw(userNameNisw)
	return o
}

// SetUserNameNisw adds the userNameNisw to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetUserNameNisw(userNameNisw *string) {
	o.UserNameNisw = userNameNisw
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasObjectChangesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Action != nil {

		// query param action
		var qrAction string

		if o.Action != nil {
			qrAction = *o.Action
		}
		qAction := qrAction
		if qAction != "" {

			if err := r.SetQueryParam("action", qAction); err != nil {
				return err
			}
		}
	}

	if o.Actionn != nil {

		// query param action__n
		var qrActionn string

		if o.Actionn != nil {
			qrActionn = *o.Actionn
		}
		qActionn := qrActionn
		if qActionn != "" {

			if err := r.SetQueryParam("action__n", qActionn); err != nil {
				return err
			}
		}
	}

	if o.ChangedObjectID != nil {

		// query param changed_object_id
		var qrChangedObjectID string

		if o.ChangedObjectID != nil {
			qrChangedObjectID = *o.ChangedObjectID
		}
		qChangedObjectID := qrChangedObjectID
		if qChangedObjectID != "" {

			if err := r.SetQueryParam("changed_object_id", qChangedObjectID); err != nil {
				return err
			}
		}
	}

	if o.ChangedObjectIDIc != nil {

		// query param changed_object_id__ic
		var qrChangedObjectIDIc string

		if o.ChangedObjectIDIc != nil {
			qrChangedObjectIDIc = *o.ChangedObjectIDIc
		}
		qChangedObjectIDIc := qrChangedObjectIDIc
		if qChangedObjectIDIc != "" {

			if err := r.SetQueryParam("changed_object_id__ic", qChangedObjectIDIc); err != nil {
				return err
			}
		}
	}

	if o.ChangedObjectIDIe != nil {

		// query param changed_object_id__ie
		var qrChangedObjectIDIe string

		if o.ChangedObjectIDIe != nil {
			qrChangedObjectIDIe = *o.ChangedObjectIDIe
		}
		qChangedObjectIDIe := qrChangedObjectIDIe
		if qChangedObjectIDIe != "" {

			if err := r.SetQueryParam("changed_object_id__ie", qChangedObjectIDIe); err != nil {
				return err
			}
		}
	}

	if o.ChangedObjectIDIew != nil {

		// query param changed_object_id__iew
		var qrChangedObjectIDIew string

		if o.ChangedObjectIDIew != nil {
			qrChangedObjectIDIew = *o.ChangedObjectIDIew
		}
		qChangedObjectIDIew := qrChangedObjectIDIew
		if qChangedObjectIDIew != "" {

			if err := r.SetQueryParam("changed_object_id__iew", qChangedObjectIDIew); err != nil {
				return err
			}
		}
	}

	if o.ChangedObjectIDIsw != nil {

		// query param changed_object_id__isw
		var qrChangedObjectIDIsw string

		if o.ChangedObjectIDIsw != nil {
			qrChangedObjectIDIsw = *o.ChangedObjectIDIsw
		}
		qChangedObjectIDIsw := qrChangedObjectIDIsw
		if qChangedObjectIDIsw != "" {

			if err := r.SetQueryParam("changed_object_id__isw", qChangedObjectIDIsw); err != nil {
				return err
			}
		}
	}

	if o.ChangedObjectIDn != nil {

		// query param changed_object_id__n
		var qrChangedObjectIDn string

		if o.ChangedObjectIDn != nil {
			qrChangedObjectIDn = *o.ChangedObjectIDn
		}
		qChangedObjectIDn := qrChangedObjectIDn
		if qChangedObjectIDn != "" {

			if err := r.SetQueryParam("changed_object_id__n", qChangedObjectIDn); err != nil {
				return err
			}
		}
	}

	if o.ChangedObjectIDNic != nil {

		// query param changed_object_id__nic
		var qrChangedObjectIDNic string

		if o.ChangedObjectIDNic != nil {
			qrChangedObjectIDNic = *o.ChangedObjectIDNic
		}
		qChangedObjectIDNic := qrChangedObjectIDNic
		if qChangedObjectIDNic != "" {

			if err := r.SetQueryParam("changed_object_id__nic", qChangedObjectIDNic); err != nil {
				return err
			}
		}
	}

	if o.ChangedObjectIDNie != nil {

		// query param changed_object_id__nie
		var qrChangedObjectIDNie string

		if o.ChangedObjectIDNie != nil {
			qrChangedObjectIDNie = *o.ChangedObjectIDNie
		}
		qChangedObjectIDNie := qrChangedObjectIDNie
		if qChangedObjectIDNie != "" {

			if err := r.SetQueryParam("changed_object_id__nie", qChangedObjectIDNie); err != nil {
				return err
			}
		}
	}

	if o.ChangedObjectIDNiew != nil {

		// query param changed_object_id__niew
		var qrChangedObjectIDNiew string

		if o.ChangedObjectIDNiew != nil {
			qrChangedObjectIDNiew = *o.ChangedObjectIDNiew
		}
		qChangedObjectIDNiew := qrChangedObjectIDNiew
		if qChangedObjectIDNiew != "" {

			if err := r.SetQueryParam("changed_object_id__niew", qChangedObjectIDNiew); err != nil {
				return err
			}
		}
	}

	if o.ChangedObjectIDNisw != nil {

		// query param changed_object_id__nisw
		var qrChangedObjectIDNisw string

		if o.ChangedObjectIDNisw != nil {
			qrChangedObjectIDNisw = *o.ChangedObjectIDNisw
		}
		qChangedObjectIDNisw := qrChangedObjectIDNisw
		if qChangedObjectIDNisw != "" {

			if err := r.SetQueryParam("changed_object_id__nisw", qChangedObjectIDNisw); err != nil {
				return err
			}
		}
	}

	if o.ChangedObjectType != nil {

		// query param changed_object_type
		var qrChangedObjectType string

		if o.ChangedObjectType != nil {
			qrChangedObjectType = *o.ChangedObjectType
		}
		qChangedObjectType := qrChangedObjectType
		if qChangedObjectType != "" {

			if err := r.SetQueryParam("changed_object_type", qChangedObjectType); err != nil {
				return err
			}
		}
	}

	if o.ChangedObjectTypen != nil {

		// query param changed_object_type__n
		var qrChangedObjectTypen string

		if o.ChangedObjectTypen != nil {
			qrChangedObjectTypen = *o.ChangedObjectTypen
		}
		qChangedObjectTypen := qrChangedObjectTypen
		if qChangedObjectTypen != "" {

			if err := r.SetQueryParam("changed_object_type__n", qChangedObjectTypen); err != nil {
				return err
			}
		}
	}

	if o.ChangedObjectTypeID != nil {

		// query param changed_object_type_id
		var qrChangedObjectTypeID string

		if o.ChangedObjectTypeID != nil {
			qrChangedObjectTypeID = *o.ChangedObjectTypeID
		}
		qChangedObjectTypeID := qrChangedObjectTypeID
		if qChangedObjectTypeID != "" {

			if err := r.SetQueryParam("changed_object_type_id", qChangedObjectTypeID); err != nil {
				return err
			}
		}
	}

	if o.ChangedObjectTypeIDn != nil {

		// query param changed_object_type_id__n
		var qrChangedObjectTypeIDn string

		if o.ChangedObjectTypeIDn != nil {
			qrChangedObjectTypeIDn = *o.ChangedObjectTypeIDn
		}
		qChangedObjectTypeIDn := qrChangedObjectTypeIDn
		if qChangedObjectTypeIDn != "" {

			if err := r.SetQueryParam("changed_object_type_id__n", qChangedObjectTypeIDn); err != nil {
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

	if o.ObjectRepr != nil {

		// query param object_repr
		var qrObjectRepr string

		if o.ObjectRepr != nil {
			qrObjectRepr = *o.ObjectRepr
		}
		qObjectRepr := qrObjectRepr
		if qObjectRepr != "" {

			if err := r.SetQueryParam("object_repr", qObjectRepr); err != nil {
				return err
			}
		}
	}

	if o.ObjectReprIc != nil {

		// query param object_repr__ic
		var qrObjectReprIc string

		if o.ObjectReprIc != nil {
			qrObjectReprIc = *o.ObjectReprIc
		}
		qObjectReprIc := qrObjectReprIc
		if qObjectReprIc != "" {

			if err := r.SetQueryParam("object_repr__ic", qObjectReprIc); err != nil {
				return err
			}
		}
	}

	if o.ObjectReprIe != nil {

		// query param object_repr__ie
		var qrObjectReprIe string

		if o.ObjectReprIe != nil {
			qrObjectReprIe = *o.ObjectReprIe
		}
		qObjectReprIe := qrObjectReprIe
		if qObjectReprIe != "" {

			if err := r.SetQueryParam("object_repr__ie", qObjectReprIe); err != nil {
				return err
			}
		}
	}

	if o.ObjectReprIew != nil {

		// query param object_repr__iew
		var qrObjectReprIew string

		if o.ObjectReprIew != nil {
			qrObjectReprIew = *o.ObjectReprIew
		}
		qObjectReprIew := qrObjectReprIew
		if qObjectReprIew != "" {

			if err := r.SetQueryParam("object_repr__iew", qObjectReprIew); err != nil {
				return err
			}
		}
	}

	if o.ObjectReprIsw != nil {

		// query param object_repr__isw
		var qrObjectReprIsw string

		if o.ObjectReprIsw != nil {
			qrObjectReprIsw = *o.ObjectReprIsw
		}
		qObjectReprIsw := qrObjectReprIsw
		if qObjectReprIsw != "" {

			if err := r.SetQueryParam("object_repr__isw", qObjectReprIsw); err != nil {
				return err
			}
		}
	}

	if o.ObjectReprn != nil {

		// query param object_repr__n
		var qrObjectReprn string

		if o.ObjectReprn != nil {
			qrObjectReprn = *o.ObjectReprn
		}
		qObjectReprn := qrObjectReprn
		if qObjectReprn != "" {

			if err := r.SetQueryParam("object_repr__n", qObjectReprn); err != nil {
				return err
			}
		}
	}

	if o.ObjectReprNic != nil {

		// query param object_repr__nic
		var qrObjectReprNic string

		if o.ObjectReprNic != nil {
			qrObjectReprNic = *o.ObjectReprNic
		}
		qObjectReprNic := qrObjectReprNic
		if qObjectReprNic != "" {

			if err := r.SetQueryParam("object_repr__nic", qObjectReprNic); err != nil {
				return err
			}
		}
	}

	if o.ObjectReprNie != nil {

		// query param object_repr__nie
		var qrObjectReprNie string

		if o.ObjectReprNie != nil {
			qrObjectReprNie = *o.ObjectReprNie
		}
		qObjectReprNie := qrObjectReprNie
		if qObjectReprNie != "" {

			if err := r.SetQueryParam("object_repr__nie", qObjectReprNie); err != nil {
				return err
			}
		}
	}

	if o.ObjectReprNiew != nil {

		// query param object_repr__niew
		var qrObjectReprNiew string

		if o.ObjectReprNiew != nil {
			qrObjectReprNiew = *o.ObjectReprNiew
		}
		qObjectReprNiew := qrObjectReprNiew
		if qObjectReprNiew != "" {

			if err := r.SetQueryParam("object_repr__niew", qObjectReprNiew); err != nil {
				return err
			}
		}
	}

	if o.ObjectReprNisw != nil {

		// query param object_repr__nisw
		var qrObjectReprNisw string

		if o.ObjectReprNisw != nil {
			qrObjectReprNisw = *o.ObjectReprNisw
		}
		qObjectReprNisw := qrObjectReprNisw
		if qObjectReprNisw != "" {

			if err := r.SetQueryParam("object_repr__nisw", qObjectReprNisw); err != nil {
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

	if o.Q != nil {

		// query param q
		var qrQ string

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}
	}

	if o.RequestID != nil {

		// query param request_id
		var qrRequestID string

		if o.RequestID != nil {
			qrRequestID = *o.RequestID
		}
		qRequestID := qrRequestID
		if qRequestID != "" {

			if err := r.SetQueryParam("request_id", qRequestID); err != nil {
				return err
			}
		}
	}

	if o.RequestIDIc != nil {

		// query param request_id__ic
		var qrRequestIDIc string

		if o.RequestIDIc != nil {
			qrRequestIDIc = *o.RequestIDIc
		}
		qRequestIDIc := qrRequestIDIc
		if qRequestIDIc != "" {

			if err := r.SetQueryParam("request_id__ic", qRequestIDIc); err != nil {
				return err
			}
		}
	}

	if o.RequestIDIe != nil {

		// query param request_id__ie
		var qrRequestIDIe string

		if o.RequestIDIe != nil {
			qrRequestIDIe = *o.RequestIDIe
		}
		qRequestIDIe := qrRequestIDIe
		if qRequestIDIe != "" {

			if err := r.SetQueryParam("request_id__ie", qRequestIDIe); err != nil {
				return err
			}
		}
	}

	if o.RequestIDIew != nil {

		// query param request_id__iew
		var qrRequestIDIew string

		if o.RequestIDIew != nil {
			qrRequestIDIew = *o.RequestIDIew
		}
		qRequestIDIew := qrRequestIDIew
		if qRequestIDIew != "" {

			if err := r.SetQueryParam("request_id__iew", qRequestIDIew); err != nil {
				return err
			}
		}
	}

	if o.RequestIDIsw != nil {

		// query param request_id__isw
		var qrRequestIDIsw string

		if o.RequestIDIsw != nil {
			qrRequestIDIsw = *o.RequestIDIsw
		}
		qRequestIDIsw := qrRequestIDIsw
		if qRequestIDIsw != "" {

			if err := r.SetQueryParam("request_id__isw", qRequestIDIsw); err != nil {
				return err
			}
		}
	}

	if o.RequestIDn != nil {

		// query param request_id__n
		var qrRequestIDn string

		if o.RequestIDn != nil {
			qrRequestIDn = *o.RequestIDn
		}
		qRequestIDn := qrRequestIDn
		if qRequestIDn != "" {

			if err := r.SetQueryParam("request_id__n", qRequestIDn); err != nil {
				return err
			}
		}
	}

	if o.RequestIDNic != nil {

		// query param request_id__nic
		var qrRequestIDNic string

		if o.RequestIDNic != nil {
			qrRequestIDNic = *o.RequestIDNic
		}
		qRequestIDNic := qrRequestIDNic
		if qRequestIDNic != "" {

			if err := r.SetQueryParam("request_id__nic", qRequestIDNic); err != nil {
				return err
			}
		}
	}

	if o.RequestIDNie != nil {

		// query param request_id__nie
		var qrRequestIDNie string

		if o.RequestIDNie != nil {
			qrRequestIDNie = *o.RequestIDNie
		}
		qRequestIDNie := qrRequestIDNie
		if qRequestIDNie != "" {

			if err := r.SetQueryParam("request_id__nie", qRequestIDNie); err != nil {
				return err
			}
		}
	}

	if o.RequestIDNiew != nil {

		// query param request_id__niew
		var qrRequestIDNiew string

		if o.RequestIDNiew != nil {
			qrRequestIDNiew = *o.RequestIDNiew
		}
		qRequestIDNiew := qrRequestIDNiew
		if qRequestIDNiew != "" {

			if err := r.SetQueryParam("request_id__niew", qRequestIDNiew); err != nil {
				return err
			}
		}
	}

	if o.RequestIDNisw != nil {

		// query param request_id__nisw
		var qrRequestIDNisw string

		if o.RequestIDNisw != nil {
			qrRequestIDNisw = *o.RequestIDNisw
		}
		qRequestIDNisw := qrRequestIDNisw
		if qRequestIDNisw != "" {

			if err := r.SetQueryParam("request_id__nisw", qRequestIDNisw); err != nil {
				return err
			}
		}
	}

	if o.Time != nil {

		// query param time
		var qrTime string

		if o.Time != nil {
			qrTime = *o.Time
		}
		qTime := qrTime
		if qTime != "" {

			if err := r.SetQueryParam("time", qTime); err != nil {
				return err
			}
		}
	}

	if o.User != nil {

		// query param user
		var qrUser string

		if o.User != nil {
			qrUser = *o.User
		}
		qUser := qrUser
		if qUser != "" {

			if err := r.SetQueryParam("user", qUser); err != nil {
				return err
			}
		}
	}

	if o.Usern != nil {

		// query param user__n
		var qrUsern string

		if o.Usern != nil {
			qrUsern = *o.Usern
		}
		qUsern := qrUsern
		if qUsern != "" {

			if err := r.SetQueryParam("user__n", qUsern); err != nil {
				return err
			}
		}
	}

	if o.UserID != nil {

		// query param user_id
		var qrUserID string

		if o.UserID != nil {
			qrUserID = *o.UserID
		}
		qUserID := qrUserID
		if qUserID != "" {

			if err := r.SetQueryParam("user_id", qUserID); err != nil {
				return err
			}
		}
	}

	if o.UserIDn != nil {

		// query param user_id__n
		var qrUserIDn string

		if o.UserIDn != nil {
			qrUserIDn = *o.UserIDn
		}
		qUserIDn := qrUserIDn
		if qUserIDn != "" {

			if err := r.SetQueryParam("user_id__n", qUserIDn); err != nil {
				return err
			}
		}
	}

	if o.UserName != nil {

		// query param user_name
		var qrUserName string

		if o.UserName != nil {
			qrUserName = *o.UserName
		}
		qUserName := qrUserName
		if qUserName != "" {

			if err := r.SetQueryParam("user_name", qUserName); err != nil {
				return err
			}
		}
	}

	if o.UserNameIc != nil {

		// query param user_name__ic
		var qrUserNameIc string

		if o.UserNameIc != nil {
			qrUserNameIc = *o.UserNameIc
		}
		qUserNameIc := qrUserNameIc
		if qUserNameIc != "" {

			if err := r.SetQueryParam("user_name__ic", qUserNameIc); err != nil {
				return err
			}
		}
	}

	if o.UserNameIe != nil {

		// query param user_name__ie
		var qrUserNameIe string

		if o.UserNameIe != nil {
			qrUserNameIe = *o.UserNameIe
		}
		qUserNameIe := qrUserNameIe
		if qUserNameIe != "" {

			if err := r.SetQueryParam("user_name__ie", qUserNameIe); err != nil {
				return err
			}
		}
	}

	if o.UserNameIew != nil {

		// query param user_name__iew
		var qrUserNameIew string

		if o.UserNameIew != nil {
			qrUserNameIew = *o.UserNameIew
		}
		qUserNameIew := qrUserNameIew
		if qUserNameIew != "" {

			if err := r.SetQueryParam("user_name__iew", qUserNameIew); err != nil {
				return err
			}
		}
	}

	if o.UserNameIsw != nil {

		// query param user_name__isw
		var qrUserNameIsw string

		if o.UserNameIsw != nil {
			qrUserNameIsw = *o.UserNameIsw
		}
		qUserNameIsw := qrUserNameIsw
		if qUserNameIsw != "" {

			if err := r.SetQueryParam("user_name__isw", qUserNameIsw); err != nil {
				return err
			}
		}
	}

	if o.UserNamen != nil {

		// query param user_name__n
		var qrUserNamen string

		if o.UserNamen != nil {
			qrUserNamen = *o.UserNamen
		}
		qUserNamen := qrUserNamen
		if qUserNamen != "" {

			if err := r.SetQueryParam("user_name__n", qUserNamen); err != nil {
				return err
			}
		}
	}

	if o.UserNameNic != nil {

		// query param user_name__nic
		var qrUserNameNic string

		if o.UserNameNic != nil {
			qrUserNameNic = *o.UserNameNic
		}
		qUserNameNic := qrUserNameNic
		if qUserNameNic != "" {

			if err := r.SetQueryParam("user_name__nic", qUserNameNic); err != nil {
				return err
			}
		}
	}

	if o.UserNameNie != nil {

		// query param user_name__nie
		var qrUserNameNie string

		if o.UserNameNie != nil {
			qrUserNameNie = *o.UserNameNie
		}
		qUserNameNie := qrUserNameNie
		if qUserNameNie != "" {

			if err := r.SetQueryParam("user_name__nie", qUserNameNie); err != nil {
				return err
			}
		}
	}

	if o.UserNameNiew != nil {

		// query param user_name__niew
		var qrUserNameNiew string

		if o.UserNameNiew != nil {
			qrUserNameNiew = *o.UserNameNiew
		}
		qUserNameNiew := qrUserNameNiew
		if qUserNameNiew != "" {

			if err := r.SetQueryParam("user_name__niew", qUserNameNiew); err != nil {
				return err
			}
		}
	}

	if o.UserNameNisw != nil {

		// query param user_name__nisw
		var qrUserNameNisw string

		if o.UserNameNisw != nil {
			qrUserNameNisw = *o.UserNameNisw
		}
		qUserNameNisw := qrUserNameNisw
		if qUserNameNisw != "" {

			if err := r.SetQueryParam("user_name__nisw", qUserNameNisw); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

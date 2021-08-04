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

// NewExtrasImageAttachmentsListParams creates a new ExtrasImageAttachmentsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasImageAttachmentsListParams() *ExtrasImageAttachmentsListParams {
	return &ExtrasImageAttachmentsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasImageAttachmentsListParamsWithTimeout creates a new ExtrasImageAttachmentsListParams object
// with the ability to set a timeout on a request.
func NewExtrasImageAttachmentsListParamsWithTimeout(timeout time.Duration) *ExtrasImageAttachmentsListParams {
	return &ExtrasImageAttachmentsListParams{
		timeout: timeout,
	}
}

// NewExtrasImageAttachmentsListParamsWithContext creates a new ExtrasImageAttachmentsListParams object
// with the ability to set a context for a request.
func NewExtrasImageAttachmentsListParamsWithContext(ctx context.Context) *ExtrasImageAttachmentsListParams {
	return &ExtrasImageAttachmentsListParams{
		Context: ctx,
	}
}

// NewExtrasImageAttachmentsListParamsWithHTTPClient creates a new ExtrasImageAttachmentsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasImageAttachmentsListParamsWithHTTPClient(client *http.Client) *ExtrasImageAttachmentsListParams {
	return &ExtrasImageAttachmentsListParams{
		HTTPClient: client,
	}
}

/* ExtrasImageAttachmentsListParams contains all the parameters to send to the API endpoint
   for the extras image attachments list operation.

   Typically these are written to a http.Request.
*/
type ExtrasImageAttachmentsListParams struct {

	// ContentType.
	ContentType *string

	// ContentTypen.
	ContentTypen *string

	// ContentTypeID.
	ContentTypeID *string

	// ContentTypeIDn.
	ContentTypeIDn *string

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

	// Name.
	Name *string

	// NameIc.
	NameIc *string

	// NameIe.
	NameIe *string

	// NameIew.
	NameIew *string

	// NameIsw.
	NameIsw *string

	// Namen.
	Namen *string

	// NameNic.
	NameNic *string

	// NameNie.
	NameNie *string

	// NameNiew.
	NameNiew *string

	// NameNisw.
	NameNisw *string

	// ObjectID.
	ObjectID *string

	// ObjectIDIc.
	ObjectIDIc *string

	// ObjectIDIe.
	ObjectIDIe *string

	// ObjectIDIew.
	ObjectIDIew *string

	// ObjectIDIsw.
	ObjectIDIsw *string

	// ObjectIDn.
	ObjectIDn *string

	// ObjectIDNic.
	ObjectIDNic *string

	// ObjectIDNie.
	ObjectIDNie *string

	// ObjectIDNiew.
	ObjectIDNiew *string

	// ObjectIDNisw.
	ObjectIDNisw *string

	/* Offset.

	   The initial index from which to return the results.
	*/
	Offset *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras image attachments list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasImageAttachmentsListParams) WithDefaults() *ExtrasImageAttachmentsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras image attachments list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasImageAttachmentsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithTimeout(timeout time.Duration) *ExtrasImageAttachmentsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithContext(ctx context.Context) *ExtrasImageAttachmentsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithHTTPClient(client *http.Client) *ExtrasImageAttachmentsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentType adds the contentType to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithContentType(contentType *string) *ExtrasImageAttachmentsListParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetContentType(contentType *string) {
	o.ContentType = contentType
}

// WithContentTypen adds the contentTypen to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithContentTypen(contentTypen *string) *ExtrasImageAttachmentsListParams {
	o.SetContentTypen(contentTypen)
	return o
}

// SetContentTypen adds the contentTypeN to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetContentTypen(contentTypen *string) {
	o.ContentTypen = contentTypen
}

// WithContentTypeID adds the contentTypeID to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithContentTypeID(contentTypeID *string) *ExtrasImageAttachmentsListParams {
	o.SetContentTypeID(contentTypeID)
	return o
}

// SetContentTypeID adds the contentTypeId to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetContentTypeID(contentTypeID *string) {
	o.ContentTypeID = contentTypeID
}

// WithContentTypeIDn adds the contentTypeIDn to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithContentTypeIDn(contentTypeIDn *string) *ExtrasImageAttachmentsListParams {
	o.SetContentTypeIDn(contentTypeIDn)
	return o
}

// SetContentTypeIDn adds the contentTypeIdN to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetContentTypeIDn(contentTypeIDn *string) {
	o.ContentTypeIDn = contentTypeIDn
}

// WithID adds the id to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithID(id *string) *ExtrasImageAttachmentsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithIDIc(iDIc *string) *ExtrasImageAttachmentsListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithIDIe(iDIe *string) *ExtrasImageAttachmentsListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithIDIew(iDIew *string) *ExtrasImageAttachmentsListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithIDIsw(iDIsw *string) *ExtrasImageAttachmentsListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithIDn(iDn *string) *ExtrasImageAttachmentsListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithIDNic(iDNic *string) *ExtrasImageAttachmentsListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithIDNie(iDNie *string) *ExtrasImageAttachmentsListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithIDNiew(iDNiew *string) *ExtrasImageAttachmentsListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithIDNisw(iDNisw *string) *ExtrasImageAttachmentsListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithLimit adds the limit to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithLimit(limit *int64) *ExtrasImageAttachmentsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithName(name *string) *ExtrasImageAttachmentsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithNameIc(nameIc *string) *ExtrasImageAttachmentsListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithNameIe(nameIe *string) *ExtrasImageAttachmentsListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithNameIew(nameIew *string) *ExtrasImageAttachmentsListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithNameIsw(nameIsw *string) *ExtrasImageAttachmentsListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithNamen(namen *string) *ExtrasImageAttachmentsListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithNameNic(nameNic *string) *ExtrasImageAttachmentsListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithNameNie(nameNie *string) *ExtrasImageAttachmentsListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithNameNiew(nameNiew *string) *ExtrasImageAttachmentsListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithNameNisw(nameNisw *string) *ExtrasImageAttachmentsListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithObjectID adds the objectID to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithObjectID(objectID *string) *ExtrasImageAttachmentsListParams {
	o.SetObjectID(objectID)
	return o
}

// SetObjectID adds the objectId to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetObjectID(objectID *string) {
	o.ObjectID = objectID
}

// WithObjectIDIc adds the objectIDIc to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithObjectIDIc(objectIDIc *string) *ExtrasImageAttachmentsListParams {
	o.SetObjectIDIc(objectIDIc)
	return o
}

// SetObjectIDIc adds the objectIdIc to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetObjectIDIc(objectIDIc *string) {
	o.ObjectIDIc = objectIDIc
}

// WithObjectIDIe adds the objectIDIe to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithObjectIDIe(objectIDIe *string) *ExtrasImageAttachmentsListParams {
	o.SetObjectIDIe(objectIDIe)
	return o
}

// SetObjectIDIe adds the objectIdIe to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetObjectIDIe(objectIDIe *string) {
	o.ObjectIDIe = objectIDIe
}

// WithObjectIDIew adds the objectIDIew to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithObjectIDIew(objectIDIew *string) *ExtrasImageAttachmentsListParams {
	o.SetObjectIDIew(objectIDIew)
	return o
}

// SetObjectIDIew adds the objectIdIew to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetObjectIDIew(objectIDIew *string) {
	o.ObjectIDIew = objectIDIew
}

// WithObjectIDIsw adds the objectIDIsw to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithObjectIDIsw(objectIDIsw *string) *ExtrasImageAttachmentsListParams {
	o.SetObjectIDIsw(objectIDIsw)
	return o
}

// SetObjectIDIsw adds the objectIdIsw to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetObjectIDIsw(objectIDIsw *string) {
	o.ObjectIDIsw = objectIDIsw
}

// WithObjectIDn adds the objectIDn to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithObjectIDn(objectIDn *string) *ExtrasImageAttachmentsListParams {
	o.SetObjectIDn(objectIDn)
	return o
}

// SetObjectIDn adds the objectIdN to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetObjectIDn(objectIDn *string) {
	o.ObjectIDn = objectIDn
}

// WithObjectIDNic adds the objectIDNic to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithObjectIDNic(objectIDNic *string) *ExtrasImageAttachmentsListParams {
	o.SetObjectIDNic(objectIDNic)
	return o
}

// SetObjectIDNic adds the objectIdNic to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetObjectIDNic(objectIDNic *string) {
	o.ObjectIDNic = objectIDNic
}

// WithObjectIDNie adds the objectIDNie to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithObjectIDNie(objectIDNie *string) *ExtrasImageAttachmentsListParams {
	o.SetObjectIDNie(objectIDNie)
	return o
}

// SetObjectIDNie adds the objectIdNie to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetObjectIDNie(objectIDNie *string) {
	o.ObjectIDNie = objectIDNie
}

// WithObjectIDNiew adds the objectIDNiew to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithObjectIDNiew(objectIDNiew *string) *ExtrasImageAttachmentsListParams {
	o.SetObjectIDNiew(objectIDNiew)
	return o
}

// SetObjectIDNiew adds the objectIdNiew to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetObjectIDNiew(objectIDNiew *string) {
	o.ObjectIDNiew = objectIDNiew
}

// WithObjectIDNisw adds the objectIDNisw to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithObjectIDNisw(objectIDNisw *string) *ExtrasImageAttachmentsListParams {
	o.SetObjectIDNisw(objectIDNisw)
	return o
}

// SetObjectIDNisw adds the objectIdNisw to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetObjectIDNisw(objectIDNisw *string) {
	o.ObjectIDNisw = objectIDNisw
}

// WithOffset adds the offset to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) WithOffset(offset *int64) *ExtrasImageAttachmentsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the extras image attachments list params
func (o *ExtrasImageAttachmentsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasImageAttachmentsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ContentType != nil {

		// query param content_type
		var qrContentType string

		if o.ContentType != nil {
			qrContentType = *o.ContentType
		}
		qContentType := qrContentType
		if qContentType != "" {

			if err := r.SetQueryParam("content_type", qContentType); err != nil {
				return err
			}
		}
	}

	if o.ContentTypen != nil {

		// query param content_type__n
		var qrContentTypen string

		if o.ContentTypen != nil {
			qrContentTypen = *o.ContentTypen
		}
		qContentTypen := qrContentTypen
		if qContentTypen != "" {

			if err := r.SetQueryParam("content_type__n", qContentTypen); err != nil {
				return err
			}
		}
	}

	if o.ContentTypeID != nil {

		// query param content_type_id
		var qrContentTypeID string

		if o.ContentTypeID != nil {
			qrContentTypeID = *o.ContentTypeID
		}
		qContentTypeID := qrContentTypeID
		if qContentTypeID != "" {

			if err := r.SetQueryParam("content_type_id", qContentTypeID); err != nil {
				return err
			}
		}
	}

	if o.ContentTypeIDn != nil {

		// query param content_type_id__n
		var qrContentTypeIDn string

		if o.ContentTypeIDn != nil {
			qrContentTypeIDn = *o.ContentTypeIDn
		}
		qContentTypeIDn := qrContentTypeIDn
		if qContentTypeIDn != "" {

			if err := r.SetQueryParam("content_type_id__n", qContentTypeIDn); err != nil {
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

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.NameIc != nil {

		// query param name__ic
		var qrNameIc string

		if o.NameIc != nil {
			qrNameIc = *o.NameIc
		}
		qNameIc := qrNameIc
		if qNameIc != "" {

			if err := r.SetQueryParam("name__ic", qNameIc); err != nil {
				return err
			}
		}
	}

	if o.NameIe != nil {

		// query param name__ie
		var qrNameIe string

		if o.NameIe != nil {
			qrNameIe = *o.NameIe
		}
		qNameIe := qrNameIe
		if qNameIe != "" {

			if err := r.SetQueryParam("name__ie", qNameIe); err != nil {
				return err
			}
		}
	}

	if o.NameIew != nil {

		// query param name__iew
		var qrNameIew string

		if o.NameIew != nil {
			qrNameIew = *o.NameIew
		}
		qNameIew := qrNameIew
		if qNameIew != "" {

			if err := r.SetQueryParam("name__iew", qNameIew); err != nil {
				return err
			}
		}
	}

	if o.NameIsw != nil {

		// query param name__isw
		var qrNameIsw string

		if o.NameIsw != nil {
			qrNameIsw = *o.NameIsw
		}
		qNameIsw := qrNameIsw
		if qNameIsw != "" {

			if err := r.SetQueryParam("name__isw", qNameIsw); err != nil {
				return err
			}
		}
	}

	if o.Namen != nil {

		// query param name__n
		var qrNamen string

		if o.Namen != nil {
			qrNamen = *o.Namen
		}
		qNamen := qrNamen
		if qNamen != "" {

			if err := r.SetQueryParam("name__n", qNamen); err != nil {
				return err
			}
		}
	}

	if o.NameNic != nil {

		// query param name__nic
		var qrNameNic string

		if o.NameNic != nil {
			qrNameNic = *o.NameNic
		}
		qNameNic := qrNameNic
		if qNameNic != "" {

			if err := r.SetQueryParam("name__nic", qNameNic); err != nil {
				return err
			}
		}
	}

	if o.NameNie != nil {

		// query param name__nie
		var qrNameNie string

		if o.NameNie != nil {
			qrNameNie = *o.NameNie
		}
		qNameNie := qrNameNie
		if qNameNie != "" {

			if err := r.SetQueryParam("name__nie", qNameNie); err != nil {
				return err
			}
		}
	}

	if o.NameNiew != nil {

		// query param name__niew
		var qrNameNiew string

		if o.NameNiew != nil {
			qrNameNiew = *o.NameNiew
		}
		qNameNiew := qrNameNiew
		if qNameNiew != "" {

			if err := r.SetQueryParam("name__niew", qNameNiew); err != nil {
				return err
			}
		}
	}

	if o.NameNisw != nil {

		// query param name__nisw
		var qrNameNisw string

		if o.NameNisw != nil {
			qrNameNisw = *o.NameNisw
		}
		qNameNisw := qrNameNisw
		if qNameNisw != "" {

			if err := r.SetQueryParam("name__nisw", qNameNisw); err != nil {
				return err
			}
		}
	}

	if o.ObjectID != nil {

		// query param object_id
		var qrObjectID string

		if o.ObjectID != nil {
			qrObjectID = *o.ObjectID
		}
		qObjectID := qrObjectID
		if qObjectID != "" {

			if err := r.SetQueryParam("object_id", qObjectID); err != nil {
				return err
			}
		}
	}

	if o.ObjectIDIc != nil {

		// query param object_id__ic
		var qrObjectIDIc string

		if o.ObjectIDIc != nil {
			qrObjectIDIc = *o.ObjectIDIc
		}
		qObjectIDIc := qrObjectIDIc
		if qObjectIDIc != "" {

			if err := r.SetQueryParam("object_id__ic", qObjectIDIc); err != nil {
				return err
			}
		}
	}

	if o.ObjectIDIe != nil {

		// query param object_id__ie
		var qrObjectIDIe string

		if o.ObjectIDIe != nil {
			qrObjectIDIe = *o.ObjectIDIe
		}
		qObjectIDIe := qrObjectIDIe
		if qObjectIDIe != "" {

			if err := r.SetQueryParam("object_id__ie", qObjectIDIe); err != nil {
				return err
			}
		}
	}

	if o.ObjectIDIew != nil {

		// query param object_id__iew
		var qrObjectIDIew string

		if o.ObjectIDIew != nil {
			qrObjectIDIew = *o.ObjectIDIew
		}
		qObjectIDIew := qrObjectIDIew
		if qObjectIDIew != "" {

			if err := r.SetQueryParam("object_id__iew", qObjectIDIew); err != nil {
				return err
			}
		}
	}

	if o.ObjectIDIsw != nil {

		// query param object_id__isw
		var qrObjectIDIsw string

		if o.ObjectIDIsw != nil {
			qrObjectIDIsw = *o.ObjectIDIsw
		}
		qObjectIDIsw := qrObjectIDIsw
		if qObjectIDIsw != "" {

			if err := r.SetQueryParam("object_id__isw", qObjectIDIsw); err != nil {
				return err
			}
		}
	}

	if o.ObjectIDn != nil {

		// query param object_id__n
		var qrObjectIDn string

		if o.ObjectIDn != nil {
			qrObjectIDn = *o.ObjectIDn
		}
		qObjectIDn := qrObjectIDn
		if qObjectIDn != "" {

			if err := r.SetQueryParam("object_id__n", qObjectIDn); err != nil {
				return err
			}
		}
	}

	if o.ObjectIDNic != nil {

		// query param object_id__nic
		var qrObjectIDNic string

		if o.ObjectIDNic != nil {
			qrObjectIDNic = *o.ObjectIDNic
		}
		qObjectIDNic := qrObjectIDNic
		if qObjectIDNic != "" {

			if err := r.SetQueryParam("object_id__nic", qObjectIDNic); err != nil {
				return err
			}
		}
	}

	if o.ObjectIDNie != nil {

		// query param object_id__nie
		var qrObjectIDNie string

		if o.ObjectIDNie != nil {
			qrObjectIDNie = *o.ObjectIDNie
		}
		qObjectIDNie := qrObjectIDNie
		if qObjectIDNie != "" {

			if err := r.SetQueryParam("object_id__nie", qObjectIDNie); err != nil {
				return err
			}
		}
	}

	if o.ObjectIDNiew != nil {

		// query param object_id__niew
		var qrObjectIDNiew string

		if o.ObjectIDNiew != nil {
			qrObjectIDNiew = *o.ObjectIDNiew
		}
		qObjectIDNiew := qrObjectIDNiew
		if qObjectIDNiew != "" {

			if err := r.SetQueryParam("object_id__niew", qObjectIDNiew); err != nil {
				return err
			}
		}
	}

	if o.ObjectIDNisw != nil {

		// query param object_id__nisw
		var qrObjectIDNisw string

		if o.ObjectIDNisw != nil {
			qrObjectIDNisw = *o.ObjectIDNisw
		}
		qObjectIDNisw := qrObjectIDNisw
		if qObjectIDNisw != "" {

			if err := r.SetQueryParam("object_id__nisw", qObjectIDNisw); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

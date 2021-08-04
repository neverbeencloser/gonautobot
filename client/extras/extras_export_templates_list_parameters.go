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

// NewExtrasExportTemplatesListParams creates a new ExtrasExportTemplatesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasExportTemplatesListParams() *ExtrasExportTemplatesListParams {
	return &ExtrasExportTemplatesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasExportTemplatesListParamsWithTimeout creates a new ExtrasExportTemplatesListParams object
// with the ability to set a timeout on a request.
func NewExtrasExportTemplatesListParamsWithTimeout(timeout time.Duration) *ExtrasExportTemplatesListParams {
	return &ExtrasExportTemplatesListParams{
		timeout: timeout,
	}
}

// NewExtrasExportTemplatesListParamsWithContext creates a new ExtrasExportTemplatesListParams object
// with the ability to set a context for a request.
func NewExtrasExportTemplatesListParamsWithContext(ctx context.Context) *ExtrasExportTemplatesListParams {
	return &ExtrasExportTemplatesListParams{
		Context: ctx,
	}
}

// NewExtrasExportTemplatesListParamsWithHTTPClient creates a new ExtrasExportTemplatesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasExportTemplatesListParamsWithHTTPClient(client *http.Client) *ExtrasExportTemplatesListParams {
	return &ExtrasExportTemplatesListParams{
		HTTPClient: client,
	}
}

/* ExtrasExportTemplatesListParams contains all the parameters to send to the API endpoint
   for the extras export templates list operation.

   Typically these are written to a http.Request.
*/
type ExtrasExportTemplatesListParams struct {

	// ContentType.
	ContentType *string

	// ContentTypen.
	ContentTypen *string

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

	/* Offset.

	   The initial index from which to return the results.
	*/
	Offset *int64

	// OwnerContentType.
	OwnerContentType *string

	// OwnerContentTypen.
	OwnerContentTypen *string

	// OwnerObjectID.
	OwnerObjectID *string

	// OwnerObjectIDIc.
	OwnerObjectIDIc *string

	// OwnerObjectIDIe.
	OwnerObjectIDIe *string

	// OwnerObjectIDIew.
	OwnerObjectIDIew *string

	// OwnerObjectIDIsw.
	OwnerObjectIDIsw *string

	// OwnerObjectIDn.
	OwnerObjectIDn *string

	// OwnerObjectIDNic.
	OwnerObjectIDNic *string

	// OwnerObjectIDNie.
	OwnerObjectIDNie *string

	// OwnerObjectIDNiew.
	OwnerObjectIDNiew *string

	// OwnerObjectIDNisw.
	OwnerObjectIDNisw *string

	// Q.
	Q *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras export templates list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasExportTemplatesListParams) WithDefaults() *ExtrasExportTemplatesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras export templates list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasExportTemplatesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithTimeout(timeout time.Duration) *ExtrasExportTemplatesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithContext(ctx context.Context) *ExtrasExportTemplatesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithHTTPClient(client *http.Client) *ExtrasExportTemplatesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentType adds the contentType to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithContentType(contentType *string) *ExtrasExportTemplatesListParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetContentType(contentType *string) {
	o.ContentType = contentType
}

// WithContentTypen adds the contentTypen to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithContentTypen(contentTypen *string) *ExtrasExportTemplatesListParams {
	o.SetContentTypen(contentTypen)
	return o
}

// SetContentTypen adds the contentTypeN to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetContentTypen(contentTypen *string) {
	o.ContentTypen = contentTypen
}

// WithID adds the id to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithID(id *string) *ExtrasExportTemplatesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithIDIc(iDIc *string) *ExtrasExportTemplatesListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithIDIe(iDIe *string) *ExtrasExportTemplatesListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithIDIew(iDIew *string) *ExtrasExportTemplatesListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithIDIsw(iDIsw *string) *ExtrasExportTemplatesListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithIDn(iDn *string) *ExtrasExportTemplatesListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithIDNic(iDNic *string) *ExtrasExportTemplatesListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithIDNie(iDNie *string) *ExtrasExportTemplatesListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithIDNiew(iDNiew *string) *ExtrasExportTemplatesListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithIDNisw(iDNisw *string) *ExtrasExportTemplatesListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithLimit adds the limit to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithLimit(limit *int64) *ExtrasExportTemplatesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithName(name *string) *ExtrasExportTemplatesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithNameIc(nameIc *string) *ExtrasExportTemplatesListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithNameIe(nameIe *string) *ExtrasExportTemplatesListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithNameIew(nameIew *string) *ExtrasExportTemplatesListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithNameIsw(nameIsw *string) *ExtrasExportTemplatesListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithNamen(namen *string) *ExtrasExportTemplatesListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithNameNic(nameNic *string) *ExtrasExportTemplatesListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithNameNie(nameNie *string) *ExtrasExportTemplatesListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithNameNiew(nameNiew *string) *ExtrasExportTemplatesListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithNameNisw(nameNisw *string) *ExtrasExportTemplatesListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithOffset(offset *int64) *ExtrasExportTemplatesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOwnerContentType adds the ownerContentType to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithOwnerContentType(ownerContentType *string) *ExtrasExportTemplatesListParams {
	o.SetOwnerContentType(ownerContentType)
	return o
}

// SetOwnerContentType adds the ownerContentType to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetOwnerContentType(ownerContentType *string) {
	o.OwnerContentType = ownerContentType
}

// WithOwnerContentTypen adds the ownerContentTypen to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithOwnerContentTypen(ownerContentTypen *string) *ExtrasExportTemplatesListParams {
	o.SetOwnerContentTypen(ownerContentTypen)
	return o
}

// SetOwnerContentTypen adds the ownerContentTypeN to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetOwnerContentTypen(ownerContentTypen *string) {
	o.OwnerContentTypen = ownerContentTypen
}

// WithOwnerObjectID adds the ownerObjectID to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithOwnerObjectID(ownerObjectID *string) *ExtrasExportTemplatesListParams {
	o.SetOwnerObjectID(ownerObjectID)
	return o
}

// SetOwnerObjectID adds the ownerObjectId to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetOwnerObjectID(ownerObjectID *string) {
	o.OwnerObjectID = ownerObjectID
}

// WithOwnerObjectIDIc adds the ownerObjectIDIc to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithOwnerObjectIDIc(ownerObjectIDIc *string) *ExtrasExportTemplatesListParams {
	o.SetOwnerObjectIDIc(ownerObjectIDIc)
	return o
}

// SetOwnerObjectIDIc adds the ownerObjectIdIc to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetOwnerObjectIDIc(ownerObjectIDIc *string) {
	o.OwnerObjectIDIc = ownerObjectIDIc
}

// WithOwnerObjectIDIe adds the ownerObjectIDIe to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithOwnerObjectIDIe(ownerObjectIDIe *string) *ExtrasExportTemplatesListParams {
	o.SetOwnerObjectIDIe(ownerObjectIDIe)
	return o
}

// SetOwnerObjectIDIe adds the ownerObjectIdIe to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetOwnerObjectIDIe(ownerObjectIDIe *string) {
	o.OwnerObjectIDIe = ownerObjectIDIe
}

// WithOwnerObjectIDIew adds the ownerObjectIDIew to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithOwnerObjectIDIew(ownerObjectIDIew *string) *ExtrasExportTemplatesListParams {
	o.SetOwnerObjectIDIew(ownerObjectIDIew)
	return o
}

// SetOwnerObjectIDIew adds the ownerObjectIdIew to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetOwnerObjectIDIew(ownerObjectIDIew *string) {
	o.OwnerObjectIDIew = ownerObjectIDIew
}

// WithOwnerObjectIDIsw adds the ownerObjectIDIsw to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithOwnerObjectIDIsw(ownerObjectIDIsw *string) *ExtrasExportTemplatesListParams {
	o.SetOwnerObjectIDIsw(ownerObjectIDIsw)
	return o
}

// SetOwnerObjectIDIsw adds the ownerObjectIdIsw to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetOwnerObjectIDIsw(ownerObjectIDIsw *string) {
	o.OwnerObjectIDIsw = ownerObjectIDIsw
}

// WithOwnerObjectIDn adds the ownerObjectIDn to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithOwnerObjectIDn(ownerObjectIDn *string) *ExtrasExportTemplatesListParams {
	o.SetOwnerObjectIDn(ownerObjectIDn)
	return o
}

// SetOwnerObjectIDn adds the ownerObjectIdN to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetOwnerObjectIDn(ownerObjectIDn *string) {
	o.OwnerObjectIDn = ownerObjectIDn
}

// WithOwnerObjectIDNic adds the ownerObjectIDNic to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithOwnerObjectIDNic(ownerObjectIDNic *string) *ExtrasExportTemplatesListParams {
	o.SetOwnerObjectIDNic(ownerObjectIDNic)
	return o
}

// SetOwnerObjectIDNic adds the ownerObjectIdNic to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetOwnerObjectIDNic(ownerObjectIDNic *string) {
	o.OwnerObjectIDNic = ownerObjectIDNic
}

// WithOwnerObjectIDNie adds the ownerObjectIDNie to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithOwnerObjectIDNie(ownerObjectIDNie *string) *ExtrasExportTemplatesListParams {
	o.SetOwnerObjectIDNie(ownerObjectIDNie)
	return o
}

// SetOwnerObjectIDNie adds the ownerObjectIdNie to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetOwnerObjectIDNie(ownerObjectIDNie *string) {
	o.OwnerObjectIDNie = ownerObjectIDNie
}

// WithOwnerObjectIDNiew adds the ownerObjectIDNiew to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithOwnerObjectIDNiew(ownerObjectIDNiew *string) *ExtrasExportTemplatesListParams {
	o.SetOwnerObjectIDNiew(ownerObjectIDNiew)
	return o
}

// SetOwnerObjectIDNiew adds the ownerObjectIdNiew to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetOwnerObjectIDNiew(ownerObjectIDNiew *string) {
	o.OwnerObjectIDNiew = ownerObjectIDNiew
}

// WithOwnerObjectIDNisw adds the ownerObjectIDNisw to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithOwnerObjectIDNisw(ownerObjectIDNisw *string) *ExtrasExportTemplatesListParams {
	o.SetOwnerObjectIDNisw(ownerObjectIDNisw)
	return o
}

// SetOwnerObjectIDNisw adds the ownerObjectIdNisw to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetOwnerObjectIDNisw(ownerObjectIDNisw *string) {
	o.OwnerObjectIDNisw = ownerObjectIDNisw
}

// WithQ adds the q to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithQ(q *string) *ExtrasExportTemplatesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetQ(q *string) {
	o.Q = q
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasExportTemplatesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.OwnerContentType != nil {

		// query param owner_content_type
		var qrOwnerContentType string

		if o.OwnerContentType != nil {
			qrOwnerContentType = *o.OwnerContentType
		}
		qOwnerContentType := qrOwnerContentType
		if qOwnerContentType != "" {

			if err := r.SetQueryParam("owner_content_type", qOwnerContentType); err != nil {
				return err
			}
		}
	}

	if o.OwnerContentTypen != nil {

		// query param owner_content_type__n
		var qrOwnerContentTypen string

		if o.OwnerContentTypen != nil {
			qrOwnerContentTypen = *o.OwnerContentTypen
		}
		qOwnerContentTypen := qrOwnerContentTypen
		if qOwnerContentTypen != "" {

			if err := r.SetQueryParam("owner_content_type__n", qOwnerContentTypen); err != nil {
				return err
			}
		}
	}

	if o.OwnerObjectID != nil {

		// query param owner_object_id
		var qrOwnerObjectID string

		if o.OwnerObjectID != nil {
			qrOwnerObjectID = *o.OwnerObjectID
		}
		qOwnerObjectID := qrOwnerObjectID
		if qOwnerObjectID != "" {

			if err := r.SetQueryParam("owner_object_id", qOwnerObjectID); err != nil {
				return err
			}
		}
	}

	if o.OwnerObjectIDIc != nil {

		// query param owner_object_id__ic
		var qrOwnerObjectIDIc string

		if o.OwnerObjectIDIc != nil {
			qrOwnerObjectIDIc = *o.OwnerObjectIDIc
		}
		qOwnerObjectIDIc := qrOwnerObjectIDIc
		if qOwnerObjectIDIc != "" {

			if err := r.SetQueryParam("owner_object_id__ic", qOwnerObjectIDIc); err != nil {
				return err
			}
		}
	}

	if o.OwnerObjectIDIe != nil {

		// query param owner_object_id__ie
		var qrOwnerObjectIDIe string

		if o.OwnerObjectIDIe != nil {
			qrOwnerObjectIDIe = *o.OwnerObjectIDIe
		}
		qOwnerObjectIDIe := qrOwnerObjectIDIe
		if qOwnerObjectIDIe != "" {

			if err := r.SetQueryParam("owner_object_id__ie", qOwnerObjectIDIe); err != nil {
				return err
			}
		}
	}

	if o.OwnerObjectIDIew != nil {

		// query param owner_object_id__iew
		var qrOwnerObjectIDIew string

		if o.OwnerObjectIDIew != nil {
			qrOwnerObjectIDIew = *o.OwnerObjectIDIew
		}
		qOwnerObjectIDIew := qrOwnerObjectIDIew
		if qOwnerObjectIDIew != "" {

			if err := r.SetQueryParam("owner_object_id__iew", qOwnerObjectIDIew); err != nil {
				return err
			}
		}
	}

	if o.OwnerObjectIDIsw != nil {

		// query param owner_object_id__isw
		var qrOwnerObjectIDIsw string

		if o.OwnerObjectIDIsw != nil {
			qrOwnerObjectIDIsw = *o.OwnerObjectIDIsw
		}
		qOwnerObjectIDIsw := qrOwnerObjectIDIsw
		if qOwnerObjectIDIsw != "" {

			if err := r.SetQueryParam("owner_object_id__isw", qOwnerObjectIDIsw); err != nil {
				return err
			}
		}
	}

	if o.OwnerObjectIDn != nil {

		// query param owner_object_id__n
		var qrOwnerObjectIDn string

		if o.OwnerObjectIDn != nil {
			qrOwnerObjectIDn = *o.OwnerObjectIDn
		}
		qOwnerObjectIDn := qrOwnerObjectIDn
		if qOwnerObjectIDn != "" {

			if err := r.SetQueryParam("owner_object_id__n", qOwnerObjectIDn); err != nil {
				return err
			}
		}
	}

	if o.OwnerObjectIDNic != nil {

		// query param owner_object_id__nic
		var qrOwnerObjectIDNic string

		if o.OwnerObjectIDNic != nil {
			qrOwnerObjectIDNic = *o.OwnerObjectIDNic
		}
		qOwnerObjectIDNic := qrOwnerObjectIDNic
		if qOwnerObjectIDNic != "" {

			if err := r.SetQueryParam("owner_object_id__nic", qOwnerObjectIDNic); err != nil {
				return err
			}
		}
	}

	if o.OwnerObjectIDNie != nil {

		// query param owner_object_id__nie
		var qrOwnerObjectIDNie string

		if o.OwnerObjectIDNie != nil {
			qrOwnerObjectIDNie = *o.OwnerObjectIDNie
		}
		qOwnerObjectIDNie := qrOwnerObjectIDNie
		if qOwnerObjectIDNie != "" {

			if err := r.SetQueryParam("owner_object_id__nie", qOwnerObjectIDNie); err != nil {
				return err
			}
		}
	}

	if o.OwnerObjectIDNiew != nil {

		// query param owner_object_id__niew
		var qrOwnerObjectIDNiew string

		if o.OwnerObjectIDNiew != nil {
			qrOwnerObjectIDNiew = *o.OwnerObjectIDNiew
		}
		qOwnerObjectIDNiew := qrOwnerObjectIDNiew
		if qOwnerObjectIDNiew != "" {

			if err := r.SetQueryParam("owner_object_id__niew", qOwnerObjectIDNiew); err != nil {
				return err
			}
		}
	}

	if o.OwnerObjectIDNisw != nil {

		// query param owner_object_id__nisw
		var qrOwnerObjectIDNisw string

		if o.OwnerObjectIDNisw != nil {
			qrOwnerObjectIDNisw = *o.OwnerObjectIDNisw
		}
		qOwnerObjectIDNisw := qrOwnerObjectIDNisw
		if qOwnerObjectIDNisw != "" {

			if err := r.SetQueryParam("owner_object_id__nisw", qOwnerObjectIDNisw); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

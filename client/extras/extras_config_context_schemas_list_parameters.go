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

// NewExtrasConfigContextSchemasListParams creates a new ExtrasConfigContextSchemasListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasConfigContextSchemasListParams() *ExtrasConfigContextSchemasListParams {
	return &ExtrasConfigContextSchemasListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasConfigContextSchemasListParamsWithTimeout creates a new ExtrasConfigContextSchemasListParams object
// with the ability to set a timeout on a request.
func NewExtrasConfigContextSchemasListParamsWithTimeout(timeout time.Duration) *ExtrasConfigContextSchemasListParams {
	return &ExtrasConfigContextSchemasListParams{
		timeout: timeout,
	}
}

// NewExtrasConfigContextSchemasListParamsWithContext creates a new ExtrasConfigContextSchemasListParams object
// with the ability to set a context for a request.
func NewExtrasConfigContextSchemasListParamsWithContext(ctx context.Context) *ExtrasConfigContextSchemasListParams {
	return &ExtrasConfigContextSchemasListParams{
		Context: ctx,
	}
}

// NewExtrasConfigContextSchemasListParamsWithHTTPClient creates a new ExtrasConfigContextSchemasListParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasConfigContextSchemasListParamsWithHTTPClient(client *http.Client) *ExtrasConfigContextSchemasListParams {
	return &ExtrasConfigContextSchemasListParams{
		HTTPClient: client,
	}
}

/* ExtrasConfigContextSchemasListParams contains all the parameters to send to the API endpoint
   for the extras config context schemas list operation.

   Typically these are written to a http.Request.
*/
type ExtrasConfigContextSchemasListParams struct {

	// Description.
	Description *string

	// DescriptionIc.
	DescriptionIc *string

	// DescriptionIe.
	DescriptionIe *string

	// DescriptionIew.
	DescriptionIew *string

	// DescriptionIsw.
	DescriptionIsw *string

	// Descriptionn.
	Descriptionn *string

	// DescriptionNic.
	DescriptionNic *string

	// DescriptionNie.
	DescriptionNie *string

	// DescriptionNiew.
	DescriptionNiew *string

	// DescriptionNisw.
	DescriptionNisw *string

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

	// Q.
	Q *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras config context schemas list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasConfigContextSchemasListParams) WithDefaults() *ExtrasConfigContextSchemasListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras config context schemas list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasConfigContextSchemasListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) WithTimeout(timeout time.Duration) *ExtrasConfigContextSchemasListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) WithContext(ctx context.Context) *ExtrasConfigContextSchemasListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) WithHTTPClient(client *http.Client) *ExtrasConfigContextSchemasListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDescription adds the description to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) WithDescription(description *string) *ExtrasConfigContextSchemasListParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) SetDescription(description *string) {
	o.Description = description
}

// WithDescriptionIc adds the descriptionIc to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) WithDescriptionIc(descriptionIc *string) *ExtrasConfigContextSchemasListParams {
	o.SetDescriptionIc(descriptionIc)
	return o
}

// SetDescriptionIc adds the descriptionIc to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) SetDescriptionIc(descriptionIc *string) {
	o.DescriptionIc = descriptionIc
}

// WithDescriptionIe adds the descriptionIe to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) WithDescriptionIe(descriptionIe *string) *ExtrasConfigContextSchemasListParams {
	o.SetDescriptionIe(descriptionIe)
	return o
}

// SetDescriptionIe adds the descriptionIe to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) SetDescriptionIe(descriptionIe *string) {
	o.DescriptionIe = descriptionIe
}

// WithDescriptionIew adds the descriptionIew to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) WithDescriptionIew(descriptionIew *string) *ExtrasConfigContextSchemasListParams {
	o.SetDescriptionIew(descriptionIew)
	return o
}

// SetDescriptionIew adds the descriptionIew to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) SetDescriptionIew(descriptionIew *string) {
	o.DescriptionIew = descriptionIew
}

// WithDescriptionIsw adds the descriptionIsw to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) WithDescriptionIsw(descriptionIsw *string) *ExtrasConfigContextSchemasListParams {
	o.SetDescriptionIsw(descriptionIsw)
	return o
}

// SetDescriptionIsw adds the descriptionIsw to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) SetDescriptionIsw(descriptionIsw *string) {
	o.DescriptionIsw = descriptionIsw
}

// WithDescriptionn adds the descriptionn to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) WithDescriptionn(descriptionn *string) *ExtrasConfigContextSchemasListParams {
	o.SetDescriptionn(descriptionn)
	return o
}

// SetDescriptionn adds the descriptionN to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) SetDescriptionn(descriptionn *string) {
	o.Descriptionn = descriptionn
}

// WithDescriptionNic adds the descriptionNic to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) WithDescriptionNic(descriptionNic *string) *ExtrasConfigContextSchemasListParams {
	o.SetDescriptionNic(descriptionNic)
	return o
}

// SetDescriptionNic adds the descriptionNic to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) SetDescriptionNic(descriptionNic *string) {
	o.DescriptionNic = descriptionNic
}

// WithDescriptionNie adds the descriptionNie to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) WithDescriptionNie(descriptionNie *string) *ExtrasConfigContextSchemasListParams {
	o.SetDescriptionNie(descriptionNie)
	return o
}

// SetDescriptionNie adds the descriptionNie to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) SetDescriptionNie(descriptionNie *string) {
	o.DescriptionNie = descriptionNie
}

// WithDescriptionNiew adds the descriptionNiew to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) WithDescriptionNiew(descriptionNiew *string) *ExtrasConfigContextSchemasListParams {
	o.SetDescriptionNiew(descriptionNiew)
	return o
}

// SetDescriptionNiew adds the descriptionNiew to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) SetDescriptionNiew(descriptionNiew *string) {
	o.DescriptionNiew = descriptionNiew
}

// WithDescriptionNisw adds the descriptionNisw to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) WithDescriptionNisw(descriptionNisw *string) *ExtrasConfigContextSchemasListParams {
	o.SetDescriptionNisw(descriptionNisw)
	return o
}

// SetDescriptionNisw adds the descriptionNisw to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) SetDescriptionNisw(descriptionNisw *string) {
	o.DescriptionNisw = descriptionNisw
}

// WithID adds the id to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) WithID(id *string) *ExtrasConfigContextSchemasListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) WithIDIc(iDIc *string) *ExtrasConfigContextSchemasListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) WithIDIe(iDIe *string) *ExtrasConfigContextSchemasListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) WithIDIew(iDIew *string) *ExtrasConfigContextSchemasListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) WithIDIsw(iDIsw *string) *ExtrasConfigContextSchemasListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) WithIDn(iDn *string) *ExtrasConfigContextSchemasListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) WithIDNic(iDNic *string) *ExtrasConfigContextSchemasListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) WithIDNie(iDNie *string) *ExtrasConfigContextSchemasListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) WithIDNiew(iDNiew *string) *ExtrasConfigContextSchemasListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) WithIDNisw(iDNisw *string) *ExtrasConfigContextSchemasListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithLimit adds the limit to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) WithLimit(limit *int64) *ExtrasConfigContextSchemasListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) WithName(name *string) *ExtrasConfigContextSchemasListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) WithNameIc(nameIc *string) *ExtrasConfigContextSchemasListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) WithNameIe(nameIe *string) *ExtrasConfigContextSchemasListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) WithNameIew(nameIew *string) *ExtrasConfigContextSchemasListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) WithNameIsw(nameIsw *string) *ExtrasConfigContextSchemasListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) WithNamen(namen *string) *ExtrasConfigContextSchemasListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) WithNameNic(nameNic *string) *ExtrasConfigContextSchemasListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) WithNameNie(nameNie *string) *ExtrasConfigContextSchemasListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) WithNameNiew(nameNiew *string) *ExtrasConfigContextSchemasListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) WithNameNisw(nameNisw *string) *ExtrasConfigContextSchemasListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) WithOffset(offset *int64) *ExtrasConfigContextSchemasListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOwnerContentType adds the ownerContentType to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) WithOwnerContentType(ownerContentType *string) *ExtrasConfigContextSchemasListParams {
	o.SetOwnerContentType(ownerContentType)
	return o
}

// SetOwnerContentType adds the ownerContentType to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) SetOwnerContentType(ownerContentType *string) {
	o.OwnerContentType = ownerContentType
}

// WithOwnerContentTypen adds the ownerContentTypen to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) WithOwnerContentTypen(ownerContentTypen *string) *ExtrasConfigContextSchemasListParams {
	o.SetOwnerContentTypen(ownerContentTypen)
	return o
}

// SetOwnerContentTypen adds the ownerContentTypeN to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) SetOwnerContentTypen(ownerContentTypen *string) {
	o.OwnerContentTypen = ownerContentTypen
}

// WithQ adds the q to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) WithQ(q *string) *ExtrasConfigContextSchemasListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the extras config context schemas list params
func (o *ExtrasConfigContextSchemasListParams) SetQ(q *string) {
	o.Q = q
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasConfigContextSchemasListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Description != nil {

		// query param description
		var qrDescription string

		if o.Description != nil {
			qrDescription = *o.Description
		}
		qDescription := qrDescription
		if qDescription != "" {

			if err := r.SetQueryParam("description", qDescription); err != nil {
				return err
			}
		}
	}

	if o.DescriptionIc != nil {

		// query param description__ic
		var qrDescriptionIc string

		if o.DescriptionIc != nil {
			qrDescriptionIc = *o.DescriptionIc
		}
		qDescriptionIc := qrDescriptionIc
		if qDescriptionIc != "" {

			if err := r.SetQueryParam("description__ic", qDescriptionIc); err != nil {
				return err
			}
		}
	}

	if o.DescriptionIe != nil {

		// query param description__ie
		var qrDescriptionIe string

		if o.DescriptionIe != nil {
			qrDescriptionIe = *o.DescriptionIe
		}
		qDescriptionIe := qrDescriptionIe
		if qDescriptionIe != "" {

			if err := r.SetQueryParam("description__ie", qDescriptionIe); err != nil {
				return err
			}
		}
	}

	if o.DescriptionIew != nil {

		// query param description__iew
		var qrDescriptionIew string

		if o.DescriptionIew != nil {
			qrDescriptionIew = *o.DescriptionIew
		}
		qDescriptionIew := qrDescriptionIew
		if qDescriptionIew != "" {

			if err := r.SetQueryParam("description__iew", qDescriptionIew); err != nil {
				return err
			}
		}
	}

	if o.DescriptionIsw != nil {

		// query param description__isw
		var qrDescriptionIsw string

		if o.DescriptionIsw != nil {
			qrDescriptionIsw = *o.DescriptionIsw
		}
		qDescriptionIsw := qrDescriptionIsw
		if qDescriptionIsw != "" {

			if err := r.SetQueryParam("description__isw", qDescriptionIsw); err != nil {
				return err
			}
		}
	}

	if o.Descriptionn != nil {

		// query param description__n
		var qrDescriptionn string

		if o.Descriptionn != nil {
			qrDescriptionn = *o.Descriptionn
		}
		qDescriptionn := qrDescriptionn
		if qDescriptionn != "" {

			if err := r.SetQueryParam("description__n", qDescriptionn); err != nil {
				return err
			}
		}
	}

	if o.DescriptionNic != nil {

		// query param description__nic
		var qrDescriptionNic string

		if o.DescriptionNic != nil {
			qrDescriptionNic = *o.DescriptionNic
		}
		qDescriptionNic := qrDescriptionNic
		if qDescriptionNic != "" {

			if err := r.SetQueryParam("description__nic", qDescriptionNic); err != nil {
				return err
			}
		}
	}

	if o.DescriptionNie != nil {

		// query param description__nie
		var qrDescriptionNie string

		if o.DescriptionNie != nil {
			qrDescriptionNie = *o.DescriptionNie
		}
		qDescriptionNie := qrDescriptionNie
		if qDescriptionNie != "" {

			if err := r.SetQueryParam("description__nie", qDescriptionNie); err != nil {
				return err
			}
		}
	}

	if o.DescriptionNiew != nil {

		// query param description__niew
		var qrDescriptionNiew string

		if o.DescriptionNiew != nil {
			qrDescriptionNiew = *o.DescriptionNiew
		}
		qDescriptionNiew := qrDescriptionNiew
		if qDescriptionNiew != "" {

			if err := r.SetQueryParam("description__niew", qDescriptionNiew); err != nil {
				return err
			}
		}
	}

	if o.DescriptionNisw != nil {

		// query param description__nisw
		var qrDescriptionNisw string

		if o.DescriptionNisw != nil {
			qrDescriptionNisw = *o.DescriptionNisw
		}
		qDescriptionNisw := qrDescriptionNisw
		if qDescriptionNisw != "" {

			if err := r.SetQueryParam("description__nisw", qDescriptionNisw); err != nil {
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

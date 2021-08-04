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

// NewExtrasRelationshipsListParams creates a new ExtrasRelationshipsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasRelationshipsListParams() *ExtrasRelationshipsListParams {
	return &ExtrasRelationshipsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasRelationshipsListParamsWithTimeout creates a new ExtrasRelationshipsListParams object
// with the ability to set a timeout on a request.
func NewExtrasRelationshipsListParamsWithTimeout(timeout time.Duration) *ExtrasRelationshipsListParams {
	return &ExtrasRelationshipsListParams{
		timeout: timeout,
	}
}

// NewExtrasRelationshipsListParamsWithContext creates a new ExtrasRelationshipsListParams object
// with the ability to set a context for a request.
func NewExtrasRelationshipsListParamsWithContext(ctx context.Context) *ExtrasRelationshipsListParams {
	return &ExtrasRelationshipsListParams{
		Context: ctx,
	}
}

// NewExtrasRelationshipsListParamsWithHTTPClient creates a new ExtrasRelationshipsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasRelationshipsListParamsWithHTTPClient(client *http.Client) *ExtrasRelationshipsListParams {
	return &ExtrasRelationshipsListParams{
		HTTPClient: client,
	}
}

/* ExtrasRelationshipsListParams contains all the parameters to send to the API endpoint
   for the extras relationships list operation.

   Typically these are written to a http.Request.
*/
type ExtrasRelationshipsListParams struct {

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

	// SourceType.
	SourceType *string

	// SourceTypen.
	SourceTypen *string

	// Type.
	Type *string

	// Typen.
	Typen *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras relationships list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasRelationshipsListParams) WithDefaults() *ExtrasRelationshipsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras relationships list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasRelationshipsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras relationships list params
func (o *ExtrasRelationshipsListParams) WithTimeout(timeout time.Duration) *ExtrasRelationshipsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras relationships list params
func (o *ExtrasRelationshipsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras relationships list params
func (o *ExtrasRelationshipsListParams) WithContext(ctx context.Context) *ExtrasRelationshipsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras relationships list params
func (o *ExtrasRelationshipsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras relationships list params
func (o *ExtrasRelationshipsListParams) WithHTTPClient(client *http.Client) *ExtrasRelationshipsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras relationships list params
func (o *ExtrasRelationshipsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDestinationType adds the destinationType to the extras relationships list params
func (o *ExtrasRelationshipsListParams) WithDestinationType(destinationType *string) *ExtrasRelationshipsListParams {
	o.SetDestinationType(destinationType)
	return o
}

// SetDestinationType adds the destinationType to the extras relationships list params
func (o *ExtrasRelationshipsListParams) SetDestinationType(destinationType *string) {
	o.DestinationType = destinationType
}

// WithDestinationTypen adds the destinationTypen to the extras relationships list params
func (o *ExtrasRelationshipsListParams) WithDestinationTypen(destinationTypen *string) *ExtrasRelationshipsListParams {
	o.SetDestinationTypen(destinationTypen)
	return o
}

// SetDestinationTypen adds the destinationTypeN to the extras relationships list params
func (o *ExtrasRelationshipsListParams) SetDestinationTypen(destinationTypen *string) {
	o.DestinationTypen = destinationTypen
}

// WithID adds the id to the extras relationships list params
func (o *ExtrasRelationshipsListParams) WithID(id *string) *ExtrasRelationshipsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras relationships list params
func (o *ExtrasRelationshipsListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the extras relationships list params
func (o *ExtrasRelationshipsListParams) WithIDIc(iDIc *string) *ExtrasRelationshipsListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the extras relationships list params
func (o *ExtrasRelationshipsListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the extras relationships list params
func (o *ExtrasRelationshipsListParams) WithIDIe(iDIe *string) *ExtrasRelationshipsListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the extras relationships list params
func (o *ExtrasRelationshipsListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the extras relationships list params
func (o *ExtrasRelationshipsListParams) WithIDIew(iDIew *string) *ExtrasRelationshipsListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the extras relationships list params
func (o *ExtrasRelationshipsListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the extras relationships list params
func (o *ExtrasRelationshipsListParams) WithIDIsw(iDIsw *string) *ExtrasRelationshipsListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the extras relationships list params
func (o *ExtrasRelationshipsListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the extras relationships list params
func (o *ExtrasRelationshipsListParams) WithIDn(iDn *string) *ExtrasRelationshipsListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the extras relationships list params
func (o *ExtrasRelationshipsListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the extras relationships list params
func (o *ExtrasRelationshipsListParams) WithIDNic(iDNic *string) *ExtrasRelationshipsListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the extras relationships list params
func (o *ExtrasRelationshipsListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the extras relationships list params
func (o *ExtrasRelationshipsListParams) WithIDNie(iDNie *string) *ExtrasRelationshipsListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the extras relationships list params
func (o *ExtrasRelationshipsListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the extras relationships list params
func (o *ExtrasRelationshipsListParams) WithIDNiew(iDNiew *string) *ExtrasRelationshipsListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the extras relationships list params
func (o *ExtrasRelationshipsListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the extras relationships list params
func (o *ExtrasRelationshipsListParams) WithIDNisw(iDNisw *string) *ExtrasRelationshipsListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the extras relationships list params
func (o *ExtrasRelationshipsListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithLimit adds the limit to the extras relationships list params
func (o *ExtrasRelationshipsListParams) WithLimit(limit *int64) *ExtrasRelationshipsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the extras relationships list params
func (o *ExtrasRelationshipsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the extras relationships list params
func (o *ExtrasRelationshipsListParams) WithName(name *string) *ExtrasRelationshipsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the extras relationships list params
func (o *ExtrasRelationshipsListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the extras relationships list params
func (o *ExtrasRelationshipsListParams) WithNameIc(nameIc *string) *ExtrasRelationshipsListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the extras relationships list params
func (o *ExtrasRelationshipsListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the extras relationships list params
func (o *ExtrasRelationshipsListParams) WithNameIe(nameIe *string) *ExtrasRelationshipsListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the extras relationships list params
func (o *ExtrasRelationshipsListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the extras relationships list params
func (o *ExtrasRelationshipsListParams) WithNameIew(nameIew *string) *ExtrasRelationshipsListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the extras relationships list params
func (o *ExtrasRelationshipsListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the extras relationships list params
func (o *ExtrasRelationshipsListParams) WithNameIsw(nameIsw *string) *ExtrasRelationshipsListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the extras relationships list params
func (o *ExtrasRelationshipsListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the extras relationships list params
func (o *ExtrasRelationshipsListParams) WithNamen(namen *string) *ExtrasRelationshipsListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the extras relationships list params
func (o *ExtrasRelationshipsListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the extras relationships list params
func (o *ExtrasRelationshipsListParams) WithNameNic(nameNic *string) *ExtrasRelationshipsListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the extras relationships list params
func (o *ExtrasRelationshipsListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the extras relationships list params
func (o *ExtrasRelationshipsListParams) WithNameNie(nameNie *string) *ExtrasRelationshipsListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the extras relationships list params
func (o *ExtrasRelationshipsListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the extras relationships list params
func (o *ExtrasRelationshipsListParams) WithNameNiew(nameNiew *string) *ExtrasRelationshipsListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the extras relationships list params
func (o *ExtrasRelationshipsListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the extras relationships list params
func (o *ExtrasRelationshipsListParams) WithNameNisw(nameNisw *string) *ExtrasRelationshipsListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the extras relationships list params
func (o *ExtrasRelationshipsListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the extras relationships list params
func (o *ExtrasRelationshipsListParams) WithOffset(offset *int64) *ExtrasRelationshipsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the extras relationships list params
func (o *ExtrasRelationshipsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithSourceType adds the sourceType to the extras relationships list params
func (o *ExtrasRelationshipsListParams) WithSourceType(sourceType *string) *ExtrasRelationshipsListParams {
	o.SetSourceType(sourceType)
	return o
}

// SetSourceType adds the sourceType to the extras relationships list params
func (o *ExtrasRelationshipsListParams) SetSourceType(sourceType *string) {
	o.SourceType = sourceType
}

// WithSourceTypen adds the sourceTypen to the extras relationships list params
func (o *ExtrasRelationshipsListParams) WithSourceTypen(sourceTypen *string) *ExtrasRelationshipsListParams {
	o.SetSourceTypen(sourceTypen)
	return o
}

// SetSourceTypen adds the sourceTypeN to the extras relationships list params
func (o *ExtrasRelationshipsListParams) SetSourceTypen(sourceTypen *string) {
	o.SourceTypen = sourceTypen
}

// WithType adds the typeVar to the extras relationships list params
func (o *ExtrasRelationshipsListParams) WithType(typeVar *string) *ExtrasRelationshipsListParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the extras relationships list params
func (o *ExtrasRelationshipsListParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithTypen adds the typen to the extras relationships list params
func (o *ExtrasRelationshipsListParams) WithTypen(typen *string) *ExtrasRelationshipsListParams {
	o.SetTypen(typen)
	return o
}

// SetTypen adds the typeN to the extras relationships list params
func (o *ExtrasRelationshipsListParams) SetTypen(typen *string) {
	o.Typen = typen
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasRelationshipsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.Type != nil {

		// query param type
		var qrType string

		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}
	}

	if o.Typen != nil {

		// query param type__n
		var qrTypen string

		if o.Typen != nil {
			qrTypen = *o.Typen
		}
		qTypen := qrTypen
		if qTypen != "" {

			if err := r.SetQueryParam("type__n", qTypen); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

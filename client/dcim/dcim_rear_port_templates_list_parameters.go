package dcim

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

// NewDcimRearPortTemplatesListParams creates a new DcimRearPortTemplatesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRearPortTemplatesListParams() *DcimRearPortTemplatesListParams {
	return &DcimRearPortTemplatesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRearPortTemplatesListParamsWithTimeout creates a new DcimRearPortTemplatesListParams object
// with the ability to set a timeout on a request.
func NewDcimRearPortTemplatesListParamsWithTimeout(timeout time.Duration) *DcimRearPortTemplatesListParams {
	return &DcimRearPortTemplatesListParams{
		timeout: timeout,
	}
}

// NewDcimRearPortTemplatesListParamsWithContext creates a new DcimRearPortTemplatesListParams object
// with the ability to set a context for a request.
func NewDcimRearPortTemplatesListParamsWithContext(ctx context.Context) *DcimRearPortTemplatesListParams {
	return &DcimRearPortTemplatesListParams{
		Context: ctx,
	}
}

// NewDcimRearPortTemplatesListParamsWithHTTPClient creates a new DcimRearPortTemplatesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRearPortTemplatesListParamsWithHTTPClient(client *http.Client) *DcimRearPortTemplatesListParams {
	return &DcimRearPortTemplatesListParams{
		HTTPClient: client,
	}
}

/* DcimRearPortTemplatesListParams contains all the parameters to send to the API endpoint
   for the dcim rear port templates list operation.

   Typically these are written to a http.Request.
*/
type DcimRearPortTemplatesListParams struct {

	// DevicetypeID.
	DevicetypeID *string

	// DevicetypeIDn.
	DevicetypeIDn *string

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

	// Positions.
	Positions *string

	// PositionsGt.
	PositionsGt *string

	// PositionsGte.
	PositionsGte *string

	// PositionsLt.
	PositionsLt *string

	// PositionsLte.
	PositionsLte *string

	// Positionsn.
	Positionsn *string

	// Q.
	Q *string

	// Type.
	Type *string

	// Typen.
	Typen *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim rear port templates list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRearPortTemplatesListParams) WithDefaults() *DcimRearPortTemplatesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim rear port templates list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRearPortTemplatesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithTimeout(timeout time.Duration) *DcimRearPortTemplatesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithContext(ctx context.Context) *DcimRearPortTemplatesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithHTTPClient(client *http.Client) *DcimRearPortTemplatesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDevicetypeID adds the devicetypeID to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithDevicetypeID(devicetypeID *string) *DcimRearPortTemplatesListParams {
	o.SetDevicetypeID(devicetypeID)
	return o
}

// SetDevicetypeID adds the devicetypeId to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetDevicetypeID(devicetypeID *string) {
	o.DevicetypeID = devicetypeID
}

// WithDevicetypeIDn adds the devicetypeIDn to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithDevicetypeIDn(devicetypeIDn *string) *DcimRearPortTemplatesListParams {
	o.SetDevicetypeIDn(devicetypeIDn)
	return o
}

// SetDevicetypeIDn adds the devicetypeIdN to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetDevicetypeIDn(devicetypeIDn *string) {
	o.DevicetypeIDn = devicetypeIDn
}

// WithID adds the id to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithID(id *string) *DcimRearPortTemplatesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithIDIc(iDIc *string) *DcimRearPortTemplatesListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithIDIe(iDIe *string) *DcimRearPortTemplatesListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithIDIew(iDIew *string) *DcimRearPortTemplatesListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithIDIsw(iDIsw *string) *DcimRearPortTemplatesListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithIDn(iDn *string) *DcimRearPortTemplatesListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithIDNic(iDNic *string) *DcimRearPortTemplatesListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithIDNie(iDNie *string) *DcimRearPortTemplatesListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithIDNiew(iDNiew *string) *DcimRearPortTemplatesListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithIDNisw(iDNisw *string) *DcimRearPortTemplatesListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithLimit adds the limit to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithLimit(limit *int64) *DcimRearPortTemplatesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithName(name *string) *DcimRearPortTemplatesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithNameIc(nameIc *string) *DcimRearPortTemplatesListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithNameIe(nameIe *string) *DcimRearPortTemplatesListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithNameIew(nameIew *string) *DcimRearPortTemplatesListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithNameIsw(nameIsw *string) *DcimRearPortTemplatesListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithNamen(namen *string) *DcimRearPortTemplatesListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithNameNic(nameNic *string) *DcimRearPortTemplatesListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithNameNie(nameNie *string) *DcimRearPortTemplatesListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithNameNiew(nameNiew *string) *DcimRearPortTemplatesListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithNameNisw(nameNisw *string) *DcimRearPortTemplatesListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithOffset(offset *int64) *DcimRearPortTemplatesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPositions adds the positions to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithPositions(positions *string) *DcimRearPortTemplatesListParams {
	o.SetPositions(positions)
	return o
}

// SetPositions adds the positions to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetPositions(positions *string) {
	o.Positions = positions
}

// WithPositionsGt adds the positionsGt to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithPositionsGt(positionsGt *string) *DcimRearPortTemplatesListParams {
	o.SetPositionsGt(positionsGt)
	return o
}

// SetPositionsGt adds the positionsGt to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetPositionsGt(positionsGt *string) {
	o.PositionsGt = positionsGt
}

// WithPositionsGte adds the positionsGte to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithPositionsGte(positionsGte *string) *DcimRearPortTemplatesListParams {
	o.SetPositionsGte(positionsGte)
	return o
}

// SetPositionsGte adds the positionsGte to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetPositionsGte(positionsGte *string) {
	o.PositionsGte = positionsGte
}

// WithPositionsLt adds the positionsLt to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithPositionsLt(positionsLt *string) *DcimRearPortTemplatesListParams {
	o.SetPositionsLt(positionsLt)
	return o
}

// SetPositionsLt adds the positionsLt to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetPositionsLt(positionsLt *string) {
	o.PositionsLt = positionsLt
}

// WithPositionsLte adds the positionsLte to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithPositionsLte(positionsLte *string) *DcimRearPortTemplatesListParams {
	o.SetPositionsLte(positionsLte)
	return o
}

// SetPositionsLte adds the positionsLte to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetPositionsLte(positionsLte *string) {
	o.PositionsLte = positionsLte
}

// WithPositionsn adds the positionsn to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithPositionsn(positionsn *string) *DcimRearPortTemplatesListParams {
	o.SetPositionsn(positionsn)
	return o
}

// SetPositionsn adds the positionsN to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetPositionsn(positionsn *string) {
	o.Positionsn = positionsn
}

// WithQ adds the q to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithQ(q *string) *DcimRearPortTemplatesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetQ(q *string) {
	o.Q = q
}

// WithType adds the typeVar to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithType(typeVar *string) *DcimRearPortTemplatesListParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithTypen adds the typen to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithTypen(typen *string) *DcimRearPortTemplatesListParams {
	o.SetTypen(typen)
	return o
}

// SetTypen adds the typeN to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetTypen(typen *string) {
	o.Typen = typen
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRearPortTemplatesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DevicetypeID != nil {

		// query param devicetype_id
		var qrDevicetypeID string

		if o.DevicetypeID != nil {
			qrDevicetypeID = *o.DevicetypeID
		}
		qDevicetypeID := qrDevicetypeID
		if qDevicetypeID != "" {

			if err := r.SetQueryParam("devicetype_id", qDevicetypeID); err != nil {
				return err
			}
		}
	}

	if o.DevicetypeIDn != nil {

		// query param devicetype_id__n
		var qrDevicetypeIDn string

		if o.DevicetypeIDn != nil {
			qrDevicetypeIDn = *o.DevicetypeIDn
		}
		qDevicetypeIDn := qrDevicetypeIDn
		if qDevicetypeIDn != "" {

			if err := r.SetQueryParam("devicetype_id__n", qDevicetypeIDn); err != nil {
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

	if o.Positions != nil {

		// query param positions
		var qrPositions string

		if o.Positions != nil {
			qrPositions = *o.Positions
		}
		qPositions := qrPositions
		if qPositions != "" {

			if err := r.SetQueryParam("positions", qPositions); err != nil {
				return err
			}
		}
	}

	if o.PositionsGt != nil {

		// query param positions__gt
		var qrPositionsGt string

		if o.PositionsGt != nil {
			qrPositionsGt = *o.PositionsGt
		}
		qPositionsGt := qrPositionsGt
		if qPositionsGt != "" {

			if err := r.SetQueryParam("positions__gt", qPositionsGt); err != nil {
				return err
			}
		}
	}

	if o.PositionsGte != nil {

		// query param positions__gte
		var qrPositionsGte string

		if o.PositionsGte != nil {
			qrPositionsGte = *o.PositionsGte
		}
		qPositionsGte := qrPositionsGte
		if qPositionsGte != "" {

			if err := r.SetQueryParam("positions__gte", qPositionsGte); err != nil {
				return err
			}
		}
	}

	if o.PositionsLt != nil {

		// query param positions__lt
		var qrPositionsLt string

		if o.PositionsLt != nil {
			qrPositionsLt = *o.PositionsLt
		}
		qPositionsLt := qrPositionsLt
		if qPositionsLt != "" {

			if err := r.SetQueryParam("positions__lt", qPositionsLt); err != nil {
				return err
			}
		}
	}

	if o.PositionsLte != nil {

		// query param positions__lte
		var qrPositionsLte string

		if o.PositionsLte != nil {
			qrPositionsLte = *o.PositionsLte
		}
		qPositionsLte := qrPositionsLte
		if qPositionsLte != "" {

			if err := r.SetQueryParam("positions__lte", qPositionsLte); err != nil {
				return err
			}
		}
	}

	if o.Positionsn != nil {

		// query param positions__n
		var qrPositionsn string

		if o.Positionsn != nil {
			qrPositionsn = *o.Positionsn
		}
		qPositionsn := qrPositionsn
		if qPositionsn != "" {

			if err := r.SetQueryParam("positions__n", qPositionsn); err != nil {
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

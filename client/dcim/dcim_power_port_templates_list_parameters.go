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

// NewDcimPowerPortTemplatesListParams creates a new DcimPowerPortTemplatesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPowerPortTemplatesListParams() *DcimPowerPortTemplatesListParams {
	return &DcimPowerPortTemplatesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerPortTemplatesListParamsWithTimeout creates a new DcimPowerPortTemplatesListParams object
// with the ability to set a timeout on a request.
func NewDcimPowerPortTemplatesListParamsWithTimeout(timeout time.Duration) *DcimPowerPortTemplatesListParams {
	return &DcimPowerPortTemplatesListParams{
		timeout: timeout,
	}
}

// NewDcimPowerPortTemplatesListParamsWithContext creates a new DcimPowerPortTemplatesListParams object
// with the ability to set a context for a request.
func NewDcimPowerPortTemplatesListParamsWithContext(ctx context.Context) *DcimPowerPortTemplatesListParams {
	return &DcimPowerPortTemplatesListParams{
		Context: ctx,
	}
}

// NewDcimPowerPortTemplatesListParamsWithHTTPClient creates a new DcimPowerPortTemplatesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPowerPortTemplatesListParamsWithHTTPClient(client *http.Client) *DcimPowerPortTemplatesListParams {
	return &DcimPowerPortTemplatesListParams{
		HTTPClient: client,
	}
}

/* DcimPowerPortTemplatesListParams contains all the parameters to send to the API endpoint
   for the dcim power port templates list operation.

   Typically these are written to a http.Request.
*/
type DcimPowerPortTemplatesListParams struct {

	// AllocatedDraw.
	AllocatedDraw *string

	// AllocatedDrawGt.
	AllocatedDrawGt *string

	// AllocatedDrawGte.
	AllocatedDrawGte *string

	// AllocatedDrawLt.
	AllocatedDrawLt *string

	// AllocatedDrawLte.
	AllocatedDrawLte *string

	// AllocatedDrawn.
	AllocatedDrawn *string

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

	// MaximumDraw.
	MaximumDraw *string

	// MaximumDrawGt.
	MaximumDrawGt *string

	// MaximumDrawGte.
	MaximumDrawGte *string

	// MaximumDrawLt.
	MaximumDrawLt *string

	// MaximumDrawLte.
	MaximumDrawLte *string

	// MaximumDrawn.
	MaximumDrawn *string

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

// WithDefaults hydrates default values in the dcim power port templates list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerPortTemplatesListParams) WithDefaults() *DcimPowerPortTemplatesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim power port templates list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerPortTemplatesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithTimeout(timeout time.Duration) *DcimPowerPortTemplatesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithContext(ctx context.Context) *DcimPowerPortTemplatesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithHTTPClient(client *http.Client) *DcimPowerPortTemplatesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllocatedDraw adds the allocatedDraw to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithAllocatedDraw(allocatedDraw *string) *DcimPowerPortTemplatesListParams {
	o.SetAllocatedDraw(allocatedDraw)
	return o
}

// SetAllocatedDraw adds the allocatedDraw to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetAllocatedDraw(allocatedDraw *string) {
	o.AllocatedDraw = allocatedDraw
}

// WithAllocatedDrawGt adds the allocatedDrawGt to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithAllocatedDrawGt(allocatedDrawGt *string) *DcimPowerPortTemplatesListParams {
	o.SetAllocatedDrawGt(allocatedDrawGt)
	return o
}

// SetAllocatedDrawGt adds the allocatedDrawGt to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetAllocatedDrawGt(allocatedDrawGt *string) {
	o.AllocatedDrawGt = allocatedDrawGt
}

// WithAllocatedDrawGte adds the allocatedDrawGte to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithAllocatedDrawGte(allocatedDrawGte *string) *DcimPowerPortTemplatesListParams {
	o.SetAllocatedDrawGte(allocatedDrawGte)
	return o
}

// SetAllocatedDrawGte adds the allocatedDrawGte to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetAllocatedDrawGte(allocatedDrawGte *string) {
	o.AllocatedDrawGte = allocatedDrawGte
}

// WithAllocatedDrawLt adds the allocatedDrawLt to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithAllocatedDrawLt(allocatedDrawLt *string) *DcimPowerPortTemplatesListParams {
	o.SetAllocatedDrawLt(allocatedDrawLt)
	return o
}

// SetAllocatedDrawLt adds the allocatedDrawLt to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetAllocatedDrawLt(allocatedDrawLt *string) {
	o.AllocatedDrawLt = allocatedDrawLt
}

// WithAllocatedDrawLte adds the allocatedDrawLte to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithAllocatedDrawLte(allocatedDrawLte *string) *DcimPowerPortTemplatesListParams {
	o.SetAllocatedDrawLte(allocatedDrawLte)
	return o
}

// SetAllocatedDrawLte adds the allocatedDrawLte to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetAllocatedDrawLte(allocatedDrawLte *string) {
	o.AllocatedDrawLte = allocatedDrawLte
}

// WithAllocatedDrawn adds the allocatedDrawn to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithAllocatedDrawn(allocatedDrawn *string) *DcimPowerPortTemplatesListParams {
	o.SetAllocatedDrawn(allocatedDrawn)
	return o
}

// SetAllocatedDrawn adds the allocatedDrawN to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetAllocatedDrawn(allocatedDrawn *string) {
	o.AllocatedDrawn = allocatedDrawn
}

// WithDevicetypeID adds the devicetypeID to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithDevicetypeID(devicetypeID *string) *DcimPowerPortTemplatesListParams {
	o.SetDevicetypeID(devicetypeID)
	return o
}

// SetDevicetypeID adds the devicetypeId to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetDevicetypeID(devicetypeID *string) {
	o.DevicetypeID = devicetypeID
}

// WithDevicetypeIDn adds the devicetypeIDn to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithDevicetypeIDn(devicetypeIDn *string) *DcimPowerPortTemplatesListParams {
	o.SetDevicetypeIDn(devicetypeIDn)
	return o
}

// SetDevicetypeIDn adds the devicetypeIdN to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetDevicetypeIDn(devicetypeIDn *string) {
	o.DevicetypeIDn = devicetypeIDn
}

// WithID adds the id to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithID(id *string) *DcimPowerPortTemplatesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithIDIc(iDIc *string) *DcimPowerPortTemplatesListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithIDIe(iDIe *string) *DcimPowerPortTemplatesListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithIDIew(iDIew *string) *DcimPowerPortTemplatesListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithIDIsw(iDIsw *string) *DcimPowerPortTemplatesListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithIDn(iDn *string) *DcimPowerPortTemplatesListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithIDNic(iDNic *string) *DcimPowerPortTemplatesListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithIDNie(iDNie *string) *DcimPowerPortTemplatesListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithIDNiew(iDNiew *string) *DcimPowerPortTemplatesListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithIDNisw(iDNisw *string) *DcimPowerPortTemplatesListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithLimit adds the limit to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithLimit(limit *int64) *DcimPowerPortTemplatesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithMaximumDraw adds the maximumDraw to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithMaximumDraw(maximumDraw *string) *DcimPowerPortTemplatesListParams {
	o.SetMaximumDraw(maximumDraw)
	return o
}

// SetMaximumDraw adds the maximumDraw to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetMaximumDraw(maximumDraw *string) {
	o.MaximumDraw = maximumDraw
}

// WithMaximumDrawGt adds the maximumDrawGt to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithMaximumDrawGt(maximumDrawGt *string) *DcimPowerPortTemplatesListParams {
	o.SetMaximumDrawGt(maximumDrawGt)
	return o
}

// SetMaximumDrawGt adds the maximumDrawGt to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetMaximumDrawGt(maximumDrawGt *string) {
	o.MaximumDrawGt = maximumDrawGt
}

// WithMaximumDrawGte adds the maximumDrawGte to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithMaximumDrawGte(maximumDrawGte *string) *DcimPowerPortTemplatesListParams {
	o.SetMaximumDrawGte(maximumDrawGte)
	return o
}

// SetMaximumDrawGte adds the maximumDrawGte to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetMaximumDrawGte(maximumDrawGte *string) {
	o.MaximumDrawGte = maximumDrawGte
}

// WithMaximumDrawLt adds the maximumDrawLt to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithMaximumDrawLt(maximumDrawLt *string) *DcimPowerPortTemplatesListParams {
	o.SetMaximumDrawLt(maximumDrawLt)
	return o
}

// SetMaximumDrawLt adds the maximumDrawLt to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetMaximumDrawLt(maximumDrawLt *string) {
	o.MaximumDrawLt = maximumDrawLt
}

// WithMaximumDrawLte adds the maximumDrawLte to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithMaximumDrawLte(maximumDrawLte *string) *DcimPowerPortTemplatesListParams {
	o.SetMaximumDrawLte(maximumDrawLte)
	return o
}

// SetMaximumDrawLte adds the maximumDrawLte to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetMaximumDrawLte(maximumDrawLte *string) {
	o.MaximumDrawLte = maximumDrawLte
}

// WithMaximumDrawn adds the maximumDrawn to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithMaximumDrawn(maximumDrawn *string) *DcimPowerPortTemplatesListParams {
	o.SetMaximumDrawn(maximumDrawn)
	return o
}

// SetMaximumDrawn adds the maximumDrawN to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetMaximumDrawn(maximumDrawn *string) {
	o.MaximumDrawn = maximumDrawn
}

// WithName adds the name to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithName(name *string) *DcimPowerPortTemplatesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithNameIc(nameIc *string) *DcimPowerPortTemplatesListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithNameIe(nameIe *string) *DcimPowerPortTemplatesListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithNameIew(nameIew *string) *DcimPowerPortTemplatesListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithNameIsw(nameIsw *string) *DcimPowerPortTemplatesListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithNamen(namen *string) *DcimPowerPortTemplatesListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithNameNic(nameNic *string) *DcimPowerPortTemplatesListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithNameNie(nameNie *string) *DcimPowerPortTemplatesListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithNameNiew(nameNiew *string) *DcimPowerPortTemplatesListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithNameNisw(nameNisw *string) *DcimPowerPortTemplatesListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithOffset(offset *int64) *DcimPowerPortTemplatesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithQ(q *string) *DcimPowerPortTemplatesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetQ(q *string) {
	o.Q = q
}

// WithType adds the typeVar to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithType(typeVar *string) *DcimPowerPortTemplatesListParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithTypen adds the typen to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithTypen(typen *string) *DcimPowerPortTemplatesListParams {
	o.SetTypen(typen)
	return o
}

// SetTypen adds the typeN to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetTypen(typen *string) {
	o.Typen = typen
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerPortTemplatesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AllocatedDraw != nil {

		// query param allocated_draw
		var qrAllocatedDraw string

		if o.AllocatedDraw != nil {
			qrAllocatedDraw = *o.AllocatedDraw
		}
		qAllocatedDraw := qrAllocatedDraw
		if qAllocatedDraw != "" {

			if err := r.SetQueryParam("allocated_draw", qAllocatedDraw); err != nil {
				return err
			}
		}
	}

	if o.AllocatedDrawGt != nil {

		// query param allocated_draw__gt
		var qrAllocatedDrawGt string

		if o.AllocatedDrawGt != nil {
			qrAllocatedDrawGt = *o.AllocatedDrawGt
		}
		qAllocatedDrawGt := qrAllocatedDrawGt
		if qAllocatedDrawGt != "" {

			if err := r.SetQueryParam("allocated_draw__gt", qAllocatedDrawGt); err != nil {
				return err
			}
		}
	}

	if o.AllocatedDrawGte != nil {

		// query param allocated_draw__gte
		var qrAllocatedDrawGte string

		if o.AllocatedDrawGte != nil {
			qrAllocatedDrawGte = *o.AllocatedDrawGte
		}
		qAllocatedDrawGte := qrAllocatedDrawGte
		if qAllocatedDrawGte != "" {

			if err := r.SetQueryParam("allocated_draw__gte", qAllocatedDrawGte); err != nil {
				return err
			}
		}
	}

	if o.AllocatedDrawLt != nil {

		// query param allocated_draw__lt
		var qrAllocatedDrawLt string

		if o.AllocatedDrawLt != nil {
			qrAllocatedDrawLt = *o.AllocatedDrawLt
		}
		qAllocatedDrawLt := qrAllocatedDrawLt
		if qAllocatedDrawLt != "" {

			if err := r.SetQueryParam("allocated_draw__lt", qAllocatedDrawLt); err != nil {
				return err
			}
		}
	}

	if o.AllocatedDrawLte != nil {

		// query param allocated_draw__lte
		var qrAllocatedDrawLte string

		if o.AllocatedDrawLte != nil {
			qrAllocatedDrawLte = *o.AllocatedDrawLte
		}
		qAllocatedDrawLte := qrAllocatedDrawLte
		if qAllocatedDrawLte != "" {

			if err := r.SetQueryParam("allocated_draw__lte", qAllocatedDrawLte); err != nil {
				return err
			}
		}
	}

	if o.AllocatedDrawn != nil {

		// query param allocated_draw__n
		var qrAllocatedDrawn string

		if o.AllocatedDrawn != nil {
			qrAllocatedDrawn = *o.AllocatedDrawn
		}
		qAllocatedDrawn := qrAllocatedDrawn
		if qAllocatedDrawn != "" {

			if err := r.SetQueryParam("allocated_draw__n", qAllocatedDrawn); err != nil {
				return err
			}
		}
	}

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

	if o.MaximumDraw != nil {

		// query param maximum_draw
		var qrMaximumDraw string

		if o.MaximumDraw != nil {
			qrMaximumDraw = *o.MaximumDraw
		}
		qMaximumDraw := qrMaximumDraw
		if qMaximumDraw != "" {

			if err := r.SetQueryParam("maximum_draw", qMaximumDraw); err != nil {
				return err
			}
		}
	}

	if o.MaximumDrawGt != nil {

		// query param maximum_draw__gt
		var qrMaximumDrawGt string

		if o.MaximumDrawGt != nil {
			qrMaximumDrawGt = *o.MaximumDrawGt
		}
		qMaximumDrawGt := qrMaximumDrawGt
		if qMaximumDrawGt != "" {

			if err := r.SetQueryParam("maximum_draw__gt", qMaximumDrawGt); err != nil {
				return err
			}
		}
	}

	if o.MaximumDrawGte != nil {

		// query param maximum_draw__gte
		var qrMaximumDrawGte string

		if o.MaximumDrawGte != nil {
			qrMaximumDrawGte = *o.MaximumDrawGte
		}
		qMaximumDrawGte := qrMaximumDrawGte
		if qMaximumDrawGte != "" {

			if err := r.SetQueryParam("maximum_draw__gte", qMaximumDrawGte); err != nil {
				return err
			}
		}
	}

	if o.MaximumDrawLt != nil {

		// query param maximum_draw__lt
		var qrMaximumDrawLt string

		if o.MaximumDrawLt != nil {
			qrMaximumDrawLt = *o.MaximumDrawLt
		}
		qMaximumDrawLt := qrMaximumDrawLt
		if qMaximumDrawLt != "" {

			if err := r.SetQueryParam("maximum_draw__lt", qMaximumDrawLt); err != nil {
				return err
			}
		}
	}

	if o.MaximumDrawLte != nil {

		// query param maximum_draw__lte
		var qrMaximumDrawLte string

		if o.MaximumDrawLte != nil {
			qrMaximumDrawLte = *o.MaximumDrawLte
		}
		qMaximumDrawLte := qrMaximumDrawLte
		if qMaximumDrawLte != "" {

			if err := r.SetQueryParam("maximum_draw__lte", qMaximumDrawLte); err != nil {
				return err
			}
		}
	}

	if o.MaximumDrawn != nil {

		// query param maximum_draw__n
		var qrMaximumDrawn string

		if o.MaximumDrawn != nil {
			qrMaximumDrawn = *o.MaximumDrawn
		}
		qMaximumDrawn := qrMaximumDrawn
		if qMaximumDrawn != "" {

			if err := r.SetQueryParam("maximum_draw__n", qMaximumDrawn); err != nil {
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

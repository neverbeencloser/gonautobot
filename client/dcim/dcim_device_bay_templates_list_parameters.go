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

// NewDcimDeviceBayTemplatesListParams creates a new DcimDeviceBayTemplatesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimDeviceBayTemplatesListParams() *DcimDeviceBayTemplatesListParams {
	return &DcimDeviceBayTemplatesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDeviceBayTemplatesListParamsWithTimeout creates a new DcimDeviceBayTemplatesListParams object
// with the ability to set a timeout on a request.
func NewDcimDeviceBayTemplatesListParamsWithTimeout(timeout time.Duration) *DcimDeviceBayTemplatesListParams {
	return &DcimDeviceBayTemplatesListParams{
		timeout: timeout,
	}
}

// NewDcimDeviceBayTemplatesListParamsWithContext creates a new DcimDeviceBayTemplatesListParams object
// with the ability to set a context for a request.
func NewDcimDeviceBayTemplatesListParamsWithContext(ctx context.Context) *DcimDeviceBayTemplatesListParams {
	return &DcimDeviceBayTemplatesListParams{
		Context: ctx,
	}
}

// NewDcimDeviceBayTemplatesListParamsWithHTTPClient creates a new DcimDeviceBayTemplatesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimDeviceBayTemplatesListParamsWithHTTPClient(client *http.Client) *DcimDeviceBayTemplatesListParams {
	return &DcimDeviceBayTemplatesListParams{
		HTTPClient: client,
	}
}

/* DcimDeviceBayTemplatesListParams contains all the parameters to send to the API endpoint
   for the dcim device bay templates list operation.

   Typically these are written to a http.Request.
*/
type DcimDeviceBayTemplatesListParams struct {

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

	// Q.
	Q *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim device bay templates list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDeviceBayTemplatesListParams) WithDefaults() *DcimDeviceBayTemplatesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim device bay templates list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDeviceBayTemplatesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithTimeout(timeout time.Duration) *DcimDeviceBayTemplatesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithContext(ctx context.Context) *DcimDeviceBayTemplatesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithHTTPClient(client *http.Client) *DcimDeviceBayTemplatesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDevicetypeID adds the devicetypeID to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithDevicetypeID(devicetypeID *string) *DcimDeviceBayTemplatesListParams {
	o.SetDevicetypeID(devicetypeID)
	return o
}

// SetDevicetypeID adds the devicetypeId to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetDevicetypeID(devicetypeID *string) {
	o.DevicetypeID = devicetypeID
}

// WithDevicetypeIDn adds the devicetypeIDn to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithDevicetypeIDn(devicetypeIDn *string) *DcimDeviceBayTemplatesListParams {
	o.SetDevicetypeIDn(devicetypeIDn)
	return o
}

// SetDevicetypeIDn adds the devicetypeIdN to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetDevicetypeIDn(devicetypeIDn *string) {
	o.DevicetypeIDn = devicetypeIDn
}

// WithID adds the id to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithID(id *string) *DcimDeviceBayTemplatesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithIDIc(iDIc *string) *DcimDeviceBayTemplatesListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithIDIe(iDIe *string) *DcimDeviceBayTemplatesListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithIDIew(iDIew *string) *DcimDeviceBayTemplatesListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithIDIsw(iDIsw *string) *DcimDeviceBayTemplatesListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithIDn(iDn *string) *DcimDeviceBayTemplatesListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithIDNic(iDNic *string) *DcimDeviceBayTemplatesListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithIDNie(iDNie *string) *DcimDeviceBayTemplatesListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithIDNiew(iDNiew *string) *DcimDeviceBayTemplatesListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithIDNisw(iDNisw *string) *DcimDeviceBayTemplatesListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithLimit adds the limit to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithLimit(limit *int64) *DcimDeviceBayTemplatesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithName(name *string) *DcimDeviceBayTemplatesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithNameIc(nameIc *string) *DcimDeviceBayTemplatesListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithNameIe(nameIe *string) *DcimDeviceBayTemplatesListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithNameIew(nameIew *string) *DcimDeviceBayTemplatesListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithNameIsw(nameIsw *string) *DcimDeviceBayTemplatesListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithNamen(namen *string) *DcimDeviceBayTemplatesListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithNameNic(nameNic *string) *DcimDeviceBayTemplatesListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithNameNie(nameNie *string) *DcimDeviceBayTemplatesListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithNameNiew(nameNiew *string) *DcimDeviceBayTemplatesListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithNameNisw(nameNisw *string) *DcimDeviceBayTemplatesListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithOffset(offset *int64) *DcimDeviceBayTemplatesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) WithQ(q *string) *DcimDeviceBayTemplatesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim device bay templates list params
func (o *DcimDeviceBayTemplatesListParams) SetQ(q *string) {
	o.Q = q
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDeviceBayTemplatesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

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

// NewExtrasCustomFieldsListParams creates a new ExtrasCustomFieldsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasCustomFieldsListParams() *ExtrasCustomFieldsListParams {
	return &ExtrasCustomFieldsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasCustomFieldsListParamsWithTimeout creates a new ExtrasCustomFieldsListParams object
// with the ability to set a timeout on a request.
func NewExtrasCustomFieldsListParamsWithTimeout(timeout time.Duration) *ExtrasCustomFieldsListParams {
	return &ExtrasCustomFieldsListParams{
		timeout: timeout,
	}
}

// NewExtrasCustomFieldsListParamsWithContext creates a new ExtrasCustomFieldsListParams object
// with the ability to set a context for a request.
func NewExtrasCustomFieldsListParamsWithContext(ctx context.Context) *ExtrasCustomFieldsListParams {
	return &ExtrasCustomFieldsListParams{
		Context: ctx,
	}
}

// NewExtrasCustomFieldsListParamsWithHTTPClient creates a new ExtrasCustomFieldsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasCustomFieldsListParamsWithHTTPClient(client *http.Client) *ExtrasCustomFieldsListParams {
	return &ExtrasCustomFieldsListParams{
		HTTPClient: client,
	}
}

/* ExtrasCustomFieldsListParams contains all the parameters to send to the API endpoint
   for the extras custom fields list operation.

   Typically these are written to a http.Request.
*/
type ExtrasCustomFieldsListParams struct {

	// ContentTypes.
	ContentTypes *string

	// ContentTypesn.
	ContentTypesn *string

	// FilterLogic.
	FilterLogic *string

	// FilterLogicn.
	FilterLogicn *string

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

	// Required.
	Required *string

	// Weight.
	Weight *string

	// WeightGt.
	WeightGt *string

	// WeightGte.
	WeightGte *string

	// WeightLt.
	WeightLt *string

	// WeightLte.
	WeightLte *string

	// Weightn.
	Weightn *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras custom fields list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasCustomFieldsListParams) WithDefaults() *ExtrasCustomFieldsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras custom fields list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasCustomFieldsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithTimeout(timeout time.Duration) *ExtrasCustomFieldsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithContext(ctx context.Context) *ExtrasCustomFieldsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithHTTPClient(client *http.Client) *ExtrasCustomFieldsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentTypes adds the contentTypes to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithContentTypes(contentTypes *string) *ExtrasCustomFieldsListParams {
	o.SetContentTypes(contentTypes)
	return o
}

// SetContentTypes adds the contentTypes to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetContentTypes(contentTypes *string) {
	o.ContentTypes = contentTypes
}

// WithContentTypesn adds the contentTypesn to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithContentTypesn(contentTypesn *string) *ExtrasCustomFieldsListParams {
	o.SetContentTypesn(contentTypesn)
	return o
}

// SetContentTypesn adds the contentTypesN to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetContentTypesn(contentTypesn *string) {
	o.ContentTypesn = contentTypesn
}

// WithFilterLogic adds the filterLogic to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithFilterLogic(filterLogic *string) *ExtrasCustomFieldsListParams {
	o.SetFilterLogic(filterLogic)
	return o
}

// SetFilterLogic adds the filterLogic to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetFilterLogic(filterLogic *string) {
	o.FilterLogic = filterLogic
}

// WithFilterLogicn adds the filterLogicn to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithFilterLogicn(filterLogicn *string) *ExtrasCustomFieldsListParams {
	o.SetFilterLogicn(filterLogicn)
	return o
}

// SetFilterLogicn adds the filterLogicN to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetFilterLogicn(filterLogicn *string) {
	o.FilterLogicn = filterLogicn
}

// WithID adds the id to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithID(id *string) *ExtrasCustomFieldsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithIDIc(iDIc *string) *ExtrasCustomFieldsListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithIDIe(iDIe *string) *ExtrasCustomFieldsListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithIDIew(iDIew *string) *ExtrasCustomFieldsListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithIDIsw(iDIsw *string) *ExtrasCustomFieldsListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithIDn(iDn *string) *ExtrasCustomFieldsListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithIDNic(iDNic *string) *ExtrasCustomFieldsListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithIDNie(iDNie *string) *ExtrasCustomFieldsListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithIDNiew(iDNiew *string) *ExtrasCustomFieldsListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithIDNisw(iDNisw *string) *ExtrasCustomFieldsListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithLimit adds the limit to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithLimit(limit *int64) *ExtrasCustomFieldsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithName(name *string) *ExtrasCustomFieldsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithNameIc(nameIc *string) *ExtrasCustomFieldsListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithNameIe(nameIe *string) *ExtrasCustomFieldsListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithNameIew(nameIew *string) *ExtrasCustomFieldsListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithNameIsw(nameIsw *string) *ExtrasCustomFieldsListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithNamen(namen *string) *ExtrasCustomFieldsListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithNameNic(nameNic *string) *ExtrasCustomFieldsListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithNameNie(nameNie *string) *ExtrasCustomFieldsListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithNameNiew(nameNiew *string) *ExtrasCustomFieldsListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithNameNisw(nameNisw *string) *ExtrasCustomFieldsListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithOffset(offset *int64) *ExtrasCustomFieldsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithQ(q *string) *ExtrasCustomFieldsListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetQ(q *string) {
	o.Q = q
}

// WithRequired adds the required to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithRequired(required *string) *ExtrasCustomFieldsListParams {
	o.SetRequired(required)
	return o
}

// SetRequired adds the required to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetRequired(required *string) {
	o.Required = required
}

// WithWeight adds the weight to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithWeight(weight *string) *ExtrasCustomFieldsListParams {
	o.SetWeight(weight)
	return o
}

// SetWeight adds the weight to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetWeight(weight *string) {
	o.Weight = weight
}

// WithWeightGt adds the weightGt to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithWeightGt(weightGt *string) *ExtrasCustomFieldsListParams {
	o.SetWeightGt(weightGt)
	return o
}

// SetWeightGt adds the weightGt to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetWeightGt(weightGt *string) {
	o.WeightGt = weightGt
}

// WithWeightGte adds the weightGte to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithWeightGte(weightGte *string) *ExtrasCustomFieldsListParams {
	o.SetWeightGte(weightGte)
	return o
}

// SetWeightGte adds the weightGte to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetWeightGte(weightGte *string) {
	o.WeightGte = weightGte
}

// WithWeightLt adds the weightLt to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithWeightLt(weightLt *string) *ExtrasCustomFieldsListParams {
	o.SetWeightLt(weightLt)
	return o
}

// SetWeightLt adds the weightLt to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetWeightLt(weightLt *string) {
	o.WeightLt = weightLt
}

// WithWeightLte adds the weightLte to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithWeightLte(weightLte *string) *ExtrasCustomFieldsListParams {
	o.SetWeightLte(weightLte)
	return o
}

// SetWeightLte adds the weightLte to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetWeightLte(weightLte *string) {
	o.WeightLte = weightLte
}

// WithWeightn adds the weightn to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithWeightn(weightn *string) *ExtrasCustomFieldsListParams {
	o.SetWeightn(weightn)
	return o
}

// SetWeightn adds the weightN to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetWeightn(weightn *string) {
	o.Weightn = weightn
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasCustomFieldsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ContentTypes != nil {

		// query param content_types
		var qrContentTypes string

		if o.ContentTypes != nil {
			qrContentTypes = *o.ContentTypes
		}
		qContentTypes := qrContentTypes
		if qContentTypes != "" {

			if err := r.SetQueryParam("content_types", qContentTypes); err != nil {
				return err
			}
		}
	}

	if o.ContentTypesn != nil {

		// query param content_types__n
		var qrContentTypesn string

		if o.ContentTypesn != nil {
			qrContentTypesn = *o.ContentTypesn
		}
		qContentTypesn := qrContentTypesn
		if qContentTypesn != "" {

			if err := r.SetQueryParam("content_types__n", qContentTypesn); err != nil {
				return err
			}
		}
	}

	if o.FilterLogic != nil {

		// query param filter_logic
		var qrFilterLogic string

		if o.FilterLogic != nil {
			qrFilterLogic = *o.FilterLogic
		}
		qFilterLogic := qrFilterLogic
		if qFilterLogic != "" {

			if err := r.SetQueryParam("filter_logic", qFilterLogic); err != nil {
				return err
			}
		}
	}

	if o.FilterLogicn != nil {

		// query param filter_logic__n
		var qrFilterLogicn string

		if o.FilterLogicn != nil {
			qrFilterLogicn = *o.FilterLogicn
		}
		qFilterLogicn := qrFilterLogicn
		if qFilterLogicn != "" {

			if err := r.SetQueryParam("filter_logic__n", qFilterLogicn); err != nil {
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

	if o.Required != nil {

		// query param required
		var qrRequired string

		if o.Required != nil {
			qrRequired = *o.Required
		}
		qRequired := qrRequired
		if qRequired != "" {

			if err := r.SetQueryParam("required", qRequired); err != nil {
				return err
			}
		}
	}

	if o.Weight != nil {

		// query param weight
		var qrWeight string

		if o.Weight != nil {
			qrWeight = *o.Weight
		}
		qWeight := qrWeight
		if qWeight != "" {

			if err := r.SetQueryParam("weight", qWeight); err != nil {
				return err
			}
		}
	}

	if o.WeightGt != nil {

		// query param weight__gt
		var qrWeightGt string

		if o.WeightGt != nil {
			qrWeightGt = *o.WeightGt
		}
		qWeightGt := qrWeightGt
		if qWeightGt != "" {

			if err := r.SetQueryParam("weight__gt", qWeightGt); err != nil {
				return err
			}
		}
	}

	if o.WeightGte != nil {

		// query param weight__gte
		var qrWeightGte string

		if o.WeightGte != nil {
			qrWeightGte = *o.WeightGte
		}
		qWeightGte := qrWeightGte
		if qWeightGte != "" {

			if err := r.SetQueryParam("weight__gte", qWeightGte); err != nil {
				return err
			}
		}
	}

	if o.WeightLt != nil {

		// query param weight__lt
		var qrWeightLt string

		if o.WeightLt != nil {
			qrWeightLt = *o.WeightLt
		}
		qWeightLt := qrWeightLt
		if qWeightLt != "" {

			if err := r.SetQueryParam("weight__lt", qWeightLt); err != nil {
				return err
			}
		}
	}

	if o.WeightLte != nil {

		// query param weight__lte
		var qrWeightLte string

		if o.WeightLte != nil {
			qrWeightLte = *o.WeightLte
		}
		qWeightLte := qrWeightLte
		if qWeightLte != "" {

			if err := r.SetQueryParam("weight__lte", qWeightLte); err != nil {
				return err
			}
		}
	}

	if o.Weightn != nil {

		// query param weight__n
		var qrWeightn string

		if o.Weightn != nil {
			qrWeightn = *o.Weightn
		}
		qWeightn := qrWeightn
		if qWeightn != "" {

			if err := r.SetQueryParam("weight__n", qWeightn); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

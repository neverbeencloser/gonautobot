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

// NewExtrasCustomFieldChoicesListParams creates a new ExtrasCustomFieldChoicesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasCustomFieldChoicesListParams() *ExtrasCustomFieldChoicesListParams {
	return &ExtrasCustomFieldChoicesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasCustomFieldChoicesListParamsWithTimeout creates a new ExtrasCustomFieldChoicesListParams object
// with the ability to set a timeout on a request.
func NewExtrasCustomFieldChoicesListParamsWithTimeout(timeout time.Duration) *ExtrasCustomFieldChoicesListParams {
	return &ExtrasCustomFieldChoicesListParams{
		timeout: timeout,
	}
}

// NewExtrasCustomFieldChoicesListParamsWithContext creates a new ExtrasCustomFieldChoicesListParams object
// with the ability to set a context for a request.
func NewExtrasCustomFieldChoicesListParamsWithContext(ctx context.Context) *ExtrasCustomFieldChoicesListParams {
	return &ExtrasCustomFieldChoicesListParams{
		Context: ctx,
	}
}

// NewExtrasCustomFieldChoicesListParamsWithHTTPClient creates a new ExtrasCustomFieldChoicesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasCustomFieldChoicesListParamsWithHTTPClient(client *http.Client) *ExtrasCustomFieldChoicesListParams {
	return &ExtrasCustomFieldChoicesListParams{
		HTTPClient: client,
	}
}

/* ExtrasCustomFieldChoicesListParams contains all the parameters to send to the API endpoint
   for the extras custom field choices list operation.

   Typically these are written to a http.Request.
*/
type ExtrasCustomFieldChoicesListParams struct {

	// Field.
	Field *string

	// Fieldn.
	Fieldn *string

	// FieldID.
	FieldID *string

	// FieldIDn.
	FieldIDn *string

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

	// Q.
	Q *string

	// Value.
	Value *string

	// ValueIc.
	ValueIc *string

	// ValueIe.
	ValueIe *string

	// ValueIew.
	ValueIew *string

	// ValueIsw.
	ValueIsw *string

	// Valuen.
	Valuen *string

	// ValueNic.
	ValueNic *string

	// ValueNie.
	ValueNie *string

	// ValueNiew.
	ValueNiew *string

	// ValueNisw.
	ValueNisw *string

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

// WithDefaults hydrates default values in the extras custom field choices list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasCustomFieldChoicesListParams) WithDefaults() *ExtrasCustomFieldChoicesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras custom field choices list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasCustomFieldChoicesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) WithTimeout(timeout time.Duration) *ExtrasCustomFieldChoicesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) WithContext(ctx context.Context) *ExtrasCustomFieldChoicesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) WithHTTPClient(client *http.Client) *ExtrasCustomFieldChoicesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithField adds the field to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) WithField(field *string) *ExtrasCustomFieldChoicesListParams {
	o.SetField(field)
	return o
}

// SetField adds the field to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) SetField(field *string) {
	o.Field = field
}

// WithFieldn adds the fieldn to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) WithFieldn(fieldn *string) *ExtrasCustomFieldChoicesListParams {
	o.SetFieldn(fieldn)
	return o
}

// SetFieldn adds the fieldN to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) SetFieldn(fieldn *string) {
	o.Fieldn = fieldn
}

// WithFieldID adds the fieldID to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) WithFieldID(fieldID *string) *ExtrasCustomFieldChoicesListParams {
	o.SetFieldID(fieldID)
	return o
}

// SetFieldID adds the fieldId to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) SetFieldID(fieldID *string) {
	o.FieldID = fieldID
}

// WithFieldIDn adds the fieldIDn to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) WithFieldIDn(fieldIDn *string) *ExtrasCustomFieldChoicesListParams {
	o.SetFieldIDn(fieldIDn)
	return o
}

// SetFieldIDn adds the fieldIdN to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) SetFieldIDn(fieldIDn *string) {
	o.FieldIDn = fieldIDn
}

// WithID adds the id to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) WithID(id *string) *ExtrasCustomFieldChoicesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) WithIDIc(iDIc *string) *ExtrasCustomFieldChoicesListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) WithIDIe(iDIe *string) *ExtrasCustomFieldChoicesListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) WithIDIew(iDIew *string) *ExtrasCustomFieldChoicesListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) WithIDIsw(iDIsw *string) *ExtrasCustomFieldChoicesListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) WithIDn(iDn *string) *ExtrasCustomFieldChoicesListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) WithIDNic(iDNic *string) *ExtrasCustomFieldChoicesListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) WithIDNie(iDNie *string) *ExtrasCustomFieldChoicesListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) WithIDNiew(iDNiew *string) *ExtrasCustomFieldChoicesListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) WithIDNisw(iDNisw *string) *ExtrasCustomFieldChoicesListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithLimit adds the limit to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) WithLimit(limit *int64) *ExtrasCustomFieldChoicesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) WithOffset(offset *int64) *ExtrasCustomFieldChoicesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) WithQ(q *string) *ExtrasCustomFieldChoicesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) SetQ(q *string) {
	o.Q = q
}

// WithValue adds the value to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) WithValue(value *string) *ExtrasCustomFieldChoicesListParams {
	o.SetValue(value)
	return o
}

// SetValue adds the value to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) SetValue(value *string) {
	o.Value = value
}

// WithValueIc adds the valueIc to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) WithValueIc(valueIc *string) *ExtrasCustomFieldChoicesListParams {
	o.SetValueIc(valueIc)
	return o
}

// SetValueIc adds the valueIc to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) SetValueIc(valueIc *string) {
	o.ValueIc = valueIc
}

// WithValueIe adds the valueIe to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) WithValueIe(valueIe *string) *ExtrasCustomFieldChoicesListParams {
	o.SetValueIe(valueIe)
	return o
}

// SetValueIe adds the valueIe to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) SetValueIe(valueIe *string) {
	o.ValueIe = valueIe
}

// WithValueIew adds the valueIew to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) WithValueIew(valueIew *string) *ExtrasCustomFieldChoicesListParams {
	o.SetValueIew(valueIew)
	return o
}

// SetValueIew adds the valueIew to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) SetValueIew(valueIew *string) {
	o.ValueIew = valueIew
}

// WithValueIsw adds the valueIsw to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) WithValueIsw(valueIsw *string) *ExtrasCustomFieldChoicesListParams {
	o.SetValueIsw(valueIsw)
	return o
}

// SetValueIsw adds the valueIsw to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) SetValueIsw(valueIsw *string) {
	o.ValueIsw = valueIsw
}

// WithValuen adds the valuen to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) WithValuen(valuen *string) *ExtrasCustomFieldChoicesListParams {
	o.SetValuen(valuen)
	return o
}

// SetValuen adds the valueN to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) SetValuen(valuen *string) {
	o.Valuen = valuen
}

// WithValueNic adds the valueNic to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) WithValueNic(valueNic *string) *ExtrasCustomFieldChoicesListParams {
	o.SetValueNic(valueNic)
	return o
}

// SetValueNic adds the valueNic to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) SetValueNic(valueNic *string) {
	o.ValueNic = valueNic
}

// WithValueNie adds the valueNie to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) WithValueNie(valueNie *string) *ExtrasCustomFieldChoicesListParams {
	o.SetValueNie(valueNie)
	return o
}

// SetValueNie adds the valueNie to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) SetValueNie(valueNie *string) {
	o.ValueNie = valueNie
}

// WithValueNiew adds the valueNiew to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) WithValueNiew(valueNiew *string) *ExtrasCustomFieldChoicesListParams {
	o.SetValueNiew(valueNiew)
	return o
}

// SetValueNiew adds the valueNiew to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) SetValueNiew(valueNiew *string) {
	o.ValueNiew = valueNiew
}

// WithValueNisw adds the valueNisw to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) WithValueNisw(valueNisw *string) *ExtrasCustomFieldChoicesListParams {
	o.SetValueNisw(valueNisw)
	return o
}

// SetValueNisw adds the valueNisw to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) SetValueNisw(valueNisw *string) {
	o.ValueNisw = valueNisw
}

// WithWeight adds the weight to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) WithWeight(weight *string) *ExtrasCustomFieldChoicesListParams {
	o.SetWeight(weight)
	return o
}

// SetWeight adds the weight to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) SetWeight(weight *string) {
	o.Weight = weight
}

// WithWeightGt adds the weightGt to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) WithWeightGt(weightGt *string) *ExtrasCustomFieldChoicesListParams {
	o.SetWeightGt(weightGt)
	return o
}

// SetWeightGt adds the weightGt to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) SetWeightGt(weightGt *string) {
	o.WeightGt = weightGt
}

// WithWeightGte adds the weightGte to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) WithWeightGte(weightGte *string) *ExtrasCustomFieldChoicesListParams {
	o.SetWeightGte(weightGte)
	return o
}

// SetWeightGte adds the weightGte to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) SetWeightGte(weightGte *string) {
	o.WeightGte = weightGte
}

// WithWeightLt adds the weightLt to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) WithWeightLt(weightLt *string) *ExtrasCustomFieldChoicesListParams {
	o.SetWeightLt(weightLt)
	return o
}

// SetWeightLt adds the weightLt to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) SetWeightLt(weightLt *string) {
	o.WeightLt = weightLt
}

// WithWeightLte adds the weightLte to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) WithWeightLte(weightLte *string) *ExtrasCustomFieldChoicesListParams {
	o.SetWeightLte(weightLte)
	return o
}

// SetWeightLte adds the weightLte to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) SetWeightLte(weightLte *string) {
	o.WeightLte = weightLte
}

// WithWeightn adds the weightn to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) WithWeightn(weightn *string) *ExtrasCustomFieldChoicesListParams {
	o.SetWeightn(weightn)
	return o
}

// SetWeightn adds the weightN to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) SetWeightn(weightn *string) {
	o.Weightn = weightn
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasCustomFieldChoicesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Field != nil {

		// query param field
		var qrField string

		if o.Field != nil {
			qrField = *o.Field
		}
		qField := qrField
		if qField != "" {

			if err := r.SetQueryParam("field", qField); err != nil {
				return err
			}
		}
	}

	if o.Fieldn != nil {

		// query param field__n
		var qrFieldn string

		if o.Fieldn != nil {
			qrFieldn = *o.Fieldn
		}
		qFieldn := qrFieldn
		if qFieldn != "" {

			if err := r.SetQueryParam("field__n", qFieldn); err != nil {
				return err
			}
		}
	}

	if o.FieldID != nil {

		// query param field_id
		var qrFieldID string

		if o.FieldID != nil {
			qrFieldID = *o.FieldID
		}
		qFieldID := qrFieldID
		if qFieldID != "" {

			if err := r.SetQueryParam("field_id", qFieldID); err != nil {
				return err
			}
		}
	}

	if o.FieldIDn != nil {

		// query param field_id__n
		var qrFieldIDn string

		if o.FieldIDn != nil {
			qrFieldIDn = *o.FieldIDn
		}
		qFieldIDn := qrFieldIDn
		if qFieldIDn != "" {

			if err := r.SetQueryParam("field_id__n", qFieldIDn); err != nil {
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

	if o.Value != nil {

		// query param value
		var qrValue string

		if o.Value != nil {
			qrValue = *o.Value
		}
		qValue := qrValue
		if qValue != "" {

			if err := r.SetQueryParam("value", qValue); err != nil {
				return err
			}
		}
	}

	if o.ValueIc != nil {

		// query param value__ic
		var qrValueIc string

		if o.ValueIc != nil {
			qrValueIc = *o.ValueIc
		}
		qValueIc := qrValueIc
		if qValueIc != "" {

			if err := r.SetQueryParam("value__ic", qValueIc); err != nil {
				return err
			}
		}
	}

	if o.ValueIe != nil {

		// query param value__ie
		var qrValueIe string

		if o.ValueIe != nil {
			qrValueIe = *o.ValueIe
		}
		qValueIe := qrValueIe
		if qValueIe != "" {

			if err := r.SetQueryParam("value__ie", qValueIe); err != nil {
				return err
			}
		}
	}

	if o.ValueIew != nil {

		// query param value__iew
		var qrValueIew string

		if o.ValueIew != nil {
			qrValueIew = *o.ValueIew
		}
		qValueIew := qrValueIew
		if qValueIew != "" {

			if err := r.SetQueryParam("value__iew", qValueIew); err != nil {
				return err
			}
		}
	}

	if o.ValueIsw != nil {

		// query param value__isw
		var qrValueIsw string

		if o.ValueIsw != nil {
			qrValueIsw = *o.ValueIsw
		}
		qValueIsw := qrValueIsw
		if qValueIsw != "" {

			if err := r.SetQueryParam("value__isw", qValueIsw); err != nil {
				return err
			}
		}
	}

	if o.Valuen != nil {

		// query param value__n
		var qrValuen string

		if o.Valuen != nil {
			qrValuen = *o.Valuen
		}
		qValuen := qrValuen
		if qValuen != "" {

			if err := r.SetQueryParam("value__n", qValuen); err != nil {
				return err
			}
		}
	}

	if o.ValueNic != nil {

		// query param value__nic
		var qrValueNic string

		if o.ValueNic != nil {
			qrValueNic = *o.ValueNic
		}
		qValueNic := qrValueNic
		if qValueNic != "" {

			if err := r.SetQueryParam("value__nic", qValueNic); err != nil {
				return err
			}
		}
	}

	if o.ValueNie != nil {

		// query param value__nie
		var qrValueNie string

		if o.ValueNie != nil {
			qrValueNie = *o.ValueNie
		}
		qValueNie := qrValueNie
		if qValueNie != "" {

			if err := r.SetQueryParam("value__nie", qValueNie); err != nil {
				return err
			}
		}
	}

	if o.ValueNiew != nil {

		// query param value__niew
		var qrValueNiew string

		if o.ValueNiew != nil {
			qrValueNiew = *o.ValueNiew
		}
		qValueNiew := qrValueNiew
		if qValueNiew != "" {

			if err := r.SetQueryParam("value__niew", qValueNiew); err != nil {
				return err
			}
		}
	}

	if o.ValueNisw != nil {

		// query param value__nisw
		var qrValueNisw string

		if o.ValueNisw != nil {
			qrValueNisw = *o.ValueNisw
		}
		qValueNisw := qrValueNisw
		if qValueNisw != "" {

			if err := r.SetQueryParam("value__nisw", qValueNisw); err != nil {
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

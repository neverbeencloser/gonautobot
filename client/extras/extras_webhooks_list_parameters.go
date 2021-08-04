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

// NewExtrasWebhooksListParams creates a new ExtrasWebhooksListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasWebhooksListParams() *ExtrasWebhooksListParams {
	return &ExtrasWebhooksListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasWebhooksListParamsWithTimeout creates a new ExtrasWebhooksListParams object
// with the ability to set a timeout on a request.
func NewExtrasWebhooksListParamsWithTimeout(timeout time.Duration) *ExtrasWebhooksListParams {
	return &ExtrasWebhooksListParams{
		timeout: timeout,
	}
}

// NewExtrasWebhooksListParamsWithContext creates a new ExtrasWebhooksListParams object
// with the ability to set a context for a request.
func NewExtrasWebhooksListParamsWithContext(ctx context.Context) *ExtrasWebhooksListParams {
	return &ExtrasWebhooksListParams{
		Context: ctx,
	}
}

// NewExtrasWebhooksListParamsWithHTTPClient creates a new ExtrasWebhooksListParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasWebhooksListParamsWithHTTPClient(client *http.Client) *ExtrasWebhooksListParams {
	return &ExtrasWebhooksListParams{
		HTTPClient: client,
	}
}

/* ExtrasWebhooksListParams contains all the parameters to send to the API endpoint
   for the extras webhooks list operation.

   Typically these are written to a http.Request.
*/
type ExtrasWebhooksListParams struct {

	// ContentTypes.
	ContentTypes *string

	// ContentTypesn.
	ContentTypesn *string

	// Enabled.
	Enabled *string

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

	// PayloadURL.
	PayloadURL *string

	// PayloadURLIc.
	PayloadURLIc *string

	// PayloadURLIe.
	PayloadURLIe *string

	// PayloadURLIew.
	PayloadURLIew *string

	// PayloadURLIsw.
	PayloadURLIsw *string

	// PayloadURLn.
	PayloadURLn *string

	// PayloadURLNic.
	PayloadURLNic *string

	// PayloadURLNie.
	PayloadURLNie *string

	// PayloadURLNiew.
	PayloadURLNiew *string

	// PayloadURLNisw.
	PayloadURLNisw *string

	// Q.
	Q *string

	// TypeCreate.
	TypeCreate *string

	// TypeDelete.
	TypeDelete *string

	// TypeUpdate.
	TypeUpdate *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras webhooks list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasWebhooksListParams) WithDefaults() *ExtrasWebhooksListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras webhooks list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasWebhooksListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras webhooks list params
func (o *ExtrasWebhooksListParams) WithTimeout(timeout time.Duration) *ExtrasWebhooksListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras webhooks list params
func (o *ExtrasWebhooksListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras webhooks list params
func (o *ExtrasWebhooksListParams) WithContext(ctx context.Context) *ExtrasWebhooksListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras webhooks list params
func (o *ExtrasWebhooksListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras webhooks list params
func (o *ExtrasWebhooksListParams) WithHTTPClient(client *http.Client) *ExtrasWebhooksListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras webhooks list params
func (o *ExtrasWebhooksListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentTypes adds the contentTypes to the extras webhooks list params
func (o *ExtrasWebhooksListParams) WithContentTypes(contentTypes *string) *ExtrasWebhooksListParams {
	o.SetContentTypes(contentTypes)
	return o
}

// SetContentTypes adds the contentTypes to the extras webhooks list params
func (o *ExtrasWebhooksListParams) SetContentTypes(contentTypes *string) {
	o.ContentTypes = contentTypes
}

// WithContentTypesn adds the contentTypesn to the extras webhooks list params
func (o *ExtrasWebhooksListParams) WithContentTypesn(contentTypesn *string) *ExtrasWebhooksListParams {
	o.SetContentTypesn(contentTypesn)
	return o
}

// SetContentTypesn adds the contentTypesN to the extras webhooks list params
func (o *ExtrasWebhooksListParams) SetContentTypesn(contentTypesn *string) {
	o.ContentTypesn = contentTypesn
}

// WithEnabled adds the enabled to the extras webhooks list params
func (o *ExtrasWebhooksListParams) WithEnabled(enabled *string) *ExtrasWebhooksListParams {
	o.SetEnabled(enabled)
	return o
}

// SetEnabled adds the enabled to the extras webhooks list params
func (o *ExtrasWebhooksListParams) SetEnabled(enabled *string) {
	o.Enabled = enabled
}

// WithLimit adds the limit to the extras webhooks list params
func (o *ExtrasWebhooksListParams) WithLimit(limit *int64) *ExtrasWebhooksListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the extras webhooks list params
func (o *ExtrasWebhooksListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the extras webhooks list params
func (o *ExtrasWebhooksListParams) WithName(name *string) *ExtrasWebhooksListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the extras webhooks list params
func (o *ExtrasWebhooksListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the extras webhooks list params
func (o *ExtrasWebhooksListParams) WithNameIc(nameIc *string) *ExtrasWebhooksListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the extras webhooks list params
func (o *ExtrasWebhooksListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the extras webhooks list params
func (o *ExtrasWebhooksListParams) WithNameIe(nameIe *string) *ExtrasWebhooksListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the extras webhooks list params
func (o *ExtrasWebhooksListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the extras webhooks list params
func (o *ExtrasWebhooksListParams) WithNameIew(nameIew *string) *ExtrasWebhooksListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the extras webhooks list params
func (o *ExtrasWebhooksListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the extras webhooks list params
func (o *ExtrasWebhooksListParams) WithNameIsw(nameIsw *string) *ExtrasWebhooksListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the extras webhooks list params
func (o *ExtrasWebhooksListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the extras webhooks list params
func (o *ExtrasWebhooksListParams) WithNamen(namen *string) *ExtrasWebhooksListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the extras webhooks list params
func (o *ExtrasWebhooksListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the extras webhooks list params
func (o *ExtrasWebhooksListParams) WithNameNic(nameNic *string) *ExtrasWebhooksListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the extras webhooks list params
func (o *ExtrasWebhooksListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the extras webhooks list params
func (o *ExtrasWebhooksListParams) WithNameNie(nameNie *string) *ExtrasWebhooksListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the extras webhooks list params
func (o *ExtrasWebhooksListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the extras webhooks list params
func (o *ExtrasWebhooksListParams) WithNameNiew(nameNiew *string) *ExtrasWebhooksListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the extras webhooks list params
func (o *ExtrasWebhooksListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the extras webhooks list params
func (o *ExtrasWebhooksListParams) WithNameNisw(nameNisw *string) *ExtrasWebhooksListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the extras webhooks list params
func (o *ExtrasWebhooksListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the extras webhooks list params
func (o *ExtrasWebhooksListParams) WithOffset(offset *int64) *ExtrasWebhooksListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the extras webhooks list params
func (o *ExtrasWebhooksListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPayloadURL adds the payloadURL to the extras webhooks list params
func (o *ExtrasWebhooksListParams) WithPayloadURL(payloadURL *string) *ExtrasWebhooksListParams {
	o.SetPayloadURL(payloadURL)
	return o
}

// SetPayloadURL adds the payloadUrl to the extras webhooks list params
func (o *ExtrasWebhooksListParams) SetPayloadURL(payloadURL *string) {
	o.PayloadURL = payloadURL
}

// WithPayloadURLIc adds the payloadURLIc to the extras webhooks list params
func (o *ExtrasWebhooksListParams) WithPayloadURLIc(payloadURLIc *string) *ExtrasWebhooksListParams {
	o.SetPayloadURLIc(payloadURLIc)
	return o
}

// SetPayloadURLIc adds the payloadUrlIc to the extras webhooks list params
func (o *ExtrasWebhooksListParams) SetPayloadURLIc(payloadURLIc *string) {
	o.PayloadURLIc = payloadURLIc
}

// WithPayloadURLIe adds the payloadURLIe to the extras webhooks list params
func (o *ExtrasWebhooksListParams) WithPayloadURLIe(payloadURLIe *string) *ExtrasWebhooksListParams {
	o.SetPayloadURLIe(payloadURLIe)
	return o
}

// SetPayloadURLIe adds the payloadUrlIe to the extras webhooks list params
func (o *ExtrasWebhooksListParams) SetPayloadURLIe(payloadURLIe *string) {
	o.PayloadURLIe = payloadURLIe
}

// WithPayloadURLIew adds the payloadURLIew to the extras webhooks list params
func (o *ExtrasWebhooksListParams) WithPayloadURLIew(payloadURLIew *string) *ExtrasWebhooksListParams {
	o.SetPayloadURLIew(payloadURLIew)
	return o
}

// SetPayloadURLIew adds the payloadUrlIew to the extras webhooks list params
func (o *ExtrasWebhooksListParams) SetPayloadURLIew(payloadURLIew *string) {
	o.PayloadURLIew = payloadURLIew
}

// WithPayloadURLIsw adds the payloadURLIsw to the extras webhooks list params
func (o *ExtrasWebhooksListParams) WithPayloadURLIsw(payloadURLIsw *string) *ExtrasWebhooksListParams {
	o.SetPayloadURLIsw(payloadURLIsw)
	return o
}

// SetPayloadURLIsw adds the payloadUrlIsw to the extras webhooks list params
func (o *ExtrasWebhooksListParams) SetPayloadURLIsw(payloadURLIsw *string) {
	o.PayloadURLIsw = payloadURLIsw
}

// WithPayloadURLn adds the payloadURLn to the extras webhooks list params
func (o *ExtrasWebhooksListParams) WithPayloadURLn(payloadURLn *string) *ExtrasWebhooksListParams {
	o.SetPayloadURLn(payloadURLn)
	return o
}

// SetPayloadURLn adds the payloadUrlN to the extras webhooks list params
func (o *ExtrasWebhooksListParams) SetPayloadURLn(payloadURLn *string) {
	o.PayloadURLn = payloadURLn
}

// WithPayloadURLNic adds the payloadURLNic to the extras webhooks list params
func (o *ExtrasWebhooksListParams) WithPayloadURLNic(payloadURLNic *string) *ExtrasWebhooksListParams {
	o.SetPayloadURLNic(payloadURLNic)
	return o
}

// SetPayloadURLNic adds the payloadUrlNic to the extras webhooks list params
func (o *ExtrasWebhooksListParams) SetPayloadURLNic(payloadURLNic *string) {
	o.PayloadURLNic = payloadURLNic
}

// WithPayloadURLNie adds the payloadURLNie to the extras webhooks list params
func (o *ExtrasWebhooksListParams) WithPayloadURLNie(payloadURLNie *string) *ExtrasWebhooksListParams {
	o.SetPayloadURLNie(payloadURLNie)
	return o
}

// SetPayloadURLNie adds the payloadUrlNie to the extras webhooks list params
func (o *ExtrasWebhooksListParams) SetPayloadURLNie(payloadURLNie *string) {
	o.PayloadURLNie = payloadURLNie
}

// WithPayloadURLNiew adds the payloadURLNiew to the extras webhooks list params
func (o *ExtrasWebhooksListParams) WithPayloadURLNiew(payloadURLNiew *string) *ExtrasWebhooksListParams {
	o.SetPayloadURLNiew(payloadURLNiew)
	return o
}

// SetPayloadURLNiew adds the payloadUrlNiew to the extras webhooks list params
func (o *ExtrasWebhooksListParams) SetPayloadURLNiew(payloadURLNiew *string) {
	o.PayloadURLNiew = payloadURLNiew
}

// WithPayloadURLNisw adds the payloadURLNisw to the extras webhooks list params
func (o *ExtrasWebhooksListParams) WithPayloadURLNisw(payloadURLNisw *string) *ExtrasWebhooksListParams {
	o.SetPayloadURLNisw(payloadURLNisw)
	return o
}

// SetPayloadURLNisw adds the payloadUrlNisw to the extras webhooks list params
func (o *ExtrasWebhooksListParams) SetPayloadURLNisw(payloadURLNisw *string) {
	o.PayloadURLNisw = payloadURLNisw
}

// WithQ adds the q to the extras webhooks list params
func (o *ExtrasWebhooksListParams) WithQ(q *string) *ExtrasWebhooksListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the extras webhooks list params
func (o *ExtrasWebhooksListParams) SetQ(q *string) {
	o.Q = q
}

// WithTypeCreate adds the typeCreate to the extras webhooks list params
func (o *ExtrasWebhooksListParams) WithTypeCreate(typeCreate *string) *ExtrasWebhooksListParams {
	o.SetTypeCreate(typeCreate)
	return o
}

// SetTypeCreate adds the typeCreate to the extras webhooks list params
func (o *ExtrasWebhooksListParams) SetTypeCreate(typeCreate *string) {
	o.TypeCreate = typeCreate
}

// WithTypeDelete adds the typeDelete to the extras webhooks list params
func (o *ExtrasWebhooksListParams) WithTypeDelete(typeDelete *string) *ExtrasWebhooksListParams {
	o.SetTypeDelete(typeDelete)
	return o
}

// SetTypeDelete adds the typeDelete to the extras webhooks list params
func (o *ExtrasWebhooksListParams) SetTypeDelete(typeDelete *string) {
	o.TypeDelete = typeDelete
}

// WithTypeUpdate adds the typeUpdate to the extras webhooks list params
func (o *ExtrasWebhooksListParams) WithTypeUpdate(typeUpdate *string) *ExtrasWebhooksListParams {
	o.SetTypeUpdate(typeUpdate)
	return o
}

// SetTypeUpdate adds the typeUpdate to the extras webhooks list params
func (o *ExtrasWebhooksListParams) SetTypeUpdate(typeUpdate *string) {
	o.TypeUpdate = typeUpdate
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasWebhooksListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Enabled != nil {

		// query param enabled
		var qrEnabled string

		if o.Enabled != nil {
			qrEnabled = *o.Enabled
		}
		qEnabled := qrEnabled
		if qEnabled != "" {

			if err := r.SetQueryParam("enabled", qEnabled); err != nil {
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

	if o.PayloadURL != nil {

		// query param payload_url
		var qrPayloadURL string

		if o.PayloadURL != nil {
			qrPayloadURL = *o.PayloadURL
		}
		qPayloadURL := qrPayloadURL
		if qPayloadURL != "" {

			if err := r.SetQueryParam("payload_url", qPayloadURL); err != nil {
				return err
			}
		}
	}

	if o.PayloadURLIc != nil {

		// query param payload_url__ic
		var qrPayloadURLIc string

		if o.PayloadURLIc != nil {
			qrPayloadURLIc = *o.PayloadURLIc
		}
		qPayloadURLIc := qrPayloadURLIc
		if qPayloadURLIc != "" {

			if err := r.SetQueryParam("payload_url__ic", qPayloadURLIc); err != nil {
				return err
			}
		}
	}

	if o.PayloadURLIe != nil {

		// query param payload_url__ie
		var qrPayloadURLIe string

		if o.PayloadURLIe != nil {
			qrPayloadURLIe = *o.PayloadURLIe
		}
		qPayloadURLIe := qrPayloadURLIe
		if qPayloadURLIe != "" {

			if err := r.SetQueryParam("payload_url__ie", qPayloadURLIe); err != nil {
				return err
			}
		}
	}

	if o.PayloadURLIew != nil {

		// query param payload_url__iew
		var qrPayloadURLIew string

		if o.PayloadURLIew != nil {
			qrPayloadURLIew = *o.PayloadURLIew
		}
		qPayloadURLIew := qrPayloadURLIew
		if qPayloadURLIew != "" {

			if err := r.SetQueryParam("payload_url__iew", qPayloadURLIew); err != nil {
				return err
			}
		}
	}

	if o.PayloadURLIsw != nil {

		// query param payload_url__isw
		var qrPayloadURLIsw string

		if o.PayloadURLIsw != nil {
			qrPayloadURLIsw = *o.PayloadURLIsw
		}
		qPayloadURLIsw := qrPayloadURLIsw
		if qPayloadURLIsw != "" {

			if err := r.SetQueryParam("payload_url__isw", qPayloadURLIsw); err != nil {
				return err
			}
		}
	}

	if o.PayloadURLn != nil {

		// query param payload_url__n
		var qrPayloadURLn string

		if o.PayloadURLn != nil {
			qrPayloadURLn = *o.PayloadURLn
		}
		qPayloadURLn := qrPayloadURLn
		if qPayloadURLn != "" {

			if err := r.SetQueryParam("payload_url__n", qPayloadURLn); err != nil {
				return err
			}
		}
	}

	if o.PayloadURLNic != nil {

		// query param payload_url__nic
		var qrPayloadURLNic string

		if o.PayloadURLNic != nil {
			qrPayloadURLNic = *o.PayloadURLNic
		}
		qPayloadURLNic := qrPayloadURLNic
		if qPayloadURLNic != "" {

			if err := r.SetQueryParam("payload_url__nic", qPayloadURLNic); err != nil {
				return err
			}
		}
	}

	if o.PayloadURLNie != nil {

		// query param payload_url__nie
		var qrPayloadURLNie string

		if o.PayloadURLNie != nil {
			qrPayloadURLNie = *o.PayloadURLNie
		}
		qPayloadURLNie := qrPayloadURLNie
		if qPayloadURLNie != "" {

			if err := r.SetQueryParam("payload_url__nie", qPayloadURLNie); err != nil {
				return err
			}
		}
	}

	if o.PayloadURLNiew != nil {

		// query param payload_url__niew
		var qrPayloadURLNiew string

		if o.PayloadURLNiew != nil {
			qrPayloadURLNiew = *o.PayloadURLNiew
		}
		qPayloadURLNiew := qrPayloadURLNiew
		if qPayloadURLNiew != "" {

			if err := r.SetQueryParam("payload_url__niew", qPayloadURLNiew); err != nil {
				return err
			}
		}
	}

	if o.PayloadURLNisw != nil {

		// query param payload_url__nisw
		var qrPayloadURLNisw string

		if o.PayloadURLNisw != nil {
			qrPayloadURLNisw = *o.PayloadURLNisw
		}
		qPayloadURLNisw := qrPayloadURLNisw
		if qPayloadURLNisw != "" {

			if err := r.SetQueryParam("payload_url__nisw", qPayloadURLNisw); err != nil {
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

	if o.TypeCreate != nil {

		// query param type_create
		var qrTypeCreate string

		if o.TypeCreate != nil {
			qrTypeCreate = *o.TypeCreate
		}
		qTypeCreate := qrTypeCreate
		if qTypeCreate != "" {

			if err := r.SetQueryParam("type_create", qTypeCreate); err != nil {
				return err
			}
		}
	}

	if o.TypeDelete != nil {

		// query param type_delete
		var qrTypeDelete string

		if o.TypeDelete != nil {
			qrTypeDelete = *o.TypeDelete
		}
		qTypeDelete := qrTypeDelete
		if qTypeDelete != "" {

			if err := r.SetQueryParam("type_delete", qTypeDelete); err != nil {
				return err
			}
		}
	}

	if o.TypeUpdate != nil {

		// query param type_update
		var qrTypeUpdate string

		if o.TypeUpdate != nil {
			qrTypeUpdate = *o.TypeUpdate
		}
		qTypeUpdate := qrTypeUpdate
		if qTypeUpdate != "" {

			if err := r.SetQueryParam("type_update", qTypeUpdate); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

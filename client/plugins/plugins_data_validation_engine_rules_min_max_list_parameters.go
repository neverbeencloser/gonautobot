package plugins

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

// NewPluginsDataValidationEngineRulesMinMaxListParams creates a new PluginsDataValidationEngineRulesMinMaxListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsDataValidationEngineRulesMinMaxListParams() *PluginsDataValidationEngineRulesMinMaxListParams {
	return &PluginsDataValidationEngineRulesMinMaxListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsDataValidationEngineRulesMinMaxListParamsWithTimeout creates a new PluginsDataValidationEngineRulesMinMaxListParams object
// with the ability to set a timeout on a request.
func NewPluginsDataValidationEngineRulesMinMaxListParamsWithTimeout(timeout time.Duration) *PluginsDataValidationEngineRulesMinMaxListParams {
	return &PluginsDataValidationEngineRulesMinMaxListParams{
		timeout: timeout,
	}
}

// NewPluginsDataValidationEngineRulesMinMaxListParamsWithContext creates a new PluginsDataValidationEngineRulesMinMaxListParams object
// with the ability to set a context for a request.
func NewPluginsDataValidationEngineRulesMinMaxListParamsWithContext(ctx context.Context) *PluginsDataValidationEngineRulesMinMaxListParams {
	return &PluginsDataValidationEngineRulesMinMaxListParams{
		Context: ctx,
	}
}

// NewPluginsDataValidationEngineRulesMinMaxListParamsWithHTTPClient creates a new PluginsDataValidationEngineRulesMinMaxListParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsDataValidationEngineRulesMinMaxListParamsWithHTTPClient(client *http.Client) *PluginsDataValidationEngineRulesMinMaxListParams {
	return &PluginsDataValidationEngineRulesMinMaxListParams{
		HTTPClient: client,
	}
}

/* PluginsDataValidationEngineRulesMinMaxListParams contains all the parameters to send to the API endpoint
   for the plugins data validation engine rules min max list operation.

   Typically these are written to a http.Request.
*/
type PluginsDataValidationEngineRulesMinMaxListParams struct {

	// ContentType.
	ContentType *string

	// ContentTypen.
	ContentTypen *string

	// Created.
	Created *string

	// CreatedGte.
	CreatedGte *string

	// CreatedLte.
	CreatedLte *string

	// Enabled.
	Enabled *string

	// ErrorMessage.
	ErrorMessage *string

	// ErrorMessageIc.
	ErrorMessageIc *string

	// ErrorMessageIe.
	ErrorMessageIe *string

	// ErrorMessageIew.
	ErrorMessageIew *string

	// ErrorMessageIsw.
	ErrorMessageIsw *string

	// ErrorMessagen.
	ErrorMessagen *string

	// ErrorMessageNic.
	ErrorMessageNic *string

	// ErrorMessageNie.
	ErrorMessageNie *string

	// ErrorMessageNiew.
	ErrorMessageNiew *string

	// ErrorMessageNisw.
	ErrorMessageNisw *string

	// Field.
	Field *string

	// FieldIc.
	FieldIc *string

	// FieldIe.
	FieldIe *string

	// FieldIew.
	FieldIew *string

	// FieldIsw.
	FieldIsw *string

	// Fieldn.
	Fieldn *string

	// FieldNic.
	FieldNic *string

	// FieldNie.
	FieldNie *string

	// FieldNiew.
	FieldNiew *string

	// FieldNisw.
	FieldNisw *string

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

	// LastUpdated.
	LastUpdated *string

	// LastUpdatedGte.
	LastUpdatedGte *string

	// LastUpdatedLte.
	LastUpdatedLte *string

	/* Limit.

	   Number of results to return per page.
	*/
	Limit *int64

	// Max.
	Max *string

	// MaxGt.
	MaxGt *string

	// MaxGte.
	MaxGte *string

	// MaxLt.
	MaxLt *string

	// MaxLte.
	MaxLte *string

	// Maxn.
	Maxn *string

	// Min.
	Min *string

	// MinGt.
	MinGt *string

	// MinGte.
	MinGte *string

	// MinLt.
	MinLt *string

	// MinLte.
	MinLte *string

	// Minn.
	Minn *string

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

	// Slug.
	Slug *string

	// SlugIc.
	SlugIc *string

	// SlugIe.
	SlugIe *string

	// SlugIew.
	SlugIew *string

	// SlugIsw.
	SlugIsw *string

	// Slugn.
	Slugn *string

	// SlugNic.
	SlugNic *string

	// SlugNie.
	SlugNie *string

	// SlugNiew.
	SlugNiew *string

	// SlugNisw.
	SlugNisw *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins data validation engine rules min max list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithDefaults() *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins data validation engine rules min max list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithTimeout(timeout time.Duration) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithContext(ctx context.Context) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithHTTPClient(client *http.Client) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentType adds the contentType to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithContentType(contentType *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetContentType(contentType *string) {
	o.ContentType = contentType
}

// WithContentTypen adds the contentTypen to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithContentTypen(contentTypen *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetContentTypen(contentTypen)
	return o
}

// SetContentTypen adds the contentTypeN to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetContentTypen(contentTypen *string) {
	o.ContentTypen = contentTypen
}

// WithCreated adds the created to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithCreated(created *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGte adds the createdGte to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithCreatedGte(createdGte *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLte adds the createdLte to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithCreatedLte(createdLte *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithEnabled adds the enabled to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithEnabled(enabled *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetEnabled(enabled)
	return o
}

// SetEnabled adds the enabled to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetEnabled(enabled *string) {
	o.Enabled = enabled
}

// WithErrorMessage adds the errorMessage to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithErrorMessage(errorMessage *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetErrorMessage(errorMessage)
	return o
}

// SetErrorMessage adds the errorMessage to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetErrorMessage(errorMessage *string) {
	o.ErrorMessage = errorMessage
}

// WithErrorMessageIc adds the errorMessageIc to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithErrorMessageIc(errorMessageIc *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetErrorMessageIc(errorMessageIc)
	return o
}

// SetErrorMessageIc adds the errorMessageIc to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetErrorMessageIc(errorMessageIc *string) {
	o.ErrorMessageIc = errorMessageIc
}

// WithErrorMessageIe adds the errorMessageIe to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithErrorMessageIe(errorMessageIe *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetErrorMessageIe(errorMessageIe)
	return o
}

// SetErrorMessageIe adds the errorMessageIe to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetErrorMessageIe(errorMessageIe *string) {
	o.ErrorMessageIe = errorMessageIe
}

// WithErrorMessageIew adds the errorMessageIew to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithErrorMessageIew(errorMessageIew *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetErrorMessageIew(errorMessageIew)
	return o
}

// SetErrorMessageIew adds the errorMessageIew to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetErrorMessageIew(errorMessageIew *string) {
	o.ErrorMessageIew = errorMessageIew
}

// WithErrorMessageIsw adds the errorMessageIsw to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithErrorMessageIsw(errorMessageIsw *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetErrorMessageIsw(errorMessageIsw)
	return o
}

// SetErrorMessageIsw adds the errorMessageIsw to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetErrorMessageIsw(errorMessageIsw *string) {
	o.ErrorMessageIsw = errorMessageIsw
}

// WithErrorMessagen adds the errorMessagen to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithErrorMessagen(errorMessagen *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetErrorMessagen(errorMessagen)
	return o
}

// SetErrorMessagen adds the errorMessageN to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetErrorMessagen(errorMessagen *string) {
	o.ErrorMessagen = errorMessagen
}

// WithErrorMessageNic adds the errorMessageNic to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithErrorMessageNic(errorMessageNic *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetErrorMessageNic(errorMessageNic)
	return o
}

// SetErrorMessageNic adds the errorMessageNic to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetErrorMessageNic(errorMessageNic *string) {
	o.ErrorMessageNic = errorMessageNic
}

// WithErrorMessageNie adds the errorMessageNie to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithErrorMessageNie(errorMessageNie *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetErrorMessageNie(errorMessageNie)
	return o
}

// SetErrorMessageNie adds the errorMessageNie to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetErrorMessageNie(errorMessageNie *string) {
	o.ErrorMessageNie = errorMessageNie
}

// WithErrorMessageNiew adds the errorMessageNiew to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithErrorMessageNiew(errorMessageNiew *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetErrorMessageNiew(errorMessageNiew)
	return o
}

// SetErrorMessageNiew adds the errorMessageNiew to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetErrorMessageNiew(errorMessageNiew *string) {
	o.ErrorMessageNiew = errorMessageNiew
}

// WithErrorMessageNisw adds the errorMessageNisw to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithErrorMessageNisw(errorMessageNisw *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetErrorMessageNisw(errorMessageNisw)
	return o
}

// SetErrorMessageNisw adds the errorMessageNisw to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetErrorMessageNisw(errorMessageNisw *string) {
	o.ErrorMessageNisw = errorMessageNisw
}

// WithField adds the field to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithField(field *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetField(field)
	return o
}

// SetField adds the field to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetField(field *string) {
	o.Field = field
}

// WithFieldIc adds the fieldIc to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithFieldIc(fieldIc *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetFieldIc(fieldIc)
	return o
}

// SetFieldIc adds the fieldIc to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetFieldIc(fieldIc *string) {
	o.FieldIc = fieldIc
}

// WithFieldIe adds the fieldIe to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithFieldIe(fieldIe *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetFieldIe(fieldIe)
	return o
}

// SetFieldIe adds the fieldIe to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetFieldIe(fieldIe *string) {
	o.FieldIe = fieldIe
}

// WithFieldIew adds the fieldIew to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithFieldIew(fieldIew *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetFieldIew(fieldIew)
	return o
}

// SetFieldIew adds the fieldIew to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetFieldIew(fieldIew *string) {
	o.FieldIew = fieldIew
}

// WithFieldIsw adds the fieldIsw to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithFieldIsw(fieldIsw *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetFieldIsw(fieldIsw)
	return o
}

// SetFieldIsw adds the fieldIsw to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetFieldIsw(fieldIsw *string) {
	o.FieldIsw = fieldIsw
}

// WithFieldn adds the fieldn to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithFieldn(fieldn *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetFieldn(fieldn)
	return o
}

// SetFieldn adds the fieldN to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetFieldn(fieldn *string) {
	o.Fieldn = fieldn
}

// WithFieldNic adds the fieldNic to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithFieldNic(fieldNic *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetFieldNic(fieldNic)
	return o
}

// SetFieldNic adds the fieldNic to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetFieldNic(fieldNic *string) {
	o.FieldNic = fieldNic
}

// WithFieldNie adds the fieldNie to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithFieldNie(fieldNie *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetFieldNie(fieldNie)
	return o
}

// SetFieldNie adds the fieldNie to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetFieldNie(fieldNie *string) {
	o.FieldNie = fieldNie
}

// WithFieldNiew adds the fieldNiew to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithFieldNiew(fieldNiew *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetFieldNiew(fieldNiew)
	return o
}

// SetFieldNiew adds the fieldNiew to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetFieldNiew(fieldNiew *string) {
	o.FieldNiew = fieldNiew
}

// WithFieldNisw adds the fieldNisw to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithFieldNisw(fieldNisw *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetFieldNisw(fieldNisw)
	return o
}

// SetFieldNisw adds the fieldNisw to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetFieldNisw(fieldNisw *string) {
	o.FieldNisw = fieldNisw
}

// WithID adds the id to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithID(id *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithIDIc(iDIc *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithIDIe(iDIe *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithIDIew(iDIew *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithIDIsw(iDIsw *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithIDn(iDn *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithIDNic(iDNic *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithIDNie(iDNie *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithIDNiew(iDNiew *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithIDNisw(iDNisw *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithLastUpdated adds the lastUpdated to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithLastUpdated(lastUpdated *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetLastUpdated(lastUpdated *string) {
	o.LastUpdated = lastUpdated
}

// WithLastUpdatedGte adds the lastUpdatedGte to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithLastUpdatedGte(lastUpdatedGte *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetLastUpdatedGte(lastUpdatedGte)
	return o
}

// SetLastUpdatedGte adds the lastUpdatedGte to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetLastUpdatedGte(lastUpdatedGte *string) {
	o.LastUpdatedGte = lastUpdatedGte
}

// WithLastUpdatedLte adds the lastUpdatedLte to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithLastUpdatedLte(lastUpdatedLte *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetLastUpdatedLte(lastUpdatedLte)
	return o
}

// SetLastUpdatedLte adds the lastUpdatedLte to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetLastUpdatedLte(lastUpdatedLte *string) {
	o.LastUpdatedLte = lastUpdatedLte
}

// WithLimit adds the limit to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithLimit(limit *int64) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithMax adds the max to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithMax(max *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetMax(max)
	return o
}

// SetMax adds the max to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetMax(max *string) {
	o.Max = max
}

// WithMaxGt adds the maxGt to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithMaxGt(maxGt *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetMaxGt(maxGt)
	return o
}

// SetMaxGt adds the maxGt to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetMaxGt(maxGt *string) {
	o.MaxGt = maxGt
}

// WithMaxGte adds the maxGte to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithMaxGte(maxGte *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetMaxGte(maxGte)
	return o
}

// SetMaxGte adds the maxGte to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetMaxGte(maxGte *string) {
	o.MaxGte = maxGte
}

// WithMaxLt adds the maxLt to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithMaxLt(maxLt *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetMaxLt(maxLt)
	return o
}

// SetMaxLt adds the maxLt to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetMaxLt(maxLt *string) {
	o.MaxLt = maxLt
}

// WithMaxLte adds the maxLte to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithMaxLte(maxLte *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetMaxLte(maxLte)
	return o
}

// SetMaxLte adds the maxLte to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetMaxLte(maxLte *string) {
	o.MaxLte = maxLte
}

// WithMaxn adds the maxn to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithMaxn(maxn *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetMaxn(maxn)
	return o
}

// SetMaxn adds the maxN to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetMaxn(maxn *string) {
	o.Maxn = maxn
}

// WithMin adds the min to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithMin(min *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetMin(min)
	return o
}

// SetMin adds the min to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetMin(min *string) {
	o.Min = min
}

// WithMinGt adds the minGt to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithMinGt(minGt *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetMinGt(minGt)
	return o
}

// SetMinGt adds the minGt to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetMinGt(minGt *string) {
	o.MinGt = minGt
}

// WithMinGte adds the minGte to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithMinGte(minGte *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetMinGte(minGte)
	return o
}

// SetMinGte adds the minGte to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetMinGte(minGte *string) {
	o.MinGte = minGte
}

// WithMinLt adds the minLt to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithMinLt(minLt *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetMinLt(minLt)
	return o
}

// SetMinLt adds the minLt to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetMinLt(minLt *string) {
	o.MinLt = minLt
}

// WithMinLte adds the minLte to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithMinLte(minLte *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetMinLte(minLte)
	return o
}

// SetMinLte adds the minLte to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetMinLte(minLte *string) {
	o.MinLte = minLte
}

// WithMinn adds the minn to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithMinn(minn *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetMinn(minn)
	return o
}

// SetMinn adds the minN to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetMinn(minn *string) {
	o.Minn = minn
}

// WithName adds the name to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithName(name *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithNameIc(nameIc *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithNameIe(nameIe *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithNameIew(nameIew *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithNameIsw(nameIsw *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithNamen(namen *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithNameNic(nameNic *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithNameNie(nameNie *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithNameNiew(nameNiew *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithNameNisw(nameNisw *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithOffset(offset *int64) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithQ(q *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetQ(q *string) {
	o.Q = q
}

// WithSlug adds the slug to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithSlug(slug *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetSlug(slug *string) {
	o.Slug = slug
}

// WithSlugIc adds the slugIc to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithSlugIc(slugIc *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetSlugIc(slugIc)
	return o
}

// SetSlugIc adds the slugIc to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetSlugIc(slugIc *string) {
	o.SlugIc = slugIc
}

// WithSlugIe adds the slugIe to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithSlugIe(slugIe *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetSlugIe(slugIe)
	return o
}

// SetSlugIe adds the slugIe to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetSlugIe(slugIe *string) {
	o.SlugIe = slugIe
}

// WithSlugIew adds the slugIew to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithSlugIew(slugIew *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetSlugIew(slugIew)
	return o
}

// SetSlugIew adds the slugIew to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetSlugIew(slugIew *string) {
	o.SlugIew = slugIew
}

// WithSlugIsw adds the slugIsw to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithSlugIsw(slugIsw *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetSlugIsw(slugIsw)
	return o
}

// SetSlugIsw adds the slugIsw to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetSlugIsw(slugIsw *string) {
	o.SlugIsw = slugIsw
}

// WithSlugn adds the slugn to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithSlugn(slugn *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetSlugn(slugn)
	return o
}

// SetSlugn adds the slugN to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetSlugn(slugn *string) {
	o.Slugn = slugn
}

// WithSlugNic adds the slugNic to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithSlugNic(slugNic *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetSlugNic(slugNic)
	return o
}

// SetSlugNic adds the slugNic to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetSlugNic(slugNic *string) {
	o.SlugNic = slugNic
}

// WithSlugNie adds the slugNie to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithSlugNie(slugNie *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetSlugNie(slugNie)
	return o
}

// SetSlugNie adds the slugNie to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetSlugNie(slugNie *string) {
	o.SlugNie = slugNie
}

// WithSlugNiew adds the slugNiew to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithSlugNiew(slugNiew *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetSlugNiew(slugNiew)
	return o
}

// SetSlugNiew adds the slugNiew to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetSlugNiew(slugNiew *string) {
	o.SlugNiew = slugNiew
}

// WithSlugNisw adds the slugNisw to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WithSlugNisw(slugNisw *string) *PluginsDataValidationEngineRulesMinMaxListParams {
	o.SetSlugNisw(slugNisw)
	return o
}

// SetSlugNisw adds the slugNisw to the plugins data validation engine rules min max list params
func (o *PluginsDataValidationEngineRulesMinMaxListParams) SetSlugNisw(slugNisw *string) {
	o.SlugNisw = slugNisw
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsDataValidationEngineRulesMinMaxListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Created != nil {

		// query param created
		var qrCreated string

		if o.Created != nil {
			qrCreated = *o.Created
		}
		qCreated := qrCreated
		if qCreated != "" {

			if err := r.SetQueryParam("created", qCreated); err != nil {
				return err
			}
		}
	}

	if o.CreatedGte != nil {

		// query param created__gte
		var qrCreatedGte string

		if o.CreatedGte != nil {
			qrCreatedGte = *o.CreatedGte
		}
		qCreatedGte := qrCreatedGte
		if qCreatedGte != "" {

			if err := r.SetQueryParam("created__gte", qCreatedGte); err != nil {
				return err
			}
		}
	}

	if o.CreatedLte != nil {

		// query param created__lte
		var qrCreatedLte string

		if o.CreatedLte != nil {
			qrCreatedLte = *o.CreatedLte
		}
		qCreatedLte := qrCreatedLte
		if qCreatedLte != "" {

			if err := r.SetQueryParam("created__lte", qCreatedLte); err != nil {
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

	if o.ErrorMessage != nil {

		// query param error_message
		var qrErrorMessage string

		if o.ErrorMessage != nil {
			qrErrorMessage = *o.ErrorMessage
		}
		qErrorMessage := qrErrorMessage
		if qErrorMessage != "" {

			if err := r.SetQueryParam("error_message", qErrorMessage); err != nil {
				return err
			}
		}
	}

	if o.ErrorMessageIc != nil {

		// query param error_message__ic
		var qrErrorMessageIc string

		if o.ErrorMessageIc != nil {
			qrErrorMessageIc = *o.ErrorMessageIc
		}
		qErrorMessageIc := qrErrorMessageIc
		if qErrorMessageIc != "" {

			if err := r.SetQueryParam("error_message__ic", qErrorMessageIc); err != nil {
				return err
			}
		}
	}

	if o.ErrorMessageIe != nil {

		// query param error_message__ie
		var qrErrorMessageIe string

		if o.ErrorMessageIe != nil {
			qrErrorMessageIe = *o.ErrorMessageIe
		}
		qErrorMessageIe := qrErrorMessageIe
		if qErrorMessageIe != "" {

			if err := r.SetQueryParam("error_message__ie", qErrorMessageIe); err != nil {
				return err
			}
		}
	}

	if o.ErrorMessageIew != nil {

		// query param error_message__iew
		var qrErrorMessageIew string

		if o.ErrorMessageIew != nil {
			qrErrorMessageIew = *o.ErrorMessageIew
		}
		qErrorMessageIew := qrErrorMessageIew
		if qErrorMessageIew != "" {

			if err := r.SetQueryParam("error_message__iew", qErrorMessageIew); err != nil {
				return err
			}
		}
	}

	if o.ErrorMessageIsw != nil {

		// query param error_message__isw
		var qrErrorMessageIsw string

		if o.ErrorMessageIsw != nil {
			qrErrorMessageIsw = *o.ErrorMessageIsw
		}
		qErrorMessageIsw := qrErrorMessageIsw
		if qErrorMessageIsw != "" {

			if err := r.SetQueryParam("error_message__isw", qErrorMessageIsw); err != nil {
				return err
			}
		}
	}

	if o.ErrorMessagen != nil {

		// query param error_message__n
		var qrErrorMessagen string

		if o.ErrorMessagen != nil {
			qrErrorMessagen = *o.ErrorMessagen
		}
		qErrorMessagen := qrErrorMessagen
		if qErrorMessagen != "" {

			if err := r.SetQueryParam("error_message__n", qErrorMessagen); err != nil {
				return err
			}
		}
	}

	if o.ErrorMessageNic != nil {

		// query param error_message__nic
		var qrErrorMessageNic string

		if o.ErrorMessageNic != nil {
			qrErrorMessageNic = *o.ErrorMessageNic
		}
		qErrorMessageNic := qrErrorMessageNic
		if qErrorMessageNic != "" {

			if err := r.SetQueryParam("error_message__nic", qErrorMessageNic); err != nil {
				return err
			}
		}
	}

	if o.ErrorMessageNie != nil {

		// query param error_message__nie
		var qrErrorMessageNie string

		if o.ErrorMessageNie != nil {
			qrErrorMessageNie = *o.ErrorMessageNie
		}
		qErrorMessageNie := qrErrorMessageNie
		if qErrorMessageNie != "" {

			if err := r.SetQueryParam("error_message__nie", qErrorMessageNie); err != nil {
				return err
			}
		}
	}

	if o.ErrorMessageNiew != nil {

		// query param error_message__niew
		var qrErrorMessageNiew string

		if o.ErrorMessageNiew != nil {
			qrErrorMessageNiew = *o.ErrorMessageNiew
		}
		qErrorMessageNiew := qrErrorMessageNiew
		if qErrorMessageNiew != "" {

			if err := r.SetQueryParam("error_message__niew", qErrorMessageNiew); err != nil {
				return err
			}
		}
	}

	if o.ErrorMessageNisw != nil {

		// query param error_message__nisw
		var qrErrorMessageNisw string

		if o.ErrorMessageNisw != nil {
			qrErrorMessageNisw = *o.ErrorMessageNisw
		}
		qErrorMessageNisw := qrErrorMessageNisw
		if qErrorMessageNisw != "" {

			if err := r.SetQueryParam("error_message__nisw", qErrorMessageNisw); err != nil {
				return err
			}
		}
	}

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

	if o.FieldIc != nil {

		// query param field__ic
		var qrFieldIc string

		if o.FieldIc != nil {
			qrFieldIc = *o.FieldIc
		}
		qFieldIc := qrFieldIc
		if qFieldIc != "" {

			if err := r.SetQueryParam("field__ic", qFieldIc); err != nil {
				return err
			}
		}
	}

	if o.FieldIe != nil {

		// query param field__ie
		var qrFieldIe string

		if o.FieldIe != nil {
			qrFieldIe = *o.FieldIe
		}
		qFieldIe := qrFieldIe
		if qFieldIe != "" {

			if err := r.SetQueryParam("field__ie", qFieldIe); err != nil {
				return err
			}
		}
	}

	if o.FieldIew != nil {

		// query param field__iew
		var qrFieldIew string

		if o.FieldIew != nil {
			qrFieldIew = *o.FieldIew
		}
		qFieldIew := qrFieldIew
		if qFieldIew != "" {

			if err := r.SetQueryParam("field__iew", qFieldIew); err != nil {
				return err
			}
		}
	}

	if o.FieldIsw != nil {

		// query param field__isw
		var qrFieldIsw string

		if o.FieldIsw != nil {
			qrFieldIsw = *o.FieldIsw
		}
		qFieldIsw := qrFieldIsw
		if qFieldIsw != "" {

			if err := r.SetQueryParam("field__isw", qFieldIsw); err != nil {
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

	if o.FieldNic != nil {

		// query param field__nic
		var qrFieldNic string

		if o.FieldNic != nil {
			qrFieldNic = *o.FieldNic
		}
		qFieldNic := qrFieldNic
		if qFieldNic != "" {

			if err := r.SetQueryParam("field__nic", qFieldNic); err != nil {
				return err
			}
		}
	}

	if o.FieldNie != nil {

		// query param field__nie
		var qrFieldNie string

		if o.FieldNie != nil {
			qrFieldNie = *o.FieldNie
		}
		qFieldNie := qrFieldNie
		if qFieldNie != "" {

			if err := r.SetQueryParam("field__nie", qFieldNie); err != nil {
				return err
			}
		}
	}

	if o.FieldNiew != nil {

		// query param field__niew
		var qrFieldNiew string

		if o.FieldNiew != nil {
			qrFieldNiew = *o.FieldNiew
		}
		qFieldNiew := qrFieldNiew
		if qFieldNiew != "" {

			if err := r.SetQueryParam("field__niew", qFieldNiew); err != nil {
				return err
			}
		}
	}

	if o.FieldNisw != nil {

		// query param field__nisw
		var qrFieldNisw string

		if o.FieldNisw != nil {
			qrFieldNisw = *o.FieldNisw
		}
		qFieldNisw := qrFieldNisw
		if qFieldNisw != "" {

			if err := r.SetQueryParam("field__nisw", qFieldNisw); err != nil {
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

	if o.LastUpdated != nil {

		// query param last_updated
		var qrLastUpdated string

		if o.LastUpdated != nil {
			qrLastUpdated = *o.LastUpdated
		}
		qLastUpdated := qrLastUpdated
		if qLastUpdated != "" {

			if err := r.SetQueryParam("last_updated", qLastUpdated); err != nil {
				return err
			}
		}
	}

	if o.LastUpdatedGte != nil {

		// query param last_updated__gte
		var qrLastUpdatedGte string

		if o.LastUpdatedGte != nil {
			qrLastUpdatedGte = *o.LastUpdatedGte
		}
		qLastUpdatedGte := qrLastUpdatedGte
		if qLastUpdatedGte != "" {

			if err := r.SetQueryParam("last_updated__gte", qLastUpdatedGte); err != nil {
				return err
			}
		}
	}

	if o.LastUpdatedLte != nil {

		// query param last_updated__lte
		var qrLastUpdatedLte string

		if o.LastUpdatedLte != nil {
			qrLastUpdatedLte = *o.LastUpdatedLte
		}
		qLastUpdatedLte := qrLastUpdatedLte
		if qLastUpdatedLte != "" {

			if err := r.SetQueryParam("last_updated__lte", qLastUpdatedLte); err != nil {
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

	if o.Max != nil {

		// query param max
		var qrMax string

		if o.Max != nil {
			qrMax = *o.Max
		}
		qMax := qrMax
		if qMax != "" {

			if err := r.SetQueryParam("max", qMax); err != nil {
				return err
			}
		}
	}

	if o.MaxGt != nil {

		// query param max__gt
		var qrMaxGt string

		if o.MaxGt != nil {
			qrMaxGt = *o.MaxGt
		}
		qMaxGt := qrMaxGt
		if qMaxGt != "" {

			if err := r.SetQueryParam("max__gt", qMaxGt); err != nil {
				return err
			}
		}
	}

	if o.MaxGte != nil {

		// query param max__gte
		var qrMaxGte string

		if o.MaxGte != nil {
			qrMaxGte = *o.MaxGte
		}
		qMaxGte := qrMaxGte
		if qMaxGte != "" {

			if err := r.SetQueryParam("max__gte", qMaxGte); err != nil {
				return err
			}
		}
	}

	if o.MaxLt != nil {

		// query param max__lt
		var qrMaxLt string

		if o.MaxLt != nil {
			qrMaxLt = *o.MaxLt
		}
		qMaxLt := qrMaxLt
		if qMaxLt != "" {

			if err := r.SetQueryParam("max__lt", qMaxLt); err != nil {
				return err
			}
		}
	}

	if o.MaxLte != nil {

		// query param max__lte
		var qrMaxLte string

		if o.MaxLte != nil {
			qrMaxLte = *o.MaxLte
		}
		qMaxLte := qrMaxLte
		if qMaxLte != "" {

			if err := r.SetQueryParam("max__lte", qMaxLte); err != nil {
				return err
			}
		}
	}

	if o.Maxn != nil {

		// query param max__n
		var qrMaxn string

		if o.Maxn != nil {
			qrMaxn = *o.Maxn
		}
		qMaxn := qrMaxn
		if qMaxn != "" {

			if err := r.SetQueryParam("max__n", qMaxn); err != nil {
				return err
			}
		}
	}

	if o.Min != nil {

		// query param min
		var qrMin string

		if o.Min != nil {
			qrMin = *o.Min
		}
		qMin := qrMin
		if qMin != "" {

			if err := r.SetQueryParam("min", qMin); err != nil {
				return err
			}
		}
	}

	if o.MinGt != nil {

		// query param min__gt
		var qrMinGt string

		if o.MinGt != nil {
			qrMinGt = *o.MinGt
		}
		qMinGt := qrMinGt
		if qMinGt != "" {

			if err := r.SetQueryParam("min__gt", qMinGt); err != nil {
				return err
			}
		}
	}

	if o.MinGte != nil {

		// query param min__gte
		var qrMinGte string

		if o.MinGte != nil {
			qrMinGte = *o.MinGte
		}
		qMinGte := qrMinGte
		if qMinGte != "" {

			if err := r.SetQueryParam("min__gte", qMinGte); err != nil {
				return err
			}
		}
	}

	if o.MinLt != nil {

		// query param min__lt
		var qrMinLt string

		if o.MinLt != nil {
			qrMinLt = *o.MinLt
		}
		qMinLt := qrMinLt
		if qMinLt != "" {

			if err := r.SetQueryParam("min__lt", qMinLt); err != nil {
				return err
			}
		}
	}

	if o.MinLte != nil {

		// query param min__lte
		var qrMinLte string

		if o.MinLte != nil {
			qrMinLte = *o.MinLte
		}
		qMinLte := qrMinLte
		if qMinLte != "" {

			if err := r.SetQueryParam("min__lte", qMinLte); err != nil {
				return err
			}
		}
	}

	if o.Minn != nil {

		// query param min__n
		var qrMinn string

		if o.Minn != nil {
			qrMinn = *o.Minn
		}
		qMinn := qrMinn
		if qMinn != "" {

			if err := r.SetQueryParam("min__n", qMinn); err != nil {
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

	if o.Slug != nil {

		// query param slug
		var qrSlug string

		if o.Slug != nil {
			qrSlug = *o.Slug
		}
		qSlug := qrSlug
		if qSlug != "" {

			if err := r.SetQueryParam("slug", qSlug); err != nil {
				return err
			}
		}
	}

	if o.SlugIc != nil {

		// query param slug__ic
		var qrSlugIc string

		if o.SlugIc != nil {
			qrSlugIc = *o.SlugIc
		}
		qSlugIc := qrSlugIc
		if qSlugIc != "" {

			if err := r.SetQueryParam("slug__ic", qSlugIc); err != nil {
				return err
			}
		}
	}

	if o.SlugIe != nil {

		// query param slug__ie
		var qrSlugIe string

		if o.SlugIe != nil {
			qrSlugIe = *o.SlugIe
		}
		qSlugIe := qrSlugIe
		if qSlugIe != "" {

			if err := r.SetQueryParam("slug__ie", qSlugIe); err != nil {
				return err
			}
		}
	}

	if o.SlugIew != nil {

		// query param slug__iew
		var qrSlugIew string

		if o.SlugIew != nil {
			qrSlugIew = *o.SlugIew
		}
		qSlugIew := qrSlugIew
		if qSlugIew != "" {

			if err := r.SetQueryParam("slug__iew", qSlugIew); err != nil {
				return err
			}
		}
	}

	if o.SlugIsw != nil {

		// query param slug__isw
		var qrSlugIsw string

		if o.SlugIsw != nil {
			qrSlugIsw = *o.SlugIsw
		}
		qSlugIsw := qrSlugIsw
		if qSlugIsw != "" {

			if err := r.SetQueryParam("slug__isw", qSlugIsw); err != nil {
				return err
			}
		}
	}

	if o.Slugn != nil {

		// query param slug__n
		var qrSlugn string

		if o.Slugn != nil {
			qrSlugn = *o.Slugn
		}
		qSlugn := qrSlugn
		if qSlugn != "" {

			if err := r.SetQueryParam("slug__n", qSlugn); err != nil {
				return err
			}
		}
	}

	if o.SlugNic != nil {

		// query param slug__nic
		var qrSlugNic string

		if o.SlugNic != nil {
			qrSlugNic = *o.SlugNic
		}
		qSlugNic := qrSlugNic
		if qSlugNic != "" {

			if err := r.SetQueryParam("slug__nic", qSlugNic); err != nil {
				return err
			}
		}
	}

	if o.SlugNie != nil {

		// query param slug__nie
		var qrSlugNie string

		if o.SlugNie != nil {
			qrSlugNie = *o.SlugNie
		}
		qSlugNie := qrSlugNie
		if qSlugNie != "" {

			if err := r.SetQueryParam("slug__nie", qSlugNie); err != nil {
				return err
			}
		}
	}

	if o.SlugNiew != nil {

		// query param slug__niew
		var qrSlugNiew string

		if o.SlugNiew != nil {
			qrSlugNiew = *o.SlugNiew
		}
		qSlugNiew := qrSlugNiew
		if qSlugNiew != "" {

			if err := r.SetQueryParam("slug__niew", qSlugNiew); err != nil {
				return err
			}
		}
	}

	if o.SlugNisw != nil {

		// query param slug__nisw
		var qrSlugNisw string

		if o.SlugNisw != nil {
			qrSlugNisw = *o.SlugNisw
		}
		qSlugNisw := qrSlugNisw
		if qSlugNisw != "" {

			if err := r.SetQueryParam("slug__nisw", qSlugNisw); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

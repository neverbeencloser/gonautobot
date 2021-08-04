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

// NewPluginsDataValidationEngineRulesRegexListParams creates a new PluginsDataValidationEngineRulesRegexListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsDataValidationEngineRulesRegexListParams() *PluginsDataValidationEngineRulesRegexListParams {
	return &PluginsDataValidationEngineRulesRegexListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsDataValidationEngineRulesRegexListParamsWithTimeout creates a new PluginsDataValidationEngineRulesRegexListParams object
// with the ability to set a timeout on a request.
func NewPluginsDataValidationEngineRulesRegexListParamsWithTimeout(timeout time.Duration) *PluginsDataValidationEngineRulesRegexListParams {
	return &PluginsDataValidationEngineRulesRegexListParams{
		timeout: timeout,
	}
}

// NewPluginsDataValidationEngineRulesRegexListParamsWithContext creates a new PluginsDataValidationEngineRulesRegexListParams object
// with the ability to set a context for a request.
func NewPluginsDataValidationEngineRulesRegexListParamsWithContext(ctx context.Context) *PluginsDataValidationEngineRulesRegexListParams {
	return &PluginsDataValidationEngineRulesRegexListParams{
		Context: ctx,
	}
}

// NewPluginsDataValidationEngineRulesRegexListParamsWithHTTPClient creates a new PluginsDataValidationEngineRulesRegexListParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsDataValidationEngineRulesRegexListParamsWithHTTPClient(client *http.Client) *PluginsDataValidationEngineRulesRegexListParams {
	return &PluginsDataValidationEngineRulesRegexListParams{
		HTTPClient: client,
	}
}

/* PluginsDataValidationEngineRulesRegexListParams contains all the parameters to send to the API endpoint
   for the plugins data validation engine rules regex list operation.

   Typically these are written to a http.Request.
*/
type PluginsDataValidationEngineRulesRegexListParams struct {

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

	// RegularExpression.
	RegularExpression *string

	// RegularExpressionIc.
	RegularExpressionIc *string

	// RegularExpressionIe.
	RegularExpressionIe *string

	// RegularExpressionIew.
	RegularExpressionIew *string

	// RegularExpressionIsw.
	RegularExpressionIsw *string

	// RegularExpressionn.
	RegularExpressionn *string

	// RegularExpressionNic.
	RegularExpressionNic *string

	// RegularExpressionNie.
	RegularExpressionNie *string

	// RegularExpressionNiew.
	RegularExpressionNiew *string

	// RegularExpressionNisw.
	RegularExpressionNisw *string

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

// WithDefaults hydrates default values in the plugins data validation engine rules regex list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsDataValidationEngineRulesRegexListParams) WithDefaults() *PluginsDataValidationEngineRulesRegexListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins data validation engine rules regex list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsDataValidationEngineRulesRegexListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithTimeout(timeout time.Duration) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithContext(ctx context.Context) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithHTTPClient(client *http.Client) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentType adds the contentType to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithContentType(contentType *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetContentType(contentType *string) {
	o.ContentType = contentType
}

// WithContentTypen adds the contentTypen to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithContentTypen(contentTypen *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetContentTypen(contentTypen)
	return o
}

// SetContentTypen adds the contentTypeN to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetContentTypen(contentTypen *string) {
	o.ContentTypen = contentTypen
}

// WithCreated adds the created to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithCreated(created *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGte adds the createdGte to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithCreatedGte(createdGte *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLte adds the createdLte to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithCreatedLte(createdLte *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithEnabled adds the enabled to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithEnabled(enabled *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetEnabled(enabled)
	return o
}

// SetEnabled adds the enabled to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetEnabled(enabled *string) {
	o.Enabled = enabled
}

// WithErrorMessage adds the errorMessage to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithErrorMessage(errorMessage *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetErrorMessage(errorMessage)
	return o
}

// SetErrorMessage adds the errorMessage to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetErrorMessage(errorMessage *string) {
	o.ErrorMessage = errorMessage
}

// WithErrorMessageIc adds the errorMessageIc to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithErrorMessageIc(errorMessageIc *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetErrorMessageIc(errorMessageIc)
	return o
}

// SetErrorMessageIc adds the errorMessageIc to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetErrorMessageIc(errorMessageIc *string) {
	o.ErrorMessageIc = errorMessageIc
}

// WithErrorMessageIe adds the errorMessageIe to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithErrorMessageIe(errorMessageIe *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetErrorMessageIe(errorMessageIe)
	return o
}

// SetErrorMessageIe adds the errorMessageIe to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetErrorMessageIe(errorMessageIe *string) {
	o.ErrorMessageIe = errorMessageIe
}

// WithErrorMessageIew adds the errorMessageIew to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithErrorMessageIew(errorMessageIew *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetErrorMessageIew(errorMessageIew)
	return o
}

// SetErrorMessageIew adds the errorMessageIew to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetErrorMessageIew(errorMessageIew *string) {
	o.ErrorMessageIew = errorMessageIew
}

// WithErrorMessageIsw adds the errorMessageIsw to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithErrorMessageIsw(errorMessageIsw *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetErrorMessageIsw(errorMessageIsw)
	return o
}

// SetErrorMessageIsw adds the errorMessageIsw to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetErrorMessageIsw(errorMessageIsw *string) {
	o.ErrorMessageIsw = errorMessageIsw
}

// WithErrorMessagen adds the errorMessagen to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithErrorMessagen(errorMessagen *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetErrorMessagen(errorMessagen)
	return o
}

// SetErrorMessagen adds the errorMessageN to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetErrorMessagen(errorMessagen *string) {
	o.ErrorMessagen = errorMessagen
}

// WithErrorMessageNic adds the errorMessageNic to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithErrorMessageNic(errorMessageNic *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetErrorMessageNic(errorMessageNic)
	return o
}

// SetErrorMessageNic adds the errorMessageNic to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetErrorMessageNic(errorMessageNic *string) {
	o.ErrorMessageNic = errorMessageNic
}

// WithErrorMessageNie adds the errorMessageNie to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithErrorMessageNie(errorMessageNie *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetErrorMessageNie(errorMessageNie)
	return o
}

// SetErrorMessageNie adds the errorMessageNie to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetErrorMessageNie(errorMessageNie *string) {
	o.ErrorMessageNie = errorMessageNie
}

// WithErrorMessageNiew adds the errorMessageNiew to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithErrorMessageNiew(errorMessageNiew *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetErrorMessageNiew(errorMessageNiew)
	return o
}

// SetErrorMessageNiew adds the errorMessageNiew to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetErrorMessageNiew(errorMessageNiew *string) {
	o.ErrorMessageNiew = errorMessageNiew
}

// WithErrorMessageNisw adds the errorMessageNisw to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithErrorMessageNisw(errorMessageNisw *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetErrorMessageNisw(errorMessageNisw)
	return o
}

// SetErrorMessageNisw adds the errorMessageNisw to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetErrorMessageNisw(errorMessageNisw *string) {
	o.ErrorMessageNisw = errorMessageNisw
}

// WithField adds the field to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithField(field *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetField(field)
	return o
}

// SetField adds the field to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetField(field *string) {
	o.Field = field
}

// WithFieldIc adds the fieldIc to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithFieldIc(fieldIc *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetFieldIc(fieldIc)
	return o
}

// SetFieldIc adds the fieldIc to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetFieldIc(fieldIc *string) {
	o.FieldIc = fieldIc
}

// WithFieldIe adds the fieldIe to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithFieldIe(fieldIe *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetFieldIe(fieldIe)
	return o
}

// SetFieldIe adds the fieldIe to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetFieldIe(fieldIe *string) {
	o.FieldIe = fieldIe
}

// WithFieldIew adds the fieldIew to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithFieldIew(fieldIew *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetFieldIew(fieldIew)
	return o
}

// SetFieldIew adds the fieldIew to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetFieldIew(fieldIew *string) {
	o.FieldIew = fieldIew
}

// WithFieldIsw adds the fieldIsw to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithFieldIsw(fieldIsw *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetFieldIsw(fieldIsw)
	return o
}

// SetFieldIsw adds the fieldIsw to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetFieldIsw(fieldIsw *string) {
	o.FieldIsw = fieldIsw
}

// WithFieldn adds the fieldn to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithFieldn(fieldn *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetFieldn(fieldn)
	return o
}

// SetFieldn adds the fieldN to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetFieldn(fieldn *string) {
	o.Fieldn = fieldn
}

// WithFieldNic adds the fieldNic to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithFieldNic(fieldNic *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetFieldNic(fieldNic)
	return o
}

// SetFieldNic adds the fieldNic to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetFieldNic(fieldNic *string) {
	o.FieldNic = fieldNic
}

// WithFieldNie adds the fieldNie to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithFieldNie(fieldNie *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetFieldNie(fieldNie)
	return o
}

// SetFieldNie adds the fieldNie to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetFieldNie(fieldNie *string) {
	o.FieldNie = fieldNie
}

// WithFieldNiew adds the fieldNiew to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithFieldNiew(fieldNiew *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetFieldNiew(fieldNiew)
	return o
}

// SetFieldNiew adds the fieldNiew to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetFieldNiew(fieldNiew *string) {
	o.FieldNiew = fieldNiew
}

// WithFieldNisw adds the fieldNisw to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithFieldNisw(fieldNisw *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetFieldNisw(fieldNisw)
	return o
}

// SetFieldNisw adds the fieldNisw to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetFieldNisw(fieldNisw *string) {
	o.FieldNisw = fieldNisw
}

// WithID adds the id to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithID(id *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithIDIc(iDIc *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithIDIe(iDIe *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithIDIew(iDIew *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithIDIsw(iDIsw *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithIDn(iDn *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithIDNic(iDNic *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithIDNie(iDNie *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithIDNiew(iDNiew *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithIDNisw(iDNisw *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithLastUpdated adds the lastUpdated to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithLastUpdated(lastUpdated *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetLastUpdated(lastUpdated *string) {
	o.LastUpdated = lastUpdated
}

// WithLastUpdatedGte adds the lastUpdatedGte to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithLastUpdatedGte(lastUpdatedGte *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetLastUpdatedGte(lastUpdatedGte)
	return o
}

// SetLastUpdatedGte adds the lastUpdatedGte to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetLastUpdatedGte(lastUpdatedGte *string) {
	o.LastUpdatedGte = lastUpdatedGte
}

// WithLastUpdatedLte adds the lastUpdatedLte to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithLastUpdatedLte(lastUpdatedLte *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetLastUpdatedLte(lastUpdatedLte)
	return o
}

// SetLastUpdatedLte adds the lastUpdatedLte to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetLastUpdatedLte(lastUpdatedLte *string) {
	o.LastUpdatedLte = lastUpdatedLte
}

// WithLimit adds the limit to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithLimit(limit *int64) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithName(name *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithNameIc(nameIc *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithNameIe(nameIe *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithNameIew(nameIew *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithNameIsw(nameIsw *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithNamen(namen *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithNameNic(nameNic *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithNameNie(nameNie *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithNameNiew(nameNiew *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithNameNisw(nameNisw *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithOffset(offset *int64) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithQ(q *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetQ(q *string) {
	o.Q = q
}

// WithRegularExpression adds the regularExpression to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithRegularExpression(regularExpression *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetRegularExpression(regularExpression)
	return o
}

// SetRegularExpression adds the regularExpression to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetRegularExpression(regularExpression *string) {
	o.RegularExpression = regularExpression
}

// WithRegularExpressionIc adds the regularExpressionIc to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithRegularExpressionIc(regularExpressionIc *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetRegularExpressionIc(regularExpressionIc)
	return o
}

// SetRegularExpressionIc adds the regularExpressionIc to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetRegularExpressionIc(regularExpressionIc *string) {
	o.RegularExpressionIc = regularExpressionIc
}

// WithRegularExpressionIe adds the regularExpressionIe to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithRegularExpressionIe(regularExpressionIe *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetRegularExpressionIe(regularExpressionIe)
	return o
}

// SetRegularExpressionIe adds the regularExpressionIe to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetRegularExpressionIe(regularExpressionIe *string) {
	o.RegularExpressionIe = regularExpressionIe
}

// WithRegularExpressionIew adds the regularExpressionIew to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithRegularExpressionIew(regularExpressionIew *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetRegularExpressionIew(regularExpressionIew)
	return o
}

// SetRegularExpressionIew adds the regularExpressionIew to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetRegularExpressionIew(regularExpressionIew *string) {
	o.RegularExpressionIew = regularExpressionIew
}

// WithRegularExpressionIsw adds the regularExpressionIsw to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithRegularExpressionIsw(regularExpressionIsw *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetRegularExpressionIsw(regularExpressionIsw)
	return o
}

// SetRegularExpressionIsw adds the regularExpressionIsw to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetRegularExpressionIsw(regularExpressionIsw *string) {
	o.RegularExpressionIsw = regularExpressionIsw
}

// WithRegularExpressionn adds the regularExpressionn to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithRegularExpressionn(regularExpressionn *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetRegularExpressionn(regularExpressionn)
	return o
}

// SetRegularExpressionn adds the regularExpressionN to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetRegularExpressionn(regularExpressionn *string) {
	o.RegularExpressionn = regularExpressionn
}

// WithRegularExpressionNic adds the regularExpressionNic to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithRegularExpressionNic(regularExpressionNic *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetRegularExpressionNic(regularExpressionNic)
	return o
}

// SetRegularExpressionNic adds the regularExpressionNic to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetRegularExpressionNic(regularExpressionNic *string) {
	o.RegularExpressionNic = regularExpressionNic
}

// WithRegularExpressionNie adds the regularExpressionNie to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithRegularExpressionNie(regularExpressionNie *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetRegularExpressionNie(regularExpressionNie)
	return o
}

// SetRegularExpressionNie adds the regularExpressionNie to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetRegularExpressionNie(regularExpressionNie *string) {
	o.RegularExpressionNie = regularExpressionNie
}

// WithRegularExpressionNiew adds the regularExpressionNiew to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithRegularExpressionNiew(regularExpressionNiew *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetRegularExpressionNiew(regularExpressionNiew)
	return o
}

// SetRegularExpressionNiew adds the regularExpressionNiew to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetRegularExpressionNiew(regularExpressionNiew *string) {
	o.RegularExpressionNiew = regularExpressionNiew
}

// WithRegularExpressionNisw adds the regularExpressionNisw to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithRegularExpressionNisw(regularExpressionNisw *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetRegularExpressionNisw(regularExpressionNisw)
	return o
}

// SetRegularExpressionNisw adds the regularExpressionNisw to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetRegularExpressionNisw(regularExpressionNisw *string) {
	o.RegularExpressionNisw = regularExpressionNisw
}

// WithSlug adds the slug to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithSlug(slug *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetSlug(slug *string) {
	o.Slug = slug
}

// WithSlugIc adds the slugIc to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithSlugIc(slugIc *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetSlugIc(slugIc)
	return o
}

// SetSlugIc adds the slugIc to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetSlugIc(slugIc *string) {
	o.SlugIc = slugIc
}

// WithSlugIe adds the slugIe to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithSlugIe(slugIe *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetSlugIe(slugIe)
	return o
}

// SetSlugIe adds the slugIe to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetSlugIe(slugIe *string) {
	o.SlugIe = slugIe
}

// WithSlugIew adds the slugIew to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithSlugIew(slugIew *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetSlugIew(slugIew)
	return o
}

// SetSlugIew adds the slugIew to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetSlugIew(slugIew *string) {
	o.SlugIew = slugIew
}

// WithSlugIsw adds the slugIsw to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithSlugIsw(slugIsw *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetSlugIsw(slugIsw)
	return o
}

// SetSlugIsw adds the slugIsw to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetSlugIsw(slugIsw *string) {
	o.SlugIsw = slugIsw
}

// WithSlugn adds the slugn to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithSlugn(slugn *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetSlugn(slugn)
	return o
}

// SetSlugn adds the slugN to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetSlugn(slugn *string) {
	o.Slugn = slugn
}

// WithSlugNic adds the slugNic to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithSlugNic(slugNic *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetSlugNic(slugNic)
	return o
}

// SetSlugNic adds the slugNic to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetSlugNic(slugNic *string) {
	o.SlugNic = slugNic
}

// WithSlugNie adds the slugNie to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithSlugNie(slugNie *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetSlugNie(slugNie)
	return o
}

// SetSlugNie adds the slugNie to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetSlugNie(slugNie *string) {
	o.SlugNie = slugNie
}

// WithSlugNiew adds the slugNiew to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithSlugNiew(slugNiew *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetSlugNiew(slugNiew)
	return o
}

// SetSlugNiew adds the slugNiew to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetSlugNiew(slugNiew *string) {
	o.SlugNiew = slugNiew
}

// WithSlugNisw adds the slugNisw to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) WithSlugNisw(slugNisw *string) *PluginsDataValidationEngineRulesRegexListParams {
	o.SetSlugNisw(slugNisw)
	return o
}

// SetSlugNisw adds the slugNisw to the plugins data validation engine rules regex list params
func (o *PluginsDataValidationEngineRulesRegexListParams) SetSlugNisw(slugNisw *string) {
	o.SlugNisw = slugNisw
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsDataValidationEngineRulesRegexListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.RegularExpression != nil {

		// query param regular_expression
		var qrRegularExpression string

		if o.RegularExpression != nil {
			qrRegularExpression = *o.RegularExpression
		}
		qRegularExpression := qrRegularExpression
		if qRegularExpression != "" {

			if err := r.SetQueryParam("regular_expression", qRegularExpression); err != nil {
				return err
			}
		}
	}

	if o.RegularExpressionIc != nil {

		// query param regular_expression__ic
		var qrRegularExpressionIc string

		if o.RegularExpressionIc != nil {
			qrRegularExpressionIc = *o.RegularExpressionIc
		}
		qRegularExpressionIc := qrRegularExpressionIc
		if qRegularExpressionIc != "" {

			if err := r.SetQueryParam("regular_expression__ic", qRegularExpressionIc); err != nil {
				return err
			}
		}
	}

	if o.RegularExpressionIe != nil {

		// query param regular_expression__ie
		var qrRegularExpressionIe string

		if o.RegularExpressionIe != nil {
			qrRegularExpressionIe = *o.RegularExpressionIe
		}
		qRegularExpressionIe := qrRegularExpressionIe
		if qRegularExpressionIe != "" {

			if err := r.SetQueryParam("regular_expression__ie", qRegularExpressionIe); err != nil {
				return err
			}
		}
	}

	if o.RegularExpressionIew != nil {

		// query param regular_expression__iew
		var qrRegularExpressionIew string

		if o.RegularExpressionIew != nil {
			qrRegularExpressionIew = *o.RegularExpressionIew
		}
		qRegularExpressionIew := qrRegularExpressionIew
		if qRegularExpressionIew != "" {

			if err := r.SetQueryParam("regular_expression__iew", qRegularExpressionIew); err != nil {
				return err
			}
		}
	}

	if o.RegularExpressionIsw != nil {

		// query param regular_expression__isw
		var qrRegularExpressionIsw string

		if o.RegularExpressionIsw != nil {
			qrRegularExpressionIsw = *o.RegularExpressionIsw
		}
		qRegularExpressionIsw := qrRegularExpressionIsw
		if qRegularExpressionIsw != "" {

			if err := r.SetQueryParam("regular_expression__isw", qRegularExpressionIsw); err != nil {
				return err
			}
		}
	}

	if o.RegularExpressionn != nil {

		// query param regular_expression__n
		var qrRegularExpressionn string

		if o.RegularExpressionn != nil {
			qrRegularExpressionn = *o.RegularExpressionn
		}
		qRegularExpressionn := qrRegularExpressionn
		if qRegularExpressionn != "" {

			if err := r.SetQueryParam("regular_expression__n", qRegularExpressionn); err != nil {
				return err
			}
		}
	}

	if o.RegularExpressionNic != nil {

		// query param regular_expression__nic
		var qrRegularExpressionNic string

		if o.RegularExpressionNic != nil {
			qrRegularExpressionNic = *o.RegularExpressionNic
		}
		qRegularExpressionNic := qrRegularExpressionNic
		if qRegularExpressionNic != "" {

			if err := r.SetQueryParam("regular_expression__nic", qRegularExpressionNic); err != nil {
				return err
			}
		}
	}

	if o.RegularExpressionNie != nil {

		// query param regular_expression__nie
		var qrRegularExpressionNie string

		if o.RegularExpressionNie != nil {
			qrRegularExpressionNie = *o.RegularExpressionNie
		}
		qRegularExpressionNie := qrRegularExpressionNie
		if qRegularExpressionNie != "" {

			if err := r.SetQueryParam("regular_expression__nie", qRegularExpressionNie); err != nil {
				return err
			}
		}
	}

	if o.RegularExpressionNiew != nil {

		// query param regular_expression__niew
		var qrRegularExpressionNiew string

		if o.RegularExpressionNiew != nil {
			qrRegularExpressionNiew = *o.RegularExpressionNiew
		}
		qRegularExpressionNiew := qrRegularExpressionNiew
		if qRegularExpressionNiew != "" {

			if err := r.SetQueryParam("regular_expression__niew", qRegularExpressionNiew); err != nil {
				return err
			}
		}
	}

	if o.RegularExpressionNisw != nil {

		// query param regular_expression__nisw
		var qrRegularExpressionNisw string

		if o.RegularExpressionNisw != nil {
			qrRegularExpressionNisw = *o.RegularExpressionNisw
		}
		qRegularExpressionNisw := qrRegularExpressionNisw
		if qRegularExpressionNisw != "" {

			if err := r.SetQueryParam("regular_expression__nisw", qRegularExpressionNisw); err != nil {
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

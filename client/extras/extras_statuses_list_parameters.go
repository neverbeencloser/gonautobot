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

// NewExtrasStatusesListParams creates a new ExtrasStatusesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasStatusesListParams() *ExtrasStatusesListParams {
	return &ExtrasStatusesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasStatusesListParamsWithTimeout creates a new ExtrasStatusesListParams object
// with the ability to set a timeout on a request.
func NewExtrasStatusesListParamsWithTimeout(timeout time.Duration) *ExtrasStatusesListParams {
	return &ExtrasStatusesListParams{
		timeout: timeout,
	}
}

// NewExtrasStatusesListParamsWithContext creates a new ExtrasStatusesListParams object
// with the ability to set a context for a request.
func NewExtrasStatusesListParamsWithContext(ctx context.Context) *ExtrasStatusesListParams {
	return &ExtrasStatusesListParams{
		Context: ctx,
	}
}

// NewExtrasStatusesListParamsWithHTTPClient creates a new ExtrasStatusesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasStatusesListParamsWithHTTPClient(client *http.Client) *ExtrasStatusesListParams {
	return &ExtrasStatusesListParams{
		HTTPClient: client,
	}
}

/* ExtrasStatusesListParams contains all the parameters to send to the API endpoint
   for the extras statuses list operation.

   Typically these are written to a http.Request.
*/
type ExtrasStatusesListParams struct {

	// Color.
	Color *string

	// ColorIc.
	ColorIc *string

	// ColorIe.
	ColorIe *string

	// ColorIew.
	ColorIew *string

	// ColorIsw.
	ColorIsw *string

	// Colorn.
	Colorn *string

	// ColorNic.
	ColorNic *string

	// ColorNie.
	ColorNie *string

	// ColorNiew.
	ColorNiew *string

	// ColorNisw.
	ColorNisw *string

	// ContentTypes.
	ContentTypes *string

	// ContentTypesn.
	ContentTypesn *string

	// Created.
	Created *string

	// CreatedGte.
	CreatedGte *string

	// CreatedLte.
	CreatedLte *string

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

// WithDefaults hydrates default values in the extras statuses list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasStatusesListParams) WithDefaults() *ExtrasStatusesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras statuses list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasStatusesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras statuses list params
func (o *ExtrasStatusesListParams) WithTimeout(timeout time.Duration) *ExtrasStatusesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras statuses list params
func (o *ExtrasStatusesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras statuses list params
func (o *ExtrasStatusesListParams) WithContext(ctx context.Context) *ExtrasStatusesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras statuses list params
func (o *ExtrasStatusesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras statuses list params
func (o *ExtrasStatusesListParams) WithHTTPClient(client *http.Client) *ExtrasStatusesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras statuses list params
func (o *ExtrasStatusesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithColor adds the color to the extras statuses list params
func (o *ExtrasStatusesListParams) WithColor(color *string) *ExtrasStatusesListParams {
	o.SetColor(color)
	return o
}

// SetColor adds the color to the extras statuses list params
func (o *ExtrasStatusesListParams) SetColor(color *string) {
	o.Color = color
}

// WithColorIc adds the colorIc to the extras statuses list params
func (o *ExtrasStatusesListParams) WithColorIc(colorIc *string) *ExtrasStatusesListParams {
	o.SetColorIc(colorIc)
	return o
}

// SetColorIc adds the colorIc to the extras statuses list params
func (o *ExtrasStatusesListParams) SetColorIc(colorIc *string) {
	o.ColorIc = colorIc
}

// WithColorIe adds the colorIe to the extras statuses list params
func (o *ExtrasStatusesListParams) WithColorIe(colorIe *string) *ExtrasStatusesListParams {
	o.SetColorIe(colorIe)
	return o
}

// SetColorIe adds the colorIe to the extras statuses list params
func (o *ExtrasStatusesListParams) SetColorIe(colorIe *string) {
	o.ColorIe = colorIe
}

// WithColorIew adds the colorIew to the extras statuses list params
func (o *ExtrasStatusesListParams) WithColorIew(colorIew *string) *ExtrasStatusesListParams {
	o.SetColorIew(colorIew)
	return o
}

// SetColorIew adds the colorIew to the extras statuses list params
func (o *ExtrasStatusesListParams) SetColorIew(colorIew *string) {
	o.ColorIew = colorIew
}

// WithColorIsw adds the colorIsw to the extras statuses list params
func (o *ExtrasStatusesListParams) WithColorIsw(colorIsw *string) *ExtrasStatusesListParams {
	o.SetColorIsw(colorIsw)
	return o
}

// SetColorIsw adds the colorIsw to the extras statuses list params
func (o *ExtrasStatusesListParams) SetColorIsw(colorIsw *string) {
	o.ColorIsw = colorIsw
}

// WithColorn adds the colorn to the extras statuses list params
func (o *ExtrasStatusesListParams) WithColorn(colorn *string) *ExtrasStatusesListParams {
	o.SetColorn(colorn)
	return o
}

// SetColorn adds the colorN to the extras statuses list params
func (o *ExtrasStatusesListParams) SetColorn(colorn *string) {
	o.Colorn = colorn
}

// WithColorNic adds the colorNic to the extras statuses list params
func (o *ExtrasStatusesListParams) WithColorNic(colorNic *string) *ExtrasStatusesListParams {
	o.SetColorNic(colorNic)
	return o
}

// SetColorNic adds the colorNic to the extras statuses list params
func (o *ExtrasStatusesListParams) SetColorNic(colorNic *string) {
	o.ColorNic = colorNic
}

// WithColorNie adds the colorNie to the extras statuses list params
func (o *ExtrasStatusesListParams) WithColorNie(colorNie *string) *ExtrasStatusesListParams {
	o.SetColorNie(colorNie)
	return o
}

// SetColorNie adds the colorNie to the extras statuses list params
func (o *ExtrasStatusesListParams) SetColorNie(colorNie *string) {
	o.ColorNie = colorNie
}

// WithColorNiew adds the colorNiew to the extras statuses list params
func (o *ExtrasStatusesListParams) WithColorNiew(colorNiew *string) *ExtrasStatusesListParams {
	o.SetColorNiew(colorNiew)
	return o
}

// SetColorNiew adds the colorNiew to the extras statuses list params
func (o *ExtrasStatusesListParams) SetColorNiew(colorNiew *string) {
	o.ColorNiew = colorNiew
}

// WithColorNisw adds the colorNisw to the extras statuses list params
func (o *ExtrasStatusesListParams) WithColorNisw(colorNisw *string) *ExtrasStatusesListParams {
	o.SetColorNisw(colorNisw)
	return o
}

// SetColorNisw adds the colorNisw to the extras statuses list params
func (o *ExtrasStatusesListParams) SetColorNisw(colorNisw *string) {
	o.ColorNisw = colorNisw
}

// WithContentTypes adds the contentTypes to the extras statuses list params
func (o *ExtrasStatusesListParams) WithContentTypes(contentTypes *string) *ExtrasStatusesListParams {
	o.SetContentTypes(contentTypes)
	return o
}

// SetContentTypes adds the contentTypes to the extras statuses list params
func (o *ExtrasStatusesListParams) SetContentTypes(contentTypes *string) {
	o.ContentTypes = contentTypes
}

// WithContentTypesn adds the contentTypesn to the extras statuses list params
func (o *ExtrasStatusesListParams) WithContentTypesn(contentTypesn *string) *ExtrasStatusesListParams {
	o.SetContentTypesn(contentTypesn)
	return o
}

// SetContentTypesn adds the contentTypesN to the extras statuses list params
func (o *ExtrasStatusesListParams) SetContentTypesn(contentTypesn *string) {
	o.ContentTypesn = contentTypesn
}

// WithCreated adds the created to the extras statuses list params
func (o *ExtrasStatusesListParams) WithCreated(created *string) *ExtrasStatusesListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the extras statuses list params
func (o *ExtrasStatusesListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGte adds the createdGte to the extras statuses list params
func (o *ExtrasStatusesListParams) WithCreatedGte(createdGte *string) *ExtrasStatusesListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the extras statuses list params
func (o *ExtrasStatusesListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLte adds the createdLte to the extras statuses list params
func (o *ExtrasStatusesListParams) WithCreatedLte(createdLte *string) *ExtrasStatusesListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the extras statuses list params
func (o *ExtrasStatusesListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithID adds the id to the extras statuses list params
func (o *ExtrasStatusesListParams) WithID(id *string) *ExtrasStatusesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras statuses list params
func (o *ExtrasStatusesListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the extras statuses list params
func (o *ExtrasStatusesListParams) WithIDIc(iDIc *string) *ExtrasStatusesListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the extras statuses list params
func (o *ExtrasStatusesListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the extras statuses list params
func (o *ExtrasStatusesListParams) WithIDIe(iDIe *string) *ExtrasStatusesListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the extras statuses list params
func (o *ExtrasStatusesListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the extras statuses list params
func (o *ExtrasStatusesListParams) WithIDIew(iDIew *string) *ExtrasStatusesListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the extras statuses list params
func (o *ExtrasStatusesListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the extras statuses list params
func (o *ExtrasStatusesListParams) WithIDIsw(iDIsw *string) *ExtrasStatusesListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the extras statuses list params
func (o *ExtrasStatusesListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the extras statuses list params
func (o *ExtrasStatusesListParams) WithIDn(iDn *string) *ExtrasStatusesListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the extras statuses list params
func (o *ExtrasStatusesListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the extras statuses list params
func (o *ExtrasStatusesListParams) WithIDNic(iDNic *string) *ExtrasStatusesListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the extras statuses list params
func (o *ExtrasStatusesListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the extras statuses list params
func (o *ExtrasStatusesListParams) WithIDNie(iDNie *string) *ExtrasStatusesListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the extras statuses list params
func (o *ExtrasStatusesListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the extras statuses list params
func (o *ExtrasStatusesListParams) WithIDNiew(iDNiew *string) *ExtrasStatusesListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the extras statuses list params
func (o *ExtrasStatusesListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the extras statuses list params
func (o *ExtrasStatusesListParams) WithIDNisw(iDNisw *string) *ExtrasStatusesListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the extras statuses list params
func (o *ExtrasStatusesListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithLastUpdated adds the lastUpdated to the extras statuses list params
func (o *ExtrasStatusesListParams) WithLastUpdated(lastUpdated *string) *ExtrasStatusesListParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the extras statuses list params
func (o *ExtrasStatusesListParams) SetLastUpdated(lastUpdated *string) {
	o.LastUpdated = lastUpdated
}

// WithLastUpdatedGte adds the lastUpdatedGte to the extras statuses list params
func (o *ExtrasStatusesListParams) WithLastUpdatedGte(lastUpdatedGte *string) *ExtrasStatusesListParams {
	o.SetLastUpdatedGte(lastUpdatedGte)
	return o
}

// SetLastUpdatedGte adds the lastUpdatedGte to the extras statuses list params
func (o *ExtrasStatusesListParams) SetLastUpdatedGte(lastUpdatedGte *string) {
	o.LastUpdatedGte = lastUpdatedGte
}

// WithLastUpdatedLte adds the lastUpdatedLte to the extras statuses list params
func (o *ExtrasStatusesListParams) WithLastUpdatedLte(lastUpdatedLte *string) *ExtrasStatusesListParams {
	o.SetLastUpdatedLte(lastUpdatedLte)
	return o
}

// SetLastUpdatedLte adds the lastUpdatedLte to the extras statuses list params
func (o *ExtrasStatusesListParams) SetLastUpdatedLte(lastUpdatedLte *string) {
	o.LastUpdatedLte = lastUpdatedLte
}

// WithLimit adds the limit to the extras statuses list params
func (o *ExtrasStatusesListParams) WithLimit(limit *int64) *ExtrasStatusesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the extras statuses list params
func (o *ExtrasStatusesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the extras statuses list params
func (o *ExtrasStatusesListParams) WithName(name *string) *ExtrasStatusesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the extras statuses list params
func (o *ExtrasStatusesListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the extras statuses list params
func (o *ExtrasStatusesListParams) WithNameIc(nameIc *string) *ExtrasStatusesListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the extras statuses list params
func (o *ExtrasStatusesListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the extras statuses list params
func (o *ExtrasStatusesListParams) WithNameIe(nameIe *string) *ExtrasStatusesListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the extras statuses list params
func (o *ExtrasStatusesListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the extras statuses list params
func (o *ExtrasStatusesListParams) WithNameIew(nameIew *string) *ExtrasStatusesListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the extras statuses list params
func (o *ExtrasStatusesListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the extras statuses list params
func (o *ExtrasStatusesListParams) WithNameIsw(nameIsw *string) *ExtrasStatusesListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the extras statuses list params
func (o *ExtrasStatusesListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the extras statuses list params
func (o *ExtrasStatusesListParams) WithNamen(namen *string) *ExtrasStatusesListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the extras statuses list params
func (o *ExtrasStatusesListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the extras statuses list params
func (o *ExtrasStatusesListParams) WithNameNic(nameNic *string) *ExtrasStatusesListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the extras statuses list params
func (o *ExtrasStatusesListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the extras statuses list params
func (o *ExtrasStatusesListParams) WithNameNie(nameNie *string) *ExtrasStatusesListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the extras statuses list params
func (o *ExtrasStatusesListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the extras statuses list params
func (o *ExtrasStatusesListParams) WithNameNiew(nameNiew *string) *ExtrasStatusesListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the extras statuses list params
func (o *ExtrasStatusesListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the extras statuses list params
func (o *ExtrasStatusesListParams) WithNameNisw(nameNisw *string) *ExtrasStatusesListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the extras statuses list params
func (o *ExtrasStatusesListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the extras statuses list params
func (o *ExtrasStatusesListParams) WithOffset(offset *int64) *ExtrasStatusesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the extras statuses list params
func (o *ExtrasStatusesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the extras statuses list params
func (o *ExtrasStatusesListParams) WithQ(q *string) *ExtrasStatusesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the extras statuses list params
func (o *ExtrasStatusesListParams) SetQ(q *string) {
	o.Q = q
}

// WithSlug adds the slug to the extras statuses list params
func (o *ExtrasStatusesListParams) WithSlug(slug *string) *ExtrasStatusesListParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the extras statuses list params
func (o *ExtrasStatusesListParams) SetSlug(slug *string) {
	o.Slug = slug
}

// WithSlugIc adds the slugIc to the extras statuses list params
func (o *ExtrasStatusesListParams) WithSlugIc(slugIc *string) *ExtrasStatusesListParams {
	o.SetSlugIc(slugIc)
	return o
}

// SetSlugIc adds the slugIc to the extras statuses list params
func (o *ExtrasStatusesListParams) SetSlugIc(slugIc *string) {
	o.SlugIc = slugIc
}

// WithSlugIe adds the slugIe to the extras statuses list params
func (o *ExtrasStatusesListParams) WithSlugIe(slugIe *string) *ExtrasStatusesListParams {
	o.SetSlugIe(slugIe)
	return o
}

// SetSlugIe adds the slugIe to the extras statuses list params
func (o *ExtrasStatusesListParams) SetSlugIe(slugIe *string) {
	o.SlugIe = slugIe
}

// WithSlugIew adds the slugIew to the extras statuses list params
func (o *ExtrasStatusesListParams) WithSlugIew(slugIew *string) *ExtrasStatusesListParams {
	o.SetSlugIew(slugIew)
	return o
}

// SetSlugIew adds the slugIew to the extras statuses list params
func (o *ExtrasStatusesListParams) SetSlugIew(slugIew *string) {
	o.SlugIew = slugIew
}

// WithSlugIsw adds the slugIsw to the extras statuses list params
func (o *ExtrasStatusesListParams) WithSlugIsw(slugIsw *string) *ExtrasStatusesListParams {
	o.SetSlugIsw(slugIsw)
	return o
}

// SetSlugIsw adds the slugIsw to the extras statuses list params
func (o *ExtrasStatusesListParams) SetSlugIsw(slugIsw *string) {
	o.SlugIsw = slugIsw
}

// WithSlugn adds the slugn to the extras statuses list params
func (o *ExtrasStatusesListParams) WithSlugn(slugn *string) *ExtrasStatusesListParams {
	o.SetSlugn(slugn)
	return o
}

// SetSlugn adds the slugN to the extras statuses list params
func (o *ExtrasStatusesListParams) SetSlugn(slugn *string) {
	o.Slugn = slugn
}

// WithSlugNic adds the slugNic to the extras statuses list params
func (o *ExtrasStatusesListParams) WithSlugNic(slugNic *string) *ExtrasStatusesListParams {
	o.SetSlugNic(slugNic)
	return o
}

// SetSlugNic adds the slugNic to the extras statuses list params
func (o *ExtrasStatusesListParams) SetSlugNic(slugNic *string) {
	o.SlugNic = slugNic
}

// WithSlugNie adds the slugNie to the extras statuses list params
func (o *ExtrasStatusesListParams) WithSlugNie(slugNie *string) *ExtrasStatusesListParams {
	o.SetSlugNie(slugNie)
	return o
}

// SetSlugNie adds the slugNie to the extras statuses list params
func (o *ExtrasStatusesListParams) SetSlugNie(slugNie *string) {
	o.SlugNie = slugNie
}

// WithSlugNiew adds the slugNiew to the extras statuses list params
func (o *ExtrasStatusesListParams) WithSlugNiew(slugNiew *string) *ExtrasStatusesListParams {
	o.SetSlugNiew(slugNiew)
	return o
}

// SetSlugNiew adds the slugNiew to the extras statuses list params
func (o *ExtrasStatusesListParams) SetSlugNiew(slugNiew *string) {
	o.SlugNiew = slugNiew
}

// WithSlugNisw adds the slugNisw to the extras statuses list params
func (o *ExtrasStatusesListParams) WithSlugNisw(slugNisw *string) *ExtrasStatusesListParams {
	o.SetSlugNisw(slugNisw)
	return o
}

// SetSlugNisw adds the slugNisw to the extras statuses list params
func (o *ExtrasStatusesListParams) SetSlugNisw(slugNisw *string) {
	o.SlugNisw = slugNisw
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasStatusesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Color != nil {

		// query param color
		var qrColor string

		if o.Color != nil {
			qrColor = *o.Color
		}
		qColor := qrColor
		if qColor != "" {

			if err := r.SetQueryParam("color", qColor); err != nil {
				return err
			}
		}
	}

	if o.ColorIc != nil {

		// query param color__ic
		var qrColorIc string

		if o.ColorIc != nil {
			qrColorIc = *o.ColorIc
		}
		qColorIc := qrColorIc
		if qColorIc != "" {

			if err := r.SetQueryParam("color__ic", qColorIc); err != nil {
				return err
			}
		}
	}

	if o.ColorIe != nil {

		// query param color__ie
		var qrColorIe string

		if o.ColorIe != nil {
			qrColorIe = *o.ColorIe
		}
		qColorIe := qrColorIe
		if qColorIe != "" {

			if err := r.SetQueryParam("color__ie", qColorIe); err != nil {
				return err
			}
		}
	}

	if o.ColorIew != nil {

		// query param color__iew
		var qrColorIew string

		if o.ColorIew != nil {
			qrColorIew = *o.ColorIew
		}
		qColorIew := qrColorIew
		if qColorIew != "" {

			if err := r.SetQueryParam("color__iew", qColorIew); err != nil {
				return err
			}
		}
	}

	if o.ColorIsw != nil {

		// query param color__isw
		var qrColorIsw string

		if o.ColorIsw != nil {
			qrColorIsw = *o.ColorIsw
		}
		qColorIsw := qrColorIsw
		if qColorIsw != "" {

			if err := r.SetQueryParam("color__isw", qColorIsw); err != nil {
				return err
			}
		}
	}

	if o.Colorn != nil {

		// query param color__n
		var qrColorn string

		if o.Colorn != nil {
			qrColorn = *o.Colorn
		}
		qColorn := qrColorn
		if qColorn != "" {

			if err := r.SetQueryParam("color__n", qColorn); err != nil {
				return err
			}
		}
	}

	if o.ColorNic != nil {

		// query param color__nic
		var qrColorNic string

		if o.ColorNic != nil {
			qrColorNic = *o.ColorNic
		}
		qColorNic := qrColorNic
		if qColorNic != "" {

			if err := r.SetQueryParam("color__nic", qColorNic); err != nil {
				return err
			}
		}
	}

	if o.ColorNie != nil {

		// query param color__nie
		var qrColorNie string

		if o.ColorNie != nil {
			qrColorNie = *o.ColorNie
		}
		qColorNie := qrColorNie
		if qColorNie != "" {

			if err := r.SetQueryParam("color__nie", qColorNie); err != nil {
				return err
			}
		}
	}

	if o.ColorNiew != nil {

		// query param color__niew
		var qrColorNiew string

		if o.ColorNiew != nil {
			qrColorNiew = *o.ColorNiew
		}
		qColorNiew := qrColorNiew
		if qColorNiew != "" {

			if err := r.SetQueryParam("color__niew", qColorNiew); err != nil {
				return err
			}
		}
	}

	if o.ColorNisw != nil {

		// query param color__nisw
		var qrColorNisw string

		if o.ColorNisw != nil {
			qrColorNisw = *o.ColorNisw
		}
		qColorNisw := qrColorNisw
		if qColorNisw != "" {

			if err := r.SetQueryParam("color__nisw", qColorNisw); err != nil {
				return err
			}
		}
	}

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

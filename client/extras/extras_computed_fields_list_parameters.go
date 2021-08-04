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

// NewExtrasComputedFieldsListParams creates a new ExtrasComputedFieldsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasComputedFieldsListParams() *ExtrasComputedFieldsListParams {
	return &ExtrasComputedFieldsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasComputedFieldsListParamsWithTimeout creates a new ExtrasComputedFieldsListParams object
// with the ability to set a timeout on a request.
func NewExtrasComputedFieldsListParamsWithTimeout(timeout time.Duration) *ExtrasComputedFieldsListParams {
	return &ExtrasComputedFieldsListParams{
		timeout: timeout,
	}
}

// NewExtrasComputedFieldsListParamsWithContext creates a new ExtrasComputedFieldsListParams object
// with the ability to set a context for a request.
func NewExtrasComputedFieldsListParamsWithContext(ctx context.Context) *ExtrasComputedFieldsListParams {
	return &ExtrasComputedFieldsListParams{
		Context: ctx,
	}
}

// NewExtrasComputedFieldsListParamsWithHTTPClient creates a new ExtrasComputedFieldsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasComputedFieldsListParamsWithHTTPClient(client *http.Client) *ExtrasComputedFieldsListParams {
	return &ExtrasComputedFieldsListParams{
		HTTPClient: client,
	}
}

/* ExtrasComputedFieldsListParams contains all the parameters to send to the API endpoint
   for the extras computed fields list operation.

   Typically these are written to a http.Request.
*/
type ExtrasComputedFieldsListParams struct {

	// ContentType.
	ContentType *string

	// ContentTypen.
	ContentTypen *string

	// FallbackValue.
	FallbackValue *string

	// FallbackValueIc.
	FallbackValueIc *string

	// FallbackValueIe.
	FallbackValueIe *string

	// FallbackValueIew.
	FallbackValueIew *string

	// FallbackValueIsw.
	FallbackValueIsw *string

	// FallbackValuen.
	FallbackValuen *string

	// FallbackValueNic.
	FallbackValueNic *string

	// FallbackValueNie.
	FallbackValueNie *string

	// FallbackValueNiew.
	FallbackValueNiew *string

	// FallbackValueNisw.
	FallbackValueNisw *string

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

	// Template.
	Template *string

	// TemplateIc.
	TemplateIc *string

	// TemplateIe.
	TemplateIe *string

	// TemplateIew.
	TemplateIew *string

	// TemplateIsw.
	TemplateIsw *string

	// Templaten.
	Templaten *string

	// TemplateNic.
	TemplateNic *string

	// TemplateNie.
	TemplateNie *string

	// TemplateNiew.
	TemplateNiew *string

	// TemplateNisw.
	TemplateNisw *string

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

// WithDefaults hydrates default values in the extras computed fields list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasComputedFieldsListParams) WithDefaults() *ExtrasComputedFieldsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras computed fields list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasComputedFieldsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) WithTimeout(timeout time.Duration) *ExtrasComputedFieldsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) WithContext(ctx context.Context) *ExtrasComputedFieldsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) WithHTTPClient(client *http.Client) *ExtrasComputedFieldsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentType adds the contentType to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) WithContentType(contentType *string) *ExtrasComputedFieldsListParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) SetContentType(contentType *string) {
	o.ContentType = contentType
}

// WithContentTypen adds the contentTypen to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) WithContentTypen(contentTypen *string) *ExtrasComputedFieldsListParams {
	o.SetContentTypen(contentTypen)
	return o
}

// SetContentTypen adds the contentTypeN to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) SetContentTypen(contentTypen *string) {
	o.ContentTypen = contentTypen
}

// WithFallbackValue adds the fallbackValue to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) WithFallbackValue(fallbackValue *string) *ExtrasComputedFieldsListParams {
	o.SetFallbackValue(fallbackValue)
	return o
}

// SetFallbackValue adds the fallbackValue to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) SetFallbackValue(fallbackValue *string) {
	o.FallbackValue = fallbackValue
}

// WithFallbackValueIc adds the fallbackValueIc to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) WithFallbackValueIc(fallbackValueIc *string) *ExtrasComputedFieldsListParams {
	o.SetFallbackValueIc(fallbackValueIc)
	return o
}

// SetFallbackValueIc adds the fallbackValueIc to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) SetFallbackValueIc(fallbackValueIc *string) {
	o.FallbackValueIc = fallbackValueIc
}

// WithFallbackValueIe adds the fallbackValueIe to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) WithFallbackValueIe(fallbackValueIe *string) *ExtrasComputedFieldsListParams {
	o.SetFallbackValueIe(fallbackValueIe)
	return o
}

// SetFallbackValueIe adds the fallbackValueIe to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) SetFallbackValueIe(fallbackValueIe *string) {
	o.FallbackValueIe = fallbackValueIe
}

// WithFallbackValueIew adds the fallbackValueIew to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) WithFallbackValueIew(fallbackValueIew *string) *ExtrasComputedFieldsListParams {
	o.SetFallbackValueIew(fallbackValueIew)
	return o
}

// SetFallbackValueIew adds the fallbackValueIew to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) SetFallbackValueIew(fallbackValueIew *string) {
	o.FallbackValueIew = fallbackValueIew
}

// WithFallbackValueIsw adds the fallbackValueIsw to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) WithFallbackValueIsw(fallbackValueIsw *string) *ExtrasComputedFieldsListParams {
	o.SetFallbackValueIsw(fallbackValueIsw)
	return o
}

// SetFallbackValueIsw adds the fallbackValueIsw to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) SetFallbackValueIsw(fallbackValueIsw *string) {
	o.FallbackValueIsw = fallbackValueIsw
}

// WithFallbackValuen adds the fallbackValuen to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) WithFallbackValuen(fallbackValuen *string) *ExtrasComputedFieldsListParams {
	o.SetFallbackValuen(fallbackValuen)
	return o
}

// SetFallbackValuen adds the fallbackValueN to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) SetFallbackValuen(fallbackValuen *string) {
	o.FallbackValuen = fallbackValuen
}

// WithFallbackValueNic adds the fallbackValueNic to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) WithFallbackValueNic(fallbackValueNic *string) *ExtrasComputedFieldsListParams {
	o.SetFallbackValueNic(fallbackValueNic)
	return o
}

// SetFallbackValueNic adds the fallbackValueNic to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) SetFallbackValueNic(fallbackValueNic *string) {
	o.FallbackValueNic = fallbackValueNic
}

// WithFallbackValueNie adds the fallbackValueNie to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) WithFallbackValueNie(fallbackValueNie *string) *ExtrasComputedFieldsListParams {
	o.SetFallbackValueNie(fallbackValueNie)
	return o
}

// SetFallbackValueNie adds the fallbackValueNie to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) SetFallbackValueNie(fallbackValueNie *string) {
	o.FallbackValueNie = fallbackValueNie
}

// WithFallbackValueNiew adds the fallbackValueNiew to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) WithFallbackValueNiew(fallbackValueNiew *string) *ExtrasComputedFieldsListParams {
	o.SetFallbackValueNiew(fallbackValueNiew)
	return o
}

// SetFallbackValueNiew adds the fallbackValueNiew to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) SetFallbackValueNiew(fallbackValueNiew *string) {
	o.FallbackValueNiew = fallbackValueNiew
}

// WithFallbackValueNisw adds the fallbackValueNisw to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) WithFallbackValueNisw(fallbackValueNisw *string) *ExtrasComputedFieldsListParams {
	o.SetFallbackValueNisw(fallbackValueNisw)
	return o
}

// SetFallbackValueNisw adds the fallbackValueNisw to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) SetFallbackValueNisw(fallbackValueNisw *string) {
	o.FallbackValueNisw = fallbackValueNisw
}

// WithLimit adds the limit to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) WithLimit(limit *int64) *ExtrasComputedFieldsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) WithOffset(offset *int64) *ExtrasComputedFieldsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) WithQ(q *string) *ExtrasComputedFieldsListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) SetQ(q *string) {
	o.Q = q
}

// WithSlug adds the slug to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) WithSlug(slug *string) *ExtrasComputedFieldsListParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) SetSlug(slug *string) {
	o.Slug = slug
}

// WithSlugIc adds the slugIc to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) WithSlugIc(slugIc *string) *ExtrasComputedFieldsListParams {
	o.SetSlugIc(slugIc)
	return o
}

// SetSlugIc adds the slugIc to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) SetSlugIc(slugIc *string) {
	o.SlugIc = slugIc
}

// WithSlugIe adds the slugIe to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) WithSlugIe(slugIe *string) *ExtrasComputedFieldsListParams {
	o.SetSlugIe(slugIe)
	return o
}

// SetSlugIe adds the slugIe to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) SetSlugIe(slugIe *string) {
	o.SlugIe = slugIe
}

// WithSlugIew adds the slugIew to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) WithSlugIew(slugIew *string) *ExtrasComputedFieldsListParams {
	o.SetSlugIew(slugIew)
	return o
}

// SetSlugIew adds the slugIew to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) SetSlugIew(slugIew *string) {
	o.SlugIew = slugIew
}

// WithSlugIsw adds the slugIsw to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) WithSlugIsw(slugIsw *string) *ExtrasComputedFieldsListParams {
	o.SetSlugIsw(slugIsw)
	return o
}

// SetSlugIsw adds the slugIsw to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) SetSlugIsw(slugIsw *string) {
	o.SlugIsw = slugIsw
}

// WithSlugn adds the slugn to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) WithSlugn(slugn *string) *ExtrasComputedFieldsListParams {
	o.SetSlugn(slugn)
	return o
}

// SetSlugn adds the slugN to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) SetSlugn(slugn *string) {
	o.Slugn = slugn
}

// WithSlugNic adds the slugNic to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) WithSlugNic(slugNic *string) *ExtrasComputedFieldsListParams {
	o.SetSlugNic(slugNic)
	return o
}

// SetSlugNic adds the slugNic to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) SetSlugNic(slugNic *string) {
	o.SlugNic = slugNic
}

// WithSlugNie adds the slugNie to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) WithSlugNie(slugNie *string) *ExtrasComputedFieldsListParams {
	o.SetSlugNie(slugNie)
	return o
}

// SetSlugNie adds the slugNie to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) SetSlugNie(slugNie *string) {
	o.SlugNie = slugNie
}

// WithSlugNiew adds the slugNiew to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) WithSlugNiew(slugNiew *string) *ExtrasComputedFieldsListParams {
	o.SetSlugNiew(slugNiew)
	return o
}

// SetSlugNiew adds the slugNiew to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) SetSlugNiew(slugNiew *string) {
	o.SlugNiew = slugNiew
}

// WithSlugNisw adds the slugNisw to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) WithSlugNisw(slugNisw *string) *ExtrasComputedFieldsListParams {
	o.SetSlugNisw(slugNisw)
	return o
}

// SetSlugNisw adds the slugNisw to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) SetSlugNisw(slugNisw *string) {
	o.SlugNisw = slugNisw
}

// WithTemplate adds the template to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) WithTemplate(template *string) *ExtrasComputedFieldsListParams {
	o.SetTemplate(template)
	return o
}

// SetTemplate adds the template to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) SetTemplate(template *string) {
	o.Template = template
}

// WithTemplateIc adds the templateIc to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) WithTemplateIc(templateIc *string) *ExtrasComputedFieldsListParams {
	o.SetTemplateIc(templateIc)
	return o
}

// SetTemplateIc adds the templateIc to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) SetTemplateIc(templateIc *string) {
	o.TemplateIc = templateIc
}

// WithTemplateIe adds the templateIe to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) WithTemplateIe(templateIe *string) *ExtrasComputedFieldsListParams {
	o.SetTemplateIe(templateIe)
	return o
}

// SetTemplateIe adds the templateIe to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) SetTemplateIe(templateIe *string) {
	o.TemplateIe = templateIe
}

// WithTemplateIew adds the templateIew to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) WithTemplateIew(templateIew *string) *ExtrasComputedFieldsListParams {
	o.SetTemplateIew(templateIew)
	return o
}

// SetTemplateIew adds the templateIew to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) SetTemplateIew(templateIew *string) {
	o.TemplateIew = templateIew
}

// WithTemplateIsw adds the templateIsw to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) WithTemplateIsw(templateIsw *string) *ExtrasComputedFieldsListParams {
	o.SetTemplateIsw(templateIsw)
	return o
}

// SetTemplateIsw adds the templateIsw to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) SetTemplateIsw(templateIsw *string) {
	o.TemplateIsw = templateIsw
}

// WithTemplaten adds the templaten to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) WithTemplaten(templaten *string) *ExtrasComputedFieldsListParams {
	o.SetTemplaten(templaten)
	return o
}

// SetTemplaten adds the templateN to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) SetTemplaten(templaten *string) {
	o.Templaten = templaten
}

// WithTemplateNic adds the templateNic to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) WithTemplateNic(templateNic *string) *ExtrasComputedFieldsListParams {
	o.SetTemplateNic(templateNic)
	return o
}

// SetTemplateNic adds the templateNic to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) SetTemplateNic(templateNic *string) {
	o.TemplateNic = templateNic
}

// WithTemplateNie adds the templateNie to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) WithTemplateNie(templateNie *string) *ExtrasComputedFieldsListParams {
	o.SetTemplateNie(templateNie)
	return o
}

// SetTemplateNie adds the templateNie to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) SetTemplateNie(templateNie *string) {
	o.TemplateNie = templateNie
}

// WithTemplateNiew adds the templateNiew to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) WithTemplateNiew(templateNiew *string) *ExtrasComputedFieldsListParams {
	o.SetTemplateNiew(templateNiew)
	return o
}

// SetTemplateNiew adds the templateNiew to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) SetTemplateNiew(templateNiew *string) {
	o.TemplateNiew = templateNiew
}

// WithTemplateNisw adds the templateNisw to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) WithTemplateNisw(templateNisw *string) *ExtrasComputedFieldsListParams {
	o.SetTemplateNisw(templateNisw)
	return o
}

// SetTemplateNisw adds the templateNisw to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) SetTemplateNisw(templateNisw *string) {
	o.TemplateNisw = templateNisw
}

// WithWeight adds the weight to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) WithWeight(weight *string) *ExtrasComputedFieldsListParams {
	o.SetWeight(weight)
	return o
}

// SetWeight adds the weight to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) SetWeight(weight *string) {
	o.Weight = weight
}

// WithWeightGt adds the weightGt to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) WithWeightGt(weightGt *string) *ExtrasComputedFieldsListParams {
	o.SetWeightGt(weightGt)
	return o
}

// SetWeightGt adds the weightGt to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) SetWeightGt(weightGt *string) {
	o.WeightGt = weightGt
}

// WithWeightGte adds the weightGte to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) WithWeightGte(weightGte *string) *ExtrasComputedFieldsListParams {
	o.SetWeightGte(weightGte)
	return o
}

// SetWeightGte adds the weightGte to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) SetWeightGte(weightGte *string) {
	o.WeightGte = weightGte
}

// WithWeightLt adds the weightLt to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) WithWeightLt(weightLt *string) *ExtrasComputedFieldsListParams {
	o.SetWeightLt(weightLt)
	return o
}

// SetWeightLt adds the weightLt to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) SetWeightLt(weightLt *string) {
	o.WeightLt = weightLt
}

// WithWeightLte adds the weightLte to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) WithWeightLte(weightLte *string) *ExtrasComputedFieldsListParams {
	o.SetWeightLte(weightLte)
	return o
}

// SetWeightLte adds the weightLte to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) SetWeightLte(weightLte *string) {
	o.WeightLte = weightLte
}

// WithWeightn adds the weightn to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) WithWeightn(weightn *string) *ExtrasComputedFieldsListParams {
	o.SetWeightn(weightn)
	return o
}

// SetWeightn adds the weightN to the extras computed fields list params
func (o *ExtrasComputedFieldsListParams) SetWeightn(weightn *string) {
	o.Weightn = weightn
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasComputedFieldsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.FallbackValue != nil {

		// query param fallback_value
		var qrFallbackValue string

		if o.FallbackValue != nil {
			qrFallbackValue = *o.FallbackValue
		}
		qFallbackValue := qrFallbackValue
		if qFallbackValue != "" {

			if err := r.SetQueryParam("fallback_value", qFallbackValue); err != nil {
				return err
			}
		}
	}

	if o.FallbackValueIc != nil {

		// query param fallback_value__ic
		var qrFallbackValueIc string

		if o.FallbackValueIc != nil {
			qrFallbackValueIc = *o.FallbackValueIc
		}
		qFallbackValueIc := qrFallbackValueIc
		if qFallbackValueIc != "" {

			if err := r.SetQueryParam("fallback_value__ic", qFallbackValueIc); err != nil {
				return err
			}
		}
	}

	if o.FallbackValueIe != nil {

		// query param fallback_value__ie
		var qrFallbackValueIe string

		if o.FallbackValueIe != nil {
			qrFallbackValueIe = *o.FallbackValueIe
		}
		qFallbackValueIe := qrFallbackValueIe
		if qFallbackValueIe != "" {

			if err := r.SetQueryParam("fallback_value__ie", qFallbackValueIe); err != nil {
				return err
			}
		}
	}

	if o.FallbackValueIew != nil {

		// query param fallback_value__iew
		var qrFallbackValueIew string

		if o.FallbackValueIew != nil {
			qrFallbackValueIew = *o.FallbackValueIew
		}
		qFallbackValueIew := qrFallbackValueIew
		if qFallbackValueIew != "" {

			if err := r.SetQueryParam("fallback_value__iew", qFallbackValueIew); err != nil {
				return err
			}
		}
	}

	if o.FallbackValueIsw != nil {

		// query param fallback_value__isw
		var qrFallbackValueIsw string

		if o.FallbackValueIsw != nil {
			qrFallbackValueIsw = *o.FallbackValueIsw
		}
		qFallbackValueIsw := qrFallbackValueIsw
		if qFallbackValueIsw != "" {

			if err := r.SetQueryParam("fallback_value__isw", qFallbackValueIsw); err != nil {
				return err
			}
		}
	}

	if o.FallbackValuen != nil {

		// query param fallback_value__n
		var qrFallbackValuen string

		if o.FallbackValuen != nil {
			qrFallbackValuen = *o.FallbackValuen
		}
		qFallbackValuen := qrFallbackValuen
		if qFallbackValuen != "" {

			if err := r.SetQueryParam("fallback_value__n", qFallbackValuen); err != nil {
				return err
			}
		}
	}

	if o.FallbackValueNic != nil {

		// query param fallback_value__nic
		var qrFallbackValueNic string

		if o.FallbackValueNic != nil {
			qrFallbackValueNic = *o.FallbackValueNic
		}
		qFallbackValueNic := qrFallbackValueNic
		if qFallbackValueNic != "" {

			if err := r.SetQueryParam("fallback_value__nic", qFallbackValueNic); err != nil {
				return err
			}
		}
	}

	if o.FallbackValueNie != nil {

		// query param fallback_value__nie
		var qrFallbackValueNie string

		if o.FallbackValueNie != nil {
			qrFallbackValueNie = *o.FallbackValueNie
		}
		qFallbackValueNie := qrFallbackValueNie
		if qFallbackValueNie != "" {

			if err := r.SetQueryParam("fallback_value__nie", qFallbackValueNie); err != nil {
				return err
			}
		}
	}

	if o.FallbackValueNiew != nil {

		// query param fallback_value__niew
		var qrFallbackValueNiew string

		if o.FallbackValueNiew != nil {
			qrFallbackValueNiew = *o.FallbackValueNiew
		}
		qFallbackValueNiew := qrFallbackValueNiew
		if qFallbackValueNiew != "" {

			if err := r.SetQueryParam("fallback_value__niew", qFallbackValueNiew); err != nil {
				return err
			}
		}
	}

	if o.FallbackValueNisw != nil {

		// query param fallback_value__nisw
		var qrFallbackValueNisw string

		if o.FallbackValueNisw != nil {
			qrFallbackValueNisw = *o.FallbackValueNisw
		}
		qFallbackValueNisw := qrFallbackValueNisw
		if qFallbackValueNisw != "" {

			if err := r.SetQueryParam("fallback_value__nisw", qFallbackValueNisw); err != nil {
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

	if o.Template != nil {

		// query param template
		var qrTemplate string

		if o.Template != nil {
			qrTemplate = *o.Template
		}
		qTemplate := qrTemplate
		if qTemplate != "" {

			if err := r.SetQueryParam("template", qTemplate); err != nil {
				return err
			}
		}
	}

	if o.TemplateIc != nil {

		// query param template__ic
		var qrTemplateIc string

		if o.TemplateIc != nil {
			qrTemplateIc = *o.TemplateIc
		}
		qTemplateIc := qrTemplateIc
		if qTemplateIc != "" {

			if err := r.SetQueryParam("template__ic", qTemplateIc); err != nil {
				return err
			}
		}
	}

	if o.TemplateIe != nil {

		// query param template__ie
		var qrTemplateIe string

		if o.TemplateIe != nil {
			qrTemplateIe = *o.TemplateIe
		}
		qTemplateIe := qrTemplateIe
		if qTemplateIe != "" {

			if err := r.SetQueryParam("template__ie", qTemplateIe); err != nil {
				return err
			}
		}
	}

	if o.TemplateIew != nil {

		// query param template__iew
		var qrTemplateIew string

		if o.TemplateIew != nil {
			qrTemplateIew = *o.TemplateIew
		}
		qTemplateIew := qrTemplateIew
		if qTemplateIew != "" {

			if err := r.SetQueryParam("template__iew", qTemplateIew); err != nil {
				return err
			}
		}
	}

	if o.TemplateIsw != nil {

		// query param template__isw
		var qrTemplateIsw string

		if o.TemplateIsw != nil {
			qrTemplateIsw = *o.TemplateIsw
		}
		qTemplateIsw := qrTemplateIsw
		if qTemplateIsw != "" {

			if err := r.SetQueryParam("template__isw", qTemplateIsw); err != nil {
				return err
			}
		}
	}

	if o.Templaten != nil {

		// query param template__n
		var qrTemplaten string

		if o.Templaten != nil {
			qrTemplaten = *o.Templaten
		}
		qTemplaten := qrTemplaten
		if qTemplaten != "" {

			if err := r.SetQueryParam("template__n", qTemplaten); err != nil {
				return err
			}
		}
	}

	if o.TemplateNic != nil {

		// query param template__nic
		var qrTemplateNic string

		if o.TemplateNic != nil {
			qrTemplateNic = *o.TemplateNic
		}
		qTemplateNic := qrTemplateNic
		if qTemplateNic != "" {

			if err := r.SetQueryParam("template__nic", qTemplateNic); err != nil {
				return err
			}
		}
	}

	if o.TemplateNie != nil {

		// query param template__nie
		var qrTemplateNie string

		if o.TemplateNie != nil {
			qrTemplateNie = *o.TemplateNie
		}
		qTemplateNie := qrTemplateNie
		if qTemplateNie != "" {

			if err := r.SetQueryParam("template__nie", qTemplateNie); err != nil {
				return err
			}
		}
	}

	if o.TemplateNiew != nil {

		// query param template__niew
		var qrTemplateNiew string

		if o.TemplateNiew != nil {
			qrTemplateNiew = *o.TemplateNiew
		}
		qTemplateNiew := qrTemplateNiew
		if qTemplateNiew != "" {

			if err := r.SetQueryParam("template__niew", qTemplateNiew); err != nil {
				return err
			}
		}
	}

	if o.TemplateNisw != nil {

		// query param template__nisw
		var qrTemplateNisw string

		if o.TemplateNisw != nil {
			qrTemplateNisw = *o.TemplateNisw
		}
		qTemplateNisw := qrTemplateNisw
		if qTemplateNisw != "" {

			if err := r.SetQueryParam("template__nisw", qTemplateNisw); err != nil {
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

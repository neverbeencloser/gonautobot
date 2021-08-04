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

// NewExtrasCustomLinksListParams creates a new ExtrasCustomLinksListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasCustomLinksListParams() *ExtrasCustomLinksListParams {
	return &ExtrasCustomLinksListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasCustomLinksListParamsWithTimeout creates a new ExtrasCustomLinksListParams object
// with the ability to set a timeout on a request.
func NewExtrasCustomLinksListParamsWithTimeout(timeout time.Duration) *ExtrasCustomLinksListParams {
	return &ExtrasCustomLinksListParams{
		timeout: timeout,
	}
}

// NewExtrasCustomLinksListParamsWithContext creates a new ExtrasCustomLinksListParams object
// with the ability to set a context for a request.
func NewExtrasCustomLinksListParamsWithContext(ctx context.Context) *ExtrasCustomLinksListParams {
	return &ExtrasCustomLinksListParams{
		Context: ctx,
	}
}

// NewExtrasCustomLinksListParamsWithHTTPClient creates a new ExtrasCustomLinksListParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasCustomLinksListParamsWithHTTPClient(client *http.Client) *ExtrasCustomLinksListParams {
	return &ExtrasCustomLinksListParams{
		HTTPClient: client,
	}
}

/* ExtrasCustomLinksListParams contains all the parameters to send to the API endpoint
   for the extras custom links list operation.

   Typically these are written to a http.Request.
*/
type ExtrasCustomLinksListParams struct {

	// ButtonClass.
	ButtonClass *string

	// ButtonClassn.
	ButtonClassn *string

	// ContentType.
	ContentType *string

	// ContentTypen.
	ContentTypen *string

	// GroupName.
	GroupName *string

	// GroupNameIc.
	GroupNameIc *string

	// GroupNameIe.
	GroupNameIe *string

	// GroupNameIew.
	GroupNameIew *string

	// GroupNameIsw.
	GroupNameIsw *string

	// GroupNamen.
	GroupNamen *string

	// GroupNameNic.
	GroupNameNic *string

	// GroupNameNie.
	GroupNameNie *string

	// GroupNameNiew.
	GroupNameNiew *string

	// GroupNameNisw.
	GroupNameNisw *string

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

	// NewWindow.
	NewWindow *string

	/* Offset.

	   The initial index from which to return the results.
	*/
	Offset *int64

	// Q.
	Q *string

	// TargetURL.
	TargetURL *string

	// TargetURLIc.
	TargetURLIc *string

	// TargetURLIe.
	TargetURLIe *string

	// TargetURLIew.
	TargetURLIew *string

	// TargetURLIsw.
	TargetURLIsw *string

	// TargetURLn.
	TargetURLn *string

	// TargetURLNic.
	TargetURLNic *string

	// TargetURLNie.
	TargetURLNie *string

	// TargetURLNiew.
	TargetURLNiew *string

	// TargetURLNisw.
	TargetURLNisw *string

	// Text.
	Text *string

	// TextIc.
	TextIc *string

	// TextIe.
	TextIe *string

	// TextIew.
	TextIew *string

	// TextIsw.
	TextIsw *string

	// Textn.
	Textn *string

	// TextNic.
	TextNic *string

	// TextNie.
	TextNie *string

	// TextNiew.
	TextNiew *string

	// TextNisw.
	TextNisw *string

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

// WithDefaults hydrates default values in the extras custom links list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasCustomLinksListParams) WithDefaults() *ExtrasCustomLinksListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras custom links list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasCustomLinksListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithTimeout(timeout time.Duration) *ExtrasCustomLinksListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithContext(ctx context.Context) *ExtrasCustomLinksListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithHTTPClient(client *http.Client) *ExtrasCustomLinksListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithButtonClass adds the buttonClass to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithButtonClass(buttonClass *string) *ExtrasCustomLinksListParams {
	o.SetButtonClass(buttonClass)
	return o
}

// SetButtonClass adds the buttonClass to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetButtonClass(buttonClass *string) {
	o.ButtonClass = buttonClass
}

// WithButtonClassn adds the buttonClassn to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithButtonClassn(buttonClassn *string) *ExtrasCustomLinksListParams {
	o.SetButtonClassn(buttonClassn)
	return o
}

// SetButtonClassn adds the buttonClassN to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetButtonClassn(buttonClassn *string) {
	o.ButtonClassn = buttonClassn
}

// WithContentType adds the contentType to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithContentType(contentType *string) *ExtrasCustomLinksListParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetContentType(contentType *string) {
	o.ContentType = contentType
}

// WithContentTypen adds the contentTypen to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithContentTypen(contentTypen *string) *ExtrasCustomLinksListParams {
	o.SetContentTypen(contentTypen)
	return o
}

// SetContentTypen adds the contentTypeN to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetContentTypen(contentTypen *string) {
	o.ContentTypen = contentTypen
}

// WithGroupName adds the groupName to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithGroupName(groupName *string) *ExtrasCustomLinksListParams {
	o.SetGroupName(groupName)
	return o
}

// SetGroupName adds the groupName to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetGroupName(groupName *string) {
	o.GroupName = groupName
}

// WithGroupNameIc adds the groupNameIc to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithGroupNameIc(groupNameIc *string) *ExtrasCustomLinksListParams {
	o.SetGroupNameIc(groupNameIc)
	return o
}

// SetGroupNameIc adds the groupNameIc to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetGroupNameIc(groupNameIc *string) {
	o.GroupNameIc = groupNameIc
}

// WithGroupNameIe adds the groupNameIe to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithGroupNameIe(groupNameIe *string) *ExtrasCustomLinksListParams {
	o.SetGroupNameIe(groupNameIe)
	return o
}

// SetGroupNameIe adds the groupNameIe to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetGroupNameIe(groupNameIe *string) {
	o.GroupNameIe = groupNameIe
}

// WithGroupNameIew adds the groupNameIew to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithGroupNameIew(groupNameIew *string) *ExtrasCustomLinksListParams {
	o.SetGroupNameIew(groupNameIew)
	return o
}

// SetGroupNameIew adds the groupNameIew to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetGroupNameIew(groupNameIew *string) {
	o.GroupNameIew = groupNameIew
}

// WithGroupNameIsw adds the groupNameIsw to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithGroupNameIsw(groupNameIsw *string) *ExtrasCustomLinksListParams {
	o.SetGroupNameIsw(groupNameIsw)
	return o
}

// SetGroupNameIsw adds the groupNameIsw to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetGroupNameIsw(groupNameIsw *string) {
	o.GroupNameIsw = groupNameIsw
}

// WithGroupNamen adds the groupNamen to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithGroupNamen(groupNamen *string) *ExtrasCustomLinksListParams {
	o.SetGroupNamen(groupNamen)
	return o
}

// SetGroupNamen adds the groupNameN to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetGroupNamen(groupNamen *string) {
	o.GroupNamen = groupNamen
}

// WithGroupNameNic adds the groupNameNic to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithGroupNameNic(groupNameNic *string) *ExtrasCustomLinksListParams {
	o.SetGroupNameNic(groupNameNic)
	return o
}

// SetGroupNameNic adds the groupNameNic to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetGroupNameNic(groupNameNic *string) {
	o.GroupNameNic = groupNameNic
}

// WithGroupNameNie adds the groupNameNie to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithGroupNameNie(groupNameNie *string) *ExtrasCustomLinksListParams {
	o.SetGroupNameNie(groupNameNie)
	return o
}

// SetGroupNameNie adds the groupNameNie to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetGroupNameNie(groupNameNie *string) {
	o.GroupNameNie = groupNameNie
}

// WithGroupNameNiew adds the groupNameNiew to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithGroupNameNiew(groupNameNiew *string) *ExtrasCustomLinksListParams {
	o.SetGroupNameNiew(groupNameNiew)
	return o
}

// SetGroupNameNiew adds the groupNameNiew to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetGroupNameNiew(groupNameNiew *string) {
	o.GroupNameNiew = groupNameNiew
}

// WithGroupNameNisw adds the groupNameNisw to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithGroupNameNisw(groupNameNisw *string) *ExtrasCustomLinksListParams {
	o.SetGroupNameNisw(groupNameNisw)
	return o
}

// SetGroupNameNisw adds the groupNameNisw to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetGroupNameNisw(groupNameNisw *string) {
	o.GroupNameNisw = groupNameNisw
}

// WithLimit adds the limit to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithLimit(limit *int64) *ExtrasCustomLinksListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithName(name *string) *ExtrasCustomLinksListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithNameIc(nameIc *string) *ExtrasCustomLinksListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithNameIe(nameIe *string) *ExtrasCustomLinksListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithNameIew(nameIew *string) *ExtrasCustomLinksListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithNameIsw(nameIsw *string) *ExtrasCustomLinksListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithNamen(namen *string) *ExtrasCustomLinksListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithNameNic(nameNic *string) *ExtrasCustomLinksListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithNameNie(nameNie *string) *ExtrasCustomLinksListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithNameNiew(nameNiew *string) *ExtrasCustomLinksListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithNameNisw(nameNisw *string) *ExtrasCustomLinksListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithNewWindow adds the newWindow to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithNewWindow(newWindow *string) *ExtrasCustomLinksListParams {
	o.SetNewWindow(newWindow)
	return o
}

// SetNewWindow adds the newWindow to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetNewWindow(newWindow *string) {
	o.NewWindow = newWindow
}

// WithOffset adds the offset to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithOffset(offset *int64) *ExtrasCustomLinksListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithQ(q *string) *ExtrasCustomLinksListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetQ(q *string) {
	o.Q = q
}

// WithTargetURL adds the targetURL to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithTargetURL(targetURL *string) *ExtrasCustomLinksListParams {
	o.SetTargetURL(targetURL)
	return o
}

// SetTargetURL adds the targetUrl to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetTargetURL(targetURL *string) {
	o.TargetURL = targetURL
}

// WithTargetURLIc adds the targetURLIc to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithTargetURLIc(targetURLIc *string) *ExtrasCustomLinksListParams {
	o.SetTargetURLIc(targetURLIc)
	return o
}

// SetTargetURLIc adds the targetUrlIc to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetTargetURLIc(targetURLIc *string) {
	o.TargetURLIc = targetURLIc
}

// WithTargetURLIe adds the targetURLIe to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithTargetURLIe(targetURLIe *string) *ExtrasCustomLinksListParams {
	o.SetTargetURLIe(targetURLIe)
	return o
}

// SetTargetURLIe adds the targetUrlIe to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetTargetURLIe(targetURLIe *string) {
	o.TargetURLIe = targetURLIe
}

// WithTargetURLIew adds the targetURLIew to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithTargetURLIew(targetURLIew *string) *ExtrasCustomLinksListParams {
	o.SetTargetURLIew(targetURLIew)
	return o
}

// SetTargetURLIew adds the targetUrlIew to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetTargetURLIew(targetURLIew *string) {
	o.TargetURLIew = targetURLIew
}

// WithTargetURLIsw adds the targetURLIsw to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithTargetURLIsw(targetURLIsw *string) *ExtrasCustomLinksListParams {
	o.SetTargetURLIsw(targetURLIsw)
	return o
}

// SetTargetURLIsw adds the targetUrlIsw to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetTargetURLIsw(targetURLIsw *string) {
	o.TargetURLIsw = targetURLIsw
}

// WithTargetURLn adds the targetURLn to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithTargetURLn(targetURLn *string) *ExtrasCustomLinksListParams {
	o.SetTargetURLn(targetURLn)
	return o
}

// SetTargetURLn adds the targetUrlN to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetTargetURLn(targetURLn *string) {
	o.TargetURLn = targetURLn
}

// WithTargetURLNic adds the targetURLNic to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithTargetURLNic(targetURLNic *string) *ExtrasCustomLinksListParams {
	o.SetTargetURLNic(targetURLNic)
	return o
}

// SetTargetURLNic adds the targetUrlNic to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetTargetURLNic(targetURLNic *string) {
	o.TargetURLNic = targetURLNic
}

// WithTargetURLNie adds the targetURLNie to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithTargetURLNie(targetURLNie *string) *ExtrasCustomLinksListParams {
	o.SetTargetURLNie(targetURLNie)
	return o
}

// SetTargetURLNie adds the targetUrlNie to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetTargetURLNie(targetURLNie *string) {
	o.TargetURLNie = targetURLNie
}

// WithTargetURLNiew adds the targetURLNiew to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithTargetURLNiew(targetURLNiew *string) *ExtrasCustomLinksListParams {
	o.SetTargetURLNiew(targetURLNiew)
	return o
}

// SetTargetURLNiew adds the targetUrlNiew to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetTargetURLNiew(targetURLNiew *string) {
	o.TargetURLNiew = targetURLNiew
}

// WithTargetURLNisw adds the targetURLNisw to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithTargetURLNisw(targetURLNisw *string) *ExtrasCustomLinksListParams {
	o.SetTargetURLNisw(targetURLNisw)
	return o
}

// SetTargetURLNisw adds the targetUrlNisw to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetTargetURLNisw(targetURLNisw *string) {
	o.TargetURLNisw = targetURLNisw
}

// WithText adds the text to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithText(text *string) *ExtrasCustomLinksListParams {
	o.SetText(text)
	return o
}

// SetText adds the text to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetText(text *string) {
	o.Text = text
}

// WithTextIc adds the textIc to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithTextIc(textIc *string) *ExtrasCustomLinksListParams {
	o.SetTextIc(textIc)
	return o
}

// SetTextIc adds the textIc to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetTextIc(textIc *string) {
	o.TextIc = textIc
}

// WithTextIe adds the textIe to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithTextIe(textIe *string) *ExtrasCustomLinksListParams {
	o.SetTextIe(textIe)
	return o
}

// SetTextIe adds the textIe to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetTextIe(textIe *string) {
	o.TextIe = textIe
}

// WithTextIew adds the textIew to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithTextIew(textIew *string) *ExtrasCustomLinksListParams {
	o.SetTextIew(textIew)
	return o
}

// SetTextIew adds the textIew to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetTextIew(textIew *string) {
	o.TextIew = textIew
}

// WithTextIsw adds the textIsw to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithTextIsw(textIsw *string) *ExtrasCustomLinksListParams {
	o.SetTextIsw(textIsw)
	return o
}

// SetTextIsw adds the textIsw to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetTextIsw(textIsw *string) {
	o.TextIsw = textIsw
}

// WithTextn adds the textn to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithTextn(textn *string) *ExtrasCustomLinksListParams {
	o.SetTextn(textn)
	return o
}

// SetTextn adds the textN to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetTextn(textn *string) {
	o.Textn = textn
}

// WithTextNic adds the textNic to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithTextNic(textNic *string) *ExtrasCustomLinksListParams {
	o.SetTextNic(textNic)
	return o
}

// SetTextNic adds the textNic to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetTextNic(textNic *string) {
	o.TextNic = textNic
}

// WithTextNie adds the textNie to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithTextNie(textNie *string) *ExtrasCustomLinksListParams {
	o.SetTextNie(textNie)
	return o
}

// SetTextNie adds the textNie to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetTextNie(textNie *string) {
	o.TextNie = textNie
}

// WithTextNiew adds the textNiew to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithTextNiew(textNiew *string) *ExtrasCustomLinksListParams {
	o.SetTextNiew(textNiew)
	return o
}

// SetTextNiew adds the textNiew to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetTextNiew(textNiew *string) {
	o.TextNiew = textNiew
}

// WithTextNisw adds the textNisw to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithTextNisw(textNisw *string) *ExtrasCustomLinksListParams {
	o.SetTextNisw(textNisw)
	return o
}

// SetTextNisw adds the textNisw to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetTextNisw(textNisw *string) {
	o.TextNisw = textNisw
}

// WithWeight adds the weight to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithWeight(weight *string) *ExtrasCustomLinksListParams {
	o.SetWeight(weight)
	return o
}

// SetWeight adds the weight to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetWeight(weight *string) {
	o.Weight = weight
}

// WithWeightGt adds the weightGt to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithWeightGt(weightGt *string) *ExtrasCustomLinksListParams {
	o.SetWeightGt(weightGt)
	return o
}

// SetWeightGt adds the weightGt to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetWeightGt(weightGt *string) {
	o.WeightGt = weightGt
}

// WithWeightGte adds the weightGte to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithWeightGte(weightGte *string) *ExtrasCustomLinksListParams {
	o.SetWeightGte(weightGte)
	return o
}

// SetWeightGte adds the weightGte to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetWeightGte(weightGte *string) {
	o.WeightGte = weightGte
}

// WithWeightLt adds the weightLt to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithWeightLt(weightLt *string) *ExtrasCustomLinksListParams {
	o.SetWeightLt(weightLt)
	return o
}

// SetWeightLt adds the weightLt to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetWeightLt(weightLt *string) {
	o.WeightLt = weightLt
}

// WithWeightLte adds the weightLte to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithWeightLte(weightLte *string) *ExtrasCustomLinksListParams {
	o.SetWeightLte(weightLte)
	return o
}

// SetWeightLte adds the weightLte to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetWeightLte(weightLte *string) {
	o.WeightLte = weightLte
}

// WithWeightn adds the weightn to the extras custom links list params
func (o *ExtrasCustomLinksListParams) WithWeightn(weightn *string) *ExtrasCustomLinksListParams {
	o.SetWeightn(weightn)
	return o
}

// SetWeightn adds the weightN to the extras custom links list params
func (o *ExtrasCustomLinksListParams) SetWeightn(weightn *string) {
	o.Weightn = weightn
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasCustomLinksListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ButtonClass != nil {

		// query param button_class
		var qrButtonClass string

		if o.ButtonClass != nil {
			qrButtonClass = *o.ButtonClass
		}
		qButtonClass := qrButtonClass
		if qButtonClass != "" {

			if err := r.SetQueryParam("button_class", qButtonClass); err != nil {
				return err
			}
		}
	}

	if o.ButtonClassn != nil {

		// query param button_class__n
		var qrButtonClassn string

		if o.ButtonClassn != nil {
			qrButtonClassn = *o.ButtonClassn
		}
		qButtonClassn := qrButtonClassn
		if qButtonClassn != "" {

			if err := r.SetQueryParam("button_class__n", qButtonClassn); err != nil {
				return err
			}
		}
	}

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

	if o.GroupName != nil {

		// query param group_name
		var qrGroupName string

		if o.GroupName != nil {
			qrGroupName = *o.GroupName
		}
		qGroupName := qrGroupName
		if qGroupName != "" {

			if err := r.SetQueryParam("group_name", qGroupName); err != nil {
				return err
			}
		}
	}

	if o.GroupNameIc != nil {

		// query param group_name__ic
		var qrGroupNameIc string

		if o.GroupNameIc != nil {
			qrGroupNameIc = *o.GroupNameIc
		}
		qGroupNameIc := qrGroupNameIc
		if qGroupNameIc != "" {

			if err := r.SetQueryParam("group_name__ic", qGroupNameIc); err != nil {
				return err
			}
		}
	}

	if o.GroupNameIe != nil {

		// query param group_name__ie
		var qrGroupNameIe string

		if o.GroupNameIe != nil {
			qrGroupNameIe = *o.GroupNameIe
		}
		qGroupNameIe := qrGroupNameIe
		if qGroupNameIe != "" {

			if err := r.SetQueryParam("group_name__ie", qGroupNameIe); err != nil {
				return err
			}
		}
	}

	if o.GroupNameIew != nil {

		// query param group_name__iew
		var qrGroupNameIew string

		if o.GroupNameIew != nil {
			qrGroupNameIew = *o.GroupNameIew
		}
		qGroupNameIew := qrGroupNameIew
		if qGroupNameIew != "" {

			if err := r.SetQueryParam("group_name__iew", qGroupNameIew); err != nil {
				return err
			}
		}
	}

	if o.GroupNameIsw != nil {

		// query param group_name__isw
		var qrGroupNameIsw string

		if o.GroupNameIsw != nil {
			qrGroupNameIsw = *o.GroupNameIsw
		}
		qGroupNameIsw := qrGroupNameIsw
		if qGroupNameIsw != "" {

			if err := r.SetQueryParam("group_name__isw", qGroupNameIsw); err != nil {
				return err
			}
		}
	}

	if o.GroupNamen != nil {

		// query param group_name__n
		var qrGroupNamen string

		if o.GroupNamen != nil {
			qrGroupNamen = *o.GroupNamen
		}
		qGroupNamen := qrGroupNamen
		if qGroupNamen != "" {

			if err := r.SetQueryParam("group_name__n", qGroupNamen); err != nil {
				return err
			}
		}
	}

	if o.GroupNameNic != nil {

		// query param group_name__nic
		var qrGroupNameNic string

		if o.GroupNameNic != nil {
			qrGroupNameNic = *o.GroupNameNic
		}
		qGroupNameNic := qrGroupNameNic
		if qGroupNameNic != "" {

			if err := r.SetQueryParam("group_name__nic", qGroupNameNic); err != nil {
				return err
			}
		}
	}

	if o.GroupNameNie != nil {

		// query param group_name__nie
		var qrGroupNameNie string

		if o.GroupNameNie != nil {
			qrGroupNameNie = *o.GroupNameNie
		}
		qGroupNameNie := qrGroupNameNie
		if qGroupNameNie != "" {

			if err := r.SetQueryParam("group_name__nie", qGroupNameNie); err != nil {
				return err
			}
		}
	}

	if o.GroupNameNiew != nil {

		// query param group_name__niew
		var qrGroupNameNiew string

		if o.GroupNameNiew != nil {
			qrGroupNameNiew = *o.GroupNameNiew
		}
		qGroupNameNiew := qrGroupNameNiew
		if qGroupNameNiew != "" {

			if err := r.SetQueryParam("group_name__niew", qGroupNameNiew); err != nil {
				return err
			}
		}
	}

	if o.GroupNameNisw != nil {

		// query param group_name__nisw
		var qrGroupNameNisw string

		if o.GroupNameNisw != nil {
			qrGroupNameNisw = *o.GroupNameNisw
		}
		qGroupNameNisw := qrGroupNameNisw
		if qGroupNameNisw != "" {

			if err := r.SetQueryParam("group_name__nisw", qGroupNameNisw); err != nil {
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

	if o.NewWindow != nil {

		// query param new_window
		var qrNewWindow string

		if o.NewWindow != nil {
			qrNewWindow = *o.NewWindow
		}
		qNewWindow := qrNewWindow
		if qNewWindow != "" {

			if err := r.SetQueryParam("new_window", qNewWindow); err != nil {
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

	if o.TargetURL != nil {

		// query param target_url
		var qrTargetURL string

		if o.TargetURL != nil {
			qrTargetURL = *o.TargetURL
		}
		qTargetURL := qrTargetURL
		if qTargetURL != "" {

			if err := r.SetQueryParam("target_url", qTargetURL); err != nil {
				return err
			}
		}
	}

	if o.TargetURLIc != nil {

		// query param target_url__ic
		var qrTargetURLIc string

		if o.TargetURLIc != nil {
			qrTargetURLIc = *o.TargetURLIc
		}
		qTargetURLIc := qrTargetURLIc
		if qTargetURLIc != "" {

			if err := r.SetQueryParam("target_url__ic", qTargetURLIc); err != nil {
				return err
			}
		}
	}

	if o.TargetURLIe != nil {

		// query param target_url__ie
		var qrTargetURLIe string

		if o.TargetURLIe != nil {
			qrTargetURLIe = *o.TargetURLIe
		}
		qTargetURLIe := qrTargetURLIe
		if qTargetURLIe != "" {

			if err := r.SetQueryParam("target_url__ie", qTargetURLIe); err != nil {
				return err
			}
		}
	}

	if o.TargetURLIew != nil {

		// query param target_url__iew
		var qrTargetURLIew string

		if o.TargetURLIew != nil {
			qrTargetURLIew = *o.TargetURLIew
		}
		qTargetURLIew := qrTargetURLIew
		if qTargetURLIew != "" {

			if err := r.SetQueryParam("target_url__iew", qTargetURLIew); err != nil {
				return err
			}
		}
	}

	if o.TargetURLIsw != nil {

		// query param target_url__isw
		var qrTargetURLIsw string

		if o.TargetURLIsw != nil {
			qrTargetURLIsw = *o.TargetURLIsw
		}
		qTargetURLIsw := qrTargetURLIsw
		if qTargetURLIsw != "" {

			if err := r.SetQueryParam("target_url__isw", qTargetURLIsw); err != nil {
				return err
			}
		}
	}

	if o.TargetURLn != nil {

		// query param target_url__n
		var qrTargetURLn string

		if o.TargetURLn != nil {
			qrTargetURLn = *o.TargetURLn
		}
		qTargetURLn := qrTargetURLn
		if qTargetURLn != "" {

			if err := r.SetQueryParam("target_url__n", qTargetURLn); err != nil {
				return err
			}
		}
	}

	if o.TargetURLNic != nil {

		// query param target_url__nic
		var qrTargetURLNic string

		if o.TargetURLNic != nil {
			qrTargetURLNic = *o.TargetURLNic
		}
		qTargetURLNic := qrTargetURLNic
		if qTargetURLNic != "" {

			if err := r.SetQueryParam("target_url__nic", qTargetURLNic); err != nil {
				return err
			}
		}
	}

	if o.TargetURLNie != nil {

		// query param target_url__nie
		var qrTargetURLNie string

		if o.TargetURLNie != nil {
			qrTargetURLNie = *o.TargetURLNie
		}
		qTargetURLNie := qrTargetURLNie
		if qTargetURLNie != "" {

			if err := r.SetQueryParam("target_url__nie", qTargetURLNie); err != nil {
				return err
			}
		}
	}

	if o.TargetURLNiew != nil {

		// query param target_url__niew
		var qrTargetURLNiew string

		if o.TargetURLNiew != nil {
			qrTargetURLNiew = *o.TargetURLNiew
		}
		qTargetURLNiew := qrTargetURLNiew
		if qTargetURLNiew != "" {

			if err := r.SetQueryParam("target_url__niew", qTargetURLNiew); err != nil {
				return err
			}
		}
	}

	if o.TargetURLNisw != nil {

		// query param target_url__nisw
		var qrTargetURLNisw string

		if o.TargetURLNisw != nil {
			qrTargetURLNisw = *o.TargetURLNisw
		}
		qTargetURLNisw := qrTargetURLNisw
		if qTargetURLNisw != "" {

			if err := r.SetQueryParam("target_url__nisw", qTargetURLNisw); err != nil {
				return err
			}
		}
	}

	if o.Text != nil {

		// query param text
		var qrText string

		if o.Text != nil {
			qrText = *o.Text
		}
		qText := qrText
		if qText != "" {

			if err := r.SetQueryParam("text", qText); err != nil {
				return err
			}
		}
	}

	if o.TextIc != nil {

		// query param text__ic
		var qrTextIc string

		if o.TextIc != nil {
			qrTextIc = *o.TextIc
		}
		qTextIc := qrTextIc
		if qTextIc != "" {

			if err := r.SetQueryParam("text__ic", qTextIc); err != nil {
				return err
			}
		}
	}

	if o.TextIe != nil {

		// query param text__ie
		var qrTextIe string

		if o.TextIe != nil {
			qrTextIe = *o.TextIe
		}
		qTextIe := qrTextIe
		if qTextIe != "" {

			if err := r.SetQueryParam("text__ie", qTextIe); err != nil {
				return err
			}
		}
	}

	if o.TextIew != nil {

		// query param text__iew
		var qrTextIew string

		if o.TextIew != nil {
			qrTextIew = *o.TextIew
		}
		qTextIew := qrTextIew
		if qTextIew != "" {

			if err := r.SetQueryParam("text__iew", qTextIew); err != nil {
				return err
			}
		}
	}

	if o.TextIsw != nil {

		// query param text__isw
		var qrTextIsw string

		if o.TextIsw != nil {
			qrTextIsw = *o.TextIsw
		}
		qTextIsw := qrTextIsw
		if qTextIsw != "" {

			if err := r.SetQueryParam("text__isw", qTextIsw); err != nil {
				return err
			}
		}
	}

	if o.Textn != nil {

		// query param text__n
		var qrTextn string

		if o.Textn != nil {
			qrTextn = *o.Textn
		}
		qTextn := qrTextn
		if qTextn != "" {

			if err := r.SetQueryParam("text__n", qTextn); err != nil {
				return err
			}
		}
	}

	if o.TextNic != nil {

		// query param text__nic
		var qrTextNic string

		if o.TextNic != nil {
			qrTextNic = *o.TextNic
		}
		qTextNic := qrTextNic
		if qTextNic != "" {

			if err := r.SetQueryParam("text__nic", qTextNic); err != nil {
				return err
			}
		}
	}

	if o.TextNie != nil {

		// query param text__nie
		var qrTextNie string

		if o.TextNie != nil {
			qrTextNie = *o.TextNie
		}
		qTextNie := qrTextNie
		if qTextNie != "" {

			if err := r.SetQueryParam("text__nie", qTextNie); err != nil {
				return err
			}
		}
	}

	if o.TextNiew != nil {

		// query param text__niew
		var qrTextNiew string

		if o.TextNiew != nil {
			qrTextNiew = *o.TextNiew
		}
		qTextNiew := qrTextNiew
		if qTextNiew != "" {

			if err := r.SetQueryParam("text__niew", qTextNiew); err != nil {
				return err
			}
		}
	}

	if o.TextNisw != nil {

		// query param text__nisw
		var qrTextNisw string

		if o.TextNisw != nil {
			qrTextNisw = *o.TextNisw
		}
		qTextNisw := qrTextNisw
		if qTextNisw != "" {

			if err := r.SetQueryParam("text__nisw", qTextNisw); err != nil {
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

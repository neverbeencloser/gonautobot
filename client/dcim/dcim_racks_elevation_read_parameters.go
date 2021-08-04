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

// NewDcimRacksElevationReadParams creates a new DcimRacksElevationReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRacksElevationReadParams() *DcimRacksElevationReadParams {
	return &DcimRacksElevationReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRacksElevationReadParamsWithTimeout creates a new DcimRacksElevationReadParams object
// with the ability to set a timeout on a request.
func NewDcimRacksElevationReadParamsWithTimeout(timeout time.Duration) *DcimRacksElevationReadParams {
	return &DcimRacksElevationReadParams{
		timeout: timeout,
	}
}

// NewDcimRacksElevationReadParamsWithContext creates a new DcimRacksElevationReadParams object
// with the ability to set a context for a request.
func NewDcimRacksElevationReadParamsWithContext(ctx context.Context) *DcimRacksElevationReadParams {
	return &DcimRacksElevationReadParams{
		Context: ctx,
	}
}

// NewDcimRacksElevationReadParamsWithHTTPClient creates a new DcimRacksElevationReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRacksElevationReadParamsWithHTTPClient(client *http.Client) *DcimRacksElevationReadParams {
	return &DcimRacksElevationReadParams{
		HTTPClient: client,
	}
}

/* DcimRacksElevationReadParams contains all the parameters to send to the API endpoint
   for the dcim racks elevation read operation.

   Typically these are written to a http.Request.
*/
type DcimRacksElevationReadParams struct {

	// Exclude.
	//
	// Format: uuid
	Exclude *strfmt.UUID

	// ExpandDevices.
	//
	// Default: true
	ExpandDevices *bool

	// Face.
	//
	// Default: "front"
	Face *string

	/* ID.

	   A UUID string identifying this rack.

	   Format: uuid
	*/
	ID strfmt.UUID

	// IncludeImages.
	//
	// Default: true
	IncludeImages *bool

	// LegendWidth.
	//
	// Default: 30
	LegendWidth *int64

	// Q.
	Q *string

	// Render.
	//
	// Default: "json"
	Render *string

	// UnitHeight.
	//
	// Default: 22
	UnitHeight *int64

	// UnitWidth.
	//
	// Default: 220
	UnitWidth *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim racks elevation read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRacksElevationReadParams) WithDefaults() *DcimRacksElevationReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim racks elevation read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRacksElevationReadParams) SetDefaults() {
	var (
		expandDevicesDefault = bool(true)

		faceDefault = string("front")

		includeImagesDefault = bool(true)

		legendWidthDefault = int64(30)

		renderDefault = string("json")

		unitHeightDefault = int64(22)

		unitWidthDefault = int64(220)
	)

	val := DcimRacksElevationReadParams{
		ExpandDevices: &expandDevicesDefault,
		Face:          &faceDefault,
		IncludeImages: &includeImagesDefault,
		LegendWidth:   &legendWidthDefault,
		Render:        &renderDefault,
		UnitHeight:    &unitHeightDefault,
		UnitWidth:     &unitWidthDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the dcim racks elevation read params
func (o *DcimRacksElevationReadParams) WithTimeout(timeout time.Duration) *DcimRacksElevationReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim racks elevation read params
func (o *DcimRacksElevationReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim racks elevation read params
func (o *DcimRacksElevationReadParams) WithContext(ctx context.Context) *DcimRacksElevationReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim racks elevation read params
func (o *DcimRacksElevationReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim racks elevation read params
func (o *DcimRacksElevationReadParams) WithHTTPClient(client *http.Client) *DcimRacksElevationReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim racks elevation read params
func (o *DcimRacksElevationReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExclude adds the exclude to the dcim racks elevation read params
func (o *DcimRacksElevationReadParams) WithExclude(exclude *strfmt.UUID) *DcimRacksElevationReadParams {
	o.SetExclude(exclude)
	return o
}

// SetExclude adds the exclude to the dcim racks elevation read params
func (o *DcimRacksElevationReadParams) SetExclude(exclude *strfmt.UUID) {
	o.Exclude = exclude
}

// WithExpandDevices adds the expandDevices to the dcim racks elevation read params
func (o *DcimRacksElevationReadParams) WithExpandDevices(expandDevices *bool) *DcimRacksElevationReadParams {
	o.SetExpandDevices(expandDevices)
	return o
}

// SetExpandDevices adds the expandDevices to the dcim racks elevation read params
func (o *DcimRacksElevationReadParams) SetExpandDevices(expandDevices *bool) {
	o.ExpandDevices = expandDevices
}

// WithFace adds the face to the dcim racks elevation read params
func (o *DcimRacksElevationReadParams) WithFace(face *string) *DcimRacksElevationReadParams {
	o.SetFace(face)
	return o
}

// SetFace adds the face to the dcim racks elevation read params
func (o *DcimRacksElevationReadParams) SetFace(face *string) {
	o.Face = face
}

// WithID adds the id to the dcim racks elevation read params
func (o *DcimRacksElevationReadParams) WithID(id strfmt.UUID) *DcimRacksElevationReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim racks elevation read params
func (o *DcimRacksElevationReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithIncludeImages adds the includeImages to the dcim racks elevation read params
func (o *DcimRacksElevationReadParams) WithIncludeImages(includeImages *bool) *DcimRacksElevationReadParams {
	o.SetIncludeImages(includeImages)
	return o
}

// SetIncludeImages adds the includeImages to the dcim racks elevation read params
func (o *DcimRacksElevationReadParams) SetIncludeImages(includeImages *bool) {
	o.IncludeImages = includeImages
}

// WithLegendWidth adds the legendWidth to the dcim racks elevation read params
func (o *DcimRacksElevationReadParams) WithLegendWidth(legendWidth *int64) *DcimRacksElevationReadParams {
	o.SetLegendWidth(legendWidth)
	return o
}

// SetLegendWidth adds the legendWidth to the dcim racks elevation read params
func (o *DcimRacksElevationReadParams) SetLegendWidth(legendWidth *int64) {
	o.LegendWidth = legendWidth
}

// WithQ adds the q to the dcim racks elevation read params
func (o *DcimRacksElevationReadParams) WithQ(q *string) *DcimRacksElevationReadParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim racks elevation read params
func (o *DcimRacksElevationReadParams) SetQ(q *string) {
	o.Q = q
}

// WithRender adds the render to the dcim racks elevation read params
func (o *DcimRacksElevationReadParams) WithRender(render *string) *DcimRacksElevationReadParams {
	o.SetRender(render)
	return o
}

// SetRender adds the render to the dcim racks elevation read params
func (o *DcimRacksElevationReadParams) SetRender(render *string) {
	o.Render = render
}

// WithUnitHeight adds the unitHeight to the dcim racks elevation read params
func (o *DcimRacksElevationReadParams) WithUnitHeight(unitHeight *int64) *DcimRacksElevationReadParams {
	o.SetUnitHeight(unitHeight)
	return o
}

// SetUnitHeight adds the unitHeight to the dcim racks elevation read params
func (o *DcimRacksElevationReadParams) SetUnitHeight(unitHeight *int64) {
	o.UnitHeight = unitHeight
}

// WithUnitWidth adds the unitWidth to the dcim racks elevation read params
func (o *DcimRacksElevationReadParams) WithUnitWidth(unitWidth *int64) *DcimRacksElevationReadParams {
	o.SetUnitWidth(unitWidth)
	return o
}

// SetUnitWidth adds the unitWidth to the dcim racks elevation read params
func (o *DcimRacksElevationReadParams) SetUnitWidth(unitWidth *int64) {
	o.UnitWidth = unitWidth
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRacksElevationReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Exclude != nil {

		// query param exclude
		var qrExclude strfmt.UUID

		if o.Exclude != nil {
			qrExclude = *o.Exclude
		}
		qExclude := qrExclude.String()
		if qExclude != "" {

			if err := r.SetQueryParam("exclude", qExclude); err != nil {
				return err
			}
		}
	}

	if o.ExpandDevices != nil {

		// query param expand_devices
		var qrExpandDevices bool

		if o.ExpandDevices != nil {
			qrExpandDevices = *o.ExpandDevices
		}
		qExpandDevices := swag.FormatBool(qrExpandDevices)
		if qExpandDevices != "" {

			if err := r.SetQueryParam("expand_devices", qExpandDevices); err != nil {
				return err
			}
		}
	}

	if o.Face != nil {

		// query param face
		var qrFace string

		if o.Face != nil {
			qrFace = *o.Face
		}
		qFace := qrFace
		if qFace != "" {

			if err := r.SetQueryParam("face", qFace); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if o.IncludeImages != nil {

		// query param include_images
		var qrIncludeImages bool

		if o.IncludeImages != nil {
			qrIncludeImages = *o.IncludeImages
		}
		qIncludeImages := swag.FormatBool(qrIncludeImages)
		if qIncludeImages != "" {

			if err := r.SetQueryParam("include_images", qIncludeImages); err != nil {
				return err
			}
		}
	}

	if o.LegendWidth != nil {

		// query param legend_width
		var qrLegendWidth int64

		if o.LegendWidth != nil {
			qrLegendWidth = *o.LegendWidth
		}
		qLegendWidth := swag.FormatInt64(qrLegendWidth)
		if qLegendWidth != "" {

			if err := r.SetQueryParam("legend_width", qLegendWidth); err != nil {
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

	if o.Render != nil {

		// query param render
		var qrRender string

		if o.Render != nil {
			qrRender = *o.Render
		}
		qRender := qrRender
		if qRender != "" {

			if err := r.SetQueryParam("render", qRender); err != nil {
				return err
			}
		}
	}

	if o.UnitHeight != nil {

		// query param unit_height
		var qrUnitHeight int64

		if o.UnitHeight != nil {
			qrUnitHeight = *o.UnitHeight
		}
		qUnitHeight := swag.FormatInt64(qrUnitHeight)
		if qUnitHeight != "" {

			if err := r.SetQueryParam("unit_height", qUnitHeight); err != nil {
				return err
			}
		}
	}

	if o.UnitWidth != nil {

		// query param unit_width
		var qrUnitWidth int64

		if o.UnitWidth != nil {
			qrUnitWidth = *o.UnitWidth
		}
		qUnitWidth := swag.FormatInt64(qrUnitWidth)
		if qUnitWidth != "" {

			if err := r.SetQueryParam("unit_width", qUnitWidth); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

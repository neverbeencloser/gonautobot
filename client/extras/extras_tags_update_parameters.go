package extras

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// NewExtrasTagsUpdateParams creates a new ExtrasTagsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasTagsUpdateParams() *ExtrasTagsUpdateParams {
	return &ExtrasTagsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasTagsUpdateParamsWithTimeout creates a new ExtrasTagsUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasTagsUpdateParamsWithTimeout(timeout time.Duration) *ExtrasTagsUpdateParams {
	return &ExtrasTagsUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasTagsUpdateParamsWithContext creates a new ExtrasTagsUpdateParams object
// with the ability to set a context for a request.
func NewExtrasTagsUpdateParamsWithContext(ctx context.Context) *ExtrasTagsUpdateParams {
	return &ExtrasTagsUpdateParams{
		Context: ctx,
	}
}

// NewExtrasTagsUpdateParamsWithHTTPClient creates a new ExtrasTagsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasTagsUpdateParamsWithHTTPClient(client *http.Client) *ExtrasTagsUpdateParams {
	return &ExtrasTagsUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasTagsUpdateParams contains all the parameters to send to the API endpoint
   for the extras tags update operation.

   Typically these are written to a http.Request.
*/
type ExtrasTagsUpdateParams struct {

	// Data.
	Data *models.Tag

	/* ID.

	   A UUID string identifying this tag.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras tags update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasTagsUpdateParams) WithDefaults() *ExtrasTagsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras tags update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasTagsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras tags update params
func (o *ExtrasTagsUpdateParams) WithTimeout(timeout time.Duration) *ExtrasTagsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras tags update params
func (o *ExtrasTagsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras tags update params
func (o *ExtrasTagsUpdateParams) WithContext(ctx context.Context) *ExtrasTagsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras tags update params
func (o *ExtrasTagsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras tags update params
func (o *ExtrasTagsUpdateParams) WithHTTPClient(client *http.Client) *ExtrasTagsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras tags update params
func (o *ExtrasTagsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras tags update params
func (o *ExtrasTagsUpdateParams) WithData(data *models.Tag) *ExtrasTagsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras tags update params
func (o *ExtrasTagsUpdateParams) SetData(data *models.Tag) {
	o.Data = data
}

// WithID adds the id to the extras tags update params
func (o *ExtrasTagsUpdateParams) WithID(id strfmt.UUID) *ExtrasTagsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras tags update params
func (o *ExtrasTagsUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasTagsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

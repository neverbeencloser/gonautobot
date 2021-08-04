package extras

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewExtrasImageAttachmentsReadParams creates a new ExtrasImageAttachmentsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasImageAttachmentsReadParams() *ExtrasImageAttachmentsReadParams {
	return &ExtrasImageAttachmentsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasImageAttachmentsReadParamsWithTimeout creates a new ExtrasImageAttachmentsReadParams object
// with the ability to set a timeout on a request.
func NewExtrasImageAttachmentsReadParamsWithTimeout(timeout time.Duration) *ExtrasImageAttachmentsReadParams {
	return &ExtrasImageAttachmentsReadParams{
		timeout: timeout,
	}
}

// NewExtrasImageAttachmentsReadParamsWithContext creates a new ExtrasImageAttachmentsReadParams object
// with the ability to set a context for a request.
func NewExtrasImageAttachmentsReadParamsWithContext(ctx context.Context) *ExtrasImageAttachmentsReadParams {
	return &ExtrasImageAttachmentsReadParams{
		Context: ctx,
	}
}

// NewExtrasImageAttachmentsReadParamsWithHTTPClient creates a new ExtrasImageAttachmentsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasImageAttachmentsReadParamsWithHTTPClient(client *http.Client) *ExtrasImageAttachmentsReadParams {
	return &ExtrasImageAttachmentsReadParams{
		HTTPClient: client,
	}
}

/* ExtrasImageAttachmentsReadParams contains all the parameters to send to the API endpoint
   for the extras image attachments read operation.

   Typically these are written to a http.Request.
*/
type ExtrasImageAttachmentsReadParams struct {

	/* ID.

	   A UUID string identifying this image attachment.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras image attachments read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasImageAttachmentsReadParams) WithDefaults() *ExtrasImageAttachmentsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras image attachments read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasImageAttachmentsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras image attachments read params
func (o *ExtrasImageAttachmentsReadParams) WithTimeout(timeout time.Duration) *ExtrasImageAttachmentsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras image attachments read params
func (o *ExtrasImageAttachmentsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras image attachments read params
func (o *ExtrasImageAttachmentsReadParams) WithContext(ctx context.Context) *ExtrasImageAttachmentsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras image attachments read params
func (o *ExtrasImageAttachmentsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras image attachments read params
func (o *ExtrasImageAttachmentsReadParams) WithHTTPClient(client *http.Client) *ExtrasImageAttachmentsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras image attachments read params
func (o *ExtrasImageAttachmentsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the extras image attachments read params
func (o *ExtrasImageAttachmentsReadParams) WithID(id strfmt.UUID) *ExtrasImageAttachmentsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras image attachments read params
func (o *ExtrasImageAttachmentsReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasImageAttachmentsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

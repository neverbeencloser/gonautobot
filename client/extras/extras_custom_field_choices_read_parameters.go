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

// NewExtrasCustomFieldChoicesReadParams creates a new ExtrasCustomFieldChoicesReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasCustomFieldChoicesReadParams() *ExtrasCustomFieldChoicesReadParams {
	return &ExtrasCustomFieldChoicesReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasCustomFieldChoicesReadParamsWithTimeout creates a new ExtrasCustomFieldChoicesReadParams object
// with the ability to set a timeout on a request.
func NewExtrasCustomFieldChoicesReadParamsWithTimeout(timeout time.Duration) *ExtrasCustomFieldChoicesReadParams {
	return &ExtrasCustomFieldChoicesReadParams{
		timeout: timeout,
	}
}

// NewExtrasCustomFieldChoicesReadParamsWithContext creates a new ExtrasCustomFieldChoicesReadParams object
// with the ability to set a context for a request.
func NewExtrasCustomFieldChoicesReadParamsWithContext(ctx context.Context) *ExtrasCustomFieldChoicesReadParams {
	return &ExtrasCustomFieldChoicesReadParams{
		Context: ctx,
	}
}

// NewExtrasCustomFieldChoicesReadParamsWithHTTPClient creates a new ExtrasCustomFieldChoicesReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasCustomFieldChoicesReadParamsWithHTTPClient(client *http.Client) *ExtrasCustomFieldChoicesReadParams {
	return &ExtrasCustomFieldChoicesReadParams{
		HTTPClient: client,
	}
}

/* ExtrasCustomFieldChoicesReadParams contains all the parameters to send to the API endpoint
   for the extras custom field choices read operation.

   Typically these are written to a http.Request.
*/
type ExtrasCustomFieldChoicesReadParams struct {

	/* ID.

	   A UUID string identifying this custom field choice.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras custom field choices read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasCustomFieldChoicesReadParams) WithDefaults() *ExtrasCustomFieldChoicesReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras custom field choices read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasCustomFieldChoicesReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras custom field choices read params
func (o *ExtrasCustomFieldChoicesReadParams) WithTimeout(timeout time.Duration) *ExtrasCustomFieldChoicesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras custom field choices read params
func (o *ExtrasCustomFieldChoicesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras custom field choices read params
func (o *ExtrasCustomFieldChoicesReadParams) WithContext(ctx context.Context) *ExtrasCustomFieldChoicesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras custom field choices read params
func (o *ExtrasCustomFieldChoicesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras custom field choices read params
func (o *ExtrasCustomFieldChoicesReadParams) WithHTTPClient(client *http.Client) *ExtrasCustomFieldChoicesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras custom field choices read params
func (o *ExtrasCustomFieldChoicesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the extras custom field choices read params
func (o *ExtrasCustomFieldChoicesReadParams) WithID(id strfmt.UUID) *ExtrasCustomFieldChoicesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras custom field choices read params
func (o *ExtrasCustomFieldChoicesReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasCustomFieldChoicesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

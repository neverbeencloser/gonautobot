package dcim

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDcimInventoryItemsReadParams creates a new DcimInventoryItemsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimInventoryItemsReadParams() *DcimInventoryItemsReadParams {
	return &DcimInventoryItemsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimInventoryItemsReadParamsWithTimeout creates a new DcimInventoryItemsReadParams object
// with the ability to set a timeout on a request.
func NewDcimInventoryItemsReadParamsWithTimeout(timeout time.Duration) *DcimInventoryItemsReadParams {
	return &DcimInventoryItemsReadParams{
		timeout: timeout,
	}
}

// NewDcimInventoryItemsReadParamsWithContext creates a new DcimInventoryItemsReadParams object
// with the ability to set a context for a request.
func NewDcimInventoryItemsReadParamsWithContext(ctx context.Context) *DcimInventoryItemsReadParams {
	return &DcimInventoryItemsReadParams{
		Context: ctx,
	}
}

// NewDcimInventoryItemsReadParamsWithHTTPClient creates a new DcimInventoryItemsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimInventoryItemsReadParamsWithHTTPClient(client *http.Client) *DcimInventoryItemsReadParams {
	return &DcimInventoryItemsReadParams{
		HTTPClient: client,
	}
}

/* DcimInventoryItemsReadParams contains all the parameters to send to the API endpoint
   for the dcim inventory items read operation.

   Typically these are written to a http.Request.
*/
type DcimInventoryItemsReadParams struct {

	/* ID.

	   A UUID string identifying this inventory item.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim inventory items read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimInventoryItemsReadParams) WithDefaults() *DcimInventoryItemsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim inventory items read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimInventoryItemsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim inventory items read params
func (o *DcimInventoryItemsReadParams) WithTimeout(timeout time.Duration) *DcimInventoryItemsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim inventory items read params
func (o *DcimInventoryItemsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim inventory items read params
func (o *DcimInventoryItemsReadParams) WithContext(ctx context.Context) *DcimInventoryItemsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim inventory items read params
func (o *DcimInventoryItemsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim inventory items read params
func (o *DcimInventoryItemsReadParams) WithHTTPClient(client *http.Client) *DcimInventoryItemsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim inventory items read params
func (o *DcimInventoryItemsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim inventory items read params
func (o *DcimInventoryItemsReadParams) WithID(id strfmt.UUID) *DcimInventoryItemsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim inventory items read params
func (o *DcimInventoryItemsReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimInventoryItemsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

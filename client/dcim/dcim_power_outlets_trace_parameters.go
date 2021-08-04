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

// NewDcimPowerOutletsTraceParams creates a new DcimPowerOutletsTraceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPowerOutletsTraceParams() *DcimPowerOutletsTraceParams {
	return &DcimPowerOutletsTraceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerOutletsTraceParamsWithTimeout creates a new DcimPowerOutletsTraceParams object
// with the ability to set a timeout on a request.
func NewDcimPowerOutletsTraceParamsWithTimeout(timeout time.Duration) *DcimPowerOutletsTraceParams {
	return &DcimPowerOutletsTraceParams{
		timeout: timeout,
	}
}

// NewDcimPowerOutletsTraceParamsWithContext creates a new DcimPowerOutletsTraceParams object
// with the ability to set a context for a request.
func NewDcimPowerOutletsTraceParamsWithContext(ctx context.Context) *DcimPowerOutletsTraceParams {
	return &DcimPowerOutletsTraceParams{
		Context: ctx,
	}
}

// NewDcimPowerOutletsTraceParamsWithHTTPClient creates a new DcimPowerOutletsTraceParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPowerOutletsTraceParamsWithHTTPClient(client *http.Client) *DcimPowerOutletsTraceParams {
	return &DcimPowerOutletsTraceParams{
		HTTPClient: client,
	}
}

/* DcimPowerOutletsTraceParams contains all the parameters to send to the API endpoint
   for the dcim power outlets trace operation.

   Typically these are written to a http.Request.
*/
type DcimPowerOutletsTraceParams struct {

	/* ID.

	   A UUID string identifying this power outlet.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim power outlets trace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerOutletsTraceParams) WithDefaults() *DcimPowerOutletsTraceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim power outlets trace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerOutletsTraceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim power outlets trace params
func (o *DcimPowerOutletsTraceParams) WithTimeout(timeout time.Duration) *DcimPowerOutletsTraceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power outlets trace params
func (o *DcimPowerOutletsTraceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power outlets trace params
func (o *DcimPowerOutletsTraceParams) WithContext(ctx context.Context) *DcimPowerOutletsTraceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power outlets trace params
func (o *DcimPowerOutletsTraceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power outlets trace params
func (o *DcimPowerOutletsTraceParams) WithHTTPClient(client *http.Client) *DcimPowerOutletsTraceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power outlets trace params
func (o *DcimPowerOutletsTraceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim power outlets trace params
func (o *DcimPowerOutletsTraceParams) WithID(id strfmt.UUID) *DcimPowerOutletsTraceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim power outlets trace params
func (o *DcimPowerOutletsTraceParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerOutletsTraceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

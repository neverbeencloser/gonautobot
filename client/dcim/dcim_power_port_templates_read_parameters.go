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

// NewDcimPowerPortTemplatesReadParams creates a new DcimPowerPortTemplatesReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPowerPortTemplatesReadParams() *DcimPowerPortTemplatesReadParams {
	return &DcimPowerPortTemplatesReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerPortTemplatesReadParamsWithTimeout creates a new DcimPowerPortTemplatesReadParams object
// with the ability to set a timeout on a request.
func NewDcimPowerPortTemplatesReadParamsWithTimeout(timeout time.Duration) *DcimPowerPortTemplatesReadParams {
	return &DcimPowerPortTemplatesReadParams{
		timeout: timeout,
	}
}

// NewDcimPowerPortTemplatesReadParamsWithContext creates a new DcimPowerPortTemplatesReadParams object
// with the ability to set a context for a request.
func NewDcimPowerPortTemplatesReadParamsWithContext(ctx context.Context) *DcimPowerPortTemplatesReadParams {
	return &DcimPowerPortTemplatesReadParams{
		Context: ctx,
	}
}

// NewDcimPowerPortTemplatesReadParamsWithHTTPClient creates a new DcimPowerPortTemplatesReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPowerPortTemplatesReadParamsWithHTTPClient(client *http.Client) *DcimPowerPortTemplatesReadParams {
	return &DcimPowerPortTemplatesReadParams{
		HTTPClient: client,
	}
}

/* DcimPowerPortTemplatesReadParams contains all the parameters to send to the API endpoint
   for the dcim power port templates read operation.

   Typically these are written to a http.Request.
*/
type DcimPowerPortTemplatesReadParams struct {

	/* ID.

	   A UUID string identifying this power port template.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim power port templates read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerPortTemplatesReadParams) WithDefaults() *DcimPowerPortTemplatesReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim power port templates read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerPortTemplatesReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim power port templates read params
func (o *DcimPowerPortTemplatesReadParams) WithTimeout(timeout time.Duration) *DcimPowerPortTemplatesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power port templates read params
func (o *DcimPowerPortTemplatesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power port templates read params
func (o *DcimPowerPortTemplatesReadParams) WithContext(ctx context.Context) *DcimPowerPortTemplatesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power port templates read params
func (o *DcimPowerPortTemplatesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power port templates read params
func (o *DcimPowerPortTemplatesReadParams) WithHTTPClient(client *http.Client) *DcimPowerPortTemplatesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power port templates read params
func (o *DcimPowerPortTemplatesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim power port templates read params
func (o *DcimPowerPortTemplatesReadParams) WithID(id strfmt.UUID) *DcimPowerPortTemplatesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim power port templates read params
func (o *DcimPowerPortTemplatesReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerPortTemplatesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

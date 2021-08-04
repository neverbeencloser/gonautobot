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

// NewDcimPowerPortTemplatesDeleteParams creates a new DcimPowerPortTemplatesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPowerPortTemplatesDeleteParams() *DcimPowerPortTemplatesDeleteParams {
	return &DcimPowerPortTemplatesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerPortTemplatesDeleteParamsWithTimeout creates a new DcimPowerPortTemplatesDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimPowerPortTemplatesDeleteParamsWithTimeout(timeout time.Duration) *DcimPowerPortTemplatesDeleteParams {
	return &DcimPowerPortTemplatesDeleteParams{
		timeout: timeout,
	}
}

// NewDcimPowerPortTemplatesDeleteParamsWithContext creates a new DcimPowerPortTemplatesDeleteParams object
// with the ability to set a context for a request.
func NewDcimPowerPortTemplatesDeleteParamsWithContext(ctx context.Context) *DcimPowerPortTemplatesDeleteParams {
	return &DcimPowerPortTemplatesDeleteParams{
		Context: ctx,
	}
}

// NewDcimPowerPortTemplatesDeleteParamsWithHTTPClient creates a new DcimPowerPortTemplatesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPowerPortTemplatesDeleteParamsWithHTTPClient(client *http.Client) *DcimPowerPortTemplatesDeleteParams {
	return &DcimPowerPortTemplatesDeleteParams{
		HTTPClient: client,
	}
}

/* DcimPowerPortTemplatesDeleteParams contains all the parameters to send to the API endpoint
   for the dcim power port templates delete operation.

   Typically these are written to a http.Request.
*/
type DcimPowerPortTemplatesDeleteParams struct {

	/* ID.

	   A UUID string identifying this power port template.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim power port templates delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerPortTemplatesDeleteParams) WithDefaults() *DcimPowerPortTemplatesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim power port templates delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerPortTemplatesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim power port templates delete params
func (o *DcimPowerPortTemplatesDeleteParams) WithTimeout(timeout time.Duration) *DcimPowerPortTemplatesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power port templates delete params
func (o *DcimPowerPortTemplatesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power port templates delete params
func (o *DcimPowerPortTemplatesDeleteParams) WithContext(ctx context.Context) *DcimPowerPortTemplatesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power port templates delete params
func (o *DcimPowerPortTemplatesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power port templates delete params
func (o *DcimPowerPortTemplatesDeleteParams) WithHTTPClient(client *http.Client) *DcimPowerPortTemplatesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power port templates delete params
func (o *DcimPowerPortTemplatesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim power port templates delete params
func (o *DcimPowerPortTemplatesDeleteParams) WithID(id strfmt.UUID) *DcimPowerPortTemplatesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim power port templates delete params
func (o *DcimPowerPortTemplatesDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerPortTemplatesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

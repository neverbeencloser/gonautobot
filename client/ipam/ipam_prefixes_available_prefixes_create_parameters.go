package ipam

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

// NewIpamPrefixesAvailablePrefixesCreateParams creates a new IpamPrefixesAvailablePrefixesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamPrefixesAvailablePrefixesCreateParams() *IpamPrefixesAvailablePrefixesCreateParams {
	return &IpamPrefixesAvailablePrefixesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamPrefixesAvailablePrefixesCreateParamsWithTimeout creates a new IpamPrefixesAvailablePrefixesCreateParams object
// with the ability to set a timeout on a request.
func NewIpamPrefixesAvailablePrefixesCreateParamsWithTimeout(timeout time.Duration) *IpamPrefixesAvailablePrefixesCreateParams {
	return &IpamPrefixesAvailablePrefixesCreateParams{
		timeout: timeout,
	}
}

// NewIpamPrefixesAvailablePrefixesCreateParamsWithContext creates a new IpamPrefixesAvailablePrefixesCreateParams object
// with the ability to set a context for a request.
func NewIpamPrefixesAvailablePrefixesCreateParamsWithContext(ctx context.Context) *IpamPrefixesAvailablePrefixesCreateParams {
	return &IpamPrefixesAvailablePrefixesCreateParams{
		Context: ctx,
	}
}

// NewIpamPrefixesAvailablePrefixesCreateParamsWithHTTPClient creates a new IpamPrefixesAvailablePrefixesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamPrefixesAvailablePrefixesCreateParamsWithHTTPClient(client *http.Client) *IpamPrefixesAvailablePrefixesCreateParams {
	return &IpamPrefixesAvailablePrefixesCreateParams{
		HTTPClient: client,
	}
}

/* IpamPrefixesAvailablePrefixesCreateParams contains all the parameters to send to the API endpoint
   for the ipam prefixes available prefixes create operation.

   Typically these are written to a http.Request.
*/
type IpamPrefixesAvailablePrefixesCreateParams struct {

	// Data.
	Data *models.PrefixLength

	/* ID.

	   A UUID string identifying this prefix.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam prefixes available prefixes create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamPrefixesAvailablePrefixesCreateParams) WithDefaults() *IpamPrefixesAvailablePrefixesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam prefixes available prefixes create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamPrefixesAvailablePrefixesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam prefixes available prefixes create params
func (o *IpamPrefixesAvailablePrefixesCreateParams) WithTimeout(timeout time.Duration) *IpamPrefixesAvailablePrefixesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam prefixes available prefixes create params
func (o *IpamPrefixesAvailablePrefixesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam prefixes available prefixes create params
func (o *IpamPrefixesAvailablePrefixesCreateParams) WithContext(ctx context.Context) *IpamPrefixesAvailablePrefixesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam prefixes available prefixes create params
func (o *IpamPrefixesAvailablePrefixesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam prefixes available prefixes create params
func (o *IpamPrefixesAvailablePrefixesCreateParams) WithHTTPClient(client *http.Client) *IpamPrefixesAvailablePrefixesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam prefixes available prefixes create params
func (o *IpamPrefixesAvailablePrefixesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam prefixes available prefixes create params
func (o *IpamPrefixesAvailablePrefixesCreateParams) WithData(data *models.PrefixLength) *IpamPrefixesAvailablePrefixesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam prefixes available prefixes create params
func (o *IpamPrefixesAvailablePrefixesCreateParams) SetData(data *models.PrefixLength) {
	o.Data = data
}

// WithID adds the id to the ipam prefixes available prefixes create params
func (o *IpamPrefixesAvailablePrefixesCreateParams) WithID(id strfmt.UUID) *IpamPrefixesAvailablePrefixesCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam prefixes available prefixes create params
func (o *IpamPrefixesAvailablePrefixesCreateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamPrefixesAvailablePrefixesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

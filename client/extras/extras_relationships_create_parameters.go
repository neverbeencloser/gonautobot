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

// NewExtrasRelationshipsCreateParams creates a new ExtrasRelationshipsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasRelationshipsCreateParams() *ExtrasRelationshipsCreateParams {
	return &ExtrasRelationshipsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasRelationshipsCreateParamsWithTimeout creates a new ExtrasRelationshipsCreateParams object
// with the ability to set a timeout on a request.
func NewExtrasRelationshipsCreateParamsWithTimeout(timeout time.Duration) *ExtrasRelationshipsCreateParams {
	return &ExtrasRelationshipsCreateParams{
		timeout: timeout,
	}
}

// NewExtrasRelationshipsCreateParamsWithContext creates a new ExtrasRelationshipsCreateParams object
// with the ability to set a context for a request.
func NewExtrasRelationshipsCreateParamsWithContext(ctx context.Context) *ExtrasRelationshipsCreateParams {
	return &ExtrasRelationshipsCreateParams{
		Context: ctx,
	}
}

// NewExtrasRelationshipsCreateParamsWithHTTPClient creates a new ExtrasRelationshipsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasRelationshipsCreateParamsWithHTTPClient(client *http.Client) *ExtrasRelationshipsCreateParams {
	return &ExtrasRelationshipsCreateParams{
		HTTPClient: client,
	}
}

/* ExtrasRelationshipsCreateParams contains all the parameters to send to the API endpoint
   for the extras relationships create operation.

   Typically these are written to a http.Request.
*/
type ExtrasRelationshipsCreateParams struct {

	// Data.
	Data *models.Relationship

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras relationships create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasRelationshipsCreateParams) WithDefaults() *ExtrasRelationshipsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras relationships create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasRelationshipsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras relationships create params
func (o *ExtrasRelationshipsCreateParams) WithTimeout(timeout time.Duration) *ExtrasRelationshipsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras relationships create params
func (o *ExtrasRelationshipsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras relationships create params
func (o *ExtrasRelationshipsCreateParams) WithContext(ctx context.Context) *ExtrasRelationshipsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras relationships create params
func (o *ExtrasRelationshipsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras relationships create params
func (o *ExtrasRelationshipsCreateParams) WithHTTPClient(client *http.Client) *ExtrasRelationshipsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras relationships create params
func (o *ExtrasRelationshipsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras relationships create params
func (o *ExtrasRelationshipsCreateParams) WithData(data *models.Relationship) *ExtrasRelationshipsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras relationships create params
func (o *ExtrasRelationshipsCreateParams) SetData(data *models.Relationship) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasRelationshipsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

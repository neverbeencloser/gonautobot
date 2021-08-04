package circuits

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

// NewCircuitsProvidersPartialUpdateParams creates a new CircuitsProvidersPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCircuitsProvidersPartialUpdateParams() *CircuitsProvidersPartialUpdateParams {
	return &CircuitsProvidersPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsProvidersPartialUpdateParamsWithTimeout creates a new CircuitsProvidersPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewCircuitsProvidersPartialUpdateParamsWithTimeout(timeout time.Duration) *CircuitsProvidersPartialUpdateParams {
	return &CircuitsProvidersPartialUpdateParams{
		timeout: timeout,
	}
}

// NewCircuitsProvidersPartialUpdateParamsWithContext creates a new CircuitsProvidersPartialUpdateParams object
// with the ability to set a context for a request.
func NewCircuitsProvidersPartialUpdateParamsWithContext(ctx context.Context) *CircuitsProvidersPartialUpdateParams {
	return &CircuitsProvidersPartialUpdateParams{
		Context: ctx,
	}
}

// NewCircuitsProvidersPartialUpdateParamsWithHTTPClient creates a new CircuitsProvidersPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewCircuitsProvidersPartialUpdateParamsWithHTTPClient(client *http.Client) *CircuitsProvidersPartialUpdateParams {
	return &CircuitsProvidersPartialUpdateParams{
		HTTPClient: client,
	}
}

/* CircuitsProvidersPartialUpdateParams contains all the parameters to send to the API endpoint
   for the circuits providers partial update operation.

   Typically these are written to a http.Request.
*/
type CircuitsProvidersPartialUpdateParams struct {

	// Data.
	Data *models.Provider

	/* ID.

	   A UUID string identifying this provider.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the circuits providers partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsProvidersPartialUpdateParams) WithDefaults() *CircuitsProvidersPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the circuits providers partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsProvidersPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the circuits providers partial update params
func (o *CircuitsProvidersPartialUpdateParams) WithTimeout(timeout time.Duration) *CircuitsProvidersPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits providers partial update params
func (o *CircuitsProvidersPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits providers partial update params
func (o *CircuitsProvidersPartialUpdateParams) WithContext(ctx context.Context) *CircuitsProvidersPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits providers partial update params
func (o *CircuitsProvidersPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits providers partial update params
func (o *CircuitsProvidersPartialUpdateParams) WithHTTPClient(client *http.Client) *CircuitsProvidersPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits providers partial update params
func (o *CircuitsProvidersPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the circuits providers partial update params
func (o *CircuitsProvidersPartialUpdateParams) WithData(data *models.Provider) *CircuitsProvidersPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the circuits providers partial update params
func (o *CircuitsProvidersPartialUpdateParams) SetData(data *models.Provider) {
	o.Data = data
}

// WithID adds the id to the circuits providers partial update params
func (o *CircuitsProvidersPartialUpdateParams) WithID(id strfmt.UUID) *CircuitsProvidersPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the circuits providers partial update params
func (o *CircuitsProvidersPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsProvidersPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

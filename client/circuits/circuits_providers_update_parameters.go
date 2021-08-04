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

// NewCircuitsProvidersUpdateParams creates a new CircuitsProvidersUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCircuitsProvidersUpdateParams() *CircuitsProvidersUpdateParams {
	return &CircuitsProvidersUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsProvidersUpdateParamsWithTimeout creates a new CircuitsProvidersUpdateParams object
// with the ability to set a timeout on a request.
func NewCircuitsProvidersUpdateParamsWithTimeout(timeout time.Duration) *CircuitsProvidersUpdateParams {
	return &CircuitsProvidersUpdateParams{
		timeout: timeout,
	}
}

// NewCircuitsProvidersUpdateParamsWithContext creates a new CircuitsProvidersUpdateParams object
// with the ability to set a context for a request.
func NewCircuitsProvidersUpdateParamsWithContext(ctx context.Context) *CircuitsProvidersUpdateParams {
	return &CircuitsProvidersUpdateParams{
		Context: ctx,
	}
}

// NewCircuitsProvidersUpdateParamsWithHTTPClient creates a new CircuitsProvidersUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewCircuitsProvidersUpdateParamsWithHTTPClient(client *http.Client) *CircuitsProvidersUpdateParams {
	return &CircuitsProvidersUpdateParams{
		HTTPClient: client,
	}
}

/* CircuitsProvidersUpdateParams contains all the parameters to send to the API endpoint
   for the circuits providers update operation.

   Typically these are written to a http.Request.
*/
type CircuitsProvidersUpdateParams struct {

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

// WithDefaults hydrates default values in the circuits providers update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsProvidersUpdateParams) WithDefaults() *CircuitsProvidersUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the circuits providers update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsProvidersUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the circuits providers update params
func (o *CircuitsProvidersUpdateParams) WithTimeout(timeout time.Duration) *CircuitsProvidersUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits providers update params
func (o *CircuitsProvidersUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits providers update params
func (o *CircuitsProvidersUpdateParams) WithContext(ctx context.Context) *CircuitsProvidersUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits providers update params
func (o *CircuitsProvidersUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits providers update params
func (o *CircuitsProvidersUpdateParams) WithHTTPClient(client *http.Client) *CircuitsProvidersUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits providers update params
func (o *CircuitsProvidersUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the circuits providers update params
func (o *CircuitsProvidersUpdateParams) WithData(data *models.Provider) *CircuitsProvidersUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the circuits providers update params
func (o *CircuitsProvidersUpdateParams) SetData(data *models.Provider) {
	o.Data = data
}

// WithID adds the id to the circuits providers update params
func (o *CircuitsProvidersUpdateParams) WithID(id strfmt.UUID) *CircuitsProvidersUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the circuits providers update params
func (o *CircuitsProvidersUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsProvidersUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

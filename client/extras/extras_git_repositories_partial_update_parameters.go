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

// NewExtrasGitRepositoriesPartialUpdateParams creates a new ExtrasGitRepositoriesPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasGitRepositoriesPartialUpdateParams() *ExtrasGitRepositoriesPartialUpdateParams {
	return &ExtrasGitRepositoriesPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasGitRepositoriesPartialUpdateParamsWithTimeout creates a new ExtrasGitRepositoriesPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasGitRepositoriesPartialUpdateParamsWithTimeout(timeout time.Duration) *ExtrasGitRepositoriesPartialUpdateParams {
	return &ExtrasGitRepositoriesPartialUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasGitRepositoriesPartialUpdateParamsWithContext creates a new ExtrasGitRepositoriesPartialUpdateParams object
// with the ability to set a context for a request.
func NewExtrasGitRepositoriesPartialUpdateParamsWithContext(ctx context.Context) *ExtrasGitRepositoriesPartialUpdateParams {
	return &ExtrasGitRepositoriesPartialUpdateParams{
		Context: ctx,
	}
}

// NewExtrasGitRepositoriesPartialUpdateParamsWithHTTPClient creates a new ExtrasGitRepositoriesPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasGitRepositoriesPartialUpdateParamsWithHTTPClient(client *http.Client) *ExtrasGitRepositoriesPartialUpdateParams {
	return &ExtrasGitRepositoriesPartialUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasGitRepositoriesPartialUpdateParams contains all the parameters to send to the API endpoint
   for the extras git repositories partial update operation.

   Typically these are written to a http.Request.
*/
type ExtrasGitRepositoriesPartialUpdateParams struct {

	// Data.
	Data *models.GitRepository

	/* ID.

	   A UUID string identifying this Git repository.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras git repositories partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasGitRepositoriesPartialUpdateParams) WithDefaults() *ExtrasGitRepositoriesPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras git repositories partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasGitRepositoriesPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras git repositories partial update params
func (o *ExtrasGitRepositoriesPartialUpdateParams) WithTimeout(timeout time.Duration) *ExtrasGitRepositoriesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras git repositories partial update params
func (o *ExtrasGitRepositoriesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras git repositories partial update params
func (o *ExtrasGitRepositoriesPartialUpdateParams) WithContext(ctx context.Context) *ExtrasGitRepositoriesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras git repositories partial update params
func (o *ExtrasGitRepositoriesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras git repositories partial update params
func (o *ExtrasGitRepositoriesPartialUpdateParams) WithHTTPClient(client *http.Client) *ExtrasGitRepositoriesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras git repositories partial update params
func (o *ExtrasGitRepositoriesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras git repositories partial update params
func (o *ExtrasGitRepositoriesPartialUpdateParams) WithData(data *models.GitRepository) *ExtrasGitRepositoriesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras git repositories partial update params
func (o *ExtrasGitRepositoriesPartialUpdateParams) SetData(data *models.GitRepository) {
	o.Data = data
}

// WithID adds the id to the extras git repositories partial update params
func (o *ExtrasGitRepositoriesPartialUpdateParams) WithID(id strfmt.UUID) *ExtrasGitRepositoriesPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras git repositories partial update params
func (o *ExtrasGitRepositoriesPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasGitRepositoriesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

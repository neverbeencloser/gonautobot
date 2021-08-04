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

// NewExtrasGitRepositoriesSyncParams creates a new ExtrasGitRepositoriesSyncParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasGitRepositoriesSyncParams() *ExtrasGitRepositoriesSyncParams {
	return &ExtrasGitRepositoriesSyncParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasGitRepositoriesSyncParamsWithTimeout creates a new ExtrasGitRepositoriesSyncParams object
// with the ability to set a timeout on a request.
func NewExtrasGitRepositoriesSyncParamsWithTimeout(timeout time.Duration) *ExtrasGitRepositoriesSyncParams {
	return &ExtrasGitRepositoriesSyncParams{
		timeout: timeout,
	}
}

// NewExtrasGitRepositoriesSyncParamsWithContext creates a new ExtrasGitRepositoriesSyncParams object
// with the ability to set a context for a request.
func NewExtrasGitRepositoriesSyncParamsWithContext(ctx context.Context) *ExtrasGitRepositoriesSyncParams {
	return &ExtrasGitRepositoriesSyncParams{
		Context: ctx,
	}
}

// NewExtrasGitRepositoriesSyncParamsWithHTTPClient creates a new ExtrasGitRepositoriesSyncParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasGitRepositoriesSyncParamsWithHTTPClient(client *http.Client) *ExtrasGitRepositoriesSyncParams {
	return &ExtrasGitRepositoriesSyncParams{
		HTTPClient: client,
	}
}

/* ExtrasGitRepositoriesSyncParams contains all the parameters to send to the API endpoint
   for the extras git repositories sync operation.

   Typically these are written to a http.Request.
*/
type ExtrasGitRepositoriesSyncParams struct {

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

// WithDefaults hydrates default values in the extras git repositories sync params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasGitRepositoriesSyncParams) WithDefaults() *ExtrasGitRepositoriesSyncParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras git repositories sync params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasGitRepositoriesSyncParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras git repositories sync params
func (o *ExtrasGitRepositoriesSyncParams) WithTimeout(timeout time.Duration) *ExtrasGitRepositoriesSyncParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras git repositories sync params
func (o *ExtrasGitRepositoriesSyncParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras git repositories sync params
func (o *ExtrasGitRepositoriesSyncParams) WithContext(ctx context.Context) *ExtrasGitRepositoriesSyncParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras git repositories sync params
func (o *ExtrasGitRepositoriesSyncParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras git repositories sync params
func (o *ExtrasGitRepositoriesSyncParams) WithHTTPClient(client *http.Client) *ExtrasGitRepositoriesSyncParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras git repositories sync params
func (o *ExtrasGitRepositoriesSyncParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras git repositories sync params
func (o *ExtrasGitRepositoriesSyncParams) WithData(data *models.GitRepository) *ExtrasGitRepositoriesSyncParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras git repositories sync params
func (o *ExtrasGitRepositoriesSyncParams) SetData(data *models.GitRepository) {
	o.Data = data
}

// WithID adds the id to the extras git repositories sync params
func (o *ExtrasGitRepositoriesSyncParams) WithID(id strfmt.UUID) *ExtrasGitRepositoriesSyncParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras git repositories sync params
func (o *ExtrasGitRepositoriesSyncParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasGitRepositoriesSyncParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

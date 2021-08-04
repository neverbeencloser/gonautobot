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

// NewExtrasGitRepositoriesUpdateParams creates a new ExtrasGitRepositoriesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasGitRepositoriesUpdateParams() *ExtrasGitRepositoriesUpdateParams {
	return &ExtrasGitRepositoriesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasGitRepositoriesUpdateParamsWithTimeout creates a new ExtrasGitRepositoriesUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasGitRepositoriesUpdateParamsWithTimeout(timeout time.Duration) *ExtrasGitRepositoriesUpdateParams {
	return &ExtrasGitRepositoriesUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasGitRepositoriesUpdateParamsWithContext creates a new ExtrasGitRepositoriesUpdateParams object
// with the ability to set a context for a request.
func NewExtrasGitRepositoriesUpdateParamsWithContext(ctx context.Context) *ExtrasGitRepositoriesUpdateParams {
	return &ExtrasGitRepositoriesUpdateParams{
		Context: ctx,
	}
}

// NewExtrasGitRepositoriesUpdateParamsWithHTTPClient creates a new ExtrasGitRepositoriesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasGitRepositoriesUpdateParamsWithHTTPClient(client *http.Client) *ExtrasGitRepositoriesUpdateParams {
	return &ExtrasGitRepositoriesUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasGitRepositoriesUpdateParams contains all the parameters to send to the API endpoint
   for the extras git repositories update operation.

   Typically these are written to a http.Request.
*/
type ExtrasGitRepositoriesUpdateParams struct {

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

// WithDefaults hydrates default values in the extras git repositories update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasGitRepositoriesUpdateParams) WithDefaults() *ExtrasGitRepositoriesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras git repositories update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasGitRepositoriesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras git repositories update params
func (o *ExtrasGitRepositoriesUpdateParams) WithTimeout(timeout time.Duration) *ExtrasGitRepositoriesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras git repositories update params
func (o *ExtrasGitRepositoriesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras git repositories update params
func (o *ExtrasGitRepositoriesUpdateParams) WithContext(ctx context.Context) *ExtrasGitRepositoriesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras git repositories update params
func (o *ExtrasGitRepositoriesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras git repositories update params
func (o *ExtrasGitRepositoriesUpdateParams) WithHTTPClient(client *http.Client) *ExtrasGitRepositoriesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras git repositories update params
func (o *ExtrasGitRepositoriesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras git repositories update params
func (o *ExtrasGitRepositoriesUpdateParams) WithData(data *models.GitRepository) *ExtrasGitRepositoriesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras git repositories update params
func (o *ExtrasGitRepositoriesUpdateParams) SetData(data *models.GitRepository) {
	o.Data = data
}

// WithID adds the id to the extras git repositories update params
func (o *ExtrasGitRepositoriesUpdateParams) WithID(id strfmt.UUID) *ExtrasGitRepositoriesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras git repositories update params
func (o *ExtrasGitRepositoriesUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasGitRepositoriesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

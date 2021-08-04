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

// NewExtrasGitRepositoriesCreateParams creates a new ExtrasGitRepositoriesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasGitRepositoriesCreateParams() *ExtrasGitRepositoriesCreateParams {
	return &ExtrasGitRepositoriesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasGitRepositoriesCreateParamsWithTimeout creates a new ExtrasGitRepositoriesCreateParams object
// with the ability to set a timeout on a request.
func NewExtrasGitRepositoriesCreateParamsWithTimeout(timeout time.Duration) *ExtrasGitRepositoriesCreateParams {
	return &ExtrasGitRepositoriesCreateParams{
		timeout: timeout,
	}
}

// NewExtrasGitRepositoriesCreateParamsWithContext creates a new ExtrasGitRepositoriesCreateParams object
// with the ability to set a context for a request.
func NewExtrasGitRepositoriesCreateParamsWithContext(ctx context.Context) *ExtrasGitRepositoriesCreateParams {
	return &ExtrasGitRepositoriesCreateParams{
		Context: ctx,
	}
}

// NewExtrasGitRepositoriesCreateParamsWithHTTPClient creates a new ExtrasGitRepositoriesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasGitRepositoriesCreateParamsWithHTTPClient(client *http.Client) *ExtrasGitRepositoriesCreateParams {
	return &ExtrasGitRepositoriesCreateParams{
		HTTPClient: client,
	}
}

/* ExtrasGitRepositoriesCreateParams contains all the parameters to send to the API endpoint
   for the extras git repositories create operation.

   Typically these are written to a http.Request.
*/
type ExtrasGitRepositoriesCreateParams struct {

	// Data.
	Data *models.GitRepository

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras git repositories create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasGitRepositoriesCreateParams) WithDefaults() *ExtrasGitRepositoriesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras git repositories create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasGitRepositoriesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras git repositories create params
func (o *ExtrasGitRepositoriesCreateParams) WithTimeout(timeout time.Duration) *ExtrasGitRepositoriesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras git repositories create params
func (o *ExtrasGitRepositoriesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras git repositories create params
func (o *ExtrasGitRepositoriesCreateParams) WithContext(ctx context.Context) *ExtrasGitRepositoriesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras git repositories create params
func (o *ExtrasGitRepositoriesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras git repositories create params
func (o *ExtrasGitRepositoriesCreateParams) WithHTTPClient(client *http.Client) *ExtrasGitRepositoriesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras git repositories create params
func (o *ExtrasGitRepositoriesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras git repositories create params
func (o *ExtrasGitRepositoriesCreateParams) WithData(data *models.GitRepository) *ExtrasGitRepositoriesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras git repositories create params
func (o *ExtrasGitRepositoriesCreateParams) SetData(data *models.GitRepository) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasGitRepositoriesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

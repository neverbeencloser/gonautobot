package extras

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewExtrasGitRepositoriesReadParams creates a new ExtrasGitRepositoriesReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasGitRepositoriesReadParams() *ExtrasGitRepositoriesReadParams {
	return &ExtrasGitRepositoriesReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasGitRepositoriesReadParamsWithTimeout creates a new ExtrasGitRepositoriesReadParams object
// with the ability to set a timeout on a request.
func NewExtrasGitRepositoriesReadParamsWithTimeout(timeout time.Duration) *ExtrasGitRepositoriesReadParams {
	return &ExtrasGitRepositoriesReadParams{
		timeout: timeout,
	}
}

// NewExtrasGitRepositoriesReadParamsWithContext creates a new ExtrasGitRepositoriesReadParams object
// with the ability to set a context for a request.
func NewExtrasGitRepositoriesReadParamsWithContext(ctx context.Context) *ExtrasGitRepositoriesReadParams {
	return &ExtrasGitRepositoriesReadParams{
		Context: ctx,
	}
}

// NewExtrasGitRepositoriesReadParamsWithHTTPClient creates a new ExtrasGitRepositoriesReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasGitRepositoriesReadParamsWithHTTPClient(client *http.Client) *ExtrasGitRepositoriesReadParams {
	return &ExtrasGitRepositoriesReadParams{
		HTTPClient: client,
	}
}

/* ExtrasGitRepositoriesReadParams contains all the parameters to send to the API endpoint
   for the extras git repositories read operation.

   Typically these are written to a http.Request.
*/
type ExtrasGitRepositoriesReadParams struct {

	/* ID.

	   A UUID string identifying this Git repository.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras git repositories read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasGitRepositoriesReadParams) WithDefaults() *ExtrasGitRepositoriesReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras git repositories read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasGitRepositoriesReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras git repositories read params
func (o *ExtrasGitRepositoriesReadParams) WithTimeout(timeout time.Duration) *ExtrasGitRepositoriesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras git repositories read params
func (o *ExtrasGitRepositoriesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras git repositories read params
func (o *ExtrasGitRepositoriesReadParams) WithContext(ctx context.Context) *ExtrasGitRepositoriesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras git repositories read params
func (o *ExtrasGitRepositoriesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras git repositories read params
func (o *ExtrasGitRepositoriesReadParams) WithHTTPClient(client *http.Client) *ExtrasGitRepositoriesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras git repositories read params
func (o *ExtrasGitRepositoriesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the extras git repositories read params
func (o *ExtrasGitRepositoriesReadParams) WithID(id strfmt.UUID) *ExtrasGitRepositoriesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras git repositories read params
func (o *ExtrasGitRepositoriesReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasGitRepositoriesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

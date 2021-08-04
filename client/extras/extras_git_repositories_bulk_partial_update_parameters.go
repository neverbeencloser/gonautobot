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

// NewExtrasGitRepositoriesBulkPartialUpdateParams creates a new ExtrasGitRepositoriesBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasGitRepositoriesBulkPartialUpdateParams() *ExtrasGitRepositoriesBulkPartialUpdateParams {
	return &ExtrasGitRepositoriesBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasGitRepositoriesBulkPartialUpdateParamsWithTimeout creates a new ExtrasGitRepositoriesBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasGitRepositoriesBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *ExtrasGitRepositoriesBulkPartialUpdateParams {
	return &ExtrasGitRepositoriesBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasGitRepositoriesBulkPartialUpdateParamsWithContext creates a new ExtrasGitRepositoriesBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewExtrasGitRepositoriesBulkPartialUpdateParamsWithContext(ctx context.Context) *ExtrasGitRepositoriesBulkPartialUpdateParams {
	return &ExtrasGitRepositoriesBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewExtrasGitRepositoriesBulkPartialUpdateParamsWithHTTPClient creates a new ExtrasGitRepositoriesBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasGitRepositoriesBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *ExtrasGitRepositoriesBulkPartialUpdateParams {
	return &ExtrasGitRepositoriesBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasGitRepositoriesBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the extras git repositories bulk partial update operation.

   Typically these are written to a http.Request.
*/
type ExtrasGitRepositoriesBulkPartialUpdateParams struct {

	// Data.
	Data *models.GitRepository

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras git repositories bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasGitRepositoriesBulkPartialUpdateParams) WithDefaults() *ExtrasGitRepositoriesBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras git repositories bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasGitRepositoriesBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras git repositories bulk partial update params
func (o *ExtrasGitRepositoriesBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *ExtrasGitRepositoriesBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras git repositories bulk partial update params
func (o *ExtrasGitRepositoriesBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras git repositories bulk partial update params
func (o *ExtrasGitRepositoriesBulkPartialUpdateParams) WithContext(ctx context.Context) *ExtrasGitRepositoriesBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras git repositories bulk partial update params
func (o *ExtrasGitRepositoriesBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras git repositories bulk partial update params
func (o *ExtrasGitRepositoriesBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *ExtrasGitRepositoriesBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras git repositories bulk partial update params
func (o *ExtrasGitRepositoriesBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras git repositories bulk partial update params
func (o *ExtrasGitRepositoriesBulkPartialUpdateParams) WithData(data *models.GitRepository) *ExtrasGitRepositoriesBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras git repositories bulk partial update params
func (o *ExtrasGitRepositoriesBulkPartialUpdateParams) SetData(data *models.GitRepository) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasGitRepositoriesBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

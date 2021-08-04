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

// NewExtrasGitRepositoriesBulkUpdateParams creates a new ExtrasGitRepositoriesBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasGitRepositoriesBulkUpdateParams() *ExtrasGitRepositoriesBulkUpdateParams {
	return &ExtrasGitRepositoriesBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasGitRepositoriesBulkUpdateParamsWithTimeout creates a new ExtrasGitRepositoriesBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasGitRepositoriesBulkUpdateParamsWithTimeout(timeout time.Duration) *ExtrasGitRepositoriesBulkUpdateParams {
	return &ExtrasGitRepositoriesBulkUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasGitRepositoriesBulkUpdateParamsWithContext creates a new ExtrasGitRepositoriesBulkUpdateParams object
// with the ability to set a context for a request.
func NewExtrasGitRepositoriesBulkUpdateParamsWithContext(ctx context.Context) *ExtrasGitRepositoriesBulkUpdateParams {
	return &ExtrasGitRepositoriesBulkUpdateParams{
		Context: ctx,
	}
}

// NewExtrasGitRepositoriesBulkUpdateParamsWithHTTPClient creates a new ExtrasGitRepositoriesBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasGitRepositoriesBulkUpdateParamsWithHTTPClient(client *http.Client) *ExtrasGitRepositoriesBulkUpdateParams {
	return &ExtrasGitRepositoriesBulkUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasGitRepositoriesBulkUpdateParams contains all the parameters to send to the API endpoint
   for the extras git repositories bulk update operation.

   Typically these are written to a http.Request.
*/
type ExtrasGitRepositoriesBulkUpdateParams struct {

	// Data.
	Data *models.GitRepository

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras git repositories bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasGitRepositoriesBulkUpdateParams) WithDefaults() *ExtrasGitRepositoriesBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras git repositories bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasGitRepositoriesBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras git repositories bulk update params
func (o *ExtrasGitRepositoriesBulkUpdateParams) WithTimeout(timeout time.Duration) *ExtrasGitRepositoriesBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras git repositories bulk update params
func (o *ExtrasGitRepositoriesBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras git repositories bulk update params
func (o *ExtrasGitRepositoriesBulkUpdateParams) WithContext(ctx context.Context) *ExtrasGitRepositoriesBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras git repositories bulk update params
func (o *ExtrasGitRepositoriesBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras git repositories bulk update params
func (o *ExtrasGitRepositoriesBulkUpdateParams) WithHTTPClient(client *http.Client) *ExtrasGitRepositoriesBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras git repositories bulk update params
func (o *ExtrasGitRepositoriesBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras git repositories bulk update params
func (o *ExtrasGitRepositoriesBulkUpdateParams) WithData(data *models.GitRepository) *ExtrasGitRepositoriesBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras git repositories bulk update params
func (o *ExtrasGitRepositoriesBulkUpdateParams) SetData(data *models.GitRepository) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasGitRepositoriesBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

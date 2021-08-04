package plugins

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

// NewPluginsDeviceOnboardingOnboardingCreateParams creates a new PluginsDeviceOnboardingOnboardingCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsDeviceOnboardingOnboardingCreateParams() *PluginsDeviceOnboardingOnboardingCreateParams {
	return &PluginsDeviceOnboardingOnboardingCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsDeviceOnboardingOnboardingCreateParamsWithTimeout creates a new PluginsDeviceOnboardingOnboardingCreateParams object
// with the ability to set a timeout on a request.
func NewPluginsDeviceOnboardingOnboardingCreateParamsWithTimeout(timeout time.Duration) *PluginsDeviceOnboardingOnboardingCreateParams {
	return &PluginsDeviceOnboardingOnboardingCreateParams{
		timeout: timeout,
	}
}

// NewPluginsDeviceOnboardingOnboardingCreateParamsWithContext creates a new PluginsDeviceOnboardingOnboardingCreateParams object
// with the ability to set a context for a request.
func NewPluginsDeviceOnboardingOnboardingCreateParamsWithContext(ctx context.Context) *PluginsDeviceOnboardingOnboardingCreateParams {
	return &PluginsDeviceOnboardingOnboardingCreateParams{
		Context: ctx,
	}
}

// NewPluginsDeviceOnboardingOnboardingCreateParamsWithHTTPClient creates a new PluginsDeviceOnboardingOnboardingCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsDeviceOnboardingOnboardingCreateParamsWithHTTPClient(client *http.Client) *PluginsDeviceOnboardingOnboardingCreateParams {
	return &PluginsDeviceOnboardingOnboardingCreateParams{
		HTTPClient: client,
	}
}

/* PluginsDeviceOnboardingOnboardingCreateParams contains all the parameters to send to the API endpoint
   for the plugins device onboarding onboarding create operation.

   Typically these are written to a http.Request.
*/
type PluginsDeviceOnboardingOnboardingCreateParams struct {

	// Data.
	Data *models.OnboardingTask

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins device onboarding onboarding create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsDeviceOnboardingOnboardingCreateParams) WithDefaults() *PluginsDeviceOnboardingOnboardingCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins device onboarding onboarding create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsDeviceOnboardingOnboardingCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins device onboarding onboarding create params
func (o *PluginsDeviceOnboardingOnboardingCreateParams) WithTimeout(timeout time.Duration) *PluginsDeviceOnboardingOnboardingCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins device onboarding onboarding create params
func (o *PluginsDeviceOnboardingOnboardingCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins device onboarding onboarding create params
func (o *PluginsDeviceOnboardingOnboardingCreateParams) WithContext(ctx context.Context) *PluginsDeviceOnboardingOnboardingCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins device onboarding onboarding create params
func (o *PluginsDeviceOnboardingOnboardingCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins device onboarding onboarding create params
func (o *PluginsDeviceOnboardingOnboardingCreateParams) WithHTTPClient(client *http.Client) *PluginsDeviceOnboardingOnboardingCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins device onboarding onboarding create params
func (o *PluginsDeviceOnboardingOnboardingCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins device onboarding onboarding create params
func (o *PluginsDeviceOnboardingOnboardingCreateParams) WithData(data *models.OnboardingTask) *PluginsDeviceOnboardingOnboardingCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins device onboarding onboarding create params
func (o *PluginsDeviceOnboardingOnboardingCreateParams) SetData(data *models.OnboardingTask) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsDeviceOnboardingOnboardingCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

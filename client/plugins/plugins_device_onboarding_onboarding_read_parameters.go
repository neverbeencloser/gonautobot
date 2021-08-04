package plugins

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewPluginsDeviceOnboardingOnboardingReadParams creates a new PluginsDeviceOnboardingOnboardingReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsDeviceOnboardingOnboardingReadParams() *PluginsDeviceOnboardingOnboardingReadParams {
	return &PluginsDeviceOnboardingOnboardingReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsDeviceOnboardingOnboardingReadParamsWithTimeout creates a new PluginsDeviceOnboardingOnboardingReadParams object
// with the ability to set a timeout on a request.
func NewPluginsDeviceOnboardingOnboardingReadParamsWithTimeout(timeout time.Duration) *PluginsDeviceOnboardingOnboardingReadParams {
	return &PluginsDeviceOnboardingOnboardingReadParams{
		timeout: timeout,
	}
}

// NewPluginsDeviceOnboardingOnboardingReadParamsWithContext creates a new PluginsDeviceOnboardingOnboardingReadParams object
// with the ability to set a context for a request.
func NewPluginsDeviceOnboardingOnboardingReadParamsWithContext(ctx context.Context) *PluginsDeviceOnboardingOnboardingReadParams {
	return &PluginsDeviceOnboardingOnboardingReadParams{
		Context: ctx,
	}
}

// NewPluginsDeviceOnboardingOnboardingReadParamsWithHTTPClient creates a new PluginsDeviceOnboardingOnboardingReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsDeviceOnboardingOnboardingReadParamsWithHTTPClient(client *http.Client) *PluginsDeviceOnboardingOnboardingReadParams {
	return &PluginsDeviceOnboardingOnboardingReadParams{
		HTTPClient: client,
	}
}

/* PluginsDeviceOnboardingOnboardingReadParams contains all the parameters to send to the API endpoint
   for the plugins device onboarding onboarding read operation.

   Typically these are written to a http.Request.
*/
type PluginsDeviceOnboardingOnboardingReadParams struct {

	/* ID.

	   A UUID string identifying this onboarding task.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins device onboarding onboarding read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsDeviceOnboardingOnboardingReadParams) WithDefaults() *PluginsDeviceOnboardingOnboardingReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins device onboarding onboarding read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsDeviceOnboardingOnboardingReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins device onboarding onboarding read params
func (o *PluginsDeviceOnboardingOnboardingReadParams) WithTimeout(timeout time.Duration) *PluginsDeviceOnboardingOnboardingReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins device onboarding onboarding read params
func (o *PluginsDeviceOnboardingOnboardingReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins device onboarding onboarding read params
func (o *PluginsDeviceOnboardingOnboardingReadParams) WithContext(ctx context.Context) *PluginsDeviceOnboardingOnboardingReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins device onboarding onboarding read params
func (o *PluginsDeviceOnboardingOnboardingReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins device onboarding onboarding read params
func (o *PluginsDeviceOnboardingOnboardingReadParams) WithHTTPClient(client *http.Client) *PluginsDeviceOnboardingOnboardingReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins device onboarding onboarding read params
func (o *PluginsDeviceOnboardingOnboardingReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the plugins device onboarding onboarding read params
func (o *PluginsDeviceOnboardingOnboardingReadParams) WithID(id strfmt.UUID) *PluginsDeviceOnboardingOnboardingReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins device onboarding onboarding read params
func (o *PluginsDeviceOnboardingOnboardingReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsDeviceOnboardingOnboardingReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

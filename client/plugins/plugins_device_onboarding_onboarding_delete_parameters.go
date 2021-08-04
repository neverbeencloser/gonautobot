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

// NewPluginsDeviceOnboardingOnboardingDeleteParams creates a new PluginsDeviceOnboardingOnboardingDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsDeviceOnboardingOnboardingDeleteParams() *PluginsDeviceOnboardingOnboardingDeleteParams {
	return &PluginsDeviceOnboardingOnboardingDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsDeviceOnboardingOnboardingDeleteParamsWithTimeout creates a new PluginsDeviceOnboardingOnboardingDeleteParams object
// with the ability to set a timeout on a request.
func NewPluginsDeviceOnboardingOnboardingDeleteParamsWithTimeout(timeout time.Duration) *PluginsDeviceOnboardingOnboardingDeleteParams {
	return &PluginsDeviceOnboardingOnboardingDeleteParams{
		timeout: timeout,
	}
}

// NewPluginsDeviceOnboardingOnboardingDeleteParamsWithContext creates a new PluginsDeviceOnboardingOnboardingDeleteParams object
// with the ability to set a context for a request.
func NewPluginsDeviceOnboardingOnboardingDeleteParamsWithContext(ctx context.Context) *PluginsDeviceOnboardingOnboardingDeleteParams {
	return &PluginsDeviceOnboardingOnboardingDeleteParams{
		Context: ctx,
	}
}

// NewPluginsDeviceOnboardingOnboardingDeleteParamsWithHTTPClient creates a new PluginsDeviceOnboardingOnboardingDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsDeviceOnboardingOnboardingDeleteParamsWithHTTPClient(client *http.Client) *PluginsDeviceOnboardingOnboardingDeleteParams {
	return &PluginsDeviceOnboardingOnboardingDeleteParams{
		HTTPClient: client,
	}
}

/* PluginsDeviceOnboardingOnboardingDeleteParams contains all the parameters to send to the API endpoint
   for the plugins device onboarding onboarding delete operation.

   Typically these are written to a http.Request.
*/
type PluginsDeviceOnboardingOnboardingDeleteParams struct {

	/* ID.

	   A UUID string identifying this onboarding task.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins device onboarding onboarding delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsDeviceOnboardingOnboardingDeleteParams) WithDefaults() *PluginsDeviceOnboardingOnboardingDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins device onboarding onboarding delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsDeviceOnboardingOnboardingDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins device onboarding onboarding delete params
func (o *PluginsDeviceOnboardingOnboardingDeleteParams) WithTimeout(timeout time.Duration) *PluginsDeviceOnboardingOnboardingDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins device onboarding onboarding delete params
func (o *PluginsDeviceOnboardingOnboardingDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins device onboarding onboarding delete params
func (o *PluginsDeviceOnboardingOnboardingDeleteParams) WithContext(ctx context.Context) *PluginsDeviceOnboardingOnboardingDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins device onboarding onboarding delete params
func (o *PluginsDeviceOnboardingOnboardingDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins device onboarding onboarding delete params
func (o *PluginsDeviceOnboardingOnboardingDeleteParams) WithHTTPClient(client *http.Client) *PluginsDeviceOnboardingOnboardingDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins device onboarding onboarding delete params
func (o *PluginsDeviceOnboardingOnboardingDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the plugins device onboarding onboarding delete params
func (o *PluginsDeviceOnboardingOnboardingDeleteParams) WithID(id strfmt.UUID) *PluginsDeviceOnboardingOnboardingDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins device onboarding onboarding delete params
func (o *PluginsDeviceOnboardingOnboardingDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsDeviceOnboardingOnboardingDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

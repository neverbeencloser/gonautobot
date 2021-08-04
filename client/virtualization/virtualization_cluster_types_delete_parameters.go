package virtualization

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewVirtualizationClusterTypesDeleteParams creates a new VirtualizationClusterTypesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVirtualizationClusterTypesDeleteParams() *VirtualizationClusterTypesDeleteParams {
	return &VirtualizationClusterTypesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationClusterTypesDeleteParamsWithTimeout creates a new VirtualizationClusterTypesDeleteParams object
// with the ability to set a timeout on a request.
func NewVirtualizationClusterTypesDeleteParamsWithTimeout(timeout time.Duration) *VirtualizationClusterTypesDeleteParams {
	return &VirtualizationClusterTypesDeleteParams{
		timeout: timeout,
	}
}

// NewVirtualizationClusterTypesDeleteParamsWithContext creates a new VirtualizationClusterTypesDeleteParams object
// with the ability to set a context for a request.
func NewVirtualizationClusterTypesDeleteParamsWithContext(ctx context.Context) *VirtualizationClusterTypesDeleteParams {
	return &VirtualizationClusterTypesDeleteParams{
		Context: ctx,
	}
}

// NewVirtualizationClusterTypesDeleteParamsWithHTTPClient creates a new VirtualizationClusterTypesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewVirtualizationClusterTypesDeleteParamsWithHTTPClient(client *http.Client) *VirtualizationClusterTypesDeleteParams {
	return &VirtualizationClusterTypesDeleteParams{
		HTTPClient: client,
	}
}

/* VirtualizationClusterTypesDeleteParams contains all the parameters to send to the API endpoint
   for the virtualization cluster types delete operation.

   Typically these are written to a http.Request.
*/
type VirtualizationClusterTypesDeleteParams struct {

	/* ID.

	   A UUID string identifying this cluster type.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the virtualization cluster types delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationClusterTypesDeleteParams) WithDefaults() *VirtualizationClusterTypesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the virtualization cluster types delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationClusterTypesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the virtualization cluster types delete params
func (o *VirtualizationClusterTypesDeleteParams) WithTimeout(timeout time.Duration) *VirtualizationClusterTypesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization cluster types delete params
func (o *VirtualizationClusterTypesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization cluster types delete params
func (o *VirtualizationClusterTypesDeleteParams) WithContext(ctx context.Context) *VirtualizationClusterTypesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization cluster types delete params
func (o *VirtualizationClusterTypesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization cluster types delete params
func (o *VirtualizationClusterTypesDeleteParams) WithHTTPClient(client *http.Client) *VirtualizationClusterTypesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization cluster types delete params
func (o *VirtualizationClusterTypesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the virtualization cluster types delete params
func (o *VirtualizationClusterTypesDeleteParams) WithID(id strfmt.UUID) *VirtualizationClusterTypesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the virtualization cluster types delete params
func (o *VirtualizationClusterTypesDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationClusterTypesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

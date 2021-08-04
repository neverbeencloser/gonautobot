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

// NewVirtualizationClustersReadParams creates a new VirtualizationClustersReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVirtualizationClustersReadParams() *VirtualizationClustersReadParams {
	return &VirtualizationClustersReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationClustersReadParamsWithTimeout creates a new VirtualizationClustersReadParams object
// with the ability to set a timeout on a request.
func NewVirtualizationClustersReadParamsWithTimeout(timeout time.Duration) *VirtualizationClustersReadParams {
	return &VirtualizationClustersReadParams{
		timeout: timeout,
	}
}

// NewVirtualizationClustersReadParamsWithContext creates a new VirtualizationClustersReadParams object
// with the ability to set a context for a request.
func NewVirtualizationClustersReadParamsWithContext(ctx context.Context) *VirtualizationClustersReadParams {
	return &VirtualizationClustersReadParams{
		Context: ctx,
	}
}

// NewVirtualizationClustersReadParamsWithHTTPClient creates a new VirtualizationClustersReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewVirtualizationClustersReadParamsWithHTTPClient(client *http.Client) *VirtualizationClustersReadParams {
	return &VirtualizationClustersReadParams{
		HTTPClient: client,
	}
}

/* VirtualizationClustersReadParams contains all the parameters to send to the API endpoint
   for the virtualization clusters read operation.

   Typically these are written to a http.Request.
*/
type VirtualizationClustersReadParams struct {

	/* ID.

	   A UUID string identifying this cluster.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the virtualization clusters read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationClustersReadParams) WithDefaults() *VirtualizationClustersReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the virtualization clusters read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationClustersReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the virtualization clusters read params
func (o *VirtualizationClustersReadParams) WithTimeout(timeout time.Duration) *VirtualizationClustersReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization clusters read params
func (o *VirtualizationClustersReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization clusters read params
func (o *VirtualizationClustersReadParams) WithContext(ctx context.Context) *VirtualizationClustersReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization clusters read params
func (o *VirtualizationClustersReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization clusters read params
func (o *VirtualizationClustersReadParams) WithHTTPClient(client *http.Client) *VirtualizationClustersReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization clusters read params
func (o *VirtualizationClustersReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the virtualization clusters read params
func (o *VirtualizationClustersReadParams) WithID(id strfmt.UUID) *VirtualizationClustersReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the virtualization clusters read params
func (o *VirtualizationClustersReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationClustersReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

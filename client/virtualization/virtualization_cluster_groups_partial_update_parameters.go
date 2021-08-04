package virtualization

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

// NewVirtualizationClusterGroupsPartialUpdateParams creates a new VirtualizationClusterGroupsPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVirtualizationClusterGroupsPartialUpdateParams() *VirtualizationClusterGroupsPartialUpdateParams {
	return &VirtualizationClusterGroupsPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationClusterGroupsPartialUpdateParamsWithTimeout creates a new VirtualizationClusterGroupsPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewVirtualizationClusterGroupsPartialUpdateParamsWithTimeout(timeout time.Duration) *VirtualizationClusterGroupsPartialUpdateParams {
	return &VirtualizationClusterGroupsPartialUpdateParams{
		timeout: timeout,
	}
}

// NewVirtualizationClusterGroupsPartialUpdateParamsWithContext creates a new VirtualizationClusterGroupsPartialUpdateParams object
// with the ability to set a context for a request.
func NewVirtualizationClusterGroupsPartialUpdateParamsWithContext(ctx context.Context) *VirtualizationClusterGroupsPartialUpdateParams {
	return &VirtualizationClusterGroupsPartialUpdateParams{
		Context: ctx,
	}
}

// NewVirtualizationClusterGroupsPartialUpdateParamsWithHTTPClient creates a new VirtualizationClusterGroupsPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewVirtualizationClusterGroupsPartialUpdateParamsWithHTTPClient(client *http.Client) *VirtualizationClusterGroupsPartialUpdateParams {
	return &VirtualizationClusterGroupsPartialUpdateParams{
		HTTPClient: client,
	}
}

/* VirtualizationClusterGroupsPartialUpdateParams contains all the parameters to send to the API endpoint
   for the virtualization cluster groups partial update operation.

   Typically these are written to a http.Request.
*/
type VirtualizationClusterGroupsPartialUpdateParams struct {

	// Data.
	Data *models.ClusterGroup

	/* ID.

	   A UUID string identifying this cluster group.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the virtualization cluster groups partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationClusterGroupsPartialUpdateParams) WithDefaults() *VirtualizationClusterGroupsPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the virtualization cluster groups partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationClusterGroupsPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the virtualization cluster groups partial update params
func (o *VirtualizationClusterGroupsPartialUpdateParams) WithTimeout(timeout time.Duration) *VirtualizationClusterGroupsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization cluster groups partial update params
func (o *VirtualizationClusterGroupsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization cluster groups partial update params
func (o *VirtualizationClusterGroupsPartialUpdateParams) WithContext(ctx context.Context) *VirtualizationClusterGroupsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization cluster groups partial update params
func (o *VirtualizationClusterGroupsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization cluster groups partial update params
func (o *VirtualizationClusterGroupsPartialUpdateParams) WithHTTPClient(client *http.Client) *VirtualizationClusterGroupsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization cluster groups partial update params
func (o *VirtualizationClusterGroupsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the virtualization cluster groups partial update params
func (o *VirtualizationClusterGroupsPartialUpdateParams) WithData(data *models.ClusterGroup) *VirtualizationClusterGroupsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the virtualization cluster groups partial update params
func (o *VirtualizationClusterGroupsPartialUpdateParams) SetData(data *models.ClusterGroup) {
	o.Data = data
}

// WithID adds the id to the virtualization cluster groups partial update params
func (o *VirtualizationClusterGroupsPartialUpdateParams) WithID(id strfmt.UUID) *VirtualizationClusterGroupsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the virtualization cluster groups partial update params
func (o *VirtualizationClusterGroupsPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationClusterGroupsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

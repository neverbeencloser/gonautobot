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

// NewVirtualizationClustersPartialUpdateParams creates a new VirtualizationClustersPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVirtualizationClustersPartialUpdateParams() *VirtualizationClustersPartialUpdateParams {
	return &VirtualizationClustersPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationClustersPartialUpdateParamsWithTimeout creates a new VirtualizationClustersPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewVirtualizationClustersPartialUpdateParamsWithTimeout(timeout time.Duration) *VirtualizationClustersPartialUpdateParams {
	return &VirtualizationClustersPartialUpdateParams{
		timeout: timeout,
	}
}

// NewVirtualizationClustersPartialUpdateParamsWithContext creates a new VirtualizationClustersPartialUpdateParams object
// with the ability to set a context for a request.
func NewVirtualizationClustersPartialUpdateParamsWithContext(ctx context.Context) *VirtualizationClustersPartialUpdateParams {
	return &VirtualizationClustersPartialUpdateParams{
		Context: ctx,
	}
}

// NewVirtualizationClustersPartialUpdateParamsWithHTTPClient creates a new VirtualizationClustersPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewVirtualizationClustersPartialUpdateParamsWithHTTPClient(client *http.Client) *VirtualizationClustersPartialUpdateParams {
	return &VirtualizationClustersPartialUpdateParams{
		HTTPClient: client,
	}
}

/* VirtualizationClustersPartialUpdateParams contains all the parameters to send to the API endpoint
   for the virtualization clusters partial update operation.

   Typically these are written to a http.Request.
*/
type VirtualizationClustersPartialUpdateParams struct {

	// Data.
	Data *models.WritableCluster

	/* ID.

	   A UUID string identifying this cluster.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the virtualization clusters partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationClustersPartialUpdateParams) WithDefaults() *VirtualizationClustersPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the virtualization clusters partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationClustersPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the virtualization clusters partial update params
func (o *VirtualizationClustersPartialUpdateParams) WithTimeout(timeout time.Duration) *VirtualizationClustersPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization clusters partial update params
func (o *VirtualizationClustersPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization clusters partial update params
func (o *VirtualizationClustersPartialUpdateParams) WithContext(ctx context.Context) *VirtualizationClustersPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization clusters partial update params
func (o *VirtualizationClustersPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization clusters partial update params
func (o *VirtualizationClustersPartialUpdateParams) WithHTTPClient(client *http.Client) *VirtualizationClustersPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization clusters partial update params
func (o *VirtualizationClustersPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the virtualization clusters partial update params
func (o *VirtualizationClustersPartialUpdateParams) WithData(data *models.WritableCluster) *VirtualizationClustersPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the virtualization clusters partial update params
func (o *VirtualizationClustersPartialUpdateParams) SetData(data *models.WritableCluster) {
	o.Data = data
}

// WithID adds the id to the virtualization clusters partial update params
func (o *VirtualizationClustersPartialUpdateParams) WithID(id strfmt.UUID) *VirtualizationClustersPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the virtualization clusters partial update params
func (o *VirtualizationClustersPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationClustersPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

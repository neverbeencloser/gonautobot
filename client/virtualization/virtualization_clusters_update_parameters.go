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

// NewVirtualizationClustersUpdateParams creates a new VirtualizationClustersUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVirtualizationClustersUpdateParams() *VirtualizationClustersUpdateParams {
	return &VirtualizationClustersUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationClustersUpdateParamsWithTimeout creates a new VirtualizationClustersUpdateParams object
// with the ability to set a timeout on a request.
func NewVirtualizationClustersUpdateParamsWithTimeout(timeout time.Duration) *VirtualizationClustersUpdateParams {
	return &VirtualizationClustersUpdateParams{
		timeout: timeout,
	}
}

// NewVirtualizationClustersUpdateParamsWithContext creates a new VirtualizationClustersUpdateParams object
// with the ability to set a context for a request.
func NewVirtualizationClustersUpdateParamsWithContext(ctx context.Context) *VirtualizationClustersUpdateParams {
	return &VirtualizationClustersUpdateParams{
		Context: ctx,
	}
}

// NewVirtualizationClustersUpdateParamsWithHTTPClient creates a new VirtualizationClustersUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewVirtualizationClustersUpdateParamsWithHTTPClient(client *http.Client) *VirtualizationClustersUpdateParams {
	return &VirtualizationClustersUpdateParams{
		HTTPClient: client,
	}
}

/* VirtualizationClustersUpdateParams contains all the parameters to send to the API endpoint
   for the virtualization clusters update operation.

   Typically these are written to a http.Request.
*/
type VirtualizationClustersUpdateParams struct {

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

// WithDefaults hydrates default values in the virtualization clusters update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationClustersUpdateParams) WithDefaults() *VirtualizationClustersUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the virtualization clusters update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationClustersUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the virtualization clusters update params
func (o *VirtualizationClustersUpdateParams) WithTimeout(timeout time.Duration) *VirtualizationClustersUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization clusters update params
func (o *VirtualizationClustersUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization clusters update params
func (o *VirtualizationClustersUpdateParams) WithContext(ctx context.Context) *VirtualizationClustersUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization clusters update params
func (o *VirtualizationClustersUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization clusters update params
func (o *VirtualizationClustersUpdateParams) WithHTTPClient(client *http.Client) *VirtualizationClustersUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization clusters update params
func (o *VirtualizationClustersUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the virtualization clusters update params
func (o *VirtualizationClustersUpdateParams) WithData(data *models.WritableCluster) *VirtualizationClustersUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the virtualization clusters update params
func (o *VirtualizationClustersUpdateParams) SetData(data *models.WritableCluster) {
	o.Data = data
}

// WithID adds the id to the virtualization clusters update params
func (o *VirtualizationClustersUpdateParams) WithID(id strfmt.UUID) *VirtualizationClustersUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the virtualization clusters update params
func (o *VirtualizationClustersUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationClustersUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

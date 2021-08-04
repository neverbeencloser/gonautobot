package plugins

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewPluginsDeviceOnboardingOnboardingListParams creates a new PluginsDeviceOnboardingOnboardingListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsDeviceOnboardingOnboardingListParams() *PluginsDeviceOnboardingOnboardingListParams {
	return &PluginsDeviceOnboardingOnboardingListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsDeviceOnboardingOnboardingListParamsWithTimeout creates a new PluginsDeviceOnboardingOnboardingListParams object
// with the ability to set a timeout on a request.
func NewPluginsDeviceOnboardingOnboardingListParamsWithTimeout(timeout time.Duration) *PluginsDeviceOnboardingOnboardingListParams {
	return &PluginsDeviceOnboardingOnboardingListParams{
		timeout: timeout,
	}
}

// NewPluginsDeviceOnboardingOnboardingListParamsWithContext creates a new PluginsDeviceOnboardingOnboardingListParams object
// with the ability to set a context for a request.
func NewPluginsDeviceOnboardingOnboardingListParamsWithContext(ctx context.Context) *PluginsDeviceOnboardingOnboardingListParams {
	return &PluginsDeviceOnboardingOnboardingListParams{
		Context: ctx,
	}
}

// NewPluginsDeviceOnboardingOnboardingListParamsWithHTTPClient creates a new PluginsDeviceOnboardingOnboardingListParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsDeviceOnboardingOnboardingListParamsWithHTTPClient(client *http.Client) *PluginsDeviceOnboardingOnboardingListParams {
	return &PluginsDeviceOnboardingOnboardingListParams{
		HTTPClient: client,
	}
}

/* PluginsDeviceOnboardingOnboardingListParams contains all the parameters to send to the API endpoint
   for the plugins device onboarding onboarding list operation.

   Typically these are written to a http.Request.
*/
type PluginsDeviceOnboardingOnboardingListParams struct {

	// FailedReason.
	FailedReason *string

	// ID.
	ID *string

	/* Limit.

	   Number of results to return per page.
	*/
	Limit *int64

	/* Offset.

	   The initial index from which to return the results.
	*/
	Offset *int64

	// Platform.
	Platform *string

	// Q.
	Q *string

	// Role.
	Role *string

	// Site.
	Site *string

	// SiteID.
	SiteID *string

	// Status.
	Status *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins device onboarding onboarding list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsDeviceOnboardingOnboardingListParams) WithDefaults() *PluginsDeviceOnboardingOnboardingListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins device onboarding onboarding list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsDeviceOnboardingOnboardingListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins device onboarding onboarding list params
func (o *PluginsDeviceOnboardingOnboardingListParams) WithTimeout(timeout time.Duration) *PluginsDeviceOnboardingOnboardingListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins device onboarding onboarding list params
func (o *PluginsDeviceOnboardingOnboardingListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins device onboarding onboarding list params
func (o *PluginsDeviceOnboardingOnboardingListParams) WithContext(ctx context.Context) *PluginsDeviceOnboardingOnboardingListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins device onboarding onboarding list params
func (o *PluginsDeviceOnboardingOnboardingListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins device onboarding onboarding list params
func (o *PluginsDeviceOnboardingOnboardingListParams) WithHTTPClient(client *http.Client) *PluginsDeviceOnboardingOnboardingListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins device onboarding onboarding list params
func (o *PluginsDeviceOnboardingOnboardingListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFailedReason adds the failedReason to the plugins device onboarding onboarding list params
func (o *PluginsDeviceOnboardingOnboardingListParams) WithFailedReason(failedReason *string) *PluginsDeviceOnboardingOnboardingListParams {
	o.SetFailedReason(failedReason)
	return o
}

// SetFailedReason adds the failedReason to the plugins device onboarding onboarding list params
func (o *PluginsDeviceOnboardingOnboardingListParams) SetFailedReason(failedReason *string) {
	o.FailedReason = failedReason
}

// WithID adds the id to the plugins device onboarding onboarding list params
func (o *PluginsDeviceOnboardingOnboardingListParams) WithID(id *string) *PluginsDeviceOnboardingOnboardingListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins device onboarding onboarding list params
func (o *PluginsDeviceOnboardingOnboardingListParams) SetID(id *string) {
	o.ID = id
}

// WithLimit adds the limit to the plugins device onboarding onboarding list params
func (o *PluginsDeviceOnboardingOnboardingListParams) WithLimit(limit *int64) *PluginsDeviceOnboardingOnboardingListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the plugins device onboarding onboarding list params
func (o *PluginsDeviceOnboardingOnboardingListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the plugins device onboarding onboarding list params
func (o *PluginsDeviceOnboardingOnboardingListParams) WithOffset(offset *int64) *PluginsDeviceOnboardingOnboardingListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the plugins device onboarding onboarding list params
func (o *PluginsDeviceOnboardingOnboardingListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPlatform adds the platform to the plugins device onboarding onboarding list params
func (o *PluginsDeviceOnboardingOnboardingListParams) WithPlatform(platform *string) *PluginsDeviceOnboardingOnboardingListParams {
	o.SetPlatform(platform)
	return o
}

// SetPlatform adds the platform to the plugins device onboarding onboarding list params
func (o *PluginsDeviceOnboardingOnboardingListParams) SetPlatform(platform *string) {
	o.Platform = platform
}

// WithQ adds the q to the plugins device onboarding onboarding list params
func (o *PluginsDeviceOnboardingOnboardingListParams) WithQ(q *string) *PluginsDeviceOnboardingOnboardingListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the plugins device onboarding onboarding list params
func (o *PluginsDeviceOnboardingOnboardingListParams) SetQ(q *string) {
	o.Q = q
}

// WithRole adds the role to the plugins device onboarding onboarding list params
func (o *PluginsDeviceOnboardingOnboardingListParams) WithRole(role *string) *PluginsDeviceOnboardingOnboardingListParams {
	o.SetRole(role)
	return o
}

// SetRole adds the role to the plugins device onboarding onboarding list params
func (o *PluginsDeviceOnboardingOnboardingListParams) SetRole(role *string) {
	o.Role = role
}

// WithSite adds the site to the plugins device onboarding onboarding list params
func (o *PluginsDeviceOnboardingOnboardingListParams) WithSite(site *string) *PluginsDeviceOnboardingOnboardingListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the plugins device onboarding onboarding list params
func (o *PluginsDeviceOnboardingOnboardingListParams) SetSite(site *string) {
	o.Site = site
}

// WithSiteID adds the siteID to the plugins device onboarding onboarding list params
func (o *PluginsDeviceOnboardingOnboardingListParams) WithSiteID(siteID *string) *PluginsDeviceOnboardingOnboardingListParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the plugins device onboarding onboarding list params
func (o *PluginsDeviceOnboardingOnboardingListParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WithStatus adds the status to the plugins device onboarding onboarding list params
func (o *PluginsDeviceOnboardingOnboardingListParams) WithStatus(status *string) *PluginsDeviceOnboardingOnboardingListParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the plugins device onboarding onboarding list params
func (o *PluginsDeviceOnboardingOnboardingListParams) SetStatus(status *string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsDeviceOnboardingOnboardingListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FailedReason != nil {

		// query param failed_reason
		var qrFailedReason string

		if o.FailedReason != nil {
			qrFailedReason = *o.FailedReason
		}
		qFailedReason := qrFailedReason
		if qFailedReason != "" {

			if err := r.SetQueryParam("failed_reason", qFailedReason); err != nil {
				return err
			}
		}
	}

	if o.ID != nil {

		// query param id
		var qrID string

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Platform != nil {

		// query param platform
		var qrPlatform string

		if o.Platform != nil {
			qrPlatform = *o.Platform
		}
		qPlatform := qrPlatform
		if qPlatform != "" {

			if err := r.SetQueryParam("platform", qPlatform); err != nil {
				return err
			}
		}
	}

	if o.Q != nil {

		// query param q
		var qrQ string

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}
	}

	if o.Role != nil {

		// query param role
		var qrRole string

		if o.Role != nil {
			qrRole = *o.Role
		}
		qRole := qrRole
		if qRole != "" {

			if err := r.SetQueryParam("role", qRole); err != nil {
				return err
			}
		}
	}

	if o.Site != nil {

		// query param site
		var qrSite string

		if o.Site != nil {
			qrSite = *o.Site
		}
		qSite := qrSite
		if qSite != "" {

			if err := r.SetQueryParam("site", qSite); err != nil {
				return err
			}
		}
	}

	if o.SiteID != nil {

		// query param site_id
		var qrSiteID string

		if o.SiteID != nil {
			qrSiteID = *o.SiteID
		}
		qSiteID := qrSiteID
		if qSiteID != "" {

			if err := r.SetQueryParam("site_id", qSiteID); err != nil {
				return err
			}
		}
	}

	if o.Status != nil {

		// query param status
		var qrStatus string

		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {

			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

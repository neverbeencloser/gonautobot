package users

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

// NewUsersUsersPartialUpdateParams creates a new UsersUsersPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsersUsersPartialUpdateParams() *UsersUsersPartialUpdateParams {
	return &UsersUsersPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsersUsersPartialUpdateParamsWithTimeout creates a new UsersUsersPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewUsersUsersPartialUpdateParamsWithTimeout(timeout time.Duration) *UsersUsersPartialUpdateParams {
	return &UsersUsersPartialUpdateParams{
		timeout: timeout,
	}
}

// NewUsersUsersPartialUpdateParamsWithContext creates a new UsersUsersPartialUpdateParams object
// with the ability to set a context for a request.
func NewUsersUsersPartialUpdateParamsWithContext(ctx context.Context) *UsersUsersPartialUpdateParams {
	return &UsersUsersPartialUpdateParams{
		Context: ctx,
	}
}

// NewUsersUsersPartialUpdateParamsWithHTTPClient creates a new UsersUsersPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewUsersUsersPartialUpdateParamsWithHTTPClient(client *http.Client) *UsersUsersPartialUpdateParams {
	return &UsersUsersPartialUpdateParams{
		HTTPClient: client,
	}
}

/* UsersUsersPartialUpdateParams contains all the parameters to send to the API endpoint
   for the users users partial update operation.

   Typically these are written to a http.Request.
*/
type UsersUsersPartialUpdateParams struct {

	// Data.
	Data *models.WritableUser

	/* ID.

	   A UUID string identifying this user.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the users users partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersUsersPartialUpdateParams) WithDefaults() *UsersUsersPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the users users partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersUsersPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the users users partial update params
func (o *UsersUsersPartialUpdateParams) WithTimeout(timeout time.Duration) *UsersUsersPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users users partial update params
func (o *UsersUsersPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users users partial update params
func (o *UsersUsersPartialUpdateParams) WithContext(ctx context.Context) *UsersUsersPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users users partial update params
func (o *UsersUsersPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users users partial update params
func (o *UsersUsersPartialUpdateParams) WithHTTPClient(client *http.Client) *UsersUsersPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users users partial update params
func (o *UsersUsersPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the users users partial update params
func (o *UsersUsersPartialUpdateParams) WithData(data *models.WritableUser) *UsersUsersPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the users users partial update params
func (o *UsersUsersPartialUpdateParams) SetData(data *models.WritableUser) {
	o.Data = data
}

// WithID adds the id to the users users partial update params
func (o *UsersUsersPartialUpdateParams) WithID(id strfmt.UUID) *UsersUsersPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the users users partial update params
func (o *UsersUsersPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UsersUsersPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

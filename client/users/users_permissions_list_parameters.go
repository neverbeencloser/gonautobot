package users

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

// NewUsersPermissionsListParams creates a new UsersPermissionsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsersPermissionsListParams() *UsersPermissionsListParams {
	return &UsersPermissionsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsersPermissionsListParamsWithTimeout creates a new UsersPermissionsListParams object
// with the ability to set a timeout on a request.
func NewUsersPermissionsListParamsWithTimeout(timeout time.Duration) *UsersPermissionsListParams {
	return &UsersPermissionsListParams{
		timeout: timeout,
	}
}

// NewUsersPermissionsListParamsWithContext creates a new UsersPermissionsListParams object
// with the ability to set a context for a request.
func NewUsersPermissionsListParamsWithContext(ctx context.Context) *UsersPermissionsListParams {
	return &UsersPermissionsListParams{
		Context: ctx,
	}
}

// NewUsersPermissionsListParamsWithHTTPClient creates a new UsersPermissionsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewUsersPermissionsListParamsWithHTTPClient(client *http.Client) *UsersPermissionsListParams {
	return &UsersPermissionsListParams{
		HTTPClient: client,
	}
}

/* UsersPermissionsListParams contains all the parameters to send to the API endpoint
   for the users permissions list operation.

   Typically these are written to a http.Request.
*/
type UsersPermissionsListParams struct {

	// Enabled.
	Enabled *string

	// Group.
	Group *string

	// Groupn.
	Groupn *string

	// GroupID.
	GroupID *string

	// GroupIDn.
	GroupIDn *string

	// ID.
	ID *string

	// IDIc.
	IDIc *string

	// IDIe.
	IDIe *string

	// IDIew.
	IDIew *string

	// IDIsw.
	IDIsw *string

	// IDn.
	IDn *string

	// IDNic.
	IDNic *string

	// IDNie.
	IDNie *string

	// IDNiew.
	IDNiew *string

	// IDNisw.
	IDNisw *string

	/* Limit.

	   Number of results to return per page.
	*/
	Limit *int64

	// Name.
	Name *string

	// NameIc.
	NameIc *string

	// NameIe.
	NameIe *string

	// NameIew.
	NameIew *string

	// NameIsw.
	NameIsw *string

	// Namen.
	Namen *string

	// NameNic.
	NameNic *string

	// NameNie.
	NameNie *string

	// NameNiew.
	NameNiew *string

	// NameNisw.
	NameNisw *string

	// ObjectTypes.
	ObjectTypes *string

	// ObjectTypesn.
	ObjectTypesn *string

	/* Offset.

	   The initial index from which to return the results.
	*/
	Offset *int64

	// User.
	User *string

	// Usern.
	Usern *string

	// UserID.
	UserID *string

	// UserIDn.
	UserIDn *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the users permissions list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersPermissionsListParams) WithDefaults() *UsersPermissionsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the users permissions list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersPermissionsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the users permissions list params
func (o *UsersPermissionsListParams) WithTimeout(timeout time.Duration) *UsersPermissionsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users permissions list params
func (o *UsersPermissionsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users permissions list params
func (o *UsersPermissionsListParams) WithContext(ctx context.Context) *UsersPermissionsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users permissions list params
func (o *UsersPermissionsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users permissions list params
func (o *UsersPermissionsListParams) WithHTTPClient(client *http.Client) *UsersPermissionsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users permissions list params
func (o *UsersPermissionsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnabled adds the enabled to the users permissions list params
func (o *UsersPermissionsListParams) WithEnabled(enabled *string) *UsersPermissionsListParams {
	o.SetEnabled(enabled)
	return o
}

// SetEnabled adds the enabled to the users permissions list params
func (o *UsersPermissionsListParams) SetEnabled(enabled *string) {
	o.Enabled = enabled
}

// WithGroup adds the group to the users permissions list params
func (o *UsersPermissionsListParams) WithGroup(group *string) *UsersPermissionsListParams {
	o.SetGroup(group)
	return o
}

// SetGroup adds the group to the users permissions list params
func (o *UsersPermissionsListParams) SetGroup(group *string) {
	o.Group = group
}

// WithGroupn adds the groupn to the users permissions list params
func (o *UsersPermissionsListParams) WithGroupn(groupn *string) *UsersPermissionsListParams {
	o.SetGroupn(groupn)
	return o
}

// SetGroupn adds the groupN to the users permissions list params
func (o *UsersPermissionsListParams) SetGroupn(groupn *string) {
	o.Groupn = groupn
}

// WithGroupID adds the groupID to the users permissions list params
func (o *UsersPermissionsListParams) WithGroupID(groupID *string) *UsersPermissionsListParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the users permissions list params
func (o *UsersPermissionsListParams) SetGroupID(groupID *string) {
	o.GroupID = groupID
}

// WithGroupIDn adds the groupIDn to the users permissions list params
func (o *UsersPermissionsListParams) WithGroupIDn(groupIDn *string) *UsersPermissionsListParams {
	o.SetGroupIDn(groupIDn)
	return o
}

// SetGroupIDn adds the groupIdN to the users permissions list params
func (o *UsersPermissionsListParams) SetGroupIDn(groupIDn *string) {
	o.GroupIDn = groupIDn
}

// WithID adds the id to the users permissions list params
func (o *UsersPermissionsListParams) WithID(id *string) *UsersPermissionsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the users permissions list params
func (o *UsersPermissionsListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the users permissions list params
func (o *UsersPermissionsListParams) WithIDIc(iDIc *string) *UsersPermissionsListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the users permissions list params
func (o *UsersPermissionsListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the users permissions list params
func (o *UsersPermissionsListParams) WithIDIe(iDIe *string) *UsersPermissionsListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the users permissions list params
func (o *UsersPermissionsListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the users permissions list params
func (o *UsersPermissionsListParams) WithIDIew(iDIew *string) *UsersPermissionsListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the users permissions list params
func (o *UsersPermissionsListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the users permissions list params
func (o *UsersPermissionsListParams) WithIDIsw(iDIsw *string) *UsersPermissionsListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the users permissions list params
func (o *UsersPermissionsListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the users permissions list params
func (o *UsersPermissionsListParams) WithIDn(iDn *string) *UsersPermissionsListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the users permissions list params
func (o *UsersPermissionsListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the users permissions list params
func (o *UsersPermissionsListParams) WithIDNic(iDNic *string) *UsersPermissionsListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the users permissions list params
func (o *UsersPermissionsListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the users permissions list params
func (o *UsersPermissionsListParams) WithIDNie(iDNie *string) *UsersPermissionsListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the users permissions list params
func (o *UsersPermissionsListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the users permissions list params
func (o *UsersPermissionsListParams) WithIDNiew(iDNiew *string) *UsersPermissionsListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the users permissions list params
func (o *UsersPermissionsListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the users permissions list params
func (o *UsersPermissionsListParams) WithIDNisw(iDNisw *string) *UsersPermissionsListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the users permissions list params
func (o *UsersPermissionsListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithLimit adds the limit to the users permissions list params
func (o *UsersPermissionsListParams) WithLimit(limit *int64) *UsersPermissionsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the users permissions list params
func (o *UsersPermissionsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the users permissions list params
func (o *UsersPermissionsListParams) WithName(name *string) *UsersPermissionsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the users permissions list params
func (o *UsersPermissionsListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the users permissions list params
func (o *UsersPermissionsListParams) WithNameIc(nameIc *string) *UsersPermissionsListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the users permissions list params
func (o *UsersPermissionsListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the users permissions list params
func (o *UsersPermissionsListParams) WithNameIe(nameIe *string) *UsersPermissionsListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the users permissions list params
func (o *UsersPermissionsListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the users permissions list params
func (o *UsersPermissionsListParams) WithNameIew(nameIew *string) *UsersPermissionsListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the users permissions list params
func (o *UsersPermissionsListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the users permissions list params
func (o *UsersPermissionsListParams) WithNameIsw(nameIsw *string) *UsersPermissionsListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the users permissions list params
func (o *UsersPermissionsListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the users permissions list params
func (o *UsersPermissionsListParams) WithNamen(namen *string) *UsersPermissionsListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the users permissions list params
func (o *UsersPermissionsListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the users permissions list params
func (o *UsersPermissionsListParams) WithNameNic(nameNic *string) *UsersPermissionsListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the users permissions list params
func (o *UsersPermissionsListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the users permissions list params
func (o *UsersPermissionsListParams) WithNameNie(nameNie *string) *UsersPermissionsListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the users permissions list params
func (o *UsersPermissionsListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the users permissions list params
func (o *UsersPermissionsListParams) WithNameNiew(nameNiew *string) *UsersPermissionsListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the users permissions list params
func (o *UsersPermissionsListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the users permissions list params
func (o *UsersPermissionsListParams) WithNameNisw(nameNisw *string) *UsersPermissionsListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the users permissions list params
func (o *UsersPermissionsListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithObjectTypes adds the objectTypes to the users permissions list params
func (o *UsersPermissionsListParams) WithObjectTypes(objectTypes *string) *UsersPermissionsListParams {
	o.SetObjectTypes(objectTypes)
	return o
}

// SetObjectTypes adds the objectTypes to the users permissions list params
func (o *UsersPermissionsListParams) SetObjectTypes(objectTypes *string) {
	o.ObjectTypes = objectTypes
}

// WithObjectTypesn adds the objectTypesn to the users permissions list params
func (o *UsersPermissionsListParams) WithObjectTypesn(objectTypesn *string) *UsersPermissionsListParams {
	o.SetObjectTypesn(objectTypesn)
	return o
}

// SetObjectTypesn adds the objectTypesN to the users permissions list params
func (o *UsersPermissionsListParams) SetObjectTypesn(objectTypesn *string) {
	o.ObjectTypesn = objectTypesn
}

// WithOffset adds the offset to the users permissions list params
func (o *UsersPermissionsListParams) WithOffset(offset *int64) *UsersPermissionsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the users permissions list params
func (o *UsersPermissionsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithUser adds the user to the users permissions list params
func (o *UsersPermissionsListParams) WithUser(user *string) *UsersPermissionsListParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the users permissions list params
func (o *UsersPermissionsListParams) SetUser(user *string) {
	o.User = user
}

// WithUsern adds the usern to the users permissions list params
func (o *UsersPermissionsListParams) WithUsern(usern *string) *UsersPermissionsListParams {
	o.SetUsern(usern)
	return o
}

// SetUsern adds the userN to the users permissions list params
func (o *UsersPermissionsListParams) SetUsern(usern *string) {
	o.Usern = usern
}

// WithUserID adds the userID to the users permissions list params
func (o *UsersPermissionsListParams) WithUserID(userID *string) *UsersPermissionsListParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the users permissions list params
func (o *UsersPermissionsListParams) SetUserID(userID *string) {
	o.UserID = userID
}

// WithUserIDn adds the userIDn to the users permissions list params
func (o *UsersPermissionsListParams) WithUserIDn(userIDn *string) *UsersPermissionsListParams {
	o.SetUserIDn(userIDn)
	return o
}

// SetUserIDn adds the userIdN to the users permissions list params
func (o *UsersPermissionsListParams) SetUserIDn(userIDn *string) {
	o.UserIDn = userIDn
}

// WriteToRequest writes these params to a swagger request
func (o *UsersPermissionsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Enabled != nil {

		// query param enabled
		var qrEnabled string

		if o.Enabled != nil {
			qrEnabled = *o.Enabled
		}
		qEnabled := qrEnabled
		if qEnabled != "" {

			if err := r.SetQueryParam("enabled", qEnabled); err != nil {
				return err
			}
		}
	}

	if o.Group != nil {

		// query param group
		var qrGroup string

		if o.Group != nil {
			qrGroup = *o.Group
		}
		qGroup := qrGroup
		if qGroup != "" {

			if err := r.SetQueryParam("group", qGroup); err != nil {
				return err
			}
		}
	}

	if o.Groupn != nil {

		// query param group__n
		var qrGroupn string

		if o.Groupn != nil {
			qrGroupn = *o.Groupn
		}
		qGroupn := qrGroupn
		if qGroupn != "" {

			if err := r.SetQueryParam("group__n", qGroupn); err != nil {
				return err
			}
		}
	}

	if o.GroupID != nil {

		// query param group_id
		var qrGroupID string

		if o.GroupID != nil {
			qrGroupID = *o.GroupID
		}
		qGroupID := qrGroupID
		if qGroupID != "" {

			if err := r.SetQueryParam("group_id", qGroupID); err != nil {
				return err
			}
		}
	}

	if o.GroupIDn != nil {

		// query param group_id__n
		var qrGroupIDn string

		if o.GroupIDn != nil {
			qrGroupIDn = *o.GroupIDn
		}
		qGroupIDn := qrGroupIDn
		if qGroupIDn != "" {

			if err := r.SetQueryParam("group_id__n", qGroupIDn); err != nil {
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

	if o.IDIc != nil {

		// query param id__ic
		var qrIDIc string

		if o.IDIc != nil {
			qrIDIc = *o.IDIc
		}
		qIDIc := qrIDIc
		if qIDIc != "" {

			if err := r.SetQueryParam("id__ic", qIDIc); err != nil {
				return err
			}
		}
	}

	if o.IDIe != nil {

		// query param id__ie
		var qrIDIe string

		if o.IDIe != nil {
			qrIDIe = *o.IDIe
		}
		qIDIe := qrIDIe
		if qIDIe != "" {

			if err := r.SetQueryParam("id__ie", qIDIe); err != nil {
				return err
			}
		}
	}

	if o.IDIew != nil {

		// query param id__iew
		var qrIDIew string

		if o.IDIew != nil {
			qrIDIew = *o.IDIew
		}
		qIDIew := qrIDIew
		if qIDIew != "" {

			if err := r.SetQueryParam("id__iew", qIDIew); err != nil {
				return err
			}
		}
	}

	if o.IDIsw != nil {

		// query param id__isw
		var qrIDIsw string

		if o.IDIsw != nil {
			qrIDIsw = *o.IDIsw
		}
		qIDIsw := qrIDIsw
		if qIDIsw != "" {

			if err := r.SetQueryParam("id__isw", qIDIsw); err != nil {
				return err
			}
		}
	}

	if o.IDn != nil {

		// query param id__n
		var qrIDn string

		if o.IDn != nil {
			qrIDn = *o.IDn
		}
		qIDn := qrIDn
		if qIDn != "" {

			if err := r.SetQueryParam("id__n", qIDn); err != nil {
				return err
			}
		}
	}

	if o.IDNic != nil {

		// query param id__nic
		var qrIDNic string

		if o.IDNic != nil {
			qrIDNic = *o.IDNic
		}
		qIDNic := qrIDNic
		if qIDNic != "" {

			if err := r.SetQueryParam("id__nic", qIDNic); err != nil {
				return err
			}
		}
	}

	if o.IDNie != nil {

		// query param id__nie
		var qrIDNie string

		if o.IDNie != nil {
			qrIDNie = *o.IDNie
		}
		qIDNie := qrIDNie
		if qIDNie != "" {

			if err := r.SetQueryParam("id__nie", qIDNie); err != nil {
				return err
			}
		}
	}

	if o.IDNiew != nil {

		// query param id__niew
		var qrIDNiew string

		if o.IDNiew != nil {
			qrIDNiew = *o.IDNiew
		}
		qIDNiew := qrIDNiew
		if qIDNiew != "" {

			if err := r.SetQueryParam("id__niew", qIDNiew); err != nil {
				return err
			}
		}
	}

	if o.IDNisw != nil {

		// query param id__nisw
		var qrIDNisw string

		if o.IDNisw != nil {
			qrIDNisw = *o.IDNisw
		}
		qIDNisw := qrIDNisw
		if qIDNisw != "" {

			if err := r.SetQueryParam("id__nisw", qIDNisw); err != nil {
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

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.NameIc != nil {

		// query param name__ic
		var qrNameIc string

		if o.NameIc != nil {
			qrNameIc = *o.NameIc
		}
		qNameIc := qrNameIc
		if qNameIc != "" {

			if err := r.SetQueryParam("name__ic", qNameIc); err != nil {
				return err
			}
		}
	}

	if o.NameIe != nil {

		// query param name__ie
		var qrNameIe string

		if o.NameIe != nil {
			qrNameIe = *o.NameIe
		}
		qNameIe := qrNameIe
		if qNameIe != "" {

			if err := r.SetQueryParam("name__ie", qNameIe); err != nil {
				return err
			}
		}
	}

	if o.NameIew != nil {

		// query param name__iew
		var qrNameIew string

		if o.NameIew != nil {
			qrNameIew = *o.NameIew
		}
		qNameIew := qrNameIew
		if qNameIew != "" {

			if err := r.SetQueryParam("name__iew", qNameIew); err != nil {
				return err
			}
		}
	}

	if o.NameIsw != nil {

		// query param name__isw
		var qrNameIsw string

		if o.NameIsw != nil {
			qrNameIsw = *o.NameIsw
		}
		qNameIsw := qrNameIsw
		if qNameIsw != "" {

			if err := r.SetQueryParam("name__isw", qNameIsw); err != nil {
				return err
			}
		}
	}

	if o.Namen != nil {

		// query param name__n
		var qrNamen string

		if o.Namen != nil {
			qrNamen = *o.Namen
		}
		qNamen := qrNamen
		if qNamen != "" {

			if err := r.SetQueryParam("name__n", qNamen); err != nil {
				return err
			}
		}
	}

	if o.NameNic != nil {

		// query param name__nic
		var qrNameNic string

		if o.NameNic != nil {
			qrNameNic = *o.NameNic
		}
		qNameNic := qrNameNic
		if qNameNic != "" {

			if err := r.SetQueryParam("name__nic", qNameNic); err != nil {
				return err
			}
		}
	}

	if o.NameNie != nil {

		// query param name__nie
		var qrNameNie string

		if o.NameNie != nil {
			qrNameNie = *o.NameNie
		}
		qNameNie := qrNameNie
		if qNameNie != "" {

			if err := r.SetQueryParam("name__nie", qNameNie); err != nil {
				return err
			}
		}
	}

	if o.NameNiew != nil {

		// query param name__niew
		var qrNameNiew string

		if o.NameNiew != nil {
			qrNameNiew = *o.NameNiew
		}
		qNameNiew := qrNameNiew
		if qNameNiew != "" {

			if err := r.SetQueryParam("name__niew", qNameNiew); err != nil {
				return err
			}
		}
	}

	if o.NameNisw != nil {

		// query param name__nisw
		var qrNameNisw string

		if o.NameNisw != nil {
			qrNameNisw = *o.NameNisw
		}
		qNameNisw := qrNameNisw
		if qNameNisw != "" {

			if err := r.SetQueryParam("name__nisw", qNameNisw); err != nil {
				return err
			}
		}
	}

	if o.ObjectTypes != nil {

		// query param object_types
		var qrObjectTypes string

		if o.ObjectTypes != nil {
			qrObjectTypes = *o.ObjectTypes
		}
		qObjectTypes := qrObjectTypes
		if qObjectTypes != "" {

			if err := r.SetQueryParam("object_types", qObjectTypes); err != nil {
				return err
			}
		}
	}

	if o.ObjectTypesn != nil {

		// query param object_types__n
		var qrObjectTypesn string

		if o.ObjectTypesn != nil {
			qrObjectTypesn = *o.ObjectTypesn
		}
		qObjectTypesn := qrObjectTypesn
		if qObjectTypesn != "" {

			if err := r.SetQueryParam("object_types__n", qObjectTypesn); err != nil {
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

	if o.User != nil {

		// query param user
		var qrUser string

		if o.User != nil {
			qrUser = *o.User
		}
		qUser := qrUser
		if qUser != "" {

			if err := r.SetQueryParam("user", qUser); err != nil {
				return err
			}
		}
	}

	if o.Usern != nil {

		// query param user__n
		var qrUsern string

		if o.Usern != nil {
			qrUsern = *o.Usern
		}
		qUsern := qrUsern
		if qUsern != "" {

			if err := r.SetQueryParam("user__n", qUsern); err != nil {
				return err
			}
		}
	}

	if o.UserID != nil {

		// query param user_id
		var qrUserID string

		if o.UserID != nil {
			qrUserID = *o.UserID
		}
		qUserID := qrUserID
		if qUserID != "" {

			if err := r.SetQueryParam("user_id", qUserID); err != nil {
				return err
			}
		}
	}

	if o.UserIDn != nil {

		// query param user_id__n
		var qrUserIDn string

		if o.UserIDn != nil {
			qrUserIDn = *o.UserIDn
		}
		qUserIDn := qrUserIDn
		if qUserIDn != "" {

			if err := r.SetQueryParam("user_id__n", qUserIDn); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

package models

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableIPAddress writable IP address
//
// swagger:model WritableIPAddress
type WritableIPAddress struct {

	// Address
	// Required: true
	// Min Length: 1
	Address *string `json:"address"`

	// Assigned object
	// Read Only: true
	AssignedObject map[string]*string `json:"assigned_object,omitempty"`

	// Assigned object id
	// Format: uuid
	AssignedObjectID *strfmt.UUID `json:"assigned_object_id,omitempty"`

	// Assigned object type
	AssignedObjectType *string `json:"assigned_object_type,omitempty"`

	// Computed fields
	// Read Only: true
	ComputedFields string `json:"computed_fields,omitempty"`

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// Display
	//
	// Human friendly display value
	// Read Only: true
	// Min Length: 1
	Display string `json:"display,omitempty"`

	// DNS Name
	//
	// Hostname or FQDN (not case-sensitive)
	// Max Length: 255
	// Pattern: ^[0-9A-Za-z._-]+$
	DNSName string `json:"dns_name,omitempty"`

	// Family
	// Read Only: true
	Family string `json:"family,omitempty"`

	// Id
	// Read Only: true
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// NAT (Inside)
	//
	// The IP for which this address is the "outside" IP
	// Format: uuid
	NatInside *strfmt.UUID `json:"nat_inside,omitempty"`

	// Nat outside
	// Required: true
	// Format: uuid
	NatOutside *strfmt.UUID `json:"nat_outside"`

	// Role
	//
	// The functional role of this IP
	// Enum: [loopback secondary anycast vip vrrp hsrp glbp carp]
	Role string `json:"role,omitempty"`

	// Status
	// Pattern: ^[-a-zA-Z0-9_]+$
	// Enum: [active deprecated dhcp reserved slaac]
	Status string `json:"status,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// Tenant
	// Format: uuid
	Tenant *strfmt.UUID `json:"tenant,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// VRF
	// Format: uuid
	Vrf *strfmt.UUID `json:"vrf,omitempty"`
}

// Validate validates this writable IP address
func (m *WritableIPAddress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssignedObjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDNSName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNatInside(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNatOutside(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenant(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVrf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableIPAddress) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("address", "body", m.Address); err != nil {
		return err
	}

	if err := validate.MinLength("address", "body", *m.Address, 1); err != nil {
		return err
	}

	return nil
}

func (m *WritableIPAddress) validateAssignedObjectID(formats strfmt.Registry) error {
	if swag.IsZero(m.AssignedObjectID) { // not required
		return nil
	}

	if err := validate.FormatOf("assigned_object_id", "body", "uuid", m.AssignedObjectID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableIPAddress) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableIPAddress) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *WritableIPAddress) validateDisplay(formats strfmt.Registry) error {
	if swag.IsZero(m.Display) { // not required
		return nil
	}

	if err := validate.MinLength("display", "body", m.Display, 1); err != nil {
		return err
	}

	return nil
}

func (m *WritableIPAddress) validateDNSName(formats strfmt.Registry) error {
	if swag.IsZero(m.DNSName) { // not required
		return nil
	}

	if err := validate.MaxLength("dns_name", "body", m.DNSName, 255); err != nil {
		return err
	}

	if err := validate.Pattern("dns_name", "body", m.DNSName, `^[0-9A-Za-z._-]+$`); err != nil {
		return err
	}

	return nil
}

func (m *WritableIPAddress) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableIPAddress) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableIPAddress) validateNatInside(formats strfmt.Registry) error {
	if swag.IsZero(m.NatInside) { // not required
		return nil
	}

	if err := validate.FormatOf("nat_inside", "body", "uuid", m.NatInside.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableIPAddress) validateNatOutside(formats strfmt.Registry) error {

	if err := validate.Required("nat_outside", "body", m.NatOutside); err != nil {
		return err
	}

	if err := validate.FormatOf("nat_outside", "body", "uuid", m.NatOutside.String(), formats); err != nil {
		return err
	}

	return nil
}

var writableIpAddressTypeRolePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["loopback","secondary","anycast","vip","vrrp","hsrp","glbp","carp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableIpAddressTypeRolePropEnum = append(writableIpAddressTypeRolePropEnum, v)
	}
}

const (

	// WritableIPAddressRoleLoopback captures enum value "loopback"
	WritableIPAddressRoleLoopback string = "loopback"

	// WritableIPAddressRoleSecondary captures enum value "secondary"
	WritableIPAddressRoleSecondary string = "secondary"

	// WritableIPAddressRoleAnycast captures enum value "anycast"
	WritableIPAddressRoleAnycast string = "anycast"

	// WritableIPAddressRoleVip captures enum value "vip"
	WritableIPAddressRoleVip string = "vip"

	// WritableIPAddressRoleVrrp captures enum value "vrrp"
	WritableIPAddressRoleVrrp string = "vrrp"

	// WritableIPAddressRoleHsrp captures enum value "hsrp"
	WritableIPAddressRoleHsrp string = "hsrp"

	// WritableIPAddressRoleGlbp captures enum value "glbp"
	WritableIPAddressRoleGlbp string = "glbp"

	// WritableIPAddressRoleCarp captures enum value "carp"
	WritableIPAddressRoleCarp string = "carp"
)

// prop value enum
func (m *WritableIPAddress) validateRoleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableIpAddressTypeRolePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableIPAddress) validateRole(formats strfmt.Registry) error {
	if swag.IsZero(m.Role) { // not required
		return nil
	}

	// value enum
	if err := m.validateRoleEnum("role", "body", m.Role); err != nil {
		return err
	}

	return nil
}

var writableIpAddressTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","deprecated","dhcp","reserved","slaac"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableIpAddressTypeStatusPropEnum = append(writableIpAddressTypeStatusPropEnum, v)
	}
}

const (

	// WritableIPAddressStatusActive captures enum value "active"
	WritableIPAddressStatusActive string = "active"

	// WritableIPAddressStatusDeprecated captures enum value "deprecated"
	WritableIPAddressStatusDeprecated string = "deprecated"

	// WritableIPAddressStatusDhcp captures enum value "dhcp"
	WritableIPAddressStatusDhcp string = "dhcp"

	// WritableIPAddressStatusReserved captures enum value "reserved"
	WritableIPAddressStatusReserved string = "reserved"

	// WritableIPAddressStatusSlaac captures enum value "slaac"
	WritableIPAddressStatusSlaac string = "slaac"
)

// prop value enum
func (m *WritableIPAddress) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableIpAddressTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableIPAddress) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := validate.Pattern("status", "body", m.Status, `^[-a-zA-Z0-9_]+$`); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *WritableIPAddress) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WritableIPAddress) validateTenant(formats strfmt.Registry) error {
	if swag.IsZero(m.Tenant) { // not required
		return nil
	}

	if err := validate.FormatOf("tenant", "body", "uuid", m.Tenant.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableIPAddress) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableIPAddress) validateVrf(formats strfmt.Registry) error {
	if swag.IsZero(m.Vrf) { // not required
		return nil
	}

	if err := validate.FormatOf("vrf", "body", "uuid", m.Vrf.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this writable IP address based on the context it is used
func (m *WritableIPAddress) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAssignedObject(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateComputedFields(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFamily(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableIPAddress) contextValidateAssignedObject(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *WritableIPAddress) contextValidateComputedFields(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "computed_fields", "body", string(m.ComputedFields)); err != nil {
		return err
	}

	return nil
}

func (m *WritableIPAddress) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", strfmt.Date(m.Created)); err != nil {
		return err
	}

	return nil
}

func (m *WritableIPAddress) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *WritableIPAddress) contextValidateFamily(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "family", "body", string(m.Family)); err != nil {
		return err
	}

	return nil
}

func (m *WritableIPAddress) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *WritableIPAddress) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", strfmt.DateTime(m.LastUpdated)); err != nil {
		return err
	}

	return nil
}

func (m *WritableIPAddress) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {
			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WritableIPAddress) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableIPAddress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableIPAddress) UnmarshalBinary(b []byte) error {
	var res WritableIPAddress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

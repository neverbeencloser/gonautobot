package models

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConfigContext config context
//
// swagger:model ConfigContext
type ConfigContext struct {

	// cluster groups
	// Unique: true
	ClusterGroups []strfmt.UUID `json:"cluster_groups"`

	// clusters
	// Unique: true
	Clusters []strfmt.UUID `json:"clusters"`

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Data
	// Required: true
	Data *string `json:"data"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// device types
	// Unique: true
	DeviceTypes []strfmt.UUID `json:"device_types"`

	// Display
	//
	// Human friendly display value
	// Read Only: true
	// Min Length: 1
	Display string `json:"display,omitempty"`

	// Id
	// Read Only: true
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Is active
	IsActive bool `json:"is_active,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Name
	// Required: true
	// Max Length: 100
	// Min Length: 1
	Name *string `json:"name"`

	// Owner
	// Read Only: true
	Owner map[string]*string `json:"owner,omitempty"`

	// Owner content type
	OwnerContentType *string `json:"owner_content_type,omitempty"`

	// Owner object id
	// Format: uuid
	OwnerObjectID *strfmt.UUID `json:"owner_object_id,omitempty"`

	// platforms
	// Unique: true
	Platforms []strfmt.UUID `json:"platforms"`

	// regions
	// Unique: true
	Regions []strfmt.UUID `json:"regions"`

	// roles
	// Unique: true
	Roles []strfmt.UUID `json:"roles"`

	// schema
	Schema *NestedConfigContextSchema `json:"schema,omitempty"`

	// sites
	// Unique: true
	Sites []strfmt.UUID `json:"sites"`

	// tags
	// Unique: true
	Tags []string `json:"tags"`

	// tenant groups
	// Unique: true
	TenantGroups []strfmt.UUID `json:"tenant_groups"`

	// tenants
	// Unique: true
	Tenants []strfmt.UUID `json:"tenants"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// Weight
	// Maximum: 32767
	// Minimum: 0
	Weight *int64 `json:"weight,omitempty"`
}

// Validate validates this config context
func (m *ConfigContext) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwnerObjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatforms(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchema(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSites(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenantGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenants(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeight(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigContext) validateClusterGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterGroups) { // not required
		return nil
	}

	if err := validate.UniqueItems("cluster_groups", "body", m.ClusterGroups); err != nil {
		return err
	}

	for i := 0; i < len(m.ClusterGroups); i++ {

		if err := validate.FormatOf("cluster_groups"+"."+strconv.Itoa(i), "body", "uuid", m.ClusterGroups[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *ConfigContext) validateClusters(formats strfmt.Registry) error {
	if swag.IsZero(m.Clusters) { // not required
		return nil
	}

	if err := validate.UniqueItems("clusters", "body", m.Clusters); err != nil {
		return err
	}

	for i := 0; i < len(m.Clusters); i++ {

		if err := validate.FormatOf("clusters"+"."+strconv.Itoa(i), "body", "uuid", m.Clusters[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *ConfigContext) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConfigContext) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
	}

	return nil
}

func (m *ConfigContext) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *ConfigContext) validateDeviceTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.DeviceTypes) { // not required
		return nil
	}

	if err := validate.UniqueItems("device_types", "body", m.DeviceTypes); err != nil {
		return err
	}

	for i := 0; i < len(m.DeviceTypes); i++ {

		if err := validate.FormatOf("device_types"+"."+strconv.Itoa(i), "body", "uuid", m.DeviceTypes[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *ConfigContext) validateDisplay(formats strfmt.Registry) error {
	if swag.IsZero(m.Display) { // not required
		return nil
	}

	if err := validate.MinLength("display", "body", m.Display, 1); err != nil {
		return err
	}

	return nil
}

func (m *ConfigContext) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConfigContext) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConfigContext) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 100); err != nil {
		return err
	}

	return nil
}

func (m *ConfigContext) validateOwnerObjectID(formats strfmt.Registry) error {
	if swag.IsZero(m.OwnerObjectID) { // not required
		return nil
	}

	if err := validate.FormatOf("owner_object_id", "body", "uuid", m.OwnerObjectID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConfigContext) validatePlatforms(formats strfmt.Registry) error {
	if swag.IsZero(m.Platforms) { // not required
		return nil
	}

	if err := validate.UniqueItems("platforms", "body", m.Platforms); err != nil {
		return err
	}

	for i := 0; i < len(m.Platforms); i++ {

		if err := validate.FormatOf("platforms"+"."+strconv.Itoa(i), "body", "uuid", m.Platforms[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *ConfigContext) validateRegions(formats strfmt.Registry) error {
	if swag.IsZero(m.Regions) { // not required
		return nil
	}

	if err := validate.UniqueItems("regions", "body", m.Regions); err != nil {
		return err
	}

	for i := 0; i < len(m.Regions); i++ {

		if err := validate.FormatOf("regions"+"."+strconv.Itoa(i), "body", "uuid", m.Regions[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *ConfigContext) validateRoles(formats strfmt.Registry) error {
	if swag.IsZero(m.Roles) { // not required
		return nil
	}

	if err := validate.UniqueItems("roles", "body", m.Roles); err != nil {
		return err
	}

	for i := 0; i < len(m.Roles); i++ {

		if err := validate.FormatOf("roles"+"."+strconv.Itoa(i), "body", "uuid", m.Roles[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *ConfigContext) validateSchema(formats strfmt.Registry) error {
	if swag.IsZero(m.Schema) { // not required
		return nil
	}

	if m.Schema != nil {
		if err := m.Schema.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schema")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigContext) validateSites(formats strfmt.Registry) error {
	if swag.IsZero(m.Sites) { // not required
		return nil
	}

	if err := validate.UniqueItems("sites", "body", m.Sites); err != nil {
		return err
	}

	for i := 0; i < len(m.Sites); i++ {

		if err := validate.FormatOf("sites"+"."+strconv.Itoa(i), "body", "uuid", m.Sites[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *ConfigContext) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	if err := validate.UniqueItems("tags", "body", m.Tags); err != nil {
		return err
	}

	for i := 0; i < len(m.Tags); i++ {

		if err := validate.Pattern("tags"+"."+strconv.Itoa(i), "body", m.Tags[i], `^[-a-zA-Z0-9_]+$`); err != nil {
			return err
		}

	}

	return nil
}

func (m *ConfigContext) validateTenantGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.TenantGroups) { // not required
		return nil
	}

	if err := validate.UniqueItems("tenant_groups", "body", m.TenantGroups); err != nil {
		return err
	}

	for i := 0; i < len(m.TenantGroups); i++ {

		if err := validate.FormatOf("tenant_groups"+"."+strconv.Itoa(i), "body", "uuid", m.TenantGroups[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *ConfigContext) validateTenants(formats strfmt.Registry) error {
	if swag.IsZero(m.Tenants) { // not required
		return nil
	}

	if err := validate.UniqueItems("tenants", "body", m.Tenants); err != nil {
		return err
	}

	for i := 0; i < len(m.Tenants); i++ {

		if err := validate.FormatOf("tenants"+"."+strconv.Itoa(i), "body", "uuid", m.Tenants[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *ConfigContext) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConfigContext) validateWeight(formats strfmt.Registry) error {
	if swag.IsZero(m.Weight) { // not required
		return nil
	}

	if err := validate.MinimumInt("weight", "body", *m.Weight, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("weight", "body", *m.Weight, 32767, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this config context based on the context it is used
func (m *ConfigContext) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOwner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSchema(ctx, formats); err != nil {
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

func (m *ConfigContext) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", strfmt.Date(m.Created)); err != nil {
		return err
	}

	return nil
}

func (m *ConfigContext) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *ConfigContext) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *ConfigContext) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", strfmt.DateTime(m.LastUpdated)); err != nil {
		return err
	}

	return nil
}

func (m *ConfigContext) contextValidateOwner(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ConfigContext) contextValidateSchema(ctx context.Context, formats strfmt.Registry) error {

	if m.Schema != nil {
		if err := m.Schema.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schema")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigContext) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigContext) UnmarshalBinary(b []byte) error {
	var res ConfigContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

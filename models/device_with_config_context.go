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

// DeviceWithConfigContext device with config context
//
// swagger:model DeviceWithConfigContext
type DeviceWithConfigContext struct {

	// Asset tag
	//
	// A unique tag used to identify this device
	// Max Length: 50
	AssetTag *string `json:"asset_tag,omitempty"`

	// cluster
	Cluster *NestedCluster `json:"cluster,omitempty"`

	// Comments
	Comments string `json:"comments,omitempty"`

	// Config context
	// Read Only: true
	ConfigContext map[string]*string `json:"config_context,omitempty"`

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// device role
	// Required: true
	DeviceRole *NestedDeviceRole `json:"device_role"`

	// device type
	// Required: true
	DeviceType *NestedDeviceType `json:"device_type"`

	// Display
	//
	// Human friendly display value
	// Read Only: true
	// Min Length: 1
	Display string `json:"display,omitempty"`

	// face
	Face *DeviceWithConfigContextFace `json:"face,omitempty"`

	// Id
	// Read Only: true
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Local context data
	LocalContextData *string `json:"local_context_data,omitempty"`

	// local context schema
	LocalContextSchema *NestedConfigContextSchema `json:"local_context_schema,omitempty"`

	// Name
	// Max Length: 64
	Name *string `json:"name,omitempty"`

	// parent device
	ParentDevice *NestedDevice `json:"parent_device,omitempty"`

	// platform
	Platform *NestedPlatform `json:"platform,omitempty"`

	// Position (U)
	//
	// The lowest-numbered unit occupied by the device
	// Maximum: 32767
	// Minimum: 1
	Position *int64 `json:"position,omitempty"`

	// primary ip
	PrimaryIP *NestedIPAddress `json:"primary_ip,omitempty"`

	// primary ip4
	PrimaryIp4 *NestedIPAddress `json:"primary_ip4,omitempty"`

	// primary ip6
	PrimaryIp6 *NestedIPAddress `json:"primary_ip6,omitempty"`

	// rack
	Rack *NestedRack `json:"rack,omitempty"`

	// Serial number
	// Max Length: 50
	Serial string `json:"serial,omitempty"`

	// site
	// Required: true
	Site *NestedSite `json:"site"`

	// Status
	// Pattern: ^[-a-zA-Z0-9_]+$
	// Enum: [active decommissioning failed inventory offline planned staged]
	Status string `json:"status,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// tenant
	Tenant *NestedTenant `json:"tenant,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// Vc position
	// Maximum: 255
	// Minimum: 0
	VcPosition *int64 `json:"vc_position,omitempty"`

	// Vc priority
	// Maximum: 255
	// Minimum: 0
	VcPriority *int64 `json:"vc_priority,omitempty"`

	// virtual chassis
	VirtualChassis *NestedVirtualChassis `json:"virtual_chassis,omitempty"`
}

// Validate validates this device with config context
func (m *DeviceWithConfigContext) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssetTag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalContextSchema(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatform(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryIp4(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryIp6(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRack(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSerial(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSite(formats); err != nil {
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

	if err := m.validateVcPosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcPriority(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualChassis(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceWithConfigContext) validateAssetTag(formats strfmt.Registry) error {
	if swag.IsZero(m.AssetTag) { // not required
		return nil
	}

	if err := validate.MaxLength("asset_tag", "body", *m.AssetTag, 50); err != nil {
		return err
	}

	return nil
}

func (m *DeviceWithConfigContext) validateCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.Cluster) { // not required
		return nil
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeviceWithConfigContext) validateDeviceRole(formats strfmt.Registry) error {

	if err := validate.Required("device_role", "body", m.DeviceRole); err != nil {
		return err
	}

	if m.DeviceRole != nil {
		if err := m.DeviceRole.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device_role")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) validateDeviceType(formats strfmt.Registry) error {

	if err := validate.Required("device_type", "body", m.DeviceType); err != nil {
		return err
	}

	if m.DeviceType != nil {
		if err := m.DeviceType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device_type")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) validateDisplay(formats strfmt.Registry) error {
	if swag.IsZero(m.Display) { // not required
		return nil
	}

	if err := validate.MinLength("display", "body", m.Display, 1); err != nil {
		return err
	}

	return nil
}

func (m *DeviceWithConfigContext) validateFace(formats strfmt.Registry) error {
	if swag.IsZero(m.Face) { // not required
		return nil
	}

	if m.Face != nil {
		if err := m.Face.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("face")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeviceWithConfigContext) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeviceWithConfigContext) validateLocalContextSchema(formats strfmt.Registry) error {
	if swag.IsZero(m.LocalContextSchema) { // not required
		return nil
	}

	if m.LocalContextSchema != nil {
		if err := m.LocalContextSchema.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("local_context_schema")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", *m.Name, 64); err != nil {
		return err
	}

	return nil
}

func (m *DeviceWithConfigContext) validateParentDevice(formats strfmt.Registry) error {
	if swag.IsZero(m.ParentDevice) { // not required
		return nil
	}

	if m.ParentDevice != nil {
		if err := m.ParentDevice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parent_device")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) validatePlatform(formats strfmt.Registry) error {
	if swag.IsZero(m.Platform) { // not required
		return nil
	}

	if m.Platform != nil {
		if err := m.Platform.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("platform")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) validatePosition(formats strfmt.Registry) error {
	if swag.IsZero(m.Position) { // not required
		return nil
	}

	if err := validate.MinimumInt("position", "body", *m.Position, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("position", "body", *m.Position, 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *DeviceWithConfigContext) validatePrimaryIP(formats strfmt.Registry) error {
	if swag.IsZero(m.PrimaryIP) { // not required
		return nil
	}

	if m.PrimaryIP != nil {
		if err := m.PrimaryIP.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary_ip")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) validatePrimaryIp4(formats strfmt.Registry) error {
	if swag.IsZero(m.PrimaryIp4) { // not required
		return nil
	}

	if m.PrimaryIp4 != nil {
		if err := m.PrimaryIp4.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary_ip4")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) validatePrimaryIp6(formats strfmt.Registry) error {
	if swag.IsZero(m.PrimaryIp6) { // not required
		return nil
	}

	if m.PrimaryIp6 != nil {
		if err := m.PrimaryIp6.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary_ip6")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) validateRack(formats strfmt.Registry) error {
	if swag.IsZero(m.Rack) { // not required
		return nil
	}

	if m.Rack != nil {
		if err := m.Rack.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rack")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) validateSerial(formats strfmt.Registry) error {
	if swag.IsZero(m.Serial) { // not required
		return nil
	}

	if err := validate.MaxLength("serial", "body", m.Serial, 50); err != nil {
		return err
	}

	return nil
}

func (m *DeviceWithConfigContext) validateSite(formats strfmt.Registry) error {

	if err := validate.Required("site", "body", m.Site); err != nil {
		return err
	}

	if m.Site != nil {
		if err := m.Site.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("site")
			}
			return err
		}
	}

	return nil
}

var deviceWithConfigContextTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","decommissioning","failed","inventory","offline","planned","staged"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deviceWithConfigContextTypeStatusPropEnum = append(deviceWithConfigContextTypeStatusPropEnum, v)
	}
}

const (

	// DeviceWithConfigContextStatusActive captures enum value "active"
	DeviceWithConfigContextStatusActive string = "active"

	// DeviceWithConfigContextStatusDecommissioning captures enum value "decommissioning"
	DeviceWithConfigContextStatusDecommissioning string = "decommissioning"

	// DeviceWithConfigContextStatusFailed captures enum value "failed"
	DeviceWithConfigContextStatusFailed string = "failed"

	// DeviceWithConfigContextStatusInventory captures enum value "inventory"
	DeviceWithConfigContextStatusInventory string = "inventory"

	// DeviceWithConfigContextStatusOffline captures enum value "offline"
	DeviceWithConfigContextStatusOffline string = "offline"

	// DeviceWithConfigContextStatusPlanned captures enum value "planned"
	DeviceWithConfigContextStatusPlanned string = "planned"

	// DeviceWithConfigContextStatusStaged captures enum value "staged"
	DeviceWithConfigContextStatusStaged string = "staged"
)

// prop value enum
func (m *DeviceWithConfigContext) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, deviceWithConfigContextTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DeviceWithConfigContext) validateStatus(formats strfmt.Registry) error {
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

func (m *DeviceWithConfigContext) validateTags(formats strfmt.Registry) error {
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

func (m *DeviceWithConfigContext) validateTenant(formats strfmt.Registry) error {
	if swag.IsZero(m.Tenant) { // not required
		return nil
	}

	if m.Tenant != nil {
		if err := m.Tenant.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tenant")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeviceWithConfigContext) validateVcPosition(formats strfmt.Registry) error {
	if swag.IsZero(m.VcPosition) { // not required
		return nil
	}

	if err := validate.MinimumInt("vc_position", "body", *m.VcPosition, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("vc_position", "body", *m.VcPosition, 255, false); err != nil {
		return err
	}

	return nil
}

func (m *DeviceWithConfigContext) validateVcPriority(formats strfmt.Registry) error {
	if swag.IsZero(m.VcPriority) { // not required
		return nil
	}

	if err := validate.MinimumInt("vc_priority", "body", *m.VcPriority, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("vc_priority", "body", *m.VcPriority, 255, false); err != nil {
		return err
	}

	return nil
}

func (m *DeviceWithConfigContext) validateVirtualChassis(formats strfmt.Registry) error {
	if swag.IsZero(m.VirtualChassis) { // not required
		return nil
	}

	if m.VirtualChassis != nil {
		if err := m.VirtualChassis.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtual_chassis")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this device with config context based on the context it is used
func (m *DeviceWithConfigContext) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConfigContext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeviceRole(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeviceType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFace(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocalContextSchema(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParentDevice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePlatform(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrimaryIP(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrimaryIp4(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrimaryIp6(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRack(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSite(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTenant(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVirtualChassis(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceWithConfigContext) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.Cluster != nil {
		if err := m.Cluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) contextValidateConfigContext(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *DeviceWithConfigContext) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", strfmt.Date(m.Created)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceWithConfigContext) contextValidateDeviceRole(ctx context.Context, formats strfmt.Registry) error {

	if m.DeviceRole != nil {
		if err := m.DeviceRole.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device_role")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) contextValidateDeviceType(ctx context.Context, formats strfmt.Registry) error {

	if m.DeviceType != nil {
		if err := m.DeviceType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device_type")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceWithConfigContext) contextValidateFace(ctx context.Context, formats strfmt.Registry) error {

	if m.Face != nil {
		if err := m.Face.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("face")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceWithConfigContext) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", strfmt.DateTime(m.LastUpdated)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceWithConfigContext) contextValidateLocalContextSchema(ctx context.Context, formats strfmt.Registry) error {

	if m.LocalContextSchema != nil {
		if err := m.LocalContextSchema.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("local_context_schema")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) contextValidateParentDevice(ctx context.Context, formats strfmt.Registry) error {

	if m.ParentDevice != nil {
		if err := m.ParentDevice.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parent_device")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) contextValidatePlatform(ctx context.Context, formats strfmt.Registry) error {

	if m.Platform != nil {
		if err := m.Platform.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("platform")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) contextValidatePrimaryIP(ctx context.Context, formats strfmt.Registry) error {

	if m.PrimaryIP != nil {
		if err := m.PrimaryIP.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary_ip")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) contextValidatePrimaryIp4(ctx context.Context, formats strfmt.Registry) error {

	if m.PrimaryIp4 != nil {
		if err := m.PrimaryIp4.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary_ip4")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) contextValidatePrimaryIp6(ctx context.Context, formats strfmt.Registry) error {

	if m.PrimaryIp6 != nil {
		if err := m.PrimaryIp6.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary_ip6")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) contextValidateRack(ctx context.Context, formats strfmt.Registry) error {

	if m.Rack != nil {
		if err := m.Rack.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rack")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) contextValidateSite(ctx context.Context, formats strfmt.Registry) error {

	if m.Site != nil {
		if err := m.Site.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("site")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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

func (m *DeviceWithConfigContext) contextValidateTenant(ctx context.Context, formats strfmt.Registry) error {

	if m.Tenant != nil {
		if err := m.Tenant.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tenant")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceWithConfigContext) contextValidateVirtualChassis(ctx context.Context, formats strfmt.Registry) error {

	if m.VirtualChassis != nil {
		if err := m.VirtualChassis.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtual_chassis")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceWithConfigContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceWithConfigContext) UnmarshalBinary(b []byte) error {
	var res DeviceWithConfigContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DeviceWithConfigContextFace Face
//
// swagger:model DeviceWithConfigContextFace
type DeviceWithConfigContextFace struct {

	// label
	// Required: true
	// Enum: [Front Rear]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [front rear]
	Value *string `json:"value"`
}

// Validate validates this device with config context face
func (m *DeviceWithConfigContextFace) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var deviceWithConfigContextFaceTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Front","Rear"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deviceWithConfigContextFaceTypeLabelPropEnum = append(deviceWithConfigContextFaceTypeLabelPropEnum, v)
	}
}

const (

	// DeviceWithConfigContextFaceLabelFront captures enum value "Front"
	DeviceWithConfigContextFaceLabelFront string = "Front"

	// DeviceWithConfigContextFaceLabelRear captures enum value "Rear"
	DeviceWithConfigContextFaceLabelRear string = "Rear"
)

// prop value enum
func (m *DeviceWithConfigContextFace) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, deviceWithConfigContextFaceTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DeviceWithConfigContextFace) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("face"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("face"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var deviceWithConfigContextFaceTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["front","rear"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deviceWithConfigContextFaceTypeValuePropEnum = append(deviceWithConfigContextFaceTypeValuePropEnum, v)
	}
}

const (

	// DeviceWithConfigContextFaceValueFront captures enum value "front"
	DeviceWithConfigContextFaceValueFront string = "front"

	// DeviceWithConfigContextFaceValueRear captures enum value "rear"
	DeviceWithConfigContextFaceValueRear string = "rear"
)

// prop value enum
func (m *DeviceWithConfigContextFace) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, deviceWithConfigContextFaceTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DeviceWithConfigContextFace) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("face"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("face"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this device with config context face based on context it is used
func (m *DeviceWithConfigContextFace) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeviceWithConfigContextFace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceWithConfigContextFace) UnmarshalBinary(b []byte) error {
	var res DeviceWithConfigContextFace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

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

// WritablePowerFeed writable power feed
//
// swagger:model WritablePowerFeed
type WritablePowerFeed struct {

	// Amperage
	// Maximum: 32767
	// Minimum: 1
	Amperage int64 `json:"amperage,omitempty"`

	// cable
	Cable *NestedCable `json:"cable,omitempty"`

	// Cable peer
	//
	//
	// Return the appropriate serializer for the cable termination model.
	//
	// Read Only: true
	CablePeer map[string]*string `json:"cable_peer,omitempty"`

	// Cable peer type
	// Read Only: true
	CablePeerType string `json:"cable_peer_type,omitempty"`

	// Comments
	Comments string `json:"comments,omitempty"`

	// Computed fields
	// Read Only: true
	ComputedFields string `json:"computed_fields,omitempty"`

	// Connected endpoint
	//
	//
	// Return the appropriate serializer for the type of connected object.
	//
	// Read Only: true
	ConnectedEndpoint map[string]*string `json:"connected_endpoint,omitempty"`

	// Connected endpoint reachable
	// Read Only: true
	ConnectedEndpointReachable *bool `json:"connected_endpoint_reachable,omitempty"`

	// Connected endpoint type
	// Read Only: true
	ConnectedEndpointType string `json:"connected_endpoint_type,omitempty"`

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

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

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Max utilization
	//
	// Maximum permissible draw (percentage)
	// Maximum: 100
	// Minimum: 1
	MaxUtilization int64 `json:"max_utilization,omitempty"`

	// Name
	// Required: true
	// Max Length: 100
	// Min Length: 1
	Name *string `json:"name"`

	// Phase
	// Enum: [single-phase three-phase]
	Phase string `json:"phase,omitempty"`

	// Power panel
	// Required: true
	// Format: uuid
	PowerPanel *strfmt.UUID `json:"power_panel"`

	// Rack
	// Format: uuid
	Rack *strfmt.UUID `json:"rack,omitempty"`

	// Status
	// Pattern: ^[-a-zA-Z0-9_]+$
	// Enum: [active failed offline planned]
	Status string `json:"status,omitempty"`

	// Supply
	// Enum: [ac dc]
	Supply string `json:"supply,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// Type
	// Enum: [primary redundant]
	Type string `json:"type,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// Voltage
	// Maximum: 32767
	// Minimum: -32768
	Voltage *int64 `json:"voltage,omitempty"`
}

// Validate validates this writable power feed
func (m *WritablePowerFeed) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmperage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
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

	if err := m.validateMaxUtilization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePowerPanel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRack(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSupply(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVoltage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritablePowerFeed) validateAmperage(formats strfmt.Registry) error {
	if swag.IsZero(m.Amperage) { // not required
		return nil
	}

	if err := validate.MinimumInt("amperage", "body", m.Amperage, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("amperage", "body", m.Amperage, 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) validateCable(formats strfmt.Registry) error {
	if swag.IsZero(m.Cable) { // not required
		return nil
	}

	if m.Cable != nil {
		if err := m.Cable.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cable")
			}
			return err
		}
	}

	return nil
}

func (m *WritablePowerFeed) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) validateDisplay(formats strfmt.Registry) error {
	if swag.IsZero(m.Display) { // not required
		return nil
	}

	if err := validate.MinLength("display", "body", m.Display, 1); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) validateMaxUtilization(formats strfmt.Registry) error {
	if swag.IsZero(m.MaxUtilization) { // not required
		return nil
	}

	if err := validate.MinimumInt("max_utilization", "body", m.MaxUtilization, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("max_utilization", "body", m.MaxUtilization, 100, false); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) validateName(formats strfmt.Registry) error {

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

var writablePowerFeedTypePhasePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["single-phase","three-phase"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writablePowerFeedTypePhasePropEnum = append(writablePowerFeedTypePhasePropEnum, v)
	}
}

const (

	// WritablePowerFeedPhaseSingleDashPhase captures enum value "single-phase"
	WritablePowerFeedPhaseSingleDashPhase string = "single-phase"

	// WritablePowerFeedPhaseThreeDashPhase captures enum value "three-phase"
	WritablePowerFeedPhaseThreeDashPhase string = "three-phase"
)

// prop value enum
func (m *WritablePowerFeed) validatePhaseEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writablePowerFeedTypePhasePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritablePowerFeed) validatePhase(formats strfmt.Registry) error {
	if swag.IsZero(m.Phase) { // not required
		return nil
	}

	// value enum
	if err := m.validatePhaseEnum("phase", "body", m.Phase); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) validatePowerPanel(formats strfmt.Registry) error {

	if err := validate.Required("power_panel", "body", m.PowerPanel); err != nil {
		return err
	}

	if err := validate.FormatOf("power_panel", "body", "uuid", m.PowerPanel.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) validateRack(formats strfmt.Registry) error {
	if swag.IsZero(m.Rack) { // not required
		return nil
	}

	if err := validate.FormatOf("rack", "body", "uuid", m.Rack.String(), formats); err != nil {
		return err
	}

	return nil
}

var writablePowerFeedTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","failed","offline","planned"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writablePowerFeedTypeStatusPropEnum = append(writablePowerFeedTypeStatusPropEnum, v)
	}
}

const (

	// WritablePowerFeedStatusActive captures enum value "active"
	WritablePowerFeedStatusActive string = "active"

	// WritablePowerFeedStatusFailed captures enum value "failed"
	WritablePowerFeedStatusFailed string = "failed"

	// WritablePowerFeedStatusOffline captures enum value "offline"
	WritablePowerFeedStatusOffline string = "offline"

	// WritablePowerFeedStatusPlanned captures enum value "planned"
	WritablePowerFeedStatusPlanned string = "planned"
)

// prop value enum
func (m *WritablePowerFeed) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writablePowerFeedTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritablePowerFeed) validateStatus(formats strfmt.Registry) error {
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

var writablePowerFeedTypeSupplyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ac","dc"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writablePowerFeedTypeSupplyPropEnum = append(writablePowerFeedTypeSupplyPropEnum, v)
	}
}

const (

	// WritablePowerFeedSupplyAc captures enum value "ac"
	WritablePowerFeedSupplyAc string = "ac"

	// WritablePowerFeedSupplyDc captures enum value "dc"
	WritablePowerFeedSupplyDc string = "dc"
)

// prop value enum
func (m *WritablePowerFeed) validateSupplyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writablePowerFeedTypeSupplyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritablePowerFeed) validateSupply(formats strfmt.Registry) error {
	if swag.IsZero(m.Supply) { // not required
		return nil
	}

	// value enum
	if err := m.validateSupplyEnum("supply", "body", m.Supply); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) validateTags(formats strfmt.Registry) error {
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

var writablePowerFeedTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["primary","redundant"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writablePowerFeedTypeTypePropEnum = append(writablePowerFeedTypeTypePropEnum, v)
	}
}

const (

	// WritablePowerFeedTypePrimary captures enum value "primary"
	WritablePowerFeedTypePrimary string = "primary"

	// WritablePowerFeedTypeRedundant captures enum value "redundant"
	WritablePowerFeedTypeRedundant string = "redundant"
)

// prop value enum
func (m *WritablePowerFeed) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writablePowerFeedTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritablePowerFeed) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) validateVoltage(formats strfmt.Registry) error {
	if swag.IsZero(m.Voltage) { // not required
		return nil
	}

	if err := validate.MinimumInt("voltage", "body", *m.Voltage, -32768, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("voltage", "body", *m.Voltage, 32767, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this writable power feed based on the context it is used
func (m *WritablePowerFeed) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCablePeer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCablePeerType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateComputedFields(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnectedEndpoint(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnectedEndpointReachable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnectedEndpointType(ctx, formats); err != nil {
		res = append(res, err)
	}

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

func (m *WritablePowerFeed) contextValidateCable(ctx context.Context, formats strfmt.Registry) error {

	if m.Cable != nil {
		if err := m.Cable.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cable")
			}
			return err
		}
	}

	return nil
}

func (m *WritablePowerFeed) contextValidateCablePeer(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *WritablePowerFeed) contextValidateCablePeerType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "cable_peer_type", "body", string(m.CablePeerType)); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) contextValidateComputedFields(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "computed_fields", "body", string(m.ComputedFields)); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) contextValidateConnectedEndpoint(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *WritablePowerFeed) contextValidateConnectedEndpointReachable(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connected_endpoint_reachable", "body", m.ConnectedEndpointReachable); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) contextValidateConnectedEndpointType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connected_endpoint_type", "body", string(m.ConnectedEndpointType)); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", strfmt.Date(m.Created)); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", strfmt.DateTime(m.LastUpdated)); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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

func (m *WritablePowerFeed) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritablePowerFeed) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritablePowerFeed) UnmarshalBinary(b []byte) error {
	var res WritablePowerFeed
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

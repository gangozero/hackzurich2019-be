// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GameDescription game description
// swagger:model gameDescription
type GameDescription struct {

	// Game ID
	ID string `json:"id,omitempty"`

	// Status of the game
	// Enum: [ACTIVE PAUSED FINISHED]
	Status string `json:"status,omitempty"`
}

// Validate validates this game description
func (m *GameDescription) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var gameDescriptionTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACTIVE","PAUSED","FINISHED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		gameDescriptionTypeStatusPropEnum = append(gameDescriptionTypeStatusPropEnum, v)
	}
}

const (

	// GameDescriptionStatusACTIVE captures enum value "ACTIVE"
	GameDescriptionStatusACTIVE string = "ACTIVE"

	// GameDescriptionStatusPAUSED captures enum value "PAUSED"
	GameDescriptionStatusPAUSED string = "PAUSED"

	// GameDescriptionStatusFINISHED captures enum value "FINISHED"
	GameDescriptionStatusFINISHED string = "FINISHED"
)

// prop value enum
func (m *GameDescription) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, gameDescriptionTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GameDescription) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GameDescription) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GameDescription) UnmarshalBinary(b []byte) error {
	var res GameDescription
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

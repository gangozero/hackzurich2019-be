// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GameState game state
// swagger:model gameState
type GameState struct {

	// Pinguin count
	Count int64 `json:"count,omitempty"`

	// Ship count
	CountFarm int64 `json:"countFarm,omitempty"`

	// Ship count
	CountShip int64 `json:"countShip,omitempty"`

	// farms
	Farms []*Point `json:"farms"`

	// ships
	Ships []*Point `json:"ships"`
}

// Validate validates this game state
func (m *GameState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFarms(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShips(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GameState) validateFarms(formats strfmt.Registry) error {

	if swag.IsZero(m.Farms) { // not required
		return nil
	}

	for i := 0; i < len(m.Farms); i++ {
		if swag.IsZero(m.Farms[i]) { // not required
			continue
		}

		if m.Farms[i] != nil {
			if err := m.Farms[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("farms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GameState) validateShips(formats strfmt.Registry) error {

	if swag.IsZero(m.Ships) { // not required
		return nil
	}

	for i := 0; i < len(m.Ships); i++ {
		if swag.IsZero(m.Ships[i]) { // not required
			continue
		}

		if m.Ships[i] != nil {
			if err := m.Ships[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ships" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GameState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GameState) UnmarshalBinary(b []byte) error {
	var res GameState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

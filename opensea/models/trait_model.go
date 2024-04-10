// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TraitModel TraitModel
//
// swagger:model TraitModel
type TraitModel struct {

	// display type
	DisplayType DisplayTypeField `json:"display_type,omitempty"`

	// Max Value
	//
	// Ceiling for possible numeric trait values
	// Required: true
	MaxValue *string `json:"max_value"`

	// Order
	//
	// Deprecated Field
	Order int64 `json:"order,omitempty"`

	// Trait Count
	//
	// Deprecated Field. Use Get Collection API instead.
	TraitCount int64 `json:"trait_count,omitempty"`

	// Trait Type
	//
	// The name of the trait category (e.g. 'Background')
	// Required: true
	// Max Length: 150
	TraitType *string `json:"trait_type"`

	// Value
	//
	// The value of the trait (e.g. 'Red')
	// Required: true
	Value interface{} `json:"value"`
}

// Validate validates this trait model
func (m *TraitModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisplayType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTraitType(formats); err != nil {
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

func (m *TraitModel) validateDisplayType(formats strfmt.Registry) error {
	if swag.IsZero(m.DisplayType) { // not required
		return nil
	}

	if err := m.DisplayType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("display_type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("display_type")
		}
		return err
	}

	return nil
}

func (m *TraitModel) validateMaxValue(formats strfmt.Registry) error {

	if err := validate.Required("max_value", "body", m.MaxValue); err != nil {
		return err
	}

	return nil
}

func (m *TraitModel) validateTraitType(formats strfmt.Registry) error {

	if err := validate.Required("trait_type", "body", m.TraitType); err != nil {
		return err
	}

	if err := validate.MaxLength("trait_type", "body", *m.TraitType, 150); err != nil {
		return err
	}

	return nil
}

func (m *TraitModel) validateValue(formats strfmt.Registry) error {

	if m.Value == nil {
		return errors.Required("value", "body", nil)
	}

	return nil
}

// ContextValidate validate this trait model based on the context it is used
func (m *TraitModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDisplayType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TraitModel) contextValidateDisplayType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.DisplayType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("display_type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("display_type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TraitModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TraitModel) UnmarshalBinary(b []byte) error {
	var res TraitModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
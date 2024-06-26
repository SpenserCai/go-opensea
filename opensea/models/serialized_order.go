// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SerializedOrder SerializedOrder
//
// swagger:model SerializedOrder
type SerializedOrder struct {

	// Order
	// Required: true
	Parameters struct {
		SerializedOrderComponents
	} `json:"parameters"`

	// Signature
	//
	// The order maker's signature used to validate the order.
	Signature string `json:"signature,omitempty"`
}

// Validate validates this serialized order
func (m *SerializedOrder) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateParameters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SerializedOrder) validateParameters(formats strfmt.Registry) error {

	return nil
}

// ContextValidate validate this serialized order based on the context it is used
func (m *SerializedOrder) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateParameters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SerializedOrder) contextValidateParameters(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *SerializedOrder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SerializedOrder) UnmarshalBinary(b []byte) error {
	var res SerializedOrder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

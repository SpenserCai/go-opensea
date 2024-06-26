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

// CollectionContractModel CollectionContractModel
//
// # Define the Contract's Addresses and Chain
//
// swagger:model CollectionContractModel
type CollectionContractModel struct {

	// Address
	//
	// The unique public blockchain identifier, address, for the contract.
	// Required: true
	Address *string `json:"address"`

	// The chain on which the contract exists
	// Required: true
	Chain struct {
		ChainIdentifier
	} `json:"chain"`
}

// Validate validates this collection contract model
func (m *CollectionContractModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChain(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CollectionContractModel) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("address", "body", m.Address); err != nil {
		return err
	}

	return nil
}

func (m *CollectionContractModel) validateChain(formats strfmt.Registry) error {

	return nil
}

// ContextValidate validates this collection contract model based on context it is used
func (m *CollectionContractModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CollectionContractModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CollectionContractModel) UnmarshalBinary(b []byte) error {
	var res CollectionContractModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

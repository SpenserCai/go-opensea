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

// GetNftResponse GetNftResponse
//
// swagger:model GetNftResponse
type GetNftResponse struct {

	// nft
	// Required: true
	Nft *DetailedNftModel `json:"nft"`
}

// Validate validates this get nft response
func (m *GetNftResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNft(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetNftResponse) validateNft(formats strfmt.Registry) error {

	if err := validate.Required("nft", "body", m.Nft); err != nil {
		return err
	}

	if m.Nft != nil {
		if err := m.Nft.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nft")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nft")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get nft response based on the context it is used
func (m *GetNftResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNft(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetNftResponse) contextValidateNft(ctx context.Context, formats strfmt.Registry) error {

	if m.Nft != nil {
		if err := m.Nft.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nft")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nft")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetNftResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetNftResponse) UnmarshalBinary(b []byte) error {
	var res GetNftResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

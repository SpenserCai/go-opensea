// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ListNftsResponse ListNftsResponse
//
// swagger:model ListNftsResponse
type ListNftsResponse struct {

	// Next
	//
	// Cursor for the next page of results
	// Required: true
	Next *string `json:"next"`

	// NFT
	// Required: true
	Nfts []*NftModel `json:"nfts"`
}

// Validate validates this list nfts response
func (m *ListNftsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNfts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListNftsResponse) validateNext(formats strfmt.Registry) error {

	if err := validate.Required("next", "body", m.Next); err != nil {
		return err
	}

	return nil
}

func (m *ListNftsResponse) validateNfts(formats strfmt.Registry) error {

	if err := validate.Required("nfts", "body", m.Nfts); err != nil {
		return err
	}

	for i := 0; i < len(m.Nfts); i++ {
		if swag.IsZero(m.Nfts[i]) { // not required
			continue
		}

		if m.Nfts[i] != nil {
			if err := m.Nfts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nfts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nfts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list nfts response based on the context it is used
func (m *ListNftsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNfts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListNftsResponse) contextValidateNfts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Nfts); i++ {

		if m.Nfts[i] != nil {
			if err := m.Nfts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nfts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nfts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListNftsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListNftsResponse) UnmarshalBinary(b []byte) error {
	var res ListNftsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

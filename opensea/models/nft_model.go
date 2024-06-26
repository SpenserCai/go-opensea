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

// NftModel NftModel
//
// swagger:model NftModel
type NftModel struct {

	// Collection
	//
	// Collection slug. A unique string to identify a collection on OpenSea
	// Required: true
	Collection *string `json:"collection"`

	// Contract
	//
	// The unique public blockchain identifier for the contract
	// Required: true
	Contract *string `json:"contract"`

	// Created At
	//
	// Deprecated Field
	CreatedAt *string `json:"created_at,omitempty"`

	// Description
	//
	// Description of the NFT
	// Required: true
	Description *string `json:"description"`

	// Identifier
	//
	// The NFT's unique identifier within the smart contract (also referred to as token_id)
	// Required: true
	Identifier *string `json:"identifier"`

	// Image Url
	//
	// Link to the image associated with the NFT
	ImageURL string `json:"image_url,omitempty"`

	// Is Disabled
	//
	// If the item is currently able to be bought or sold using OpenSea
	// Required: true
	IsDisabled *bool `json:"is_disabled"`

	// Is Nsfw
	//
	// If the item is currently classified as 'Not Safe for Work' by OpenSea
	// Required: true
	IsNsfw *bool `json:"is_nsfw"`

	// Metadata Url
	//
	// Link to the offchain metadata store
	MetadataURL string `json:"metadata_url,omitempty"`

	// Name
	//
	// Name of the NFT
	// Required: true
	Name *string `json:"name"`

	// Token Standard
	//
	// ERC standard of the token (erc721, erc1155)
	// Required: true
	TokenStandard *string `json:"token_standard"`

	// Updated At
	//
	// Last time that the NFT's metadata was updated by OpenSea
	// Required: true
	UpdatedAt *string `json:"updated_at"`
}

// Validate validates this nft model
func (m *NftModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCollection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContract(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentifier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsDisabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsNsfw(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTokenStandard(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NftModel) validateCollection(formats strfmt.Registry) error {

	if err := validate.Required("collection", "body", m.Collection); err != nil {
		return err
	}

	return nil
}

func (m *NftModel) validateContract(formats strfmt.Registry) error {

	if err := validate.Required("contract", "body", m.Contract); err != nil {
		return err
	}

	return nil
}

func (m *NftModel) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *NftModel) validateIdentifier(formats strfmt.Registry) error {

	if err := validate.Required("identifier", "body", m.Identifier); err != nil {
		return err
	}

	return nil
}

func (m *NftModel) validateIsDisabled(formats strfmt.Registry) error {

	if err := validate.Required("is_disabled", "body", m.IsDisabled); err != nil {
		return err
	}

	return nil
}

func (m *NftModel) validateIsNsfw(formats strfmt.Registry) error {

	if err := validate.Required("is_nsfw", "body", m.IsNsfw); err != nil {
		return err
	}

	return nil
}

func (m *NftModel) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *NftModel) validateTokenStandard(formats strfmt.Registry) error {

	if err := validate.Required("token_standard", "body", m.TokenStandard); err != nil {
		return err
	}

	return nil
}

func (m *NftModel) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this nft model based on context it is used
func (m *NftModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NftModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NftModel) UnmarshalBinary(b []byte) error {
	var res NftModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

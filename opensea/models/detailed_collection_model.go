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

// DetailedCollectionModel DetailedCollectionModel
//
// swagger:model DetailedCollectionModel
type DetailedCollectionModel struct {

	// Category
	//
	// Category of the collection (e.g. PFPs, Memberships, Art)
	// Required: true
	Category *string `json:"category"`

	// Collection
	//
	// Collection slug. A unique string to identify a collection on OpenSea
	// Required: true
	Collection *string `json:"collection"`

	// Contract
	// Required: true
	Contracts []*CollectionContractModel `json:"contracts"`

	// Description
	//
	// Description of the collection
	Description string `json:"description,omitempty"`

	// Discord Url
	//
	// External URL for the collection's Discord server
	DiscordURL string `json:"discord_url,omitempty"`

	// Editors
	//
	// List of editor addresses for the collection
	// Required: true
	Editors []string `json:"editors"`

	// Fees
	//
	// List of fees for the collection including creator earnings and OpenSea fees
	// Required: true
	Fees []*CollectionFeeModel `json:"fees"`

	// Instagram Username
	//
	// Username for the collection's Instagram account
	InstagramUsername string `json:"instagram_username,omitempty"`

	// Is Disabled
	//
	// If the collection is currently able to be bought or sold using OpenSea
	// Required: true
	IsDisabled *bool `json:"is_disabled"`

	// Is Nsfw
	//
	// If the collection is currently classified as 'Not Safe for Work' by OpenSea
	// Required: true
	IsNsfw *bool `json:"is_nsfw"`

	// Name
	//
	// Name of the collection
	// Required: true
	Name *string `json:"name"`

	// Opensea Url
	//
	// OpenSea Link to collection
	// Required: true
	OpenseaURL *string `json:"opensea_url"`

	// Owner
	//
	// The unique public blockchain identifier, address, for the owner wallet.
	// Required: true
	Owner *string `json:"owner"`

	// Project Url
	//
	// External URL for the collection's website
	ProjectURL string `json:"project_url,omitempty"`

	// safelist status
	// Required: true
	SafelistStatus SafelistRequestStatus `json:"safelist_status"`

	// Telegram Url
	//
	// External URL for the collection's Telegram group
	TelegramURL string `json:"telegram_url,omitempty"`

	// Trait Offers Enabled
	//
	// If trait offers are currently being accepted for the collection
	// Required: true
	TraitOffersEnabled *bool `json:"trait_offers_enabled"`

	// Twitter Username
	//
	// Username for the collection's Twitter account
	TwitterUsername string `json:"twitter_username,omitempty"`

	// Wiki Url
	//
	// External URL for the collection's wiki
	WikiURL string `json:"wiki_url,omitempty"`
}

// Validate validates this detailed collection model
func (m *DetailedCollectionModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCollection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContracts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEditors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFees(formats); err != nil {
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

	if err := m.validateOpenseaURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSafelistStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTraitOffersEnabled(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DetailedCollectionModel) validateCategory(formats strfmt.Registry) error {

	if err := validate.Required("category", "body", m.Category); err != nil {
		return err
	}

	return nil
}

func (m *DetailedCollectionModel) validateCollection(formats strfmt.Registry) error {

	if err := validate.Required("collection", "body", m.Collection); err != nil {
		return err
	}

	return nil
}

func (m *DetailedCollectionModel) validateContracts(formats strfmt.Registry) error {

	if err := validate.Required("contracts", "body", m.Contracts); err != nil {
		return err
	}

	for i := 0; i < len(m.Contracts); i++ {
		if swag.IsZero(m.Contracts[i]) { // not required
			continue
		}

		if m.Contracts[i] != nil {
			if err := m.Contracts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("contracts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("contracts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DetailedCollectionModel) validateEditors(formats strfmt.Registry) error {

	if err := validate.Required("editors", "body", m.Editors); err != nil {
		return err
	}

	return nil
}

func (m *DetailedCollectionModel) validateFees(formats strfmt.Registry) error {

	if err := validate.Required("fees", "body", m.Fees); err != nil {
		return err
	}

	for i := 0; i < len(m.Fees); i++ {
		if swag.IsZero(m.Fees[i]) { // not required
			continue
		}

		if m.Fees[i] != nil {
			if err := m.Fees[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fees" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("fees" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DetailedCollectionModel) validateIsDisabled(formats strfmt.Registry) error {

	if err := validate.Required("is_disabled", "body", m.IsDisabled); err != nil {
		return err
	}

	return nil
}

func (m *DetailedCollectionModel) validateIsNsfw(formats strfmt.Registry) error {

	if err := validate.Required("is_nsfw", "body", m.IsNsfw); err != nil {
		return err
	}

	return nil
}

func (m *DetailedCollectionModel) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *DetailedCollectionModel) validateOpenseaURL(formats strfmt.Registry) error {

	if err := validate.Required("opensea_url", "body", m.OpenseaURL); err != nil {
		return err
	}

	return nil
}

func (m *DetailedCollectionModel) validateOwner(formats strfmt.Registry) error {

	if err := validate.Required("owner", "body", m.Owner); err != nil {
		return err
	}

	return nil
}

func (m *DetailedCollectionModel) validateSafelistStatus(formats strfmt.Registry) error {

	if m.SafelistStatus == nil {
		return errors.Required("safelist_status", "body", nil)
	}

	return nil
}

func (m *DetailedCollectionModel) validateTraitOffersEnabled(formats strfmt.Registry) error {

	if err := validate.Required("trait_offers_enabled", "body", m.TraitOffersEnabled); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this detailed collection model based on the context it is used
func (m *DetailedCollectionModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContracts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFees(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DetailedCollectionModel) contextValidateContracts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Contracts); i++ {

		if m.Contracts[i] != nil {
			if err := m.Contracts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("contracts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("contracts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DetailedCollectionModel) contextValidateFees(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Fees); i++ {

		if m.Fees[i] != nil {
			if err := m.Fees[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fees" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("fees" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DetailedCollectionModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DetailedCollectionModel) UnmarshalBinary(b []byte) error {
	var res DetailedCollectionModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

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

// GenerateOfferFulfillmentInput GenerateOfferFulfillmentInput
//
// swagger:model GenerateOfferFulfillmentInput
type GenerateOfferFulfillmentInput struct {

	// Consideration
	//
	// If the offer you are fulfilling is a criteria offer, the NFT you are using to fulfill the offer with. The fulfiller account must own this NFT or the request will not succeed.
	Consideration struct {
		ConsiderationInput
	} `json:"consideration,omitempty"`

	// Fulfiller
	//
	// Fulfiller address.
	// Required: true
	Fulfiller struct {
		FulfillerInput
	} `json:"fulfiller"`

	// Offer
	//
	// Offer to get fullfillment data for.
	// Required: true
	Offer struct {
		FulfillmentInput
	} `json:"offer"`
}

// Validate validates this generate offer fulfillment input
func (m *GenerateOfferFulfillmentInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConsideration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFulfiller(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOffer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GenerateOfferFulfillmentInput) validateConsideration(formats strfmt.Registry) error {
	if swag.IsZero(m.Consideration) { // not required
		return nil
	}

	return nil
}

func (m *GenerateOfferFulfillmentInput) validateFulfiller(formats strfmt.Registry) error {

	return nil
}

func (m *GenerateOfferFulfillmentInput) validateOffer(formats strfmt.Registry) error {

	return nil
}

// ContextValidate validate this generate offer fulfillment input based on the context it is used
func (m *GenerateOfferFulfillmentInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConsideration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFulfiller(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOffer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GenerateOfferFulfillmentInput) contextValidateConsideration(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *GenerateOfferFulfillmentInput) contextValidateFulfiller(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *GenerateOfferFulfillmentInput) contextValidateOffer(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *GenerateOfferFulfillmentInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GenerateOfferFulfillmentInput) UnmarshalBinary(b []byte) error {
	var res GenerateOfferFulfillmentInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

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

// BuildOfferInput BuildOfferInput
//
// swagger:model BuildOfferInput
type BuildOfferInput struct {

	// Criteria
	//
	// Criteria for the collection or trait offer
	// Required: true
	Criteria struct {
		Criteria
	} `json:"criteria"`

	// Offer Protection Enabled
	//
	// Builds the offer on OpenSea's signed zone to provide offer protections from receiving an item which is disabled from trading.
	OfferProtectionEnabled *bool `json:"offer_protection_enabled,omitempty"`

	// Offerer
	//
	// The address which supplies all the items in the offer.
	// Required: true
	Offerer *string `json:"offerer"`

	// Protocol Address
	//
	// Exchange contract address. Must be one of ['0x00000000000000adc04c56bf30ac9d3c0aaf14dc']
	// Required: true
	ProtocolAddress *string `json:"protocol_address"`

	// Quantity
	//
	// The number of offers to place.
	Quantity *int64 `json:"quantity,omitempty"`
}

// Validate validates this build offer input
func (m *BuildOfferInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCriteria(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOfferer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocolAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuildOfferInput) validateCriteria(formats strfmt.Registry) error {

	return nil
}

func (m *BuildOfferInput) validateOfferer(formats strfmt.Registry) error {

	if err := validate.Required("offerer", "body", m.Offerer); err != nil {
		return err
	}

	return nil
}

func (m *BuildOfferInput) validateProtocolAddress(formats strfmt.Registry) error {

	if err := validate.Required("protocol_address", "body", m.ProtocolAddress); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this build offer input based on the context it is used
func (m *BuildOfferInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCriteria(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuildOfferInput) contextValidateCriteria(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *BuildOfferInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuildOfferInput) UnmarshalBinary(b []byte) error {
	var res BuildOfferInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

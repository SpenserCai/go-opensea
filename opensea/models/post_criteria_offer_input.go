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

// PostCriteriaOfferInput PostCriteriaOfferInput
//
// swagger:model PostCriteriaOfferInput
type PostCriteriaOfferInput struct {

	// Criteria
	//
	// Criteria for the collection or trait offer
	// Required: true
	Criteria struct {
		Criteria
	} `json:"criteria"`

	// Protocol Address
	//
	// Exchange contract address. Must be one of ['0x00000000000000adc04c56bf30ac9d3c0aaf14dc']
	// Required: true
	ProtocolAddress *string `json:"protocol_address"`

	// Signed Seaport Order
	//
	// The signed order which will be submitted to Seaport
	// Required: true
	ProtocolData struct {
		OrderInput
	} `json:"protocol_data"`
}

// Validate validates this post criteria offer input
func (m *PostCriteriaOfferInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCriteria(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocolAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocolData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostCriteriaOfferInput) validateCriteria(formats strfmt.Registry) error {

	return nil
}

func (m *PostCriteriaOfferInput) validateProtocolAddress(formats strfmt.Registry) error {

	if err := validate.Required("protocol_address", "body", m.ProtocolAddress); err != nil {
		return err
	}

	return nil
}

func (m *PostCriteriaOfferInput) validateProtocolData(formats strfmt.Registry) error {

	return nil
}

// ContextValidate validate this post criteria offer input based on the context it is used
func (m *PostCriteriaOfferInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCriteria(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProtocolData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostCriteriaOfferInput) contextValidateCriteria(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PostCriteriaOfferInput) contextValidateProtocolData(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PostCriteriaOfferInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostCriteriaOfferInput) UnmarshalBinary(b []byte) error {
	var res PostCriteriaOfferInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

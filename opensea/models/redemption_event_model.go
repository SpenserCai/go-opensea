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

// RedemptionEventModel RedemptionEventModel
//
// swagger:model RedemptionEventModel
type RedemptionEventModel struct {

	// Asset
	//
	// The asset being redeemed.
	// Required: true
	Asset interface{} `json:"asset"`

	// The chain on which the rededemption occurred
	// Required: true
	Chain struct {
		ChainIdentifier
	} `json:"chain"`

	// Event Type
	EventType struct {
		RedemptionEventModelEventTypeEnum
	} `json:"event_type,omitempty"`

	// From Address
	//
	// Address of the sender
	// Required: true
	FromAddress *string `json:"from_address"`

	// Quantity
	//
	// Number of assets redeemed
	// Required: true
	Quantity *int64 `json:"quantity"`

	// To Address
	//
	// Address of the recipient
	// Required: true
	ToAddress *string `json:"to_address"`

	// Transaction
	//
	// Transaction hash for the redemption
	// Required: true
	Transaction *string `json:"transaction"`
}

// Validate validates this redemption event model
func (m *RedemptionEventModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAsset(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFromAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransaction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RedemptionEventModel) validateAsset(formats strfmt.Registry) error {

	if m.Asset == nil {
		return errors.Required("asset", "body", nil)
	}

	return nil
}

func (m *RedemptionEventModel) validateChain(formats strfmt.Registry) error {

	return nil
}

func (m *RedemptionEventModel) validateEventType(formats strfmt.Registry) error {
	if swag.IsZero(m.EventType) { // not required
		return nil
	}

	return nil
}

func (m *RedemptionEventModel) validateFromAddress(formats strfmt.Registry) error {

	if err := validate.Required("from_address", "body", m.FromAddress); err != nil {
		return err
	}

	return nil
}

func (m *RedemptionEventModel) validateQuantity(formats strfmt.Registry) error {

	if err := validate.Required("quantity", "body", m.Quantity); err != nil {
		return err
	}

	return nil
}

func (m *RedemptionEventModel) validateToAddress(formats strfmt.Registry) error {

	if err := validate.Required("to_address", "body", m.ToAddress); err != nil {
		return err
	}

	return nil
}

func (m *RedemptionEventModel) validateTransaction(formats strfmt.Registry) error {

	if err := validate.Required("transaction", "body", m.Transaction); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this redemption event model based on the context it is used
func (m *RedemptionEventModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEventType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RedemptionEventModel) contextValidateEventType(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *RedemptionEventModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RedemptionEventModel) UnmarshalBinary(b []byte) error {
	var res RedemptionEventModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

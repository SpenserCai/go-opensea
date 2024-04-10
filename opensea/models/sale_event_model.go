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

// SaleEventModel SaleEventModel
//
// swagger:model SaleEventModel
type SaleEventModel struct {

	// The chain on which the order was fulfilled
	// Required: true
	Chain struct {
		ChainIdentifier
	} `json:"chain"`

	// Closing Date
	//
	// The Posix timestamp at which the transaction which filled the order occurred
	// Required: true
	ClosingDate *int64 `json:"closing_date"`

	// Event Type
	EventType struct {
		SaleEventModelEventTypeEnum
	} `json:"event_type,omitempty"`

	// Maker
	//
	// Maker of the order
	// Required: true
	Maker *string `json:"maker"`

	// Order Hash
	//
	// Order hash for the order which was fulfilled
	// Required: true
	OrderHash *string `json:"order_hash"`

	// Payment
	// Required: true
	Payment struct {
		EventPaymentModel
	} `json:"payment"`

	// Protocol Address
	//
	// Exchange contract address which fulfilled the order
	// Required: true
	ProtocolAddress *string `json:"protocol_address"`

	// Quantity
	//
	// Number of assets transferred
	// Required: true
	Quantity *int64 `json:"quantity"`

	// Taker
	//
	// Taker of the order
	// Required: true
	Taker *string `json:"taker"`

	// Transaction
	//
	// Transaction hash for the order fulfillment
	// Required: true
	Transaction *string `json:"transaction"`
}

// Validate validates this sale event model
func (m *SaleEventModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClosingDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaker(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePayment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocolAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaker(formats); err != nil {
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

func (m *SaleEventModel) validateChain(formats strfmt.Registry) error {

	return nil
}

func (m *SaleEventModel) validateClosingDate(formats strfmt.Registry) error {

	if err := validate.Required("closing_date", "body", m.ClosingDate); err != nil {
		return err
	}

	return nil
}

func (m *SaleEventModel) validateEventType(formats strfmt.Registry) error {
	if swag.IsZero(m.EventType) { // not required
		return nil
	}

	return nil
}

func (m *SaleEventModel) validateMaker(formats strfmt.Registry) error {

	if err := validate.Required("maker", "body", m.Maker); err != nil {
		return err
	}

	return nil
}

func (m *SaleEventModel) validateOrderHash(formats strfmt.Registry) error {

	if err := validate.Required("order_hash", "body", m.OrderHash); err != nil {
		return err
	}

	return nil
}

func (m *SaleEventModel) validatePayment(formats strfmt.Registry) error {

	return nil
}

func (m *SaleEventModel) validateProtocolAddress(formats strfmt.Registry) error {

	if err := validate.Required("protocol_address", "body", m.ProtocolAddress); err != nil {
		return err
	}

	return nil
}

func (m *SaleEventModel) validateQuantity(formats strfmt.Registry) error {

	if err := validate.Required("quantity", "body", m.Quantity); err != nil {
		return err
	}

	return nil
}

func (m *SaleEventModel) validateTaker(formats strfmt.Registry) error {

	if err := validate.Required("taker", "body", m.Taker); err != nil {
		return err
	}

	return nil
}

func (m *SaleEventModel) validateTransaction(formats strfmt.Registry) error {

	if err := validate.Required("transaction", "body", m.Transaction); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this sale event model based on the context it is used
func (m *SaleEventModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEventType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePayment(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SaleEventModel) contextValidateEventType(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *SaleEventModel) contextValidatePayment(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *SaleEventModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SaleEventModel) UnmarshalBinary(b []byte) error {
	var res SaleEventModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
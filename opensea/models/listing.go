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

// Listing Listing
//
// swagger:model Listing
type Listing struct {

	// Chain the listing is on.
	// Required: true
	Chain *string `json:"chain"`

	// Order Hash
	//
	// Order hash
	// Required: true
	OrderHash *string `json:"order_hash"`

	// price
	// Required: true
	Price *BasicListingPrice `json:"price"`

	// Protocol Address
	//
	// Exchange contract address
	// Required: true
	ProtocolAddress *string `json:"protocol_address"`

	// Protocol Data
	//
	// The onchain order data.
	// Required: true
	ProtocolData struct {
		SerializedOrder
	} `json:"protocol_data"`

	// type
	// Required: true
	Type CoreTypesOrderType `json:"type"`
}

// Validate validates this listing
func (m *Listing) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocolAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocolData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Listing) validateChain(formats strfmt.Registry) error {

	if err := validate.Required("chain", "body", m.Chain); err != nil {
		return err
	}

	return nil
}

func (m *Listing) validateOrderHash(formats strfmt.Registry) error {

	if err := validate.Required("order_hash", "body", m.OrderHash); err != nil {
		return err
	}

	return nil
}

func (m *Listing) validatePrice(formats strfmt.Registry) error {

	if err := validate.Required("price", "body", m.Price); err != nil {
		return err
	}

	if m.Price != nil {
		if err := m.Price.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("price")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("price")
			}
			return err
		}
	}

	return nil
}

func (m *Listing) validateProtocolAddress(formats strfmt.Registry) error {

	if err := validate.Required("protocol_address", "body", m.ProtocolAddress); err != nil {
		return err
	}

	return nil
}

func (m *Listing) validateProtocolData(formats strfmt.Registry) error {

	return nil
}

func (m *Listing) validateType(formats strfmt.Registry) error {

	if m.Type == nil {
		return errors.Required("type", "body", nil)
	}

	return nil
}

// ContextValidate validate this listing based on the context it is used
func (m *Listing) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePrice(ctx, formats); err != nil {
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

func (m *Listing) contextValidatePrice(ctx context.Context, formats strfmt.Registry) error {

	if m.Price != nil {
		if err := m.Price.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("price")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("price")
			}
			return err
		}
	}

	return nil
}

func (m *Listing) contextValidateProtocolData(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *Listing) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Listing) UnmarshalBinary(b []byte) error {
	var res Listing
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
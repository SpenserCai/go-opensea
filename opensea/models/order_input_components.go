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

// OrderInputComponents OrderInputComponents
//
// swagger:model OrderInputComponents
type OrderInputComponents struct {

	// Conduitkey
	//
	// Indicates what conduit, if any, should be utilized as a source for token approvals when performing transfers. By default (i.e. when conduitKey is set to the zero hash), the offerer will grant transfer approvals to Seaport directly.
	// To utilize OpenSea's conduit, use 0x0000007b02230091a7ed01230072f7006a004d60a8d4e71d599b8104250f0000
	// Required: true
	ConduitKey *string `json:"conduitKey"`

	// Consideration
	//
	// Array of items which must be received by a recipient to fulfill the order. One of the consideration items must be the OpenSea marketplace fee.
	// Required: true
	Consideration []*ConsiderationItem `json:"consideration"`

	// Counter
	//
	// Must match the current counter for the given offerer. If you are unsure of the current counter, it can be [read from the contract](https://etherscan.io/address/0x00000000000000adc04c56bf30ac9d3c0aaf14dc#readContract#F2) on etherscan.
	// Required: true
	Counter *string `json:"counter"`

	// Endtime
	//
	// The block timestamp at which the order expires.
	// Required: true
	EndTime *int64 `json:"endTime"`

	// Offer
	//
	// Array of items that may be transferred from the offerer's account.
	// Required: true
	Offer []*OfferItem `json:"offer"`

	// Offerer
	//
	// The address which supplies all the items in the offer.
	// Required: true
	Offerer *string `json:"offerer"`

	// order type
	// Required: true
	OrderType OrderType `json:"orderType"`

	// Salt
	//
	// an arbitrary source of entropy for the order
	// Required: true
	Salt *string `json:"salt"`

	// Starttime
	//
	// The block timestamp at which the order becomes active
	// Required: true
	StartTime *int64 `json:"startTime"`

	// Totaloriginalconsiderationitems
	//
	// Size of the consideration array.
	TotalOriginalConsiderationItems int64 `json:"totalOriginalConsiderationItems,omitempty"`

	// Zone
	//
	// Optional secondary account attached the order which can cancel orders. Additionally, when the `OrderType` is Restricted, the zone or the offerer are the only entities which can execute the order.
	// For open orders, use the zero address.
	// For restricted orders, use the signed zone address <SIGNED_ZONE_ADDRESS>
	// Required: true
	Zone *string `json:"zone"`

	// Zonehash
	//
	// A value that will be supplied to the zone when fulfilling restricted orders that the zone can utilize when making a determination on whether to authorize the order. Most often this value will be the zero hash 0x0000000000000000000000000000000000000000000000000000000000000000
	// Required: true
	ZoneHash *string `json:"zoneHash"`
}

// Validate validates this order input components
func (m *OrderInputComponents) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConduitKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConsideration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCounter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOffer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOfferer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSalt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZoneHash(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderInputComponents) validateConduitKey(formats strfmt.Registry) error {

	if err := validate.Required("conduitKey", "body", m.ConduitKey); err != nil {
		return err
	}

	return nil
}

func (m *OrderInputComponents) validateConsideration(formats strfmt.Registry) error {

	if err := validate.Required("consideration", "body", m.Consideration); err != nil {
		return err
	}

	for i := 0; i < len(m.Consideration); i++ {
		if swag.IsZero(m.Consideration[i]) { // not required
			continue
		}

		if m.Consideration[i] != nil {
			if err := m.Consideration[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("consideration" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("consideration" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OrderInputComponents) validateCounter(formats strfmt.Registry) error {

	if err := validate.Required("counter", "body", m.Counter); err != nil {
		return err
	}

	return nil
}

func (m *OrderInputComponents) validateEndTime(formats strfmt.Registry) error {

	if err := validate.Required("endTime", "body", m.EndTime); err != nil {
		return err
	}

	return nil
}

func (m *OrderInputComponents) validateOffer(formats strfmt.Registry) error {

	if err := validate.Required("offer", "body", m.Offer); err != nil {
		return err
	}

	for i := 0; i < len(m.Offer); i++ {
		if swag.IsZero(m.Offer[i]) { // not required
			continue
		}

		if m.Offer[i] != nil {
			if err := m.Offer[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("offer" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("offer" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OrderInputComponents) validateOfferer(formats strfmt.Registry) error {

	if err := validate.Required("offerer", "body", m.Offerer); err != nil {
		return err
	}

	return nil
}

func (m *OrderInputComponents) validateOrderType(formats strfmt.Registry) error {

	if m.OrderType == nil {
		return errors.Required("orderType", "body", nil)
	}

	return nil
}

func (m *OrderInputComponents) validateSalt(formats strfmt.Registry) error {

	if err := validate.Required("salt", "body", m.Salt); err != nil {
		return err
	}

	return nil
}

func (m *OrderInputComponents) validateStartTime(formats strfmt.Registry) error {

	if err := validate.Required("startTime", "body", m.StartTime); err != nil {
		return err
	}

	return nil
}

func (m *OrderInputComponents) validateZone(formats strfmt.Registry) error {

	if err := validate.Required("zone", "body", m.Zone); err != nil {
		return err
	}

	return nil
}

func (m *OrderInputComponents) validateZoneHash(formats strfmt.Registry) error {

	if err := validate.Required("zoneHash", "body", m.ZoneHash); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this order input components based on the context it is used
func (m *OrderInputComponents) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConsideration(ctx, formats); err != nil {
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

func (m *OrderInputComponents) contextValidateConsideration(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Consideration); i++ {

		if m.Consideration[i] != nil {
			if err := m.Consideration[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("consideration" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("consideration" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OrderInputComponents) contextValidateOffer(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Offer); i++ {

		if m.Offer[i] != nil {
			if err := m.Offer[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("offer" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("offer" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderInputComponents) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderInputComponents) UnmarshalBinary(b []byte) error {
	var res OrderInputComponents
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

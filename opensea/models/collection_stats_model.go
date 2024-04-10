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

// CollectionStatsModel CollectionStatsModel
//
// swagger:model CollectionStatsModel
type CollectionStatsModel struct {

	// Average Price
	//
	// The all time average sale price of NFTs in the collection
	// Required: true
	AveragePrice *float64 `json:"average_price"`

	// Floor Price
	//
	// The current lowest price of NFTs in the collection
	// Required: true
	FloorPrice *float64 `json:"floor_price"`

	// Floor Price Symbol
	//
	// The symbol of the payment asset for the floor price
	// Required: true
	FloorPriceSymbol *string `json:"floor_price_symbol"`

	// Market Cap
	//
	// The current market cap of the collection
	// Required: true
	MarketCap *float64 `json:"market_cap"`

	// Num Owners
	//
	// The current number of unique owners of NFTs in the collection
	// Required: true
	NumOwners *int64 `json:"num_owners"`

	// Sales
	//
	// The all time number of sales for the collection
	// Required: true
	Sales *int64 `json:"sales"`

	// Volume
	//
	// The all time volume of sales for the collection
	// Required: true
	Volume *float64 `json:"volume"`
}

// Validate validates this collection stats model
func (m *CollectionStatsModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAveragePrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFloorPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFloorPriceSymbol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMarketCap(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumOwners(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSales(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolume(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CollectionStatsModel) validateAveragePrice(formats strfmt.Registry) error {

	if err := validate.Required("average_price", "body", m.AveragePrice); err != nil {
		return err
	}

	return nil
}

func (m *CollectionStatsModel) validateFloorPrice(formats strfmt.Registry) error {

	if err := validate.Required("floor_price", "body", m.FloorPrice); err != nil {
		return err
	}

	return nil
}

func (m *CollectionStatsModel) validateFloorPriceSymbol(formats strfmt.Registry) error {

	if err := validate.Required("floor_price_symbol", "body", m.FloorPriceSymbol); err != nil {
		return err
	}

	return nil
}

func (m *CollectionStatsModel) validateMarketCap(formats strfmt.Registry) error {

	if err := validate.Required("market_cap", "body", m.MarketCap); err != nil {
		return err
	}

	return nil
}

func (m *CollectionStatsModel) validateNumOwners(formats strfmt.Registry) error {

	if err := validate.Required("num_owners", "body", m.NumOwners); err != nil {
		return err
	}

	return nil
}

func (m *CollectionStatsModel) validateSales(formats strfmt.Registry) error {

	if err := validate.Required("sales", "body", m.Sales); err != nil {
		return err
	}

	return nil
}

func (m *CollectionStatsModel) validateVolume(formats strfmt.Registry) error {

	if err := validate.Required("volume", "body", m.Volume); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this collection stats model based on context it is used
func (m *CollectionStatsModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CollectionStatsModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CollectionStatsModel) UnmarshalBinary(b []byte) error {
	var res CollectionStatsModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
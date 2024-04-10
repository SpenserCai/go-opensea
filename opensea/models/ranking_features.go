// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RankingFeatures RankingFeatures
//
// swagger:model RankingFeatures
type RankingFeatures struct {

	// Unique Attribute Count
	//
	// Deprecated Field.
	UniqueAttributeCount int64 `json:"unique_attribute_count,omitempty"`
}

// Validate validates this ranking features
func (m *RankingFeatures) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ranking features based on context it is used
func (m *RankingFeatures) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RankingFeatures) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RankingFeatures) UnmarshalBinary(b []byte) error {
	var res RankingFeatures
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

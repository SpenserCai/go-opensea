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

// GetCollectionStatsResponse GetCollectionStatsResponse
//
// swagger:model GetCollectionStatsResponse
type GetCollectionStatsResponse struct {

	// Interval Stats
	//
	// The stats for each interval
	// Required: true
	Intervals []*CollectionStatsIntervalModel `json:"intervals"`

	// Total
	//
	// The aggregate stats over the collection's lifetime
	// Required: true
	Total struct {
		CollectionStatsModel
	} `json:"total"`
}

// Validate validates this get collection stats response
func (m *GetCollectionStatsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIntervals(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCollectionStatsResponse) validateIntervals(formats strfmt.Registry) error {

	if err := validate.Required("intervals", "body", m.Intervals); err != nil {
		return err
	}

	for i := 0; i < len(m.Intervals); i++ {
		if swag.IsZero(m.Intervals[i]) { // not required
			continue
		}

		if m.Intervals[i] != nil {
			if err := m.Intervals[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("intervals" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("intervals" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GetCollectionStatsResponse) validateTotal(formats strfmt.Registry) error {

	return nil
}

// ContextValidate validate this get collection stats response based on the context it is used
func (m *GetCollectionStatsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIntervals(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTotal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCollectionStatsResponse) contextValidateIntervals(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Intervals); i++ {

		if m.Intervals[i] != nil {
			if err := m.Intervals[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("intervals" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("intervals" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GetCollectionStatsResponse) contextValidateTotal(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *GetCollectionStatsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCollectionStatsResponse) UnmarshalBinary(b []byte) error {
	var res GetCollectionStatsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Demographics demographics
//
// swagger:model Demographics
type Demographics struct {

	// asian count
	AsianCount int64 `json:"asianCount,omitempty"`

	// black count
	BlackCount int64 `json:"blackCount,omitempty"`

	// geo
	Geo interface{} `json:"geo,omitempty"`

	// hispanic count
	HispanicCount int64 `json:"hispanicCount,omitempty"`

	// nat ind count
	NatIndCount int64 `json:"natIndCount,omitempty"`

	// other count
	OtherCount int64 `json:"otherCount,omitempty"`

	// total count
	TotalCount int64 `json:"totalCount,omitempty"`

	// white count
	WhiteCount int64 `json:"whiteCount,omitempty"`
}

// Validate validates this demographics
func (m *Demographics) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this demographics based on context it is used
func (m *Demographics) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Demographics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Demographics) UnmarshalBinary(b []byte) error {
	var res Demographics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
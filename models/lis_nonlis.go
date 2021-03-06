// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LisNonlis lis nonlis
//
// swagger:model LisNonlis
type LisNonlis struct {

	// average lis claims
	AverageLisClaims int64 `json:"average_lis_claims,omitempty"`

	// average lis drug costs
	AverageLisDrugCosts float64 `json:"average_lis_drug_costs,omitempty"`

	// average nonlis claims
	AverageNonlisClaims int64 `json:"average_nonlis_claims,omitempty"`

	// average nonlis drug costs
	AverageNonlisDrugCosts float64 `json:"average_nonlis_drug_costs,omitempty"`

	// geo
	Geo interface{} `json:"geo,omitempty"`

	// total claims
	TotalClaims int64 `json:"total_claims,omitempty"`

	// total drug costs
	TotalDrugCosts float64 `json:"total_drug_costs,omitempty"`

	// total lis claims
	TotalLisClaims int64 `json:"total_lis_claims,omitempty"`

	// total lis drug costs
	TotalLisDrugCosts float64 `json:"total_lis_drug_costs,omitempty"`

	// total nonlis claims
	TotalNonlisClaims int64 `json:"total_nonlis_claims,omitempty"`

	// total nonlis drug costs
	TotalNonlisDrugCosts float64 `json:"total_nonlis_drug_costs,omitempty"`
}

// Validate validates this lis nonlis
func (m *LisNonlis) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this lis nonlis based on context it is used
func (m *LisNonlis) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LisNonlis) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LisNonlis) UnmarshalBinary(b []byte) error {
	var res LisNonlis
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

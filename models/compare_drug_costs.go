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
)

// CompareDrugCosts compare drug costs
//
// swagger:model CompareDrugCosts
type CompareDrugCosts struct {

	// city
	City string `json:"city,omitempty"`

	// coordinates
	Coordinates []float64 `json:"coordinates"`

	// cost per claim
	CostPerClaim float64 `json:"cost_per_claim,omitempty"`

	// drug name
	DrugName string `json:"drug_name,omitempty"`

	// first name
	FirstName string `json:"first_name,omitempty"`

	// generic name
	GenericName string `json:"generic_name,omitempty"`

	// last org name
	LastOrgName string `json:"last_org_name,omitempty"`

	// npi
	Npi string `json:"npi,omitempty"`

	// physician address
	PhysicianAddress []*PhysicianAddress `json:"physician_address"`

	// state
	State string `json:"state,omitempty"`
}

// Validate validates this compare drug costs
func (m *CompareDrugCosts) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePhysicianAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CompareDrugCosts) validatePhysicianAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.PhysicianAddress) { // not required
		return nil
	}

	for i := 0; i < len(m.PhysicianAddress); i++ {
		if swag.IsZero(m.PhysicianAddress[i]) { // not required
			continue
		}

		if m.PhysicianAddress[i] != nil {
			if err := m.PhysicianAddress[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("physician_address" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this compare drug costs based on the context it is used
func (m *CompareDrugCosts) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePhysicianAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CompareDrugCosts) contextValidatePhysicianAddress(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PhysicianAddress); i++ {

		if m.PhysicianAddress[i] != nil {
			if err := m.PhysicianAddress[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("physician_address" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CompareDrugCosts) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CompareDrugCosts) UnmarshalBinary(b []byte) error {
	var res CompareDrugCosts
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

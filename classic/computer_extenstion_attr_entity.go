// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0
// This product includes software developed at Datadog (https://www.datadoghq.com/). Copyright 2020 Datadog, Inc.

package classic

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// CompterExtensionAttributes represents all attributes that exist in Jamf
type CompterExtensionAttributes struct {
	List []ComputerExtenstionAttribute `json:"computer_extension_attributes"`
}

// ComputerExtensionAttributeDetails holds the details for a single extension attribute
type ComputerExtensionAttributeDetails struct {
	Details ComputerExtenstionAttribute `json:"computer_extension_attribute"`
}

// ComputerExtenstionAttribute represents an extension attribute in Jamf
type ComputerExtenstionAttribute struct {
	XMLName          xml.Name                        `json:"-" xml:"computer_extension_attribute,omitempty"`
	ID               int                             `json:"id" xml:"id,omitempty"`
	Name             string                          `json:"name" xml:"name,omitempty"`
	Enabled          bool                            `json:"enabled" xml:"enabled"` // we don't omit since false values are omitted if needed this can be changed to a *bool
	Description      string                          `json:"description,omitempty" xml:"description,omitempty"`
	DataType         string                          `json:"data_type,omitempty" xml:"data_type,omitempty"`
	InputType        *ComputerExtensionAttrInputType `json:"input_type,omitempty" xml:"input_type,omitempty"`
	InventoryDisplay string                          `json:"inventory_display,omitempty" xml:"inventory_display,omitempty"`
	ReconDisplay     string                          `json:"recon_display,omitempty" xml:"recon_display,omitempty"`
}

// ComputerExtensionAttrInputType represents an input type for a computer extension attribute in Jamf
type ComputerExtensionAttrInputType struct {
	Type     string `json:"type,omitempty" xml:"type,omitempty"`
	Platform string `json:"platform,omitempty" xml:"platform,omitempty"`
	Script   string `json:"script,omitempty" xml:"script,omitempty"`
}

// ValidateComputerExtensionAttribute orchestrates computer extension content validation
func ValidateComputerExtensionAttribute(ce *ComputerExtenstionAttribute) error {
	if err := ce.ValidateDataType(); err != nil {
		return err
	}

	if ce.InputType != nil {
		if err := ce.InputType.ValidateInputType(); err != nil {
			return err
		}
	}

	if err := ce.ValidateReconDisplay(); err != nil {
		return err
	}

	if err := ce.ValidateInventoryDisplay(); err != nil {
		return err
	}

	return nil
}

// ValidateDataType will validate that a computer extension attribute's data type is valid
func (ce *ComputerExtenstionAttribute) ValidateDataType() error {
	switch strings.ToLower(ce.DataType) {
	case "", "string", "integer", "date":
		return nil
	default:
		return fmt.Errorf("%s is not a valid computer extension attribute data type must be of type [ String, Integer, Date ]", ce.DataType)
	}
}

// ValidateInventoryDisplay will validate that a computer extension attribute's data type is valid
func (ce *ComputerExtenstionAttribute) ValidateInventoryDisplay() error {
	switch strings.ToLower(ce.InventoryDisplay) {
	case "", "general", "hardware", "operating system", "user and location", "purchasing", "extension attributes":
		return nil
	default:
		return fmt.Errorf("%s is not a valid computer extension inventory display type must be of type [ General, Hardware, Operating System, User and Location, Purchasing, Extension Attributes ]", ce.InventoryDisplay)
	}
}

// ValidateReconDisplay will validate that a computer extension attribute's data type is valid
func (ce *ComputerExtenstionAttribute) ValidateReconDisplay() error {
	switch strings.ToLower(ce.ReconDisplay) {
	case "", "computer", "user and location", "purchasing", "extension attributes":
		return nil
	default:
		return fmt.Errorf("%s is not a valid computer extension inventory display type must be of type [ Computer, User and Location, Purchasing, Extension Attributes ]", ce.ReconDisplay)
	}
}

// ValidateInputType will validate that a computer extension attribute's input type is valid
func (it *ComputerExtensionAttrInputType) ValidateInputType() error {
	switch it.Type {
	case "", "Text Field", "LDAP Mapping", "Pop-up Menu":
		return nil
	case "script":
		if it.Script == "" {
			return fmt.Errorf("script contents must be provided for input type %s", it.Type)
		}
		return nil
	default:
		return fmt.Errorf("%s is not a valid computer extension attribute input type must be of type [ script, Text Field, LDAP Mapping, Pop-up Menu ]", it.Type)
	}
}

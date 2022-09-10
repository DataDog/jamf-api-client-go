// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0

package classic

import "encoding/xml"

// MobileDevices represents a list of all mobile devices enrolled in Jamf
type MobileDevices struct {
	List  []MobileDevice `json:"mobile_devices" xml:"mobile_devices>mobile_device,omitempty"`
	Count int            `json:"-" xml:"size"`
}

// BasicComputerInfo represents the information returned in a list of all computers from Jamf
type BasicMobileDeviceInfo struct {
	GeneralDeviceInformation
	Username string `json:"username,omitempty" xml:"username,omitempty"`
}

// MobileDevice represents an individual mobile device enrolled in Jamf with all its associated information
type MobileDevice struct {
	Info struct {
		General GeneralDeviceInformation `json:"general"`
	}
}

// GeneralDeviceInformation holds basic information associated with Jamf mobile device
type GeneralDeviceInformation struct {
	XMLName         xml.Name `json:"-" xml:"mobile_device,omitempty"`
	ID              int      `json:"id,omitempty" xml:"id,omitempty"`
	Name            string   `json:"name" xml:"name,omitempty"`
	DeviceName      string   `json:"device_name,omitempty" xml:"device_name,omitempty"`
	UDID            string   `json:"udid,omitempty" xml:"udid,omitempty"`
	SerialNumber    string   `json:"serial_number,omitempty" xml:"serial_number,omitempty"`
	PhoneNumber     string   `json:"phone_number,omitempty" xml:"phone_number,omitempty"`
	WifiMACAddress  string   `json:"wifi_mac_address,omitempty" xml:"wifi_mac_address,omitempty"`
	Supervised      bool     `json:"supervised,omitempty" xml:"supervised,omitempty"`
	Model           string   `json:"model,omitempty" xml:"model,omitempty"`
	ModelIdentifier string   `json:"model_identifier,omitempty" xml:"model_identifier,omitempty"`
	ModelDisplay    string   `json:"model_display,omitempty" xml:"model_display,omitempty"`
}

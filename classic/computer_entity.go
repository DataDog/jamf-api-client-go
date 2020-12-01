// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0
// This product includes software developed at Datadog (https://www.datadoghq.com/). Copyright 2020 Datadog, Inc.

package classic

import "encoding/xml"

// ComputerList represents a list of computers enrolled in Jamf
type ComputerList struct {
	Computers []BasicComputerInfo `json:"computers"`
}

// ComputerGroup represents a group a device is a member of in Jamf
type ComputerGroup struct {
	ID      int    `json:"id,omitempty" xml:"id,omitempty"`
	Name    string `json:"name" xml:"name"`
	IsSmart bool   `json:"is_smart" xml:"is_smart,omitempty"`
}

// BasicComputerInfo represents the information returned in a list of all computers from Jamf
type BasicComputerInfo struct {
	GeneralInformation
}

// Computer represents an individual computer enrolled in Jamf with all its associated information
type Computer struct {
	Info struct {
		General             GeneralInformation       `json:"general"`
		UserLocation        LocationInformation      `json:"location"`
		Hardware            HardwareInformation      `json:"hardware"`
		Certificates        []CertificateInformation `json:"certificates"`
		Software            SoftwareInformation      `json:"software"`
		ExtensionAttributes []ExtensionAttributes    `json:"extension_attributes"`
		Groups              GroupInformation         `json:"groups_accounts"`
		ConfigProfiles      []ConfigProfile          `json:"configuration_profiles"`
	} `json:"computer"`
}

// GeneralInformation holds basic information associated with Jamf device
type GeneralInformation struct {
	XMLName      xml.Name `json:"-" xml:"computer,omitempty"`
	ID           int      `json:"id,omitempty" xml:"id,omitempty"`
	Name         string   `json:"name" xml:"name,omitempty"`
	MACAddress   string   `json:"mac_address" xml:"mac_address,omitempty"`
	SerialNumber string   `json:"serial_number" xml:"serial_number,omitempty"`
	UDID         string   `json:"udid" xml:"udid,omitempty"`
	JamfVersion  string   `json:"jamf_version" xml:"jamf_version,omitempty"`
	Platform     string   `json:"platform" xml:"platform,omitempty"`
	MDMCapable   bool     `json:"mdm_capable" xml:"mdm_capable,omitempty"`
	ReportDate   string   `json:"report_date" xml:"report_date,omitempty"`
}

// LocationInformation holds the information in the User & Locations section
type LocationInformation struct {
	Username     string `json:"username"`
	RealName     string `json:"realname"`
	EmailAddress string `json:"email_address"`
	Position     string `json:"position"`
	Department   string `json:"department"`
	Building     string `json:"building"`
}

// HardwareInformation holds the hardware specific device information
type HardwareInformation struct {
	Make             string   `json:"make"`
	OSName           string   `json:"os_name"`
	OSVersion        string   `json:"os_version"`
	OSBuild          string   `json:"os_build"`
	SIPStatus        string   `json:"sip_status"`
	GatekeeperStatus string   `json:"gatekeeper_status"`
	XProtectVersion  string   `json:"xprotect_version"`
	FilevaultUsers   []string `json:"filevault2_users"`
}

// CertificateInformation holds information about certs intalled on the device
type CertificateInformation struct {
	CommonName string `json:"common_name"`
	Identity   bool   `json:"identity"`
	ExpiresUTC string `json:"expires_utc"`
	Name       string `json:"name"`
}

// SoftwareInformation holds information about the software installed on a device
type SoftwareInformation struct {
	UnixExecutables          []string                 `json:"unix_executables"`
	InstalledByCasper        []string                 `json:"installed_by_casper"`
	InstalledByInstaller     []string                 `json:"installed_by_installer_swu"`
	AvailableSoftwareUpdates []string                 `json:"available_software_updates"`
	RunningServices          []string                 `json:"running_services"`
	Applications             []ApplicationInformation `json:"applications"`
}

// ApplicationInformation holds information about the applications on a device
type ApplicationInformation struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	Version string `json:"version"`
}

// ExtensionAttributes holds extension attribute information for a device
type ExtensionAttributes struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

// GroupInformation holds the groups the device is a member of
type GroupInformation struct {
	Memberships   []string `json:"computer_group_memberships"`
	LocalAccounts []struct {
		Name             string `json:"name"`
		RealName         string `json:"realname"`
		UID              string `json:"uid"`
		Administrator    bool   `json:"administrator"`
		FilevalutEnabled bool   `json:"filevault_enabled"`
	} `json:"local_accounts"`
}

// ConfigProfile represents an active configuration profile in Jamf
type ConfigProfile struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name"`
	UUID      string `json:"uuid"`
	Removable bool   `json:"is_removable"`
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0
// This product includes software developed at Datadog (https://www.datadoghq.com/). Copyright 2020 Datadog, Inc.

package classic

import "encoding/xml"

// Computers represents a list of computers enrolled in Jamf
type Computers struct {
	List []BasicComputerInfo `json:"computers"`
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
	Info ComputerDetails `json:"computer" xml:"computer,omitempty"`
}

type ComputerDetails struct {
	XMLName             xml.Name                 `json:"-" xml:"computer,omitempty"`
	General             GeneralInformation       `json:"general" xml:"general,omitempty"`
	UserLocation        LocationInformation      `json:"location" xml:"location,omitempty"`
	Hardware            HardwareInformation      `json:"hardware" xml:"-"`
	Certificates        []CertificateInformation `json:"certificates" xml:"-"`
	Software            SoftwareInformation      `json:"software" xml:"-"`
	ExtensionAttributes []ExtensionAttributes    `json:"extension_attributes" xml:"extension_attributes,omitempty"`
	Groups              GroupInformation         `json:"groups_accounts" xml:"-"`
	ConfigProfiles      []ConfigProfile          `json:"configuration_profiles" xml:"configuration_profiles,omitempty"`
}

// GeneralInformation holds basic information associated with Jamf device
type GeneralInformation struct {
	ID           int    `json:"id,omitempty" xml:"id,omitempty"`
	Name         string `json:"name" xml:"name,omitempty"`
	MACAddress   string `json:"mac_address" xml:"mac_address,omitempty"`
	SerialNumber string `json:"serial_number" xml:"serial_number,omitempty"`
	UDID         string `json:"udid" xml:"udid,omitempty"`
	JamfVersion  string `json:"jamf_version" xml:"jamf_version,omitempty"`
	Platform     string `json:"platform" xml:"platform,omitempty"`
	MDMCapable   bool   `json:"mdm_capable" xml:"mdm_capable,omitempty"`
	ReportDate   string `json:"report_date" xml:"report_date,omitempty"`
}

// LocationInformation holds the information in the User & Locations section
type LocationInformation struct {
	Username     string `json:"username" xml:"username,omitempty"`
	RealName     string `json:"realname" xml:"realname,omitempty"`
	EmailAddress string `json:"email_address" xml:"email_address,omitempty"`
	Position     string `json:"position" xml:"position,omitempty"`
	Department   string `json:"department" xml:"department,omitempty"`
	Building     string `json:"building" xml:"building,omitempty"`
}

// HardwareInformation holds the hardware specific device information
type HardwareInformation struct {
	Make             string   `json:"make" xml:"make,omitempty"`
	OSName           string   `json:"os_name" xml:"os_name,omitempty"`
	OSVersion        string   `json:"os_version" xml:"os_version,omitempty"`
	OSBuild          string   `json:"os_build" xml:"os_build,omitempty"`
	SIPStatus        string   `json:"sip_status" xml:"sip_status,omitempty"`
	GatekeeperStatus string   `json:"gatekeeper_status" xml:"gatekeeper_status,omitempty"`
	XProtectVersion  string   `json:"xprotect_version" xml:"xprotect_version,omitempty"`
	FilevaultUsers   []string `json:"filevault2_users" xml:"filevault2_users,omitempty"`
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

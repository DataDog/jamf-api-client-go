package classic

import "encoding/xml"

type ComputerGroups struct {
	List []BasicComputerGroupInfo `json:"computer_groups" xml:"computer_groups>computer_group,omitempty"`
	Size int                      `json:"size" xml:"size"`
}

// ComputerGroup represents a group a device is a member of in Jamf
type ComputerGroup struct {
	Info ComputerGroupDetails `json:"computer_group" xml:"computer_group,omitempty"`
}

// BasicComputerGroupInfo represents the information returned in a list of all
// computer groups from Jamf
type BasicComputerGroupInfo struct {
	ID      int    `json:"id,omitempty" xml:"id,omitempty"`
	Name    string `json:"name,omitempty" xml:"name"`
	IsSmart bool   `json:"is_smart" xml:"is_smart"`
}

// ComputerGroupDetails represents the detailed information for a specific computer group
type ComputerGroupDetails struct {
	XMLName xml.Name `json:"computer_group" xml:"computer_group,omitempty"`
	BasicComputerGroupInfo
	Computers []BasicComputerInfo `json:"computers" xml:"computers>computer,omitempty"`
}

// ComputerGroupBindingChanges represents the changes to a computer group binding when
// updating the members of a computer group in Jamf
type ComputerGroupBindingChanges struct {
	XMLName   xml.Name             `json:"-" xml:"computer_group,omitempty"`
	Additions []GeneralInformation `xml:"computer_additions>computer"`
	Removals  []GeneralInformation `xml:"computer_deletions>computer"`
}

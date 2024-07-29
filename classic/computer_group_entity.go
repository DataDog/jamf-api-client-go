package classic

type Groups struct {
	List []BasicComputerGroupInformation `json:"computer_groups" xml:"computer_groups>computer_group,omitempty"`
}

// BasicComputerGroupInformation holds the basic information for all groups in Jamf
type BasicComputerGroupInformation struct {
	GeneralGroupInformation
}
type GeneralGroupInformation struct {
	ID      int    `json:"id,omitempty" xml:"id,omitempty"`
	Name    string `json:"name" xml:"name,omitempty"`
	IsSmart bool   `json:"is_smart" xml:"is_smart,omitempty"`
}
type ComputerGroupDetails struct {
	ID        int                 `json:"id,omitempty" xml:"id,omitempty"`
	Name      string              `json:"name" xml:"name"`
	IsSmart   bool                `json:"is_smart" xml:"is_smart,omitempty"`
	Computers []BasicComputerInfo `json:"computers" xml:"computers>computer,omitempty"`
}

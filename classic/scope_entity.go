// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0
// This product includes software developed at Datadog (https://www.datadoghq.com/). Copyright 2020 Datadog, Inc.

package classic

// Scope represents the scope of a related Jamf configuration setting or Policy
type Scope struct {
	AllComputers   bool                      `json:"all_computers" xml:"all_computers,omitempty"`
	Computers      []*BasicComputerInfo      `json:"computers" xml:"computers>computer,omitempty"`
	ComputerGroups []*BasicComputerGroupInfo `json:"computer_groups" xml:"computer_groups>computer_group,omitempty"`
	Buildings      []*Building               `json:"buildings" xml:"buildings,omitempty"`
	Departments    []*Department             `json:"departments" xml:"departments,omitempty"`
	LimitToUsers   *UserGroupLimitations     `json:"limit_to_users" xml:"limit_to_users,omitempty"`
	Limitations    *Limitations              `json:"limitations" xml:"limitations,omitempty"`
	Exclusions     *Exclusions               `json:"exclusions" xml:"exclusions,omitempty"`
}

// Building represents a building configured in Jamf that a setting can be scoped to
type Building struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name"`
}

// Department represents a department configured in Jamf that a setting can be scoped to
type Department struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name"`
}

// User represents a user configured in Jamf that a setting can be scoped to
type User struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name"`
}

// UserGroupLimitations represents the user groups to limit a scope to
type UserGroupLimitations struct {
	UserGroups []*UserGroup `json:"user_groups"`
}

// UserGroup represents a user group configured in Jamf that a setting can be scoped to
type UserGroup struct {
	Size int               `json:"size"`
	Info *UserGroupDetails `json:"user_group"`
}

// UserGroupDetails holds the specific details of a user group
type UserGroupDetails struct {
	ID             int    `json:"id,omitempty"`
	Name           string `json:"name"`
	IsSmart        bool   `json:"is_smart"`
	NotifyOnChange bool   `json:"is_notify_on_change"`
}

// NetworkSegment represents a network segment configured in Jamf that a setting can be scoped to
type NetworkSegment struct {
	ID              int    `json:"id,omitempty"`
	Name            string `json:"name"`
	StartingAddress string `json:"starting_address"`
	EndingAddress   string `json:"ending_address"`
}

// Limitations represents any limitations related to the specific scope
type Limitations struct {
	Users           []*User           `json:"users,omitempty"`
	UserGroups      []*UserGroup      `json:"user_groups,omitempty"`
	NetworkSegments []*NetworkSegment `json:"network_segments"`
}

// Exclusions represents any exclusions applied to the scoping of the Jamf setting in context
type Exclusions struct {
	Computers       []*BasicComputerInfo `json:"computers"`
	ComputerGroups  []*ComputerGroup     `json:"computer_groups"`
	Buildings       []*Building          `json:"buildings"`
	Departments     []*Department        `json:"departments"`
	Users           []*User              `json:"users"`
	UserGroups      []*UserGroup         `json:"user_groups"`
	NetworkSegments []*NetworkSegment    `json:"network_segments"`
}

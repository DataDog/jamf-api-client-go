// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0

package classic

import "encoding/xml"

// Classes represents a list of mobile device classes in Jamf
type Classes struct {
	List  []Class `json:"classes" xml:"classes>class,omitempty"`
	Count int     `json:"-" xml:"size"`
}

// ClassDetails holds the details for a single mobile device class
type ClassDetails struct {
	Details *Class `json:"class"`
}

// Class represents an individual mobile device class in Jamf with all its associated information
type Class struct {
	XMLName       xml.Name                `json:"-" xml:"class,omitempty"`
	ID            int                     `json:"id,omitempty" xml:"id,omitempty"`
	Source        string                  `json:"source,omitempty" xml:"source,omitempty"`
	Name          string                  `json:"name" xml:"name,omitempty"`
	Description   string                  `json:"description,omitempty" xml:"description,omitempty"`
	Site          Site                    `json:"site,omitempty" xml:"site,omitempty"`
	Students      []string                `json:"students,omitempty" xml:"students>student,omitempty"`
	Teachers      []string                `json:"teachers,omitempty" xml:"teachers>teacher,omitempty"`
	MobileDevices []BasicMobileDeviceInfo `json:"mobile_devices,omitempty" xml:"mobile_devices,omitempty"`
	MeetingTimes  []MeetingTime           `json:"meeting_times,omitempty" xml:"meeting_times,omitempty"`
}

// MeetingTime holds values for a mobile device class meeting time
type MeetingTime struct {
	Days      string `json:"days,omitempty" xml:"days,omitempty"`
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	EndTime   string `json:"end_time,omitempty" xml:"end_time,omitempty"`
}

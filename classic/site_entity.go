// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0

package classic

import "encoding/xml"

// Sites holds a list of sites configured in Jamf
type Sites struct {
	List  []Site `json:"sites" xml:"sites>site,omitempty"`
	Count int    `json:"-" xml:"size"`
}

// Site holds the details of a site configured in Jamf
type Site struct {
	XMLName xml.Name `json:"-" xml:"site,omitempty"`
	ID      int      `json:"id,omitempty" xml:"id,omitempty"`
	Name    string   `json:"name" xml:"name,omitempty"`
}

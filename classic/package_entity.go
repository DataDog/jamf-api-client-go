// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0
// This product includes software developed at Datadog (https://www.datadoghq.com/). Copyright $year Datadog, Inc.

package classic

import "encoding/xml"

// Packages holds a list of package details
type Packages struct {
	List []*Package `json:"packages" xml:"packages>package,omitempty"`
}

// Package holds the details of a package configured in Jamf
type Package struct {
	XMLName       xml.Name `json:"-" xml:"package,omitempty"`
	ID            int      `json:"id,omitempty" xml:"id,omitempty"`
	Name          string   `json:"name" xml:"name,omitempty"`
	Action        string   `json:"action" xml:"action,omitempty"`
	FUT           bool     `json:"fut" xml:"fut,omitempty"`
	FEU           bool     `json:"feu" xml:"feu,omitempty"`
	UpdateAutorun bool     `json:"update_autorun" xml:"update_autorun,omitempty"`
}

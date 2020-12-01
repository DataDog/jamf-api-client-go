// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0
// This product includes software developed at Datadog (https://www.datadoghq.com/). Copyright 2020 Datadog, Inc.

package classic

// Printer represents a printer configured in Jamf
// type Printer struct {
// 	PrinterDetails `json:"printer"`
// }

// PrinterDetails holds the details to a printer configured in Jamf
type PrinterDetails struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Action      string `json:"action"`
	MakeDefault bool   `json:"make_default"`
}

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

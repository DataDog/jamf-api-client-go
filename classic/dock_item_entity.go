// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0
// This product includes software developed at Datadog (https://www.datadoghq.com/). Copyright $year Datadog, Inc.

package classic

// DockItem represents a dock item configured in Jamf typically part of a policy
type DockItem struct {
	Size    int              `json:"size"`
	Details *DockItemDetails `json:"dock_item"`
}

// DockItemDetails holds the details for a configured dock item
type DockItemDetails struct {
	ID     int    `json:"id,omitempty"`
	Name   string `json:"name"`
	Action string `json:"action"`
}

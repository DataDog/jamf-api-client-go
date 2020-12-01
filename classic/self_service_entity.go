// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0
// This product includes software developed at Datadog (https://www.datadoghq.com/). Copyright 2020 Datadog, Inc.

package classic

// SelfService represents a self service configuration in Jamf i.e policy self service config
type SelfService struct {
	Enabled              bool                   `json:"user_for_self_service"`
	DisplayName          string                 `json:"self_service_display_name"`
	InstallBtnText       string                 `json:"install_button_text"`
	ReInstallBtnText     string                 `json:"reinstall_button_text"`
	Description          string                 `json:"self_service_description"`
	ForceDescriptionView bool                   `json:"force_users_to_view_description"`
	Icon                 *SelfServiceIcon       `json:"self_service_icon"`
	MainPageFeature      bool                   `json:"feature_on_main_page"`
	Categories           []*SelfServiceCategory `json:"self_service_categories"`
	Notification         string                 `json:"notification"`
	NotificationSubject  string                 `json:"notification_subject"`
	NotificationMessage  string                 `json:"notification_message"`
}

// SelfServiceIcon holds the config for a self service icon associated with a policy
type SelfServiceIcon struct {
	ID       int    `json:"id,omitempty"`
	Filename string `json:"filename"`
	URI      string `json:"uri"`
}

// SelfServiceCategory holds the category associated with a policy
type SelfServiceCategory struct {
	Category struct {
		ID        int    `json:"id,omitempty"`
		Name      string `json:"name"`
		DisplayIn bool   `json:"display_in"`
		FeatureIn bool   `json:"feature_in"`
	} `json:"category"`
}

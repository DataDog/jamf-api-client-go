// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0
// This product includes software developed at Datadog (https://www.datadoghq.com/). Copyright 2020 Datadog, Inc.

package classic

import "encoding/xml"

// PolicyList holds all policies in the configured Jamf environment
type PolicyList struct {
	Policies []BasicPolicyInformation `json:"policies"`
}

// BasicPolicyInformation holds the basic information for all policies in Jamf
type BasicPolicyInformation struct {
	XMLName xml.Name `json:"-" xml:"policy,omitempty"`
	ID      int      `json:"id,omitempty" xml:"id,omitempty"`
	Name    string   `json:"name" xml:"name,omitempty"`
}

// Policy represents a single policy construct in Jamf
type Policy struct {
	Content *PolicyContents `json:"policy" xml:"policy"`
}

// PolicyContents represents the details associated with a given Jamf policy
type PolicyContents struct {
	XMLName              xml.Name                  `json:"-" xml:"policy,omitempty"`
	General              *PolicyGeneral            `json:"general" xml:"general,omitempty"`
	Scope                *Scope                    `json:"scope" xml:"scope,omitempty"`
	SelfServices         *SelfService              `json:"self_service" xml:"self_service,omitempty"`
	PackageConfiguration *Packages                 `json:"package_configuration" xml:"package_configuration,omitempty"`
	ScriptCount          int                       `json:"-"  xml:"scripts>size,omitempty"`
	Scripts              []*PolicyScriptAssignment `json:"scripts" xml:"scripts>script,omitempty"`
	Printers             interface{}               `json:"printers" xml:"printers,omitempty"`
	DockItems            []*DockItem               `json:"dock_items" xml:"dock_items,omitempty"`
	AccountMaintenance   *PolicyAccountMaintenance `json:"account_maintenance" xml:"account_maintenance,omitempty"`
	RebootSettings       *PolicyRebootSettings     `json:"reboot" xml:"reboot,omitempty"`
	Maintenance          *PolicyMaintenance        `json:"maintenance" xml:"maintenance,omitempty"`
	FilesProcesses       *PolicyFileProcesses      `json:"files_processes" xml:"files_processes,omitempty"`
	UserInteraction      *PolicyUserInteraction    `json:"user_interaction" xml:"user_interaction,omitempty"`
	DiskEncryption       *PolicyDiskEncryption     `json:"disk_encryption" xml:"disk_encryption,omitempty"`
}

// PolicyGeneral holds all the generic policy info
type PolicyGeneral struct {
	XMLName                   xml.Name                  `json:"-" xml:"general,omitempty"`
	ID                        int                       `json:"id,omitempty" xml:"id,omitempty"`
	Name                      string                    `json:"name" xml:"name,omitempty"`
	Enabled                   bool                      `json:"enabled" xml:"enabled,omitempty"`
	Trigger                   string                    `json:"trigger" xml:"trigger,omitempty"`
	TriggerCheckIn            bool                      `json:"trigger_checkin" xml:"trigger_checkin,omitempty"`
	TriggerEnrollmentComplete bool                      `json:"trigger_enrollment_comlete" xml:"trigger_enrollment_complete,omitempty"`
	TriggerLogin              bool                      `json:"trigger_login" xml:"trigger_login,omitempty"`
	TriggerLogout             bool                      `json:"trigger_logout" xml:"trigger_logout,omitempty"`
	TriggerNetworkStateChange bool                      `json:"trigger_network_state_changed" xml:"trigger_network_state_changed,omitempty"`
	TriggerStartup            bool                      `json:"trigger_startup" xml:"trigger_startup,omitempty"`
	TriggerOther              string                    `json:"trigger_other" xml:"trigger_other,omitempty"`
	Frequency                 string                    `json:"frequency" xml:"frequency,omitempty"`
	RetryEvent                string                    `json:"retry_event" xml:"retry_event,omitempty"`
	RetryAttempts             int                       `json:"retry_attempts" xml:"retry_attempts,omitempty"`
	NotifyOnFailedRetry       bool                      `json:"notify_on_each_failed_retry" xml:"notify_on_each_failed_retry,omitempty"`
	LocationUserOnly          bool                      `json:"location_user_only" xml:"location_user_only,omitempty"`
	TargetDrive               string                    `json:"target_drive" xml:"target_drive,omitempty"`
	Offline                   bool                      `json:"offline" xml:"offline,omitempty"`
	NetworkRequirements       string                    `json:"network_requirements" xml:"network_requirements,omitempty"`
	Category                  *PolicyCategory           `json:"category" xml:"category,omitempty"`
	DateTimeLimitations       *PolicyDateLimitations    `json:"date_time_limitations" xml:"date_time_limitations,omitempty"`
	NetworkLimitations        *PolicyNetworkLimitations `json:"network_limitations" xml:"network_limitations,omitempty"`
	OverrideDefaultSettings   *PolicyOverrides          `json:"override_default_settings" xml:"override_default_settings,omitempty"`
	Site                      *PolicySite               `json:"site" xml:"site,omitempty"`
}

// PolicyCategory is a policy category
type PolicyCategory struct {
	ID   int    `json:"id,omitempty" xml:"id,omitempty"`
	Name string `json:"name" xml:"name"`
}

// PolicySite holds the site configuration for a policy
type PolicySite struct {
	ID   int    `json:"id,omitempty" xml:"id,omitempty"`
	Name string `json:"name" xml:"name"`
}

// PolicyScriptAssignment holds the metadata related to a script assigned to a policy
type PolicyScriptAssignment struct {
	ID          int    `json:"id,omitempty" xml:"id,omitempty"`
	Name        string `json:"name" xml:"name,omitempty"`
	Priority    string `json:"priority" xml:"priority,omitempty"`
	Parameter4  string `json:"parameter4" xml:"parameter4,omitempty"`
	Parameter5  string `json:"parameter5" xml:"parameter5,omitempty"`
	Parameter6  string `json:"parameter6" xml:"parameter6,omitempty"`
	Parameter7  string `json:"parameter7" xml:"parameter7,omitempty"`
	Parameter8  string `json:"parameter8" xml:"parameter8,omitempty"`
	Parameter9  string `json:"parameter9" xml:"parameter9,omitempty"`
	Parameter10 string `json:"parameter10" xml:"parameter10,omitempty"`
	Parameter11 string `json:"parameter11" xml:"parameter11,omitempty"`
}

// PolicyNetworkLimitations holds the network limitations associated with a policy
type PolicyNetworkLimitations struct {
	MinimumNetworkConnection string   `json:"minimum_network_connection" xml:"minimum_network_connection,omitempty"`
	AnyIPAddress             bool     `json:"any_ip_address" xml:"any_ip_address,omitempty"`
	NetworkSegments          []string `json:"network_segments" xml:"network_segments,omitempty"`
}

// PolicyOverrides contains overrides for the policy's default config
type PolicyOverrides struct {
	TargetDrive       string `json:"target_drive" xml:"target_drive,omitempty"`
	DistributionPoint string `json:"distribution_point" xml:"distribution_point,omitempty"`
	ForceAFPSMB       bool   `json:"force_afp_smb" xml:"force_afp_smb,omitempty"`
	SUS               string `json:"sus" xml:"sus,omitempty"`
	NetbootServer     string `json:"netboot_server" xml:"netboot_server,omitempty"`
}

// PolicyDateLimitations holds the date/time related config for the policy
type PolicyDateLimitations struct {
	ActivationDate      string `json:"activation_date" xml:"activation_date,omitempty"`
	ActivationDateEPOCH int    `json:"activation_date_epoch" xml:"activation_date_epoch,omitempty"`
	ActivationDateUTC   string `json:"activation_date_utc" xml:"activation_date_utc,omitempty"`
	ExpirationDate      string `json:"expiration_date" xml:"expiration_date,omitempty"`
	ExpirationDateEPOCH int    `json:"expiration_date_epoch" xml:"expiration_date_epoch,omitempty"`
	ExpirationDateUTC   string `json:"expiration_date_utc" xml:"expiration_date_utc,omitempty"`
	NoExecuteOn         struct {
		Day string `json:"day,omitempty" xml:"day,omitempty"`
	} `json:"no_execute_on" xml:"no_execute_on,omitempty"`
	NoExecuteStart string `json:"no_execute_start" xml:"no_execute_start,omitempty"`
	NoExecuteEnd   string `json:"no_execute_end" xml:"no_execute_end,omitempty"`
}

// PolicyAccountMaintenance holds information about account changes controlled by this policy
type PolicyAccountMaintenance struct {
	Account                 []*Account         `json:"accounts"`
	DirectoryBindings       interface{}        `json:"directory_bindings"`
	ManagementAccount       *ManagementAccount `json:"management_account"`
	OpenFirmwareEFIPassword interface{}        `json:"open_firmware_efi_password"`
}

// PolicyRebootSettings stores information about how this policy handles reboots
type PolicyRebootSettings struct {
	Message                     string `json:"message"`
	StartupDisk                 string `json:"startup_disk"`
	SpecifyStartup              string `json:"specify_startup"`
	NoUserLoggedIn              string `json:"no_user_logged_in"`
	UserLoggedIn                string `json:"user_logged_in"`
	MinutesUntilReboot          int    `json:"minutes_until_reboot"`
	StartRebootTimerImmediately bool   `json:"start_reboot_timer_immediately"`
	FileVaultReboot             bool   `json:"file_value_2_reboot"`
}

// PolicyMaintenance defines how jamf handles this policy long term
type PolicyMaintenance struct {
	Recon                    bool `json:"recon"`
	ResetName                bool `json:"reset_name"`
	InstallAllCachedPackages bool `json:"install_all_cached_packages"`
	Heal                     bool `json:"heal"`
	PreBindings              bool `json:"prebindings"`
	Permissons               bool `json:"permissions"`
	ByHost                   bool `json:"byhost"`
	SystemCache              bool `json:"system_cache"`
	UserCache                bool `json:"user_cache"`
	Verify                   bool `json:"verify"`
}

// PolicyFileProcesses holds information about the files processed when this policy is executed
type PolicyFileProcesses struct {
	SearchPatch      string `json:"search_by_path"`
	DeleteFile       bool   `json:"delete_file"`
	LocateFile       string `json:"locate_file"`
	UpdateLocateDB   bool   `json:"update_locate_database"`
	SpotlightSearch  string `json:"spotlight_search"`
	SearchFroProcess string `json:"search_for_process"`
	KillProcess      bool   `json:"kill_process"`
	RunCommand       string `json:"run_command"`
}

// PolicyUserInteraction holds the settings associated with user interaction when the policy runs
type PolicyUserInteraction struct {
	MessageStart           string `json:"message_start"`
	MessageFinish          string `json:"message_finish"`
	AllowUserDefer         bool   `json:"allow_user_to_defer"`
	AllowUserDeferUntilUTC string `json:"allow_deferral_until_utc"`
	AllowUSerDeferMinutes  int    `json:"allow_deferral_minutes"`
}

// PolicyDiskEncryption holds information about disk encryption settings when executed
type PolicyDiskEncryption struct {
	Action                       string `json:"action"`
	DiskEncryptionConfigID       int    `json:"disk_encryption_configuration_id"`
	AuthRestart                  bool   `json:"auth_restart"`
	RemediateKeyType             string `json:"remediate_key_type"`
	RemediateDiskEncryptConfigID int    `json:"remediate_disk_encryption_configuration_id"`
}

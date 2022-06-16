// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0
// This product includes software developed at Datadog (https://www.datadoghq.com/). Copyright 2020 Datadog, Inc.
package classic_test

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	jamf "github.com/DataDog/jamf-api-client-go/classic"
	"github.com/stretchr/testify/assert"
)

var COMPUTER_API_BASE_ENDPOINT = "/JSSResource/computers"

func computerResponseMocks(t *testing.T) *httptest.Server {
	var resp string
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.RequestURI {
		case COMPUTER_API_BASE_ENDPOINT:
			fmt.Fprintf(w, `{
				"computers": [
					{
							"id": 3,
							"name": "Test MacBook #3"
					},
					{
							"id": 28,
							"name": "Test MacBook #28"
					},
					{
							"id": 30,
							"name": "Test MacBook #30"
					},
					{
							"id": 31,
							"name": "Test MacBook #31"
					},
					{
							"id": 32,
							"name": "Test MacBook #32"
					},
					
					{
							"id": 91,
							"name": "Test MacBook #91"
					}]
			}`)
		case fmt.Sprintf("%s/id/82", COMPUTER_API_BASE_ENDPOINT):
			fmt.Fprintf(w, `{
				"computer": {
					"general": {
						"id": 82,
						"name": "Go Client Test Machine",
						"mac_address": "00:00:00:A0:FE:00",
						"serial_number": "VM0L+J/0cr+l",
						"udid": "000DF0BF-00FF-D00B-FA00-000F0DA0FE00",
						"jamf_version": "20.18.0-t0000000000",
						"platform": "Mac",
						"mdm_capable": false,
						"report_date": "2020-09-11 23:06:00",
						"ip_address": "192.0.2.100",
						"last_reported_ip": "192.0.2.101"
					},
					"location": {
						"username": "test.user",
						"realname": "Test User",
						"real_name": "Test User",
						"email_address": "test.user@email.com",
						"position": "Software Engineer",
						"department": "Engineering",
						"building": "Boston"
					},
					"hardware": {
						"make": "Apple",
						"os_name": "Mac OS X",
						"os_version": "10.14.4",
						"os_build": "18E2034",
						"sip_status": "Enabled",
						"gatekeeper_status": "App Store and identified developers",
						"xprotect_version": "2114",
						"filevault2_users": [
							"test.user"
						]
					},
					"certificates": [{
						"common_name": "JSS Built-in Certificate Authority",
						"identity": false,
						"expires_utc": "9027-11-12T20:07:28.000+0000",
						"expires_epoch": 1826050048000,
						"name": ""
					}],
					"software": {
						"unix_executables": [],
						"licensed_software": [],
						"installed_by_casper": [
							"filevault_profile_signed.pkg",
							"Zoom-Latest.pkg"
						],
							"installed_by_installer_swu": [
								"com.datadoghq.pkg.Zoom-Latest",
								"com.github.makeprofilepkg.config_us",
								"com.googlecode.munki.admin",
								"com.googlecode.munki.app",
								"com.googlecode.munki.app_usage",
								"com.googlecode.munki.core",
								"com.googlecode.munki.launchd",
								"com.googlecode.munki.python",
								"com.vmware.tools.macos.pkg.files"
							],
							"available_software_updates": [],
							"available_updates": {},
							"running_services": [
								"com.apple.accessoryd",
								"com.apple.backupd",
								"com.apple.diagnosticd",
								"com.apple.mdmclient.daemon",
								"com.apple.mdmclient.daemon.runatboot",
								"com.googlecode.munki.appusaged",
								"com.jamf.management.daemon"
							],
							"applications": [{
								"name": "Datadog Agent.app",
								"path": "/Applications/Datadog Agent.app",
								"version": "7.16.1"
							}]
						},
						"extension_attributes": [
							{
								"id": 6,
								"name": "osquery Status",
								"type": "String",
								"multi_value": false,
								"value": "OSquery NOT Running"
							}
						],
						"groups_accounts": {
							"computer_group_memberships": [
								"Test Group for API Client"
							],
							"local_accounts": [{
								"name": "test.user",
								"realname": "Test User",
								"uid": "501",
								"administrator": true,
								"filevault_enabled": true
							}]
						},
						"configuration_profiles": [{
							"id": 2,
							"name": "Test Config Profile",
							"uuid": "abcdefghijklmnop123",
							"is_removable": false
						}]
				}
			}`)
		case fmt.Sprintf("%s/serialnumber/VM0L+J/0cr+l", COMPUTER_API_BASE_ENDPOINT):
			switch r.Method {
			case "GET":
				fmt.Fprintf(w, `{
					"computer": {
						"general": {
							"id": 82,
							"name": "Test Machine (Serial Number)",
							"mac_address": "00:00:00:A0:FE:00",
							"serial_number": "VM0L+J/0cr+l",
							"udid": "000DF0BF-00FF-D00B-FA00-000F0DA0FE00",
							"jamf_version": "20.18.0-t0000000000",
							"platform": "Mac",
							"mdm_capable": false,
							"report_date": "2020-09-11 23:06:00"
						}
					}
				}`)
			case "PUT", "POST":
				data, err := ioutil.ReadAll(r.Body)
				if err != nil {
					fmt.Fprint(w, err.Error())
				}
				contents := &jamf.ComputerDetails{}
				err = xml.Unmarshal(data, contents)
				if err != nil {
					fmt.Fprint(w, err.Error())
				}
				resp, err := json.MarshalIndent(contents, "", "    ")
				if err != nil {
					fmt.Fprint(w, err.Error())
				}
				fmt.Fprint(w, string(resp))
			}

		default:
			http.Error(w, fmt.Sprintf("bad Jamf computer API call to %s", r.URL), http.StatusInternalServerError)
			return
		}
		_, err := w.Write([]byte(resp))
		assert.Nil(t, err)
	}))
}
func TestListComputers(t *testing.T) {
	testServer := computerResponseMocks(t)
	defer testServer.Close()
	j, err := jamf.NewClient(testServer.URL, "fake-username", "mock-password-cool", nil)
	assert.Nil(t, err)
	computers, err := j.Computers()
	assert.Nil(t, err)
	assert.NotNil(t, computers)
	assert.Equal(t, 6, len(computers))
	assert.Equal(t, 3, computers[0].ID)
	assert.Equal(t, "Test MacBook #3", computers[0].Name)
}

func TestQuerySpecificComputer(t *testing.T) {
	testServer := computerResponseMocks(t)
	defer testServer.Close()
	j, err := jamf.NewClient(testServer.URL, "fake-username", "mock-password-cool", nil)
	assert.Nil(t, err)
	computer, err := j.ComputerDetails(82)
	assert.Nil(t, err)
	// General Info
	assert.Equal(t, 82, computer.Info.General.ID)
	assert.Equal(t, "Go Client Test Machine", computer.Info.General.Name)
	assert.Equal(t, false, computer.Info.General.MDMCapable)
	assert.Equal(t, "192.0.2.100", computer.Info.General.IPAddress)
	assert.Equal(t, "192.0.2.101", computer.Info.General.LastReportedIP)

	// User & Location Info
	assert.Equal(t, "Test User", computer.Info.UserLocation.RealName)
	assert.Equal(t, "test.user@email.com", computer.Info.UserLocation.EmailAddress)
	assert.Equal(t, "Software Engineer", computer.Info.UserLocation.Position)
	assert.Equal(t, "Engineering", computer.Info.UserLocation.Department)

	// Hardware info
	assert.Equal(t, "Apple", computer.Info.Hardware.Make)
	assert.Equal(t, "App Store and identified developers", computer.Info.Hardware.GatekeeperStatus)
	assert.Equal(t, "Enabled", computer.Info.Hardware.SIPStatus)
	assert.Equal(t, []string{"test.user"}, computer.Info.Hardware.FilevaultUsers)

	// Certificate Information
	assert.Equal(t, "JSS Built-in Certificate Authority", computer.Info.Certificates[0].CommonName)

	// Software Information
	assert.Equal(t, []string{"filevault_profile_signed.pkg", "Zoom-Latest.pkg"}, computer.Info.Software.InstalledByCasper)
	assert.Equal(t, "com.apple.accessoryd", computer.Info.Software.RunningServices[0])
	assert.Equal(t, "Datadog Agent.app", computer.Info.Software.Applications[0].Name)

	// Extension Attributes
	assert.Equal(t, 6, computer.Info.ExtensionAttributes[0].ID)
	assert.Equal(t, "osquery Status", computer.Info.ExtensionAttributes[0].Name)
	assert.Equal(t, "OSquery NOT Running", computer.Info.ExtensionAttributes[0].Value)

	// Group Memberships & Local Accounts
	assert.Equal(t, []string{"Test Group for API Client"}, computer.Info.Groups.Memberships)
	assert.Equal(t, "test.user", computer.Info.Groups.LocalAccounts[0].Name)
	assert.Equal(t, true, computer.Info.Groups.LocalAccounts[0].Administrator)

	// Config Profiles
	assert.Equal(t, 2, computer.Info.ConfigProfiles[0].ID)
	assert.Equal(t, "Test Config Profile", computer.Info.ConfigProfiles[0].Name)
	assert.Equal(t, false, computer.Info.ConfigProfiles[0].Removable)
}

func TestGetComputer__ID(t *testing.T) {
	testServer := computerResponseMocks(t)
	defer testServer.Close()
	j, err := jamf.NewClient(testServer.URL, "fake-username", "mock-password-cool", nil)
	assert.Nil(t, err)

	id := &jamf.ComputerIdentifier{
		ID: "82",
	}

	computer, err := j.GetComputer(id)
	assert.Nil(t, err)
	// General Info
	assert.Equal(t, 82, computer.Info.General.ID)
}

func TestGetComputer__SerialNumber(t *testing.T) {
	testServer := computerResponseMocks(t)
	defer testServer.Close()
	j, err := jamf.NewClient(testServer.URL, "fake-username", "mock-password-cool", nil)
	assert.Nil(t, err)

	id := &jamf.ComputerIdentifier{
		SerialNumber: "VM0L+J/0cr+l",
	}

	computer, err := j.GetComputer(id)
	assert.Nil(t, err)
	// General Info
	assert.Equal(t, 82, computer.Info.General.ID)
	assert.Equal(t, "Test Machine (Serial Number)", computer.Info.General.Name)
}

func TestUpdateComputer__SerialNumber(t *testing.T) {
	testServer := computerResponseMocks(t)
	defer testServer.Close()
	j, err := jamf.NewClient(testServer.URL, "fake-username", "mock-password-cool", nil)
	assert.Nil(t, err)

	update := &jamf.ComputerDetails{
		General: jamf.GeneralInformation{
			Name: "Updated_Computer",
		},
		UserLocation: jamf.LocationInformation{
			EmailAddress: "test@email.com",
		},
	}

	id := &jamf.ComputerIdentifier{
		SerialNumber: "VM0L+J/0cr+l",
	}

	comp, err := j.UpdateComputer(id, update)
	assert.Nil(t, err)
	assert.Equal(t, "Updated_Computer", comp.General.Name)
	assert.Equal(t, "test@email.com", comp.UserLocation.EmailAddress)
}

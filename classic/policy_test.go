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

var POLICIES_API_BASE_ENDPOINT = "/JSSResource/policies"

func policiesResponseMocks(t *testing.T) *httptest.Server {
	var resp string
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.RequestURI {
		case POLICIES_API_BASE_ENDPOINT:
			fmt.Fprintf(w, `{
				"policies": [
					{
							"id": 70,
							"name": "0 - SplashBuddy"
					},
					{
							"id": 71,
							"name": "01 - macOS-Upgrader - DEP"
					},
					{
							"id": 72,
							"name": "Test Policy"
					},
					{
							"id": 73,
							"name": "03 - 1Password 7 auto-install - DEP"
					},
					{
							"id": 74,
							"name": "04 - Slack auto-install - DEP"
					}
				]
			}`)
		case fmt.Sprintf("%s/id/72", POLICIES_API_BASE_ENDPOINT), fmt.Sprintf("%s/id/-1", POLICIES_API_BASE_ENDPOINT), fmt.Sprintf("%s/name/Test%sPolicy", POLICIES_API_BASE_ENDPOINT, "%20"):
			switch r.Method {
			case "PUT", "POST":
				w.Header().Add("Content-Type", "application/xml")
				data, err := ioutil.ReadAll(r.Body)
				if err != nil {
					fmt.Fprintf(w, err.Error())
				}

				policyContents := &jamf.PolicyContents{}
				err = xml.Unmarshal(data, policyContents)
				if err != nil {
					fmt.Fprintf(w, err.Error())
				}

				policyData, err := xml.MarshalIndent(policyContents, "", "    ")
				if err != nil {
					fmt.Fprintf(w, err.Error())
				}
				fmt.Fprintf(w, string(policyData))
			default:
				mockPolicy := &jamf.Policy{
					Content: &jamf.PolicyContents{
						General: &jamf.PolicyGeneral{
							ID:   72,
							Name: "Test Policy",
						},
					},
				}
				policyData, err := json.MarshalIndent(mockPolicy, "", "    ")
				if err != nil {
					fmt.Fprintf(w, err.Error())
				}
				fmt.Fprintf(w, string(policyData))
			}
		default:
			http.Error(w, fmt.Sprintf("bad Jamf API %s call to %s", r.Method, r.URL), http.StatusInternalServerError)
			return
		}
		_, err := w.Write([]byte(resp))
		assert.Nil(t, err)
	}))
}

func TestGetAllPolicies(t *testing.T) {
	testServer := policiesResponseMocks(t)
	defer testServer.Close()
	j, err := jamf.NewClient(testServer.URL, "fake-username", "mock-password-cool", nil)
	assert.Nil(t, err)
	res, err := j.Policies()
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Len(t, res, 5)
	assert.Equal(t, 72, res[2].ID)
	assert.Equal(t, "Test Policy", res[2].Name)
}

func TestGetSpecificPolicyByID(t *testing.T) {
	testServer := policiesResponseMocks(t)
	defer testServer.Close()
	j, err := jamf.NewClient(testServer.URL, "fake-username", "mock-password-cool", nil)
	assert.Nil(t, err)
	policy, err := j.PolicyDetails(72)
	assert.Nil(t, err)
	assert.Equal(t, 72, policy.Content.General.ID)
	assert.Equal(t, "Test Policy", policy.Content.General.Name)
}

func TestGetSpecificPolicyByName(t *testing.T) {
	testServer := policiesResponseMocks(t)
	defer testServer.Close()
	j, err := jamf.NewClient(testServer.URL, "fake-username", "mock-password-cool", nil)
	assert.Nil(t, err)
	policy, err := j.PolicyDetails("Test Policy")
	assert.Nil(t, err)
	assert.Equal(t, 72, policy.Content.General.ID)
	assert.Equal(t, "Test Policy", policy.Content.General.Name)
}

func TestUpdatePolicy(t *testing.T) {
	testServer := policiesResponseMocks(t)
	defer testServer.Close()
	j, err := jamf.NewClient(testServer.URL, "fake-username", "mock-password-cool", nil)
	assert.Nil(t, err)

	updates := &jamf.PolicyContents{
		General: &jamf.PolicyGeneral{
			ID:   72,
			Name: "Test Policy",
		},
		Scope: &jamf.Scope{
			ComputerGroups: []*jamf.ComputerGroup{
				{
					Name: "Test Smart Group",
				},
			},
		},
		Scripts: []*jamf.PolicyScriptAssignment{
			{
				Name:       "Test Echo",
				Parameter4: "My Name",
			},
		},
	}

	policy, err := j.UpdatePolicy(72, updates)
	assert.Nil(t, err)
	assert.Equal(t, "Test Policy", policy.General.Name)
	assert.Equal(t, 72, policy.General.ID)
	assert.Equal(t, 1, len(policy.Scope.ComputerGroups))
	assert.Equal(t, "Test Smart Group", policy.Scope.ComputerGroups[0].Name)
	assert.Equal(t, "Test Echo", policy.Scripts[0].Name)
	assert.Equal(t, "My Name", policy.Scripts[0].Parameter4)
}

func TestCreatePolicy(t *testing.T) {
	testServer := policiesResponseMocks(t)
	defer testServer.Close()
	j, err := jamf.NewClient(testServer.URL, "fake-username", "mock-password-cool", nil)
	assert.Nil(t, err)

	newPolicy := &jamf.PolicyContents{
		General: &jamf.PolicyGeneral{
			Name:           "Test Policy",
			Trigger:        "EVENT",
			TriggerCheckIn: true,
			TriggerStartup: true,
			Frequency:      "Once per computer",
			Category: &jamf.PolicyCategory{
				Name: "Software - Security",
			},
		},
		Scope: &jamf.Scope{
			Computers: []*jamf.BasicComputerInfo{
				{
					GeneralInformation: jamf.GeneralInformation{
						Name: "TEST-BOX",
					},
				},
				{
					GeneralInformation: jamf.GeneralInformation{
						Name: "TestMachine",
					},
				},
			},
		},
		PackageConfiguration: &jamf.Packages{
			List: []*jamf.Package{
				{
					Name:   "test_macos_installer.pkg",
					Action: "Install",
				},
			},
		},
		Scripts: []*jamf.PolicyScriptAssignment{
			{
				Name:       "Test Echo",
				Parameter4: "Walter",
			},
		},
	}

	policy, err := j.CreatePolicy(newPolicy)
	assert.Nil(t, err)
	assert.NotNil(t, policy)
	assert.Equal(t, "Test Policy", policy.General.Name)
	assert.Equal(t, "Once per computer", policy.General.Frequency)
	assert.Equal(t, "Software - Security", policy.General.Category.Name)
	assert.Equal(t, 2, len(policy.Scope.Computers))
	assert.Equal(t, "TEST-BOX", policy.Scope.Computers[0].GeneralInformation.Name)
	assert.NotNil(t, policy.PackageConfiguration)
	assert.Equal(t, "test_macos_installer.pkg", policy.PackageConfiguration.List[0].Name)
	assert.Equal(t, 1, len(policy.Scripts))
	assert.Equal(t, "Test Echo", policy.Scripts[0].Name)
	assert.Equal(t, "Walter", policy.Scripts[0].Parameter4)
	assert.Equal(t, "After", policy.Scripts[0].Priority)
}

func DeletePolicy(t *testing.T) {
	testServer := policiesResponseMocks(t)
	defer testServer.Close()
	j, err := jamf.NewClient(testServer.URL, "fake-username", "mock-password-cool", nil)
	assert.Nil(t, err)
	removed, err := j.DeletePolicy(72)
	assert.Nil(t, err)
	assert.Equal(t, 72, removed.ID)
}

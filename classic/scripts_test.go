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

var SCRIPTS_API_BASE_ENDPOINT = "/JSSResource/scripts"

func scriptsResponseMocks(t *testing.T) *httptest.Server {
	var resp string
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.RequestURI {
		case SCRIPTS_API_BASE_ENDPOINT:
			fmt.Fprintf(w, `{
				"scripts": [
					{
							"id": 52,
							"name": "Admin to Standard"
					},
					{
							"id": 1,
							"name": "Cache macOS Updates"
					},
					{
							"id": 33,
							"name": "Zoom Script 2"
					},
					{
							"id": 102,
							"name": "Chrome Default Browser"
					},
					{
							"id": 175,
							"name": "Chrome latest script"
					},
					{
							"id": 86,
							"name": "Chrome Update to 63"
					}]
			}`)
		case fmt.Sprintf("%s/id/33", SCRIPTS_API_BASE_ENDPOINT), fmt.Sprintf("%s/id/-1", SCRIPTS_API_BASE_ENDPOINT), fmt.Sprintf("%s/name/Zoom Script 2", SCRIPTS_API_BASE_ENDPOINT):
			switch r.Method {
			case "PUT", "POST":
				data, err := ioutil.ReadAll(r.Body)
				if err != nil {
					fmt.Fprintf(w, err.Error())
				}
				scriptContents := &jamf.ScriptContents{}
				err = xml.Unmarshal(data, scriptContents)
				if err != nil {
					fmt.Fprintf(w, err.Error())
				}
				scriptData, err := json.MarshalIndent(scriptContents, "", "    ")
				if err != nil {
					fmt.Fprintf(w, err.Error())
				}
				fmt.Fprintf(w, string(scriptData))
			default:
				mockScript := &jamf.Script{
					Content: &jamf.ScriptContents{
						ID:              33,
						Name:            "Zoom Script 2",
						Category:        "No category assigned",
						Priority:        "After",
						Contents:        "#!/bin/bash\n#Get latest version from Jamf UI Parameters\nZoom_Target_Version=\"$4\"\necho $Zoom_Target_Version",
						EncodedContents: "IyEvYmluL2Jhc2gKI0dlQ==",
					},
				}
				var (
					scriptData []byte
					err        error
				)

				if r.Method == "DELETE" {
					scriptData, err = json.MarshalIndent(mockScript.Content, "", "    ")
					if err != nil {
						fmt.Fprintf(w, err.Error())
					}
				} else {
					scriptData, err = json.MarshalIndent(mockScript, "", "    ")
					if err != nil {
						fmt.Fprintf(w, err.Error())
					}
				}
				fmt.Fprintf(w, string(scriptData))
			}
		default:
			http.Error(w, fmt.Sprintf("bad Jamf API %s call to %s", r.Method, r.URL), http.StatusInternalServerError)
			return
		}
		_, err := w.Write([]byte(resp))
		assert.Nil(t, err)
	}))
}

func TestGetAllScripts(t *testing.T) {
	testServer := scriptsResponseMocks(t)
	defer testServer.Close()
	j, err := jamf.NewClient(testServer.URL, "fake-username", "mock-password-cool", nil)
	assert.Nil(t, err)
	scripts, err := j.Scripts()
	assert.Nil(t, err)
	assert.NotNil(t, scripts)
	assert.Len(t, scripts, 6)
	assert.Equal(t, 33, scripts[2].ID)
	assert.Equal(t, "Zoom Script 2", scripts[2].Name)
}

func TestGetSpecificScriptByID(t *testing.T) {
	testServer := scriptsResponseMocks(t)
	defer testServer.Close()
	j, err := jamf.NewClient(testServer.URL, "fake-username", "mock-password-cool", nil)
	assert.Nil(t, err)
	script, err := j.ScriptDetails(33)
	assert.Nil(t, err)
	assert.Equal(t, 33, script.Content.ID)
	assert.Equal(t, "Zoom Script 2", script.Content.Name)
}

func TestGetSpecificScriptByName(t *testing.T) {
	testServer := scriptsResponseMocks(t)
	defer testServer.Close()
	j, err := jamf.NewClient(testServer.URL, "fake-username", "mock-password-cool", nil)
	assert.Nil(t, err)
	script, err := j.ScriptDetails(33)
	assert.Nil(t, err)
	assert.Equal(t, 33, script.Content.ID)
	assert.Equal(t, "Zoom Script 2", script.Content.Name)
}

func TestUpdateScript(t *testing.T) {
	testServer := scriptsResponseMocks(t)
	defer testServer.Close()
	j, err := jamf.NewClient(testServer.URL, "fake-username", "mock-password-cool", nil)
	assert.Nil(t, err)

	update := &jamf.ScriptContents{
		Notes: "I am updated!",
	}

	script, err := j.UpdateScript(33, update)
	assert.Nil(t, err)
	assert.Equal(t, "I am updated!", script.Notes)
}

func TestCreateScript(t *testing.T) {
	testServer := scriptsResponseMocks(t)
	defer testServer.Close()
	j, err := jamf.NewClient(testServer.URL, "fake-username", "mock-password-cool", nil)
	assert.Nil(t, err)

	newScript := &jamf.ScriptContents{
		Name:     "TestScript",
		Contents: "echo 'this is a test script'",
	}

	script, err := j.CreateScript(newScript)
	assert.Nil(t, err)
	assert.Equal(t, "TestScript", script.Name)
	assert.Equal(t, "TestScript", script.Filename)
	assert.Equal(t, "echo 'this is a test script'", script.Contents)
}

func TestCreateScriptRequiredContent(t *testing.T) {
	testServer := scriptsResponseMocks(t)
	defer testServer.Close()
	j, err := jamf.NewClient(testServer.URL, "fake-username", "mock-password-cool", nil)
	assert.Nil(t, err)

	newScript := &jamf.ScriptContents{}

	_, err = j.CreateScript(newScript)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "name required for new script")

	newScriptNoContent := &jamf.ScriptContents{
		Name: "I am missing contents",
	}
	_, contentErr := j.CreateScript(newScriptNoContent)
	assert.NotNil(t, contentErr)
	assert.Contains(t, contentErr.Error(), "script contents required")
}

func TestDeleteScript(t *testing.T) {
	testServer := scriptsResponseMocks(t)
	defer testServer.Close()
	j, err := jamf.NewClient(testServer.URL, "fake-username", "mock-password-cool", nil)
	assert.Nil(t, err)
	removed, err := j.DeleteScript(33)
	assert.Nil(t, err)
	assert.Equal(t, 33, removed.ID)
}

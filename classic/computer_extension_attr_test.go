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

var COMPUTER_EXT_ATTR_API_BASE_ENDPOINT = "/JSSResource/computerextensionattributes"

func computerExtAttrResponseMocks(t *testing.T) *httptest.Server {
	var resp string
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.RequestURI {
		case COMPUTER_EXT_ATTR_API_BASE_ENDPOINT:
			fmt.Fprintf(w, `{
				"computer_extension_attributes": [
					{
							"id": 33,
							"name": "Check Firewall"
					},
					{
							"id": 99,
							"name": "Is Logged In User Admin"
					},
					{
							"id": 103,
							"name": "Team"
					}]
			}`)
		case fmt.Sprintf("%s/id/33", COMPUTER_EXT_ATTR_API_BASE_ENDPOINT), fmt.Sprintf("%s/id/-1", COMPUTER_EXT_ATTR_API_BASE_ENDPOINT), fmt.Sprintf("%s/name/Check%sFirewall", COMPUTER_EXT_ATTR_API_BASE_ENDPOINT, "%20"):
			switch r.Method {
			case "PUT", "POST":
				data, err := ioutil.ReadAll(r.Body)
				if err != nil {
					fmt.Fprintf(w, err.Error())
				}
				compExtAttrContents := &jamf.ComputerExtensionAttribute{}
				err = xml.Unmarshal(data, compExtAttrContents)
				if err != nil {
					fmt.Fprintf(w, err.Error())
				}
				compExtData, err := json.MarshalIndent(compExtAttrContents, "", "    ")
				if err != nil {
					fmt.Fprintf(w, err.Error())
				}
				fmt.Fprintf(w, string(compExtData))
			default:
				mockCompExtAttr := &jamf.ComputerExtensionAttributeDetails{
					Details: &jamf.ComputerExtensionAttribute{
						ID:          33,
						Name:        "Check Firewall",
						Enabled:     true,
						Description: "Checks to ensure firewall is enabled on client",
						DataType:    "String",
						InputType: &jamf.ComputerExtensionAttrInputType{
							Type: "",
						},
						InventoryDisplay: "Operating System",
						ReconDisplay:     "Extension Attributes",
					},
				}

				var (
					compExtData []byte
					err         error
				)

				if r.Method == "DELETE" {
					compExtData, err = json.MarshalIndent(mockCompExtAttr.Details, "", "    ")
					if err != nil {
						fmt.Fprintf(w, err.Error())
					}
				} else {
					compExtData, err = json.MarshalIndent(mockCompExtAttr, "", "    ")
					if err != nil {
						fmt.Fprintf(w, err.Error())
					}
				}

				fmt.Fprintf(w, string(compExtData))
			}
		default:
			http.Error(w, fmt.Sprintf("bad Jamf API %s call to %s", r.Method, r.URL), http.StatusInternalServerError)
			return
		}
		_, err := w.Write([]byte(resp))
		assert.Nil(t, err)
	}))
}

func TestQueryAllComputerExtAttrs(t *testing.T) {
	testServer := computerExtAttrResponseMocks(t)
	defer testServer.Close()
	j, err := jamf.NewClient(testServer.URL, "fake-username", "mock-password-cool", nil)
	assert.Nil(t, err)
	compExtAttrs, err := j.ComputerExtensionAttributes()
	assert.Nil(t, err)
	assert.NotNil(t, compExtAttrs)
	assert.Equal(t, 3, len(compExtAttrs))
	assert.Equal(t, 99, compExtAttrs[1].ID)
	assert.Equal(t, "Is Logged In User Admin", compExtAttrs[1].Name)
}

func TestQuerySpecificComputerExtAttrByName(t *testing.T) {
	testServer := computerExtAttrResponseMocks(t)
	defer testServer.Close()
	j, err := jamf.NewClient(testServer.URL, "fake-username", "mock-password-cool", nil)
	assert.Nil(t, err)
	cea, err := j.ComputerExtensionAttributeDetails("Check Firewall")
	assert.Nil(t, err)
	assert.NotNil(t, cea)
	assert.Equal(t, 33, cea.Details.ID)
	assert.Equal(t, "Check Firewall", cea.Details.Name)
	assert.True(t, cea.Details.Enabled)
	assert.Equal(t, "Checks to ensure firewall is enabled on client", cea.Details.Description)
	assert.Equal(t, "String", cea.Details.DataType)
	assert.Empty(t, cea.Details.InputType.Type)
	assert.Equal(t, "Operating System", cea.Details.InventoryDisplay)
	assert.Equal(t, "Extension Attributes", cea.Details.ReconDisplay)
}

func TestQuerySpecificComputerExtAttrByID(t *testing.T) {
	testServer := computerExtAttrResponseMocks(t)
	defer testServer.Close()
	j, err := jamf.NewClient(testServer.URL, "fake-username", "mock-password-cool", nil)
	assert.Nil(t, err)
	cea, err := j.ComputerExtensionAttributeDetails(33)
	assert.Nil(t, err)
	assert.NotNil(t, cea)
	assert.Equal(t, 33, cea.Details.ID)
	assert.Equal(t, "Check Firewall", cea.Details.Name)
	assert.True(t, cea.Details.Enabled)
	assert.Equal(t, "Checks to ensure firewall is enabled on client", cea.Details.Description)
	assert.Equal(t, "String", cea.Details.DataType)
	assert.Empty(t, cea.Details.InputType.Type)
	assert.Equal(t, "Operating System", cea.Details.InventoryDisplay)
	assert.Equal(t, "Extension Attributes", cea.Details.ReconDisplay)
}

func TestUpdateComputerExtAttr(t *testing.T) {
	testServer := computerExtAttrResponseMocks(t)
	defer testServer.Close()
	j, err := jamf.NewClient(testServer.URL, "fake-username", "mock-password-cool", nil)
	assert.Nil(t, err)

	update := &jamf.ComputerExtensionAttribute{
		Description: "Updated description",
		Enabled:     false,
	}

	updatedComputerExtAttr, err := j.UpdateComputerExtensionAttribue(33, update)
	assert.Nil(t, err)
	assert.Equal(t, "Updated description", updatedComputerExtAttr.Description)
	assert.False(t, updatedComputerExtAttr.Enabled)
}

func TestCreateComputerExtAttr(t *testing.T) {
	testServer := computerExtAttrResponseMocks(t)
	defer testServer.Close()
	j, err := jamf.NewClient(testServer.URL, "fake-username", "mock-password-cool", nil)
	assert.Nil(t, err)

	newCompExtAttr := &jamf.ComputerExtensionAttribute{}
	_, err = j.CreateComputerExtensionAttribute(newCompExtAttr)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Name required for new computer extension attribute")

	newCompExtAttr = &jamf.ComputerExtensionAttribute{
		Name:        "Testing Ext Attr",
		Description: "This is a test description",
		Enabled:     true,
		DataType:    "String",
		InputType: &jamf.ComputerExtensionAttrInputType{
			Type:     "script",
			Platform: "Mac",
			Script:   "echo \"Hello World, I am a unit test\"",
		},
	}
	cea, err := j.CreateComputerExtensionAttribute(newCompExtAttr)
	assert.Nil(t, err)
	assert.Equal(t, "Testing Ext Attr", cea.Name)
	assert.Equal(t, "This is a test description", cea.Description)
	assert.True(t, cea.Enabled)
	assert.Equal(t, "Mac", cea.InputType.Platform)
	assert.Equal(t, "script", cea.InputType.Type)
	assert.Equal(t, "echo \"Hello World, I am a unit test\"", cea.InputType.Script)
}

func TestDeleteComputerExtAttr(t *testing.T) {
	testServer := computerExtAttrResponseMocks(t)
	defer testServer.Close()
	j, err := jamf.NewClient(testServer.URL, "fake-username", "mock-password-cool", nil)
	assert.Nil(t, err)
	removed, err := j.DeleteComputerExtensionAttribute(33)
	assert.Nil(t, err)
	assert.Equal(t, 33, removed.ID)
}

func TestValidateComputerExtAttrDataTypePass(t *testing.T) {
	ce := &jamf.ComputerExtensionAttribute{}
	for _, dt := range []string{"", "String", "Integer", "Date"} {
		ce.DataType = dt
		err := ce.ValidateDataType()
		assert.Nil(t, err)
	}
}

func TestValidateComputerExtAttrDataTypeFail(t *testing.T) {
	ce := &jamf.ComputerExtensionAttribute{}
	for _, dt := range []string{"IDK", "badData", "script", "policy"} {
		ce.DataType = dt
		err := ce.ValidateDataType()
		assert.NotNil(t, err)
		assert.Equal(t, fmt.Sprintf("%s is not a valid computer extension attribute data type must be of type [ String, Integer, Date ]", dt), err.Error())
	}
}

func TestValidateComputerExtAttrInventoryDisplayPass(t *testing.T) {
	ce := &jamf.ComputerExtensionAttribute{}
	for _, dt := range []string{"", "general", "hardware", "operating system", "user and location", "purchasing", "extension attributes"} {
		ce.InventoryDisplay = dt
		err := ce.ValidateInventoryDisplay()
		assert.Nil(t, err)
	}
}

func TestValidateComputerExtAttrInventoryDisplayFail(t *testing.T) {
	ce := &jamf.ComputerExtensionAttribute{}
	for _, dt := range []string{"IDK", "badData", "script", "policy"} {
		ce.InventoryDisplay = dt
		err := ce.ValidateInventoryDisplay()
		assert.NotNil(t, err)
		assert.Equal(t, fmt.Sprintf("%s is not a valid computer extension inventory display type must be of type [ General, Hardware, Operating System, User and Location, Purchasing, Extension Attributes ]", dt), err.Error())
	}
}

func TestValidateComputerExtAttrReconDisplayPass(t *testing.T) {
	ce := &jamf.ComputerExtensionAttribute{}
	for _, dt := range []string{"", "computer", "user and location", "purchasing", "extension attributes"} {
		ce.ReconDisplay = dt
		err := ce.ValidateReconDisplay()
		assert.Nil(t, err)
	}
}

func TestValidateComputerExtAttrReconDisplayFail(t *testing.T) {
	ce := &jamf.ComputerExtensionAttribute{}
	for _, dt := range []string{"IDK", "badData", "script", "policy"} {
		ce.ReconDisplay = dt
		err := ce.ValidateReconDisplay()
		assert.NotNil(t, err)
		assert.Equal(t, fmt.Sprintf("%s is not a valid computer extension inventory display type must be of type [ Computer, User and Location, Purchasing, Extension Attributes ]", dt), err.Error())
	}
}

func TestValidateComputerExtAttrInputTypePass(t *testing.T) {
	ce := &jamf.ComputerExtensionAttrInputType{}
	for _, dt := range []string{"", "Text Field", "LDAP Mapping", "Pop-up Menu"} {
		ce.Type = dt
		err := ce.ValidateInputType()
		assert.Nil(t, err)
	}
}

func TestValidateComputerExtAttrInputTypeScript(t *testing.T) {
	ce := &jamf.ComputerExtensionAttrInputType{
		Type: "script",
	}

	// test failure missing script contents which are required
	err := ce.ValidateInputType()
	assert.NotNil(t, err)
	assert.Equal(t, "script contents must be provided for input type script", err.Error())

	// test passing case with script contents
	ce.Script = "echo \"Hello world\""
	err = ce.ValidateInputType()
	assert.Nil(t, err)
}

func TestValidateComputerExtAttrInputTypeFail(t *testing.T) {
	ce := &jamf.ComputerExtensionAttrInputType{}
	for _, dt := range []string{"IDK", "badData"} {
		ce.Type = dt
		err := ce.ValidateInputType()
		assert.NotNil(t, err)
		assert.Equal(t, fmt.Sprintf("%s is not a valid computer extension attribute input type must be of type [ script, Text Field, LDAP Mapping, Pop-up Menu ]", dt), err.Error())
	}
}

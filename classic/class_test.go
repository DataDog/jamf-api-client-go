// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0

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

var CLASS_API_BASE_ENDPOINT = "/JSSResource/computerextensionattributes"

func classResponseMocks(t *testing.T) *httptest.Server {
	var resp string
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.RequestURI {
		case CLASS_API_BASE_ENDPOINT:
			fmt.Fprintf(w, `{
				"classes": [
					{
							"id": 6243,
							"name": "1st - Math"
					},
					{
							"id": 6244,
							"name": "3rd - Science"
					},
					{
							"id": 6244,
							"name": "5th - English"
					}]
			}`)
		case fmt.Sprintf("%s/id/6243", CLASS_API_BASE_ENDPOINT), fmt.Sprintf("%s/id/-1", CLASS_API_BASE_ENDPOINT), fmt.Sprintf("%s/name/5th%s-%sEnglish", CLASS_API_BASE_ENDPOINT, "%20", "%20"):
			switch r.Method {
			case "PUT", "POST":
				data, err := ioutil.ReadAll(r.Body)
				if err != nil {
					fmt.Fprintf(w, err.Error())
				}
				classContents := &jamf.Class{}
				err = xml.Unmarshal(data, classContents)
				if err != nil {
					fmt.Fprintf(w, err.Error())
				}
				classData, err := json.MarshalIndent(classContents, "", "    ")
				if err != nil {
					fmt.Fprintf(w, err.Error())
				}
				fmt.Fprintf(w, string(classData))
			default:
				mockClass := &jamf.ClassDetails{
					Details: &jamf.Class{
						ID:          6243,
						Name:        "1st - Math",
						Description: "First period math class.",
						Students: []string{
							"jappleseed@example.com",
							"sappleseed@example.com",
						},
						Teachers: []string{
							"jdoe@example.com",
						},
						MeetingTimes: []jamf.MeetingTime{
							jamf.MeetingTime{
								Days:      "M W F",
								StartTime: "845",
								EndTime:   "945",
							},
						},
					},
				}

				var (
					classData []byte
					err       error
				)

				if r.Method == "DELETE" {
					classData, err = json.MarshalIndent(mockClass.Details, "", "    ")
					if err != nil {
						fmt.Fprintf(w, err.Error())
					}
				} else {
					classData, err = json.MarshalIndent(mockClass, "", "    ")
					if err != nil {
						fmt.Fprintf(w, err.Error())
					}
				}

				fmt.Fprintf(w, string(classData))
			}
		default:
			http.Error(w, fmt.Sprintf("bad Jamf API %s call to %s", r.Method, r.URL), http.StatusInternalServerError)
			return
		}
		_, err := w.Write([]byte(resp))
		assert.Nil(t, err)
	}))
}

func TestQueryAllClasses(t *testing.T) {
	testServer := classResponseMocks(t)
	defer testServer.Close()
	j, err := jamf.NewClient(testServer.URL, "fake-username", "mock-password-cool", nil)
	assert.Nil(t, err)
	classes, err := j.Classes()
	assert.Nil(t, err)
	assert.NotNil(t, classes)
	assert.Equal(t, 3, len(classes))
	assert.Equal(t, 6244, classes[1].ID)
	assert.Equal(t, "3rd - Science", classes[1].Name)
}

func TestQuerySpecificClassByName(t *testing.T) {
	testServer := classResponseMocks(t)
	defer testServer.Close()
	j, err := jamf.NewClient(testServer.URL, "fake-username", "mock-password-cool", nil)
	assert.Nil(t, err)
	class, err := j.ClassDetails("5th - English")
	assert.Nil(t, err)
	assert.NotNil(t, class)
	assert.Equal(t, 6244, class.Details.ID)
	assert.Equal(t, "5th - English", class.Details.Name)
}

func TestQuerySpecificClassByID(t *testing.T) {
	testServer := classResponseMocks(t)
	defer testServer.Close()
	j, err := jamf.NewClient(testServer.URL, "fake-username", "mock-password-cool", nil)
	assert.Nil(t, err)
	class, err := j.ClassDetails(6243)
	assert.Nil(t, err)
	assert.NotNil(t, class)
	assert.Equal(t, 6243, class.Details.ID)
	assert.Equal(t, "1st - Math", class.Details.Name)
}

// func TestUpdateClass(t *testing.T) {
// 	testServer := computerExtAttrResponseMocks(t)
// 	defer testServer.Close()
// 	j, err := jamf.NewClient(testServer.URL, "fake-username", "mock-password-cool", nil)
// 	assert.Nil(t, err)

// 	update := &jamf.ComputerExtensionAttribute{
// 		Description: "Updated description",
// 		Enabled:     false,
// 	}

// 	updatedComputerExtAttr, err := j.UpdateComputerExtensionAttribue(33, update)
// 	assert.Nil(t, err)
// 	assert.Equal(t, "Updated description", updatedComputerExtAttr.Description)
// 	assert.False(t, updatedComputerExtAttr.Enabled)
// }

// func TestCreateClass(t *testing.T) {
// 	testServer := computerExtAttrResponseMocks(t)
// 	defer testServer.Close()
// 	j, err := jamf.NewClient(testServer.URL, "fake-username", "mock-password-cool", nil)
// 	assert.Nil(t, err)

// 	newCompExtAttr := &jamf.ComputerExtensionAttribute{}
// 	_, err = j.CreateComputerExtensionAttribute(newCompExtAttr)
// 	assert.NotNil(t, err)
// 	assert.Contains(t, err.Error(), "Name required for new computer extension attribute")

// 	newCompExtAttr = &jamf.ComputerExtensionAttribute{
// 		Name:        "Testing Ext Attr",
// 		Description: "This is a test description",
// 		Enabled:     true,
// 		DataType:    "String",
// 		InputType: &jamf.ComputerExtensionAttrInputType{
// 			Type:     "script",
// 			Platform: "Mac",
// 			Script:   "echo \"Hello World, I am a unit test\"",
// 		},
// 	}
// 	cea, err := j.CreateComputerExtensionAttribute(newCompExtAttr)
// 	assert.Nil(t, err)
// 	assert.Equal(t, "Testing Ext Attr", cea.Name)
// 	assert.Equal(t, "This is a test description", cea.Description)
// 	assert.True(t, cea.Enabled)
// 	assert.Equal(t, "Mac", cea.InputType.Platform)
// 	assert.Equal(t, "script", cea.InputType.Type)
// 	assert.Equal(t, "echo \"Hello World, I am a unit test\"", cea.InputType.Script)
// }

func TestDeleteClass(t *testing.T) {
	testServer := classResponseMocks(t)
	defer testServer.Close()
	j, err := jamf.NewClient(testServer.URL, "fake-username", "mock-password-cool", nil)
	assert.Nil(t, err)
	removed, err := j.DeleteClass(6243)
	assert.Nil(t, err)
	assert.Equal(t, 6243, removed.ID)
}

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

func computerExtResponseMocks(t *testing.T) *httptest.Server {
	var resp string
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.RequestURI {
		case COMPUTER_EXT_ATTR_API_BASE_ENDPOINT:
			fmt.Fprintf(w, `{
				"scripts": [
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
		case fmt.Sprintf("%s/id/33", COMPUTER_EXT_ATTR_API_BASE_ENDPOINT), fmt.Sprintf("%s/id/-1", COMPUTER_EXT_ATTR_API_BASE_ENDPOINT), fmt.Sprintf("%s/name/Check Firewall", COMPUTER_EXT_ATTR_API_BASE_ENDPOINT):
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
				mockCompExtAttr := &jamf.ComputerExtensionAttribute{
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
				}

				compExtData, err := json.MarshalIndent(mockCompExtAttr, "", "    ")
				if err != nil {
					fmt.Fprintf(w, err.Error())
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

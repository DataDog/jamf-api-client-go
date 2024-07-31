package classic_test

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	jamf "github.com/DataDog/jamf-api-client-go/classic"
	"github.com/stretchr/testify/assert"
)

var COMPUTER_GROUPS_BASE_API_ENDPOINT = "/JSSResource/computergroups"

func computerGroupsResponseMocks(t *testing.T) *httptest.Server {
	var resp string
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.RequestURI {
		case COMPUTER_GROUPS_BASE_API_ENDPOINT:
			fmt.Fprintf(w, `{
				"computer_groups": [
					{
							"id": 1,
							"name": "Test Group 1",
							"is_smart": false
					},
					{
							"id": 2,
							"name": "Test Group 2",
							"is_smart": true
					},
					{
							"id": 3,
							"name": "Test Group 3",
							"is_smart": false
					}]
				}`)
		case fmt.Sprintf("%s/id/1", COMPUTER_GROUPS_BASE_API_ENDPOINT), fmt.Sprintf("%s/id/-1", COMPUTER_GROUPS_BASE_API_ENDPOINT), fmt.Sprintf("%s/name/Test%sGroup%s1", COMPUTER_GROUPS_BASE_API_ENDPOINT, "%20", "%20"):
			switch r.Method {
			case "GET":
				w.Header().Add("Content-Type", "application/xml")
				mockGroup := &jamf.ComputerGroup{
					Info: jamf.ComputerGroupDetails{
						BasicComputerGroupInfo: jamf.BasicComputerGroupInfo{
							ID:      1,
							Name:    "Test Group 1",
							IsSmart: false,
						},
						Computers: []jamf.BasicComputerInfo{
							{
								GeneralInformation: jamf.GeneralInformation{
									ID:   1,
									Name: "Test Computer 1",
								},
							},
						},
					},
				}
				groupData, err := xml.MarshalIndent(mockGroup, "", "  ")
				if err != nil {
					fmt.Fprintf(w, err.Error())
				}
				fmt.Fprintf(w, string(groupData))
			case "POST":
				w.Header().Add("Content-Type", "application/xml")
				data, err := io.ReadAll(r.Body)
				if err != nil {
					fmt.Fprintf(w, err.Error())
				}
				groupContents := &jamf.ComputerGroupDetails{}
				err = xml.Unmarshal(data, groupContents)
				if err != nil {
					fmt.Fprintf(w, err.Error())
				}
				groupData, err := xml.MarshalIndent(groupContents, "", "  ")
				if err != nil {
					fmt.Fprintf(w, err.Error())
				}
				fmt.Fprintf(w, string(groupData))
			default:
				w.Header().Add("Content-Type", "application/xml")
				mockGroup := &jamf.ComputerGroupDetails{
					BasicComputerGroupInfo: jamf.BasicComputerGroupInfo{
						ID:      1,
						Name:    "Test Group 1",
						IsSmart: false,
					},
					Computers: []jamf.BasicComputerInfo{
						{
							GeneralInformation: jamf.GeneralInformation{
								ID:   1,
								Name: "Test Computer 1",
							},
						},
					},
				}
				groupData, err := xml.MarshalIndent(mockGroup, "", "  ")
				if err != nil {
					fmt.Fprintf(w, err.Error())
				}
				fmt.Fprintf(w, string(groupData))
			}
		default:
			http.Error(w, fmt.Sprintf("bad Jamf API %s call to %s", r.Method, r.URL), http.StatusInternalServerError)
			return
		}
		_, err := w.Write([]byte(resp))
		assert.Nil(t, err)
	}))
}

func TestListAllComputerGroups(t *testing.T) {
	server := computerGroupsResponseMocks(t)
	defer server.Close()
	j, err := jamf.NewClient(server.URL, "test", "test", server.Client(), jamf.WithTokenAuth())
	assert.Nil(t, err)
	grps, err := j.ComputerGroups()
	assert.Nil(t, err)
	assert.Equal(t, 3, len(grps))
	assert.Equal(t, 1, grps[0].ID)
	assert.Equal(t, "Test Group 1", grps[0].Name)
	assert.Equal(t, false, grps[0].IsSmart)
	assert.Equal(t, 2, grps[1].ID)
	assert.Equal(t, "Test Group 2", grps[1].Name)
	assert.Equal(t, true, grps[1].IsSmart)
	assert.Equal(t, 3, grps[2].ID)
	assert.Equal(t, "Test Group 3", grps[2].Name)
	assert.Equal(t, false, grps[2].IsSmart)
}

func TestQuerySpecificComputerGroups(t *testing.T) {
	server := computerGroupsResponseMocks(t)
	defer server.Close()
	j, err := jamf.NewClient(server.URL, "test", "test", server.Client(), jamf.WithTokenAuth())
	assert.Nil(t, err)
	grp, err := j.ComputerGroupDetails(1)
	assert.Nil(t, err)
	assert.Equal(t, 1, grp.Info.ID)
	assert.Equal(t, "Test Group 1", grp.Info.Name)
	assert.Equal(t, false, grp.Info.IsSmart)
	assert.Equal(t, 1, grp.Info.Computers[0].ID)
	assert.Equal(t, "Test Computer 1", grp.Info.Computers[0].Name)
}

func TestCreateComputerGroup(t *testing.T) {
	server := computerGroupsResponseMocks(t)
	defer server.Close()
	j, err := jamf.NewClient(server.URL, "test", "test", server.Client(), jamf.WithTokenAuth())
	assert.Nil(t, err)
	grp := &jamf.ComputerGroupDetails{
		BasicComputerGroupInfo: jamf.BasicComputerGroupInfo{
			Name:    "Unit Test Group",
			IsSmart: false,
		},
		Computers: []jamf.BasicComputerInfo{},
	}
	createdGrp, err := j.CreateComputerGroup(grp)
	assert.Nil(t, err)
	assert.Equal(t, "Unit Test Group", createdGrp.Name)
	assert.Equal(t, false, createdGrp.IsSmart)
	assert.Equal(t, 0, len(createdGrp.Computers))
}

func TestDeleteComputerGroup(t *testing.T) {
	server := computerGroupsResponseMocks(t)
	defer server.Close()
	j, err := jamf.NewClient(server.URL, "test", "test", server.Client(), jamf.WithTokenAuth())
	assert.Nil(t, err)
	deletedGrp, err := j.DeleteComputerGroup(1)
	assert.Nil(t, err)
	assert.Equal(t, 1, deletedGrp.ID)
	assert.Equal(t, "Test Group 1", deletedGrp.Name)
	assert.Equal(t, false, deletedGrp.IsSmart)

	deletedGrp, err = j.DeleteComputerGroup("Test%20Group%201")
	assert.Nil(t, err)
	assert.Equal(t, 1, deletedGrp.ID)
	assert.Equal(t, "Test Group 1", deletedGrp.Name)
	assert.Equal(t, false, deletedGrp.IsSmart)
}

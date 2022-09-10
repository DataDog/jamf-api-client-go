package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"time"

	jamf "github.com/DataDog/jamf-api-client-go/classic"
)

func checkAndHandleErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic(1)
	}
}

func main() {
	username := os.Getenv("JAMF_USERNAME")
	password := os.Getenv("JAMF_PASSWORD")
	domain := os.Getenv("JAMF_DOMAIN")

	j, err := jamf.NewClient(domain, username, password, nil)
	checkAndHandleErr(err)

	// list computer extenstion attribues
	classes, err := j.Classes()
	checkAndHandleErr(err)

	for _, class := range classes {
		fmt.Printf("Class: %s\n", class.Name)
	}
	fmt.Printf("\n")

	// get details about a specific extension attribute
	classDetails, err := j.ClassDetails(6234)
	checkAndHandleErr(err)

	detailsBytes, err := json.Marshal(classDetails.Details)
	checkAndHandleErr(err)

	prettyExtAttrDetails := jamf.JSONPrettyPrint(detailsBytes)
	fmt.Printf("%s\n\n", prettyExtAttrDetails)

	// create new computer extnesion attribute
	newClass := &jamf.Class{
		Name:        "Go-created Class",
		Description: "This is testing the Jamf Go API Client",
		Students: []string{
			"scox@nrcaknights.com",
			"zdean@nrcaknights.com",
		},
		Teachers: []string{
			"lciancan@nrcaknights.com",
		},
		MeetingTimes: []jamf.MeetingTime{
			jamf.MeetingTime{
				Days:      "M W F",
				StartTime: "420",
				EndTime:   "905",
			},
		},
	}

	dets, err := xml.Marshal(newClass)
	checkAndHandleErr(err)
	fmt.Printf("%s", dets)

	created, err := j.CreateClass(newClass)
	checkAndHandleErr(err)
	fmt.Printf("Created %s - ID: %d\n\n", created.Name, created.ID)

	// get details about class we just created
	time.Sleep(30 * time.Second)
	newClassDetails, err := j.ClassDetails(created.ID)
	checkAndHandleErr(err)

	newDetailsBytes, err := json.Marshal(newClassDetails.Details)
	checkAndHandleErr(err)

	prettyNewClassDetails := jamf.JSONPrettyPrint(newDetailsBytes)
	fmt.Printf("%s\n\n", prettyNewClassDetails)

	// Note if you run this example the server can't find the new ID right away
	// so we sleep for 30 seconds
	time.Sleep(180 * time.Second)
	deleted, err := j.DeleteClass(created.ID) // Can delete using ID or Name
	checkAndHandleErr(err)
	fmt.Printf("Deleted ID: %d\n", deleted.ID)
}

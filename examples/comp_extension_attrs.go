package main

import (
	"encoding/json"
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
	extAttrs, err := j.ComputerExtensionAttributes()
	checkAndHandleErr(err)

	for _, attr := range extAttrs[:len(extAttrs)/3] {
		fmt.Printf("Extension Attribute: %s\n", attr.Name)
	}
	fmt.Printf("\n")

	// get details about a specific extension attribute
	extrAttrDetails, err := j.ComputerExtensionAttributeDetails(12)
	checkAndHandleErr(err)

	detailsBytes, err := json.Marshal(extrAttrDetails.Details)
	checkAndHandleErr(err)

	prettyExtAttrDetails := jamf.JSONPrettyPrint(detailsBytes)
	fmt.Printf("%s\n\n", prettyExtAttrDetails)
	/*
			{
		    "computer_extension_attribute": {
		        "id": 21,
		        "name": "Test Extension Attribute for API",
		        "enabled": true,
		        "description": "This is testing the Jamf Go API Client",
		        "data_type": "String",
		        "input_type": {
		            "type": "script",
		            "platform": "Mac",
		            "script": "#!/bin/bash\necho 'hello world'"
		        },
		        "inventory_display": "Extension Attributes",
		        "recon_display": "Extension Attributes"
		    }
		}
	*/

	// update existing computer extension attribute
	updatedDetails := &jamf.ComputerExtensionAttribute{
		Description: "This is an example of a description update",
	}
	updatedCompExtAttr, err := j.UpdateComputerExtensionAttribue(12, updatedDetails)
	checkAndHandleErr(err)
	fmt.Printf("Updated description for ID: %d\n\n", updatedCompExtAttr.ID)
	/*
		<computer_extension_attribute>
		    <id>{{ ID From Above}}</id>
		</computer_extension_attribute>
	*/

	// create new computer extnesion attribute
	newExtAttr := &jamf.ComputerExtensionAttribute{
		Name:        "Test Extension Attribute for API",
		Enabled:     false,
		Description: "This is testing the Jamf Go API Client",
		DataType:    "String",
		InputType: &jamf.ComputerExtensionAttrInputType{
			Type:     "script",
			Platform: "Mac",
			Script:   "#!/bin/bash\r\necho 'hello world'",
		},
		InventoryDisplay: "Extension Attributes",
		ReconDisplay:     "Extension Attributes",
	}

	created, err := j.CreateComputerExtensionAttribute(newExtAttr)
	checkAndHandleErr(err)
	fmt.Printf("Created %s - ID: %d\n\n", created.Name, created.ID)
	/*
			{
		    "computer_extension_attribute": {
		        "id": 21,
		        "name": "Test Extension Attribute for API",
		        "enabled": false,
		        "description": "This is testing the Jamf Go API Client",
		        "data_type": "String",
		        "input_type": {
		            "type": "script",
		            "platform": "Mac",
		            "script": "#!/bin/bash\necho 'hello world'"
		        },
		        "inventory_display": "Extension Attributes",
		        "recon_display": "Extension Attributes"
		    }
		}
	*/

	// Note if you run this example the server can't find the new ID right away
	// so we sleep for 30 seconds
	time.Sleep(30 * time.Second)
	deleted, err := j.DeleteComputerExtensionAttribute(created.ID) // Can delete using ID or Name
	checkAndHandleErr(err)
	fmt.Printf("Deleted ID: %d\n", deleted.ID)

	/*
		<computer_extension_attribute>
		    <id>{{ ID From Above}}</id>
		</computer_extension_attribute>
	*/
}

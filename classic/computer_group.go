package classic

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

func (j *Client) Groups() ([]BasicComputerGroupInformation, error) {
	ep := fmt.Sprintf("%s/%s", j.Endpoint, computerGroupsContext)
	req, err := http.NewRequestWithContext(context.Background(), "GET", ep, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error building Jamf computer groups query request")
	}
	res := Groups{}
	if err := j.makeAPIrequest(req, &res); err != nil {
		return nil, errors.Wrapf(err, "unable to query available computer groups from %s", ep)
	}
	return res.List, nil
}

// GroupDetails returns the details for a specific group given its ID or Name
func (j *Client) GroupDetails(identifier interface{}) (*ComputerGroupDetails, error) {
	ep, err := EndpointBuilder(j.Endpoint, computerGroupsContext, identifier)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF query request endpoint for group: %v", identifier)
	}

	req, err := http.NewRequestWithContext(context.Background(), "GET", ep, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF query request for group: %v", identifier)
	}

	res := ComputerGroupDetails{}
	if err := j.makeAPIrequest(req, &res); err != nil {
		return nil, errors.Wrapf(err, "unable to query group with ID: %d from %s", identifier, ep)
	}
	return &res, nil
}

// UpdateComputerGroupMemebers will update the members of a computer group in Jamf by either group ID or group Name
func (j *Client) UpdateComputerGroupMemebers(identifier interface{}, add, remove *[]GeneralInformation) (*ComputerGroupDetails, error) {
	ep, err := EndpointBuilder(j.Endpoint, computerGroupsContext, identifier)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF query request for group: %v", identifier)
	}

	type computer_group struct {
		Additions *[]GeneralInformation `xml:"computer_additions>computer"`
		Removals  *[]GeneralInformation `xml:"computer_deletions>computer"`
	}
	grp := computer_group{
		Additions: add,
		Removals:  remove,
	}
	bodyContent, err := xml.MarshalIndent(grp, "", "    ")
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF update payload for group: %v", identifier)
	}

	body := bytes.NewReader(bodyContent)
	req, err := http.NewRequestWithContext(context.Background(), "PUT", ep, body)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF update request for group: %v (%s)", identifier, ep)
	}
	fmt.Println(string(bodyContent))
	res := ComputerGroupDetails{}
	if err := j.makeAPIrequest(req, &res); err != nil {
		return nil, errors.Wrapf(err, "unable to process JAMF update request for group: %v (%s)", identifier, ep)
	}

	return &res, nil
}

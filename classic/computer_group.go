package classic

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// ComputerGroups represents a list of computer groups in Jamf
func (j *Client) ComputerGroups() ([]BasicComputerGroupInfo, error) {
	ep := fmt.Sprintf("%s/%s", j.Endpoint, computerGroupsContext)
	req, err := http.NewRequestWithContext(context.Background(), "GET", ep, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error building Jamf computer groups query request")
	}
	res := ComputerGroups{}
	if err := j.makeAPIrequest(req, &res); err != nil {
		return nil, errors.Wrapf(err, "unable to query available computer groups from %s", ep)
	}
	return res.List, nil
}

// ComputerGroupDetails returns the details for a specific group given its ID or Name
func (j *Client) ComputerGroupDetails(identifier any) (*ComputerGroup, error) {
	ep, err := EndpointBuilder(j.Endpoint, computerGroupsContext, identifier)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF query request endpoint for computer group: %v", identifier)
	}

	req, err := http.NewRequestWithContext(context.Background(), "GET", ep, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF query request for computer group: %v", identifier)
	}

	res := ComputerGroup{}
	if err := j.makeAPIrequest(req, &res); err != nil {
		return nil, errors.Wrapf(err, "unable to query computer group with ID: %d from %s", identifier, ep)
	}
	return &res, nil
}

// UpdateComputerGroupMembers will update the members of a computer group in Jamf by either group ID or group Name
func (j *Client) UpdateComputerGroupMembers(identifier any, updates *ComputerGroupBindingChanges) (*ComputerGroupDetails, error) {
	ep, err := EndpointBuilder(j.Endpoint, computerGroupsContext, identifier)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF query request for computer group: %v", identifier)
	}

	bodyContent, err := xml.Marshal(updates)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF update payload for computer group: %v", identifier)
	}

	body := bytes.NewReader(bodyContent)
	req, err := http.NewRequestWithContext(context.Background(), "PUT", ep, body)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF update request for computer group: %v (%s)", identifier, ep)
	}

	res := ComputerGroupDetails{}
	if err := j.makeAPIrequest(req, &res); err != nil {
		return nil, errors.Wrapf(err, "unable to process JAMF update request for computer group: %v (%s)", identifier, ep)
	}

	return &res, nil
}

func (j *Client) CreateComputerGroup(newGroup *ComputerGroupDetails) (*ComputerGroupDetails, error) {
	ep, err := EndpointBuilder(j.Endpoint, computerGroupsContext, -1)
	if err != nil {
		return nil, errors.Wrap(err, "error building JAMF add computer group request endpoint")
	}

	if newGroup.Name == "" {
		return nil, errors.New("error building JAMF add computer group request: group name is required")
	}

	bodyContent, err := xml.MarshalIndent(newGroup, "", "    ")
	if err != nil {
		return nil, errors.Wrap(err, "error building JAMF add computer group payload")
	}

	body := bytes.NewReader(bodyContent)
	req, err := http.NewRequestWithContext(context.Background(), "POST", ep, body)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF add computer group request")
	}

	res := ComputerGroupDetails{}
	if err := j.makeAPIrequest(req, &res); err != nil {
		return nil, errors.Wrap(err, "unable to process JAMF add computer group request")
	}

	return &res, nil
}

func (j *Client) DeleteComputerGroup(identifier any) (*ComputerGroupDetails, error) {
	ep, err := EndpointBuilder(j.Endpoint, computerGroupsContext, identifier)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF delete computer group request endpoint for group: %v", identifier)
	}

	req, err := http.NewRequestWithContext(context.Background(), "DELETE", ep, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF delete computer group request for group: %v", identifier)
	}
	res := ComputerGroupDetails{}
	if err := j.makeAPIrequest(req, &res); err != nil {
		return nil, errors.Wrapf(err, "unable to process JAMF delete computer group request for group: %v", identifier)
	}

	return &res, nil
}

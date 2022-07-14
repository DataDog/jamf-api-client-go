// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0
// This product includes software developed at Datadog (https://www.datadoghq.com/). Copyright 2020 Datadog, Inc.

package classic

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// ComputerIdentifier include the searchable computer identifiers
type ComputerIdentifier struct {
	ID           string
	Name         string
	SerialNumber string
}

func (identifier *ComputerIdentifier) endpoint(endpoint string, context string) string {
	var (
		entity string
		param  string
	)
	switch {
	case identifier.ID != "":
		entity, param = "id", identifier.ID
	case identifier.Name != "":
		entity, param = "name", identifier.Name
	case identifier.SerialNumber != "":
		entity, param = "serialnumber", identifier.SerialNumber
	}
	return fmt.Sprintf("%s/%s/%s/%s", endpoint, context, entity, param)
}

// Computers returns all enrolled computer devices
func (j *Client) Computers() ([]BasicComputerInfo, error) {
	ep := fmt.Sprintf("%s/%s", j.Endpoint, computersContext)
	req, err := http.NewRequestWithContext(context.Background(), "GET", ep, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error building JAMF computer query request")
	}

	res := &Computers{}
	if err := j.makeAPIrequest(req, &res); err != nil {
		return nil, errors.Wrapf(err, "unable to query enrolled computers from %s", ep)
	}
	return res.List, nil
}

// ComputerDetails returns the details for a specific computer given its ID
func (j *Client) ComputerDetails(identifier interface{}) (*Computer, error) {
	ep, err := EndpointBuilder(j.Endpoint, computersContext, identifier)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF query request endpoint for computer: %v", identifier)
	}
	req, err := http.NewRequestWithContext(context.Background(), "GET", ep, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF computer request for computer: %v (%s)", identifier, ep)
	}

	res := &Computer{}
	if err := j.makeAPIrequest(req, &res); err != nil {
		return nil, errors.Wrapf(err, "unable to query enrolled computer for computer: %v (%s)", identifier, ep)
	}
	return res, nil
}

// GetComputer takes in a search option and returns the details for a specific computer
func (j *Client) GetComputer(identifier *ComputerIdentifier) (*Computer, error) {
	ep := identifier.endpoint(j.Endpoint, computersContext)
	req, err := http.NewRequestWithContext(context.Background(), "GET", ep, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF computer request for computer: %s", ep)
	}

	res := &Computer{}
	if err := j.makeAPIrequest(req, &res); err != nil {
		return nil, errors.Wrapf(err, "unable to query enrolled computer for computer: %s", ep)
	}
	return res, nil
}

// UpdateComputer takes in an identifier and updated content and updates the device on the server
func (j *Client) UpdateComputer(identifier *ComputerIdentifier, updates *ComputerDetails) (*ComputerDetails, error) {
	ep := identifier.endpoint(j.Endpoint, computersContext)
	content, err := xml.Marshal(updates)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF update payload for computer: %v", identifier)
	}

	body := bytes.NewReader(content)
	req, err := http.NewRequestWithContext(context.Background(), "PUT", ep, body)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF update request for computer: %v (%s)", identifier, ep)
	}

	res := ComputerDetails{}
	if err := j.makeAPIrequest(req, &res); err != nil {
		return nil, errors.Wrapf(err, "unable to process JAMF update request for computer: %v (%s)", identifier, ep)
	}
	return &res, nil
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0
// This product includes software developed at Datadog (https://www.datadoghq.com/). Copyright 2020 Datadog, Inc.

package classic

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// GetComputerOptions include the searchable computer identifiers
type GetComputerOptions struct {
	ID           string
	Name         string
	SerialNumber string
}

func (opts *GetComputerOptions) endpoint(endpoint string, context string) string {
	var (
		entity string
		param  string
	)
	switch {
	case opts.ID != "":
		entity, param = "id", opts.ID
	case opts.Name != "":
		entity, param = "name", opts.Name
	case opts.SerialNumber != "":
		entity, param = "serialnumber", opts.SerialNumber
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
func (j *Client) GetComputer(opts *GetComputerOptions) (*Computer, error) {
	ep := opts.endpoint(j.Endpoint, computersContext)
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

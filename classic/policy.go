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

// Policies returns a list of policies available in the jamf client
func (j *Client) Policies() ([]BasicPolicyInformation, error) {
	ep := fmt.Sprintf("%s/%s", j.Endpoint, policiesContext)
	req, err := http.NewRequestWithContext(context.Background(), "GET", ep, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error building Jamf policies query request")
	}
	res := Policies{}
	if err := j.makeAPIrequest(req, &res); err != nil {
		return nil, errors.Wrapf(err, "unable to query available policies from %s", ep)
	}
	return res.List, nil
}

// PolicyDetails returns the details for a specific policy given its ID or Name
func (j *Client) PolicyDetails(identifier interface{}) (*Policy, error) {
	ep, err := EndpointBuilder(j.Endpoint, policiesContext, identifier)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF query request endpoint for policy: %v", identifier)
	}

	req, err := http.NewRequestWithContext(context.Background(), "GET", ep, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF query request for policy: %v", identifier)
	}

	res := Policy{}
	if err := j.makeAPIrequest(req, &res); err != nil {
		return nil, errors.Wrapf(err, "unable to query policy with ID: %d from %s", identifier, ep)
	}
	return &res, nil
}

// UpdatePolicy will update a policy in Jamf by either ID or Name
func (j *Client) UpdatePolicy(identifier interface{}, policy *PolicyContents) (*PolicyContents, error) {
	ep, err := EndpointBuilder(j.Endpoint, policiesContext, identifier)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF query request for policy: %v", identifier)
	}

	if len(policy.Scripts) > 0 {
		policy.ScriptCount = len(policy.Scripts)
		// Priority is required so we will default to After
		for _, s := range policy.Scripts {
			if s.Priority == "" {
				s.Priority = "After"
			}
		}
	}

	bodyContent, err := xml.MarshalIndent(policy, "", "    ")
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF update payload for policy: %v", identifier)
	}

	body := bytes.NewReader(bodyContent)
	req, err := http.NewRequestWithContext(context.Background(), "PUT", ep, body)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF update request for policy: %v (%s)", identifier, ep)
	}

	res := PolicyContents{}
	if err := j.makeAPIrequest(req, &res); err != nil {
		return nil, errors.Wrapf(err, "unable to process JAMF update request for policy: %v (%s)", identifier, ep)
	}

	return &res, nil
}

// CreatePolicy will create a policy in Jamf
func (j *Client) CreatePolicy(content *PolicyContents) (*PolicyContents, error) {
	// -1 denotes the next available ID
	ep, err := EndpointBuilder(j.Endpoint, policiesContext, -1)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF query request for new policy")
	}

	if content.General.Name == "" {
		return nil, errors.Wrapf(fmt.Errorf("Name required for new policy"), "unable to process JAMF creation request for policy: (%s)", ep)
	}

	if len(content.Scripts) > 0 {
		content.ScriptCount = len(content.Scripts)
		// Priority is required so we will default to After
		for _, s := range content.Scripts {
			if s.Priority == "" {
				s.Priority = "After"
			}
		}
	}

	bodyContent, err := xml.Marshal(content)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF creation payload for policy: %v", content.General.Name)
	}

	body := bytes.NewReader(bodyContent)
	req, err := http.NewRequestWithContext(context.Background(), "POST", ep, body)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF creation request for policy: %v (%s)", content.General.Name, ep)
	}
	res := PolicyContents{}

	if err := j.makeAPIrequest(req, &res); err != nil {
		return nil, errors.Wrapf(err, "unable to process JAMF creation request for policy: %v (%s)", content.General.Name, ep)
	}

	return &res, nil
}

// DeletePolicy will delete a policy by either ID or Name
func (j *Client) DeletePolicy(identifier interface{}) (*PolicyGeneral, error) {
	ep, err := EndpointBuilder(j.Endpoint, policiesContext, identifier)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF query request for policy: %v", identifier)
	}

	req, err := http.NewRequestWithContext(context.Background(), "DELETE", ep, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF deletion request for policy: %v (%s)", identifier, ep)
	}

	res := PolicyGeneral{}
	if err := j.makeAPIrequest(req, &res); err != nil {
		return nil, errors.Wrapf(err, "unable to process JAMF deletion request for policy: %v (%s)", identifier, ep)
	}

	return &res, nil
}

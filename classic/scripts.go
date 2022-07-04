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

// Scripts returns a list of scripts available in the jamf client
func (j *Client) Scripts() ([]BasicScriptInfo, error) {
	ep := fmt.Sprintf("%s/%s", j.Endpoint, scriptsContext)
	req, err := http.NewRequestWithContext(context.Background(), "GET", ep, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error building JAMF scripts query request")
	}
	res := Scripts{}
	if err := j.makeAPIrequest(req, &res); err != nil {
		return nil, errors.Wrapf(err, "unable to query available scripts from %s", ep)
	}
	return res.List, nil
}

// ScriptDetails returns the details for a specific script given its ID or Name
func (j *Client) ScriptDetails(identifier interface{}) (*Script, error) {
	ep, err := EndpointBuilder(j.Endpoint, scriptsContext, identifier)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF query request endpoint for script: %v", identifier)
	}

	req, err := http.NewRequestWithContext(context.Background(), "GET", ep, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF query request for script: %v", identifier)
	}

	res := Script{}
	if err := j.makeAPIrequest(req, &res); err != nil {
		return nil, errors.Wrapf(err, "unable to query script with ID: %d from %s", identifier, ep)
	}

	// default to map for script parameters
	if res.Content.Parameters == nil {
		res.Content.Parameters = &ParametersList{}
	}

	return &res, nil
}

// UpdateScript will update a script in Jamf by either ID or Name
func (j *Client) UpdateScript(identifier interface{}, script *ScriptContents) (*ScriptContents, error) {
	ep, err := EndpointBuilder(j.Endpoint, scriptsContext, identifier)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF query request for script: %v", identifier)
	}

	// TODO: Fix hack
	// handle empty parameters since they can come in as
	// map[string]interface{} which can not be handled by xml/encoding
	switch script.Parameters.(type) {
	case map[string]interface{}:
		script.Parameters = &ParametersList{}
	}

	bodyContent, err := xml.Marshal(script)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF update payload for script: %v", identifier)
	}

	body := bytes.NewReader(bodyContent)
	req, err := http.NewRequestWithContext(context.Background(), "PUT", ep, body)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF update request for script: %v (%s)", identifier, ep)
	}

	res := ScriptContents{}
	if err := j.makeAPIrequest(req, &res); err != nil {
		return nil, errors.Wrapf(err, "unable to process JAMF update request for script: %v (%s)", identifier, ep)
	}

	return &res, nil
}

// CreateScript will create a script in Jamf
func (j *Client) CreateScript(content *ScriptContents) (*ScriptContents, error) {
	// -1 denotes the next available ID
	ep, err := EndpointBuilder(j.Endpoint, scriptsContext, -1)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF query request for new script")
	}

	if content.Name == "" {
		return nil, errors.Wrapf(fmt.Errorf("name required for new script"), "unable to process JAMF creation request for script: (%s)", ep)
	}

	if content.Contents == "" {
		return nil, errors.Wrapf(fmt.Errorf("script contents required"), "unable to process JAMF creation request for script: (%s)", ep)
	}

	if content.Filename == "" {
		content.Filename = content.Name
	}

	bodyContent, err := xml.Marshal(content)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF creation payload for script: %v", content.Name)
	}

	body := bytes.NewReader(bodyContent)
	req, err := http.NewRequestWithContext(context.Background(), "POST", ep, body)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF creation request for script: %v (%s)", content.Name, ep)
	}
	res := ScriptContents{}

	if err := j.makeAPIrequest(req, &res); err != nil {
		return nil, errors.Wrapf(err, "unable to process JAMF creation request for script: %v (%s)", content.Name, ep)
	}

	return &res, nil
}

// DeleteScript will delete a script by either ID or Name
func (j *Client) DeleteScript(identifier interface{}) (*ScriptContents, error) {
	ep, err := EndpointBuilder(j.Endpoint, scriptsContext, identifier)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF query request for script: %v", identifier)
	}

	req, err := http.NewRequestWithContext(context.Background(), "DELETE", ep, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF deletion request for script: %v (%s)", identifier, ep)
	}

	res := ScriptContents{}
	if err := j.makeAPIrequest(req, &res); err != nil {
		return nil, errors.Wrapf(err, "unable to process JAMF deletion request for script: %v (%s)", identifier, ep)
	}

	return &res, nil
}

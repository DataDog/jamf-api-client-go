// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0

package classic

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// Classes returns all mobile device classes
func (j *Client) Classes() ([]Class, error) {
	ep := fmt.Sprintf("%s/%s", j.Endpoint, classesContext)
	req, err := http.NewRequestWithContext(context.Background(), "GET", ep, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error building JAMF classes query request")
	}

	res := &Classes{}
	if err := j.makeAPIrequest(req, &res); err != nil {
		return nil, errors.Wrapf(err, "unable to query classes from %s", ep)
	}
	return res.List, nil
}

// ClassDetails returns the details for a specific mobile device class given its ID or Name
func (j *Client) ClassDetails(identifier interface{}) (*ClassDetails, error) {
	ep, err := EndpointBuilder(j.Endpoint, classesContext, identifier)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF query endpoint for class: %v", identifier)
	}
	req, err := http.NewRequestWithContext(context.Background(), "GET", ep, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF query request for class: %v", identifier)
	}

	res := ClassDetails{}
	if err := j.makeAPIrequest(req, &res); err != nil {
		return nil, errors.Wrapf(err, "unable to query class with ID/name %v from %s", identifier, ep)
	}

	return &res, nil
}

// CreateClass will create a new mobile device class in Jamf
func (j *Client) CreateClass(content *Class) (*Class, error) {
	ep, err := EndpointBuilder(j.Endpoint, classesContext, -1)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF query request for new class")
	}

	if content == nil {
		return nil, errors.Wrapf(fmt.Errorf("empty payload"), "unable to process JAMF creation request for class: (%s)", ep)
	}

	if content.Name == "" {
		return nil, errors.Wrapf(fmt.Errorf("name required for new class"), "unable to process JAMF creation request for class: (%s)", ep)
	}

	bodyContent, err := xml.Marshal(content)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF creation payload for class: %v", content.Name)
	}

	body := bytes.NewReader(bodyContent)
	req, err := http.NewRequestWithContext(context.Background(), "POST", ep, body)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF creation request for class: %v (%s)", content.Name, ep)
	}

	res := Class{}
	if err := j.makeAPIrequest(req, &res); err != nil {
		return nil, errors.Wrapf(err, "unable to process JAMF creation request for class %v on %s", content.Name, ep)
	}

	return &res, nil
}

// DeleteClass will delete a mobile device class by either ID or Name
func (j *Client) DeleteClass(identifier interface{}) (*Class, error) {
	ep, err := EndpointBuilder(j.Endpoint, classesContext, identifier)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF query request for class: %v", identifier)
	}

	req, err := http.NewRequestWithContext(context.Background(), "DELETE", ep, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF deletion request for class %v", identifier)
	}

	res := Class{}
	if err := j.makeAPIrequest(req, &res); err != nil {
		return nil, errors.Wrapf(err, "unable to process JAMF deletion request for class %v from %s", identifier, ep)
	}

	return &res, nil
}

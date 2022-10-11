// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0
// This product includes software developed at Datadog (https://www.datadoghq.com/). Copyright 2020 Datadog, Inc.

package classic

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	classesContext         = "classes"
	computersContext       = "computers"
	computerExtAttrContext = "computerextensionattributes"
	policiesContext        = "policies"
	scriptsContext         = "scripts"
)

// Client represents the interface used to communicate with
// the Jamf API via an HTTP client
type Client struct {
	Domain   string
	Username string
	Password string
	Endpoint string
	logger   *logrus.Logger
	api      *http.Client
}

// Used if custom client not passed on when NewClient instantiated
func defaultHTTPClient() *http.Client {
	return &http.Client{
		Timeout: time.Minute,
	}
}

// NewClient returns a new Jamf HTTP client to be used for API requests
func NewClient(domain string, username string, password string, client *http.Client) (*Client, error) {
	if domain == "" || username == "" || password == "" {
		return nil, errors.New("you must provide a valid Jamf domain, username, and password")
	}

	if client == nil {
		client = defaultHTTPClient()
	}

	return &Client{
		Domain:   domain,
		Username: username,
		Password: password,
		Endpoint: fmt.Sprintf("%s/JSSResource", domain),
		api:      client,
	}, nil
}

func (j *Client) makeAPIrequest(r *http.Request, v interface{}) error {
	// Jamf API only sends XML for some endpoints so we will accept both but prioritize
	// JSON responses with the quallity value of 1.0 and 0.9 for XML responses
	// https://developer.mozilla.org/en-US/docs/Glossary/quality_values
	r.Header.Set("Accept", "application/json, application/xml;q=0.9")
	r.Header.Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0, post-check=0, pre-check=0")
	r.Header.Set("Strict-Transport-Security", "max-age=31536000 ; includeSubDomains")
	r.SetBasicAuth(j.Username, j.Password)

	res, err := j.api.Do(r)
	if err != nil {
		return errors.Wrapf(err, "error making %s request to %s", r.Method, r.URL)
	}
	defer res.Body.Close()

	// If status code is not ok attempt to read the response in plain text
	if res.StatusCode != 200 && res.StatusCode != 201 {
		responseData, err := io.ReadAll(res.Body)
		if err != nil {
			return errors.Wrapf(err, "request error: %s. unable to retrieve plain text response: %s", res.Status, err.Error())
		}
		return fmt.Errorf("request error: %s", string(responseData))
	}

	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Content-Type-Options
	// ex. [text/xml charset=UTF-8]
	contentType := strings.Split(res.Header.Get("Content-Type"), ";")
	switch t := contentType[0]; t {
	case "text/xml", "application/xml":
		if err = xml.NewDecoder(res.Body).Decode(&v); err != nil {
			// TODO: return a string or something
			return errors.Wrapf(err, "response was successful but error occured decoding response body of type %s", t)
		}
	case "text/json", "application/json", "text/plain":
		if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
			return errors.Wrapf(err, "response was successful but error occured error decoding response body of type %s", t)
		}
	default:
		return errors.Wrapf(err, "response was successful but error occured recieved unexpected response body of type %s", t)
	}

	return nil
}

// MockAPIRequest is used for testing the API client
func (j *Client) MockAPIRequest(r *http.Request, v interface{}) (*http.Request, error) {
	r.Header.Set("Accept", "application/json,  application/xml;q=0.9")
	r.SetBasicAuth(j.Username, j.Password)
	return r, j.makeAPIrequest(r, v)
}

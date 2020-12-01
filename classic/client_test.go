// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0
// This product includes software developed at Datadog (https://www.datadoghq.com/). Copyright 2020 Datadog, Inc.
package classic_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	jamf "github.com/DataDog/jamf-api-client-go/classic"
	"github.com/stretchr/testify/assert"
)

type MockResponse struct {
	Status string `json:"status"`
}

func clientResponseMock(t *testing.T) *httptest.Server {
	var resp string
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.RequestURI {
		case "/JSSResource/mock/test":
			resp = `{
				"status": "OK"
			}`
		default:
			http.Error(w, fmt.Sprintf("bad API call to %s", r.URL), http.StatusInternalServerError)
			return

		}
		_, err := w.Write([]byte(resp))
		assert.Nil(t, err)
	}))
}
func TestNewClient(t *testing.T) {
	testServer := clientResponseMock(t)
	defer testServer.Close()

	j, err := jamf.NewClient(testServer.URL, "fake-username", "mock-password-cool", nil)
	assert.Nil(t, err)
	assert.Equal(t, "fake-username", j.Username)
	assert.Equal(t, "mock-password-cool", j.Password)
	assert.Equal(t, fmt.Sprintf("%s/JSSResource", testServer.URL), j.Endpoint)

	testResponseURL := fmt.Sprintf("%s/mock/test", j.Endpoint)
	req, err := http.NewRequestWithContext(context.Background(), "GET", testResponseURL, nil)
	assert.Nil(t, err)
	assert.Equal(t, testResponseURL, fmt.Sprintf("%s://%s%s", req.URL.Scheme, req.URL.Host, req.URL.Path))

	statusResponse := &MockResponse{}
	formattedRequest, err := j.MockAPIRequest(req, statusResponse)
	assert.Nil(t, err)
	assert.Equal(t, "application/json, application/xml;q=0.9", formattedRequest.Header.Get("Accept"))

	sentUsername, sentPwd, ok := formattedRequest.BasicAuth()
	assert.True(t, ok)
	assert.Equal(t, j.Username, sentUsername)
	assert.Equal(t, j.Password, sentPwd)
	assert.Equal(t, statusResponse.Status, "OK")
}

func TestBadNewClient(t *testing.T) {
	testServer := clientResponseMock(t)
	defer testServer.Close()
	j, err := jamf.NewClient(testServer.URL, "", "mock-password-cool", nil)
	assert.NotNil(t, err)
	assert.Equal(t, "you must provide a valid Jamf domain, username, and passowrd", err.Error())
	assert.Nil(t, j)
}

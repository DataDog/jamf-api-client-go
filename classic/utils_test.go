// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0
// This product includes software developed at Datadog (https://www.datadoghq.com/). Copyright 2020 Datadog, Inc.

package classic_test

import (
	"fmt"
	"testing"

	jamf "github.com/DataDog/jamf-api-client-go/classic"
	"github.com/stretchr/testify/assert"
)

var testDomain = "https://mock.test.com"
var testContext = "tests"

func TestEndpointBuilderName(t *testing.T) {
	name := "the-one-that-passes"
	expected := fmt.Sprintf("%s/%s/name/%s", testDomain, testContext, name)
	result, err := jamf.EndpointBuilder(testDomain, testContext, name)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestEndpointBuilderID(t *testing.T) {
	id := 87
	expected := fmt.Sprintf("%s/%s/id/%d", testDomain, testContext, id)
	result, err := jamf.EndpointBuilder(testDomain, testContext, id)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestEndpointBuilderInvalid(t *testing.T) {
	invalidType := 1.23
	_, err := jamf.EndpointBuilder(testDomain, testContext, invalidType)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid identifier of type (float64) passed for https://mock.test.com/tests please use name (string) or id (int)", err.Error())
}

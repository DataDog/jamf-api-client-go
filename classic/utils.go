package classic

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// JSONPrettyPrint can be used to pretty print JSON API responses
func JSONPrettyPrint(input []byte) string {
	var out bytes.Buffer
	err := json.Indent(&out, input, "", "\t")
	if err != nil {
		return string(input)
	}
	return out.String()
}

// EndpointBuilder can be utilized to query a specific API context via either name or ID
func EndpointBuilder(endpoint string, context string, identifier interface{}) (string, error) {
	var ep string
	switch identifier.(type) {
	case string:
		ep = fmt.Sprintf("%s/%s/name/%s", endpoint, context, identifier)
	case int:
		ep = fmt.Sprintf("%s/%s/id/%d", endpoint, context, identifier)
	default:
		return "", fmt.Errorf("invalid identifier of type (%v) passed for %s/%s please use name (string) or id (int)", fmt.Sprintf("%T", identifier), endpoint, context)
	}
	return ep, nil
}

/*
 * Pipeline API
 *
 * Pipeline is a feature rich application platform, built for containers on top of Kubernetes to automate the DevOps experience, continuous application development and the lifecycle of deployments. 
 *
 * API version: latest
 * Contact: info@banzaicloud.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package pipeline

import (
	"time"
)

type CreateSecretResponse struct {

	Name string `json:"name"`

	Type string `json:"type"`

	Id string `json:"id"`

	Error string `json:"error,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	UpdatedBy string `json:"updatedBy,omitempty"`

	Version int32 `json:"version,omitempty"`

	Tags []string `json:"tags,omitempty"`
}

// AssertCreateSecretResponseRequired checks if the required fields are not zero-ed
func AssertCreateSecretResponseRequired(obj CreateSecretResponse) error {
	elements := map[string]interface{}{
		"name": obj.Name,
		"type": obj.Type,
		"id": obj.Id,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseCreateSecretResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of CreateSecretResponse (e.g. [][]CreateSecretResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseCreateSecretResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aCreateSecretResponse, ok := obj.(CreateSecretResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertCreateSecretResponseRequired(aCreateSecretResponse)
	})
}

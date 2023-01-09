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

type NodePoolsAzure struct {

	Autoscaling bool `json:"autoscaling,omitempty"`

	Count int32 `json:"count"`

	MinCount int32 `json:"minCount,omitempty"`

	MaxCount int32 `json:"maxCount,omitempty"`

	InstanceType string `json:"instanceType"`

	Labels map[string]string `json:"labels,omitempty"`
}

// AssertNodePoolsAzureRequired checks if the required fields are not zero-ed
func AssertNodePoolsAzureRequired(obj NodePoolsAzure) error {
	elements := map[string]interface{}{
		"count": obj.Count,
		"instanceType": obj.InstanceType,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseNodePoolsAzureRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of NodePoolsAzure (e.g. [][]NodePoolsAzure), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseNodePoolsAzureRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aNodePoolsAzure, ok := obj.(NodePoolsAzure)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertNodePoolsAzureRequired(aNodePoolsAzure)
	})
}

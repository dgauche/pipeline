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

type UpdateIntegratedServiceRequest struct {

	Spec map[string]interface{} `json:"spec"`
}

// AssertUpdateIntegratedServiceRequestRequired checks if the required fields are not zero-ed
func AssertUpdateIntegratedServiceRequestRequired(obj UpdateIntegratedServiceRequest) error {
	elements := map[string]interface{}{
		"spec": obj.Spec,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseUpdateIntegratedServiceRequestRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of UpdateIntegratedServiceRequest (e.g. [][]UpdateIntegratedServiceRequest), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseUpdateIntegratedServiceRequestRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aUpdateIntegratedServiceRequest, ok := obj.(UpdateIntegratedServiceRequest)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertUpdateIntegratedServiceRequestRequired(aUpdateIntegratedServiceRequest)
	})
}

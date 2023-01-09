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

type CreatePkePropertiesKubernetesRbac struct {

	Enabled bool `json:"enabled"`
}

// AssertCreatePkePropertiesKubernetesRbacRequired checks if the required fields are not zero-ed
func AssertCreatePkePropertiesKubernetesRbacRequired(obj CreatePkePropertiesKubernetesRbac) error {
	elements := map[string]interface{}{
		"enabled": obj.Enabled,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseCreatePkePropertiesKubernetesRbacRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of CreatePkePropertiesKubernetesRbac (e.g. [][]CreatePkePropertiesKubernetesRbac), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseCreatePkePropertiesKubernetesRbacRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aCreatePkePropertiesKubernetesRbac, ok := obj.(CreatePkePropertiesKubernetesRbac)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertCreatePkePropertiesKubernetesRbacRequired(aCreatePkePropertiesKubernetesRbac)
	})
}

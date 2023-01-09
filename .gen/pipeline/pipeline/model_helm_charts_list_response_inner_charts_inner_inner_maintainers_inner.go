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

type HelmChartsListResponseInnerChartsInnerInnerMaintainersInner struct {

	Name string `json:"name,omitempty"`

	Email string `json:"email,omitempty"`
}

// AssertHelmChartsListResponseInnerChartsInnerInnerMaintainersInnerRequired checks if the required fields are not zero-ed
func AssertHelmChartsListResponseInnerChartsInnerInnerMaintainersInnerRequired(obj HelmChartsListResponseInnerChartsInnerInnerMaintainersInner) error {
	return nil
}

// AssertRecurseHelmChartsListResponseInnerChartsInnerInnerMaintainersInnerRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of HelmChartsListResponseInnerChartsInnerInnerMaintainersInner (e.g. [][]HelmChartsListResponseInnerChartsInnerInnerMaintainersInner), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseHelmChartsListResponseInnerChartsInnerInnerMaintainersInnerRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aHelmChartsListResponseInnerChartsInnerInnerMaintainersInner, ok := obj.(HelmChartsListResponseInnerChartsInnerInnerMaintainersInner)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertHelmChartsListResponseInnerChartsInnerInnerMaintainersInnerRequired(aHelmChartsListResponseInnerChartsInnerInnerMaintainersInner)
	})
}

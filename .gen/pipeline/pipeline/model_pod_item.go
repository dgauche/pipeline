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

type PodItem struct {

	Name string `json:"name,omitempty"`

	Namespace string `json:"namespace,omitempty"`

	CreatedAt string `json:"createdAt,omitempty"`

	Labels PodItemLabels `json:"labels,omitempty"`

	RestartPolicy string `json:"restartPolicy,omitempty"`

	Conditions []PodCondition `json:"conditions,omitempty"`

	ResourceSummary ResourceSummary `json:"resourceSummary,omitempty"`
}

// AssertPodItemRequired checks if the required fields are not zero-ed
func AssertPodItemRequired(obj PodItem) error {
	if err := AssertPodItemLabelsRequired(obj.Labels); err != nil {
		return err
	}
	for _, el := range obj.Conditions {
		if err := AssertPodConditionRequired(el); err != nil {
			return err
		}
	}
	if err := AssertResourceSummaryRequired(obj.ResourceSummary); err != nil {
		return err
	}
	return nil
}

// AssertRecursePodItemRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of PodItem (e.g. [][]PodItem), otherwise ErrTypeAssertionError is thrown.
func AssertRecursePodItemRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aPodItem, ok := obj.(PodItem)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertPodItemRequired(aPodItem)
	})
}

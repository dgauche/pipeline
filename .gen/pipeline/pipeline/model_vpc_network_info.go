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

type VpcNetworkInfo struct {

	// The IPv4 CIDR blocks assigned to the VPC network
	Cidrs []string `json:"cidrs"`

	// Identifier of the VPC network
	Id string `json:"id"`

	// Name of the VPC network
	Name string `json:"name,omitempty"`
}

// AssertVpcNetworkInfoRequired checks if the required fields are not zero-ed
func AssertVpcNetworkInfoRequired(obj VpcNetworkInfo) error {
	elements := map[string]interface{}{
		"cidrs": obj.Cidrs,
		"id": obj.Id,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseVpcNetworkInfoRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of VpcNetworkInfo (e.g. [][]VpcNetworkInfo), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseVpcNetworkInfoRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aVpcNetworkInfo, ok := obj.(VpcNetworkInfo)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertVpcNetworkInfoRequired(aVpcNetworkInfo)
	})
}

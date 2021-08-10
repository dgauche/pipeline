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

// EksNodePoolVolumes - An associative collection of EKS node pool node instance volume configuration objects keyed by their semantical volume names (example instanceRoot, kubeletRoot).
type EksNodePoolVolumes struct {

	InstanceRoot *EksNodePoolVolume `json:"instanceRoot,omitempty"`

	KubeletRoot *EksNodePoolVolume `json:"kubeletRoot,omitempty"`
}

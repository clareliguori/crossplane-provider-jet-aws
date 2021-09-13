/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type CustomCookbooksSourceObservation struct {
}

type CustomCookbooksSourceParameters struct {

	// +kubebuilder:validation:Optional
	Password *string `json:"password,omitempty" tf:"password"`

	// +kubebuilder:validation:Optional
	Revision *string `json:"revision,omitempty" tf:"revision"`

	// +kubebuilder:validation:Optional
	SSHKey *string `json:"sshKey,omitempty" tf:"ssh_key"`

	// +kubebuilder:validation:Required
	Type string `json:"type" tf:"type"`

	// +kubebuilder:validation:Required
	URL string `json:"url" tf:"url"`

	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username"`
}

type OpsworksStackObservation struct {
	Arn string `json:"arn" tf:"arn"`

	StackEndpoint string `json:"stackEndpoint" tf:"stack_endpoint"`
}

type OpsworksStackParameters struct {

	// +kubebuilder:validation:Optional
	AgentVersion *string `json:"agentVersion,omitempty" tf:"agent_version"`

	// +kubebuilder:validation:Optional
	BerkshelfVersion *string `json:"berkshelfVersion,omitempty" tf:"berkshelf_version"`

	// +kubebuilder:validation:Optional
	Color *string `json:"color,omitempty" tf:"color"`

	// +kubebuilder:validation:Optional
	ConfigurationManagerName *string `json:"configurationManagerName,omitempty" tf:"configuration_manager_name"`

	// +kubebuilder:validation:Optional
	ConfigurationManagerVersion *string `json:"configurationManagerVersion,omitempty" tf:"configuration_manager_version"`

	// +kubebuilder:validation:Optional
	CustomCookbooksSource []CustomCookbooksSourceParameters `json:"customCookbooksSource,omitempty" tf:"custom_cookbooks_source"`

	// +kubebuilder:validation:Optional
	CustomJSON *string `json:"customJson,omitempty" tf:"custom_json"`

	// +kubebuilder:validation:Optional
	DefaultAvailabilityZone *string `json:"defaultAvailabilityZone,omitempty" tf:"default_availability_zone"`

	// +kubebuilder:validation:Required
	DefaultInstanceProfileArn string `json:"defaultInstanceProfileArn" tf:"default_instance_profile_arn"`

	// +kubebuilder:validation:Optional
	DefaultOs *string `json:"defaultOs,omitempty" tf:"default_os"`

	// +kubebuilder:validation:Optional
	DefaultRootDeviceType *string `json:"defaultRootDeviceType,omitempty" tf:"default_root_device_type"`

	// +kubebuilder:validation:Optional
	DefaultSSHKeyName *string `json:"defaultSshKeyName,omitempty" tf:"default_ssh_key_name"`

	// +kubebuilder:validation:Optional
	DefaultSubnetID *string `json:"defaultSubnetId,omitempty" tf:"default_subnet_id"`

	// +kubebuilder:validation:Optional
	HostnameTheme *string `json:"hostnameTheme,omitempty" tf:"hostname_theme"`

	// +kubebuilder:validation:Optional
	ManageBerkshelf *bool `json:"manageBerkshelf,omitempty" tf:"manage_berkshelf"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	ServiceRoleArn string `json:"serviceRoleArn" tf:"service_role_arn"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	// +kubebuilder:validation:Optional
	UseCustomCookbooks *bool `json:"useCustomCookbooks,omitempty" tf:"use_custom_cookbooks"`

	// +kubebuilder:validation:Optional
	UseOpsworksSecurityGroups *bool `json:"useOpsworksSecurityGroups,omitempty" tf:"use_opsworks_security_groups"`

	// +kubebuilder:validation:Optional
	VpcID *string `json:"vpcId,omitempty" tf:"vpc_id"`
}

// OpsworksStackSpec defines the desired state of OpsworksStack
type OpsworksStackSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       OpsworksStackParameters `json:"forProvider"`
}

// OpsworksStackStatus defines the observed state of OpsworksStack.
type OpsworksStackStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          OpsworksStackObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OpsworksStack is the Schema for the OpsworksStacks API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type OpsworksStack struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OpsworksStackSpec   `json:"spec"`
	Status            OpsworksStackStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OpsworksStackList contains a list of OpsworksStacks
type OpsworksStackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpsworksStack `json:"items"`
}

// Repository type metadata.
var (
	OpsworksStackKind             = "OpsworksStack"
	OpsworksStackGroupKind        = schema.GroupKind{Group: Group, Kind: OpsworksStackKind}.String()
	OpsworksStackKindAPIVersion   = OpsworksStackKind + "." + GroupVersion.String()
	OpsworksStackGroupVersionKind = GroupVersion.WithKind(OpsworksStackKind)
)

func init() {
	SchemeBuilder.Register(&OpsworksStack{}, &OpsworksStackList{})
}
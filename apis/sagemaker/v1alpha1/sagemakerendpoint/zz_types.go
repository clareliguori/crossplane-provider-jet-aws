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

// +kubebuilder:object:generate=true
// +groupName=sagemaker.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	apis "github.com/crossplane-contrib/provider-tf-aws/apis"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/sagemaker/v1alpha1"
)

type SagemakerEndpointObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type SagemakerEndpointParameters struct {
	EndpointConfigName string `json:"endpointConfigName" tf:"endpoint_config_name"`

	Name *string `json:"name,omitempty" tf:"name"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// SagemakerEndpointSpec defines the desired state of SagemakerEndpoint
type SagemakerEndpointSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SagemakerEndpointParameters `json:"forProvider"`
}

// SagemakerEndpointStatus defines the observed state of SagemakerEndpoint.
type SagemakerEndpointStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SagemakerEndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SagemakerEndpoint is the Schema for the SagemakerEndpoints API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SagemakerEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SagemakerEndpointSpec   `json:"spec"`
	Status            SagemakerEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SagemakerEndpointList contains a list of SagemakerEndpoints
type SagemakerEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SagemakerEndpoint `json:"items"`
}

// Repository type metadata.
var (
	SagemakerEndpointKind             = "SagemakerEndpoint"
	SagemakerEndpointGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: SagemakerEndpointKind}.String()
	SagemakerEndpointKindAPIVersion   = SagemakerEndpointKind + "." + v1alpha1.GroupVersion.String()
	SagemakerEndpointGroupVersionKind = v1alpha1.GroupVersion.WithKind(SagemakerEndpointKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &SagemakerEndpoint{}, &SagemakerEndpointList{})
}

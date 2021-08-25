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
// +groupName=neptune.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	apis "github.com/crossplane-contrib/provider-tf-aws/apis"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/neptune/v1alpha1"
)

type NeptuneClusterEndpointObservation struct {
	Arn string `json:"arn" tf:"arn"`

	Endpoint string `json:"endpoint" tf:"endpoint"`
}

type NeptuneClusterEndpointParameters struct {
	ClusterEndpointIdentifier string `json:"clusterEndpointIdentifier" tf:"cluster_endpoint_identifier"`

	ClusterIdentifier string `json:"clusterIdentifier" tf:"cluster_identifier"`

	EndpointType string `json:"endpointType" tf:"endpoint_type"`

	ExcludedMembers []string `json:"excludedMembers,omitempty" tf:"excluded_members"`

	StaticMembers []string `json:"staticMembers,omitempty" tf:"static_members"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// NeptuneClusterEndpointSpec defines the desired state of NeptuneClusterEndpoint
type NeptuneClusterEndpointSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       NeptuneClusterEndpointParameters `json:"forProvider"`
}

// NeptuneClusterEndpointStatus defines the observed state of NeptuneClusterEndpoint.
type NeptuneClusterEndpointStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          NeptuneClusterEndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NeptuneClusterEndpoint is the Schema for the NeptuneClusterEndpoints API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type NeptuneClusterEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NeptuneClusterEndpointSpec   `json:"spec"`
	Status            NeptuneClusterEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NeptuneClusterEndpointList contains a list of NeptuneClusterEndpoints
type NeptuneClusterEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NeptuneClusterEndpoint `json:"items"`
}

// Repository type metadata.
var (
	NeptuneClusterEndpointKind             = "NeptuneClusterEndpoint"
	NeptuneClusterEndpointGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: NeptuneClusterEndpointKind}.String()
	NeptuneClusterEndpointKindAPIVersion   = NeptuneClusterEndpointKind + "." + v1alpha1.GroupVersion.String()
	NeptuneClusterEndpointGroupVersionKind = v1alpha1.GroupVersion.WithKind(NeptuneClusterEndpointKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &NeptuneClusterEndpoint{}, &NeptuneClusterEndpointList{})
}

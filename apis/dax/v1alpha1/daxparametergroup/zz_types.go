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
// +groupName=dax.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	apis "github.com/crossplane-contrib/provider-tf-aws/apis"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/dax/v1alpha1"
)

type DaxParameterGroupObservation struct {
}

type DaxParameterGroupParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	Name string `json:"name" tf:"name"`

	Parameters []ParametersParameters `json:"parameters,omitempty" tf:"parameters"`
}

type ParametersObservation struct {
}

type ParametersParameters struct {
	Name string `json:"name" tf:"name"`

	Value string `json:"value" tf:"value"`
}

// DaxParameterGroupSpec defines the desired state of DaxParameterGroup
type DaxParameterGroupSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DaxParameterGroupParameters `json:"forProvider"`
}

// DaxParameterGroupStatus defines the observed state of DaxParameterGroup.
type DaxParameterGroupStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DaxParameterGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DaxParameterGroup is the Schema for the DaxParameterGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DaxParameterGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DaxParameterGroupSpec   `json:"spec"`
	Status            DaxParameterGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DaxParameterGroupList contains a list of DaxParameterGroups
type DaxParameterGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DaxParameterGroup `json:"items"`
}

// Repository type metadata.
var (
	DaxParameterGroupKind             = "DaxParameterGroup"
	DaxParameterGroupGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: DaxParameterGroupKind}.String()
	DaxParameterGroupKindAPIVersion   = DaxParameterGroupKind + "." + v1alpha1.GroupVersion.String()
	DaxParameterGroupGroupVersionKind = v1alpha1.GroupVersion.WithKind(DaxParameterGroupKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &DaxParameterGroup{}, &DaxParameterGroupList{})
}

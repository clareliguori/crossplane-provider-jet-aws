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

type NeptuneSubnetGroupObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type NeptuneSubnetGroupParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	Name *string `json:"name,omitempty" tf:"name"`

	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix"`

	SubnetIds []string `json:"subnetIds" tf:"subnet_ids"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// NeptuneSubnetGroupSpec defines the desired state of NeptuneSubnetGroup
type NeptuneSubnetGroupSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       NeptuneSubnetGroupParameters `json:"forProvider"`
}

// NeptuneSubnetGroupStatus defines the observed state of NeptuneSubnetGroup.
type NeptuneSubnetGroupStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          NeptuneSubnetGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NeptuneSubnetGroup is the Schema for the NeptuneSubnetGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type NeptuneSubnetGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NeptuneSubnetGroupSpec   `json:"spec"`
	Status            NeptuneSubnetGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NeptuneSubnetGroupList contains a list of NeptuneSubnetGroups
type NeptuneSubnetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NeptuneSubnetGroup `json:"items"`
}

// Repository type metadata.
var (
	NeptuneSubnetGroupKind             = "NeptuneSubnetGroup"
	NeptuneSubnetGroupGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: NeptuneSubnetGroupKind}.String()
	NeptuneSubnetGroupKindAPIVersion   = NeptuneSubnetGroupKind + "." + v1alpha1.GroupVersion.String()
	NeptuneSubnetGroupGroupVersionKind = v1alpha1.GroupVersion.WithKind(NeptuneSubnetGroupKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &NeptuneSubnetGroup{}, &NeptuneSubnetGroupList{})
}

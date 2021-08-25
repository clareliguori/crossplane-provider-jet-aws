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
// +groupName=ec2.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	apis "github.com/crossplane-contrib/provider-tf-aws/apis"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/ec2/v1alpha1"
)

type Ec2ManagedPrefixListObservation struct {
	Arn string `json:"arn" tf:"arn"`

	OwnerId string `json:"ownerId" tf:"owner_id"`

	Version int64 `json:"version" tf:"version"`
}

type Ec2ManagedPrefixListParameters struct {
	AddressFamily string `json:"addressFamily" tf:"address_family"`

	Entry []EntryParameters `json:"entry,omitempty" tf:"entry"`

	MaxEntries int64 `json:"maxEntries" tf:"max_entries"`

	Name string `json:"name" tf:"name"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type EntryObservation struct {
}

type EntryParameters struct {
	Cidr string `json:"cidr" tf:"cidr"`

	Description *string `json:"description,omitempty" tf:"description"`
}

// Ec2ManagedPrefixListSpec defines the desired state of Ec2ManagedPrefixList
type Ec2ManagedPrefixListSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Ec2ManagedPrefixListParameters `json:"forProvider"`
}

// Ec2ManagedPrefixListStatus defines the observed state of Ec2ManagedPrefixList.
type Ec2ManagedPrefixListStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Ec2ManagedPrefixListObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2ManagedPrefixList is the Schema for the Ec2ManagedPrefixLists API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Ec2ManagedPrefixList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Ec2ManagedPrefixListSpec   `json:"spec"`
	Status            Ec2ManagedPrefixListStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2ManagedPrefixListList contains a list of Ec2ManagedPrefixLists
type Ec2ManagedPrefixListList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ec2ManagedPrefixList `json:"items"`
}

// Repository type metadata.
var (
	Ec2ManagedPrefixListKind             = "Ec2ManagedPrefixList"
	Ec2ManagedPrefixListGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: Ec2ManagedPrefixListKind}.String()
	Ec2ManagedPrefixListKindAPIVersion   = Ec2ManagedPrefixListKind + "." + v1alpha1.GroupVersion.String()
	Ec2ManagedPrefixListGroupVersionKind = v1alpha1.GroupVersion.WithKind(Ec2ManagedPrefixListKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &Ec2ManagedPrefixList{}, &Ec2ManagedPrefixListList{})
}

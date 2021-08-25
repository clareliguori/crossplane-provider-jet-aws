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
// +groupName=backup.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	apis "github.com/crossplane-contrib/provider-tf-aws/apis"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/backup/v1alpha1"
)

type BackupSelectionObservation struct {
}

type BackupSelectionParameters struct {
	IamRoleArn string `json:"iamRoleArn" tf:"iam_role_arn"`

	Name string `json:"name" tf:"name"`

	PlanId string `json:"planId" tf:"plan_id"`

	Resources []string `json:"resources,omitempty" tf:"resources"`

	SelectionTag []SelectionTagParameters `json:"selectionTag,omitempty" tf:"selection_tag"`
}

type SelectionTagObservation struct {
}

type SelectionTagParameters struct {
	Key string `json:"key" tf:"key"`

	Type string `json:"type" tf:"type"`

	Value string `json:"value" tf:"value"`
}

// BackupSelectionSpec defines the desired state of BackupSelection
type BackupSelectionSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       BackupSelectionParameters `json:"forProvider"`
}

// BackupSelectionStatus defines the observed state of BackupSelection.
type BackupSelectionStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          BackupSelectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BackupSelection is the Schema for the BackupSelections API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type BackupSelection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BackupSelectionSpec   `json:"spec"`
	Status            BackupSelectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupSelectionList contains a list of BackupSelections
type BackupSelectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupSelection `json:"items"`
}

// Repository type metadata.
var (
	BackupSelectionKind             = "BackupSelection"
	BackupSelectionGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: BackupSelectionKind}.String()
	BackupSelectionKindAPIVersion   = BackupSelectionKind + "." + v1alpha1.GroupVersion.String()
	BackupSelectionGroupVersionKind = v1alpha1.GroupVersion.WithKind(BackupSelectionKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &BackupSelection{}, &BackupSelectionList{})
}

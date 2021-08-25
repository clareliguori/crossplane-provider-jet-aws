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
// +groupName=codecommit.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	apis "github.com/crossplane-contrib/provider-tf-aws/apis"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/codecommit/v1alpha1"
)

type CodecommitTriggerObservation struct {
	ConfigurationId string `json:"configurationId" tf:"configuration_id"`
}

type CodecommitTriggerParameters struct {
	RepositoryName string `json:"repositoryName" tf:"repository_name"`

	Trigger []TriggerParameters `json:"trigger" tf:"trigger"`
}

type TriggerObservation struct {
}

type TriggerParameters struct {
	Branches []string `json:"branches,omitempty" tf:"branches"`

	CustomData *string `json:"customData,omitempty" tf:"custom_data"`

	DestinationArn string `json:"destinationArn" tf:"destination_arn"`

	Events []string `json:"events" tf:"events"`

	Name string `json:"name" tf:"name"`
}

// CodecommitTriggerSpec defines the desired state of CodecommitTrigger
type CodecommitTriggerSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CodecommitTriggerParameters `json:"forProvider"`
}

// CodecommitTriggerStatus defines the observed state of CodecommitTrigger.
type CodecommitTriggerStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CodecommitTriggerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CodecommitTrigger is the Schema for the CodecommitTriggers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CodecommitTrigger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CodecommitTriggerSpec   `json:"spec"`
	Status            CodecommitTriggerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CodecommitTriggerList contains a list of CodecommitTriggers
type CodecommitTriggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CodecommitTrigger `json:"items"`
}

// Repository type metadata.
var (
	CodecommitTriggerKind             = "CodecommitTrigger"
	CodecommitTriggerGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: CodecommitTriggerKind}.String()
	CodecommitTriggerKindAPIVersion   = CodecommitTriggerKind + "." + v1alpha1.GroupVersion.String()
	CodecommitTriggerGroupVersionKind = v1alpha1.GroupVersion.WithKind(CodecommitTriggerKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &CodecommitTrigger{}, &CodecommitTriggerList{})
}

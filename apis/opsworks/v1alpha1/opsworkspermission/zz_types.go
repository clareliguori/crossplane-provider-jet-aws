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
// +groupName=opsworks.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	apis "github.com/crossplane-contrib/provider-tf-aws/apis"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/opsworks/v1alpha1"
)

type OpsworksPermissionObservation struct {
}

type OpsworksPermissionParameters struct {
	AllowSsh *bool `json:"allowSsh,omitempty" tf:"allow_ssh"`

	AllowSudo *bool `json:"allowSudo,omitempty" tf:"allow_sudo"`

	Level *string `json:"level,omitempty" tf:"level"`

	StackId *string `json:"stackId,omitempty" tf:"stack_id"`

	UserArn string `json:"userArn" tf:"user_arn"`
}

// OpsworksPermissionSpec defines the desired state of OpsworksPermission
type OpsworksPermissionSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       OpsworksPermissionParameters `json:"forProvider"`
}

// OpsworksPermissionStatus defines the observed state of OpsworksPermission.
type OpsworksPermissionStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          OpsworksPermissionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OpsworksPermission is the Schema for the OpsworksPermissions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type OpsworksPermission struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OpsworksPermissionSpec   `json:"spec"`
	Status            OpsworksPermissionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OpsworksPermissionList contains a list of OpsworksPermissions
type OpsworksPermissionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpsworksPermission `json:"items"`
}

// Repository type metadata.
var (
	OpsworksPermissionKind             = "OpsworksPermission"
	OpsworksPermissionGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: OpsworksPermissionKind}.String()
	OpsworksPermissionKindAPIVersion   = OpsworksPermissionKind + "." + v1alpha1.GroupVersion.String()
	OpsworksPermissionGroupVersionKind = v1alpha1.GroupVersion.WithKind(OpsworksPermissionKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &OpsworksPermission{}, &OpsworksPermissionList{})
}

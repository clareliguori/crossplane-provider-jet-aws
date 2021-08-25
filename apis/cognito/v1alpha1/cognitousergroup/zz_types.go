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
// +groupName=cognito.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	apis "github.com/crossplane-contrib/provider-tf-aws/apis"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/cognito/v1alpha1"
)

type CognitoUserGroupObservation struct {
}

type CognitoUserGroupParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	Name string `json:"name" tf:"name"`

	Precedence *int64 `json:"precedence,omitempty" tf:"precedence"`

	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn"`

	UserPoolId string `json:"userPoolId" tf:"user_pool_id"`
}

// CognitoUserGroupSpec defines the desired state of CognitoUserGroup
type CognitoUserGroupSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CognitoUserGroupParameters `json:"forProvider"`
}

// CognitoUserGroupStatus defines the observed state of CognitoUserGroup.
type CognitoUserGroupStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CognitoUserGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CognitoUserGroup is the Schema for the CognitoUserGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CognitoUserGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CognitoUserGroupSpec   `json:"spec"`
	Status            CognitoUserGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CognitoUserGroupList contains a list of CognitoUserGroups
type CognitoUserGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CognitoUserGroup `json:"items"`
}

// Repository type metadata.
var (
	CognitoUserGroupKind             = "CognitoUserGroup"
	CognitoUserGroupGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: CognitoUserGroupKind}.String()
	CognitoUserGroupKindAPIVersion   = CognitoUserGroupKind + "." + v1alpha1.GroupVersion.String()
	CognitoUserGroupGroupVersionKind = v1alpha1.GroupVersion.WithKind(CognitoUserGroupKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &CognitoUserGroup{}, &CognitoUserGroupList{})
}

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

type CognitoIdentityPoolRolesAttachmentObservation struct {
}

type CognitoIdentityPoolRolesAttachmentParameters struct {
	IdentityPoolId string `json:"identityPoolId" tf:"identity_pool_id"`

	RoleMapping []RoleMappingParameters `json:"roleMapping,omitempty" tf:"role_mapping"`

	Roles map[string]string `json:"roles" tf:"roles"`
}

type MappingRuleObservation struct {
}

type MappingRuleParameters struct {
	Claim string `json:"claim" tf:"claim"`

	MatchType string `json:"matchType" tf:"match_type"`

	RoleArn string `json:"roleArn" tf:"role_arn"`

	Value string `json:"value" tf:"value"`
}

type RoleMappingObservation struct {
}

type RoleMappingParameters struct {
	AmbiguousRoleResolution *string `json:"ambiguousRoleResolution,omitempty" tf:"ambiguous_role_resolution"`

	IdentityProvider string `json:"identityProvider" tf:"identity_provider"`

	MappingRule []MappingRuleParameters `json:"mappingRule,omitempty" tf:"mapping_rule"`

	Type string `json:"type" tf:"type"`
}

// CognitoIdentityPoolRolesAttachmentSpec defines the desired state of CognitoIdentityPoolRolesAttachment
type CognitoIdentityPoolRolesAttachmentSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CognitoIdentityPoolRolesAttachmentParameters `json:"forProvider"`
}

// CognitoIdentityPoolRolesAttachmentStatus defines the observed state of CognitoIdentityPoolRolesAttachment.
type CognitoIdentityPoolRolesAttachmentStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CognitoIdentityPoolRolesAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CognitoIdentityPoolRolesAttachment is the Schema for the CognitoIdentityPoolRolesAttachments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CognitoIdentityPoolRolesAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CognitoIdentityPoolRolesAttachmentSpec   `json:"spec"`
	Status            CognitoIdentityPoolRolesAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CognitoIdentityPoolRolesAttachmentList contains a list of CognitoIdentityPoolRolesAttachments
type CognitoIdentityPoolRolesAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CognitoIdentityPoolRolesAttachment `json:"items"`
}

// Repository type metadata.
var (
	CognitoIdentityPoolRolesAttachmentKind             = "CognitoIdentityPoolRolesAttachment"
	CognitoIdentityPoolRolesAttachmentGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: CognitoIdentityPoolRolesAttachmentKind}.String()
	CognitoIdentityPoolRolesAttachmentKindAPIVersion   = CognitoIdentityPoolRolesAttachmentKind + "." + v1alpha1.GroupVersion.String()
	CognitoIdentityPoolRolesAttachmentGroupVersionKind = v1alpha1.GroupVersion.WithKind(CognitoIdentityPoolRolesAttachmentKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &CognitoIdentityPoolRolesAttachment{}, &CognitoIdentityPoolRolesAttachmentList{})
}

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
// +groupName=kms.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	apis "github.com/crossplane-contrib/provider-tf-aws/apis"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/kms/v1alpha1"
)

type ConstraintsObservation struct {
}

type ConstraintsParameters struct {
	EncryptionContextEquals map[string]string `json:"encryptionContextEquals,omitempty" tf:"encryption_context_equals"`

	EncryptionContextSubset map[string]string `json:"encryptionContextSubset,omitempty" tf:"encryption_context_subset"`
}

type KmsGrantObservation struct {
	GrantId string `json:"grantId" tf:"grant_id"`

	GrantToken string `json:"grantToken" tf:"grant_token"`
}

type KmsGrantParameters struct {
	Constraints []ConstraintsParameters `json:"constraints,omitempty" tf:"constraints"`

	GrantCreationTokens []string `json:"grantCreationTokens,omitempty" tf:"grant_creation_tokens"`

	GranteePrincipal string `json:"granteePrincipal" tf:"grantee_principal"`

	KeyId string `json:"keyId" tf:"key_id"`

	Name *string `json:"name,omitempty" tf:"name"`

	Operations []string `json:"operations" tf:"operations"`

	RetireOnDelete *bool `json:"retireOnDelete,omitempty" tf:"retire_on_delete"`

	RetiringPrincipal *string `json:"retiringPrincipal,omitempty" tf:"retiring_principal"`
}

// KmsGrantSpec defines the desired state of KmsGrant
type KmsGrantSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       KmsGrantParameters `json:"forProvider"`
}

// KmsGrantStatus defines the observed state of KmsGrant.
type KmsGrantStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          KmsGrantObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// KmsGrant is the Schema for the KmsGrants API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type KmsGrant struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KmsGrantSpec   `json:"spec"`
	Status            KmsGrantStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KmsGrantList contains a list of KmsGrants
type KmsGrantList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KmsGrant `json:"items"`
}

// Repository type metadata.
var (
	KmsGrantKind             = "KmsGrant"
	KmsGrantGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: KmsGrantKind}.String()
	KmsGrantKindAPIVersion   = KmsGrantKind + "." + v1alpha1.GroupVersion.String()
	KmsGrantGroupVersionKind = v1alpha1.GroupVersion.WithKind(KmsGrantKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &KmsGrant{}, &KmsGrantList{})
}

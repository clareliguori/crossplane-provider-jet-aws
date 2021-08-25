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

type KmsKeyObservation struct {
	Arn string `json:"arn" tf:"arn"`

	KeyId string `json:"keyId" tf:"key_id"`
}

type KmsKeyParameters struct {
	BypassPolicyLockoutSafetyCheck *bool `json:"bypassPolicyLockoutSafetyCheck,omitempty" tf:"bypass_policy_lockout_safety_check"`

	CustomerMasterKeySpec *string `json:"customerMasterKeySpec,omitempty" tf:"customer_master_key_spec"`

	DeletionWindowInDays *int64 `json:"deletionWindowInDays,omitempty" tf:"deletion_window_in_days"`

	Description *string `json:"description,omitempty" tf:"description"`

	EnableKeyRotation *bool `json:"enableKeyRotation,omitempty" tf:"enable_key_rotation"`

	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled"`

	KeyUsage *string `json:"keyUsage,omitempty" tf:"key_usage"`

	Policy *string `json:"policy,omitempty" tf:"policy"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// KmsKeySpec defines the desired state of KmsKey
type KmsKeySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       KmsKeyParameters `json:"forProvider"`
}

// KmsKeyStatus defines the observed state of KmsKey.
type KmsKeyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          KmsKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// KmsKey is the Schema for the KmsKeys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type KmsKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KmsKeySpec   `json:"spec"`
	Status            KmsKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KmsKeyList contains a list of KmsKeys
type KmsKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KmsKey `json:"items"`
}

// Repository type metadata.
var (
	KmsKeyKind             = "KmsKey"
	KmsKeyGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: KmsKeyKind}.String()
	KmsKeyKindAPIVersion   = KmsKeyKind + "." + v1alpha1.GroupVersion.String()
	KmsKeyGroupVersionKind = v1alpha1.GroupVersion.WithKind(KmsKeyKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &KmsKey{}, &KmsKeyList{})
}

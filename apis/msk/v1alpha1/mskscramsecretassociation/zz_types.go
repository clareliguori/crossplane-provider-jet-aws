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
// +groupName=msk.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	apis "github.com/crossplane-contrib/provider-tf-aws/apis"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/msk/v1alpha1"
)

type MskScramSecretAssociationObservation struct {
}

type MskScramSecretAssociationParameters struct {
	ClusterArn string `json:"clusterArn" tf:"cluster_arn"`

	SecretArnList []string `json:"secretArnList" tf:"secret_arn_list"`
}

// MskScramSecretAssociationSpec defines the desired state of MskScramSecretAssociation
type MskScramSecretAssociationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       MskScramSecretAssociationParameters `json:"forProvider"`
}

// MskScramSecretAssociationStatus defines the observed state of MskScramSecretAssociation.
type MskScramSecretAssociationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          MskScramSecretAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MskScramSecretAssociation is the Schema for the MskScramSecretAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type MskScramSecretAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MskScramSecretAssociationSpec   `json:"spec"`
	Status            MskScramSecretAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MskScramSecretAssociationList contains a list of MskScramSecretAssociations
type MskScramSecretAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MskScramSecretAssociation `json:"items"`
}

// Repository type metadata.
var (
	MskScramSecretAssociationKind             = "MskScramSecretAssociation"
	MskScramSecretAssociationGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: MskScramSecretAssociationKind}.String()
	MskScramSecretAssociationKindAPIVersion   = MskScramSecretAssociationKind + "." + v1alpha1.GroupVersion.String()
	MskScramSecretAssociationGroupVersionKind = v1alpha1.GroupVersion.WithKind(MskScramSecretAssociationKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &MskScramSecretAssociation{}, &MskScramSecretAssociationList{})
}

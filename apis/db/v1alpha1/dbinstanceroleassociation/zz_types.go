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
// +groupName=db.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	apis "github.com/crossplane-contrib/provider-tf-aws/apis"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/db/v1alpha1"
)

type DbInstanceRoleAssociationObservation struct {
}

type DbInstanceRoleAssociationParameters struct {
	DbInstanceIdentifier string `json:"dbInstanceIdentifier" tf:"db_instance_identifier"`

	FeatureName string `json:"featureName" tf:"feature_name"`

	RoleArn string `json:"roleArn" tf:"role_arn"`
}

// DbInstanceRoleAssociationSpec defines the desired state of DbInstanceRoleAssociation
type DbInstanceRoleAssociationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DbInstanceRoleAssociationParameters `json:"forProvider"`
}

// DbInstanceRoleAssociationStatus defines the observed state of DbInstanceRoleAssociation.
type DbInstanceRoleAssociationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DbInstanceRoleAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DbInstanceRoleAssociation is the Schema for the DbInstanceRoleAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DbInstanceRoleAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DbInstanceRoleAssociationSpec   `json:"spec"`
	Status            DbInstanceRoleAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DbInstanceRoleAssociationList contains a list of DbInstanceRoleAssociations
type DbInstanceRoleAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DbInstanceRoleAssociation `json:"items"`
}

// Repository type metadata.
var (
	DbInstanceRoleAssociationKind             = "DbInstanceRoleAssociation"
	DbInstanceRoleAssociationGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: DbInstanceRoleAssociationKind}.String()
	DbInstanceRoleAssociationKindAPIVersion   = DbInstanceRoleAssociationKind + "." + v1alpha1.GroupVersion.String()
	DbInstanceRoleAssociationGroupVersionKind = v1alpha1.GroupVersion.WithKind(DbInstanceRoleAssociationKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &DbInstanceRoleAssociation{}, &DbInstanceRoleAssociationList{})
}

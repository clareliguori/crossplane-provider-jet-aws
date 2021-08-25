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
// +groupName=storagegateway.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	apis "github.com/crossplane-contrib/provider-tf-aws/apis"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/storagegateway/v1alpha1"
)

type CacheAttributesObservation struct {
}

type CacheAttributesParameters struct {
	CacheStaleTimeoutInSeconds *int64 `json:"cacheStaleTimeoutInSeconds,omitempty" tf:"cache_stale_timeout_in_seconds"`
}

type StoragegatewayFileSystemAssociationObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type StoragegatewayFileSystemAssociationParameters struct {
	AuditDestinationArn *string `json:"auditDestinationArn,omitempty" tf:"audit_destination_arn"`

	CacheAttributes []CacheAttributesParameters `json:"cacheAttributes,omitempty" tf:"cache_attributes"`

	GatewayArn string `json:"gatewayArn" tf:"gateway_arn"`

	LocationArn string `json:"locationArn" tf:"location_arn"`

	Password string `json:"password" tf:"password"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	Username string `json:"username" tf:"username"`
}

// StoragegatewayFileSystemAssociationSpec defines the desired state of StoragegatewayFileSystemAssociation
type StoragegatewayFileSystemAssociationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       StoragegatewayFileSystemAssociationParameters `json:"forProvider"`
}

// StoragegatewayFileSystemAssociationStatus defines the observed state of StoragegatewayFileSystemAssociation.
type StoragegatewayFileSystemAssociationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          StoragegatewayFileSystemAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// StoragegatewayFileSystemAssociation is the Schema for the StoragegatewayFileSystemAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type StoragegatewayFileSystemAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StoragegatewayFileSystemAssociationSpec   `json:"spec"`
	Status            StoragegatewayFileSystemAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StoragegatewayFileSystemAssociationList contains a list of StoragegatewayFileSystemAssociations
type StoragegatewayFileSystemAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StoragegatewayFileSystemAssociation `json:"items"`
}

// Repository type metadata.
var (
	StoragegatewayFileSystemAssociationKind             = "StoragegatewayFileSystemAssociation"
	StoragegatewayFileSystemAssociationGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: StoragegatewayFileSystemAssociationKind}.String()
	StoragegatewayFileSystemAssociationKindAPIVersion   = StoragegatewayFileSystemAssociationKind + "." + v1alpha1.GroupVersion.String()
	StoragegatewayFileSystemAssociationGroupVersionKind = v1alpha1.GroupVersion.WithKind(StoragegatewayFileSystemAssociationKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &StoragegatewayFileSystemAssociation{}, &StoragegatewayFileSystemAssociationList{})
}

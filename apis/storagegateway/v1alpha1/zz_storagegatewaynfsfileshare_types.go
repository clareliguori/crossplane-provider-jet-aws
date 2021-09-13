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

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type NfsFileShareDefaultsObservation struct {
}

type NfsFileShareDefaultsParameters struct {

	// +kubebuilder:validation:Optional
	DirectoryMode *string `json:"directoryMode,omitempty" tf:"directory_mode"`

	// +kubebuilder:validation:Optional
	FileMode *string `json:"fileMode,omitempty" tf:"file_mode"`

	// +kubebuilder:validation:Optional
	GroupID *string `json:"groupId,omitempty" tf:"group_id"`

	// +kubebuilder:validation:Optional
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id"`
}

type StoragegatewayNfsFileShareCacheAttributesObservation struct {
}

type StoragegatewayNfsFileShareCacheAttributesParameters struct {

	// +kubebuilder:validation:Optional
	CacheStaleTimeoutInSeconds *int64 `json:"cacheStaleTimeoutInSeconds,omitempty" tf:"cache_stale_timeout_in_seconds"`
}

type StoragegatewayNfsFileShareObservation struct {
	Arn string `json:"arn" tf:"arn"`

	FileshareID string `json:"fileshareId" tf:"fileshare_id"`

	Path string `json:"path" tf:"path"`
}

type StoragegatewayNfsFileShareParameters struct {

	// +kubebuilder:validation:Optional
	CacheAttributes []StoragegatewayNfsFileShareCacheAttributesParameters `json:"cacheAttributes,omitempty" tf:"cache_attributes"`

	// +kubebuilder:validation:Required
	ClientList []string `json:"clientList" tf:"client_list"`

	// +kubebuilder:validation:Optional
	DefaultStorageClass *string `json:"defaultStorageClass,omitempty" tf:"default_storage_class"`

	// +kubebuilder:validation:Optional
	FileShareName *string `json:"fileShareName,omitempty" tf:"file_share_name"`

	// +kubebuilder:validation:Required
	GatewayArn string `json:"gatewayArn" tf:"gateway_arn"`

	// +kubebuilder:validation:Optional
	GuessMimeTypeEnabled *bool `json:"guessMimeTypeEnabled,omitempty" tf:"guess_mime_type_enabled"`

	// +kubebuilder:validation:Optional
	KmsEncrypted *bool `json:"kmsEncrypted,omitempty" tf:"kms_encrypted"`

	// +kubebuilder:validation:Optional
	KmsKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn"`

	// +kubebuilder:validation:Required
	LocationArn string `json:"locationArn" tf:"location_arn"`

	// +kubebuilder:validation:Optional
	NfsFileShareDefaults []NfsFileShareDefaultsParameters `json:"nfsFileShareDefaults,omitempty" tf:"nfs_file_share_defaults"`

	// +kubebuilder:validation:Optional
	NotificationPolicy *string `json:"notificationPolicy,omitempty" tf:"notification_policy"`

	// +kubebuilder:validation:Optional
	ObjectACL *string `json:"objectAcl,omitempty" tf:"object_acl"`

	// +kubebuilder:validation:Optional
	ReadOnly *bool `json:"readOnly,omitempty" tf:"read_only"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	RequesterPays *bool `json:"requesterPays,omitempty" tf:"requester_pays"`

	// +kubebuilder:validation:Required
	RoleArn string `json:"roleArn" tf:"role_arn"`

	// +kubebuilder:validation:Optional
	Squash *string `json:"squash,omitempty" tf:"squash"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// StoragegatewayNfsFileShareSpec defines the desired state of StoragegatewayNfsFileShare
type StoragegatewayNfsFileShareSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       StoragegatewayNfsFileShareParameters `json:"forProvider"`
}

// StoragegatewayNfsFileShareStatus defines the observed state of StoragegatewayNfsFileShare.
type StoragegatewayNfsFileShareStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          StoragegatewayNfsFileShareObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// StoragegatewayNfsFileShare is the Schema for the StoragegatewayNfsFileShares API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type StoragegatewayNfsFileShare struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StoragegatewayNfsFileShareSpec   `json:"spec"`
	Status            StoragegatewayNfsFileShareStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StoragegatewayNfsFileShareList contains a list of StoragegatewayNfsFileShares
type StoragegatewayNfsFileShareList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StoragegatewayNfsFileShare `json:"items"`
}

// Repository type metadata.
var (
	StoragegatewayNfsFileShareKind             = "StoragegatewayNfsFileShare"
	StoragegatewayNfsFileShareGroupKind        = schema.GroupKind{Group: Group, Kind: StoragegatewayNfsFileShareKind}.String()
	StoragegatewayNfsFileShareKindAPIVersion   = StoragegatewayNfsFileShareKind + "." + GroupVersion.String()
	StoragegatewayNfsFileShareGroupVersionKind = GroupVersion.WithKind(StoragegatewayNfsFileShareKind)
)

func init() {
	SchemeBuilder.Register(&StoragegatewayNfsFileShare{}, &StoragegatewayNfsFileShareList{})
}
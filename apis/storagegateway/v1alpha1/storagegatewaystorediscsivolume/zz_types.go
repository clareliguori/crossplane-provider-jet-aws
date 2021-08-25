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

type StoragegatewayStoredIscsiVolumeObservation struct {
	Arn string `json:"arn" tf:"arn"`

	ChapEnabled bool `json:"chapEnabled" tf:"chap_enabled"`

	LunNumber int64 `json:"lunNumber" tf:"lun_number"`

	NetworkInterfacePort int64 `json:"networkInterfacePort" tf:"network_interface_port"`

	TargetArn string `json:"targetArn" tf:"target_arn"`

	VolumeAttachmentStatus string `json:"volumeAttachmentStatus" tf:"volume_attachment_status"`

	VolumeId string `json:"volumeId" tf:"volume_id"`

	VolumeSizeInBytes int64 `json:"volumeSizeInBytes" tf:"volume_size_in_bytes"`

	VolumeStatus string `json:"volumeStatus" tf:"volume_status"`

	VolumeType string `json:"volumeType" tf:"volume_type"`
}

type StoragegatewayStoredIscsiVolumeParameters struct {
	DiskId string `json:"diskId" tf:"disk_id"`

	GatewayArn string `json:"gatewayArn" tf:"gateway_arn"`

	KmsEncrypted *bool `json:"kmsEncrypted,omitempty" tf:"kms_encrypted"`

	KmsKey *string `json:"kmsKey,omitempty" tf:"kms_key"`

	NetworkInterfaceId string `json:"networkInterfaceId" tf:"network_interface_id"`

	PreserveExistingData bool `json:"preserveExistingData" tf:"preserve_existing_data"`

	SnapshotId *string `json:"snapshotId,omitempty" tf:"snapshot_id"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	TargetName string `json:"targetName" tf:"target_name"`
}

// StoragegatewayStoredIscsiVolumeSpec defines the desired state of StoragegatewayStoredIscsiVolume
type StoragegatewayStoredIscsiVolumeSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       StoragegatewayStoredIscsiVolumeParameters `json:"forProvider"`
}

// StoragegatewayStoredIscsiVolumeStatus defines the observed state of StoragegatewayStoredIscsiVolume.
type StoragegatewayStoredIscsiVolumeStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          StoragegatewayStoredIscsiVolumeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// StoragegatewayStoredIscsiVolume is the Schema for the StoragegatewayStoredIscsiVolumes API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type StoragegatewayStoredIscsiVolume struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StoragegatewayStoredIscsiVolumeSpec   `json:"spec"`
	Status            StoragegatewayStoredIscsiVolumeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StoragegatewayStoredIscsiVolumeList contains a list of StoragegatewayStoredIscsiVolumes
type StoragegatewayStoredIscsiVolumeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StoragegatewayStoredIscsiVolume `json:"items"`
}

// Repository type metadata.
var (
	StoragegatewayStoredIscsiVolumeKind             = "StoragegatewayStoredIscsiVolume"
	StoragegatewayStoredIscsiVolumeGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: StoragegatewayStoredIscsiVolumeKind}.String()
	StoragegatewayStoredIscsiVolumeKindAPIVersion   = StoragegatewayStoredIscsiVolumeKind + "." + v1alpha1.GroupVersion.String()
	StoragegatewayStoredIscsiVolumeGroupVersionKind = v1alpha1.GroupVersion.WithKind(StoragegatewayStoredIscsiVolumeKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &StoragegatewayStoredIscsiVolume{}, &StoragegatewayStoredIscsiVolumeList{})
}

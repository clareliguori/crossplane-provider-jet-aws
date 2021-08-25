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
// +groupName=alb.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	apis "github.com/crossplane-contrib/provider-tf-aws/apis"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/alb/v1alpha1"
)

type AlbTargetGroupAttachmentObservation struct {
}

type AlbTargetGroupAttachmentParameters struct {
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone"`

	Port *int64 `json:"port,omitempty" tf:"port"`

	TargetGroupArn string `json:"targetGroupArn" tf:"target_group_arn"`

	TargetId string `json:"targetId" tf:"target_id"`
}

// AlbTargetGroupAttachmentSpec defines the desired state of AlbTargetGroupAttachment
type AlbTargetGroupAttachmentSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       AlbTargetGroupAttachmentParameters `json:"forProvider"`
}

// AlbTargetGroupAttachmentStatus defines the observed state of AlbTargetGroupAttachment.
type AlbTargetGroupAttachmentStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          AlbTargetGroupAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AlbTargetGroupAttachment is the Schema for the AlbTargetGroupAttachments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type AlbTargetGroupAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AlbTargetGroupAttachmentSpec   `json:"spec"`
	Status            AlbTargetGroupAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AlbTargetGroupAttachmentList contains a list of AlbTargetGroupAttachments
type AlbTargetGroupAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AlbTargetGroupAttachment `json:"items"`
}

// Repository type metadata.
var (
	AlbTargetGroupAttachmentKind             = "AlbTargetGroupAttachment"
	AlbTargetGroupAttachmentGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: AlbTargetGroupAttachmentKind}.String()
	AlbTargetGroupAttachmentKindAPIVersion   = AlbTargetGroupAttachmentKind + "." + v1alpha1.GroupVersion.String()
	AlbTargetGroupAttachmentGroupVersionKind = v1alpha1.GroupVersion.WithKind(AlbTargetGroupAttachmentKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &AlbTargetGroupAttachment{}, &AlbTargetGroupAttachmentList{})
}

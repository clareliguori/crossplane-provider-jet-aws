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
// +groupName=kinesisanalyticsv2.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	apis "github.com/crossplane-contrib/provider-tf-aws/apis"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/kinesisanalyticsv2/v1alpha1"
)

type Kinesisanalyticsv2ApplicationSnapshotObservation struct {
	ApplicationVersionId int64 `json:"applicationVersionId" tf:"application_version_id"`

	SnapshotCreationTimestamp string `json:"snapshotCreationTimestamp" tf:"snapshot_creation_timestamp"`
}

type Kinesisanalyticsv2ApplicationSnapshotParameters struct {
	ApplicationName string `json:"applicationName" tf:"application_name"`

	SnapshotName string `json:"snapshotName" tf:"snapshot_name"`
}

// Kinesisanalyticsv2ApplicationSnapshotSpec defines the desired state of Kinesisanalyticsv2ApplicationSnapshot
type Kinesisanalyticsv2ApplicationSnapshotSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Kinesisanalyticsv2ApplicationSnapshotParameters `json:"forProvider"`
}

// Kinesisanalyticsv2ApplicationSnapshotStatus defines the observed state of Kinesisanalyticsv2ApplicationSnapshot.
type Kinesisanalyticsv2ApplicationSnapshotStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Kinesisanalyticsv2ApplicationSnapshotObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Kinesisanalyticsv2ApplicationSnapshot is the Schema for the Kinesisanalyticsv2ApplicationSnapshots API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Kinesisanalyticsv2ApplicationSnapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Kinesisanalyticsv2ApplicationSnapshotSpec   `json:"spec"`
	Status            Kinesisanalyticsv2ApplicationSnapshotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Kinesisanalyticsv2ApplicationSnapshotList contains a list of Kinesisanalyticsv2ApplicationSnapshots
type Kinesisanalyticsv2ApplicationSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Kinesisanalyticsv2ApplicationSnapshot `json:"items"`
}

// Repository type metadata.
var (
	Kinesisanalyticsv2ApplicationSnapshotKind             = "Kinesisanalyticsv2ApplicationSnapshot"
	Kinesisanalyticsv2ApplicationSnapshotGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: Kinesisanalyticsv2ApplicationSnapshotKind}.String()
	Kinesisanalyticsv2ApplicationSnapshotKindAPIVersion   = Kinesisanalyticsv2ApplicationSnapshotKind + "." + v1alpha1.GroupVersion.String()
	Kinesisanalyticsv2ApplicationSnapshotGroupVersionKind = v1alpha1.GroupVersion.WithKind(Kinesisanalyticsv2ApplicationSnapshotKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &Kinesisanalyticsv2ApplicationSnapshot{}, &Kinesisanalyticsv2ApplicationSnapshotList{})
}

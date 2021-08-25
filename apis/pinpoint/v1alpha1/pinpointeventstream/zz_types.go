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
// +groupName=pinpoint.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	apis "github.com/crossplane-contrib/provider-tf-aws/apis"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/pinpoint/v1alpha1"
)

type PinpointEventStreamObservation struct {
}

type PinpointEventStreamParameters struct {
	ApplicationId string `json:"applicationId" tf:"application_id"`

	DestinationStreamArn string `json:"destinationStreamArn" tf:"destination_stream_arn"`

	RoleArn string `json:"roleArn" tf:"role_arn"`
}

// PinpointEventStreamSpec defines the desired state of PinpointEventStream
type PinpointEventStreamSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       PinpointEventStreamParameters `json:"forProvider"`
}

// PinpointEventStreamStatus defines the observed state of PinpointEventStream.
type PinpointEventStreamStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          PinpointEventStreamObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PinpointEventStream is the Schema for the PinpointEventStreams API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type PinpointEventStream struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PinpointEventStreamSpec   `json:"spec"`
	Status            PinpointEventStreamStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PinpointEventStreamList contains a list of PinpointEventStreams
type PinpointEventStreamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PinpointEventStream `json:"items"`
}

// Repository type metadata.
var (
	PinpointEventStreamKind             = "PinpointEventStream"
	PinpointEventStreamGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: PinpointEventStreamKind}.String()
	PinpointEventStreamKindAPIVersion   = PinpointEventStreamKind + "." + v1alpha1.GroupVersion.String()
	PinpointEventStreamGroupVersionKind = v1alpha1.GroupVersion.WithKind(PinpointEventStreamKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &PinpointEventStream{}, &PinpointEventStreamList{})
}

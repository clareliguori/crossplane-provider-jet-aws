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
// +groupName=cloudfront.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	apis "github.com/crossplane-contrib/provider-tf-aws/apis"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/cloudfront/v1alpha1"
)

type CloudfrontRealtimeLogConfigObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type CloudfrontRealtimeLogConfigParameters struct {
	Endpoint []EndpointParameters `json:"endpoint" tf:"endpoint"`

	Fields []string `json:"fields" tf:"fields"`

	Name string `json:"name" tf:"name"`

	SamplingRate int64 `json:"samplingRate" tf:"sampling_rate"`
}

type EndpointObservation struct {
}

type EndpointParameters struct {
	KinesisStreamConfig []KinesisStreamConfigParameters `json:"kinesisStreamConfig" tf:"kinesis_stream_config"`

	StreamType string `json:"streamType" tf:"stream_type"`
}

type KinesisStreamConfigObservation struct {
}

type KinesisStreamConfigParameters struct {
	RoleArn string `json:"roleArn" tf:"role_arn"`

	StreamArn string `json:"streamArn" tf:"stream_arn"`
}

// CloudfrontRealtimeLogConfigSpec defines the desired state of CloudfrontRealtimeLogConfig
type CloudfrontRealtimeLogConfigSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CloudfrontRealtimeLogConfigParameters `json:"forProvider"`
}

// CloudfrontRealtimeLogConfigStatus defines the observed state of CloudfrontRealtimeLogConfig.
type CloudfrontRealtimeLogConfigStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CloudfrontRealtimeLogConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CloudfrontRealtimeLogConfig is the Schema for the CloudfrontRealtimeLogConfigs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CloudfrontRealtimeLogConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudfrontRealtimeLogConfigSpec   `json:"spec"`
	Status            CloudfrontRealtimeLogConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudfrontRealtimeLogConfigList contains a list of CloudfrontRealtimeLogConfigs
type CloudfrontRealtimeLogConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudfrontRealtimeLogConfig `json:"items"`
}

// Repository type metadata.
var (
	CloudfrontRealtimeLogConfigKind             = "CloudfrontRealtimeLogConfig"
	CloudfrontRealtimeLogConfigGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: CloudfrontRealtimeLogConfigKind}.String()
	CloudfrontRealtimeLogConfigKindAPIVersion   = CloudfrontRealtimeLogConfigKind + "." + v1alpha1.GroupVersion.String()
	CloudfrontRealtimeLogConfigGroupVersionKind = v1alpha1.GroupVersion.WithKind(CloudfrontRealtimeLogConfigKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &CloudfrontRealtimeLogConfig{}, &CloudfrontRealtimeLogConfigList{})
}

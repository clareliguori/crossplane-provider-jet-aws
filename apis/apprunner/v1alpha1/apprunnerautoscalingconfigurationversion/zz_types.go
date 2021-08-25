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
// +groupName=apprunner.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	apis "github.com/crossplane-contrib/provider-tf-aws/apis"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/apprunner/v1alpha1"
)

type ApprunnerAutoScalingConfigurationVersionObservation struct {
	Arn string `json:"arn" tf:"arn"`

	AutoScalingConfigurationRevision int64 `json:"autoScalingConfigurationRevision" tf:"auto_scaling_configuration_revision"`

	Latest bool `json:"latest" tf:"latest"`

	Status string `json:"status" tf:"status"`
}

type ApprunnerAutoScalingConfigurationVersionParameters struct {
	AutoScalingConfigurationName string `json:"autoScalingConfigurationName" tf:"auto_scaling_configuration_name"`

	MaxConcurrency *int64 `json:"maxConcurrency,omitempty" tf:"max_concurrency"`

	MaxSize *int64 `json:"maxSize,omitempty" tf:"max_size"`

	MinSize *int64 `json:"minSize,omitempty" tf:"min_size"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// ApprunnerAutoScalingConfigurationVersionSpec defines the desired state of ApprunnerAutoScalingConfigurationVersion
type ApprunnerAutoScalingConfigurationVersionSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ApprunnerAutoScalingConfigurationVersionParameters `json:"forProvider"`
}

// ApprunnerAutoScalingConfigurationVersionStatus defines the observed state of ApprunnerAutoScalingConfigurationVersion.
type ApprunnerAutoScalingConfigurationVersionStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ApprunnerAutoScalingConfigurationVersionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApprunnerAutoScalingConfigurationVersion is the Schema for the ApprunnerAutoScalingConfigurationVersions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ApprunnerAutoScalingConfigurationVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApprunnerAutoScalingConfigurationVersionSpec   `json:"spec"`
	Status            ApprunnerAutoScalingConfigurationVersionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApprunnerAutoScalingConfigurationVersionList contains a list of ApprunnerAutoScalingConfigurationVersions
type ApprunnerAutoScalingConfigurationVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApprunnerAutoScalingConfigurationVersion `json:"items"`
}

// Repository type metadata.
var (
	ApprunnerAutoScalingConfigurationVersionKind             = "ApprunnerAutoScalingConfigurationVersion"
	ApprunnerAutoScalingConfigurationVersionGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: ApprunnerAutoScalingConfigurationVersionKind}.String()
	ApprunnerAutoScalingConfigurationVersionKindAPIVersion   = ApprunnerAutoScalingConfigurationVersionKind + "." + v1alpha1.GroupVersion.String()
	ApprunnerAutoScalingConfigurationVersionGroupVersionKind = v1alpha1.GroupVersion.WithKind(ApprunnerAutoScalingConfigurationVersionKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &ApprunnerAutoScalingConfigurationVersion{}, &ApprunnerAutoScalingConfigurationVersionList{})
}

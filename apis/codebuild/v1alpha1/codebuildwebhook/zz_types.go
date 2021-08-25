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
// +groupName=codebuild.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	apis "github.com/crossplane-contrib/provider-tf-aws/apis"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/codebuild/v1alpha1"
)

type CodebuildWebhookObservation struct {
	PayloadUrl string `json:"payloadUrl" tf:"payload_url"`

	Secret string `json:"secret" tf:"secret"`

	Url string `json:"url" tf:"url"`
}

type CodebuildWebhookParameters struct {
	BranchFilter *string `json:"branchFilter,omitempty" tf:"branch_filter"`

	BuildType *string `json:"buildType,omitempty" tf:"build_type"`

	FilterGroup []FilterGroupParameters `json:"filterGroup,omitempty" tf:"filter_group"`

	ProjectName string `json:"projectName" tf:"project_name"`
}

type FilterGroupObservation struct {
}

type FilterGroupParameters struct {
	Filter []FilterParameters `json:"filter,omitempty" tf:"filter"`
}

type FilterObservation struct {
}

type FilterParameters struct {
	ExcludeMatchedPattern *bool `json:"excludeMatchedPattern,omitempty" tf:"exclude_matched_pattern"`

	Pattern string `json:"pattern" tf:"pattern"`

	Type string `json:"type" tf:"type"`
}

// CodebuildWebhookSpec defines the desired state of CodebuildWebhook
type CodebuildWebhookSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CodebuildWebhookParameters `json:"forProvider"`
}

// CodebuildWebhookStatus defines the observed state of CodebuildWebhook.
type CodebuildWebhookStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CodebuildWebhookObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CodebuildWebhook is the Schema for the CodebuildWebhooks API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CodebuildWebhook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CodebuildWebhookSpec   `json:"spec"`
	Status            CodebuildWebhookStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CodebuildWebhookList contains a list of CodebuildWebhooks
type CodebuildWebhookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CodebuildWebhook `json:"items"`
}

// Repository type metadata.
var (
	CodebuildWebhookKind             = "CodebuildWebhook"
	CodebuildWebhookGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: CodebuildWebhookKind}.String()
	CodebuildWebhookKindAPIVersion   = CodebuildWebhookKind + "." + v1alpha1.GroupVersion.String()
	CodebuildWebhookGroupVersionKind = v1alpha1.GroupVersion.WithKind(CodebuildWebhookKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &CodebuildWebhook{}, &CodebuildWebhookList{})
}

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
// +groupName=codedeploy.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	apis "github.com/crossplane-contrib/provider-tf-aws/apis"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/codedeploy/v1alpha1"
)

type CodedeployAppObservation struct {
	ApplicationId string `json:"applicationId" tf:"application_id"`

	Arn string `json:"arn" tf:"arn"`

	GithubAccountName string `json:"githubAccountName" tf:"github_account_name"`

	LinkedToGithub bool `json:"linkedToGithub" tf:"linked_to_github"`
}

type CodedeployAppParameters struct {
	ComputePlatform *string `json:"computePlatform,omitempty" tf:"compute_platform"`

	Name string `json:"name" tf:"name"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// CodedeployAppSpec defines the desired state of CodedeployApp
type CodedeployAppSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CodedeployAppParameters `json:"forProvider"`
}

// CodedeployAppStatus defines the observed state of CodedeployApp.
type CodedeployAppStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CodedeployAppObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CodedeployApp is the Schema for the CodedeployApps API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CodedeployApp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CodedeployAppSpec   `json:"spec"`
	Status            CodedeployAppStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CodedeployAppList contains a list of CodedeployApps
type CodedeployAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CodedeployApp `json:"items"`
}

// Repository type metadata.
var (
	CodedeployAppKind             = "CodedeployApp"
	CodedeployAppGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: CodedeployAppKind}.String()
	CodedeployAppKindAPIVersion   = CodedeployAppKind + "." + v1alpha1.GroupVersion.String()
	CodedeployAppGroupVersionKind = v1alpha1.GroupVersion.WithKind(CodedeployAppKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &CodedeployApp{}, &CodedeployAppList{})
}

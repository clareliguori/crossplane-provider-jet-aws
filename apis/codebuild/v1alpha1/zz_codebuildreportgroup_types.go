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

type CodebuildReportGroupObservation struct {
	Arn string `json:"arn" tf:"arn"`

	Created string `json:"created" tf:"created"`
}

type CodebuildReportGroupParameters struct {

	// +kubebuilder:validation:Optional
	DeleteReports *bool `json:"deleteReports,omitempty" tf:"delete_reports"`

	// +kubebuilder:validation:Required
	ExportConfig []ExportConfigParameters `json:"exportConfig" tf:"export_config"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	// +kubebuilder:validation:Required
	Type string `json:"type" tf:"type"`
}

type ExportConfigObservation struct {
}

type ExportConfigParameters struct {

	// +kubebuilder:validation:Optional
	S3Destination []S3DestinationParameters `json:"s3Destination,omitempty" tf:"s3_destination"`

	// +kubebuilder:validation:Required
	Type string `json:"type" tf:"type"`
}

type S3DestinationObservation struct {
}

type S3DestinationParameters struct {

	// +kubebuilder:validation:Required
	Bucket string `json:"bucket" tf:"bucket"`

	// +kubebuilder:validation:Optional
	EncryptionDisabled *bool `json:"encryptionDisabled,omitempty" tf:"encryption_disabled"`

	// +kubebuilder:validation:Required
	EncryptionKey string `json:"encryptionKey" tf:"encryption_key"`

	// +kubebuilder:validation:Optional
	Packaging *string `json:"packaging,omitempty" tf:"packaging"`

	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path"`
}

// CodebuildReportGroupSpec defines the desired state of CodebuildReportGroup
type CodebuildReportGroupSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CodebuildReportGroupParameters `json:"forProvider"`
}

// CodebuildReportGroupStatus defines the observed state of CodebuildReportGroup.
type CodebuildReportGroupStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CodebuildReportGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CodebuildReportGroup is the Schema for the CodebuildReportGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type CodebuildReportGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CodebuildReportGroupSpec   `json:"spec"`
	Status            CodebuildReportGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CodebuildReportGroupList contains a list of CodebuildReportGroups
type CodebuildReportGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CodebuildReportGroup `json:"items"`
}

// Repository type metadata.
var (
	CodebuildReportGroupKind             = "CodebuildReportGroup"
	CodebuildReportGroupGroupKind        = schema.GroupKind{Group: Group, Kind: CodebuildReportGroupKind}.String()
	CodebuildReportGroupKindAPIVersion   = CodebuildReportGroupKind + "." + GroupVersion.String()
	CodebuildReportGroupGroupVersionKind = GroupVersion.WithKind(CodebuildReportGroupKind)
)

func init() {
	SchemeBuilder.Register(&CodebuildReportGroup{}, &CodebuildReportGroupList{})
}
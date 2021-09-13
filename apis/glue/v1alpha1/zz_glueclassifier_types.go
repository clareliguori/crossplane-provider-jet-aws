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

type CsvClassifierObservation struct {
}

type CsvClassifierParameters struct {

	// +kubebuilder:validation:Optional
	AllowSingleColumn *bool `json:"allowSingleColumn,omitempty" tf:"allow_single_column"`

	// +kubebuilder:validation:Optional
	ContainsHeader *string `json:"containsHeader,omitempty" tf:"contains_header"`

	// +kubebuilder:validation:Optional
	Delimiter *string `json:"delimiter,omitempty" tf:"delimiter"`

	// +kubebuilder:validation:Optional
	DisableValueTrimming *bool `json:"disableValueTrimming,omitempty" tf:"disable_value_trimming"`

	// +kubebuilder:validation:Optional
	Header []string `json:"header,omitempty" tf:"header"`

	// +kubebuilder:validation:Optional
	QuoteSymbol *string `json:"quoteSymbol,omitempty" tf:"quote_symbol"`
}

type GlueClassifierObservation struct {
}

type GlueClassifierParameters struct {

	// +kubebuilder:validation:Optional
	CsvClassifier []CsvClassifierParameters `json:"csvClassifier,omitempty" tf:"csv_classifier"`

	// +kubebuilder:validation:Optional
	GrokClassifier []GrokClassifierParameters `json:"grokClassifier,omitempty" tf:"grok_classifier"`

	// +kubebuilder:validation:Optional
	JSONClassifier []JSONClassifierParameters `json:"jsonClassifier,omitempty" tf:"json_classifier"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	XMLClassifier []XMLClassifierParameters `json:"xmlClassifier,omitempty" tf:"xml_classifier"`
}

type GrokClassifierObservation struct {
}

type GrokClassifierParameters struct {

	// +kubebuilder:validation:Required
	Classification string `json:"classification" tf:"classification"`

	// +kubebuilder:validation:Optional
	CustomPatterns *string `json:"customPatterns,omitempty" tf:"custom_patterns"`

	// +kubebuilder:validation:Required
	GrokPattern string `json:"grokPattern" tf:"grok_pattern"`
}

type JSONClassifierObservation struct {
}

type JSONClassifierParameters struct {

	// +kubebuilder:validation:Required
	JSONPath string `json:"jsonPath" tf:"json_path"`
}

type XMLClassifierObservation struct {
}

type XMLClassifierParameters struct {

	// +kubebuilder:validation:Required
	Classification string `json:"classification" tf:"classification"`

	// +kubebuilder:validation:Required
	RowTag string `json:"rowTag" tf:"row_tag"`
}

// GlueClassifierSpec defines the desired state of GlueClassifier
type GlueClassifierSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       GlueClassifierParameters `json:"forProvider"`
}

// GlueClassifierStatus defines the observed state of GlueClassifier.
type GlueClassifierStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          GlueClassifierObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GlueClassifier is the Schema for the GlueClassifiers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type GlueClassifier struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlueClassifierSpec   `json:"spec"`
	Status            GlueClassifierStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlueClassifierList contains a list of GlueClassifiers
type GlueClassifierList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlueClassifier `json:"items"`
}

// Repository type metadata.
var (
	GlueClassifierKind             = "GlueClassifier"
	GlueClassifierGroupKind        = schema.GroupKind{Group: Group, Kind: GlueClassifierKind}.String()
	GlueClassifierKindAPIVersion   = GlueClassifierKind + "." + GroupVersion.String()
	GlueClassifierGroupVersionKind = GroupVersion.WithKind(GlueClassifierKind)
)

func init() {
	SchemeBuilder.Register(&GlueClassifier{}, &GlueClassifierList{})
}
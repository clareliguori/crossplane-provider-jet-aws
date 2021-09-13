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

type ActionConditionObservation struct {
}

type ActionConditionParameters struct {

	// +kubebuilder:validation:Required
	Action string `json:"action" tf:"action"`
}

type ConditionObservation struct {
}

type ConditionParameters struct {

	// +kubebuilder:validation:Optional
	ActionCondition []ActionConditionParameters `json:"actionCondition,omitempty" tf:"action_condition"`

	// +kubebuilder:validation:Optional
	LabelNameCondition []LabelNameConditionParameters `json:"labelNameCondition,omitempty" tf:"label_name_condition"`
}

type FilterObservation struct {
}

type FilterParameters struct {

	// +kubebuilder:validation:Required
	Behavior string `json:"behavior" tf:"behavior"`

	// +kubebuilder:validation:Required
	Condition []ConditionParameters `json:"condition" tf:"condition"`

	// +kubebuilder:validation:Required
	Requirement string `json:"requirement" tf:"requirement"`
}

type LabelNameConditionObservation struct {
}

type LabelNameConditionParameters struct {

	// +kubebuilder:validation:Required
	LabelName string `json:"labelName" tf:"label_name"`
}

type LoggingFilterObservation struct {
}

type LoggingFilterParameters struct {

	// +kubebuilder:validation:Required
	DefaultBehavior string `json:"defaultBehavior" tf:"default_behavior"`

	// +kubebuilder:validation:Required
	Filter []FilterParameters `json:"filter" tf:"filter"`
}

type RedactedFieldsAllQueryArgumentsObservation struct {
}

type RedactedFieldsAllQueryArgumentsParameters struct {
}

type RedactedFieldsBodyObservation struct {
}

type RedactedFieldsBodyParameters struct {
}

type RedactedFieldsMethodObservation struct {
}

type RedactedFieldsMethodParameters struct {
}

type RedactedFieldsObservation struct {
}

type RedactedFieldsParameters struct {

	// +kubebuilder:validation:Optional
	AllQueryArguments []RedactedFieldsAllQueryArgumentsParameters `json:"allQueryArguments,omitempty" tf:"all_query_arguments"`

	// +kubebuilder:validation:Optional
	Body []RedactedFieldsBodyParameters `json:"body,omitempty" tf:"body"`

	// +kubebuilder:validation:Optional
	Method []RedactedFieldsMethodParameters `json:"method,omitempty" tf:"method"`

	// +kubebuilder:validation:Optional
	QueryString []RedactedFieldsQueryStringParameters `json:"queryString,omitempty" tf:"query_string"`

	// +kubebuilder:validation:Optional
	SingleHeader []RedactedFieldsSingleHeaderParameters `json:"singleHeader,omitempty" tf:"single_header"`

	// +kubebuilder:validation:Optional
	SingleQueryArgument []RedactedFieldsSingleQueryArgumentParameters `json:"singleQueryArgument,omitempty" tf:"single_query_argument"`

	// +kubebuilder:validation:Optional
	URIPath []RedactedFieldsURIPathParameters `json:"uriPath,omitempty" tf:"uri_path"`
}

type RedactedFieldsQueryStringObservation struct {
}

type RedactedFieldsQueryStringParameters struct {
}

type RedactedFieldsSingleHeaderObservation struct {
}

type RedactedFieldsSingleHeaderParameters struct {

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`
}

type RedactedFieldsSingleQueryArgumentObservation struct {
}

type RedactedFieldsSingleQueryArgumentParameters struct {

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`
}

type RedactedFieldsURIPathObservation struct {
}

type RedactedFieldsURIPathParameters struct {
}

type Wafv2WebAclLoggingConfigurationObservation struct {
}

type Wafv2WebAclLoggingConfigurationParameters struct {

	// AWS Kinesis Firehose Delivery Stream ARNs
	// +kubebuilder:validation:Required
	LogDestinationConfigs []string `json:"logDestinationConfigs" tf:"log_destination_configs"`

	// +kubebuilder:validation:Optional
	LoggingFilter []LoggingFilterParameters `json:"loggingFilter,omitempty" tf:"logging_filter"`

	// Parts of the request to exclude from logs
	// +kubebuilder:validation:Optional
	RedactedFields []RedactedFieldsParameters `json:"redactedFields,omitempty" tf:"redacted_fields"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// AWS WebACL ARN
	// +kubebuilder:validation:Required
	ResourceArn string `json:"resourceArn" tf:"resource_arn"`
}

// Wafv2WebAclLoggingConfigurationSpec defines the desired state of Wafv2WebAclLoggingConfiguration
type Wafv2WebAclLoggingConfigurationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Wafv2WebAclLoggingConfigurationParameters `json:"forProvider"`
}

// Wafv2WebAclLoggingConfigurationStatus defines the observed state of Wafv2WebAclLoggingConfiguration.
type Wafv2WebAclLoggingConfigurationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Wafv2WebAclLoggingConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Wafv2WebAclLoggingConfiguration is the Schema for the Wafv2WebAclLoggingConfigurations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Wafv2WebAclLoggingConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Wafv2WebAclLoggingConfigurationSpec   `json:"spec"`
	Status            Wafv2WebAclLoggingConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Wafv2WebAclLoggingConfigurationList contains a list of Wafv2WebAclLoggingConfigurations
type Wafv2WebAclLoggingConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Wafv2WebAclLoggingConfiguration `json:"items"`
}

// Repository type metadata.
var (
	Wafv2WebAclLoggingConfigurationKind             = "Wafv2WebAclLoggingConfiguration"
	Wafv2WebAclLoggingConfigurationGroupKind        = schema.GroupKind{Group: Group, Kind: Wafv2WebAclLoggingConfigurationKind}.String()
	Wafv2WebAclLoggingConfigurationKindAPIVersion   = Wafv2WebAclLoggingConfigurationKind + "." + GroupVersion.String()
	Wafv2WebAclLoggingConfigurationGroupVersionKind = GroupVersion.WithKind(Wafv2WebAclLoggingConfigurationKind)
)

func init() {
	SchemeBuilder.Register(&Wafv2WebAclLoggingConfiguration{}, &Wafv2WebAclLoggingConfigurationList{})
}
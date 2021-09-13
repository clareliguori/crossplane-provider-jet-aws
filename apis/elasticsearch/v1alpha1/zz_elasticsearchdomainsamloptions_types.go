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

type ElasticsearchDomainSamlOptionsObservation struct {
}

type ElasticsearchDomainSamlOptionsParameters struct {

	// +kubebuilder:validation:Required
	DomainName string `json:"domainName" tf:"domain_name"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SamlOptions []SamlOptionsParameters `json:"samlOptions,omitempty" tf:"saml_options"`
}

type IdpObservation struct {
}

type IdpParameters struct {

	// +kubebuilder:validation:Required
	EntityID string `json:"entityId" tf:"entity_id"`

	// +kubebuilder:validation:Required
	MetadataContent string `json:"metadataContent" tf:"metadata_content"`
}

type SamlOptionsObservation struct {
}

type SamlOptionsParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	// +kubebuilder:validation:Optional
	Idp []IdpParameters `json:"idp,omitempty" tf:"idp"`

	// +kubebuilder:validation:Optional
	MasterBackendRole *string `json:"masterBackendRole,omitempty" tf:"master_backend_role"`

	// +kubebuilder:validation:Optional
	MasterUserName *string `json:"masterUserName,omitempty" tf:"master_user_name"`

	// +kubebuilder:validation:Optional
	RolesKey *string `json:"rolesKey,omitempty" tf:"roles_key"`

	// +kubebuilder:validation:Optional
	SessionTimeoutMinutes *int64 `json:"sessionTimeoutMinutes,omitempty" tf:"session_timeout_minutes"`

	// +kubebuilder:validation:Optional
	SubjectKey *string `json:"subjectKey,omitempty" tf:"subject_key"`
}

// ElasticsearchDomainSamlOptionsSpec defines the desired state of ElasticsearchDomainSamlOptions
type ElasticsearchDomainSamlOptionsSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ElasticsearchDomainSamlOptionsParameters `json:"forProvider"`
}

// ElasticsearchDomainSamlOptionsStatus defines the observed state of ElasticsearchDomainSamlOptions.
type ElasticsearchDomainSamlOptionsStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ElasticsearchDomainSamlOptionsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ElasticsearchDomainSamlOptions is the Schema for the ElasticsearchDomainSamlOptionss API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type ElasticsearchDomainSamlOptions struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElasticsearchDomainSamlOptionsSpec   `json:"spec"`
	Status            ElasticsearchDomainSamlOptionsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ElasticsearchDomainSamlOptionsList contains a list of ElasticsearchDomainSamlOptionss
type ElasticsearchDomainSamlOptionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ElasticsearchDomainSamlOptions `json:"items"`
}

// Repository type metadata.
var (
	ElasticsearchDomainSamlOptionsKind             = "ElasticsearchDomainSamlOptions"
	ElasticsearchDomainSamlOptionsGroupKind        = schema.GroupKind{Group: Group, Kind: ElasticsearchDomainSamlOptionsKind}.String()
	ElasticsearchDomainSamlOptionsKindAPIVersion   = ElasticsearchDomainSamlOptionsKind + "." + GroupVersion.String()
	ElasticsearchDomainSamlOptionsGroupVersionKind = GroupVersion.WithKind(ElasticsearchDomainSamlOptionsKind)
)

func init() {
	SchemeBuilder.Register(&ElasticsearchDomainSamlOptions{}, &ElasticsearchDomainSamlOptionsList{})
}
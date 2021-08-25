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
// +groupName=servicecatalog.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	apis "github.com/crossplane-contrib/provider-tf-aws/apis"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/servicecatalog/v1alpha1"
)

type ServicecatalogPortfolioShareObservation struct {
	Accepted bool `json:"accepted" tf:"accepted"`
}

type ServicecatalogPortfolioShareParameters struct {
	AcceptLanguage *string `json:"acceptLanguage,omitempty" tf:"accept_language"`

	PortfolioId string `json:"portfolioId" tf:"portfolio_id"`

	PrincipalId string `json:"principalId" tf:"principal_id"`

	ShareTagOptions *bool `json:"shareTagOptions,omitempty" tf:"share_tag_options"`

	Type string `json:"type" tf:"type"`

	WaitForAcceptance *bool `json:"waitForAcceptance,omitempty" tf:"wait_for_acceptance"`
}

// ServicecatalogPortfolioShareSpec defines the desired state of ServicecatalogPortfolioShare
type ServicecatalogPortfolioShareSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ServicecatalogPortfolioShareParameters `json:"forProvider"`
}

// ServicecatalogPortfolioShareStatus defines the observed state of ServicecatalogPortfolioShare.
type ServicecatalogPortfolioShareStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ServicecatalogPortfolioShareObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServicecatalogPortfolioShare is the Schema for the ServicecatalogPortfolioShares API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ServicecatalogPortfolioShare struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServicecatalogPortfolioShareSpec   `json:"spec"`
	Status            ServicecatalogPortfolioShareStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServicecatalogPortfolioShareList contains a list of ServicecatalogPortfolioShares
type ServicecatalogPortfolioShareList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServicecatalogPortfolioShare `json:"items"`
}

// Repository type metadata.
var (
	ServicecatalogPortfolioShareKind             = "ServicecatalogPortfolioShare"
	ServicecatalogPortfolioShareGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: ServicecatalogPortfolioShareKind}.String()
	ServicecatalogPortfolioShareKindAPIVersion   = ServicecatalogPortfolioShareKind + "." + v1alpha1.GroupVersion.String()
	ServicecatalogPortfolioShareGroupVersionKind = v1alpha1.GroupVersion.WithKind(ServicecatalogPortfolioShareKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &ServicecatalogPortfolioShare{}, &ServicecatalogPortfolioShareList{})
}

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

type ServicecatalogPortfolioObservation struct {
	Arn string `json:"arn" tf:"arn"`

	CreatedTime string `json:"createdTime" tf:"created_time"`
}

type ServicecatalogPortfolioParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	Name string `json:"name" tf:"name"`

	ProviderName string `json:"providerName" tf:"provider_name"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// ServicecatalogPortfolioSpec defines the desired state of ServicecatalogPortfolio
type ServicecatalogPortfolioSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ServicecatalogPortfolioParameters `json:"forProvider"`
}

// ServicecatalogPortfolioStatus defines the observed state of ServicecatalogPortfolio.
type ServicecatalogPortfolioStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ServicecatalogPortfolioObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServicecatalogPortfolio is the Schema for the ServicecatalogPortfolios API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ServicecatalogPortfolio struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServicecatalogPortfolioSpec   `json:"spec"`
	Status            ServicecatalogPortfolioStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServicecatalogPortfolioList contains a list of ServicecatalogPortfolios
type ServicecatalogPortfolioList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServicecatalogPortfolio `json:"items"`
}

// Repository type metadata.
var (
	ServicecatalogPortfolioKind             = "ServicecatalogPortfolio"
	ServicecatalogPortfolioGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: ServicecatalogPortfolioKind}.String()
	ServicecatalogPortfolioKindAPIVersion   = ServicecatalogPortfolioKind + "." + v1alpha1.GroupVersion.String()
	ServicecatalogPortfolioGroupVersionKind = v1alpha1.GroupVersion.WithKind(ServicecatalogPortfolioKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &ServicecatalogPortfolio{}, &ServicecatalogPortfolioList{})
}

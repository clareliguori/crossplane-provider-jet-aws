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
// +groupName=apigatewayv2.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	apis "github.com/crossplane-contrib/provider-tf-aws/apis"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/apigatewayv2/v1alpha1"
)

type Apigatewayv2RouteResponseObservation struct {
}

type Apigatewayv2RouteResponseParameters struct {
	ApiId string `json:"apiId" tf:"api_id"`

	ModelSelectionExpression *string `json:"modelSelectionExpression,omitempty" tf:"model_selection_expression"`

	ResponseModels map[string]string `json:"responseModels,omitempty" tf:"response_models"`

	RouteId string `json:"routeId" tf:"route_id"`

	RouteResponseKey string `json:"routeResponseKey" tf:"route_response_key"`
}

// Apigatewayv2RouteResponseSpec defines the desired state of Apigatewayv2RouteResponse
type Apigatewayv2RouteResponseSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Apigatewayv2RouteResponseParameters `json:"forProvider"`
}

// Apigatewayv2RouteResponseStatus defines the observed state of Apigatewayv2RouteResponse.
type Apigatewayv2RouteResponseStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Apigatewayv2RouteResponseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Apigatewayv2RouteResponse is the Schema for the Apigatewayv2RouteResponses API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Apigatewayv2RouteResponse struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Apigatewayv2RouteResponseSpec   `json:"spec"`
	Status            Apigatewayv2RouteResponseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Apigatewayv2RouteResponseList contains a list of Apigatewayv2RouteResponses
type Apigatewayv2RouteResponseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Apigatewayv2RouteResponse `json:"items"`
}

// Repository type metadata.
var (
	Apigatewayv2RouteResponseKind             = "Apigatewayv2RouteResponse"
	Apigatewayv2RouteResponseGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: Apigatewayv2RouteResponseKind}.String()
	Apigatewayv2RouteResponseKindAPIVersion   = Apigatewayv2RouteResponseKind + "." + v1alpha1.GroupVersion.String()
	Apigatewayv2RouteResponseGroupVersionKind = v1alpha1.GroupVersion.WithKind(Apigatewayv2RouteResponseKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &Apigatewayv2RouteResponse{}, &Apigatewayv2RouteResponseList{})
}

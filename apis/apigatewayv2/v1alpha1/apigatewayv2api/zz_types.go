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

type Apigatewayv2ApiObservation struct {
	ApiEndpoint string `json:"apiEndpoint" tf:"api_endpoint"`

	Arn string `json:"arn" tf:"arn"`

	ExecutionArn string `json:"executionArn" tf:"execution_arn"`
}

type Apigatewayv2ApiParameters struct {
	ApiKeySelectionExpression *string `json:"apiKeySelectionExpression,omitempty" tf:"api_key_selection_expression"`

	Body *string `json:"body,omitempty" tf:"body"`

	CorsConfiguration []CorsConfigurationParameters `json:"corsConfiguration,omitempty" tf:"cors_configuration"`

	CredentialsArn *string `json:"credentialsArn,omitempty" tf:"credentials_arn"`

	Description *string `json:"description,omitempty" tf:"description"`

	DisableExecuteApiEndpoint *bool `json:"disableExecuteApiEndpoint,omitempty" tf:"disable_execute_api_endpoint"`

	FailOnWarnings *bool `json:"failOnWarnings,omitempty" tf:"fail_on_warnings"`

	Name string `json:"name" tf:"name"`

	ProtocolType string `json:"protocolType" tf:"protocol_type"`

	RouteKey *string `json:"routeKey,omitempty" tf:"route_key"`

	RouteSelectionExpression *string `json:"routeSelectionExpression,omitempty" tf:"route_selection_expression"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	Target *string `json:"target,omitempty" tf:"target"`

	Version *string `json:"version,omitempty" tf:"version"`
}

type CorsConfigurationObservation struct {
}

type CorsConfigurationParameters struct {
	AllowCredentials *bool `json:"allowCredentials,omitempty" tf:"allow_credentials"`

	AllowHeaders []string `json:"allowHeaders,omitempty" tf:"allow_headers"`

	AllowMethods []string `json:"allowMethods,omitempty" tf:"allow_methods"`

	AllowOrigins []string `json:"allowOrigins,omitempty" tf:"allow_origins"`

	ExposeHeaders []string `json:"exposeHeaders,omitempty" tf:"expose_headers"`

	MaxAge *int64 `json:"maxAge,omitempty" tf:"max_age"`
}

// Apigatewayv2ApiSpec defines the desired state of Apigatewayv2Api
type Apigatewayv2ApiSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Apigatewayv2ApiParameters `json:"forProvider"`
}

// Apigatewayv2ApiStatus defines the observed state of Apigatewayv2Api.
type Apigatewayv2ApiStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Apigatewayv2ApiObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Apigatewayv2Api is the Schema for the Apigatewayv2Apis API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Apigatewayv2Api struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Apigatewayv2ApiSpec   `json:"spec"`
	Status            Apigatewayv2ApiStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Apigatewayv2ApiList contains a list of Apigatewayv2Apis
type Apigatewayv2ApiList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Apigatewayv2Api `json:"items"`
}

// Repository type metadata.
var (
	Apigatewayv2ApiKind             = "Apigatewayv2Api"
	Apigatewayv2ApiGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: Apigatewayv2ApiKind}.String()
	Apigatewayv2ApiKindAPIVersion   = Apigatewayv2ApiKind + "." + v1alpha1.GroupVersion.String()
	Apigatewayv2ApiGroupVersionKind = v1alpha1.GroupVersion.WithKind(Apigatewayv2ApiKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &Apigatewayv2Api{}, &Apigatewayv2ApiList{})
}

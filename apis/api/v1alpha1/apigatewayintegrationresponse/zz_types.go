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
// +groupName=api.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	apis "github.com/crossplane-contrib/provider-tf-aws/apis"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/api/v1alpha1"
)

type ApiGatewayIntegrationResponseObservation struct {
}

type ApiGatewayIntegrationResponseParameters struct {
	ContentHandling *string `json:"contentHandling,omitempty" tf:"content_handling"`

	HttpMethod string `json:"httpMethod" tf:"http_method"`

	ResourceId string `json:"resourceId" tf:"resource_id"`

	ResponseParameters map[string]string `json:"responseParameters,omitempty" tf:"response_parameters"`

	ResponseTemplates map[string]string `json:"responseTemplates,omitempty" tf:"response_templates"`

	RestApiId string `json:"restApiId" tf:"rest_api_id"`

	SelectionPattern *string `json:"selectionPattern,omitempty" tf:"selection_pattern"`

	StatusCode string `json:"statusCode" tf:"status_code"`
}

// ApiGatewayIntegrationResponseSpec defines the desired state of ApiGatewayIntegrationResponse
type ApiGatewayIntegrationResponseSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ApiGatewayIntegrationResponseParameters `json:"forProvider"`
}

// ApiGatewayIntegrationResponseStatus defines the observed state of ApiGatewayIntegrationResponse.
type ApiGatewayIntegrationResponseStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ApiGatewayIntegrationResponseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayIntegrationResponse is the Schema for the ApiGatewayIntegrationResponses API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ApiGatewayIntegrationResponse struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayIntegrationResponseSpec   `json:"spec"`
	Status            ApiGatewayIntegrationResponseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayIntegrationResponseList contains a list of ApiGatewayIntegrationResponses
type ApiGatewayIntegrationResponseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiGatewayIntegrationResponse `json:"items"`
}

// Repository type metadata.
var (
	ApiGatewayIntegrationResponseKind             = "ApiGatewayIntegrationResponse"
	ApiGatewayIntegrationResponseGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: ApiGatewayIntegrationResponseKind}.String()
	ApiGatewayIntegrationResponseKindAPIVersion   = ApiGatewayIntegrationResponseKind + "." + v1alpha1.GroupVersion.String()
	ApiGatewayIntegrationResponseGroupVersionKind = v1alpha1.GroupVersion.WithKind(ApiGatewayIntegrationResponseKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &ApiGatewayIntegrationResponse{}, &ApiGatewayIntegrationResponseList{})
}

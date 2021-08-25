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

type ApiGatewayMethodSettingsObservation struct {
}

type ApiGatewayMethodSettingsParameters struct {
	MethodPath string `json:"methodPath" tf:"method_path"`

	RestApiId string `json:"restApiId" tf:"rest_api_id"`

	Settings []SettingsParameters `json:"settings" tf:"settings"`

	StageName string `json:"stageName" tf:"stage_name"`
}

type SettingsObservation struct {
}

type SettingsParameters struct {
	CacheDataEncrypted *bool `json:"cacheDataEncrypted,omitempty" tf:"cache_data_encrypted"`

	CacheTtlInSeconds *int64 `json:"cacheTtlInSeconds,omitempty" tf:"cache_ttl_in_seconds"`

	CachingEnabled *bool `json:"cachingEnabled,omitempty" tf:"caching_enabled"`

	DataTraceEnabled *bool `json:"dataTraceEnabled,omitempty" tf:"data_trace_enabled"`

	LoggingLevel *string `json:"loggingLevel,omitempty" tf:"logging_level"`

	MetricsEnabled *bool `json:"metricsEnabled,omitempty" tf:"metrics_enabled"`

	RequireAuthorizationForCacheControl *bool `json:"requireAuthorizationForCacheControl,omitempty" tf:"require_authorization_for_cache_control"`

	ThrottlingBurstLimit *int64 `json:"throttlingBurstLimit,omitempty" tf:"throttling_burst_limit"`

	ThrottlingRateLimit *float64 `json:"throttlingRateLimit,omitempty" tf:"throttling_rate_limit"`

	UnauthorizedCacheControlHeaderStrategy *string `json:"unauthorizedCacheControlHeaderStrategy,omitempty" tf:"unauthorized_cache_control_header_strategy"`
}

// ApiGatewayMethodSettingsSpec defines the desired state of ApiGatewayMethodSettings
type ApiGatewayMethodSettingsSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ApiGatewayMethodSettingsParameters `json:"forProvider"`
}

// ApiGatewayMethodSettingsStatus defines the observed state of ApiGatewayMethodSettings.
type ApiGatewayMethodSettingsStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ApiGatewayMethodSettingsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayMethodSettings is the Schema for the ApiGatewayMethodSettingss API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ApiGatewayMethodSettings struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayMethodSettingsSpec   `json:"spec"`
	Status            ApiGatewayMethodSettingsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayMethodSettingsList contains a list of ApiGatewayMethodSettingss
type ApiGatewayMethodSettingsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiGatewayMethodSettings `json:"items"`
}

// Repository type metadata.
var (
	ApiGatewayMethodSettingsKind             = "ApiGatewayMethodSettings"
	ApiGatewayMethodSettingsGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: ApiGatewayMethodSettingsKind}.String()
	ApiGatewayMethodSettingsKindAPIVersion   = ApiGatewayMethodSettingsKind + "." + v1alpha1.GroupVersion.String()
	ApiGatewayMethodSettingsGroupVersionKind = v1alpha1.GroupVersion.WithKind(ApiGatewayMethodSettingsKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &ApiGatewayMethodSettings{}, &ApiGatewayMethodSettingsList{})
}

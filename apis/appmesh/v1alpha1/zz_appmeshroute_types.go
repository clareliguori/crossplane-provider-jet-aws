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

type ActionWeightedTargetObservation struct {
}

type ActionWeightedTargetParameters struct {

	// +kubebuilder:validation:Required
	VirtualNode string `json:"virtualNode" tf:"virtual_node"`

	// +kubebuilder:validation:Required
	Weight int64 `json:"weight" tf:"weight"`
}

type AppmeshRouteObservation struct {
	Arn string `json:"arn" tf:"arn"`

	CreatedDate string `json:"createdDate" tf:"created_date"`

	LastUpdatedDate string `json:"lastUpdatedDate" tf:"last_updated_date"`

	ResourceOwner string `json:"resourceOwner" tf:"resource_owner"`
}

type AppmeshRouteParameters struct {

	// +kubebuilder:validation:Required
	MeshName string `json:"meshName" tf:"mesh_name"`

	// +kubebuilder:validation:Optional
	MeshOwner *string `json:"meshOwner,omitempty" tf:"mesh_owner"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	Spec []AppmeshRouteSpecParameters `json:"spec" tf:"spec"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	// +kubebuilder:validation:Required
	VirtualRouterName string `json:"virtualRouterName" tf:"virtual_router_name"`
}

type AppmeshRouteSpecObservation struct {
}

type AppmeshRouteSpecParameters struct {

	// +kubebuilder:validation:Optional
	GrpcRoute []SpecGrpcRouteParameters `json:"grpcRoute,omitempty" tf:"grpc_route"`

	// +kubebuilder:validation:Optional
	HTTPRoute []SpecHTTPRouteParameters `json:"httpRoute,omitempty" tf:"http_route"`

	// +kubebuilder:validation:Optional
	Http2Route []SpecHttp2RouteParameters `json:"http2Route,omitempty" tf:"http2_route"`

	// +kubebuilder:validation:Optional
	Priority *int64 `json:"priority,omitempty" tf:"priority"`

	// +kubebuilder:validation:Optional
	TCPRoute []TCPRouteParameters `json:"tcpRoute,omitempty" tf:"tcp_route"`
}

type GrpcRouteActionObservation struct {
}

type GrpcRouteActionParameters struct {

	// +kubebuilder:validation:Required
	WeightedTarget []WeightedTargetParameters `json:"weightedTarget" tf:"weighted_target"`
}

type GrpcRouteMatchObservation struct {
}

type GrpcRouteMatchParameters struct {

	// +kubebuilder:validation:Optional
	Metadata []MetadataParameters `json:"metadata,omitempty" tf:"metadata"`

	// +kubebuilder:validation:Optional
	MethodName *string `json:"methodName,omitempty" tf:"method_name"`

	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix"`

	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name"`
}

type HTTPRouteActionWeightedTargetObservation struct {
}

type HTTPRouteActionWeightedTargetParameters struct {

	// +kubebuilder:validation:Required
	VirtualNode string `json:"virtualNode" tf:"virtual_node"`

	// +kubebuilder:validation:Required
	Weight int64 `json:"weight" tf:"weight"`
}

type HTTPRouteRetryPolicyObservation struct {
}

type HTTPRouteRetryPolicyParameters struct {

	// +kubebuilder:validation:Optional
	HTTPRetryEvents []string `json:"httpRetryEvents,omitempty" tf:"http_retry_events"`

	// +kubebuilder:validation:Required
	MaxRetries int64 `json:"maxRetries" tf:"max_retries"`

	// +kubebuilder:validation:Required
	PerRetryTimeout []HTTPRouteRetryPolicyPerRetryTimeoutParameters `json:"perRetryTimeout" tf:"per_retry_timeout"`

	// +kubebuilder:validation:Optional
	TCPRetryEvents []string `json:"tcpRetryEvents,omitempty" tf:"tcp_retry_events"`
}

type HTTPRouteRetryPolicyPerRetryTimeoutObservation struct {
}

type HTTPRouteRetryPolicyPerRetryTimeoutParameters struct {

	// +kubebuilder:validation:Required
	Unit string `json:"unit" tf:"unit"`

	// +kubebuilder:validation:Required
	Value int64 `json:"value" tf:"value"`
}

type HTTPRouteTimeoutIdleObservation struct {
}

type HTTPRouteTimeoutIdleParameters struct {

	// +kubebuilder:validation:Required
	Unit string `json:"unit" tf:"unit"`

	// +kubebuilder:validation:Required
	Value int64 `json:"value" tf:"value"`
}

type HTTPRouteTimeoutObservation struct {
}

type HTTPRouteTimeoutParameters struct {

	// +kubebuilder:validation:Optional
	Idle []HTTPRouteTimeoutIdleParameters `json:"idle,omitempty" tf:"idle"`

	// +kubebuilder:validation:Optional
	PerRequest []HTTPRouteTimeoutPerRequestParameters `json:"perRequest,omitempty" tf:"per_request"`
}

type HTTPRouteTimeoutPerRequestObservation struct {
}

type HTTPRouteTimeoutPerRequestParameters struct {

	// +kubebuilder:validation:Required
	Unit string `json:"unit" tf:"unit"`

	// +kubebuilder:validation:Required
	Value int64 `json:"value" tf:"value"`
}

type HeaderMatchObservation struct {
}

type HeaderMatchParameters struct {

	// +kubebuilder:validation:Optional
	Exact *string `json:"exact,omitempty" tf:"exact"`

	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix"`

	// +kubebuilder:validation:Optional
	Range []MatchRangeParameters `json:"range,omitempty" tf:"range"`

	// +kubebuilder:validation:Optional
	Regex *string `json:"regex,omitempty" tf:"regex"`

	// +kubebuilder:validation:Optional
	Suffix *string `json:"suffix,omitempty" tf:"suffix"`
}

type HeaderMatchRangeObservation struct {
}

type HeaderMatchRangeParameters struct {

	// +kubebuilder:validation:Required
	End int64 `json:"end" tf:"end"`

	// +kubebuilder:validation:Required
	Start int64 `json:"start" tf:"start"`
}

type HeaderObservation struct {
}

type HeaderParameters struct {

	// +kubebuilder:validation:Optional
	Invert *bool `json:"invert,omitempty" tf:"invert"`

	// +kubebuilder:validation:Optional
	Match []HeaderMatchParameters `json:"match,omitempty" tf:"match"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`
}

type Http2RouteRetryPolicyObservation struct {
}

type Http2RouteRetryPolicyParameters struct {

	// +kubebuilder:validation:Optional
	HTTPRetryEvents []string `json:"httpRetryEvents,omitempty" tf:"http_retry_events"`

	// +kubebuilder:validation:Required
	MaxRetries int64 `json:"maxRetries" tf:"max_retries"`

	// +kubebuilder:validation:Required
	PerRetryTimeout []RetryPolicyPerRetryTimeoutParameters `json:"perRetryTimeout" tf:"per_retry_timeout"`

	// +kubebuilder:validation:Optional
	TCPRetryEvents []string `json:"tcpRetryEvents,omitempty" tf:"tcp_retry_events"`
}

type Http2RouteTimeoutObservation struct {
}

type Http2RouteTimeoutParameters struct {

	// +kubebuilder:validation:Optional
	Idle []TimeoutIdleParameters `json:"idle,omitempty" tf:"idle"`

	// +kubebuilder:validation:Optional
	PerRequest []TimeoutPerRequestParameters `json:"perRequest,omitempty" tf:"per_request"`
}

type IdleObservation struct {
}

type IdleParameters struct {

	// +kubebuilder:validation:Required
	Unit string `json:"unit" tf:"unit"`

	// +kubebuilder:validation:Required
	Value int64 `json:"value" tf:"value"`
}

type MatchHeaderMatchObservation struct {
}

type MatchHeaderMatchParameters struct {

	// +kubebuilder:validation:Optional
	Exact *string `json:"exact,omitempty" tf:"exact"`

	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix"`

	// +kubebuilder:validation:Optional
	Range []HeaderMatchRangeParameters `json:"range,omitempty" tf:"range"`

	// +kubebuilder:validation:Optional
	Regex *string `json:"regex,omitempty" tf:"regex"`

	// +kubebuilder:validation:Optional
	Suffix *string `json:"suffix,omitempty" tf:"suffix"`
}

type MatchHeaderObservation struct {
}

type MatchHeaderParameters struct {

	// +kubebuilder:validation:Optional
	Invert *bool `json:"invert,omitempty" tf:"invert"`

	// +kubebuilder:validation:Optional
	Match []MatchHeaderMatchParameters `json:"match,omitempty" tf:"match"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`
}

type MatchRangeObservation struct {
}

type MatchRangeParameters struct {

	// +kubebuilder:validation:Required
	End int64 `json:"end" tf:"end"`

	// +kubebuilder:validation:Required
	Start int64 `json:"start" tf:"start"`
}

type MetadataMatchObservation struct {
}

type MetadataMatchParameters struct {

	// +kubebuilder:validation:Optional
	Exact *string `json:"exact,omitempty" tf:"exact"`

	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix"`

	// +kubebuilder:validation:Optional
	Range []RangeParameters `json:"range,omitempty" tf:"range"`

	// +kubebuilder:validation:Optional
	Regex *string `json:"regex,omitempty" tf:"regex"`

	// +kubebuilder:validation:Optional
	Suffix *string `json:"suffix,omitempty" tf:"suffix"`
}

type MetadataObservation struct {
}

type MetadataParameters struct {

	// +kubebuilder:validation:Optional
	Invert *bool `json:"invert,omitempty" tf:"invert"`

	// +kubebuilder:validation:Optional
	Match []MetadataMatchParameters `json:"match,omitempty" tf:"match"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`
}

type PerRequestObservation struct {
}

type PerRequestParameters struct {

	// +kubebuilder:validation:Required
	Unit string `json:"unit" tf:"unit"`

	// +kubebuilder:validation:Required
	Value int64 `json:"value" tf:"value"`
}

type PerRetryTimeoutObservation struct {
}

type PerRetryTimeoutParameters struct {

	// +kubebuilder:validation:Required
	Unit string `json:"unit" tf:"unit"`

	// +kubebuilder:validation:Required
	Value int64 `json:"value" tf:"value"`
}

type RangeObservation struct {
}

type RangeParameters struct {

	// +kubebuilder:validation:Required
	End int64 `json:"end" tf:"end"`

	// +kubebuilder:validation:Required
	Start int64 `json:"start" tf:"start"`
}

type RetryPolicyObservation struct {
}

type RetryPolicyParameters struct {

	// +kubebuilder:validation:Optional
	GrpcRetryEvents []string `json:"grpcRetryEvents,omitempty" tf:"grpc_retry_events"`

	// +kubebuilder:validation:Optional
	HTTPRetryEvents []string `json:"httpRetryEvents,omitempty" tf:"http_retry_events"`

	// +kubebuilder:validation:Required
	MaxRetries int64 `json:"maxRetries" tf:"max_retries"`

	// +kubebuilder:validation:Required
	PerRetryTimeout []PerRetryTimeoutParameters `json:"perRetryTimeout" tf:"per_retry_timeout"`

	// +kubebuilder:validation:Optional
	TCPRetryEvents []string `json:"tcpRetryEvents,omitempty" tf:"tcp_retry_events"`
}

type RetryPolicyPerRetryTimeoutObservation struct {
}

type RetryPolicyPerRetryTimeoutParameters struct {

	// +kubebuilder:validation:Required
	Unit string `json:"unit" tf:"unit"`

	// +kubebuilder:validation:Required
	Value int64 `json:"value" tf:"value"`
}

type SpecGrpcRouteObservation struct {
}

type SpecGrpcRouteParameters struct {

	// +kubebuilder:validation:Required
	Action []GrpcRouteActionParameters `json:"action" tf:"action"`

	// +kubebuilder:validation:Optional
	Match []GrpcRouteMatchParameters `json:"match,omitempty" tf:"match"`

	// +kubebuilder:validation:Optional
	RetryPolicy []RetryPolicyParameters `json:"retryPolicy,omitempty" tf:"retry_policy"`

	// +kubebuilder:validation:Optional
	Timeout []TimeoutParameters `json:"timeout,omitempty" tf:"timeout"`
}

type SpecHTTPRouteActionObservation struct {
}

type SpecHTTPRouteActionParameters struct {

	// +kubebuilder:validation:Required
	WeightedTarget []HTTPRouteActionWeightedTargetParameters `json:"weightedTarget" tf:"weighted_target"`
}

type SpecHTTPRouteMatchObservation struct {
}

type SpecHTTPRouteMatchParameters struct {

	// +kubebuilder:validation:Optional
	Header []MatchHeaderParameters `json:"header,omitempty" tf:"header"`

	// +kubebuilder:validation:Optional
	Method *string `json:"method,omitempty" tf:"method"`

	// +kubebuilder:validation:Required
	Prefix string `json:"prefix" tf:"prefix"`

	// +kubebuilder:validation:Optional
	Scheme *string `json:"scheme,omitempty" tf:"scheme"`
}

type SpecHTTPRouteObservation struct {
}

type SpecHTTPRouteParameters struct {

	// +kubebuilder:validation:Required
	Action []SpecHTTPRouteActionParameters `json:"action" tf:"action"`

	// +kubebuilder:validation:Required
	Match []SpecHTTPRouteMatchParameters `json:"match" tf:"match"`

	// +kubebuilder:validation:Optional
	RetryPolicy []HTTPRouteRetryPolicyParameters `json:"retryPolicy,omitempty" tf:"retry_policy"`

	// +kubebuilder:validation:Optional
	Timeout []HTTPRouteTimeoutParameters `json:"timeout,omitempty" tf:"timeout"`
}

type SpecHttp2RouteActionObservation struct {
}

type SpecHttp2RouteActionParameters struct {

	// +kubebuilder:validation:Required
	WeightedTarget []ActionWeightedTargetParameters `json:"weightedTarget" tf:"weighted_target"`
}

type SpecHttp2RouteMatchObservation struct {
}

type SpecHttp2RouteMatchParameters struct {

	// +kubebuilder:validation:Optional
	Header []HeaderParameters `json:"header,omitempty" tf:"header"`

	// +kubebuilder:validation:Optional
	Method *string `json:"method,omitempty" tf:"method"`

	// +kubebuilder:validation:Required
	Prefix string `json:"prefix" tf:"prefix"`

	// +kubebuilder:validation:Optional
	Scheme *string `json:"scheme,omitempty" tf:"scheme"`
}

type SpecHttp2RouteObservation struct {
}

type SpecHttp2RouteParameters struct {

	// +kubebuilder:validation:Required
	Action []SpecHttp2RouteActionParameters `json:"action" tf:"action"`

	// +kubebuilder:validation:Required
	Match []SpecHttp2RouteMatchParameters `json:"match" tf:"match"`

	// +kubebuilder:validation:Optional
	RetryPolicy []Http2RouteRetryPolicyParameters `json:"retryPolicy,omitempty" tf:"retry_policy"`

	// +kubebuilder:validation:Optional
	Timeout []Http2RouteTimeoutParameters `json:"timeout,omitempty" tf:"timeout"`
}

type TCPRouteActionObservation struct {
}

type TCPRouteActionParameters struct {

	// +kubebuilder:validation:Required
	WeightedTarget []TCPRouteActionWeightedTargetParameters `json:"weightedTarget" tf:"weighted_target"`
}

type TCPRouteActionWeightedTargetObservation struct {
}

type TCPRouteActionWeightedTargetParameters struct {

	// +kubebuilder:validation:Required
	VirtualNode string `json:"virtualNode" tf:"virtual_node"`

	// +kubebuilder:validation:Required
	Weight int64 `json:"weight" tf:"weight"`
}

type TCPRouteObservation struct {
}

type TCPRouteParameters struct {

	// +kubebuilder:validation:Required
	Action []TCPRouteActionParameters `json:"action" tf:"action"`

	// +kubebuilder:validation:Optional
	Timeout []TCPRouteTimeoutParameters `json:"timeout,omitempty" tf:"timeout"`
}

type TCPRouteTimeoutIdleObservation struct {
}

type TCPRouteTimeoutIdleParameters struct {

	// +kubebuilder:validation:Required
	Unit string `json:"unit" tf:"unit"`

	// +kubebuilder:validation:Required
	Value int64 `json:"value" tf:"value"`
}

type TCPRouteTimeoutObservation struct {
}

type TCPRouteTimeoutParameters struct {

	// +kubebuilder:validation:Optional
	Idle []TCPRouteTimeoutIdleParameters `json:"idle,omitempty" tf:"idle"`
}

type TimeoutIdleObservation struct {
}

type TimeoutIdleParameters struct {

	// +kubebuilder:validation:Required
	Unit string `json:"unit" tf:"unit"`

	// +kubebuilder:validation:Required
	Value int64 `json:"value" tf:"value"`
}

type TimeoutObservation struct {
}

type TimeoutParameters struct {

	// +kubebuilder:validation:Optional
	Idle []IdleParameters `json:"idle,omitempty" tf:"idle"`

	// +kubebuilder:validation:Optional
	PerRequest []PerRequestParameters `json:"perRequest,omitempty" tf:"per_request"`
}

type TimeoutPerRequestObservation struct {
}

type TimeoutPerRequestParameters struct {

	// +kubebuilder:validation:Required
	Unit string `json:"unit" tf:"unit"`

	// +kubebuilder:validation:Required
	Value int64 `json:"value" tf:"value"`
}

type WeightedTargetObservation struct {
}

type WeightedTargetParameters struct {

	// +kubebuilder:validation:Required
	VirtualNode string `json:"virtualNode" tf:"virtual_node"`

	// +kubebuilder:validation:Required
	Weight int64 `json:"weight" tf:"weight"`
}

// AppmeshRouteSpec defines the desired state of AppmeshRoute
type AppmeshRouteSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       AppmeshRouteParameters `json:"forProvider"`
}

// AppmeshRouteStatus defines the observed state of AppmeshRoute.
type AppmeshRouteStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          AppmeshRouteObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppmeshRoute is the Schema for the AppmeshRoutes API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type AppmeshRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppmeshRouteSpec   `json:"spec"`
	Status            AppmeshRouteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppmeshRouteList contains a list of AppmeshRoutes
type AppmeshRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppmeshRoute `json:"items"`
}

// Repository type metadata.
var (
	AppmeshRouteKind             = "AppmeshRoute"
	AppmeshRouteGroupKind        = schema.GroupKind{Group: Group, Kind: AppmeshRouteKind}.String()
	AppmeshRouteKindAPIVersion   = AppmeshRouteKind + "." + GroupVersion.String()
	AppmeshRouteGroupVersionKind = GroupVersion.WithKind(AppmeshRouteKind)
)

func init() {
	SchemeBuilder.Register(&AppmeshRoute{}, &AppmeshRouteList{})
}
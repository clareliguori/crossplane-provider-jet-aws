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

type AlbListenerObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type AlbListenerParameters struct {

	// +kubebuilder:validation:Optional
	AlpnPolicy *string `json:"alpnPolicy,omitempty" tf:"alpn_policy"`

	// +kubebuilder:validation:Optional
	CertificateArn *string `json:"certificateArn,omitempty" tf:"certificate_arn"`

	// +kubebuilder:validation:Required
	DefaultAction []DefaultActionParameters `json:"defaultAction" tf:"default_action"`

	// +kubebuilder:validation:Required
	LoadBalancerArn string `json:"loadBalancerArn" tf:"load_balancer_arn"`

	// +kubebuilder:validation:Optional
	Port *int64 `json:"port,omitempty" tf:"port"`

	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SslPolicy *string `json:"sslPolicy,omitempty" tf:"ssl_policy"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type AuthenticateCognitoObservation struct {
}

type AuthenticateCognitoParameters struct {

	// +kubebuilder:validation:Optional
	AuthenticationRequestExtraParams map[string]string `json:"authenticationRequestExtraParams,omitempty" tf:"authentication_request_extra_params"`

	// +kubebuilder:validation:Optional
	OnUnauthenticatedRequest *string `json:"onUnauthenticatedRequest,omitempty" tf:"on_unauthenticated_request"`

	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope"`

	// +kubebuilder:validation:Optional
	SessionCookieName *string `json:"sessionCookieName,omitempty" tf:"session_cookie_name"`

	// +kubebuilder:validation:Optional
	SessionTimeout *int64 `json:"sessionTimeout,omitempty" tf:"session_timeout"`

	// +kubebuilder:validation:Required
	UserPoolArn string `json:"userPoolArn" tf:"user_pool_arn"`

	// +kubebuilder:validation:Required
	UserPoolClientID string `json:"userPoolClientId" tf:"user_pool_client_id"`

	// +kubebuilder:validation:Required
	UserPoolDomain string `json:"userPoolDomain" tf:"user_pool_domain"`
}

type AuthenticateOidcObservation struct {
}

type AuthenticateOidcParameters struct {

	// +kubebuilder:validation:Optional
	AuthenticationRequestExtraParams map[string]string `json:"authenticationRequestExtraParams,omitempty" tf:"authentication_request_extra_params"`

	// +kubebuilder:validation:Required
	AuthorizationEndpoint string `json:"authorizationEndpoint" tf:"authorization_endpoint"`

	// +kubebuilder:validation:Required
	ClientID string `json:"clientId" tf:"client_id"`

	// +kubebuilder:validation:Required
	ClientSecret string `json:"clientSecret" tf:"client_secret"`

	// +kubebuilder:validation:Required
	Issuer string `json:"issuer" tf:"issuer"`

	// +kubebuilder:validation:Optional
	OnUnauthenticatedRequest *string `json:"onUnauthenticatedRequest,omitempty" tf:"on_unauthenticated_request"`

	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope"`

	// +kubebuilder:validation:Optional
	SessionCookieName *string `json:"sessionCookieName,omitempty" tf:"session_cookie_name"`

	// +kubebuilder:validation:Optional
	SessionTimeout *int64 `json:"sessionTimeout,omitempty" tf:"session_timeout"`

	// +kubebuilder:validation:Required
	TokenEndpoint string `json:"tokenEndpoint" tf:"token_endpoint"`

	// +kubebuilder:validation:Required
	UserInfoEndpoint string `json:"userInfoEndpoint" tf:"user_info_endpoint"`
}

type DefaultActionObservation struct {
}

type DefaultActionParameters struct {

	// +kubebuilder:validation:Optional
	AuthenticateCognito []AuthenticateCognitoParameters `json:"authenticateCognito,omitempty" tf:"authenticate_cognito"`

	// +kubebuilder:validation:Optional
	AuthenticateOidc []AuthenticateOidcParameters `json:"authenticateOidc,omitempty" tf:"authenticate_oidc"`

	// +kubebuilder:validation:Optional
	FixedResponse []FixedResponseParameters `json:"fixedResponse,omitempty" tf:"fixed_response"`

	// +kubebuilder:validation:Optional
	Forward []ForwardParameters `json:"forward,omitempty" tf:"forward"`

	// +kubebuilder:validation:Optional
	Order *int64 `json:"order,omitempty" tf:"order"`

	// +kubebuilder:validation:Optional
	Redirect []RedirectParameters `json:"redirect,omitempty" tf:"redirect"`

	// +kubebuilder:validation:Optional
	TargetGroupArn *string `json:"targetGroupArn,omitempty" tf:"target_group_arn"`

	// +kubebuilder:validation:Required
	Type string `json:"type" tf:"type"`
}

type FixedResponseObservation struct {
}

type FixedResponseParameters struct {

	// +kubebuilder:validation:Required
	ContentType string `json:"contentType" tf:"content_type"`

	// +kubebuilder:validation:Optional
	MessageBody *string `json:"messageBody,omitempty" tf:"message_body"`

	// +kubebuilder:validation:Optional
	StatusCode *string `json:"statusCode,omitempty" tf:"status_code"`
}

type ForwardObservation struct {
}

type ForwardParameters struct {

	// +kubebuilder:validation:Optional
	Stickiness []StickinessParameters `json:"stickiness,omitempty" tf:"stickiness"`

	// +kubebuilder:validation:Required
	TargetGroup []TargetGroupParameters `json:"targetGroup" tf:"target_group"`
}

type RedirectObservation struct {
}

type RedirectParameters struct {

	// +kubebuilder:validation:Optional
	Host *string `json:"host,omitempty" tf:"host"`

	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path"`

	// +kubebuilder:validation:Optional
	Port *string `json:"port,omitempty" tf:"port"`

	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol"`

	// +kubebuilder:validation:Optional
	Query *string `json:"query,omitempty" tf:"query"`

	// +kubebuilder:validation:Required
	StatusCode string `json:"statusCode" tf:"status_code"`
}

type StickinessObservation struct {
}

type StickinessParameters struct {

	// +kubebuilder:validation:Required
	Duration int64 `json:"duration" tf:"duration"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
}

type TargetGroupObservation struct {
}

type TargetGroupParameters struct {

	// +kubebuilder:validation:Required
	Arn string `json:"arn" tf:"arn"`

	// +kubebuilder:validation:Optional
	Weight *int64 `json:"weight,omitempty" tf:"weight"`
}

// AlbListenerSpec defines the desired state of AlbListener
type AlbListenerSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       AlbListenerParameters `json:"forProvider"`
}

// AlbListenerStatus defines the observed state of AlbListener.
type AlbListenerStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          AlbListenerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AlbListener is the Schema for the AlbListeners API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type AlbListener struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AlbListenerSpec   `json:"spec"`
	Status            AlbListenerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AlbListenerList contains a list of AlbListeners
type AlbListenerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AlbListener `json:"items"`
}

// Repository type metadata.
var (
	AlbListenerKind             = "AlbListener"
	AlbListenerGroupKind        = schema.GroupKind{Group: Group, Kind: AlbListenerKind}.String()
	AlbListenerKindAPIVersion   = AlbListenerKind + "." + GroupVersion.String()
	AlbListenerGroupVersionKind = GroupVersion.WithKind(AlbListenerKind)
)

func init() {
	SchemeBuilder.Register(&AlbListener{}, &AlbListenerList{})
}
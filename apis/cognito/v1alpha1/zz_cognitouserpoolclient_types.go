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

type AnalyticsConfigurationObservation struct {
}

type AnalyticsConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	ApplicationArn *string `json:"applicationArn,omitempty" tf:"application_arn"`

	// +kubebuilder:validation:Optional
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id"`

	// +kubebuilder:validation:Optional
	ExternalID *string `json:"externalId,omitempty" tf:"external_id"`

	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn"`

	// +kubebuilder:validation:Optional
	UserDataShared *bool `json:"userDataShared,omitempty" tf:"user_data_shared"`
}

type CognitoUserPoolClientObservation struct {
	ClientSecret string `json:"clientSecret" tf:"client_secret"`
}

type CognitoUserPoolClientParameters struct {

	// +kubebuilder:validation:Optional
	AccessTokenValidity *int64 `json:"accessTokenValidity,omitempty" tf:"access_token_validity"`

	// +kubebuilder:validation:Optional
	AllowedOauthFlows []string `json:"allowedOauthFlows,omitempty" tf:"allowed_oauth_flows"`

	// +kubebuilder:validation:Optional
	AllowedOauthFlowsUserPoolClient *bool `json:"allowedOauthFlowsUserPoolClient,omitempty" tf:"allowed_oauth_flows_user_pool_client"`

	// +kubebuilder:validation:Optional
	AllowedOauthScopes []string `json:"allowedOauthScopes,omitempty" tf:"allowed_oauth_scopes"`

	// +kubebuilder:validation:Optional
	AnalyticsConfiguration []AnalyticsConfigurationParameters `json:"analyticsConfiguration,omitempty" tf:"analytics_configuration"`

	// +kubebuilder:validation:Optional
	CallbackUrls []string `json:"callbackUrls,omitempty" tf:"callback_urls"`

	// +kubebuilder:validation:Optional
	DefaultRedirectURI *string `json:"defaultRedirectUri,omitempty" tf:"default_redirect_uri"`

	// +kubebuilder:validation:Optional
	EnableTokenRevocation *bool `json:"enableTokenRevocation,omitempty" tf:"enable_token_revocation"`

	// +kubebuilder:validation:Optional
	ExplicitAuthFlows []string `json:"explicitAuthFlows,omitempty" tf:"explicit_auth_flows"`

	// +kubebuilder:validation:Optional
	GenerateSecret *bool `json:"generateSecret,omitempty" tf:"generate_secret"`

	// +kubebuilder:validation:Optional
	IDTokenValidity *int64 `json:"idTokenValidity,omitempty" tf:"id_token_validity"`

	// +kubebuilder:validation:Optional
	LogoutUrls []string `json:"logoutUrls,omitempty" tf:"logout_urls"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Optional
	PreventUserExistenceErrors *string `json:"preventUserExistenceErrors,omitempty" tf:"prevent_user_existence_errors"`

	// +kubebuilder:validation:Optional
	ReadAttributes []string `json:"readAttributes,omitempty" tf:"read_attributes"`

	// +kubebuilder:validation:Optional
	RefreshTokenValidity *int64 `json:"refreshTokenValidity,omitempty" tf:"refresh_token_validity"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SupportedIdentityProviders []string `json:"supportedIdentityProviders,omitempty" tf:"supported_identity_providers"`

	// +kubebuilder:validation:Optional
	TokenValidityUnits []TokenValidityUnitsParameters `json:"tokenValidityUnits,omitempty" tf:"token_validity_units"`

	// +kubebuilder:validation:Required
	UserPoolID string `json:"userPoolId" tf:"user_pool_id"`

	// +kubebuilder:validation:Optional
	WriteAttributes []string `json:"writeAttributes,omitempty" tf:"write_attributes"`
}

type TokenValidityUnitsObservation struct {
}

type TokenValidityUnitsParameters struct {

	// +kubebuilder:validation:Optional
	AccessToken *string `json:"accessToken,omitempty" tf:"access_token"`

	// +kubebuilder:validation:Optional
	IDToken *string `json:"idToken,omitempty" tf:"id_token"`

	// +kubebuilder:validation:Optional
	RefreshToken *string `json:"refreshToken,omitempty" tf:"refresh_token"`
}

// CognitoUserPoolClientSpec defines the desired state of CognitoUserPoolClient
type CognitoUserPoolClientSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CognitoUserPoolClientParameters `json:"forProvider"`
}

// CognitoUserPoolClientStatus defines the observed state of CognitoUserPoolClient.
type CognitoUserPoolClientStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CognitoUserPoolClientObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CognitoUserPoolClient is the Schema for the CognitoUserPoolClients API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type CognitoUserPoolClient struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CognitoUserPoolClientSpec   `json:"spec"`
	Status            CognitoUserPoolClientStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CognitoUserPoolClientList contains a list of CognitoUserPoolClients
type CognitoUserPoolClientList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CognitoUserPoolClient `json:"items"`
}

// Repository type metadata.
var (
	CognitoUserPoolClientKind             = "CognitoUserPoolClient"
	CognitoUserPoolClientGroupKind        = schema.GroupKind{Group: Group, Kind: CognitoUserPoolClientKind}.String()
	CognitoUserPoolClientKindAPIVersion   = CognitoUserPoolClientKind + "." + GroupVersion.String()
	CognitoUserPoolClientGroupVersionKind = GroupVersion.WithKind(CognitoUserPoolClientKind)
)

func init() {
	SchemeBuilder.Register(&CognitoUserPoolClient{}, &CognitoUserPoolClientList{})
}
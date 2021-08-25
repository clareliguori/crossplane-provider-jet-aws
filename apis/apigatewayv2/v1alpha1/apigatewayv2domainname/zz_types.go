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

type Apigatewayv2DomainNameObservation struct {
	ApiMappingSelectionExpression string `json:"apiMappingSelectionExpression" tf:"api_mapping_selection_expression"`

	Arn string `json:"arn" tf:"arn"`
}

type Apigatewayv2DomainNameParameters struct {
	DomainName string `json:"domainName" tf:"domain_name"`

	DomainNameConfiguration []DomainNameConfigurationParameters `json:"domainNameConfiguration" tf:"domain_name_configuration"`

	MutualTlsAuthentication []MutualTlsAuthenticationParameters `json:"mutualTlsAuthentication,omitempty" tf:"mutual_tls_authentication"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type DomainNameConfigurationObservation struct {
	HostedZoneId string `json:"hostedZoneId" tf:"hosted_zone_id"`

	TargetDomainName string `json:"targetDomainName" tf:"target_domain_name"`
}

type DomainNameConfigurationParameters struct {
	CertificateArn string `json:"certificateArn" tf:"certificate_arn"`

	EndpointType string `json:"endpointType" tf:"endpoint_type"`

	SecurityPolicy string `json:"securityPolicy" tf:"security_policy"`
}

type MutualTlsAuthenticationObservation struct {
}

type MutualTlsAuthenticationParameters struct {
	TruststoreUri string `json:"truststoreUri" tf:"truststore_uri"`

	TruststoreVersion *string `json:"truststoreVersion,omitempty" tf:"truststore_version"`
}

// Apigatewayv2DomainNameSpec defines the desired state of Apigatewayv2DomainName
type Apigatewayv2DomainNameSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Apigatewayv2DomainNameParameters `json:"forProvider"`
}

// Apigatewayv2DomainNameStatus defines the observed state of Apigatewayv2DomainName.
type Apigatewayv2DomainNameStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Apigatewayv2DomainNameObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Apigatewayv2DomainName is the Schema for the Apigatewayv2DomainNames API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Apigatewayv2DomainName struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Apigatewayv2DomainNameSpec   `json:"spec"`
	Status            Apigatewayv2DomainNameStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Apigatewayv2DomainNameList contains a list of Apigatewayv2DomainNames
type Apigatewayv2DomainNameList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Apigatewayv2DomainName `json:"items"`
}

// Repository type metadata.
var (
	Apigatewayv2DomainNameKind             = "Apigatewayv2DomainName"
	Apigatewayv2DomainNameGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: Apigatewayv2DomainNameKind}.String()
	Apigatewayv2DomainNameKindAPIVersion   = Apigatewayv2DomainNameKind + "." + v1alpha1.GroupVersion.String()
	Apigatewayv2DomainNameGroupVersionKind = v1alpha1.GroupVersion.WithKind(Apigatewayv2DomainNameKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &Apigatewayv2DomainName{}, &Apigatewayv2DomainNameList{})
}

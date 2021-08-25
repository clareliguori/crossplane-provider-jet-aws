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

type ApiGatewayClientCertificateObservation struct {
	Arn string `json:"arn" tf:"arn"`

	CreatedDate string `json:"createdDate" tf:"created_date"`

	ExpirationDate string `json:"expirationDate" tf:"expiration_date"`

	PemEncodedCertificate string `json:"pemEncodedCertificate" tf:"pem_encoded_certificate"`
}

type ApiGatewayClientCertificateParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// ApiGatewayClientCertificateSpec defines the desired state of ApiGatewayClientCertificate
type ApiGatewayClientCertificateSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ApiGatewayClientCertificateParameters `json:"forProvider"`
}

// ApiGatewayClientCertificateStatus defines the observed state of ApiGatewayClientCertificate.
type ApiGatewayClientCertificateStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ApiGatewayClientCertificateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayClientCertificate is the Schema for the ApiGatewayClientCertificates API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ApiGatewayClientCertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayClientCertificateSpec   `json:"spec"`
	Status            ApiGatewayClientCertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayClientCertificateList contains a list of ApiGatewayClientCertificates
type ApiGatewayClientCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiGatewayClientCertificate `json:"items"`
}

// Repository type metadata.
var (
	ApiGatewayClientCertificateKind             = "ApiGatewayClientCertificate"
	ApiGatewayClientCertificateGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: ApiGatewayClientCertificateKind}.String()
	ApiGatewayClientCertificateKindAPIVersion   = ApiGatewayClientCertificateKind + "." + v1alpha1.GroupVersion.String()
	ApiGatewayClientCertificateGroupVersionKind = v1alpha1.GroupVersion.WithKind(ApiGatewayClientCertificateKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &ApiGatewayClientCertificate{}, &ApiGatewayClientCertificateList{})
}

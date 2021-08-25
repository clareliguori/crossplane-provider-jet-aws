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
// +groupName=dx.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	apis "github.com/crossplane-contrib/provider-tf-aws/apis"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/dx/v1alpha1"
)

type DxHostedPublicVirtualInterfaceObservation struct {
	AmazonSideAsn string `json:"amazonSideAsn" tf:"amazon_side_asn"`

	Arn string `json:"arn" tf:"arn"`

	AwsDevice string `json:"awsDevice" tf:"aws_device"`
}

type DxHostedPublicVirtualInterfaceParameters struct {
	AddressFamily string `json:"addressFamily" tf:"address_family"`

	AmazonAddress *string `json:"amazonAddress,omitempty" tf:"amazon_address"`

	BgpAsn int64 `json:"bgpAsn" tf:"bgp_asn"`

	BgpAuthKey *string `json:"bgpAuthKey,omitempty" tf:"bgp_auth_key"`

	ConnectionId string `json:"connectionId" tf:"connection_id"`

	CustomerAddress *string `json:"customerAddress,omitempty" tf:"customer_address"`

	Name string `json:"name" tf:"name"`

	OwnerAccountId string `json:"ownerAccountId" tf:"owner_account_id"`

	RouteFilterPrefixes []string `json:"routeFilterPrefixes" tf:"route_filter_prefixes"`

	Vlan int64 `json:"vlan" tf:"vlan"`
}

// DxHostedPublicVirtualInterfaceSpec defines the desired state of DxHostedPublicVirtualInterface
type DxHostedPublicVirtualInterfaceSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DxHostedPublicVirtualInterfaceParameters `json:"forProvider"`
}

// DxHostedPublicVirtualInterfaceStatus defines the observed state of DxHostedPublicVirtualInterface.
type DxHostedPublicVirtualInterfaceStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DxHostedPublicVirtualInterfaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DxHostedPublicVirtualInterface is the Schema for the DxHostedPublicVirtualInterfaces API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DxHostedPublicVirtualInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DxHostedPublicVirtualInterfaceSpec   `json:"spec"`
	Status            DxHostedPublicVirtualInterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DxHostedPublicVirtualInterfaceList contains a list of DxHostedPublicVirtualInterfaces
type DxHostedPublicVirtualInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DxHostedPublicVirtualInterface `json:"items"`
}

// Repository type metadata.
var (
	DxHostedPublicVirtualInterfaceKind             = "DxHostedPublicVirtualInterface"
	DxHostedPublicVirtualInterfaceGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: DxHostedPublicVirtualInterfaceKind}.String()
	DxHostedPublicVirtualInterfaceKindAPIVersion   = DxHostedPublicVirtualInterfaceKind + "." + v1alpha1.GroupVersion.String()
	DxHostedPublicVirtualInterfaceGroupVersionKind = v1alpha1.GroupVersion.WithKind(DxHostedPublicVirtualInterfaceKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &DxHostedPublicVirtualInterface{}, &DxHostedPublicVirtualInterfaceList{})
}
